// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/jmrflora/bazarTudao/ent/produto"
	"github.com/jmrflora/bazarTudao/ent/stock"
)

// Produto is the model entity for the Produto schema.
type Produto struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Sku holds the value of the "sku" field.
	Sku string `json:"sku,omitempty"`
	// Nome holds the value of the "nome" field.
	Nome string `json:"nome,omitempty"`
	// QuantNoEstoque holds the value of the "quant_no_estoque" field.
	QuantNoEstoque int `json:"quant_no_estoque,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProdutoQuery when eager-loading is set.
	Edges          ProdutoEdges `json:"edges"`
	stock_produtos *int
	selectValues   sql.SelectValues
}

// ProdutoEdges holds the relations/edges for other nodes in the graph.
type ProdutoEdges struct {
	// Ordens holds the value of the ordens edge.
	Ordens []*Ordem `json:"ordens,omitempty"`
	// Stock holds the value of the stock edge.
	Stock *Stock `json:"stock,omitempty"`
	// Itens holds the value of the itens edge.
	Itens []*ItemOrdem `json:"itens,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// OrdensOrErr returns the Ordens value or an error if the edge
// was not loaded in eager-loading.
func (e ProdutoEdges) OrdensOrErr() ([]*Ordem, error) {
	if e.loadedTypes[0] {
		return e.Ordens, nil
	}
	return nil, &NotLoadedError{edge: "ordens"}
}

// StockOrErr returns the Stock value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProdutoEdges) StockOrErr() (*Stock, error) {
	if e.Stock != nil {
		return e.Stock, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: stock.Label}
	}
	return nil, &NotLoadedError{edge: "stock"}
}

// ItensOrErr returns the Itens value or an error if the edge
// was not loaded in eager-loading.
func (e ProdutoEdges) ItensOrErr() ([]*ItemOrdem, error) {
	if e.loadedTypes[2] {
		return e.Itens, nil
	}
	return nil, &NotLoadedError{edge: "itens"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Produto) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case produto.FieldID, produto.FieldQuantNoEstoque:
			values[i] = new(sql.NullInt64)
		case produto.FieldSku, produto.FieldNome:
			values[i] = new(sql.NullString)
		case produto.ForeignKeys[0]: // stock_produtos
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Produto fields.
func (pr *Produto) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case produto.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int(value.Int64)
		case produto.FieldSku:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sku", values[i])
			} else if value.Valid {
				pr.Sku = value.String
			}
		case produto.FieldNome:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nome", values[i])
			} else if value.Valid {
				pr.Nome = value.String
			}
		case produto.FieldQuantNoEstoque:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field quant_no_estoque", values[i])
			} else if value.Valid {
				pr.QuantNoEstoque = int(value.Int64)
			}
		case produto.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field stock_produtos", value)
			} else if value.Valid {
				pr.stock_produtos = new(int)
				*pr.stock_produtos = int(value.Int64)
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Produto.
// This includes values selected through modifiers, order, etc.
func (pr *Produto) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// QueryOrdens queries the "ordens" edge of the Produto entity.
func (pr *Produto) QueryOrdens() *OrdemQuery {
	return NewProdutoClient(pr.config).QueryOrdens(pr)
}

// QueryStock queries the "stock" edge of the Produto entity.
func (pr *Produto) QueryStock() *StockQuery {
	return NewProdutoClient(pr.config).QueryStock(pr)
}

// QueryItens queries the "itens" edge of the Produto entity.
func (pr *Produto) QueryItens() *ItemOrdemQuery {
	return NewProdutoClient(pr.config).QueryItens(pr)
}

// Update returns a builder for updating this Produto.
// Note that you need to call Produto.Unwrap() before calling this method if this Produto
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Produto) Update() *ProdutoUpdateOne {
	return NewProdutoClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Produto entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Produto) Unwrap() *Produto {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Produto is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Produto) String() string {
	var builder strings.Builder
	builder.WriteString("Produto(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("sku=")
	builder.WriteString(pr.Sku)
	builder.WriteString(", ")
	builder.WriteString("nome=")
	builder.WriteString(pr.Nome)
	builder.WriteString(", ")
	builder.WriteString("quant_no_estoque=")
	builder.WriteString(fmt.Sprintf("%v", pr.QuantNoEstoque))
	builder.WriteByte(')')
	return builder.String()
}

// Produtos is a parsable slice of Produto.
type Produtos []*Produto
