package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
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
		field.Int("id").StructTag(`json:"oid,omitempty"`),
		field.Time("data_ordem").Default(time.Now),
		field.Enum("status").Values("completa", "parcial", "intocada").Default("intocada"),
		field.Float("preco_da_ordem").SchemaType(map[string]string{
			dialect.Postgres: "numeric(8,2)",
		}),
	}
}

// Edges of the Ordem.
func (Ordem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("produtos", Produto.Type).Through("items", ItemOrdem.Type),
		edge.From("cliente", Cliente.Type).Ref("ordens").Unique(),
	}
}
