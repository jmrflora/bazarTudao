// Code generated by ent, DO NOT EDIT.

package stock

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the stock type in the database.
	Label = "stock"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDataMovimento holds the string denoting the data_movimento field in the database.
	FieldDataMovimento = "data_movimento"
	// FieldQuantidade holds the string denoting the quantidade field in the database.
	FieldQuantidade = "quantidade"
	// EdgeProdutos holds the string denoting the produtos edge name in mutations.
	EdgeProdutos = "produtos"
	// Table holds the table name of the stock in the database.
	Table = "stocks"
	// ProdutosTable is the table that holds the produtos relation/edge.
	ProdutosTable = "produtos"
	// ProdutosInverseTable is the table name for the Produto entity.
	// It exists in this package in order to avoid circular dependency with the "produto" package.
	ProdutosInverseTable = "produtos"
	// ProdutosColumn is the table column denoting the produtos relation/edge.
	ProdutosColumn = "stock_produtos"
)

// Columns holds all SQL columns for stock fields.
var Columns = []string{
	FieldID,
	FieldDataMovimento,
	FieldQuantidade,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultDataMovimento holds the default value on creation for the "data_movimento" field.
	DefaultDataMovimento func() time.Time
)

// OrderOption defines the ordering options for the Stock queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDataMovimento orders the results by the data_movimento field.
func ByDataMovimento(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDataMovimento, opts...).ToFunc()
}

// ByQuantidade orders the results by the quantidade field.
func ByQuantidade(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldQuantidade, opts...).ToFunc()
}

// ByProdutosCount orders the results by produtos count.
func ByProdutosCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProdutosStep(), opts...)
	}
}

// ByProdutos orders the results by produtos terms.
func ByProdutos(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProdutosStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newProdutosStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProdutosInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ProdutosTable, ProdutosColumn),
	)
}