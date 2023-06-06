package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// PriceModel holds the schema definition for the PriceModel entity.
type PriceModel struct {
	ent.Schema
}

func (PriceModel) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the PriceModel.
func (PriceModel) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("initials").NotEmpty(),
		field.String("formula").NotEmpty(),
		field.String("asinyo_formula").Default("(percentage/100) * cp").Optional(),
	}
}

// Edges of the PriceModel.
func (PriceModel) Edges() []ent.Edge {
	return nil
}
