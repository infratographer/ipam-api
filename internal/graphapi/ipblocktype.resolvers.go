package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.32

import (
	"context"

	"go.infratographer.com/permissions-api/pkg/permissions"
	"go.infratographer.com/x/gidx"

	"go.infratographer.com/ipam-api/internal/ent/generated"
)

// CreateIPBlockType is the resolver for the createIPBlockType field.
func (r *mutationResolver) CreateIPBlockType(ctx context.Context, input generated.CreateIPBlockTypeInput) (*IPBlockTypeCreatePayload, error) {
	if err := permissions.CheckAccess(ctx, input.OwnerID, actionIPBlockTypeCreate); err != nil {
		return nil, err
	}

	t, err := r.client.IPBlockType.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &IPBlockTypeCreatePayload{IPBlockType: t}, err
}

// UpdateIPBlockType is the resolver for the updateIPBlockType field.
func (r *mutationResolver) UpdateIPBlockType(ctx context.Context, id gidx.PrefixedID, input generated.UpdateIPBlockTypeInput) (*IPBlockTypeUpdatePayload, error) {
	if err := permissions.CheckAccess(ctx, id, actionIPBlockTypeUpdate); err != nil {
		return nil, err
	}

	t, err := r.client.IPBlockType.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	t, err = t.Update().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &IPBlockTypeUpdatePayload{IPBlockType: t}, err
}

// DeleteIPBlockType is the resolver for the deleteIPBlockType field.
func (r *mutationResolver) DeleteIPBlockType(ctx context.Context, id gidx.PrefixedID) (*IPBlockTypeDeletePayload, error) {
	if err := permissions.CheckAccess(ctx, id, actionIPBlockTypeDelete); err != nil {
		return nil, err
	}

	if err := r.client.IPBlockType.DeleteOneID(id).Exec(ctx); err != nil {
		return nil, err
	}

	return &IPBlockTypeDeletePayload{DeletedID: id}, nil
}

// IPBlockType is the resolver for the ip_block_type field.
func (r *queryResolver) IPBlockType(ctx context.Context, id gidx.PrefixedID) (*generated.IPBlockType, error) {
	if err := permissions.CheckAccess(ctx, id, actionIPBlockTypeGet); err != nil {
		return nil, err
	}

	return r.client.IPBlockType.Get(ctx, id)
}
