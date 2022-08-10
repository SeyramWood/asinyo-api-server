package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

func (Product) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Float("price").Default(0.00),
		field.Float("promo_price").Optional().Nillable(),
		field.Uint32("quantity").Default(1),
		field.String("unit").NotEmpty(),
		field.Text("description").NotEmpty(),
		field.String("image").NotEmpty(),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("order_details", OrderDetail.Type),
		edge.To("favourites", Favourite.Type),
		edge.From("merchant", Merchant.Type).
			Ref("products").
			Unique().
			Required(),
		edge.From("major", ProductCategoryMajor.Type).
			Ref("products").
			Unique().
			Required(),
		edge.From("minor", ProductCategoryMinor.Type).
			Ref("products").
			Unique().
			Required(),
	}
}
