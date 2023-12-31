// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tkanata/embulk-hands-on-postgresql2csv/ent/predicate"
	"github.com/tkanata/embulk-hands-on-postgresql2csv/ent/tea"
)

// TeaDelete is the builder for deleting a Tea entity.
type TeaDelete struct {
	config
	hooks    []Hook
	mutation *TeaMutation
}

// Where appends a list predicates to the TeaDelete builder.
func (td *TeaDelete) Where(ps ...predicate.Tea) *TeaDelete {
	td.mutation.Where(ps...)
	return td
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (td *TeaDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, td.sqlExec, td.mutation, td.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (td *TeaDelete) ExecX(ctx context.Context) int {
	n, err := td.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (td *TeaDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(tea.Table, sqlgraph.NewFieldSpec(tea.FieldID, field.TypeInt))
	if ps := td.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, td.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	td.mutation.done = true
	return affected, err
}

// TeaDeleteOne is the builder for deleting a single Tea entity.
type TeaDeleteOne struct {
	td *TeaDelete
}

// Where appends a list predicates to the TeaDelete builder.
func (tdo *TeaDeleteOne) Where(ps ...predicate.Tea) *TeaDeleteOne {
	tdo.td.mutation.Where(ps...)
	return tdo
}

// Exec executes the deletion query.
func (tdo *TeaDeleteOne) Exec(ctx context.Context) error {
	n, err := tdo.td.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{tea.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tdo *TeaDeleteOne) ExecX(ctx context.Context) {
	if err := tdo.Exec(ctx); err != nil {
		panic(err)
	}
}
