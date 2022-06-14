// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgefield/ent/car"
	"entgo.io/ent/entc/integration/edgefield/ent/predicate"
	"entgo.io/ent/entc/integration/edgefield/ent/rental"
	"entgo.io/ent/schema/field"
)

// CarUpdate is the builder for updating Car entities.
type CarUpdate struct {
	config
	hooks    []Hook
	mutation *CarMutation
}

// Where appends a list predicates to the CarUpdate builder.
func (cu *CarUpdate) Where(ps ...predicate.Car) *CarUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetNumber sets the "number" field.
func (cu *CarUpdate) SetNumber(s string) *CarUpdate {
	cu.mutation.SetNumber(s)
	return cu
}

// SetNillableNumber sets the "number" field if the given value is not nil.
func (cu *CarUpdate) SetNillableNumber(s *string) *CarUpdate {
	if s != nil {
		cu.SetNumber(*s)
	}
	return cu
}

// ClearNumber clears the value of the "number" field.
func (cu *CarUpdate) ClearNumber() *CarUpdate {
	cu.mutation.ClearNumber()
	return cu
}

// AddRentalIDs adds the "rentals" edge to the Rental entity by IDs.
func (cu *CarUpdate) AddRentalIDs(ids ...int) *CarUpdate {
	cu.mutation.AddRentalIDs(ids...)
	return cu
}

// AddRentals adds the "rentals" edges to the Rental entity.
func (cu *CarUpdate) AddRentals(r ...*Rental) *CarUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cu.AddRentalIDs(ids...)
}

// Mutation returns the CarMutation object of the builder.
func (cu *CarUpdate) Mutation() *CarMutation {
	return cu.mutation
}

// ClearRentals clears all "rentals" edges to the Rental entity.
func (cu *CarUpdate) ClearRentals() *CarUpdate {
	cu.mutation.ClearRentals()
	return cu
}

// RemoveRentalIDs removes the "rentals" edge to Rental entities by IDs.
func (cu *CarUpdate) RemoveRentalIDs(ids ...int) *CarUpdate {
	cu.mutation.RemoveRentalIDs(ids...)
	return cu
}

// RemoveRentals removes "rentals" edges to Rental entities.
func (cu *CarUpdate) RemoveRentals(r ...*Rental) *CarUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cu.RemoveRentalIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CarUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CarMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
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
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   car.Table,
			Columns: car.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: car.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Number(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: car.FieldNumber,
		})
	}
	if cu.mutation.NumberCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: car.FieldNumber,
		})
	}
	if cu.mutation.RentalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   car.RentalsTable,
			Columns: []string{car.RentalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: rental.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedRentalsIDs(); len(nodes) > 0 && !cu.mutation.RentalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   car.RentalsTable,
			Columns: []string{car.RentalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: rental.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RentalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   car.RentalsTable,
			Columns: []string{car.RentalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: rental.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{car.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CarUpdateOne is the builder for updating a single Car entity.
type CarUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CarMutation
}

// SetNumber sets the "number" field.
func (cuo *CarUpdateOne) SetNumber(s string) *CarUpdateOne {
	cuo.mutation.SetNumber(s)
	return cuo
}

// SetNillableNumber sets the "number" field if the given value is not nil.
func (cuo *CarUpdateOne) SetNillableNumber(s *string) *CarUpdateOne {
	if s != nil {
		cuo.SetNumber(*s)
	}
	return cuo
}

// ClearNumber clears the value of the "number" field.
func (cuo *CarUpdateOne) ClearNumber() *CarUpdateOne {
	cuo.mutation.ClearNumber()
	return cuo
}

// AddRentalIDs adds the "rentals" edge to the Rental entity by IDs.
func (cuo *CarUpdateOne) AddRentalIDs(ids ...int) *CarUpdateOne {
	cuo.mutation.AddRentalIDs(ids...)
	return cuo
}

// AddRentals adds the "rentals" edges to the Rental entity.
func (cuo *CarUpdateOne) AddRentals(r ...*Rental) *CarUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cuo.AddRentalIDs(ids...)
}

// Mutation returns the CarMutation object of the builder.
func (cuo *CarUpdateOne) Mutation() *CarMutation {
	return cuo.mutation
}

// ClearRentals clears all "rentals" edges to the Rental entity.
func (cuo *CarUpdateOne) ClearRentals() *CarUpdateOne {
	cuo.mutation.ClearRentals()
	return cuo
}

// RemoveRentalIDs removes the "rentals" edge to Rental entities by IDs.
func (cuo *CarUpdateOne) RemoveRentalIDs(ids ...int) *CarUpdateOne {
	cuo.mutation.RemoveRentalIDs(ids...)
	return cuo
}

// RemoveRentals removes "rentals" edges to Rental entities.
func (cuo *CarUpdateOne) RemoveRentals(r ...*Rental) *CarUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cuo.RemoveRentalIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CarUpdateOne) Select(field string, fields ...string) *CarUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Car entity.
func (cuo *CarUpdateOne) Save(ctx context.Context) (*Car, error) {
	var (
		err  error
		node *Car
	)
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CarMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Car)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CarMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CarUpdateOne) SaveX(ctx context.Context) *Car {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
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

func (cuo *CarUpdateOne) sqlSave(ctx context.Context) (_node *Car, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   car.Table,
			Columns: car.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: car.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Car.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, car.FieldID)
		for _, f := range fields {
			if !car.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != car.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Number(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: car.FieldNumber,
		})
	}
	if cuo.mutation.NumberCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: car.FieldNumber,
		})
	}
	if cuo.mutation.RentalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   car.RentalsTable,
			Columns: []string{car.RentalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: rental.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedRentalsIDs(); len(nodes) > 0 && !cuo.mutation.RentalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   car.RentalsTable,
			Columns: []string{car.RentalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: rental.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RentalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   car.RentalsTable,
			Columns: []string{car.RentalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: rental.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Car{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{car.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
