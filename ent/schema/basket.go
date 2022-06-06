package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Basket holds the schema definition for the Basket entity.
type Basket struct {
	ent.Schema
}

func (Basket) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Basket.
func (Basket) Fields() []ent.Field {
	return nil
}

// Edges of the Basket.
func (Basket) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("merchant", Merchant.Type).
			Ref("baskets").
			Unique(),
		edge.From("agent", Agent.Type).
			Ref("baskets").
			Unique(),
		edge.From("customer", Customer.Type).
			Ref("baskets").
			Unique(),
		edge.From("product", Product.Type).
			Ref("baskets").
			Unique(),
	}
}
