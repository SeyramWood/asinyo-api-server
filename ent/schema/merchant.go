package schema

import (
	"entgo.io/ent"
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
		field.Enum("otp").Values("active", "inactive").Optional(),
	}
}

// Edges of the Merchant.
func (Merchant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("supplier", SupplierMerchant.Type).
			Unique(),
		edge.To("retailer", RetailMerchant.Type).
			Unique(),
		edge.To("store", MerchantStore.Type).
			Unique(),
		edge.To("products", Product.Type),
		edge.To("addresses", Address.Type),
		edge.To("orders", Order.Type),
		edge.To("favourites", Favourite.Type),
	}
}
