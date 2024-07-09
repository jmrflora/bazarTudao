// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/jmrflora/bazarTudao/ent/cliente"
	"github.com/jmrflora/bazarTudao/ent/itemordem"
	"github.com/jmrflora/bazarTudao/ent/ordem"
	"github.com/jmrflora/bazarTudao/ent/produto"
)

// OrdemCreate is the builder for creating a Ordem entity.
type OrdemCreate struct {
	config
	mutation *OrdemMutation
	hooks    []Hook
}

// SetDataOrdem sets the "data_ordem" field.
func (oc *OrdemCreate) SetDataOrdem(t time.Time) *OrdemCreate {
	oc.mutation.SetDataOrdem(t)
	return oc
}

// SetNillableDataOrdem sets the "data_ordem" field if the given value is not nil.
func (oc *OrdemCreate) SetNillableDataOrdem(t *time.Time) *OrdemCreate {
	if t != nil {
		oc.SetDataOrdem(*t)
	}
	return oc
}

// SetStatus sets the "status" field.
func (oc *OrdemCreate) SetStatus(o ordem.Status) *OrdemCreate {
	oc.mutation.SetStatus(o)
	return oc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (oc *OrdemCreate) SetNillableStatus(o *ordem.Status) *OrdemCreate {
	if o != nil {
		oc.SetStatus(*o)
	}
	return oc
}

// SetPrecoDaOrdem sets the "preco_da_ordem" field.
func (oc *OrdemCreate) SetPrecoDaOrdem(f float64) *OrdemCreate {
	oc.mutation.SetPrecoDaOrdem(f)
	return oc
}

// SetID sets the "id" field.
func (oc *OrdemCreate) SetID(i int) *OrdemCreate {
	oc.mutation.SetID(i)
	return oc
}

// AddProdutoIDs adds the "produtos" edge to the Produto entity by IDs.
func (oc *OrdemCreate) AddProdutoIDs(ids ...int) *OrdemCreate {
	oc.mutation.AddProdutoIDs(ids...)
	return oc
}

// AddProdutos adds the "produtos" edges to the Produto entity.
func (oc *OrdemCreate) AddProdutos(p ...*Produto) *OrdemCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return oc.AddProdutoIDs(ids...)
}

// SetClienteID sets the "cliente" edge to the Cliente entity by ID.
func (oc *OrdemCreate) SetClienteID(id int) *OrdemCreate {
	oc.mutation.SetClienteID(id)
	return oc
}

// SetNillableClienteID sets the "cliente" edge to the Cliente entity by ID if the given value is not nil.
func (oc *OrdemCreate) SetNillableClienteID(id *int) *OrdemCreate {
	if id != nil {
		oc = oc.SetClienteID(*id)
	}
	return oc
}

// SetCliente sets the "cliente" edge to the Cliente entity.
func (oc *OrdemCreate) SetCliente(c *Cliente) *OrdemCreate {
	return oc.SetClienteID(c.ID)
}

// AddItemIDs adds the "items" edge to the ItemOrdem entity by IDs.
func (oc *OrdemCreate) AddItemIDs(ids ...int) *OrdemCreate {
	oc.mutation.AddItemIDs(ids...)
	return oc
}

// AddItems adds the "items" edges to the ItemOrdem entity.
func (oc *OrdemCreate) AddItems(i ...*ItemOrdem) *OrdemCreate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return oc.AddItemIDs(ids...)
}

// Mutation returns the OrdemMutation object of the builder.
func (oc *OrdemCreate) Mutation() *OrdemMutation {
	return oc.mutation
}

// Save creates the Ordem in the database.
func (oc *OrdemCreate) Save(ctx context.Context) (*Ordem, error) {
	oc.defaults()
	return withHooks(ctx, oc.sqlSave, oc.mutation, oc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OrdemCreate) SaveX(ctx context.Context) *Ordem {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *OrdemCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *OrdemCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oc *OrdemCreate) defaults() {
	if _, ok := oc.mutation.DataOrdem(); !ok {
		v := ordem.DefaultDataOrdem()
		oc.mutation.SetDataOrdem(v)
	}
	if _, ok := oc.mutation.Status(); !ok {
		v := ordem.DefaultStatus
		oc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oc *OrdemCreate) check() error {
	if _, ok := oc.mutation.DataOrdem(); !ok {
		return &ValidationError{Name: "data_ordem", err: errors.New(`ent: missing required field "Ordem.data_ordem"`)}
	}
	if _, ok := oc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Ordem.status"`)}
	}
	if v, ok := oc.mutation.Status(); ok {
		if err := ordem.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Ordem.status": %w`, err)}
		}
	}
	if _, ok := oc.mutation.PrecoDaOrdem(); !ok {
		return &ValidationError{Name: "preco_da_ordem", err: errors.New(`ent: missing required field "Ordem.preco_da_ordem"`)}
	}
	return nil
}

func (oc *OrdemCreate) sqlSave(ctx context.Context) (*Ordem, error) {
	if err := oc.check(); err != nil {
		return nil, err
	}
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	oc.mutation.id = &_node.ID
	oc.mutation.done = true
	return _node, nil
}

func (oc *OrdemCreate) createSpec() (*Ordem, *sqlgraph.CreateSpec) {
	var (
		_node = &Ordem{config: oc.config}
		_spec = sqlgraph.NewCreateSpec(ordem.Table, sqlgraph.NewFieldSpec(ordem.FieldID, field.TypeInt))
	)
	if id, ok := oc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := oc.mutation.DataOrdem(); ok {
		_spec.SetField(ordem.FieldDataOrdem, field.TypeTime, value)
		_node.DataOrdem = value
	}
	if value, ok := oc.mutation.Status(); ok {
		_spec.SetField(ordem.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := oc.mutation.PrecoDaOrdem(); ok {
		_spec.SetField(ordem.FieldPrecoDaOrdem, field.TypeFloat64, value)
		_node.PrecoDaOrdem = value
	}
	if nodes := oc.mutation.ProdutosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   ordem.ProdutosTable,
			Columns: ordem.ProdutosPrimaryKey,
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
	if nodes := oc.mutation.ClienteIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ordem.ClienteTable,
			Columns: []string{ordem.ClienteColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cliente.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.cliente_ordens = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   ordem.ItemsTable,
			Columns: []string{ordem.ItemsColumn},
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

// OrdemCreateBulk is the builder for creating many Ordem entities in bulk.
type OrdemCreateBulk struct {
	config
	err      error
	builders []*OrdemCreate
}

// Save creates the Ordem entities in the database.
func (ocb *OrdemCreateBulk) Save(ctx context.Context) ([]*Ordem, error) {
	if ocb.err != nil {
		return nil, ocb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Ordem, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrdemMutation)
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
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
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
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *OrdemCreateBulk) SaveX(ctx context.Context) []*Ordem {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *OrdemCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *OrdemCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}
