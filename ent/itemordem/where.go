// Code generated by ent, DO NOT EDIT.

package itemordem

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/jmrflora/bazarTudao/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.FieldLTE(FieldID, id))
}

// Quantidade applies equality check predicate on the "quantidade" field. It's identical to QuantidadeEQ.
func Quantidade(v int) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.FieldEQ(FieldQuantidade, v))
}

// Preco applies equality check predicate on the "preco" field. It's identical to PrecoEQ.
func Preco(v float64) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.FieldEQ(FieldPreco, v))
}

// OrdemID applies equality check predicate on the "ordem_id" field. It's identical to OrdemIDEQ.
func OrdemID(v int) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.FieldEQ(FieldOrdemID, v))
}

// ProdutoID applies equality check predicate on the "produto_id" field. It's identical to ProdutoIDEQ.
func ProdutoID(v int) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.FieldEQ(FieldProdutoID, v))
}

// QuantidadeEQ applies the EQ predicate on the "quantidade" field.
func QuantidadeEQ(v int) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.FieldEQ(FieldQuantidade, v))
}

// QuantidadeNEQ applies the NEQ predicate on the "quantidade" field.
func QuantidadeNEQ(v int) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.FieldNEQ(FieldQuantidade, v))
}

// QuantidadeIn applies the In predicate on the "quantidade" field.
func QuantidadeIn(vs ...int) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.FieldIn(FieldQuantidade, vs...))
}

// QuantidadeNotIn applies the NotIn predicate on the "quantidade" field.
func QuantidadeNotIn(vs ...int) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.FieldNotIn(FieldQuantidade, vs...))
}

// QuantidadeGT applies the GT predicate on the "quantidade" field.
func QuantidadeGT(v int) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.FieldGT(FieldQuantidade, v))
}

// QuantidadeGTE applies the GTE predicate on the "quantidade" field.
func QuantidadeGTE(v int) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.FieldGTE(FieldQuantidade, v))
}

// QuantidadeLT applies the LT predicate on the "quantidade" field.
func QuantidadeLT(v int) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.FieldLT(FieldQuantidade, v))
}

// QuantidadeLTE applies the LTE predicate on the "quantidade" field.
func QuantidadeLTE(v int) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.FieldLTE(FieldQuantidade, v))
}

// PrecoEQ applies the EQ predicate on the "preco" field.
func PrecoEQ(v float64) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.FieldEQ(FieldPreco, v))
}

// PrecoNEQ applies the NEQ predicate on the "preco" field.
func PrecoNEQ(v float64) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.FieldNEQ(FieldPreco, v))
}

// PrecoIn applies the In predicate on the "preco" field.
func PrecoIn(vs ...float64) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.FieldIn(FieldPreco, vs...))
}

// PrecoNotIn applies the NotIn predicate on the "preco" field.
func PrecoNotIn(vs ...float64) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.FieldNotIn(FieldPreco, vs...))
}

// PrecoGT applies the GT predicate on the "preco" field.
func PrecoGT(v float64) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.FieldGT(FieldPreco, v))
}

// PrecoGTE applies the GTE predicate on the "preco" field.
func PrecoGTE(v float64) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.FieldGTE(FieldPreco, v))
}

// PrecoLT applies the LT predicate on the "preco" field.
func PrecoLT(v float64) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.FieldLT(FieldPreco, v))
}

// PrecoLTE applies the LTE predicate on the "preco" field.
func PrecoLTE(v float64) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.FieldLTE(FieldPreco, v))
}

// OrdemIDEQ applies the EQ predicate on the "ordem_id" field.
func OrdemIDEQ(v int) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.FieldEQ(FieldOrdemID, v))
}

// OrdemIDNEQ applies the NEQ predicate on the "ordem_id" field.
func OrdemIDNEQ(v int) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.FieldNEQ(FieldOrdemID, v))
}

// OrdemIDIn applies the In predicate on the "ordem_id" field.
func OrdemIDIn(vs ...int) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.FieldIn(FieldOrdemID, vs...))
}

// OrdemIDNotIn applies the NotIn predicate on the "ordem_id" field.
func OrdemIDNotIn(vs ...int) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.FieldNotIn(FieldOrdemID, vs...))
}

// ProdutoIDEQ applies the EQ predicate on the "produto_id" field.
func ProdutoIDEQ(v int) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.FieldEQ(FieldProdutoID, v))
}

// ProdutoIDNEQ applies the NEQ predicate on the "produto_id" field.
func ProdutoIDNEQ(v int) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.FieldNEQ(FieldProdutoID, v))
}

// ProdutoIDIn applies the In predicate on the "produto_id" field.
func ProdutoIDIn(vs ...int) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.FieldIn(FieldProdutoID, vs...))
}

// ProdutoIDNotIn applies the NotIn predicate on the "produto_id" field.
func ProdutoIDNotIn(vs ...int) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.FieldNotIn(FieldProdutoID, vs...))
}

// HasOrdem applies the HasEdge predicate on the "ordem" edge.
func HasOrdem() predicate.ItemOrdem {
	return predicate.ItemOrdem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OrdemTable, OrdemColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrdemWith applies the HasEdge predicate on the "ordem" edge with a given conditions (other predicates).
func HasOrdemWith(preds ...predicate.Ordem) predicate.ItemOrdem {
	return predicate.ItemOrdem(func(s *sql.Selector) {
		step := newOrdemStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProduto applies the HasEdge predicate on the "produto" edge.
func HasProduto() predicate.ItemOrdem {
	return predicate.ItemOrdem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ProdutoTable, ProdutoColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProdutoWith applies the HasEdge predicate on the "produto" edge with a given conditions (other predicates).
func HasProdutoWith(preds ...predicate.Produto) predicate.ItemOrdem {
	return predicate.ItemOrdem(func(s *sql.Selector) {
		step := newProdutoStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ItemOrdem) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ItemOrdem) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ItemOrdem) predicate.ItemOrdem {
	return predicate.ItemOrdem(sql.NotPredicates(p))
}