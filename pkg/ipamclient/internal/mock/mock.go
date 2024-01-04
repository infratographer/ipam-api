package mock

import (
	"context"

	"github.com/hasura/go-graphql-client"
)

// GQLClient is the mock http client
type GQLClient struct {
	DoQuery  func(ctx context.Context, q interface{}, variables map[string]interface{}, options ...graphql.Option) error
	DoMutate func(ctx context.Context, m interface{}, variables map[string]interface{}, options ...graphql.Option) error
}

// Query is the mock for a gqlclient query
func (c *GQLClient) Query(ctx context.Context, q interface{}, variables map[string]interface{}, options ...graphql.Option) error {
	return c.DoQuery(ctx, q, variables, options...)
}

// Mutate is the mock for a gqlclient mutate
func (c *GQLClient) Mutate(ctx context.Context, m interface{}, variables map[string]interface{}, options ...graphql.Option) error {
	return c.DoMutate(ctx, m, variables, options...)
}
