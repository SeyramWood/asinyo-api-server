package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Agent holds the schema definition for the Agent entity.
type Agent struct {
	ent.Schema
}

func (Agent) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Agent.
func (Agent) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").NotEmpty().Unique(),
		field.Bytes("password").NotEmpty().Sensitive(),
		field.String("ghana_card").NotEmpty().Unique(),
		field.String("last_name").NotEmpty(),
		field.String("other_name").NotEmpty(),
		field.String("phone").NotEmpty().Unique(),
		field.String("other_phone").Optional().Nillable(),
		field.String("address").NotEmpty(),
		field.String("digital_address").NotEmpty(),
	}
}

// Edges of the Agent.
func (Agent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("addresses", Address.Type),
		edge.To("orders", Order.Type),
		edge.To("favourites", Favourite.Type),
	}
}
