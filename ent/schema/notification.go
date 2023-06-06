package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/SeyramWood/app/domain/models"
)

// Notification holds the schema definition for the Notification entity.
type Notification struct {
	ent.Schema
}

func (Notification) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Notification.
func (Notification) Fields() []ent.Field {
	return []ent.Field{
		field.String("event").NotEmpty(),
		field.String("activity").NotEmpty(),
		field.String("description").NotEmpty(),
		field.String("subject_type").NotEmpty(),
		field.Int("subject_id").Optional(),
		field.String("creator_type").NotEmpty(),
		field.Int("creator_id").Optional(),
		field.String("customer_read_at").Optional(),
		field.String("agent_read_at").Optional(),
		field.String("merchant_read_at").Optional(),
		field.JSON(
			"admin_read_at", []*models.AdminRead{},
		).Optional(),
		field.JSON(
			"data", &struct {
				Data any `json:"data"`
			}{},
		).Optional(),
	}
}

// Edges of the Notification.
func (Notification) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("admin", Admin.Type).
			Ref("notifications"),
		edge.From("merchant", Merchant.Type).
			Ref("notifications"),
		edge.From("agent", Agent.Type).
			Ref("notifications"),
		edge.From("customer", Customer.Type).
			Ref("notifications"),
	}
}
