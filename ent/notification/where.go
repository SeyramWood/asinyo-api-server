// Code generated by ent, DO NOT EDIT.

package notification

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/SeyramWood/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Notification {
	return predicate.Notification(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Notification {
	return predicate.Notification(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Notification {
	return predicate.Notification(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Notification {
	return predicate.Notification(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Notification {
	return predicate.Notification(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Notification {
	return predicate.Notification(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Notification {
	return predicate.Notification(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldUpdatedAt, v))
}

// Event applies equality check predicate on the "event" field. It's identical to EventEQ.
func Event(v string) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldEvent, v))
}

// Activity applies equality check predicate on the "activity" field. It's identical to ActivityEQ.
func Activity(v string) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldActivity, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldDescription, v))
}

// SubjectType applies equality check predicate on the "subject_type" field. It's identical to SubjectTypeEQ.
func SubjectType(v string) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldSubjectType, v))
}

// SubjectID applies equality check predicate on the "subject_id" field. It's identical to SubjectIDEQ.
func SubjectID(v int) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldSubjectID, v))
}

// CreatorType applies equality check predicate on the "creator_type" field. It's identical to CreatorTypeEQ.
func CreatorType(v string) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldCreatorType, v))
}

// CreatorID applies equality check predicate on the "creator_id" field. It's identical to CreatorIDEQ.
func CreatorID(v int) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldCreatorID, v))
}

// CustomerReadAt applies equality check predicate on the "customer_read_at" field. It's identical to CustomerReadAtEQ.
func CustomerReadAt(v string) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldCustomerReadAt, v))
}

// AgentReadAt applies equality check predicate on the "agent_read_at" field. It's identical to AgentReadAtEQ.
func AgentReadAt(v string) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldAgentReadAt, v))
}

