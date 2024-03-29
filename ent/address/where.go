// Code generated by ent, DO NOT EDIT.

package address

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/SeyramWood/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldUpdatedAt, v))
}

// LastName applies equality check predicate on the "last_name" field. It's identical to LastNameEQ.
func LastName(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldLastName, v))
}

// OtherName applies equality check predicate on the "other_name" field. It's identical to OtherNameEQ.
func OtherName(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldOtherName, v))
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldPhone, v))
}

// OtherPhone applies equality check predicate on the "other_phone" field. It's identical to OtherPhoneEQ.
func OtherPhone(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldOtherPhone, v))
}

// StreetName applies equality check predicate on the "street_name" field. It's identical to StreetNameEQ.
func StreetName(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldStreetName, v))
}

// StreetNumber applies equality check predicate on the "street_number" field. It's identical to StreetNumberEQ.
func StreetNumber(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldStreetNumber, v))
}

// City applies equality check predicate on the "city" field. It's identical to CityEQ.
func City(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldCity, v))
}

// District applies equality check predicate on the "district" field. It's identical to DistrictEQ.
func District(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldDistrict, v))
}

// Region applies equality check predicate on the "Region" field. It's identical to RegionEQ.
func Region(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldRegion, v))
}

// Country applies equality check predicate on the "Country" field. It's identical to CountryEQ.
func Country(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldCountry, v))
}

// Default applies equality check predicate on the "default" field. It's identical to DefaultEQ.
func Default(v bool) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldDefault, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldUpdatedAt, v))
}

// LastNameEQ applies the EQ predicate on the "last_name" field.
func LastNameEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldLastName, v))
}

// LastNameNEQ applies the NEQ predicate on the "last_name" field.
func LastNameNEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldLastName, v))
}

// LastNameIn applies the In predicate on the "last_name" field.
func LastNameIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldLastName, vs...))
}

// LastNameNotIn applies the NotIn predicate on the "last_name" field.
func LastNameNotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldLastName, vs...))
}

// LastNameGT applies the GT predicate on the "last_name" field.
func LastNameGT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldLastName, v))
}

// LastNameGTE applies the GTE predicate on the "last_name" field.
func LastNameGTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldLastName, v))
}

// LastNameLT applies the LT predicate on the "last_name" field.
func LastNameLT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldLastName, v))
}

// LastNameLTE applies the LTE predicate on the "last_name" field.
func LastNameLTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldLastName, v))
}

// LastNameContains applies the Contains predicate on the "last_name" field.
func LastNameContains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldLastName, v))
}

// LastNameHasPrefix applies the HasPrefix predicate on the "last_name" field.
func LastNameHasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldLastName, v))
}

// LastNameHasSuffix applies the HasSuffix predicate on the "last_name" field.
func LastNameHasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldLastName, v))
}

// LastNameEqualFold applies the EqualFold predicate on the "last_name" field.
func LastNameEqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldLastName, v))
}

// LastNameContainsFold applies the ContainsFold predicate on the "last_name" field.
func LastNameContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldLastName, v))
}

// OtherNameEQ applies the EQ predicate on the "other_name" field.
func OtherNameEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldOtherName, v))
}

// OtherNameNEQ applies the NEQ predicate on the "other_name" field.
func OtherNameNEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldOtherName, v))
}

// OtherNameIn applies the In predicate on the "other_name" field.
func OtherNameIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldOtherName, vs...))
}

// OtherNameNotIn applies the NotIn predicate on the "other_name" field.
func OtherNameNotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldOtherName, vs...))
}

// OtherNameGT applies the GT predicate on the "other_name" field.
func OtherNameGT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldOtherName, v))
}

// OtherNameGTE applies the GTE predicate on the "other_name" field.
func OtherNameGTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldOtherName, v))
}

// OtherNameLT applies the LT predicate on the "other_name" field.
func OtherNameLT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldOtherName, v))
}

