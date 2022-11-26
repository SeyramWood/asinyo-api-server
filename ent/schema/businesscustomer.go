package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/SeyramWood/app/domain/models"
)

// BusinessCustomer holds the schema definition for the BusinessCustomer entity.
type BusinessCustomer struct {
	ent.Schema
}

func (BusinessCustomer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the BusinessCustomer.
func (BusinessCustomer) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("phone").NotEmpty(),
		field.String("other_phone").Optional(),
		field.String("logo").Optional(),
		field.JSON("contact", &models.BusinessCustomerContact{}).Optional(),
	}
}

// Edges of the BusinessCustomer.
func (BusinessCustomer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer", Customer.Type).
			Ref("business").
			Unique().
			Required(),
	}
}
