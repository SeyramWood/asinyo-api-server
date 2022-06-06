package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

func (Customer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").NotEmpty().Unique(),
		field.Bytes("password").NotEmpty().Sensitive(),
		field.String("first_name").NotEmpty(),
		field.String("last_name").NotEmpty(),
		field.String("phone").NotEmpty(),
		field.String("other_phone").Optional().Nillable(),
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("addresses", Address.Type),
		edge.To("orders", Order.Type),
		edge.To("baskets", Basket.Type),
		edge.To("favourites", Favourite.Type),
	}
}
