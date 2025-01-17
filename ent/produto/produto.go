// Code generated by ent, DO NOT EDIT.

package produto

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the produto type in the database.
	Label = "produto"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSku holds the string denoting the sku field in the database.
	FieldSku = "sku"
	// FieldNome holds the string denoting the nome field in the database.
	FieldNome = "nome"
	// FieldQuantNoEstoque holds the string denoting the quant_no_estoque field in the database.
	FieldQuantNoEstoque = "quant_no_estoque"
	// EdgeOrdens holds the string denoting the ordens edge name in mutations.
	EdgeOrdens = "ordens"
	// EdgeStock holds the string denoting the stock edge name in mutations.
	EdgeStock = "stock"
	// EdgeItens holds the string denoting the itens edge name in mutations.
	EdgeItens = "itens"
	// Table holds the table name of the produto in the database.
	Table = "produtos"
	// OrdensTable is the table that holds the ordens relation/edge. The primary key declared below.
	OrdensTable = "item_ordems"
	// OrdensInverseTable is the table name for the Ordem entity.
	// It exists in this package in order to avoid circular dependency with the "ordem" package.
	OrdensInverseTable = "ordems"
	// StockTable is the table that holds the stock relation/edge.
	StockTable = "stocks"
	// StockInverseTable is the table name for the Stock entity.
	// It exists in this package in order to avoid circular dependency with the "stock" package.
	StockInverseTable = "stocks"
	// StockColumn is the table column denoting the stock relation/edge.
	StockColumn = "stock_produtos"
	// ItensTable is the table that holds the itens relation/edge.
	ItensTable = "item_ordems"
	// ItensInverseTable is the table name for the ItemOrdem entity.
	// It exists in this package in order to avoid circular dependency with the "itemordem" package.
	ItensInverseTable = "item_ordems"
	// ItensColumn is the table column denoting the itens relation/edge.
	ItensColumn = "produto_id"
)

// Columns holds all SQL columns for produto fields.
var Columns = []string{
	FieldID,
	FieldSku,
	FieldNome,
	FieldQuantNoEstoque,
}

var (
	// OrdensPrimaryKey and OrdensColumn2 are the table columns denoting the
	// primary key for the ordens relation (M2M).
	OrdensPrimaryKey = []string{"ordem_id", "produto_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Produto queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// BySku orders the results by the sku field.
func BySku(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSku, opts...).ToFunc()
}

// ByNome orders the results by the nome field.
func ByNome(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNome, opts...).ToFunc()
}

// ByQuantNoEstoque orders the results by the quant_no_estoque field.
func ByQuantNoEstoque(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldQuantNoEstoque, opts...).ToFunc()
}

// ByOrdensCount orders the results by ordens count.
func ByOrdensCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOrdensStep(), opts...)
	}
}

// ByOrdens orders the results by ordens terms.
func ByOrdens(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrdensStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByStockCount orders the results by stock count.
func ByStockCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newStockStep(), opts...)
	}
}

// ByStock orders the results by stock terms.
func ByStock(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStockStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByItensCount orders the results by itens count.
func ByItensCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newItensStep(), opts...)
	}
}

// ByItens orders the results by itens terms.
func ByItens(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newItensStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newOrdensStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrdensInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, OrdensTable, OrdensPrimaryKey...),
	)
}
func newStockStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StockInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, StockTable, StockColumn),
	)
}
func newItensStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ItensInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, ItensTable, ItensColumn),
	)
}
