// Code generated by ent, DO NOT EDIT.

package logistic

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/SeyramWood/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// TrackingLink applies equality check predicate on the "tracking_link" field. It's identical to TrackingLinkEQ.
func TrackingLink(v string) predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTrackingLink), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Logistic {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Logistic {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Logistic {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Logistic {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// TrackingLinkEQ applies the EQ predicate on the "tracking_link" field.
func TrackingLinkEQ(v string) predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTrackingLink), v))
	})
}

// TrackingLinkNEQ applies the NEQ predicate on the "tracking_link" field.
func TrackingLinkNEQ(v string) predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTrackingLink), v))
	})
}

// TrackingLinkIn applies the In predicate on the "tracking_link" field.
func TrackingLinkIn(vs ...string) predicate.Logistic {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTrackingLink), v...))
	})
}

// TrackingLinkNotIn applies the NotIn predicate on the "tracking_link" field.
func TrackingLinkNotIn(vs ...string) predicate.Logistic {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTrackingLink), v...))
	})
}

// TrackingLinkGT applies the GT predicate on the "tracking_link" field.
func TrackingLinkGT(v string) predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTrackingLink), v))
	})
}

// TrackingLinkGTE applies the GTE predicate on the "tracking_link" field.
func TrackingLinkGTE(v string) predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTrackingLink), v))
	})
}

// TrackingLinkLT applies the LT predicate on the "tracking_link" field.
func TrackingLinkLT(v string) predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTrackingLink), v))
	})
}

// TrackingLinkLTE applies the LTE predicate on the "tracking_link" field.
func TrackingLinkLTE(v string) predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTrackingLink), v))
	})
}

// TrackingLinkContains applies the Contains predicate on the "tracking_link" field.
func TrackingLinkContains(v string) predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTrackingLink), v))
	})
}

// TrackingLinkHasPrefix applies the HasPrefix predicate on the "tracking_link" field.
func TrackingLinkHasPrefix(v string) predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTrackingLink), v))
	})
}

// TrackingLinkHasSuffix applies the HasSuffix predicate on the "tracking_link" field.
func TrackingLinkHasSuffix(v string) predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTrackingLink), v))
	})
}

// TrackingLinkIsNil applies the IsNil predicate on the "tracking_link" field.
func TrackingLinkIsNil() predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTrackingLink)))
	})
}

// TrackingLinkNotNil applies the NotNil predicate on the "tracking_link" field.
func TrackingLinkNotNil() predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTrackingLink)))
	})
}

// TrackingLinkEqualFold applies the EqualFold predicate on the "tracking_link" field.
func TrackingLinkEqualFold(v string) predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTrackingLink), v))
	})
}

// TrackingLinkContainsFold applies the ContainsFold predicate on the "tracking_link" field.
func TrackingLinkContainsFold(v string) predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTrackingLink), v))
	})
}

// TasksIsNil applies the IsNil predicate on the "tasks" field.
func TasksIsNil() predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTasks)))
	})
}

// TasksNotNil applies the NotNil predicate on the "tasks" field.
func TasksNotNil() predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTasks)))
	})
}

// HasOrder applies the HasEdge predicate on the "order" edge.
func HasOrder() predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrderTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, OrderTable, OrderPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrderWith applies the HasEdge predicate on the "order" edge with a given conditions (other predicates).
func HasOrderWith(preds ...predicate.Order) predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrderInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, OrderTable, OrderPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Logistic) predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Logistic) predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
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
func Not(p predicate.Logistic) predicate.Logistic {
	return predicate.Logistic(func(s *sql.Selector) {
		p(s.Not())
	})
}
