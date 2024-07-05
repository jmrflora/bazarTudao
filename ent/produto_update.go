// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/jmrflora/bazarTudao/ent/itemordem"
	"github.com/jmrflora/bazarTudao/ent/ordem"
	"github.com/jmrflora/bazarTudao/ent/predicate"
	"github.com/jmrflora/bazarTudao/ent/produto"
	"github.com/jmrflora/bazarTudao/ent/stock"
)

// ProdutoUpdate is the builder for updating Produto entities.
type ProdutoUpdate struct {
	config
	hooks    []Hook
	mutation *ProdutoMutation
}

// Where appends a list predicates to the ProdutoUpdate builder.
func (pu *ProdutoUpdate) Where(ps ...predicate.Produto) *ProdutoUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetSku sets the "sku" field.
func (pu *ProdutoUpdate) SetSku(s string) *ProdutoUpdate {
	pu.mutation.SetSku(s)
	return pu
}

// SetNillableSku sets the "sku" field if the given value is not nil.
func (pu *ProdutoUpdate) SetNillableSku(s *string) *ProdutoUpdate {
	if s != nil {
		pu.SetSku(*s)
	}
	return pu
}

// SetNome sets the "nome" field.
func (pu *ProdutoUpdate) SetNome(s string) *ProdutoUpdate {
	pu.mutation.SetNome(s)
	return pu
}

// SetNillableNome sets the "nome" field if the given value is not nil.
func (pu *ProdutoUpdate) SetNillableNome(s *string) *ProdutoUpdate {
	if s != nil {
		pu.SetNome(*s)
	}
	return pu
}

// SetQuantNoEstoque sets the "quant_no_estoque" field.
func (pu *ProdutoUpdate) SetQuantNoEstoque(i int) *ProdutoUpdate {
	pu.mutation.ResetQuantNoEstoque()
	pu.mutation.SetQuantNoEstoque(i)
	return pu
}

// SetNillableQuantNoEstoque sets the "quant_no_estoque" field if the given value is not nil.
func (pu *ProdutoUpdate) SetNillableQuantNoEstoque(i *int) *ProdutoUpdate {
	if i != nil {
		pu.SetQuantNoEstoque(*i)
	}
	return pu
}

// AddQuantNoEstoque adds i to the "quant_no_estoque" field.
func (pu *ProdutoUpdate) AddQuantNoEstoque(i int) *ProdutoUpdate {
	pu.mutation.AddQuantNoEstoque(i)
	return pu
}

// AddOrdenIDs adds the "ordens" edge to the Ordem entity by IDs.
func (pu *ProdutoUpdate) AddOrdenIDs(ids ...int) *ProdutoUpdate {
	pu.mutation.AddOrdenIDs(ids...)
	return pu
}

// AddOrdens adds the "ordens" edges to the Ordem entity.
func (pu *ProdutoUpdate) AddOrdens(o ...*Ordem) *ProdutoUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return pu.AddOrdenIDs(ids...)
}

// SetStockID sets the "stock" edge to the Stock entity by ID.
func (pu *ProdutoUpdate) SetStockID(id int) *ProdutoUpdate {
	pu.mutation.SetStockID(id)
	return pu
}

// SetNillableStockID sets the "stock" edge to the Stock entity by ID if the given value is not nil.
func (pu *ProdutoUpdate) SetNillableStockID(id *int) *ProdutoUpdate {
	if id != nil {
		pu = pu.SetStockID(*id)
	}
	return pu
}

// SetStock sets the "stock" edge to the Stock entity.
func (pu *ProdutoUpdate) SetStock(s *Stock) *ProdutoUpdate {
	return pu.SetStockID(s.ID)
}

// AddItenIDs adds the "itens" edge to the ItemOrdem entity by IDs.
func (pu *ProdutoUpdate) AddItenIDs(ids ...int) *ProdutoUpdate {
	pu.mutation.AddItenIDs(ids...)
	return pu
}

// AddItens adds the "itens" edges to the ItemOrdem entity.
func (pu *ProdutoUpdate) AddItens(i ...*ItemOrdem) *ProdutoUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return pu.AddItenIDs(ids...)
}

// Mutation returns the ProdutoMutation object of the builder.
func (pu *ProdutoUpdate) Mutation() *ProdutoMutation {
	return pu.mutation
}

// ClearOrdens clears all "ordens" edges to the Ordem entity.
func (pu *ProdutoUpdate) ClearOrdens() *ProdutoUpdate {
	pu.mutation.ClearOrdens()
	return pu
}

// RemoveOrdenIDs removes the "ordens" edge to Ordem entities by IDs.
func (pu *ProdutoUpdate) RemoveOrdenIDs(ids ...int) *ProdutoUpdate {
	pu.mutation.RemoveOrdenIDs(ids...)
	return pu
}

