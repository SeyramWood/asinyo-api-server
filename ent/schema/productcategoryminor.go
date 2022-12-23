package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProductCategoryMinor holds the schema definition for the ProductCategoryMinor entity.
type ProductCategoryMinor struct {
	ent.Schema
}

func (ProductCategoryMinor) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the ProductCategoryMinor.
func (ProductCategoryMinor) Fields() []ent.Field {
	return []ent.Field{
		field.String("category").NotEmpty().Unique(),
		field.String("image").NotEmpty(),
		field.String("slug").NotEmpty(),
	}
}

// Edges of the ProductCategoryMinor.
func (ProductCategoryMinor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("major", ProductCategoryMajor.Type).
			Ref("minors").
			Unique().
			Required(),
		edge.To("products", Product.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}
