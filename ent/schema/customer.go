package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

func (Customer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").NotEmpty().Unique(),
		field.Bytes("password").NotEmpty().Sensitive(),
		field.String("type").NotEmpty(),
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("business", BusinessCustomer.Type).
			Unique().Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("individual", IndividualCustomer.Type).
			Unique().Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("addresses", Address.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("orders", Order.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("favourites", Favourite.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("notifications", Notification.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("purchase_request", PurchaseRequest.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.From("admin", Admin.Type).Ref("customers").Unique(),
	}
}
