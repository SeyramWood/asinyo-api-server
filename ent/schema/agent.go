package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/SeyramWood/app/domain/models"
)

// Agent holds the schema definition for the Agent entity.
type Agent struct {
	ent.Schema
}

func (Agent) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Agent.
func (Agent) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").NotEmpty().Unique(),
		field.Bytes("password").NotEmpty().Sensitive(),
		field.String("ghana_card").NotEmpty().Unique(),
		field.String("last_name").NotEmpty(),
		field.String("other_name").NotEmpty(),
		field.String("phone").NotEmpty().Unique(),
		field.String("other_phone").Optional(),
		field.String("address").NotEmpty(),
		field.String("digital_address").NotEmpty(),
		field.String("region").Optional(),
		field.String("district").Optional(),
		field.String("city").Optional(),
		field.Enum("default_account").Values("bank", "momo").Optional(),
		field.JSON("bank_account", &models.AgentBankAccount{}).Optional(),
		field.JSON("momo_account", &models.AgentMomoAccount{}).Optional(),
		field.Bool("verified").Default(false),
		field.JSON("compliance", &models.AgentComplianceModel{}).Optional(),
	}
}

// Edges of the Agent.
func (Agent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("addresses", Address.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("orders", Order.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("favourites", Favourite.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("store", MerchantStore.Type),
		edge.To("requests", AgentRequest.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("notifications", Notification.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}