// OtherNameLTE applies the LTE predicate on the "other_name" field.
func OtherNameLTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldOtherName, v))
}

// OtherNameContains applies the Contains predicate on the "other_name" field.
func OtherNameContains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldOtherName, v))
}

// OtherNameHasPrefix applies the HasPrefix predicate on the "other_name" field.
func OtherNameHasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldOtherName, v))
}

// OtherNameHasSuffix applies the HasSuffix predicate on the "other_name" field.
func OtherNameHasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldOtherName, v))
}

// OtherNameEqualFold applies the EqualFold predicate on the "other_name" field.
func OtherNameEqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldOtherName, v))
}

// OtherNameContainsFold applies the ContainsFold predicate on the "other_name" field.
func OtherNameContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldOtherName, v))
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldPhone, v))
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldPhone, v))
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldPhone, vs...))
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldPhone, vs...))
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldPhone, v))
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldPhone, v))
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldPhone, v))
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldPhone, v))
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldPhone, v))
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldPhone, v))
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldPhone, v))
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldPhone, v))
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldPhone, v))
}

// OtherPhoneEQ applies the EQ predicate on the "other_phone" field.
func OtherPhoneEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldOtherPhone, v))
}

// OtherPhoneNEQ applies the NEQ predicate on the "other_phone" field.
func OtherPhoneNEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldOtherPhone, v))
}

// OtherPhoneIn applies the In predicate on the "other_phone" field.
func OtherPhoneIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldOtherPhone, vs...))
}

// OtherPhoneNotIn applies the NotIn predicate on the "other_phone" field.
func OtherPhoneNotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldOtherPhone, vs...))
}

// OtherPhoneGT applies the GT predicate on the "other_phone" field.
func OtherPhoneGT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldOtherPhone, v))
}

// OtherPhoneGTE applies the GTE predicate on the "other_phone" field.
func OtherPhoneGTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldOtherPhone, v))
}

// OtherPhoneLT applies the LT predicate on the "other_phone" field.
func OtherPhoneLT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldOtherPhone, v))
}

// OtherPhoneLTE applies the LTE predicate on the "other_phone" field.
func OtherPhoneLTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldOtherPhone, v))
}

// OtherPhoneContains applies the Contains predicate on the "other_phone" field.
func OtherPhoneContains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldOtherPhone, v))
}

// OtherPhoneHasPrefix applies the HasPrefix predicate on the "other_phone" field.
func OtherPhoneHasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldOtherPhone, v))
}

// OtherPhoneHasSuffix applies the HasSuffix predicate on the "other_phone" field.
func OtherPhoneHasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldOtherPhone, v))
}

// OtherPhoneIsNil applies the IsNil predicate on the "other_phone" field.
func OtherPhoneIsNil() predicate.Address {
	return predicate.Address(sql.FieldIsNull(FieldOtherPhone))
}

// OtherPhoneNotNil applies the NotNil predicate on the "other_phone" field.
func OtherPhoneNotNil() predicate.Address {
	return predicate.Address(sql.FieldNotNull(FieldOtherPhone))
}

// OtherPhoneEqualFold applies the EqualFold predicate on the "other_phone" field.
func OtherPhoneEqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldOtherPhone, v))
}

// OtherPhoneContainsFold applies the ContainsFold predicate on the "other_phone" field.
func OtherPhoneContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldOtherPhone, v))
}

// StreetNameEQ applies the EQ predicate on the "street_name" field.
func StreetNameEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldStreetName, v))
}

// StreetNameNEQ applies the NEQ predicate on the "street_name" field.
func StreetNameNEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldStreetName, v))
}

// StreetNameIn applies the In predicate on the "street_name" field.
func StreetNameIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldStreetName, vs...))
}

// StreetNameNotIn applies the NotIn predicate on the "street_name" field.
func StreetNameNotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldStreetName, vs...))
}

