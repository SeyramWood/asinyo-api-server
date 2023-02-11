package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

func (Order) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.String("order_number").NotEmpty(),
		field.String("currency").NotEmpty(),
		field.Float("amount").Default(0.00),
		field.Float("delivery_fee").Default(0.00),
		field.String("reference").Nillable().Optional(),
		field.String("channel").Nillable().Optional(),
		field.String("paid_at").Nillable().Optional(),
		field.Enum("delivery_method").
			Values("HOD", "PSD"),
		field.Enum("payment_method").
			Values("ONLINE", "POD").
			Default("ONLINE"),
		field.Enum("status").
			Values("pending", "in_progress", "fulfilled", "canceled").
			Default("pending"),
		field.JSON("store_tasks_created", []int{}).Optional(),
		field.Time("delivered_at").Nillable().Optional(),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("details", OrderDetail.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.From("merchant", Merchant.Type).
			Ref("orders").
			Unique(),
		edge.From("agent", Agent.Type).
			Ref("orders").
			Unique(),
		edge.From("customer", Customer.Type).
			Ref("orders").
			Unique(),
		edge.From("address", Address.Type).
			Ref("orders").
			Unique(),
		edge.From("pickup", PickupStation.Type).
			Ref("orders").
			Unique(),
		edge.From("stores", MerchantStore.Type).
			Ref("orders"),
		edge.From("logistic", Logistic.Type).
			Ref("order"),
	}
}
