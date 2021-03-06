// Code generated by entc, DO NOT EDIT.

package suppliermerchant

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/SeyramWood/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// GhanaCard applies equality check predicate on the "ghana_card" field. It's identical to GhanaCardEQ.
func GhanaCard(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGhanaCard), v))
	})
}

// LastName applies equality check predicate on the "last_name" field. It's identical to LastNameEQ.
func LastName(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastName), v))
	})
}

// OtherName applies equality check predicate on the "other_name" field. It's identical to OtherNameEQ.
func OtherName(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOtherName), v))
	})
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhone), v))
	})
}

// OtherPhone applies equality check predicate on the "other_phone" field. It's identical to OtherPhoneEQ.
func OtherPhone(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOtherPhone), v))
	})
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddress), v))
	})
}

// DigitalAddress applies equality check predicate on the "digital_address" field. It's identical to DigitalAddressEQ.
func DigitalAddress(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDigitalAddress), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.SupplierMerchant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.SupplierMerchant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.SupplierMerchant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.SupplierMerchant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// GhanaCardEQ applies the EQ predicate on the "ghana_card" field.
func GhanaCardEQ(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGhanaCard), v))
	})
}

// GhanaCardNEQ applies the NEQ predicate on the "ghana_card" field.
func GhanaCardNEQ(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGhanaCard), v))
	})
}

// GhanaCardIn applies the In predicate on the "ghana_card" field.
func GhanaCardIn(vs ...string) predicate.SupplierMerchant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGhanaCard), v...))
	})
}

// GhanaCardNotIn applies the NotIn predicate on the "ghana_card" field.
func GhanaCardNotIn(vs ...string) predicate.SupplierMerchant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGhanaCard), v...))
	})
}

// GhanaCardGT applies the GT predicate on the "ghana_card" field.
func GhanaCardGT(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGhanaCard), v))
	})
}

// GhanaCardGTE applies the GTE predicate on the "ghana_card" field.
func GhanaCardGTE(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGhanaCard), v))
	})
}

// GhanaCardLT applies the LT predicate on the "ghana_card" field.
func GhanaCardLT(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGhanaCard), v))
	})
}

// GhanaCardLTE applies the LTE predicate on the "ghana_card" field.
func GhanaCardLTE(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGhanaCard), v))
	})
}

// GhanaCardContains applies the Contains predicate on the "ghana_card" field.
func GhanaCardContains(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldGhanaCard), v))
	})
}

// GhanaCardHasPrefix applies the HasPrefix predicate on the "ghana_card" field.
func GhanaCardHasPrefix(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldGhanaCard), v))
	})
}

// GhanaCardHasSuffix applies the HasSuffix predicate on the "ghana_card" field.
func GhanaCardHasSuffix(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldGhanaCard), v))
	})
}

// GhanaCardEqualFold applies the EqualFold predicate on the "ghana_card" field.
func GhanaCardEqualFold(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldGhanaCard), v))
	})
}

// GhanaCardContainsFold applies the ContainsFold predicate on the "ghana_card" field.
func GhanaCardContainsFold(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldGhanaCard), v))
	})
}

// LastNameEQ applies the EQ predicate on the "last_name" field.
func LastNameEQ(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastName), v))
	})
}

// LastNameNEQ applies the NEQ predicate on the "last_name" field.
func LastNameNEQ(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLastName), v))
	})
}

// LastNameIn applies the In predicate on the "last_name" field.
func LastNameIn(vs ...string) predicate.SupplierMerchant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLastName), v...))
	})
}

// LastNameNotIn applies the NotIn predicate on the "last_name" field.
func LastNameNotIn(vs ...string) predicate.SupplierMerchant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLastName), v...))
	})
}

// LastNameGT applies the GT predicate on the "last_name" field.
func LastNameGT(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLastName), v))
	})
}

// LastNameGTE applies the GTE predicate on the "last_name" field.
func LastNameGTE(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLastName), v))
	})
}

// LastNameLT applies the LT predicate on the "last_name" field.
func LastNameLT(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLastName), v))
	})
}

// LastNameLTE applies the LTE predicate on the "last_name" field.
func LastNameLTE(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLastName), v))
	})
}

