package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// IndividualCustomer holds the schema definition for the IndividualCustomer entity.
type IndividualCustomer struct {
	ent.Schema
}

func (IndividualCustomer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the IndividualCustomer.
func (IndividualCustomer) Fields() []ent.Field {
	return []ent.Field{
		field.String("last_name").NotEmpty(),
		field.String("other_name").NotEmpty(),
		field.String("phone").NotEmpty(),
		field.String("other_phone").Optional(),
	}
}

// Edges of the IndividualCustomer.
func (IndividualCustomer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer", Customer.Type).
			Ref("individual").
			Unique().
			Required(),
	}
}
