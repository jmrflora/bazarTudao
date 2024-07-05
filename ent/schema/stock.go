package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Stock holds the schema definition for the Stock entity.
type Stock struct {
	ent.Schema
}

// Fields of the Stock.
func (Stock) Fields() []ent.Field {
	return []ent.Field{
		field.Time("data_movimento").Default(time.Now),
		field.Int("quantidade"),
	}
}

// Edges of the Stock.
func (Stock) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("produtos", Produto.Type),
	}
}
