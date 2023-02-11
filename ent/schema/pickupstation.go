package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PickupStation holds the schema definition for the PickupStation entity.
type PickupStation struct {
	ent.Schema
}

func (PickupStation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the PickupStation.
func (PickupStation) Fields() []ent.Field {
	return []ent.Field{
		field.String("region").NotEmpty(),
		field.String("city").NotEmpty(),
		field.String("name").NotEmpty(),
		field.String("address").NotEmpty(),
	}
}

// Edges of the PickupStation.
func (PickupStation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("orders", Order.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}
