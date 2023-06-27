package ipamclient

import (
	"context"
	"net/http"

	// "github.com/3th1nk/cidr"

	"github.com/3th1nk/cidr"
	"github.com/shurcooL/graphql"
	"go.infratographer.com/x/gidx"
	"golang.org/x/exp/slices"
)

// GQLClient is an interface for a graphql client
type GQLClient interface {
	Query(ctx context.Context, q interface{}, variables map[string]interface{}) error
	Mutate(ctx context.Context, m interface{}, variables map[string]interface{}) error
}

// Client creates a new lb api client against a specific endpoint
type Client struct {
	gqlCli     GQLClient
	httpClient *http.Client
}

// ClientOption is a function that modifies a client
type ClientOption func(*Client)

// NewClient creates a new lb api client
func NewClient(url string, opts ...ClientOption) *Client {
	c := &Client{
		httpClient: http.DefaultClient,
	}

	for _, opt := range opts {
		opt(c)
	}

	c.gqlCli = graphql.NewClient(url, c.httpClient)

	return c
}

// WithHTTPClient functional option to set the http client
func WithHTTPClient(cli *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = cli
	}
}

// GetIPAddress returns an IP Address by id
func (c *Client) GetIPAddress(ctx context.Context, id string) (*GetIPAddress, error) {
	_, err := gidx.Parse(id)
	if err != nil {
		return nil, err
	}

	vars := map[string]interface{}{
		"id": id,
	}

	var ipa GetIPAddress
	if err := c.gqlCli.Query(ctx, &ipa, vars); err != nil {
		return nil, err
	}

	return &ipa, nil
}

// GetIPBlock returns an IP Block by id
func (c *Client) GetIPBlock(ctx context.Context, id string) (*GetIPBlock, error) {
	_, err := gidx.Parse(id)
	if err != nil {
		return nil, err
	}

	vars := map[string]interface{}{
		"id": id,
	}

	var ipb GetIPBlock
	if err := c.gqlCli.Query(ctx, &ipb, vars); err != nil {
		return nil, err
	}

	return &ipb, nil
}

// GetNextAvailableAddressFromBlock returns the next available address from a block
func (c *Client) GetNextAvailableAddressFromBlock(ctx context.Context, id string) (string, error) {
	var nxt string

	var err error

	block, err := c.GetIPBlock(ctx, id)
	if err != nil {
		return "", err
	}

	prefix := block.IPBlock.Prefix
	ips := []string{}

	for _, ip := range block.IPBlock.IPAddress.Edges {
		ips = append(ips, ip.Node.IP)
	}

	pf, err := cidr.Parse(prefix)
	if err != nil {
		return "", err
	}

	bgn, end := pf.IPRange()
	network := pf.Network().String()
	broadcast := pf.Broadcast().String()

	for {
		if !slices.Contains(ips, bgn.String()) && bgn.String() != network && bgn.String() != broadcast {
			nxt = bgn.String()
			break
		}

		cidr.IPIncr(bgn)

		if bgn.Equal(end) {
			err = ErrNoAvailableIPs
		}
	}

	return nxt, err
}

// CreateIPAddressFromBlock creates an IP Address from the next available address in a given block
func (c *Client) CreateIPAddressFromBlock(ctx context.Context, blockid string, nodeid string, nodeownerid string, reserve bool) (*CreateIPAddress, error) {
	_, err := gidx.Parse(blockid)
	if err != nil {
		return nil, err
	}

	_, err = gidx.Parse(nodeid)
	if err != nil {
		return nil, err
	}

	_, err = gidx.Parse(nodeownerid)
	if err != nil {
		return nil, err
	}

	ipz, err := c.GetNextAvailableAddressFromBlock(ctx, blockid)
	if err != nil {
		return nil, err
	}

	vars := map[string]interface{}{
		"input": CreateIPAddressInput{
			IP:          ipz,
			IPBlockID:   blockid,
			NodeID:      nodeid,
			NodeOwnerID: nodeownerid,
			Reserved:    reserve,
		},
	}

	var ipn CreateIPAddress
	if err := c.gqlCli.Mutate(ctx, &ipn, vars); err != nil {
		return nil, err
	}

	return &ipn, err
}

// DeleteIPAddress deletes an IP Address by id
func (c *Client) DeleteIPAddress(ctx context.Context, id string) (*DeleteIPAddress, error) {
	_, err := gidx.Parse(id)
	if err != nil {
		return nil, err
	}

	vars := map[string]interface{}{
		"id": id,
	}

	var ipd DeleteIPAddress
	// if err := c.gqlCli.Query(ctx, &ipb, vars); err != nil {
	// 	return nil, err
	// }
	if err := c.gqlCli.Mutate(ctx, &ipd, vars); err != nil {
		return nil, err
	}

	return &ipd, nil
}
