package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Envio holds the schema definition for the Envio entity.
type Envio struct {
	ent.Schema
}

// Fields of the Envio.
func (Envio) Fields() []ent.Field {
	return []ent.Field{
		field.Time("data").Default(time.Now),
	}
}

// Edges of the Envio.
func (Envio) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("itens", ItemOrdem.Type),
	}
}
