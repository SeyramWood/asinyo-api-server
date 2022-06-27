package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OrderDetail holds the schema definition for the OrderDetail entity.
type OrderDetail struct {
	ent.Schema
}

func (OrderDetail) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the OrderDetail.
func (OrderDetail) Fields() []ent.Field {
	return []ent.Field{
		field.Float("price").Default(0.00),
		field.Float("promo_price").Default(0.00),
		field.Float("amount").Default(0.00),
		field.Int("quantity").Default(0),
	}
}

// Edges of the OrderDetail.
func (OrderDetail) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Order", Order.Type).
			Ref("details").
			Unique().
			Required(),
		edge.From("product", Product.Type).
			Ref("orders").
			Unique().
			Required(),
		edge.From("store", MerchantStore.Type).
			Ref("orders").
			Unique().
			Required(),
	}
}
