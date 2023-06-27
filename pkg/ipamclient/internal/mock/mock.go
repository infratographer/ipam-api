package mock

import (
	"context"
)

// GQLClient is the mock http client
type GQLClient struct {
	DoQuery  func(ctx context.Context, q interface{}, variables map[string]interface{}) error
	DoMutate func(ctx context.Context, m interface{}, variables map[string]interface{}) error
}

// Query is the mock for a gqlclient query
func (c *GQLClient) Query(ctx context.Context, q interface{}, variables map[string]interface{}) error {
	return c.DoQuery(ctx, q, variables)
}

// Mutate is the mock for a gqlclient mutate
func (c *GQLClient) Mutate(ctx context.Context, m interface{}, variables map[string]interface{}) error {
	return c.DoMutate(ctx, m, variables)
}