// RemoveOrdens removes "ordens" edges to Ordem entities.
func (pu *ProdutoUpdate) RemoveOrdens(o ...*Ordem) *ProdutoUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return pu.RemoveOrdenIDs(ids...)
}

// ClearStock clears the "stock" edge to the Stock entity.
func (pu *ProdutoUpdate) ClearStock() *ProdutoUpdate {
	pu.mutation.ClearStock()
	return pu
}

// ClearItens clears all "itens" edges to the ItemOrdem entity.
func (pu *ProdutoUpdate) ClearItens() *ProdutoUpdate {
	pu.mutation.ClearItens()
	return pu
}

// RemoveItenIDs removes the "itens" edge to ItemOrdem entities by IDs.
func (pu *ProdutoUpdate) RemoveItenIDs(ids ...int) *ProdutoUpdate {
	pu.mutation.RemoveItenIDs(ids...)
	return pu
}

// RemoveItens removes "itens" edges to ItemOrdem entities.
func (pu *ProdutoUpdate) RemoveItens(i ...*ItemOrdem) *ProdutoUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return pu.RemoveItenIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProdutoUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProdutoUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProdutoUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProdutoUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *ProdutoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(produto.Table, produto.Columns, sqlgraph.NewFieldSpec(produto.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Sku(); ok {
		_spec.SetField(produto.FieldSku, field.TypeString, value)
	}
	if value, ok := pu.mutation.Nome(); ok {
		_spec.SetField(produto.FieldNome, field.TypeString, value)
	}
	if value, ok := pu.mutation.QuantNoEstoque(); ok {
		_spec.SetField(produto.FieldQuantNoEstoque, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedQuantNoEstoque(); ok {
		_spec.AddField(produto.FieldQuantNoEstoque, field.TypeInt, value)
	}
	if pu.mutation.OrdensCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedOrdensIDs(); len(nodes) > 0 && !pu.mutation.OrdensCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.OrdensIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.StockCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   produto.StockTable,
			Columns: []string{produto.StockColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(stock.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.StockIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.ItensCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedItensIDs(); len(nodes) > 0 && !pu.mutation.ItensCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ItensIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{produto.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ProdutoUpdateOne is the builder for updating a single Produto entity.
type ProdutoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProdutoMutation
}

// SetSku sets the "sku" field.
func (puo *ProdutoUpdateOne) SetSku(s string) *ProdutoUpdateOne {
	puo.mutation.SetSku(s)
	return puo
}

// SetNillableSku sets the "sku" field if the given value is not nil.
func (puo *ProdutoUpdateOne) SetNillableSku(s *string) *ProdutoUpdateOne {
	if s != nil {
		puo.SetSku(*s)
	}
	return puo
}

// SetNome sets the "nome" field.
func (puo *ProdutoUpdateOne) SetNome(s string) *ProdutoUpdateOne {
	puo.mutation.SetNome(s)
	return puo
}

// SetNillableNome sets the "nome" field if the given value is not nil.
func (puo *ProdutoUpdateOne) SetNillableNome(s *string) *ProdutoUpdateOne {
	if s != nil {
		puo.SetNome(*s)
	}
	return puo
}

// SetQuantNoEstoque sets the "quant_no_estoque" field.
func (puo *ProdutoUpdateOne) SetQuantNoEstoque(i int) *ProdutoUpdateOne {
	puo.mutation.ResetQuantNoEstoque()
	puo.mutation.SetQuantNoEstoque(i)
	return puo
}

// SetNillableQuantNoEstoque sets the "quant_no_estoque" field if the given value is not nil.
func (puo *ProdutoUpdateOne) SetNillableQuantNoEstoque(i *int) *ProdutoUpdateOne {
	if i != nil {
		puo.SetQuantNoEstoque(*i)
	}
	return puo
}

// AddQuantNoEstoque adds i to the "quant_no_estoque" field.
func (puo *ProdutoUpdateOne) AddQuantNoEstoque(i int) *ProdutoUpdateOne {
	puo.mutation.AddQuantNoEstoque(i)
	return puo
}

// AddOrdenIDs adds the "ordens" edge to the Ordem entity by IDs.
func (puo *ProdutoUpdateOne) AddOrdenIDs(ids ...int) *ProdutoUpdateOne {
	puo.mutation.AddOrdenIDs(ids...)
	return puo
}

// AddOrdens adds the "ordens" edges to the Ordem entity.
func (puo *ProdutoUpdateOne) AddOrdens(o ...*Ordem) *ProdutoUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return puo.AddOrdenIDs(ids...)
}

// SetStockID sets the "stock" edge to the Stock entity by ID.
func (puo *ProdutoUpdateOne) SetStockID(id int) *ProdutoUpdateOne {
	puo.mutation.SetStockID(id)
	return puo
}

// SetNillableStockID sets the "stock" edge to the Stock entity by ID if the given value is not nil.
func (puo *ProdutoUpdateOne) SetNillableStockID(id *int) *ProdutoUpdateOne {
	if id != nil {
		puo = puo.SetStockID(*id)
	}
	return puo
}

// SetStock sets the "stock" edge to the Stock entity.
func (puo *ProdutoUpdateOne) SetStock(s *Stock) *ProdutoUpdateOne {
	return puo.SetStockID(s.ID)
}

// AddItenIDs adds the "itens" edge to the ItemOrdem entity by IDs.
func (puo *ProdutoUpdateOne) AddItenIDs(ids ...int) *ProdutoUpdateOne {
	puo.mutation.AddItenIDs(ids...)
	return puo
}

// AddItens adds the "itens" edges to the ItemOrdem entity.
func (puo *ProdutoUpdateOne) AddItens(i ...*ItemOrdem) *ProdutoUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return puo.AddItenIDs(ids...)
}

// Mutation returns the ProdutoMutation object of the builder.
func (puo *ProdutoUpdateOne) Mutation() *ProdutoMutation {
	return puo.mutation
}

// ClearOrdens clears all "ordens" edges to the Ordem entity.
func (puo *ProdutoUpdateOne) ClearOrdens() *ProdutoUpdateOne {
	puo.mutation.ClearOrdens()
	return puo
}

// RemoveOrdenIDs removes the "ordens" edge to Ordem entities by IDs.
func (puo *ProdutoUpdateOne) RemoveOrdenIDs(ids ...int) *ProdutoUpdateOne {
	puo.mutation.RemoveOrdenIDs(ids...)
	return puo
}

// RemoveOrdens removes "ordens" edges to Ordem entities.
func (puo *ProdutoUpdateOne) RemoveOrdens(o ...*Ordem) *ProdutoUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return puo.RemoveOrdenIDs(ids...)
}

// ClearStock clears the "stock" edge to the Stock entity.
func (puo *ProdutoUpdateOne) ClearStock() *ProdutoUpdateOne {
	puo.mutation.ClearStock()
	return puo
}

// ClearItens clears all "itens" edges to the ItemOrdem entity.
func (puo *ProdutoUpdateOne) ClearItens() *ProdutoUpdateOne {
	puo.mutation.ClearItens()
	return puo
}

// RemoveItenIDs removes the "itens" edge to ItemOrdem entities by IDs.
func (puo *ProdutoUpdateOne) RemoveItenIDs(ids ...int) *ProdutoUpdateOne {
	puo.mutation.RemoveItenIDs(ids...)
	return puo
}

// RemoveItens removes "itens" edges to ItemOrdem entities.
func (puo *ProdutoUpdateOne) RemoveItens(i ...*ItemOrdem) *ProdutoUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return puo.RemoveItenIDs(ids...)
}

// Where appends a list predicates to the ProdutoUpdate builder.
func (puo *ProdutoUpdateOne) Where(ps ...predicate.Produto) *ProdutoUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProdutoUpdateOne) Select(field string, fields ...string) *ProdutoUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Produto entity.
func (puo *ProdutoUpdateOne) Save(ctx context.Context) (*Produto, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProdutoUpdateOne) SaveX(ctx context.Context) *Produto {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProdutoUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProdutoUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *ProdutoUpdateOne) sqlSave(ctx context.Context) (_node *Produto, err error) {
	_spec := sqlgraph.NewUpdateSpec(produto.Table, produto.Columns, sqlgraph.NewFieldSpec(produto.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Produto.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, produto.FieldID)
		for _, f := range fields {
			if !produto.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != produto.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Sku(); ok {
		_spec.SetField(produto.FieldSku, field.TypeString, value)
	}
	if value, ok := puo.mutation.Nome(); ok {
		_spec.SetField(produto.FieldNome, field.TypeString, value)
	}
	if value, ok := puo.mutation.QuantNoEstoque(); ok {
		_spec.SetField(produto.FieldQuantNoEstoque, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedQuantNoEstoque(); ok {
		_spec.AddField(produto.FieldQuantNoEstoque, field.TypeInt, value)
	}
	if puo.mutation.OrdensCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedOrdensIDs(); len(nodes) > 0 && !puo.mutation.OrdensCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.OrdensIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.StockCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   produto.StockTable,
			Columns: []string{produto.StockColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(stock.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.StockIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.ItensCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedItensIDs(); len(nodes) > 0 && !puo.mutation.ItensCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ItensIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Produto{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{produto.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
