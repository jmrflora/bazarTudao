package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ItemOrdem holds the schema definition for the ItemOrdem entity.
type ItemOrdem struct {
	ent.Schema
}

// Fields of the ItemOrdem.
func (ItemOrdem) Fields() []ent.Field {
	return []ent.Field{
		field.Int("quantidade"),
		field.Float("preco_unitario").SchemaType(map[string]string{
			dialect.Postgres: "numeric(8,2)",
		}),
		field.Float("preco_total").SchemaType(map[string]string{
			dialect.Postgres: "numeric(8,2)",
		}),
		field.Int("ordem_id"),
		field.Int("produto_id"),
	}
}

// Edges of the ItemOrdem
func (ItemOrdem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("ordem", Ordem.Type).
			Required().
			Unique().
			Field("ordem_id"),
		edge.To("produto", Produto.Type).
			Required().
			Unique().
			Field("produto_id"),
		edge.From("envio", Envio.Type).Ref("itens").Unique(),
	}
}
