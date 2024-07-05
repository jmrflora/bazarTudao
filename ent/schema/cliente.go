package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Cliente holds the schema definition for the Cliente entity.
type Cliente struct {
	ent.Schema
}

// Fields of the Cliente.
func (Cliente) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").Unique(),
		field.String("nome"),
		field.String("cpf").Unique(),
		field.String("telefone"),
		field.String("endereco_entrega"),
	}
}

// Edges of the Cliente.
func (Cliente) Edges() []ent.Edge {
	return []ent.Edge{edge.To("ordens", Ordem.Type)}
}
