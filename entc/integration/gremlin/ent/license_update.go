// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/entc/integration/gremlin/ent/license"
	"entgo.io/ent/entc/integration/gremlin/ent/predicate"
)

// LicenseUpdate is the builder for updating License entities.
type LicenseUpdate struct {
	config
	hooks    []Hook
	mutation *LicenseMutation
}

// Where appends a list predicates to the LicenseUpdate builder.
func (lu *LicenseUpdate) Where(ps ...predicate.License) *LicenseUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// Mutation returns the LicenseMutation object of the builder.
func (lu *LicenseUpdate) Mutation() *LicenseMutation {
	return lu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LicenseUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(lu.hooks) == 0 {
		affected, err = lu.gremlinSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LicenseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lu.mutation = mutation
			affected, err = lu.gremlinSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lu.hooks) - 1; i >= 0; i-- {
			if lu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LicenseUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LicenseUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LicenseUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lu *LicenseUpdate) gremlinSave(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := lu.gremlin().Query()
	if err := lu.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	if err, ok := isConstantError(res); ok {
		return 0, err
	}
	return res.ReadInt()
}

func (lu *LicenseUpdate) gremlin() *dsl.Traversal {
	v := g.V().HasLabel(license.Label)
	for _, p := range lu.mutation.predicates {
		p(v)
	}
	var (
		trs []*dsl.Traversal
	)
	v.Count()
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// LicenseUpdateOne is the builder for updating a single License entity.
type LicenseUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LicenseMutation
}

// Mutation returns the LicenseMutation object of the builder.
func (luo *LicenseUpdateOne) Mutation() *LicenseMutation {
	return luo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LicenseUpdateOne) Select(field string, fields ...string) *LicenseUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated License entity.
func (luo *LicenseUpdateOne) Save(ctx context.Context) (*License, error) {
	var (
		err  error
		node *License
	)
	if len(luo.hooks) == 0 {
		node, err = luo.gremlinSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LicenseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			luo.mutation = mutation
			node, err = luo.gremlinSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(luo.hooks) - 1; i >= 0; i-- {
			if luo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = luo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, luo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*License)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from LicenseMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LicenseUpdateOne) SaveX(ctx context.Context) *License {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LicenseUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LicenseUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (luo *LicenseUpdateOne) gremlinSave(ctx context.Context) (*License, error) {
	res := &gremlin.Response{}
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "License.id" for update`)}
	}
	query, bindings := luo.gremlin(id).Query()
	if err := luo.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	l := &License{config: luo.config}
	if err := l.FromResponse(res); err != nil {
		return nil, err
	}
	return l, nil
}

func (luo *LicenseUpdateOne) gremlin(id int) *dsl.Traversal {
	v := g.V(id)
	var (
		trs []*dsl.Traversal
	)
	if len(luo.fields) > 0 {
		fields := make([]interface{}, 0, len(luo.fields)+1)
		fields = append(fields, true)
		for _, f := range luo.fields {
			fields = append(fields, f)
		}
		v.ValueMap(fields...)
	} else {
		v.ValueMap(true)
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}
