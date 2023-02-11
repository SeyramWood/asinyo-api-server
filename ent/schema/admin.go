package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Admin holds the schema definition for the Admin entity.
type Admin struct {
	ent.Schema
}

func (Admin) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Admin.
func (Admin) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").NotEmpty().Unique(),
		field.Bytes("password").NotEmpty().Sensitive(),
		field.String("last_name").NotEmpty(),
		field.String("other_name").NotEmpty(),
		field.Enum("status").Values("offline", "online").Default("offline"),
		field.String("last_active").Optional(),
	}
}

// Edges of the Admin.
func (Admin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roles", Role.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}
