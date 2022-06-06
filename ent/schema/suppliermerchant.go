package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SupplierMerchant holds the schema definition for the SupplierMerchant entity.
type SupplierMerchant struct {
	ent.Schema
}

func (SupplierMerchant) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the SupplierMerchant.
func (SupplierMerchant) Fields() []ent.Field {
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

// Edges of the SupplierMerchant.
func (SupplierMerchant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("merchant", Merchant.Type).
			Ref("supplier").
			Unique().
			Required(),
	}
}
