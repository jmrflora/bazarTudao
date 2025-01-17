// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/jmrflora/bazarTudao/ent/cliente"
)

// Cliente is the model entity for the Cliente schema.
type Cliente struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Nome holds the value of the "nome" field.
	Nome string `json:"nome,omitempty"`
	// Cpf holds the value of the "cpf" field.
	Cpf string `json:"cpf,omitempty"`
	// Telefone holds the value of the "telefone" field.
	Telefone string `json:"telefone,omitempty"`
	// EnderecoEntrega holds the value of the "endereco_entrega" field.
	EnderecoEntrega string `json:"endereco_entrega,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ClienteQuery when eager-loading is set.
	Edges        ClienteEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ClienteEdges holds the relations/edges for other nodes in the graph.
type ClienteEdges struct {
	// Ordens holds the value of the ordens edge.
	Ordens []*Ordem `json:"ordens,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OrdensOrErr returns the Ordens value or an error if the edge
// was not loaded in eager-loading.
func (e ClienteEdges) OrdensOrErr() ([]*Ordem, error) {
	if e.loadedTypes[0] {
		return e.Ordens, nil
	}
	return nil, &NotLoadedError{edge: "ordens"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Cliente) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case cliente.FieldID:
			values[i] = new(sql.NullInt64)
		case cliente.FieldEmail, cliente.FieldNome, cliente.FieldCpf, cliente.FieldTelefone, cliente.FieldEnderecoEntrega:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Cliente fields.
func (c *Cliente) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case cliente.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case cliente.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				c.Email = value.String
			}
		case cliente.FieldNome:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nome", values[i])
			} else if value.Valid {
				c.Nome = value.String
			}
		case cliente.FieldCpf:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cpf", values[i])
			} else if value.Valid {
				c.Cpf = value.String
			}
		case cliente.FieldTelefone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field telefone", values[i])
			} else if value.Valid {
				c.Telefone = value.String
			}
		case cliente.FieldEnderecoEntrega:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field endereco_entrega", values[i])
			} else if value.Valid {
				c.EnderecoEntrega = value.String
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Cliente.
// This includes values selected through modifiers, order, etc.
func (c *Cliente) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryOrdens queries the "ordens" edge of the Cliente entity.
func (c *Cliente) QueryOrdens() *OrdemQuery {
	return NewClienteClient(c.config).QueryOrdens(c)
}

// Update returns a builder for updating this Cliente.
// Note that you need to call Cliente.Unwrap() before calling this method if this Cliente
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Cliente) Update() *ClienteUpdateOne {
	return NewClienteClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Cliente entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Cliente) Unwrap() *Cliente {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Cliente is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Cliente) String() string {
	var builder strings.Builder
	builder.WriteString("Cliente(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("email=")
	builder.WriteString(c.Email)
	builder.WriteString(", ")
	builder.WriteString("nome=")
	builder.WriteString(c.Nome)
	builder.WriteString(", ")
	builder.WriteString("cpf=")
	builder.WriteString(c.Cpf)
	builder.WriteString(", ")
	builder.WriteString("telefone=")
	builder.WriteString(c.Telefone)
	builder.WriteString(", ")
	builder.WriteString("endereco_entrega=")
	builder.WriteString(c.EnderecoEntrega)
	builder.WriteByte(')')
	return builder.String()
}

// Clientes is a parsable slice of Cliente.
type Clientes []*Cliente
