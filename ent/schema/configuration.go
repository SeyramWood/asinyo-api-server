package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Configuration holds the schema definition for the Configuration entity.
type Configuration struct {
	ent.Schema
}

func (Configuration) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Configuration.
func (Configuration) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.JSON("data", &struct {
			Data any `json:"data"`
		}{}),
	}
}

// Edges of the Configuration.
func (Configuration) Edges() []ent.Edge {
	return nil
}
