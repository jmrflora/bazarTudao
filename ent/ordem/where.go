// Code generated by ent, DO NOT EDIT.

package ordem

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/jmrflora/bazarTudao/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Ordem {
	return predicate.Ordem(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Ordem {
	return predicate.Ordem(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Ordem {
	return predicate.Ordem(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Ordem {
	return predicate.Ordem(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Ordem {
	return predicate.Ordem(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Ordem {
	return predicate.Ordem(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Ordem {
	return predicate.Ordem(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Ordem {
	return predicate.Ordem(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Ordem {
	return predicate.Ordem(sql.FieldLTE(FieldID, id))
}

// DataOrdem applies equality check predicate on the "data_ordem" field. It's identical to DataOrdemEQ.
func DataOrdem(v time.Time) predicate.Ordem {
	return predicate.Ordem(sql.FieldEQ(FieldDataOrdem, v))
}

// DataOrdemEQ applies the EQ predicate on the "data_ordem" field.
func DataOrdemEQ(v time.Time) predicate.Ordem {
	return predicate.Ordem(sql.FieldEQ(FieldDataOrdem, v))
}

// DataOrdemNEQ applies the NEQ predicate on the "data_ordem" field.
func DataOrdemNEQ(v time.Time) predicate.Ordem {
	return predicate.Ordem(sql.FieldNEQ(FieldDataOrdem, v))
}

// DataOrdemIn applies the In predicate on the "data_ordem" field.
func DataOrdemIn(vs ...time.Time) predicate.Ordem {
	return predicate.Ordem(sql.FieldIn(FieldDataOrdem, vs...))
}

// DataOrdemNotIn applies the NotIn predicate on the "data_ordem" field.
func DataOrdemNotIn(vs ...time.Time) predicate.Ordem {
	return predicate.Ordem(sql.FieldNotIn(FieldDataOrdem, vs...))
}

// DataOrdemGT applies the GT predicate on the "data_ordem" field.
func DataOrdemGT(v time.Time) predicate.Ordem {
	return predicate.Ordem(sql.FieldGT(FieldDataOrdem, v))
}

// DataOrdemGTE applies the GTE predicate on the "data_ordem" field.
func DataOrdemGTE(v time.Time) predicate.Ordem {
	return predicate.Ordem(sql.FieldGTE(FieldDataOrdem, v))
}

// DataOrdemLT applies the LT predicate on the "data_ordem" field.
func DataOrdemLT(v time.Time) predicate.Ordem {
	return predicate.Ordem(sql.FieldLT(FieldDataOrdem, v))
}

// DataOrdemLTE applies the LTE predicate on the "data_ordem" field.
func DataOrdemLTE(v time.Time) predicate.Ordem {
	return predicate.Ordem(sql.FieldLTE(FieldDataOrdem, v))
}

// HasProdutos applies the HasEdge predicate on the "produtos" edge.
func HasProdutos() predicate.Ordem {
	return predicate.Ordem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ProdutosTable, ProdutosPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProdutosWith applies the HasEdge predicate on the "produtos" edge with a given conditions (other predicates).
func HasProdutosWith(preds ...predicate.Produto) predicate.Ordem {
	return predicate.Ordem(func(s *sql.Selector) {
		step := newProdutosStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasClientes applies the HasEdge predicate on the "clientes" edge.
func HasClientes() predicate.Ordem {
	return predicate.Ordem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ClientesTable, ClientesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasClientesWith applies the HasEdge predicate on the "clientes" edge with a given conditions (other predicates).
func HasClientesWith(preds ...predicate.Cliente) predicate.Ordem {
	return predicate.Ordem(func(s *sql.Selector) {
		step := newClientesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasItems applies the HasEdge predicate on the "items" edge.
func HasItems() predicate.Ordem {
	return predicate.Ordem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, ItemsTable, ItemsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasItemsWith applies the HasEdge predicate on the "items" edge with a given conditions (other predicates).
func HasItemsWith(preds ...predicate.ItemOrdem) predicate.Ordem {
	return predicate.Ordem(func(s *sql.Selector) {
		step := newItemsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Ordem) predicate.Ordem {
	return predicate.Ordem(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Ordem) predicate.Ordem {
	return predicate.Ordem(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Ordem) predicate.Ordem {
	return predicate.Ordem(sql.NotPredicates(p))
}