// MerchantReadAt applies equality check predicate on the "merchant_read_at" field. It's identical to MerchantReadAtEQ.
func MerchantReadAt(v string) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldMerchantReadAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Notification {
	return predicate.Notification(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Notification {
	return predicate.Notification(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Notification {
	return predicate.Notification(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Notification {
	return predicate.Notification(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Notification {
	return predicate.Notification(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Notification {
	return predicate.Notification(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Notification {
	return predicate.Notification(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Notification {
	return predicate.Notification(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Notification {
	return predicate.Notification(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Notification {
	return predicate.Notification(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Notification {
	return predicate.Notification(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Notification {
	return predicate.Notification(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Notification {
	return predicate.Notification(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Notification {
	return predicate.Notification(sql.FieldLTE(FieldUpdatedAt, v))
}

// EventEQ applies the EQ predicate on the "event" field.
func EventEQ(v string) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldEvent, v))
}

// EventNEQ applies the NEQ predicate on the "event" field.
func EventNEQ(v string) predicate.Notification {
	return predicate.Notification(sql.FieldNEQ(FieldEvent, v))
}

// EventIn applies the In predicate on the "event" field.
func EventIn(vs ...string) predicate.Notification {
	return predicate.Notification(sql.FieldIn(FieldEvent, vs...))
}

// EventNotIn applies the NotIn predicate on the "event" field.
func EventNotIn(vs ...string) predicate.Notification {
	return predicate.Notification(sql.FieldNotIn(FieldEvent, vs...))
}

// EventGT applies the GT predicate on the "event" field.
func EventGT(v string) predicate.Notification {
	return predicate.Notification(sql.FieldGT(FieldEvent, v))
}

// EventGTE applies the GTE predicate on the "event" field.
func EventGTE(v string) predicate.Notification {
	return predicate.Notification(sql.FieldGTE(FieldEvent, v))
}

// EventLT applies the LT predicate on the "event" field.
func EventLT(v string) predicate.Notification {
	return predicate.Notification(sql.FieldLT(FieldEvent, v))
}

// EventLTE applies the LTE predicate on the "event" field.
func EventLTE(v string) predicate.Notification {
	return predicate.Notification(sql.FieldLTE(FieldEvent, v))
}

// EventContains applies the Contains predicate on the "event" field.
func EventContains(v string) predicate.Notification {
	return predicate.Notification(sql.FieldContains(FieldEvent, v))
}

// EventHasPrefix applies the HasPrefix predicate on the "event" field.
func EventHasPrefix(v string) predicate.Notification {
	return predicate.Notification(sql.FieldHasPrefix(FieldEvent, v))
}

// EventHasSuffix applies the HasSuffix predicate on the "event" field.
func EventHasSuffix(v string) predicate.Notification {
	return predicate.Notification(sql.FieldHasSuffix(FieldEvent, v))
}

// EventEqualFold applies the EqualFold predicate on the "event" field.
func EventEqualFold(v string) predicate.Notification {
	return predicate.Notification(sql.FieldEqualFold(FieldEvent, v))
}

// EventContainsFold applies the ContainsFold predicate on the "event" field.
func EventContainsFold(v string) predicate.Notification {
	return predicate.Notification(sql.FieldContainsFold(FieldEvent, v))
}

// ActivityEQ applies the EQ predicate on the "activity" field.
func ActivityEQ(v string) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldActivity, v))
}

// ActivityNEQ applies the NEQ predicate on the "activity" field.
func ActivityNEQ(v string) predicate.Notification {
	return predicate.Notification(sql.FieldNEQ(FieldActivity, v))
}

// ActivityIn applies the In predicate on the "activity" field.
func ActivityIn(vs ...string) predicate.Notification {
	return predicate.Notification(sql.FieldIn(FieldActivity, vs...))
}

// ActivityNotIn applies the NotIn predicate on the "activity" field.
func ActivityNotIn(vs ...string) predicate.Notification {
	return predicate.Notification(sql.FieldNotIn(FieldActivity, vs...))
}

// ActivityGT applies the GT predicate on the "activity" field.
func ActivityGT(v string) predicate.Notification {
	return predicate.Notification(sql.FieldGT(FieldActivity, v))
}

// ActivityGTE applies the GTE predicate on the "activity" field.
func ActivityGTE(v string) predicate.Notification {
	return predicate.Notification(sql.FieldGTE(FieldActivity, v))
}

// ActivityLT applies the LT predicate on the "activity" field.
func ActivityLT(v string) predicate.Notification {
	return predicate.Notification(sql.FieldLT(FieldActivity, v))
}

// ActivityLTE applies the LTE predicate on the "activity" field.
func ActivityLTE(v string) predicate.Notification {
	return predicate.Notification(sql.FieldLTE(FieldActivity, v))
}

// ActivityContains applies the Contains predicate on the "activity" field.
func ActivityContains(v string) predicate.Notification {
	return predicate.Notification(sql.FieldContains(FieldActivity, v))
}

// ActivityHasPrefix applies the HasPrefix predicate on the "activity" field.
func ActivityHasPrefix(v string) predicate.Notification {
	return predicate.Notification(sql.FieldHasPrefix(FieldActivity, v))
}

// ActivityHasSuffix applies the HasSuffix predicate on the "activity" field.
func ActivityHasSuffix(v string) predicate.Notification {
	return predicate.Notification(sql.FieldHasSuffix(FieldActivity, v))
}

// ActivityEqualFold applies the EqualFold predicate on the "activity" field.
func ActivityEqualFold(v string) predicate.Notification {
	return predicate.Notification(sql.FieldEqualFold(FieldActivity, v))
}

// ActivityContainsFold applies the ContainsFold predicate on the "activity" field.
func ActivityContainsFold(v string) predicate.Notification {
	return predicate.Notification(sql.FieldContainsFold(FieldActivity, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Notification {
	return predicate.Notification(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Notification {
	return predicate.Notification(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Notification {
	return predicate.Notification(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Notification {
	return predicate.Notification(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Notification {
	return predicate.Notification(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Notification {
	return predicate.Notification(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Notification {
	return predicate.Notification(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Notification {
	return predicate.Notification(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Notification {
	return predicate.Notification(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Notification {
	return predicate.Notification(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Notification {
	return predicate.Notification(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Notification {
	return predicate.Notification(sql.FieldContainsFold(FieldDescription, v))
}

// SubjectTypeEQ applies the EQ predicate on the "subject_type" field.
func SubjectTypeEQ(v string) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldSubjectType, v))
}

// SubjectTypeNEQ applies the NEQ predicate on the "subject_type" field.
func SubjectTypeNEQ(v string) predicate.Notification {
	return predicate.Notification(sql.FieldNEQ(FieldSubjectType, v))
}

// SubjectTypeIn applies the In predicate on the "subject_type" field.
func SubjectTypeIn(vs ...string) predicate.Notification {
	return predicate.Notification(sql.FieldIn(FieldSubjectType, vs...))
}

// SubjectTypeNotIn applies the NotIn predicate on the "subject_type" field.
func SubjectTypeNotIn(vs ...string) predicate.Notification {
	return predicate.Notification(sql.FieldNotIn(FieldSubjectType, vs...))
}

// SubjectTypeGT applies the GT predicate on the "subject_type" field.
func SubjectTypeGT(v string) predicate.Notification {
	return predicate.Notification(sql.FieldGT(FieldSubjectType, v))
}

// SubjectTypeGTE applies the GTE predicate on the "subject_type" field.
func SubjectTypeGTE(v string) predicate.Notification {
	return predicate.Notification(sql.FieldGTE(FieldSubjectType, v))
}

// SubjectTypeLT applies the LT predicate on the "subject_type" field.
func SubjectTypeLT(v string) predicate.Notification {
	return predicate.Notification(sql.FieldLT(FieldSubjectType, v))
}

// SubjectTypeLTE applies the LTE predicate on the "subject_type" field.
func SubjectTypeLTE(v string) predicate.Notification {
	return predicate.Notification(sql.FieldLTE(FieldSubjectType, v))
}

// SubjectTypeContains applies the Contains predicate on the "subject_type" field.
func SubjectTypeContains(v string) predicate.Notification {
	return predicate.Notification(sql.FieldContains(FieldSubjectType, v))
}

// SubjectTypeHasPrefix applies the HasPrefix predicate on the "subject_type" field.
func SubjectTypeHasPrefix(v string) predicate.Notification {
	return predicate.Notification(sql.FieldHasPrefix(FieldSubjectType, v))
}

// SubjectTypeHasSuffix applies the HasSuffix predicate on the "subject_type" field.
func SubjectTypeHasSuffix(v string) predicate.Notification {
	return predicate.Notification(sql.FieldHasSuffix(FieldSubjectType, v))
}

// SubjectTypeEqualFold applies the EqualFold predicate on the "subject_type" field.
func SubjectTypeEqualFold(v string) predicate.Notification {
	return predicate.Notification(sql.FieldEqualFold(FieldSubjectType, v))
}

// SubjectTypeContainsFold applies the ContainsFold predicate on the "subject_type" field.
func SubjectTypeContainsFold(v string) predicate.Notification {
	return predicate.Notification(sql.FieldContainsFold(FieldSubjectType, v))
}

// SubjectIDEQ applies the EQ predicate on the "subject_id" field.
func SubjectIDEQ(v int) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldSubjectID, v))
}

// SubjectIDNEQ applies the NEQ predicate on the "subject_id" field.
func SubjectIDNEQ(v int) predicate.Notification {
	return predicate.Notification(sql.FieldNEQ(FieldSubjectID, v))
}

// SubjectIDIn applies the In predicate on the "subject_id" field.
func SubjectIDIn(vs ...int) predicate.Notification {
	return predicate.Notification(sql.FieldIn(FieldSubjectID, vs...))
}

// SubjectIDNotIn applies the NotIn predicate on the "subject_id" field.
func SubjectIDNotIn(vs ...int) predicate.Notification {
	return predicate.Notification(sql.FieldNotIn(FieldSubjectID, vs...))
}

// SubjectIDGT applies the GT predicate on the "subject_id" field.
func SubjectIDGT(v int) predicate.Notification {
	return predicate.Notification(sql.FieldGT(FieldSubjectID, v))
}

// SubjectIDGTE applies the GTE predicate on the "subject_id" field.
func SubjectIDGTE(v int) predicate.Notification {
	return predicate.Notification(sql.FieldGTE(FieldSubjectID, v))
}

// SubjectIDLT applies the LT predicate on the "subject_id" field.
func SubjectIDLT(v int) predicate.Notification {
	return predicate.Notification(sql.FieldLT(FieldSubjectID, v))
}

// SubjectIDLTE applies the LTE predicate on the "subject_id" field.
func SubjectIDLTE(v int) predicate.Notification {
	return predicate.Notification(sql.FieldLTE(FieldSubjectID, v))
}

// SubjectIDIsNil applies the IsNil predicate on the "subject_id" field.
func SubjectIDIsNil() predicate.Notification {
	return predicate.Notification(sql.FieldIsNull(FieldSubjectID))
}

// SubjectIDNotNil applies the NotNil predicate on the "subject_id" field.
func SubjectIDNotNil() predicate.Notification {
	return predicate.Notification(sql.FieldNotNull(FieldSubjectID))
}

// CreatorTypeEQ applies the EQ predicate on the "creator_type" field.
func CreatorTypeEQ(v string) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldCreatorType, v))
}

// CreatorTypeNEQ applies the NEQ predicate on the "creator_type" field.
func CreatorTypeNEQ(v string) predicate.Notification {
	return predicate.Notification(sql.FieldNEQ(FieldCreatorType, v))
}

// CreatorTypeIn applies the In predicate on the "creator_type" field.
func CreatorTypeIn(vs ...string) predicate.Notification {
	return predicate.Notification(sql.FieldIn(FieldCreatorType, vs...))
}

// CreatorTypeNotIn applies the NotIn predicate on the "creator_type" field.
func CreatorTypeNotIn(vs ...string) predicate.Notification {
	return predicate.Notification(sql.FieldNotIn(FieldCreatorType, vs...))
}

// CreatorTypeGT applies the GT predicate on the "creator_type" field.
func CreatorTypeGT(v string) predicate.Notification {
	return predicate.Notification(sql.FieldGT(FieldCreatorType, v))
}

// CreatorTypeGTE applies the GTE predicate on the "creator_type" field.
func CreatorTypeGTE(v string) predicate.Notification {
	return predicate.Notification(sql.FieldGTE(FieldCreatorType, v))
}

// CreatorTypeLT applies the LT predicate on the "creator_type" field.
func CreatorTypeLT(v string) predicate.Notification {
	return predicate.Notification(sql.FieldLT(FieldCreatorType, v))
}

// CreatorTypeLTE applies the LTE predicate on the "creator_type" field.
func CreatorTypeLTE(v string) predicate.Notification {
	return predicate.Notification(sql.FieldLTE(FieldCreatorType, v))
}

// CreatorTypeContains applies the Contains predicate on the "creator_type" field.
func CreatorTypeContains(v string) predicate.Notification {
	return predicate.Notification(sql.FieldContains(FieldCreatorType, v))
}

// CreatorTypeHasPrefix applies the HasPrefix predicate on the "creator_type" field.
func CreatorTypeHasPrefix(v string) predicate.Notification {
	return predicate.Notification(sql.FieldHasPrefix(FieldCreatorType, v))
}

// CreatorTypeHasSuffix applies the HasSuffix predicate on the "creator_type" field.
func CreatorTypeHasSuffix(v string) predicate.Notification {
	return predicate.Notification(sql.FieldHasSuffix(FieldCreatorType, v))
}

// CreatorTypeEqualFold applies the EqualFold predicate on the "creator_type" field.
func CreatorTypeEqualFold(v string) predicate.Notification {
	return predicate.Notification(sql.FieldEqualFold(FieldCreatorType, v))
}

// CreatorTypeContainsFold applies the ContainsFold predicate on the "creator_type" field.
func CreatorTypeContainsFold(v string) predicate.Notification {
	return predicate.Notification(sql.FieldContainsFold(FieldCreatorType, v))
}

// CreatorIDEQ applies the EQ predicate on the "creator_id" field.
func CreatorIDEQ(v int) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldCreatorID, v))
}

// CreatorIDNEQ applies the NEQ predicate on the "creator_id" field.
func CreatorIDNEQ(v int) predicate.Notification {
	return predicate.Notification(sql.FieldNEQ(FieldCreatorID, v))
}

// CreatorIDIn applies the In predicate on the "creator_id" field.
func CreatorIDIn(vs ...int) predicate.Notification {
	return predicate.Notification(sql.FieldIn(FieldCreatorID, vs...))
}

// CreatorIDNotIn applies the NotIn predicate on the "creator_id" field.
func CreatorIDNotIn(vs ...int) predicate.Notification {
	return predicate.Notification(sql.FieldNotIn(FieldCreatorID, vs...))
}

// CreatorIDGT applies the GT predicate on the "creator_id" field.
func CreatorIDGT(v int) predicate.Notification {
	return predicate.Notification(sql.FieldGT(FieldCreatorID, v))
}

// CreatorIDGTE applies the GTE predicate on the "creator_id" field.
func CreatorIDGTE(v int) predicate.Notification {
	return predicate.Notification(sql.FieldGTE(FieldCreatorID, v))
}

// CreatorIDLT applies the LT predicate on the "creator_id" field.
func CreatorIDLT(v int) predicate.Notification {
	return predicate.Notification(sql.FieldLT(FieldCreatorID, v))
}

// CreatorIDLTE applies the LTE predicate on the "creator_id" field.
func CreatorIDLTE(v int) predicate.Notification {
	return predicate.Notification(sql.FieldLTE(FieldCreatorID, v))
}

// CreatorIDIsNil applies the IsNil predicate on the "creator_id" field.
func CreatorIDIsNil() predicate.Notification {
	return predicate.Notification(sql.FieldIsNull(FieldCreatorID))
}

// CreatorIDNotNil applies the NotNil predicate on the "creator_id" field.
func CreatorIDNotNil() predicate.Notification {
	return predicate.Notification(sql.FieldNotNull(FieldCreatorID))
}

// CustomerReadAtEQ applies the EQ predicate on the "customer_read_at" field.
func CustomerReadAtEQ(v string) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldCustomerReadAt, v))
}

// CustomerReadAtNEQ applies the NEQ predicate on the "customer_read_at" field.
func CustomerReadAtNEQ(v string) predicate.Notification {
	return predicate.Notification(sql.FieldNEQ(FieldCustomerReadAt, v))
}

// CustomerReadAtIn applies the In predicate on the "customer_read_at" field.
func CustomerReadAtIn(vs ...string) predicate.Notification {
	return predicate.Notification(sql.FieldIn(FieldCustomerReadAt, vs...))
}

// CustomerReadAtNotIn applies the NotIn predicate on the "customer_read_at" field.
func CustomerReadAtNotIn(vs ...string) predicate.Notification {
	return predicate.Notification(sql.FieldNotIn(FieldCustomerReadAt, vs...))
}

// CustomerReadAtGT applies the GT predicate on the "customer_read_at" field.
func CustomerReadAtGT(v string) predicate.Notification {
	return predicate.Notification(sql.FieldGT(FieldCustomerReadAt, v))
}

// CustomerReadAtGTE applies the GTE predicate on the "customer_read_at" field.
func CustomerReadAtGTE(v string) predicate.Notification {
	return predicate.Notification(sql.FieldGTE(FieldCustomerReadAt, v))
}

// CustomerReadAtLT applies the LT predicate on the "customer_read_at" field.
func CustomerReadAtLT(v string) predicate.Notification {
	return predicate.Notification(sql.FieldLT(FieldCustomerReadAt, v))
}

// CustomerReadAtLTE applies the LTE predicate on the "customer_read_at" field.
func CustomerReadAtLTE(v string) predicate.Notification {
	return predicate.Notification(sql.FieldLTE(FieldCustomerReadAt, v))
}

// CustomerReadAtContains applies the Contains predicate on the "customer_read_at" field.
func CustomerReadAtContains(v string) predicate.Notification {
	return predicate.Notification(sql.FieldContains(FieldCustomerReadAt, v))
}

// CustomerReadAtHasPrefix applies the HasPrefix predicate on the "customer_read_at" field.
func CustomerReadAtHasPrefix(v string) predicate.Notification {
	return predicate.Notification(sql.FieldHasPrefix(FieldCustomerReadAt, v))
}

// CustomerReadAtHasSuffix applies the HasSuffix predicate on the "customer_read_at" field.
func CustomerReadAtHasSuffix(v string) predicate.Notification {
	return predicate.Notification(sql.FieldHasSuffix(FieldCustomerReadAt, v))
}

// CustomerReadAtIsNil applies the IsNil predicate on the "customer_read_at" field.
func CustomerReadAtIsNil() predicate.Notification {
	return predicate.Notification(sql.FieldIsNull(FieldCustomerReadAt))
}

// CustomerReadAtNotNil applies the NotNil predicate on the "customer_read_at" field.
func CustomerReadAtNotNil() predicate.Notification {
	return predicate.Notification(sql.FieldNotNull(FieldCustomerReadAt))
}

// CustomerReadAtEqualFold applies the EqualFold predicate on the "customer_read_at" field.
func CustomerReadAtEqualFold(v string) predicate.Notification {
	return predicate.Notification(sql.FieldEqualFold(FieldCustomerReadAt, v))
}

// CustomerReadAtContainsFold applies the ContainsFold predicate on the "customer_read_at" field.
func CustomerReadAtContainsFold(v string) predicate.Notification {
	return predicate.Notification(sql.FieldContainsFold(FieldCustomerReadAt, v))
}

// AgentReadAtEQ applies the EQ predicate on the "agent_read_at" field.
func AgentReadAtEQ(v string) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldAgentReadAt, v))
}

// AgentReadAtNEQ applies the NEQ predicate on the "agent_read_at" field.
func AgentReadAtNEQ(v string) predicate.Notification {
	return predicate.Notification(sql.FieldNEQ(FieldAgentReadAt, v))
}

// AgentReadAtIn applies the In predicate on the "agent_read_at" field.
func AgentReadAtIn(vs ...string) predicate.Notification {
	return predicate.Notification(sql.FieldIn(FieldAgentReadAt, vs...))
}

// AgentReadAtNotIn applies the NotIn predicate on the "agent_read_at" field.
func AgentReadAtNotIn(vs ...string) predicate.Notification {
	return predicate.Notification(sql.FieldNotIn(FieldAgentReadAt, vs...))
}

// AgentReadAtGT applies the GT predicate on the "agent_read_at" field.
func AgentReadAtGT(v string) predicate.Notification {
	return predicate.Notification(sql.FieldGT(FieldAgentReadAt, v))
}

// AgentReadAtGTE applies the GTE predicate on the "agent_read_at" field.
func AgentReadAtGTE(v string) predicate.Notification {
	return predicate.Notification(sql.FieldGTE(FieldAgentReadAt, v))
}

// AgentReadAtLT applies the LT predicate on the "agent_read_at" field.
func AgentReadAtLT(v string) predicate.Notification {
	return predicate.Notification(sql.FieldLT(FieldAgentReadAt, v))
}

// AgentReadAtLTE applies the LTE predicate on the "agent_read_at" field.
func AgentReadAtLTE(v string) predicate.Notification {
	return predicate.Notification(sql.FieldLTE(FieldAgentReadAt, v))
}

// AgentReadAtContains applies the Contains predicate on the "agent_read_at" field.
func AgentReadAtContains(v string) predicate.Notification {
	return predicate.Notification(sql.FieldContains(FieldAgentReadAt, v))
}

// AgentReadAtHasPrefix applies the HasPrefix predicate on the "agent_read_at" field.
func AgentReadAtHasPrefix(v string) predicate.Notification {
	return predicate.Notification(sql.FieldHasPrefix(FieldAgentReadAt, v))
}

// AgentReadAtHasSuffix applies the HasSuffix predicate on the "agent_read_at" field.
func AgentReadAtHasSuffix(v string) predicate.Notification {
	return predicate.Notification(sql.FieldHasSuffix(FieldAgentReadAt, v))
}

// AgentReadAtIsNil applies the IsNil predicate on the "agent_read_at" field.
func AgentReadAtIsNil() predicate.Notification {
	return predicate.Notification(sql.FieldIsNull(FieldAgentReadAt))
}

// AgentReadAtNotNil applies the NotNil predicate on the "agent_read_at" field.
func AgentReadAtNotNil() predicate.Notification {
	return predicate.Notification(sql.FieldNotNull(FieldAgentReadAt))
}

// AgentReadAtEqualFold applies the EqualFold predicate on the "agent_read_at" field.
func AgentReadAtEqualFold(v string) predicate.Notification {
	return predicate.Notification(sql.FieldEqualFold(FieldAgentReadAt, v))
}

// AgentReadAtContainsFold applies the ContainsFold predicate on the "agent_read_at" field.
func AgentReadAtContainsFold(v string) predicate.Notification {
	return predicate.Notification(sql.FieldContainsFold(FieldAgentReadAt, v))
}

// MerchantReadAtEQ applies the EQ predicate on the "merchant_read_at" field.
func MerchantReadAtEQ(v string) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldMerchantReadAt, v))
}

// MerchantReadAtNEQ applies the NEQ predicate on the "merchant_read_at" field.
func MerchantReadAtNEQ(v string) predicate.Notification {
	return predicate.Notification(sql.FieldNEQ(FieldMerchantReadAt, v))
}

// MerchantReadAtIn applies the In predicate on the "merchant_read_at" field.
func MerchantReadAtIn(vs ...string) predicate.Notification {
	return predicate.Notification(sql.FieldIn(FieldMerchantReadAt, vs...))
}

// MerchantReadAtNotIn applies the NotIn predicate on the "merchant_read_at" field.
func MerchantReadAtNotIn(vs ...string) predicate.Notification {
	return predicate.Notification(sql.FieldNotIn(FieldMerchantReadAt, vs...))
}

// MerchantReadAtGT applies the GT predicate on the "merchant_read_at" field.
func MerchantReadAtGT(v string) predicate.Notification {
	return predicate.Notification(sql.FieldGT(FieldMerchantReadAt, v))
}

// MerchantReadAtGTE applies the GTE predicate on the "merchant_read_at" field.
func MerchantReadAtGTE(v string) predicate.Notification {
	return predicate.Notification(sql.FieldGTE(FieldMerchantReadAt, v))
}

// MerchantReadAtLT applies the LT predicate on the "merchant_read_at" field.
func MerchantReadAtLT(v string) predicate.Notification {
	return predicate.Notification(sql.FieldLT(FieldMerchantReadAt, v))
}

// MerchantReadAtLTE applies the LTE predicate on the "merchant_read_at" field.
func MerchantReadAtLTE(v string) predicate.Notification {
	return predicate.Notification(sql.FieldLTE(FieldMerchantReadAt, v))
}

// MerchantReadAtContains applies the Contains predicate on the "merchant_read_at" field.
func MerchantReadAtContains(v string) predicate.Notification {
	return predicate.Notification(sql.FieldContains(FieldMerchantReadAt, v))
}

// MerchantReadAtHasPrefix applies the HasPrefix predicate on the "merchant_read_at" field.
func MerchantReadAtHasPrefix(v string) predicate.Notification {
	return predicate.Notification(sql.FieldHasPrefix(FieldMerchantReadAt, v))
}

// MerchantReadAtHasSuffix applies the HasSuffix predicate on the "merchant_read_at" field.
func MerchantReadAtHasSuffix(v string) predicate.Notification {
	return predicate.Notification(sql.FieldHasSuffix(FieldMerchantReadAt, v))
}

// MerchantReadAtIsNil applies the IsNil predicate on the "merchant_read_at" field.
func MerchantReadAtIsNil() predicate.Notification {
	return predicate.Notification(sql.FieldIsNull(FieldMerchantReadAt))
}

// MerchantReadAtNotNil applies the NotNil predicate on the "merchant_read_at" field.
func MerchantReadAtNotNil() predicate.Notification {
	return predicate.Notification(sql.FieldNotNull(FieldMerchantReadAt))
}

// MerchantReadAtEqualFold applies the EqualFold predicate on the "merchant_read_at" field.
func MerchantReadAtEqualFold(v string) predicate.Notification {
	return predicate.Notification(sql.FieldEqualFold(FieldMerchantReadAt, v))
}

// MerchantReadAtContainsFold applies the ContainsFold predicate on the "merchant_read_at" field.
func MerchantReadAtContainsFold(v string) predicate.Notification {
	return predicate.Notification(sql.FieldContainsFold(FieldMerchantReadAt, v))
}

// AdminReadAtIsNil applies the IsNil predicate on the "admin_read_at" field.
func AdminReadAtIsNil() predicate.Notification {
	return predicate.Notification(sql.FieldIsNull(FieldAdminReadAt))
}

// AdminReadAtNotNil applies the NotNil predicate on the "admin_read_at" field.
func AdminReadAtNotNil() predicate.Notification {
	return predicate.Notification(sql.FieldNotNull(FieldAdminReadAt))
}

// DataIsNil applies the IsNil predicate on the "data" field.
func DataIsNil() predicate.Notification {
	return predicate.Notification(sql.FieldIsNull(FieldData))
}

// DataNotNil applies the NotNil predicate on the "data" field.
func DataNotNil() predicate.Notification {
	return predicate.Notification(sql.FieldNotNull(FieldData))
}

// HasAdmin applies the HasEdge predicate on the "admin" edge.
func HasAdmin() predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, AdminTable, AdminPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAdminWith applies the HasEdge predicate on the "admin" edge with a given conditions (other predicates).
func HasAdminWith(preds ...predicate.Admin) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		step := newAdminStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMerchant applies the HasEdge predicate on the "merchant" edge.
func HasMerchant() predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, MerchantTable, MerchantPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMerchantWith applies the HasEdge predicate on the "merchant" edge with a given conditions (other predicates).
func HasMerchantWith(preds ...predicate.Merchant) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		step := newMerchantStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAgent applies the HasEdge predicate on the "agent" edge.
func HasAgent() predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, AgentTable, AgentPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAgentWith applies the HasEdge predicate on the "agent" edge with a given conditions (other predicates).
func HasAgentWith(preds ...predicate.Agent) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		step := newAgentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCustomer applies the HasEdge predicate on the "customer" edge.
func HasCustomer() predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, CustomerTable, CustomerPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCustomerWith applies the HasEdge predicate on the "customer" edge with a given conditions (other predicates).
func HasCustomerWith(preds ...predicate.Customer) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		step := newCustomerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Notification) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Notification) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
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
func Not(p predicate.Notification) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		p(s.Not())
	})
}
