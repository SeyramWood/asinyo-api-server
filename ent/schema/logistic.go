package schema

import (
	"entgo.io/ent"
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
		field.String("tracking_link").Optional(),
		field.JSON("tasks", &models.TookanMultiTaskResponse{}).Optional(),
	}
}

// Edges of the Logistic.
func (Logistic) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("order", Order.Type),
	}
}
