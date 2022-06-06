package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

func (Order) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("status").
			Values("in_progress", "shipping", "delivered").
			Default("in_progress"),
		field.Time("delivered_at").Nillable().Optional(),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("merchant", Merchant.Type).
			Ref("orders").
			Unique(),
		edge.From("agent", Agent.Type).
			Ref("orders").
			Unique(),
		edge.From("customer", Customer.Type).
			Ref("orders").
			Unique(),
		edge.From("address", Address.Type).
			Ref("orders").
			Unique(),
		edge.From("product", Product.Type).
			Ref("orders").
			Unique(),
	}
}
