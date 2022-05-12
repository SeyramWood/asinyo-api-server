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
		field.String("Name").NotEmpty(),
		field.Float("Price").Default(0.00),
		field.Float("PromoPrice").Default(0.00),
		field.Text("Description").NotEmpty(),
		field.String("Image").NotEmpty(),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("major", ProductCategoryMajor.Type).
			Ref("products").
			Required(),
		edge.From("minor", ProductCategoryMinor.Type).
			Ref("products").
			Required(),
		edge.From("mechant", Merchant.Type).
			Ref("products").
			Required(),
		edge.From("supplier", SupplierMerchant.Type).
			Ref("products").
			Required(),
		edge.From("retailer", RetailMerchant.Type).
			Ref("products").
			Required(),
	}
}
