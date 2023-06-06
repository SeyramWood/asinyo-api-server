// Code generated by ent, DO NOT EDIT.

package permission

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/SeyramWood/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldUpdatedAt, v))
}

// Permission applies equality check predicate on the "permission" field. It's identical to PermissionEQ.
func Permission(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldPermission, v))
}

// Slug applies equality check predicate on the "slug" field. It's identical to SlugEQ.
func Slug(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldSlug, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldUpdatedAt, v))
}

// PermissionEQ applies the EQ predicate on the "permission" field.
func PermissionEQ(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldPermission, v))
}

// PermissionNEQ applies the NEQ predicate on the "permission" field.
func PermissionNEQ(v string) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldPermission, v))
}

// PermissionIn applies the In predicate on the "permission" field.
func PermissionIn(vs ...string) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldPermission, vs...))
}

// PermissionNotIn applies the NotIn predicate on the "permission" field.
func PermissionNotIn(vs ...string) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldPermission, vs...))
}

// PermissionGT applies the GT predicate on the "permission" field.
func PermissionGT(v string) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldPermission, v))
}

// PermissionGTE applies the GTE predicate on the "permission" field.
func PermissionGTE(v string) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldPermission, v))
}

// PermissionLT applies the LT predicate on the "permission" field.
func PermissionLT(v string) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldPermission, v))
}

// PermissionLTE applies the LTE predicate on the "permission" field.
func PermissionLTE(v string) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldPermission, v))
}

// PermissionContains applies the Contains predicate on the "permission" field.
func PermissionContains(v string) predicate.Permission {
	return predicate.Permission(sql.FieldContains(FieldPermission, v))
}

// PermissionHasPrefix applies the HasPrefix predicate on the "permission" field.
func PermissionHasPrefix(v string) predicate.Permission {
	return predicate.Permission(sql.FieldHasPrefix(FieldPermission, v))
}

// PermissionHasSuffix applies the HasSuffix predicate on the "permission" field.
func PermissionHasSuffix(v string) predicate.Permission {
	return predicate.Permission(sql.FieldHasSuffix(FieldPermission, v))
}

// PermissionEqualFold applies the EqualFold predicate on the "permission" field.
func PermissionEqualFold(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEqualFold(FieldPermission, v))
}

// PermissionContainsFold applies the ContainsFold predicate on the "permission" field.
func PermissionContainsFold(v string) predicate.Permission {
	return predicate.Permission(sql.FieldContainsFold(FieldPermission, v))
}

// SlugEQ applies the EQ predicate on the "slug" field.
func SlugEQ(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldSlug, v))
}

// SlugNEQ applies the NEQ predicate on the "slug" field.
func SlugNEQ(v string) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldSlug, v))
}

// SlugIn applies the In predicate on the "slug" field.
func SlugIn(vs ...string) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldSlug, vs...))
}

// SlugNotIn applies the NotIn predicate on the "slug" field.
func SlugNotIn(vs ...string) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldSlug, vs...))
}

// SlugGT applies the GT predicate on the "slug" field.
func SlugGT(v string) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldSlug, v))
}

// SlugGTE applies the GTE predicate on the "slug" field.
func SlugGTE(v string) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldSlug, v))
}

// SlugLT applies the LT predicate on the "slug" field.
func SlugLT(v string) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldSlug, v))
}

// SlugLTE applies the LTE predicate on the "slug" field.
func SlugLTE(v string) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldSlug, v))
}

// SlugContains applies the Contains predicate on the "slug" field.
func SlugContains(v string) predicate.Permission {
	return predicate.Permission(sql.FieldContains(FieldSlug, v))
}

// SlugHasPrefix applies the HasPrefix predicate on the "slug" field.
func SlugHasPrefix(v string) predicate.Permission {
	return predicate.Permission(sql.FieldHasPrefix(FieldSlug, v))
}

// SlugHasSuffix applies the HasSuffix predicate on the "slug" field.
func SlugHasSuffix(v string) predicate.Permission {
	return predicate.Permission(sql.FieldHasSuffix(FieldSlug, v))
}

// SlugEqualFold applies the EqualFold predicate on the "slug" field.
func SlugEqualFold(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEqualFold(FieldSlug, v))
}

// SlugContainsFold applies the ContainsFold predicate on the "slug" field.
func SlugContainsFold(v string) predicate.Permission {
	return predicate.Permission(sql.FieldContainsFold(FieldSlug, v))
}

// HasRole applies the HasEdge predicate on the "role" edge.
func HasRole() predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, RoleTable, RolePrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoleWith applies the HasEdge predicate on the "role" edge with a given conditions (other predicates).
func HasRoleWith(preds ...predicate.Role) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		step := newRoleStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Permission) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Permission) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Permission) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		p(s.Not())
	})
}
