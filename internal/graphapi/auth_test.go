package graphapi_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.infratographer.com/permissions-api/pkg/permissions"
	"go.infratographer.com/x/echojwtx"
	"go.infratographer.com/x/gidx"
	"go.infratographer.com/x/testing/auth"
)

func TestGetIPBlockWithAuth(t *testing.T) {
	oauthCLI, issuer, oAuthClose := auth.OAuthTestClient("urn:test:loadbalancer", "")
	defer oAuthClose()

	srv, err := newTestServer(
		withAuthConfig(
			&echojwtx.AuthConfig{
				Issuer: issuer,
			},
		),
		withPermissions(
			permissions.WithDefaultChecker(permissions.DefaultAllowChecker),
		),
	)

	require.NoError(t, err)
	require.NotNil(t, srv)

	defer srv.Close()

	ctx := context.Background()
	ipbt := (&IPBlockTypeBuilder{
		OwnerID: gidx.MustNewID("testown"),
	}).MustNew(ctx)

	resp, err := graphTestClient(
		withHTTPClient(oauthCLI),
		withSrvURL(srv.URL+"/query"),
	).GetIPBlockType(ctx, ipbt.ID)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, ipbt.ID, resp.IPBlockType.ID)
}

func TestGetLoadbalancerNoAuth(t *testing.T) {
	_, issuer, oAuthClose := auth.OAuthTestClient("urn:test:loadbalancer", "")
	defer oAuthClose()

	srv, err := newTestServer(
		withAuthConfig(
			&echojwtx.AuthConfig{
				Issuer: issuer,
			},
		),
		withPermissions(
			permissions.WithDefaultChecker(permissions.DefaultAllowChecker),
		),
	)

	require.NoError(t, err)
	require.NotNil(t, srv)

	defer srv.Close()

	ctx := context.Background()
	ipbt := (&IPBlockTypeBuilder{
		OwnerID: gidx.MustNewID("testown"),
	}).MustNew(ctx)

	resp, err := graphTestClient(
		withHTTPClient(http.DefaultClient),
		withSrvURL(srv.URL+"/query"),
	).GetIPBlockType(ctx, ipbt.ID)

	require.Error(t, err, "Expected an authorization error")
	require.Nil(t, resp)
	assert.ErrorContains(t, err, `{"networkErrors":{"code":401`)
}
