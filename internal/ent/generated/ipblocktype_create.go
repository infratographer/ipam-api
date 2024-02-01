// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/ipam-api/internal/ent/generated/ipblock"
	"go.infratographer.com/ipam-api/internal/ent/generated/ipblocktype"
	"go.infratographer.com/x/gidx"
)

// IPBlockTypeCreate is the builder for creating a IPBlockType entity.
type IPBlockTypeCreate struct {
	config
	mutation *IPBlockTypeMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ibtc *IPBlockTypeCreate) SetCreatedAt(t time.Time) *IPBlockTypeCreate {
	ibtc.mutation.SetCreatedAt(t)
	return ibtc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ibtc *IPBlockTypeCreate) SetNillableCreatedAt(t *time.Time) *IPBlockTypeCreate {
	if t != nil {
		ibtc.SetCreatedAt(*t)
	}
	return ibtc
}

// SetUpdatedAt sets the "updated_at" field.
func (ibtc *IPBlockTypeCreate) SetUpdatedAt(t time.Time) *IPBlockTypeCreate {
	ibtc.mutation.SetUpdatedAt(t)
	return ibtc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ibtc *IPBlockTypeCreate) SetNillableUpdatedAt(t *time.Time) *IPBlockTypeCreate {
	if t != nil {
		ibtc.SetUpdatedAt(*t)
	}
	return ibtc
}

// SetName sets the "name" field.
func (ibtc *IPBlockTypeCreate) SetName(s string) *IPBlockTypeCreate {
	ibtc.mutation.SetName(s)
	return ibtc
}

// SetOwnerID sets the "owner_id" field.
func (ibtc *IPBlockTypeCreate) SetOwnerID(gi gidx.PrefixedID) *IPBlockTypeCreate {
	ibtc.mutation.SetOwnerID(gi)
	return ibtc
}

// SetID sets the "id" field.
func (ibtc *IPBlockTypeCreate) SetID(gi gidx.PrefixedID) *IPBlockTypeCreate {
	ibtc.mutation.SetID(gi)
	return ibtc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ibtc *IPBlockTypeCreate) SetNillableID(gi *gidx.PrefixedID) *IPBlockTypeCreate {
	if gi != nil {
		ibtc.SetID(*gi)
	}
	return ibtc
}

// AddIPBlockIDs adds the "ip_block" edge to the IPBlock entity by IDs.
func (ibtc *IPBlockTypeCreate) AddIPBlockIDs(ids ...gidx.PrefixedID) *IPBlockTypeCreate {
	ibtc.mutation.AddIPBlockIDs(ids...)
	return ibtc
}

// AddIPBlock adds the "ip_block" edges to the IPBlock entity.
func (ibtc *IPBlockTypeCreate) AddIPBlock(i ...*IPBlock) *IPBlockTypeCreate {
	ids := make([]gidx.PrefixedID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ibtc.AddIPBlockIDs(ids...)
}

// Mutation returns the IPBlockTypeMutation object of the builder.
func (ibtc *IPBlockTypeCreate) Mutation() *IPBlockTypeMutation {
	return ibtc.mutation
}

// Save creates the IPBlockType in the database.
func (ibtc *IPBlockTypeCreate) Save(ctx context.Context) (*IPBlockType, error) {
	ibtc.defaults()
	return withHooks(ctx, ibtc.sqlSave, ibtc.mutation, ibtc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ibtc *IPBlockTypeCreate) SaveX(ctx context.Context) *IPBlockType {
	v, err := ibtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ibtc *IPBlockTypeCreate) Exec(ctx context.Context) error {
	_, err := ibtc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ibtc *IPBlockTypeCreate) ExecX(ctx context.Context) {
	if err := ibtc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ibtc *IPBlockTypeCreate) defaults() {
	if _, ok := ibtc.mutation.CreatedAt(); !ok {
		v := ipblocktype.DefaultCreatedAt()
		ibtc.mutation.SetCreatedAt(v)
	}
	if _, ok := ibtc.mutation.UpdatedAt(); !ok {
		v := ipblocktype.DefaultUpdatedAt()
		ibtc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ibtc.mutation.ID(); !ok {
		v := ipblocktype.DefaultID()
		ibtc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ibtc *IPBlockTypeCreate) check() error {
	if _, ok := ibtc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`generated: missing required field "IPBlockType.created_at"`)}
	}
	if _, ok := ibtc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`generated: missing required field "IPBlockType.updated_at"`)}
	}
	if _, ok := ibtc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`generated: missing required field "IPBlockType.name"`)}
	}
	if v, ok := ibtc.mutation.Name(); ok {
		if err := ipblocktype.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "IPBlockType.name": %w`, err)}
		}
	}
	if _, ok := ibtc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner_id", err: errors.New(`generated: missing required field "IPBlockType.owner_id"`)}
	}
	return nil
}

func (ibtc *IPBlockTypeCreate) sqlSave(ctx context.Context) (*IPBlockType, error) {
	if err := ibtc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ibtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ibtc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*gidx.PrefixedID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	ibtc.mutation.id = &_node.ID
	ibtc.mutation.done = true
	return _node, nil
}

func (ibtc *IPBlockTypeCreate) createSpec() (*IPBlockType, *sqlgraph.CreateSpec) {
	var (
		_node = &IPBlockType{config: ibtc.config}
		_spec = sqlgraph.NewCreateSpec(ipblocktype.Table, sqlgraph.NewFieldSpec(ipblocktype.FieldID, field.TypeString))
	)
	if id, ok := ibtc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ibtc.mutation.CreatedAt(); ok {
		_spec.SetField(ipblocktype.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ibtc.mutation.UpdatedAt(); ok {
		_spec.SetField(ipblocktype.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ibtc.mutation.Name(); ok {
		_spec.SetField(ipblocktype.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ibtc.mutation.OwnerID(); ok {
		_spec.SetField(ipblocktype.FieldOwnerID, field.TypeString, value)
		_node.OwnerID = value
	}
	if nodes := ibtc.mutation.IPBlockIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   ipblocktype.IPBlockTable,
			Columns: []string{ipblocktype.IPBlockColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ipblock.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// IPBlockTypeCreateBulk is the builder for creating many IPBlockType entities in bulk.
type IPBlockTypeCreateBulk struct {
	config
	err      error
	builders []*IPBlockTypeCreate
}

// Save creates the IPBlockType entities in the database.
func (ibtcb *IPBlockTypeCreateBulk) Save(ctx context.Context) ([]*IPBlockType, error) {
	if ibtcb.err != nil {
		return nil, ibtcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ibtcb.builders))
	nodes := make([]*IPBlockType, len(ibtcb.builders))
	mutators := make([]Mutator, len(ibtcb.builders))
	for i := range ibtcb.builders {
		func(i int, root context.Context) {
			builder := ibtcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*IPBlockTypeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ibtcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ibtcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ibtcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ibtcb *IPBlockTypeCreateBulk) SaveX(ctx context.Context) []*IPBlockType {
	v, err := ibtcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ibtcb *IPBlockTypeCreateBulk) Exec(ctx context.Context) error {
	_, err := ibtcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ibtcb *IPBlockTypeCreateBulk) ExecX(ctx context.Context) {
	if err := ibtcb.Exec(ctx); err != nil {
		panic(err)
	}
}