// LastNameContains applies the Contains predicate on the "last_name" field.
func LastNameContains(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLastName), v))
	})
}

// LastNameHasPrefix applies the HasPrefix predicate on the "last_name" field.
func LastNameHasPrefix(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLastName), v))
	})
}

// LastNameHasSuffix applies the HasSuffix predicate on the "last_name" field.
func LastNameHasSuffix(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLastName), v))
	})
}

// LastNameEqualFold applies the EqualFold predicate on the "last_name" field.
func LastNameEqualFold(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLastName), v))
	})
}

// LastNameContainsFold applies the ContainsFold predicate on the "last_name" field.
func LastNameContainsFold(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLastName), v))
	})
}

// OtherNameEQ applies the EQ predicate on the "other_name" field.
func OtherNameEQ(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOtherName), v))
	})
}

// OtherNameNEQ applies the NEQ predicate on the "other_name" field.
func OtherNameNEQ(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOtherName), v))
	})
}

// OtherNameIn applies the In predicate on the "other_name" field.
func OtherNameIn(vs ...string) predicate.SupplierMerchant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOtherName), v...))
	})
}

// OtherNameNotIn applies the NotIn predicate on the "other_name" field.
func OtherNameNotIn(vs ...string) predicate.SupplierMerchant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOtherName), v...))
	})
}

// OtherNameGT applies the GT predicate on the "other_name" field.
func OtherNameGT(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOtherName), v))
	})
}

// OtherNameGTE applies the GTE predicate on the "other_name" field.
func OtherNameGTE(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOtherName), v))
	})
}

// OtherNameLT applies the LT predicate on the "other_name" field.
func OtherNameLT(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOtherName), v))
	})
}

// OtherNameLTE applies the LTE predicate on the "other_name" field.
func OtherNameLTE(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOtherName), v))
	})
}

// OtherNameContains applies the Contains predicate on the "other_name" field.
func OtherNameContains(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOtherName), v))
	})
}

// OtherNameHasPrefix applies the HasPrefix predicate on the "other_name" field.
func OtherNameHasPrefix(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOtherName), v))
	})
}

// OtherNameHasSuffix applies the HasSuffix predicate on the "other_name" field.
func OtherNameHasSuffix(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOtherName), v))
	})
}

// OtherNameEqualFold applies the EqualFold predicate on the "other_name" field.
func OtherNameEqualFold(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOtherName), v))
	})
}

// OtherNameContainsFold applies the ContainsFold predicate on the "other_name" field.
func OtherNameContainsFold(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOtherName), v))
	})
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhone), v))
	})
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPhone), v))
	})
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.SupplierMerchant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPhone), v...))
	})
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.SupplierMerchant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPhone), v...))
	})
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPhone), v))
	})
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPhone), v))
	})
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPhone), v))
	})
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPhone), v))
	})
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPhone), v))
	})
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPhone), v))
	})
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPhone), v))
	})
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPhone), v))
	})
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPhone), v))
	})
}

// OtherPhoneEQ applies the EQ predicate on the "other_phone" field.
func OtherPhoneEQ(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOtherPhone), v))
	})
}

// OtherPhoneNEQ applies the NEQ predicate on the "other_phone" field.
func OtherPhoneNEQ(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOtherPhone), v))
	})
}

// OtherPhoneIn applies the In predicate on the "other_phone" field.
func OtherPhoneIn(vs ...string) predicate.SupplierMerchant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOtherPhone), v...))
	})
}

// OtherPhoneNotIn applies the NotIn predicate on the "other_phone" field.
func OtherPhoneNotIn(vs ...string) predicate.SupplierMerchant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOtherPhone), v...))
	})
}

// OtherPhoneGT applies the GT predicate on the "other_phone" field.
func OtherPhoneGT(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOtherPhone), v))
	})
}

// OtherPhoneGTE applies the GTE predicate on the "other_phone" field.
func OtherPhoneGTE(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOtherPhone), v))
	})
}

// OtherPhoneLT applies the LT predicate on the "other_phone" field.
func OtherPhoneLT(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOtherPhone), v))
	})
}

// OtherPhoneLTE applies the LTE predicate on the "other_phone" field.
func OtherPhoneLTE(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOtherPhone), v))
	})
}

