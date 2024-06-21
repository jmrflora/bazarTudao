// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/jmrflora/bazarTudao/ent/predicate"
	"github.com/jmrflora/bazarTudao/ent/produto"
)

// ProdutoDelete is the builder for deleting a Produto entity.
type ProdutoDelete struct {
	config
	hooks    []Hook
	mutation *ProdutoMutation
}

// Where appends a list predicates to the ProdutoDelete builder.
func (pd *ProdutoDelete) Where(ps ...predicate.Produto) *ProdutoDelete {
	pd.mutation.Where(ps...)
	return pd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pd *ProdutoDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pd.sqlExec, pd.mutation, pd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pd *ProdutoDelete) ExecX(ctx context.Context) int {
	n, err := pd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pd *ProdutoDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(produto.Table, sqlgraph.NewFieldSpec(produto.FieldID, field.TypeInt))
	if ps := pd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pd.mutation.done = true
	return affected, err
}

// ProdutoDeleteOne is the builder for deleting a single Produto entity.
type ProdutoDeleteOne struct {
	pd *ProdutoDelete
}

// Where appends a list predicates to the ProdutoDelete builder.
func (pdo *ProdutoDeleteOne) Where(ps ...predicate.Produto) *ProdutoDeleteOne {
	pdo.pd.mutation.Where(ps...)
	return pdo
}

// Exec executes the deletion query.
func (pdo *ProdutoDeleteOne) Exec(ctx context.Context) error {
	n, err := pdo.pd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{produto.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pdo *ProdutoDeleteOne) ExecX(ctx context.Context) {
	if err := pdo.Exec(ctx); err != nil {
		panic(err)
	}
}