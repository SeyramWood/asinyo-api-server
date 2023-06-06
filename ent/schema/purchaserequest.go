package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PurchaseRequest holds the schema definition for the PurchaseRequest entity.
type PurchaseRequest struct {
	ent.Schema
}

func (PurchaseRequest) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the PurchaseRequest.
func (PurchaseRequest) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("signed").NotEmpty(),
		field.Text("description").Optional(),
		field.String("file").Optional(),
	}
}

// Edges of the PurchaseRequest.
func (PurchaseRequest) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("order", Order.Type),
		edge.From("customer", Customer.Type).Ref("purchase_request").Unique().Required(),
	}
}
