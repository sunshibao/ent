// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package entv1

import (
	"context"
	"errors"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/entc/integration/migrate/entv1/car"
	"github.com/facebookincubator/ent/entc/integration/migrate/entv1/predicate"
	"github.com/facebookincubator/ent/entc/integration/migrate/entv1/user"
	"github.com/facebookincubator/ent/schema/field"
)

// CarUpdate is the builder for updating Car entities.
type CarUpdate struct {
	config
	owner        map[int]struct{}
	clearedOwner bool
	predicates   []predicate.Car
}

// Where adds a new predicate for the builder.
func (cu *CarUpdate) Where(ps ...predicate.Car) *CarUpdate {
	cu.predicates = append(cu.predicates, ps...)
	return cu
}

// SetOwnerID sets the owner edge to User by id.
func (cu *CarUpdate) SetOwnerID(id int) *CarUpdate {
	if cu.owner == nil {
		cu.owner = make(map[int]struct{})
	}
	cu.owner[id] = struct{}{}
	return cu
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (cu *CarUpdate) SetNillableOwnerID(id *int) *CarUpdate {
	if id != nil {
		cu = cu.SetOwnerID(*id)
	}
	return cu
}

// SetOwner sets the owner edge to User.
func (cu *CarUpdate) SetOwner(u *User) *CarUpdate {
	return cu.SetOwnerID(u.ID)
}

// ClearOwner clears the owner edge to User.
func (cu *CarUpdate) ClearOwner() *CarUpdate {
	cu.clearedOwner = true
	return cu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (cu *CarUpdate) Save(ctx context.Context) (int, error) {
	if len(cu.owner) > 1 {
		return 0, errors.New("entv1: multiple assignments on a unique edge \"owner\"")
	}
	return cu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CarUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CarUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CarUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CarUpdate) sqlSave(ctx context.Context) (n int, err error) {
	spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   car.Table,
			Columns: car.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: car.FieldID,
			},
		},
	}
	if ps := cu.predicates; len(ps) > 0 {
		spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if cu.clearedOwner {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   car.OwnerTable,
			Columns: []string{car.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		spec.Edges.Clear = append(spec.Edges.Clear, edge)
	}
	if nodes := cu.owner; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   car.OwnerTable,
			Columns: []string{car.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges.Add = append(spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// CarUpdateOne is the builder for updating a single Car entity.
type CarUpdateOne struct {
	config
	id           int
	owner        map[int]struct{}
	clearedOwner bool
}

// SetOwnerID sets the owner edge to User by id.
func (cuo *CarUpdateOne) SetOwnerID(id int) *CarUpdateOne {
	if cuo.owner == nil {
		cuo.owner = make(map[int]struct{})
	}
	cuo.owner[id] = struct{}{}
	return cuo
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (cuo *CarUpdateOne) SetNillableOwnerID(id *int) *CarUpdateOne {
	if id != nil {
		cuo = cuo.SetOwnerID(*id)
	}
	return cuo
}

// SetOwner sets the owner edge to User.
func (cuo *CarUpdateOne) SetOwner(u *User) *CarUpdateOne {
	return cuo.SetOwnerID(u.ID)
}

// ClearOwner clears the owner edge to User.
func (cuo *CarUpdateOne) ClearOwner() *CarUpdateOne {
	cuo.clearedOwner = true
	return cuo
}

// Save executes the query and returns the updated entity.
func (cuo *CarUpdateOne) Save(ctx context.Context) (*Car, error) {
	if len(cuo.owner) > 1 {
		return nil, errors.New("entv1: multiple assignments on a unique edge \"owner\"")
	}
	return cuo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CarUpdateOne) SaveX(ctx context.Context) *Car {
	c, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return c
}

// Exec executes the query on the entity.
func (cuo *CarUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CarUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CarUpdateOne) sqlSave(ctx context.Context) (c *Car, err error) {
	spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   car.Table,
			Columns: car.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  cuo.id,
				Type:   field.TypeInt,
				Column: car.FieldID,
			},
		},
	}
	if cuo.clearedOwner {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   car.OwnerTable,
			Columns: []string{car.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		spec.Edges.Clear = append(spec.Edges.Clear, edge)
	}
	if nodes := cuo.owner; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   car.OwnerTable,
			Columns: []string{car.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges.Add = append(spec.Edges.Add, edge)
	}
	c = &Car{config: cuo.config}
	spec.Assign = c.assignValues
	spec.ScanValues = c.scanValues()
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return c, nil
}
