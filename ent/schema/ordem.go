package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Ordem holds the schema definition for the Ordem entity.
type Ordem struct {
	ent.Schema
}

// Fields of the Ordem.
func (Ordem) Fields() []ent.Field {
	return []ent.Field{
		field.Time("data_ordem"),
	}
}

// Edges of the Ordem.
func (Ordem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("produtos", Produto.Type).Through("items", ItemOrdem.Type),
		edge.From("clientes", Cliente.Type).Ref("ordens").Unique(),
	}
}
