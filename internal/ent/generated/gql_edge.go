// Copyright Infratographer, Inc. and/or licensed to Infratographer, Inc. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (ia *IPAddress) IPBlock(ctx context.Context) (*IPBlock, error) {
	result, err := ia.Edges.IPBlockOrErr()
	if IsNotLoaded(err) {
		result, err = ia.QueryIPBlock().Only(ctx)
	}
	return result, err
}

func (ib *IPBlock) IPBlockType(ctx context.Context) (*IPBlockType, error) {
	result, err := ib.Edges.IPBlockTypeOrErr()
	if IsNotLoaded(err) {
		result, err = ib.QueryIPBlockType().Only(ctx)
	}
	return result, err
}

func (ib *IPBlock) IPAddress(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *IPAddressOrder, where *IPAddressWhereInput,
) (*IPAddressConnection, error) {
	opts := []IPAddressPaginateOption{
		WithIPAddressOrder(orderBy),
		WithIPAddressFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := ib.Edges.totalCount[1][alias]
	if nodes, err := ib.NamedIPAddress(alias); err == nil || hasTotalCount {
		pager, err := newIPAddressPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &IPAddressConnection{Edges: []*IPAddressEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return ib.QueryIPAddress().Paginate(ctx, after, first, before, last, opts...)
}

func (ibt *IPBlockType) IPBlock(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *IPBlockOrder, where *IPBlockWhereInput,
) (*IPBlockConnection, error) {
	opts := []IPBlockPaginateOption{
		WithIPBlockOrder(orderBy),
		WithIPBlockFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := ibt.Edges.totalCount[0][alias]
	if nodes, err := ibt.NamedIPBlock(alias); err == nil || hasTotalCount {
		pager, err := newIPBlockPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &IPBlockConnection{Edges: []*IPBlockEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return ibt.QueryIPBlock().Paginate(ctx, after, first, before, last, opts...)
}
