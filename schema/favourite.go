package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Favourite holds the schema definition for the Favourite entity.
type Favourite struct {
	ent.Schema
}

func (Favourite) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Favourite.
func (Favourite) Fields() []ent.Field {
	return nil
}

// Edges of the Favourite.
func (Favourite) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("merchant", Merchant.Type).
			Ref("favourites").
			Unique(),
		edge.From("agent", Agent.Type).
			Ref("favourites").
			Unique(),
		edge.From("customer", Customer.Type).
			Ref("favourites").
			Unique(),
		edge.From("product", Product.Type).
			Ref("favourites").
			Unique(),
	}
}
