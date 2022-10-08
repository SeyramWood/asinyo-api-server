// Code generated by ent, DO NOT EDIT.

package pickupstation

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/SeyramWood/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// Region applies equality check predicate on the "region" field. It's identical to RegionEQ.
func Region(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRegion), v))
	})
}

// City applies equality check predicate on the "city" field. It's identical to CityEQ.
func City(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCity), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddress), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.PickupStation {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.PickupStation {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.PickupStation {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.PickupStation {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// RegionEQ applies the EQ predicate on the "region" field.
func RegionEQ(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRegion), v))
	})
}

// RegionNEQ applies the NEQ predicate on the "region" field.
func RegionNEQ(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRegion), v))
	})
}

// RegionIn applies the In predicate on the "region" field.
func RegionIn(vs ...string) predicate.PickupStation {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRegion), v...))
	})
}

// RegionNotIn applies the NotIn predicate on the "region" field.
func RegionNotIn(vs ...string) predicate.PickupStation {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRegion), v...))
	})
}

// RegionGT applies the GT predicate on the "region" field.
func RegionGT(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRegion), v))
	})
}

// RegionGTE applies the GTE predicate on the "region" field.
func RegionGTE(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRegion), v))
	})
}

// RegionLT applies the LT predicate on the "region" field.
func RegionLT(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRegion), v))
	})
}

// RegionLTE applies the LTE predicate on the "region" field.
func RegionLTE(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRegion), v))
	})
}

// RegionContains applies the Contains predicate on the "region" field.
func RegionContains(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRegion), v))
	})
}

// RegionHasPrefix applies the HasPrefix predicate on the "region" field.
func RegionHasPrefix(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRegion), v))
	})
}

// RegionHasSuffix applies the HasSuffix predicate on the "region" field.
func RegionHasSuffix(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRegion), v))
	})
}

// RegionEqualFold applies the EqualFold predicate on the "region" field.
func RegionEqualFold(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRegion), v))
	})
}

// RegionContainsFold applies the ContainsFold predicate on the "region" field.
func RegionContainsFold(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRegion), v))
	})
}

// CityEQ applies the EQ predicate on the "city" field.
func CityEQ(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCity), v))
	})
}

// CityNEQ applies the NEQ predicate on the "city" field.
func CityNEQ(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCity), v))
	})
}

// CityIn applies the In predicate on the "city" field.
func CityIn(vs ...string) predicate.PickupStation {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCity), v...))
	})
}

// CityNotIn applies the NotIn predicate on the "city" field.
func CityNotIn(vs ...string) predicate.PickupStation {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCity), v...))
	})
}

// CityGT applies the GT predicate on the "city" field.
func CityGT(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCity), v))
	})
}

// CityGTE applies the GTE predicate on the "city" field.
func CityGTE(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCity), v))
	})
}

// CityLT applies the LT predicate on the "city" field.
func CityLT(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCity), v))
	})
}

// CityLTE applies the LTE predicate on the "city" field.
func CityLTE(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCity), v))
	})
}

// CityContains applies the Contains predicate on the "city" field.
func CityContains(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCity), v))
	})
}

// CityHasPrefix applies the HasPrefix predicate on the "city" field.
func CityHasPrefix(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCity), v))
	})
}

// CityHasSuffix applies the HasSuffix predicate on the "city" field.
func CityHasSuffix(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCity), v))
	})
}

// CityEqualFold applies the EqualFold predicate on the "city" field.
func CityEqualFold(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCity), v))
	})
}

// CityContainsFold applies the ContainsFold predicate on the "city" field.
func CityContainsFold(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCity), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.PickupStation {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.PickupStation {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddress), v))
	})
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAddress), v))
	})
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.PickupStation {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAddress), v...))
	})
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.PickupStation {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAddress), v...))
	})
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAddress), v))
	})
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAddress), v))
	})
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAddress), v))
	})
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAddress), v))
	})
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAddress), v))
	})
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAddress), v))
	})
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAddress), v))
	})
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAddress), v))
	})
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAddress), v))
	})
}

// HasOrders applies the HasEdge predicate on the "orders" edge.
func HasOrders() predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrdersTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OrdersTable, OrdersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrdersWith applies the HasEdge predicate on the "orders" edge with a given conditions (other predicates).
func HasOrdersWith(preds ...predicate.Order) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrdersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OrdersTable, OrdersColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PickupStation) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PickupStation) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
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
func Not(p predicate.PickupStation) predicate.PickupStation {
	return predicate.PickupStation(func(s *sql.Selector) {
		p(s.Not())
	})
}
