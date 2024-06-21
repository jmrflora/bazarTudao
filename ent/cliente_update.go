// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/jmrflora/bazarTudao/ent/cliente"
	"github.com/jmrflora/bazarTudao/ent/ordem"
	"github.com/jmrflora/bazarTudao/ent/predicate"
)

// ClienteUpdate is the builder for updating Cliente entities.
type ClienteUpdate struct {
	config
	hooks    []Hook
	mutation *ClienteMutation
}

// Where appends a list predicates to the ClienteUpdate builder.
func (cu *ClienteUpdate) Where(ps ...predicate.Cliente) *ClienteUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetEmail sets the "email" field.
func (cu *ClienteUpdate) SetEmail(s string) *ClienteUpdate {
	cu.mutation.SetEmail(s)
	return cu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (cu *ClienteUpdate) SetNillableEmail(s *string) *ClienteUpdate {
	if s != nil {
		cu.SetEmail(*s)
	}
	return cu
}

// SetNome sets the "nome" field.
func (cu *ClienteUpdate) SetNome(s string) *ClienteUpdate {
	cu.mutation.SetNome(s)
	return cu
}

// SetNillableNome sets the "nome" field if the given value is not nil.
func (cu *ClienteUpdate) SetNillableNome(s *string) *ClienteUpdate {
	if s != nil {
		cu.SetNome(*s)
	}
	return cu
}

// SetCpf sets the "cpf" field.
func (cu *ClienteUpdate) SetCpf(s string) *ClienteUpdate {
	cu.mutation.SetCpf(s)
	return cu
}

// SetNillableCpf sets the "cpf" field if the given value is not nil.
func (cu *ClienteUpdate) SetNillableCpf(s *string) *ClienteUpdate {
	if s != nil {
		cu.SetCpf(*s)
	}
	return cu
}

// SetTelefone sets the "telefone" field.
func (cu *ClienteUpdate) SetTelefone(s string) *ClienteUpdate {
	cu.mutation.SetTelefone(s)
	return cu
}

// SetNillableTelefone sets the "telefone" field if the given value is not nil.
func (cu *ClienteUpdate) SetNillableTelefone(s *string) *ClienteUpdate {
	if s != nil {
		cu.SetTelefone(*s)
	}
	return cu
}

// AddOrdenIDs adds the "ordens" edge to the Ordem entity by IDs.
func (cu *ClienteUpdate) AddOrdenIDs(ids ...int) *ClienteUpdate {
	cu.mutation.AddOrdenIDs(ids...)
	return cu
}

// AddOrdens adds the "ordens" edges to the Ordem entity.
func (cu *ClienteUpdate) AddOrdens(o ...*Ordem) *ClienteUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return cu.AddOrdenIDs(ids...)
}

// Mutation returns the ClienteMutation object of the builder.
func (cu *ClienteUpdate) Mutation() *ClienteMutation {
	return cu.mutation
}

// ClearOrdens clears all "ordens" edges to the Ordem entity.
func (cu *ClienteUpdate) ClearOrdens() *ClienteUpdate {
	cu.mutation.ClearOrdens()
	return cu
}

// RemoveOrdenIDs removes the "ordens" edge to Ordem entities by IDs.
func (cu *ClienteUpdate) RemoveOrdenIDs(ids ...int) *ClienteUpdate {
	cu.mutation.RemoveOrdenIDs(ids...)
	return cu
}

