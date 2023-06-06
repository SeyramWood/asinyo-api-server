package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/domain/services"
)

// MerchantStore holds the schema definition for the MerchantStore entity.
type MerchantStore struct {
	ent.Schema
}

func (MerchantStore) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the MerchantStore.
func (MerchantStore) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Text("about").NotEmpty(),
		field.String("slogan").NotEmpty(),
		field.Text("description").NotEmpty(),
		field.String("logo").NotEmpty(),
		field.JSON("images", []string{}).Optional(),
		field.Enum("default_account").Values("bank", "momo").Optional(),
		field.JSON("bank_account", &models.MerchantBankAccount{}).Optional(),
		field.JSON("momo_account", &models.MerchantMomoAccount{}).Optional(),
		field.JSON("address", &models.MerchantStoreAddress{}).Optional(),
		field.JSON("coordinate", &services.Coordinate{}).Optional(),
		field.String("merchant_type").NotEmpty(),
		field.Bool("permit_agent").Default(true),
	}
}

// Edges of the MerchantStore.
func (MerchantStore) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("merchant", Merchant.Type).
			Ref("store").
			Unique(),
		edge.From("agent", Agent.Type).
			Ref("store").
			Unique(),
		edge.To("requests", AgentRequest.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("orders", Order.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("order_details", OrderDetail.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}
