package graphapi_test

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"go.uber.org/zap"

	"go.infratographer.com/permissions-api/pkg/permissions"
	"go.infratographer.com/x/echojwtx"
	"go.infratographer.com/x/echox"
	"go.infratographer.com/x/events"
	"go.infratographer.com/x/goosex"
	"go.infratographer.com/x/testing/eventtools"

	"go.infratographer.com/x/testing/containersx"

	"go.infratographer.com/ipam-api/db"
	ent "go.infratographer.com/ipam-api/internal/ent/generated"
	"go.infratographer.com/ipam-api/internal/graphapi"
	"go.infratographer.com/ipam-api/internal/testclient"
)

const (
	ipBlockTypePrefix = "testipt"
	locationPrefix    = "testloc"
	ownerPrefix       = "testtnt"
	nodePrefix        = "testnod"
)

var (
	TestDBURI   = os.Getenv("IPAMAPI_TESTDB_URI")
	EntClient   *ent.Client
	DBContainer *containersx.DBContainer
)

func TestMain(m *testing.M) {
	// setup the database if needed
	setupDB()
	// run the tests
	code := m.Run()
	// teardown the database
	teardownDB()
	// return the test response code
	os.Exit(code)
}

func parseDBURI(ctx context.Context) (string, string, *containersx.DBContainer) {
	switch {
	// if you don't pass in a database we default to an in memory sqlite
	case TestDBURI == "":
		return dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1", nil
	case strings.HasPrefix(TestDBURI, "sqlite://"):
		return dialect.SQLite, strings.TrimPrefix(TestDBURI, "sqlite://"), nil
	case strings.HasPrefix(TestDBURI, "postgres://"), strings.HasPrefix(TestDBURI, "postgresql://"):
		return dialect.Postgres, TestDBURI, nil
	case strings.HasPrefix(TestDBURI, "docker://"):
		dbImage := strings.TrimPrefix(TestDBURI, "docker://")

		switch {
		case strings.HasPrefix(dbImage, "cockroach"), strings.HasPrefix(dbImage, "cockroachdb"), strings.HasPrefix(dbImage, "crdb"):
			cntr, err := containersx.NewCockroachDB(ctx, dbImage)
			errPanic("error starting db test container", err)

			return dialect.Postgres, cntr.URI, cntr
		case strings.HasPrefix(dbImage, "postgres"):
			cntr, err := containersx.NewPostgresDB(ctx, dbImage,
				postgres.WithInitScripts(filepath.Join("testdata", "postgres_init.sh")),
			)
			errPanic("error starting db test container", err)

			return dialect.Postgres, cntr.URI, cntr
		default:
			panic("invalid testcontainer URI, uri: " + TestDBURI)
		}

	default:
		panic("invalid DB URI, uri: " + TestDBURI)
	}
}

func setupDB() {
	// don't setup the datastore if we already have one
	if EntClient != nil {
		return
	}

	ctx := context.Background()

	dia, uri, cntr := parseDBURI(ctx)

	nats, err := eventtools.NewNatsServer()
	if err != nil {
		errPanic("failed to start nats server", err)
	}

	pub, err := events.NewPublisher(nats.PublisherConfig)
	if err != nil {
		errPanic("failed to create events publisher", err)
	}

	c, err := ent.Open(dia, uri, ent.Debug(), ent.EventsPublisher(pub))
	if err != nil {
		errPanic("failed terminating test db container after failing to connect to the db", cntr.Container.Terminate(ctx))
		errPanic("failed opening connection to database:", err)
	}

	switch dia {
	case dialect.SQLite:
		// Run automatic migrations for SQLite
		errPanic("failed creating db scema", c.Schema.Create(ctx))
	case dialect.Postgres:
		log.Println("Running database migrations")
		goosex.MigrateUp(uri, db.Migrations)
	}

	EntClient = c
}

func teardownDB() {
	ctx := context.Background()

	if EntClient != nil {
		errPanic("teardown failed to close database connection", EntClient.Close())
	}

	if DBContainer != nil {
		errPanic("teardown failed to terminate test db container", DBContainer.Container.Terminate(ctx))
	}
}

func errPanic(msg string, err error) {
	if err != nil {
		log.Panicf("%s err: %s", msg, err.Error())
	}
}

type graphClient struct {
	srvURL     string
	httpClient *http.Client
}

type graphClientOptions func(*graphClient)

func withSrvURL(url string) graphClientOptions {
	return func(g *graphClient) {
		g.srvURL = url
	}
}

func withHTTPClient(httpcli *http.Client) graphClientOptions {
	return func(g *graphClient) {
		g.httpClient = httpcli
	}
}

func graphTestClient(options ...graphClientOptions) testclient.TestClient {
	g := &graphClient{
		srvURL: "graph",
		httpClient: &http.Client{Transport: localRoundTripper{handler: handler.NewDefaultServer(
			graphapi.NewExecutableSchema(
				graphapi.Config{Resolvers: graphapi.NewResolver(EntClient, zap.NewNop().Sugar())},
			))}},
	}

	for _, opt := range options {
		opt(g)
	}

	return testclient.NewClient(g.httpClient, g.srvURL)
}

// localRoundTripper is an http.RoundTripper that executes HTTP transactions
// by using handler directly, instead of going over an HTTP connection.
type localRoundTripper struct {
	handler http.Handler
}

func (l localRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	w := httptest.NewRecorder()
	l.handler.ServeHTTP(w, req)

	return w.Result(), nil
}

type testServerConfig struct {
	echoConfig        echox.Config
	handlerMiddleware []echo.MiddlewareFunc
}

type testServerOption func(*testServerConfig) error

func withAuthConfig(authConfig *echojwtx.AuthConfig) testServerOption {
	return func(tsc *testServerConfig) error {
		auth, err := echojwtx.NewAuth(context.Background(), *authConfig)
		if err != nil {
			return err
		}

		tsc.echoConfig = tsc.echoConfig.WithMiddleware(auth.Middleware())

		return nil
	}
}

func withPermissions(options ...permissions.Option) testServerOption {
	return func(tsc *testServerConfig) error {
		perms, err := permissions.New(permissions.Config{}, options...)
		if err != nil {
			return err
		}

		tsc.handlerMiddleware = append(tsc.handlerMiddleware, perms.Middleware())

		return nil
	}
}

func newTestServer(options ...testServerOption) (*httptest.Server, error) {
	tsc := new(testServerConfig)

	for _, opt := range options {
		if err := opt(tsc); err != nil {
			return nil, err
		}
	}

	srv, err := echox.NewServer(zap.NewNop(), tsc.echoConfig, nil)
	if err != nil {
		return nil, err
	}

	r := graphapi.NewResolver(EntClient, zap.NewNop().Sugar())
	srv.AddHandler(r.Handler(false, tsc.handlerMiddleware...))

	return httptest.NewServer(srv.Handler()), nil
}

func newString(s string) *string {
	return &s
}

func newBool(b bool) *bool {
	return &b
}
