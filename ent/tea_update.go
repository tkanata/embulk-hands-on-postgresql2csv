// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tkanata/embulk-hands-on-postgresql2csv/ent/predicate"
	"github.com/tkanata/embulk-hands-on-postgresql2csv/ent/tea"
)

// TeaUpdate is the builder for updating Tea entities.
type TeaUpdate struct {
	config
	hooks    []Hook
	mutation *TeaMutation
}

// Where appends a list predicates to the TeaUpdate builder.
func (tu *TeaUpdate) Where(ps ...predicate.Tea) *TeaUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetName sets the "name" field.
func (tu *TeaUpdate) SetName(s string) *TeaUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetColor sets the "color" field.
func (tu *TeaUpdate) SetColor(s string) *TeaUpdate {
	tu.mutation.SetColor(s)
	return tu
}

// Mutation returns the TeaMutation object of the builder.
func (tu *TeaUpdate) Mutation() *TeaMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TeaUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TeaUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TeaUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TeaUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TeaUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(tea.Table, tea.Columns, sqlgraph.NewFieldSpec(tea.FieldID, field.TypeString))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.SetField(tea.FieldName, field.TypeString, value)
	}
	if value, ok := tu.mutation.Color(); ok {
		_spec.SetField(tea.FieldColor, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tea.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TeaUpdateOne is the builder for updating a single Tea entity.
type TeaUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TeaMutation
}

// SetName sets the "name" field.
func (tuo *TeaUpdateOne) SetName(s string) *TeaUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetColor sets the "color" field.
func (tuo *TeaUpdateOne) SetColor(s string) *TeaUpdateOne {
	tuo.mutation.SetColor(s)
	return tuo
}

// Mutation returns the TeaMutation object of the builder.
func (tuo *TeaUpdateOne) Mutation() *TeaMutation {
	return tuo.mutation
}

// Where appends a list predicates to the TeaUpdate builder.
func (tuo *TeaUpdateOne) Where(ps ...predicate.Tea) *TeaUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TeaUpdateOne) Select(field string, fields ...string) *TeaUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Tea entity.
func (tuo *TeaUpdateOne) Save(ctx context.Context) (*Tea, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TeaUpdateOne) SaveX(ctx context.Context) *Tea {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TeaUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TeaUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TeaUpdateOne) sqlSave(ctx context.Context) (_node *Tea, err error) {
	_spec := sqlgraph.NewUpdateSpec(tea.Table, tea.Columns, sqlgraph.NewFieldSpec(tea.FieldID, field.TypeString))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Tea.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tea.FieldID)
		for _, f := range fields {
			if !tea.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tea.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.SetField(tea.FieldName, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Color(); ok {
		_spec.SetField(tea.FieldColor, field.TypeString, value)
	}
	_node = &Tea{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tea.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}