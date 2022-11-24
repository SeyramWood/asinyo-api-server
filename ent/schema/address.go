package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/SeyramWood/app/domain/services"
)

// Address holds the schema definition for the Address entity.
type Address struct {
	ent.Schema
}

func (Address) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Address.
func (Address) Fields() []ent.Field {
	return []ent.Field{
		field.String("last_name").NotEmpty(),
		field.String("other_name").NotEmpty(),
		field.String("phone").NotEmpty(),
		field.String("other_phone").Optional(),
		field.String("digital_address").Optional(),
		field.String("street_name").Optional(),
		field.String("street_number").Optional(),
		field.String("city").NotEmpty(),
		field.String("district").Optional(),
		field.String("Region").NotEmpty(),
		field.String("Country").Default("Ghana"),
		field.Text("address").NotEmpty(),
		field.Bool("default").Default(false),
		field.JSON("coordinate", &services.Coordinate{}).Optional(),
	}
}

// Edges of the Address.
func (Address) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("merchant", Merchant.Type).
			Ref("addresses").
			Unique(),
		edge.From("agent", Agent.Type).
			Ref("addresses").
			Unique(),
		edge.From("customer", Customer.Type).
			Ref("addresses").
			Unique(),
		edge.To("orders", Order.Type),
	}
}
