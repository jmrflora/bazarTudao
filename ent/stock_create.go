// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/jmrflora/bazarTudao/ent/produto"
	"github.com/jmrflora/bazarTudao/ent/stock"
)

// StockCreate is the builder for creating a Stock entity.
type StockCreate struct {
	config
	mutation *StockMutation
	hooks    []Hook
}

// SetDataMovimento sets the "data_movimento" field.
func (sc *StockCreate) SetDataMovimento(t time.Time) *StockCreate {
	sc.mutation.SetDataMovimento(t)
	return sc
}

// SetNillableDataMovimento sets the "data_movimento" field if the given value is not nil.
func (sc *StockCreate) SetNillableDataMovimento(t *time.Time) *StockCreate {
	if t != nil {
		sc.SetDataMovimento(*t)
	}
	return sc
}

// SetQuantidade sets the "quantidade" field.
func (sc *StockCreate) SetQuantidade(i int) *StockCreate {
	sc.mutation.SetQuantidade(i)
	return sc
}

// AddProdutoIDs adds the "produtos" edge to the Produto entity by IDs.
func (sc *StockCreate) AddProdutoIDs(ids ...int) *StockCreate {
	sc.mutation.AddProdutoIDs(ids...)
	return sc
}

// AddProdutos adds the "produtos" edges to the Produto entity.
func (sc *StockCreate) AddProdutos(p ...*Produto) *StockCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return sc.AddProdutoIDs(ids...)
}

// Mutation returns the StockMutation object of the builder.
func (sc *StockCreate) Mutation() *StockMutation {
	return sc.mutation
}

// Save creates the Stock in the database.
func (sc *StockCreate) Save(ctx context.Context) (*Stock, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StockCreate) SaveX(ctx context.Context) *Stock {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *StockCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *StockCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *StockCreate) defaults() {
	if _, ok := sc.mutation.DataMovimento(); !ok {
		v := stock.DefaultDataMovimento()
		sc.mutation.SetDataMovimento(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *StockCreate) check() error {
	if _, ok := sc.mutation.DataMovimento(); !ok {
		return &ValidationError{Name: "data_movimento", err: errors.New(`ent: missing required field "Stock.data_movimento"`)}
	}
	if _, ok := sc.mutation.Quantidade(); !ok {
		return &ValidationError{Name: "quantidade", err: errors.New(`ent: missing required field "Stock.quantidade"`)}
	}
	return nil
}

func (sc *StockCreate) sqlSave(ctx context.Context) (*Stock, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *StockCreate) createSpec() (*Stock, *sqlgraph.CreateSpec) {
	var (
		_node = &Stock{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(stock.Table, sqlgraph.NewFieldSpec(stock.FieldID, field.TypeInt))
	)
	if value, ok := sc.mutation.DataMovimento(); ok {
		_spec.SetField(stock.FieldDataMovimento, field.TypeTime, value)
		_node.DataMovimento = value
	}
	if value, ok := sc.mutation.Quantidade(); ok {
		_spec.SetField(stock.FieldQuantidade, field.TypeInt, value)
		_node.Quantidade = value
	}
	if nodes := sc.mutation.ProdutosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   stock.ProdutosTable,
			Columns: []string{stock.ProdutosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(produto.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// StockCreateBulk is the builder for creating many Stock entities in bulk.
type StockCreateBulk struct {
	config
	err      error
	builders []*StockCreate
}

// Save creates the Stock entities in the database.
func (scb *StockCreateBulk) Save(ctx context.Context) ([]*Stock, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Stock, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StockMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *StockCreateBulk) SaveX(ctx context.Context) []*Stock {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *StockCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *StockCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}