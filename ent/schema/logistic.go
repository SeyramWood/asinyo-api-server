package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/SeyramWood/app/domain/models"
)

// Logistic holds the schema definition for the Logistic entity.
type Logistic struct {
	ent.Schema
}

func (Logistic) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Logistic.
func (Logistic) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("task", &models.TookanPickupAndDeliveryTaskResponse{}).Optional(),
	}
}

// Edges of the Logistic.
func (Logistic) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("order", Order.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.From("store", MerchantStore.Type).
			Ref("logistics").
			Unique(),
	}
}