// OtherPhoneContains applies the Contains predicate on the "other_phone" field.
func OtherPhoneContains(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOtherPhone), v))
	})
}

// OtherPhoneHasPrefix applies the HasPrefix predicate on the "other_phone" field.
func OtherPhoneHasPrefix(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOtherPhone), v))
	})
}

// OtherPhoneHasSuffix applies the HasSuffix predicate on the "other_phone" field.
func OtherPhoneHasSuffix(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOtherPhone), v))
	})
}

// OtherPhoneIsNil applies the IsNil predicate on the "other_phone" field.
func OtherPhoneIsNil() predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOtherPhone)))
	})
}

// OtherPhoneNotNil applies the NotNil predicate on the "other_phone" field.
func OtherPhoneNotNil() predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOtherPhone)))
	})
}

// OtherPhoneEqualFold applies the EqualFold predicate on the "other_phone" field.
func OtherPhoneEqualFold(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOtherPhone), v))
	})
}

// OtherPhoneContainsFold applies the ContainsFold predicate on the "other_phone" field.
func OtherPhoneContainsFold(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOtherPhone), v))
	})
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddress), v))
	})
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAddress), v))
	})
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.SupplierMerchant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAddress), v...))
	})
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.SupplierMerchant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAddress), v...))
	})
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAddress), v))
	})
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAddress), v))
	})
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAddress), v))
	})
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAddress), v))
	})
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAddress), v))
	})
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAddress), v))
	})
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAddress), v))
	})
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAddress), v))
	})
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAddress), v))
	})
}

// DigitalAddressEQ applies the EQ predicate on the "digital_address" field.
func DigitalAddressEQ(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDigitalAddress), v))
	})
}

// DigitalAddressNEQ applies the NEQ predicate on the "digital_address" field.
func DigitalAddressNEQ(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDigitalAddress), v))
	})
}

// DigitalAddressIn applies the In predicate on the "digital_address" field.
func DigitalAddressIn(vs ...string) predicate.SupplierMerchant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDigitalAddress), v...))
	})
}

// DigitalAddressNotIn applies the NotIn predicate on the "digital_address" field.
func DigitalAddressNotIn(vs ...string) predicate.SupplierMerchant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDigitalAddress), v...))
	})
}

// DigitalAddressGT applies the GT predicate on the "digital_address" field.
func DigitalAddressGT(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDigitalAddress), v))
	})
}

// DigitalAddressGTE applies the GTE predicate on the "digital_address" field.
func DigitalAddressGTE(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDigitalAddress), v))
	})
}

// DigitalAddressLT applies the LT predicate on the "digital_address" field.
func DigitalAddressLT(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDigitalAddress), v))
	})
}

// DigitalAddressLTE applies the LTE predicate on the "digital_address" field.
func DigitalAddressLTE(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDigitalAddress), v))
	})
}

// DigitalAddressContains applies the Contains predicate on the "digital_address" field.
func DigitalAddressContains(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDigitalAddress), v))
	})
}

// DigitalAddressHasPrefix applies the HasPrefix predicate on the "digital_address" field.
func DigitalAddressHasPrefix(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDigitalAddress), v))
	})
}

// DigitalAddressHasSuffix applies the HasSuffix predicate on the "digital_address" field.
func DigitalAddressHasSuffix(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDigitalAddress), v))
	})
}

// DigitalAddressEqualFold applies the EqualFold predicate on the "digital_address" field.
func DigitalAddressEqualFold(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDigitalAddress), v))
	})
}

// DigitalAddressContainsFold applies the ContainsFold predicate on the "digital_address" field.
func DigitalAddressContainsFold(v string) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDigitalAddress), v))
	})
}

// HasMerchant applies the HasEdge predicate on the "merchant" edge.
func HasMerchant() predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MerchantTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, MerchantTable, MerchantColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMerchantWith applies the HasEdge predicate on the "merchant" edge with a given conditions (other predicates).
func HasMerchantWith(preds ...predicate.Merchant) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MerchantInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, MerchantTable, MerchantColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SupplierMerchant) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SupplierMerchant) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
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
func Not(p predicate.SupplierMerchant) predicate.SupplierMerchant {
	return predicate.SupplierMerchant(func(s *sql.Selector) {
		p(s.Not())
	})
}
