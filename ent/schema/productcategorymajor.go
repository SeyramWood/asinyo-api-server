package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProductCategoryMajor holds the schema definition for the ProductCategoryMajor entity.
type ProductCategoryMajor struct {
	ent.Schema
}

func (ProductCategoryMajor) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the ProductCategoryMajor.
func (ProductCategoryMajor) Fields() []ent.Field {
	return []ent.Field{
		field.String("category").NotEmpty().Unique(),
		field.String("slug").NotEmpty(),
	}
}

// Edges of the ProductCategoryMajor.
func (ProductCategoryMajor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("minors", ProductCategoryMinor.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("products", Product.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}