// StreetNameGT applies the GT predicate on the "street_name" field.
func StreetNameGT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldStreetName, v))
}

// StreetNameGTE applies the GTE predicate on the "street_name" field.
func StreetNameGTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldStreetName, v))
}

// StreetNameLT applies the LT predicate on the "street_name" field.
func StreetNameLT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldStreetName, v))
}

// StreetNameLTE applies the LTE predicate on the "street_name" field.
func StreetNameLTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldStreetName, v))
}

// StreetNameContains applies the Contains predicate on the "street_name" field.
func StreetNameContains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldStreetName, v))
}

// StreetNameHasPrefix applies the HasPrefix predicate on the "street_name" field.
func StreetNameHasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldStreetName, v))
}

// StreetNameHasSuffix applies the HasSuffix predicate on the "street_name" field.
func StreetNameHasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldStreetName, v))
}

// StreetNameIsNil applies the IsNil predicate on the "street_name" field.
func StreetNameIsNil() predicate.Address {
	return predicate.Address(sql.FieldIsNull(FieldStreetName))
}

// StreetNameNotNil applies the NotNil predicate on the "street_name" field.
func StreetNameNotNil() predicate.Address {
	return predicate.Address(sql.FieldNotNull(FieldStreetName))
}

// StreetNameEqualFold applies the EqualFold predicate on the "street_name" field.
func StreetNameEqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldStreetName, v))
}

// StreetNameContainsFold applies the ContainsFold predicate on the "street_name" field.
func StreetNameContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldStreetName, v))
}

// StreetNumberEQ applies the EQ predicate on the "street_number" field.
func StreetNumberEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldStreetNumber, v))
}

// StreetNumberNEQ applies the NEQ predicate on the "street_number" field.
func StreetNumberNEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldStreetNumber, v))
}

// StreetNumberIn applies the In predicate on the "street_number" field.
func StreetNumberIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldStreetNumber, vs...))
}

// StreetNumberNotIn applies the NotIn predicate on the "street_number" field.
func StreetNumberNotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldStreetNumber, vs...))
}

// StreetNumberGT applies the GT predicate on the "street_number" field.
func StreetNumberGT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldStreetNumber, v))
}

// StreetNumberGTE applies the GTE predicate on the "street_number" field.
func StreetNumberGTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldStreetNumber, v))
}

// StreetNumberLT applies the LT predicate on the "street_number" field.
func StreetNumberLT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldStreetNumber, v))
}

// StreetNumberLTE applies the LTE predicate on the "street_number" field.
func StreetNumberLTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldStreetNumber, v))
}

// StreetNumberContains applies the Contains predicate on the "street_number" field.
func StreetNumberContains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldStreetNumber, v))
}

// StreetNumberHasPrefix applies the HasPrefix predicate on the "street_number" field.
func StreetNumberHasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldStreetNumber, v))
}

// StreetNumberHasSuffix applies the HasSuffix predicate on the "street_number" field.
func StreetNumberHasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldStreetNumber, v))
}

// StreetNumberIsNil applies the IsNil predicate on the "street_number" field.
func StreetNumberIsNil() predicate.Address {
	return predicate.Address(sql.FieldIsNull(FieldStreetNumber))
}

// StreetNumberNotNil applies the NotNil predicate on the "street_number" field.
func StreetNumberNotNil() predicate.Address {
	return predicate.Address(sql.FieldNotNull(FieldStreetNumber))
}

// StreetNumberEqualFold applies the EqualFold predicate on the "street_number" field.
func StreetNumberEqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldStreetNumber, v))
}

// StreetNumberContainsFold applies the ContainsFold predicate on the "street_number" field.
func StreetNumberContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldStreetNumber, v))
}

// CityEQ applies the EQ predicate on the "city" field.
func CityEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldCity, v))
}

// CityNEQ applies the NEQ predicate on the "city" field.
func CityNEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldCity, v))
}

