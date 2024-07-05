// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/jmrflora/bazarTudao/ent/envio"
	"github.com/jmrflora/bazarTudao/ent/predicate"
)

// EnvioDelete is the builder for deleting a Envio entity.
type EnvioDelete struct {
	config
	hooks    []Hook
	mutation *EnvioMutation
}

// Where appends a list predicates to the EnvioDelete builder.
func (ed *EnvioDelete) Where(ps ...predicate.Envio) *EnvioDelete {
	ed.mutation.Where(ps...)
	return ed
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ed *EnvioDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ed.sqlExec, ed.mutation, ed.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ed *EnvioDelete) ExecX(ctx context.Context) int {
	n, err := ed.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ed *EnvioDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(envio.Table, sqlgraph.NewFieldSpec(envio.FieldID, field.TypeInt))
	if ps := ed.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ed.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ed.mutation.done = true
	return affected, err
}

// EnvioDeleteOne is the builder for deleting a single Envio entity.
type EnvioDeleteOne struct {
	ed *EnvioDelete
}

// Where appends a list predicates to the EnvioDelete builder.
func (edo *EnvioDeleteOne) Where(ps ...predicate.Envio) *EnvioDeleteOne {
	edo.ed.mutation.Where(ps...)
	return edo
}

// Exec executes the deletion query.
func (edo *EnvioDeleteOne) Exec(ctx context.Context) error {
	n, err := edo.ed.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{envio.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (edo *EnvioDeleteOne) ExecX(ctx context.Context) {
	if err := edo.Exec(ctx); err != nil {
		panic(err)
	}
}