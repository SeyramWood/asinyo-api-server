package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Logistic holds the schema definition for the Logistic entity.
type Logistic struct {
	ent.Schema
}

func (Logistic) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Logistic.
func (Logistic) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").Default("Asinyo").NotEmpty(),
		field.JSON("task", &struct {
			Data any `json:"data"`
		}{}).Optional(),
	}
}

// Edges of the Logistic.
func (Logistic) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order", Order.Type).
			Ref("logistic").
			Unique(),
	}
}
