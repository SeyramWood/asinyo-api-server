package schema

import (
	"entgo.io/ent"
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
		field.String("other_phone").Optional().Nillable(),
		field.String("address").NotEmpty(),
		field.String("digital_address").NotEmpty(),
<<<<<<< HEAD
		field.String("region").NotEmpty().Nillable(),
		field.String("district").NotEmpty().Nillable(),
		field.String("city").NotEmpty().Nillable(),
=======
		field.String("region").Optional().Nillable(),
		field.String("district").Optional().Nillable(),
		field.String("city").Optional().Nillable(),
>>>>>>> dev
		field.Enum("default_account").Values("bank", "momo").Optional(),
		field.JSON("bank_account", &models.MerchantBankAccount{}).Optional(),
		field.JSON("momo_account", &models.MerchantMomoAccount{}).Optional(),
		field.Bool("verified").Default(false),
		field.JSON("compliance", &models.AgentComplianceModel{}).Optional(),
	}
}

// Edges of the Agent.
func (Agent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("addresses", Address.Type),
		edge.To("orders", Order.Type),
		edge.To("favourites", Favourite.Type),
		edge.To("store", MerchantStore.Type),
		edge.To("requests", AgentRequest.Type),
	}
}
