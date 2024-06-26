package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.38

import (
	"context"

	"go.infratographer.com/ipam-api/internal/ent/generated"
	"go.infratographer.com/ipam-api/internal/ent/generated/ipaddress"
	"go.infratographer.com/permissions-api/pkg/permissions"
	"go.infratographer.com/x/gidx"
)

// Node is the resolver for the node field.
func (r *iPAddressResolver) Node(ctx context.Context, obj *generated.IPAddress) (*IPAddressable, error) {
	return &IPAddressable{ID: obj.NodeID}, nil
}

// IPAddresses is the resolver for the IPAddresses field.
func (r *iPAddressableResolver) IPAddresses(ctx context.Context, obj *IPAddressable) ([]*generated.IPAddress, error) {
	if err := permissions.CheckAccess(ctx, obj.ID, actionIPBlockGet); err != nil {
		return nil, err
	}

	m, err := r.client.IPAddress.Query().Where(ipaddress.NodeID(obj.ID)).All(ctx)
	return m, err
}

// CreateIPAddress is the resolver for the createIPAddress field.
func (r *mutationResolver) CreateIPAddress(ctx context.Context, input generated.CreateIPAddressInput) (*IPAddressCreatePayload, error) {
	if err := permissions.CheckAccess(ctx, input.NodeOwnerID, actionIPBlockCreate); err != nil {
		return nil, err
	}

	t, err := r.client.IPAddress.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &IPAddressCreatePayload{IPAddress: t}, err
}

// UpdateIPAddress is the resolver for the updateIPAddress field.
func (r *mutationResolver) UpdateIPAddress(ctx context.Context, id gidx.PrefixedID, input generated.UpdateIPAddressInput) (*IPAddressUpdatePayload, error) {
	if err := permissions.CheckAccess(ctx, id, actionIPBlockUpdate); err != nil {
		return nil, err
	}

	t, err := r.client.IPAddress.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	t, err = t.Update().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &IPAddressUpdatePayload{IPAddress: t}, err
}

// DeleteIPAddress is the resolver for the deleteIPAddress field.
func (r *mutationResolver) DeleteIPAddress(ctx context.Context, id gidx.PrefixedID) (*IPAddressDeletePayload, error) {
	if err := permissions.CheckAccess(ctx, id, actionIPBlockDelete); err != nil {
		return nil, err
	}

	if err := r.client.IPAddress.DeleteOneID(id).Exec(ctx); err != nil {
		return nil, err
	}

	return &IPAddressDeletePayload{DeletedID: id}, nil
}

// IPAddress is the resolver for the ip_address field.
func (r *queryResolver) IPAddress(ctx context.Context, id gidx.PrefixedID) (*generated.IPAddress, error) {
	if err := permissions.CheckAccess(ctx, id, actionIPBlockGet); err != nil {
		return nil, err
	}

	return r.client.IPAddress.Get(ctx, id)
}

// IPAddressByNode is the resolver for the ipAddressByNode field.
func (r *queryResolver) IPAddressByNode(ctx context.Context, id gidx.PrefixedID) (*generated.IPAddress, error) {
	if err := permissions.CheckAccess(ctx, id, actionIPBlockGet); err != nil {
		return nil, err
	}

	return r.client.IPAddress.Query().Where(ipaddress.NodeID(id)).First(ctx)
}

// IPAddressable returns IPAddressableResolver implementation.
func (r *Resolver) IPAddressable() IPAddressableResolver { return &iPAddressableResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type iPAddressableResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
