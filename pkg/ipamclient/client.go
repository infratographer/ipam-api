package ipamclient

import (
	"context"
	"net/http"

	"github.com/3th1nk/cidr"
	graphql "github.com/hasura/go-graphql-client"
	"go.infratographer.com/x/gidx"
	"golang.org/x/exp/slices"
)

// GQLClient is an interface for a graphql client
type GQLClient interface {
	Query(ctx context.Context, q interface{}, variables map[string]interface{}, options ...graphql.Option) error
	Mutate(ctx context.Context, m interface{}, variables map[string]interface{}, options ...graphql.Option) error
}

// Client creates a new lb api client against a specific endpoint
type Client struct {
	gqlCli     GQLClient
	httpClient *http.Client
}

// Option is a function that modifies a client
type Option func(*Client)

// New creates a new ipam api client
func New(url string, opts ...Option) *Client {
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
func WithHTTPClient(cli *http.Client) Option {
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
		"id": graphql.ID(id),
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
		"id": graphql.ID(id),
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
			return nxt, err
		}

		cidr.IPIncr(bgn)

		if bgn.Equal(end) {
			err = ErrNoAvailableIPs
			return "", err
		}
	}
}

// CreateIPAddressFromBlock creates an IP Address from the next available address in a given block
func (c *Client) CreateIPAddressFromBlock(ctx context.Context, blockid string, nodeid string, nodeownerid string, reserve bool) (*CreateIPAddressResult, error) {
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

	return &ipn.CreateIPAddressResult, err
}

// DeleteIPAddress deletes an IP Address by id
func (c *Client) DeleteIPAddress(ctx context.Context, id string) (*DeleteIPAddress, error) {
	_, err := gidx.Parse(id)
	if err != nil {
		return nil, err
	}

	vars := map[string]interface{}{
		"id": graphql.ID(id),
	}

	var ipd DeleteIPAddress
	if err := c.gqlCli.Mutate(ctx, &ipd, vars); err != nil {
		return nil, err
	}

	return &ipd, nil
}

// GetIPAddresses returns a list of loadbalancer IP Addresses from node id
func (c *Client) GetIPAddresses(ctx context.Context, nodeID string) ([]IPAddressNode, error) {
	_, err := gidx.Parse(nodeID)
	if err != nil {
		return nil, err
	}

	vars := map[string]interface{}{
		"id": graphql.ID(nodeID),
	}

	var nodeIPs GetIPAddressesByNode
	if err := c.gqlCli.Query(ctx, &nodeIPs, vars); err != nil {
		return nil, err
	}

	return nodeIPs.NodeIPAddress.LoadBalancerFragment.IPAddresses, nil
}

// GetIPAddressByNode returns an IP Address by node id
func (c *Client) GetIPAddressByNode(ctx context.Context, nodeID string) (*GetIPAddressByNode, error) {
	_, err := gidx.Parse(nodeID)
	if err != nil {
		return nil, err
	}

	vars := map[string]interface{}{
		"id": graphql.ID(nodeID),
	}

	var nodeIP GetIPAddressByNode
	if err := c.gqlCli.Query(ctx, &nodeIP, vars); err != nil {
		return nil, err
	}

	return &nodeIP, nil
}