// CityIn applies the In predicate on the "city" field.
func CityIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldCity, vs...))
}

// CityNotIn applies the NotIn predicate on the "city" field.
func CityNotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldCity, vs...))
}

// CityGT applies the GT predicate on the "city" field.
func CityGT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldCity, v))
}

// CityGTE applies the GTE predicate on the "city" field.
func CityGTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldCity, v))
}

// CityLT applies the LT predicate on the "city" field.
func CityLT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldCity, v))
}

// CityLTE applies the LTE predicate on the "city" field.
func CityLTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldCity, v))
}

// CityContains applies the Contains predicate on the "city" field.
func CityContains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldCity, v))
}

// CityHasPrefix applies the HasPrefix predicate on the "city" field.
func CityHasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldCity, v))
}

// CityHasSuffix applies the HasSuffix predicate on the "city" field.
func CityHasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldCity, v))
}

// CityEqualFold applies the EqualFold predicate on the "city" field.
func CityEqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldCity, v))
}

// CityContainsFold applies the ContainsFold predicate on the "city" field.
func CityContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldCity, v))
}

// DistrictEQ applies the EQ predicate on the "district" field.
func DistrictEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldDistrict, v))
}

// DistrictNEQ applies the NEQ predicate on the "district" field.
func DistrictNEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldDistrict, v))
}

// DistrictIn applies the In predicate on the "district" field.
func DistrictIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldDistrict, vs...))
}

// DistrictNotIn applies the NotIn predicate on the "district" field.
func DistrictNotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldDistrict, vs...))
}

// DistrictGT applies the GT predicate on the "district" field.
func DistrictGT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldDistrict, v))
}

// DistrictGTE applies the GTE predicate on the "district" field.
func DistrictGTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldDistrict, v))
}

// DistrictLT applies the LT predicate on the "district" field.
func DistrictLT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldDistrict, v))
}

// DistrictLTE applies the LTE predicate on the "district" field.
func DistrictLTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldDistrict, v))
}

// DistrictContains applies the Contains predicate on the "district" field.
func DistrictContains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldDistrict, v))
}

// DistrictHasPrefix applies the HasPrefix predicate on the "district" field.
func DistrictHasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldDistrict, v))
}

// DistrictHasSuffix applies the HasSuffix predicate on the "district" field.
func DistrictHasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldDistrict, v))
}

// DistrictIsNil applies the IsNil predicate on the "district" field.
func DistrictIsNil() predicate.Address {
	return predicate.Address(sql.FieldIsNull(FieldDistrict))
}

// DistrictNotNil applies the NotNil predicate on the "district" field.
func DistrictNotNil() predicate.Address {
	return predicate.Address(sql.FieldNotNull(FieldDistrict))
}

// DistrictEqualFold applies the EqualFold predicate on the "district" field.
func DistrictEqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldDistrict, v))
}

// DistrictContainsFold applies the ContainsFold predicate on the "district" field.
func DistrictContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldDistrict, v))
}

// RegionEQ applies the EQ predicate on the "Region" field.
func RegionEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldRegion, v))
}

// RegionNEQ applies the NEQ predicate on the "Region" field.
func RegionNEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldRegion, v))
}

// RegionIn applies the In predicate on the "Region" field.
func RegionIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldRegion, vs...))
}

// RegionNotIn applies the NotIn predicate on the "Region" field.
func RegionNotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldRegion, vs...))
}

// RegionGT applies the GT predicate on the "Region" field.
func RegionGT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldRegion, v))
}

// RegionGTE applies the GTE predicate on the "Region" field.
func RegionGTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldRegion, v))
}

// RegionLT applies the LT predicate on the "Region" field.
func RegionLT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldRegion, v))
}

// RegionLTE applies the LTE predicate on the "Region" field.
func RegionLTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldRegion, v))
}