// RemoveOrdens removes "ordens" edges to Ordem entities.
func (cu *ClienteUpdate) RemoveOrdens(o ...*Ordem) *ClienteUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return cu.RemoveOrdenIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ClienteUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ClienteUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ClienteUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ClienteUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *ClienteUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(cliente.Table, cliente.Columns, sqlgraph.NewFieldSpec(cliente.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Email(); ok {
		_spec.SetField(cliente.FieldEmail, field.TypeString, value)
	}
	if value, ok := cu.mutation.Nome(); ok {
		_spec.SetField(cliente.FieldNome, field.TypeString, value)
	}
	if value, ok := cu.mutation.Cpf(); ok {
		_spec.SetField(cliente.FieldCpf, field.TypeString, value)
	}
	if value, ok := cu.mutation.Telefone(); ok {
		_spec.SetField(cliente.FieldTelefone, field.TypeString, value)
	}
	if cu.mutation.OrdensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cliente.OrdensTable,
			Columns: []string{cliente.OrdensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ordem.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedOrdensIDs(); len(nodes) > 0 && !cu.mutation.OrdensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cliente.OrdensTable,
			Columns: []string{cliente.OrdensColumn},
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
	if nodes := cu.mutation.OrdensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cliente.OrdensTable,
			Columns: []string{cliente.OrdensColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cliente.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ClienteUpdateOne is the builder for updating a single Cliente entity.
type ClienteUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ClienteMutation
}

// SetEmail sets the "email" field.
func (cuo *ClienteUpdateOne) SetEmail(s string) *ClienteUpdateOne {
	cuo.mutation.SetEmail(s)
	return cuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (cuo *ClienteUpdateOne) SetNillableEmail(s *string) *ClienteUpdateOne {
	if s != nil {
		cuo.SetEmail(*s)
	}
	return cuo
}

// SetNome sets the "nome" field.
func (cuo *ClienteUpdateOne) SetNome(s string) *ClienteUpdateOne {
	cuo.mutation.SetNome(s)
	return cuo
}

// SetNillableNome sets the "nome" field if the given value is not nil.
func (cuo *ClienteUpdateOne) SetNillableNome(s *string) *ClienteUpdateOne {
	if s != nil {
		cuo.SetNome(*s)
	}
	return cuo
}

// SetCpf sets the "cpf" field.
func (cuo *ClienteUpdateOne) SetCpf(s string) *ClienteUpdateOne {
	cuo.mutation.SetCpf(s)
	return cuo
}

// SetNillableCpf sets the "cpf" field if the given value is not nil.
func (cuo *ClienteUpdateOne) SetNillableCpf(s *string) *ClienteUpdateOne {
	if s != nil {
		cuo.SetCpf(*s)
	}
	return cuo
}

// SetTelefone sets the "telefone" field.
func (cuo *ClienteUpdateOne) SetTelefone(s string) *ClienteUpdateOne {
	cuo.mutation.SetTelefone(s)
	return cuo
}

// SetNillableTelefone sets the "telefone" field if the given value is not nil.
func (cuo *ClienteUpdateOne) SetNillableTelefone(s *string) *ClienteUpdateOne {
	if s != nil {
		cuo.SetTelefone(*s)
	}
	return cuo
}

// AddOrdenIDs adds the "ordens" edge to the Ordem entity by IDs.
func (cuo *ClienteUpdateOne) AddOrdenIDs(ids ...int) *ClienteUpdateOne {
	cuo.mutation.AddOrdenIDs(ids...)
	return cuo
}

// AddOrdens adds the "ordens" edges to the Ordem entity.
func (cuo *ClienteUpdateOne) AddOrdens(o ...*Ordem) *ClienteUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return cuo.AddOrdenIDs(ids...)
}

// Mutation returns the ClienteMutation object of the builder.
func (cuo *ClienteUpdateOne) Mutation() *ClienteMutation {
	return cuo.mutation
}

// ClearOrdens clears all "ordens" edges to the Ordem entity.
func (cuo *ClienteUpdateOne) ClearOrdens() *ClienteUpdateOne {
	cuo.mutation.ClearOrdens()
	return cuo
}

// RemoveOrdenIDs removes the "ordens" edge to Ordem entities by IDs.
func (cuo *ClienteUpdateOne) RemoveOrdenIDs(ids ...int) *ClienteUpdateOne {
	cuo.mutation.RemoveOrdenIDs(ids...)
	return cuo
}

// RemoveOrdens removes "ordens" edges to Ordem entities.
func (cuo *ClienteUpdateOne) RemoveOrdens(o ...*Ordem) *ClienteUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return cuo.RemoveOrdenIDs(ids...)
}

// Where appends a list predicates to the ClienteUpdate builder.
func (cuo *ClienteUpdateOne) Where(ps ...predicate.Cliente) *ClienteUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ClienteUpdateOne) Select(field string, fields ...string) *ClienteUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Cliente entity.
func (cuo *ClienteUpdateOne) Save(ctx context.Context) (*Cliente, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ClienteUpdateOne) SaveX(ctx context.Context) *Cliente {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ClienteUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ClienteUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *ClienteUpdateOne) sqlSave(ctx context.Context) (_node *Cliente, err error) {
	_spec := sqlgraph.NewUpdateSpec(cliente.Table, cliente.Columns, sqlgraph.NewFieldSpec(cliente.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Cliente.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cliente.FieldID)
		for _, f := range fields {
			if !cliente.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != cliente.FieldID {
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
	if value, ok := cuo.mutation.Email(); ok {
		_spec.SetField(cliente.FieldEmail, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Nome(); ok {
		_spec.SetField(cliente.FieldNome, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Cpf(); ok {
		_spec.SetField(cliente.FieldCpf, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Telefone(); ok {
		_spec.SetField(cliente.FieldTelefone, field.TypeString, value)
	}
	if cuo.mutation.OrdensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cliente.OrdensTable,
			Columns: []string{cliente.OrdensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ordem.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedOrdensIDs(); len(nodes) > 0 && !cuo.mutation.OrdensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cliente.OrdensTable,
			Columns: []string{cliente.OrdensColumn},
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
	if nodes := cuo.mutation.OrdensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cliente.OrdensTable,
			Columns: []string{cliente.OrdensColumn},
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
	_node = &Cliente{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cliente.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
