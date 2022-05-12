package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// RetailMerchant holds the schema definition for the RetailMerchant entity.
type RetailMerchant struct {
	ent.Schema
}

func (RetailMerchant) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the RetailMerchant.
func (RetailMerchant) Fields() []ent.Field {
	return []ent.Field{
		field.String("ghana_card").NotEmpty().Unique(),
		field.String("last_name").NotEmpty(),
		field.String("other_name").NotEmpty(),
		field.String("phone").NotEmpty().Unique(),
		field.String("other_phone").Optional().Nillable(),
		field.String("address").NotEmpty(),
		field.String("digital_address").NotEmpty(),
	}
}

// Edges of the RetailMerchant.
func (RetailMerchant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("products", Product.Type).
			Unique(),
		edge.From("merchant", Merchant.Type).
			Ref("retailer").
			Unique().
			Required(),
	}
}
