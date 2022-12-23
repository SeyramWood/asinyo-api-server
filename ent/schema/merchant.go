package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Merchant holds the schema definition for the Merchant entity.
type Merchant struct {
	ent.Schema
}

func (Merchant) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Merchant.
func (Merchant) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").NotEmpty().Unique(),
		field.Bytes("password").NotEmpty().Sensitive(),
		field.String("type").NotEmpty(),
		field.Bool("otp").Optional(),
	}
}

// Edges of the Merchant.
func (Merchant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("supplier", SupplierMerchant.Type).
			Unique().Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("retailer", RetailMerchant.Type).
			Unique().Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("store", MerchantStore.Type).
			Unique().Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("products", Product.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("addresses", Address.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("orders", Order.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("favourites", Favourite.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}