// RegionContains applies the Contains predicate on the "Region" field.
func RegionContains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldRegion, v))
}

// RegionHasPrefix applies the HasPrefix predicate on the "Region" field.
func RegionHasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldRegion, v))
}

// RegionHasSuffix applies the HasSuffix predicate on the "Region" field.
func RegionHasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldRegion, v))
}

// RegionEqualFold applies the EqualFold predicate on the "Region" field.
func RegionEqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldRegion, v))
}

// RegionContainsFold applies the ContainsFold predicate on the "Region" field.
func RegionContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldRegion, v))
}

// CountryEQ applies the EQ predicate on the "Country" field.
func CountryEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldCountry, v))
}

// CountryNEQ applies the NEQ predicate on the "Country" field.
func CountryNEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldCountry, v))
}

// CountryIn applies the In predicate on the "Country" field.
func CountryIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldCountry, vs...))
}

// CountryNotIn applies the NotIn predicate on the "Country" field.
func CountryNotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldCountry, vs...))
}

// CountryGT applies the GT predicate on the "Country" field.
func CountryGT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldCountry, v))
}

// CountryGTE applies the GTE predicate on the "Country" field.
func CountryGTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldCountry, v))
}

// CountryLT applies the LT predicate on the "Country" field.
func CountryLT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldCountry, v))
}

// CountryLTE applies the LTE predicate on the "Country" field.
func CountryLTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldCountry, v))
}

// CountryContains applies the Contains predicate on the "Country" field.
func CountryContains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldCountry, v))
}

// CountryHasPrefix applies the HasPrefix predicate on the "Country" field.
func CountryHasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldCountry, v))
}

// CountryHasSuffix applies the HasSuffix predicate on the "Country" field.
func CountryHasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldCountry, v))
}

// CountryEqualFold applies the EqualFold predicate on the "Country" field.
func CountryEqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldCountry, v))
}

// CountryContainsFold applies the ContainsFold predicate on the "Country" field.
func CountryContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldCountry, v))
}

// DefaultEQ applies the EQ predicate on the "default" field.
func DefaultEQ(v bool) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldDefault, v))
}

// DefaultNEQ applies the NEQ predicate on the "default" field.
func DefaultNEQ(v bool) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldDefault, v))
}

// CoordinateIsNil applies the IsNil predicate on the "coordinate" field.
func CoordinateIsNil() predicate.Address {
	return predicate.Address(sql.FieldIsNull(FieldCoordinate))
}

// CoordinateNotNil applies the NotNil predicate on the "coordinate" field.
func CoordinateNotNil() predicate.Address {
	return predicate.Address(sql.FieldNotNull(FieldCoordinate))
}

// HasMerchant applies the HasEdge predicate on the "merchant" edge.
func HasMerchant() predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MerchantTable, MerchantColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMerchantWith applies the HasEdge predicate on the "merchant" edge with a given conditions (other predicates).
func HasMerchantWith(preds ...predicate.Merchant) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		step := newMerchantStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAgent applies the HasEdge predicate on the "agent" edge.
func HasAgent() predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AgentTable, AgentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAgentWith applies the HasEdge predicate on the "agent" edge with a given conditions (other predicates).
func HasAgentWith(preds ...predicate.Agent) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		step := newAgentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCustomer applies the HasEdge predicate on the "customer" edge.
func HasCustomer() predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CustomerTable, CustomerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCustomerWith applies the HasEdge predicate on the "customer" edge with a given conditions (other predicates).
func HasCustomerWith(preds ...predicate.Customer) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		step := newCustomerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrders applies the HasEdge predicate on the "orders" edge.
func HasOrders() predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OrdersTable, OrdersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrdersWith applies the HasEdge predicate on the "orders" edge with a given conditions (other predicates).
func HasOrdersWith(preds ...predicate.Order) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		step := newOrdersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Address) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Address) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
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
func Not(p predicate.Address) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		p(s.Not())
	})
}
