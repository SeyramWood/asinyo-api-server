package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// type TimeMixin struct {
// 	// We embed the `mixin.Schema` to avoid
// 	// implementing the rest of the methods.
// 	mixin.Schema
// }

// func (TimeMixin) Fields() []ent.Field {
// 	return []ent.Field{
// 		field.Time("created_at").
// 			Immutable().
// 			Default(time.Now),
// 		field.Time("updated_at").
// 			Default(time.Now).
// 			UpdateDefault(time.Now),
// 	}
// }

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
		field.String("email").NotEmpty().Unique(),
		field.Bytes("password").NotEmpty().Sensitive(),
		field.String("first_name").NotEmpty(),
		field.String("last_name").NotEmpty(),
		field.String("phone").NotEmpty(),
		field.String("other_phone").Optional().Nillable(),
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return nil
}
