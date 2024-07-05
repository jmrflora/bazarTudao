package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Produto holds the schema definition for the Produto entity.
type Produto struct {
	ent.Schema
}

// Fields of the Produto.
func (Produto) Fields() []ent.Field {
	return []ent.Field{
		field.String("sku").Unique(),
		field.String("nome"),
		field.Int("quant_no_estoque"),
	}
}

// Edges of the Produto.
func (Produto) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("ordens", Ordem.Type).Ref("produtos").Through("itens", ItemOrdem.Type),
		edge.From("stock", Stock.Type).Ref("produtos").Unique(),
	}
}
