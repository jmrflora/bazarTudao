// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/jmrflora/bazarTudao/ent/itemordem"
	"github.com/jmrflora/bazarTudao/ent/ordem"
	"github.com/jmrflora/bazarTudao/ent/produto"
	"github.com/jmrflora/bazarTudao/ent/stock"
)

// ProdutoCreate is the builder for creating a Produto entity.
type ProdutoCreate struct {
	config
	mutation *ProdutoMutation
	hooks    []Hook
}

// SetSku sets the "sku" field.
func (pc *ProdutoCreate) SetSku(s string) *ProdutoCreate {
	pc.mutation.SetSku(s)
	return pc
}

// SetNome sets the "nome" field.
func (pc *ProdutoCreate) SetNome(s string) *ProdutoCreate {
	pc.mutation.SetNome(s)
	return pc
}

// SetQuantNoEstoque sets the "quant_no_estoque" field.
func (pc *ProdutoCreate) SetQuantNoEstoque(i int) *ProdutoCreate {
	pc.mutation.SetQuantNoEstoque(i)
	return pc
}

// AddOrdenIDs adds the "ordens" edge to the Ordem entity by IDs.
func (pc *ProdutoCreate) AddOrdenIDs(ids ...int) *ProdutoCreate {
	pc.mutation.AddOrdenIDs(ids...)
	return pc
}

// AddOrdens adds the "ordens" edges to the Ordem entity.
func (pc *ProdutoCreate) AddOrdens(o ...*Ordem) *ProdutoCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return pc.AddOrdenIDs(ids...)
}

// AddStockIDs adds the "stock" edge to the Stock entity by IDs.
func (pc *ProdutoCreate) AddStockIDs(ids ...int) *ProdutoCreate {
	pc.mutation.AddStockIDs(ids...)
	return pc
}

// AddStock adds the "stock" edges to the Stock entity.
func (pc *ProdutoCreate) AddStock(s ...*Stock) *ProdutoCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pc.AddStockIDs(ids...)
}

// AddItenIDs adds the "itens" edge to the ItemOrdem entity by IDs.
func (pc *ProdutoCreate) AddItenIDs(ids ...int) *ProdutoCreate {
	pc.mutation.AddItenIDs(ids...)
	return pc
}

// AddItens adds the "itens" edges to the ItemOrdem entity.
func (pc *ProdutoCreate) AddItens(i ...*ItemOrdem) *ProdutoCreate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return pc.AddItenIDs(ids...)
}

// Mutation returns the ProdutoMutation object of the builder.
func (pc *ProdutoCreate) Mutation() *ProdutoMutation {
	return pc.mutation
}

// Save creates the Produto in the database.
func (pc *ProdutoCreate) Save(ctx context.Context) (*Produto, error) {
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProdutoCreate) SaveX(ctx context.Context) *Produto {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProdutoCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProdutoCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProdutoCreate) check() error {
	if _, ok := pc.mutation.Sku(); !ok {
		return &ValidationError{Name: "sku", err: errors.New(`ent: missing required field "Produto.sku"`)}
	}
	if _, ok := pc.mutation.Nome(); !ok {
		return &ValidationError{Name: "nome", err: errors.New(`ent: missing required field "Produto.nome"`)}
	}
	if _, ok := pc.mutation.QuantNoEstoque(); !ok {
		return &ValidationError{Name: "quant_no_estoque", err: errors.New(`ent: missing required field "Produto.quant_no_estoque"`)}
	}
	return nil
}

func (pc *ProdutoCreate) sqlSave(ctx context.Context) (*Produto, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *ProdutoCreate) createSpec() (*Produto, *sqlgraph.CreateSpec) {
	var (
		_node = &Produto{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(produto.Table, sqlgraph.NewFieldSpec(produto.FieldID, field.TypeInt))
	)
	if value, ok := pc.mutation.Sku(); ok {
		_spec.SetField(produto.FieldSku, field.TypeString, value)
		_node.Sku = value
	}
	if value, ok := pc.mutation.Nome(); ok {
		_spec.SetField(produto.FieldNome, field.TypeString, value)
		_node.Nome = value
	}
	if value, ok := pc.mutation.QuantNoEstoque(); ok {
		_spec.SetField(produto.FieldQuantNoEstoque, field.TypeInt, value)
		_node.QuantNoEstoque = value
	}
	if nodes := pc.mutation.OrdensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   produto.OrdensTable,
			Columns: produto.OrdensPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ordem.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.StockIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   produto.StockTable,
			Columns: []string{produto.StockColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(stock.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ItensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   produto.ItensTable,
			Columns: []string{produto.ItensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(itemordem.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProdutoCreateBulk is the builder for creating many Produto entities in bulk.
type ProdutoCreateBulk struct {
	config
	err      error
	builders []*ProdutoCreate
}

// Save creates the Produto entities in the database.
func (pcb *ProdutoCreateBulk) Save(ctx context.Context) ([]*Produto, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Produto, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProdutoMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProdutoCreateBulk) SaveX(ctx context.Context) []*Produto {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProdutoCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProdutoCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
