// Code generated by ent, DO NOT EDIT.

package billinginvoice

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/alpacahq/alpacadecimal"
	"github.com/openmeterio/openmeter/openmeter/billing/invoice"
	"github.com/openmeterio/openmeter/openmeter/ent/db/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldContainsFold(FieldID, id))
}

// Namespace applies equality check predicate on the "namespace" field. It's identical to NamespaceEQ.
func Namespace(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldEQ(FieldNamespace, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldEQ(FieldDeletedAt, v))
}

// Key applies equality check predicate on the "key" field. It's identical to KeyEQ.
func Key(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldEQ(FieldKey, v))
}

// CustomerID applies equality check predicate on the "customer_id" field. It's identical to CustomerIDEQ.
func CustomerID(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldEQ(FieldCustomerID, v))
}

// BillingProfileID applies equality check predicate on the "billing_profile_id" field. It's identical to BillingProfileIDEQ.
func BillingProfileID(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldEQ(FieldBillingProfileID, v))
}

// VoidedAt applies equality check predicate on the "voided_at" field. It's identical to VoidedAtEQ.
func VoidedAt(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldEQ(FieldVoidedAt, v))
}

// Currency applies equality check predicate on the "currency" field. It's identical to CurrencyEQ.
func Currency(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldEQ(FieldCurrency, v))
}

// TotalAmount applies equality check predicate on the "total_amount" field. It's identical to TotalAmountEQ.
func TotalAmount(v alpacadecimal.Decimal) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldEQ(FieldTotalAmount, v))
}

// DueDate applies equality check predicate on the "due_date" field. It's identical to DueDateEQ.
func DueDate(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldEQ(FieldDueDate, v))
}

// PeriodStart applies equality check predicate on the "period_start" field. It's identical to PeriodStartEQ.
func PeriodStart(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldEQ(FieldPeriodStart, v))
}

// PeriodEnd applies equality check predicate on the "period_end" field. It's identical to PeriodEndEQ.
func PeriodEnd(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldEQ(FieldPeriodEnd, v))
}

// NamespaceEQ applies the EQ predicate on the "namespace" field.
func NamespaceEQ(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldEQ(FieldNamespace, v))
}

// NamespaceNEQ applies the NEQ predicate on the "namespace" field.
func NamespaceNEQ(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldNEQ(FieldNamespace, v))
}

// NamespaceIn applies the In predicate on the "namespace" field.
func NamespaceIn(vs ...string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldIn(FieldNamespace, vs...))
}

// NamespaceNotIn applies the NotIn predicate on the "namespace" field.
func NamespaceNotIn(vs ...string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldNotIn(FieldNamespace, vs...))
}

// NamespaceGT applies the GT predicate on the "namespace" field.
func NamespaceGT(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldGT(FieldNamespace, v))
}

// NamespaceGTE applies the GTE predicate on the "namespace" field.
func NamespaceGTE(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldGTE(FieldNamespace, v))
}

// NamespaceLT applies the LT predicate on the "namespace" field.
func NamespaceLT(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldLT(FieldNamespace, v))
}

// NamespaceLTE applies the LTE predicate on the "namespace" field.
func NamespaceLTE(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldLTE(FieldNamespace, v))
}

// NamespaceContains applies the Contains predicate on the "namespace" field.
func NamespaceContains(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldContains(FieldNamespace, v))
}

// NamespaceHasPrefix applies the HasPrefix predicate on the "namespace" field.
func NamespaceHasPrefix(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldHasPrefix(FieldNamespace, v))
}

// NamespaceHasSuffix applies the HasSuffix predicate on the "namespace" field.
func NamespaceHasSuffix(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldHasSuffix(FieldNamespace, v))
}

// NamespaceEqualFold applies the EqualFold predicate on the "namespace" field.
func NamespaceEqualFold(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldEqualFold(FieldNamespace, v))
}

// NamespaceContainsFold applies the ContainsFold predicate on the "namespace" field.
func NamespaceContainsFold(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldContainsFold(FieldNamespace, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldNotNull(FieldDeletedAt))
}

// MetadataIsNil applies the IsNil predicate on the "metadata" field.
func MetadataIsNil() predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldIsNull(FieldMetadata))
}

// MetadataNotNil applies the NotNil predicate on the "metadata" field.
func MetadataNotNil() predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldNotNull(FieldMetadata))
}

// KeyEQ applies the EQ predicate on the "key" field.
func KeyEQ(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldEQ(FieldKey, v))
}

// KeyNEQ applies the NEQ predicate on the "key" field.
func KeyNEQ(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldNEQ(FieldKey, v))
}

// KeyIn applies the In predicate on the "key" field.
func KeyIn(vs ...string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldIn(FieldKey, vs...))
}

// KeyNotIn applies the NotIn predicate on the "key" field.
func KeyNotIn(vs ...string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldNotIn(FieldKey, vs...))
}

// KeyGT applies the GT predicate on the "key" field.
func KeyGT(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldGT(FieldKey, v))
}

// KeyGTE applies the GTE predicate on the "key" field.
func KeyGTE(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldGTE(FieldKey, v))
}

// KeyLT applies the LT predicate on the "key" field.
func KeyLT(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldLT(FieldKey, v))
}

// KeyLTE applies the LTE predicate on the "key" field.
func KeyLTE(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldLTE(FieldKey, v))
}

// KeyContains applies the Contains predicate on the "key" field.
func KeyContains(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldContains(FieldKey, v))
}

// KeyHasPrefix applies the HasPrefix predicate on the "key" field.
func KeyHasPrefix(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldHasPrefix(FieldKey, v))
}

// KeyHasSuffix applies the HasSuffix predicate on the "key" field.
func KeyHasSuffix(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldHasSuffix(FieldKey, v))
}

// KeyEqualFold applies the EqualFold predicate on the "key" field.
func KeyEqualFold(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldEqualFold(FieldKey, v))
}

// KeyContainsFold applies the ContainsFold predicate on the "key" field.
func KeyContainsFold(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldContainsFold(FieldKey, v))
}

// CustomerIDEQ applies the EQ predicate on the "customer_id" field.
func CustomerIDEQ(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldEQ(FieldCustomerID, v))
}

// CustomerIDNEQ applies the NEQ predicate on the "customer_id" field.
func CustomerIDNEQ(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldNEQ(FieldCustomerID, v))
}

// CustomerIDIn applies the In predicate on the "customer_id" field.
func CustomerIDIn(vs ...string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldIn(FieldCustomerID, vs...))
}

// CustomerIDNotIn applies the NotIn predicate on the "customer_id" field.
func CustomerIDNotIn(vs ...string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldNotIn(FieldCustomerID, vs...))
}

// CustomerIDGT applies the GT predicate on the "customer_id" field.
func CustomerIDGT(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldGT(FieldCustomerID, v))
}

// CustomerIDGTE applies the GTE predicate on the "customer_id" field.
func CustomerIDGTE(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldGTE(FieldCustomerID, v))
}

// CustomerIDLT applies the LT predicate on the "customer_id" field.
func CustomerIDLT(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldLT(FieldCustomerID, v))
}

// CustomerIDLTE applies the LTE predicate on the "customer_id" field.
func CustomerIDLTE(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldLTE(FieldCustomerID, v))
}

// CustomerIDContains applies the Contains predicate on the "customer_id" field.
func CustomerIDContains(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldContains(FieldCustomerID, v))
}

// CustomerIDHasPrefix applies the HasPrefix predicate on the "customer_id" field.
func CustomerIDHasPrefix(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldHasPrefix(FieldCustomerID, v))
}

// CustomerIDHasSuffix applies the HasSuffix predicate on the "customer_id" field.
func CustomerIDHasSuffix(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldHasSuffix(FieldCustomerID, v))
}

// CustomerIDEqualFold applies the EqualFold predicate on the "customer_id" field.
func CustomerIDEqualFold(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldEqualFold(FieldCustomerID, v))
}

// CustomerIDContainsFold applies the ContainsFold predicate on the "customer_id" field.
func CustomerIDContainsFold(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldContainsFold(FieldCustomerID, v))
}

// BillingProfileIDEQ applies the EQ predicate on the "billing_profile_id" field.
func BillingProfileIDEQ(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldEQ(FieldBillingProfileID, v))
}

// BillingProfileIDNEQ applies the NEQ predicate on the "billing_profile_id" field.
func BillingProfileIDNEQ(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldNEQ(FieldBillingProfileID, v))
}

// BillingProfileIDIn applies the In predicate on the "billing_profile_id" field.
func BillingProfileIDIn(vs ...string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldIn(FieldBillingProfileID, vs...))
}

// BillingProfileIDNotIn applies the NotIn predicate on the "billing_profile_id" field.
func BillingProfileIDNotIn(vs ...string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldNotIn(FieldBillingProfileID, vs...))
}

// BillingProfileIDGT applies the GT predicate on the "billing_profile_id" field.
func BillingProfileIDGT(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldGT(FieldBillingProfileID, v))
}

// BillingProfileIDGTE applies the GTE predicate on the "billing_profile_id" field.
func BillingProfileIDGTE(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldGTE(FieldBillingProfileID, v))
}

// BillingProfileIDLT applies the LT predicate on the "billing_profile_id" field.
func BillingProfileIDLT(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldLT(FieldBillingProfileID, v))
}

// BillingProfileIDLTE applies the LTE predicate on the "billing_profile_id" field.
func BillingProfileIDLTE(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldLTE(FieldBillingProfileID, v))
}

// BillingProfileIDContains applies the Contains predicate on the "billing_profile_id" field.
func BillingProfileIDContains(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldContains(FieldBillingProfileID, v))
}

// BillingProfileIDHasPrefix applies the HasPrefix predicate on the "billing_profile_id" field.
func BillingProfileIDHasPrefix(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldHasPrefix(FieldBillingProfileID, v))
}

// BillingProfileIDHasSuffix applies the HasSuffix predicate on the "billing_profile_id" field.
func BillingProfileIDHasSuffix(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldHasSuffix(FieldBillingProfileID, v))
}

// BillingProfileIDEqualFold applies the EqualFold predicate on the "billing_profile_id" field.
func BillingProfileIDEqualFold(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldEqualFold(FieldBillingProfileID, v))
}

// BillingProfileIDContainsFold applies the ContainsFold predicate on the "billing_profile_id" field.
func BillingProfileIDContainsFold(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldContainsFold(FieldBillingProfileID, v))
}

// VoidedAtEQ applies the EQ predicate on the "voided_at" field.
func VoidedAtEQ(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldEQ(FieldVoidedAt, v))
}

// VoidedAtNEQ applies the NEQ predicate on the "voided_at" field.
func VoidedAtNEQ(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldNEQ(FieldVoidedAt, v))
}

// VoidedAtIn applies the In predicate on the "voided_at" field.
func VoidedAtIn(vs ...time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldIn(FieldVoidedAt, vs...))
}

// VoidedAtNotIn applies the NotIn predicate on the "voided_at" field.
func VoidedAtNotIn(vs ...time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldNotIn(FieldVoidedAt, vs...))
}

// VoidedAtGT applies the GT predicate on the "voided_at" field.
func VoidedAtGT(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldGT(FieldVoidedAt, v))
}

// VoidedAtGTE applies the GTE predicate on the "voided_at" field.
func VoidedAtGTE(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldGTE(FieldVoidedAt, v))
}

// VoidedAtLT applies the LT predicate on the "voided_at" field.
func VoidedAtLT(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldLT(FieldVoidedAt, v))
}

// VoidedAtLTE applies the LTE predicate on the "voided_at" field.
func VoidedAtLTE(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldLTE(FieldVoidedAt, v))
}

// VoidedAtIsNil applies the IsNil predicate on the "voided_at" field.
func VoidedAtIsNil() predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldIsNull(FieldVoidedAt))
}

// VoidedAtNotNil applies the NotNil predicate on the "voided_at" field.
func VoidedAtNotNil() predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldNotNull(FieldVoidedAt))
}

// CurrencyEQ applies the EQ predicate on the "currency" field.
func CurrencyEQ(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldEQ(FieldCurrency, v))
}

// CurrencyNEQ applies the NEQ predicate on the "currency" field.
func CurrencyNEQ(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldNEQ(FieldCurrency, v))
}

// CurrencyIn applies the In predicate on the "currency" field.
func CurrencyIn(vs ...string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldIn(FieldCurrency, vs...))
}

// CurrencyNotIn applies the NotIn predicate on the "currency" field.
func CurrencyNotIn(vs ...string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldNotIn(FieldCurrency, vs...))
}

// CurrencyGT applies the GT predicate on the "currency" field.
func CurrencyGT(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldGT(FieldCurrency, v))
}

// CurrencyGTE applies the GTE predicate on the "currency" field.
func CurrencyGTE(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldGTE(FieldCurrency, v))
}

// CurrencyLT applies the LT predicate on the "currency" field.
func CurrencyLT(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldLT(FieldCurrency, v))
}

// CurrencyLTE applies the LTE predicate on the "currency" field.
func CurrencyLTE(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldLTE(FieldCurrency, v))
}

// CurrencyContains applies the Contains predicate on the "currency" field.
func CurrencyContains(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldContains(FieldCurrency, v))
}

// CurrencyHasPrefix applies the HasPrefix predicate on the "currency" field.
func CurrencyHasPrefix(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldHasPrefix(FieldCurrency, v))
}

// CurrencyHasSuffix applies the HasSuffix predicate on the "currency" field.
func CurrencyHasSuffix(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldHasSuffix(FieldCurrency, v))
}

// CurrencyEqualFold applies the EqualFold predicate on the "currency" field.
func CurrencyEqualFold(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldEqualFold(FieldCurrency, v))
}

// CurrencyContainsFold applies the ContainsFold predicate on the "currency" field.
func CurrencyContainsFold(v string) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldContainsFold(FieldCurrency, v))
}

// TotalAmountEQ applies the EQ predicate on the "total_amount" field.
func TotalAmountEQ(v alpacadecimal.Decimal) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldEQ(FieldTotalAmount, v))
}

// TotalAmountNEQ applies the NEQ predicate on the "total_amount" field.
func TotalAmountNEQ(v alpacadecimal.Decimal) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldNEQ(FieldTotalAmount, v))
}

// TotalAmountIn applies the In predicate on the "total_amount" field.
func TotalAmountIn(vs ...alpacadecimal.Decimal) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldIn(FieldTotalAmount, vs...))
}

// TotalAmountNotIn applies the NotIn predicate on the "total_amount" field.
func TotalAmountNotIn(vs ...alpacadecimal.Decimal) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldNotIn(FieldTotalAmount, vs...))
}

// TotalAmountGT applies the GT predicate on the "total_amount" field.
func TotalAmountGT(v alpacadecimal.Decimal) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldGT(FieldTotalAmount, v))
}

// TotalAmountGTE applies the GTE predicate on the "total_amount" field.
func TotalAmountGTE(v alpacadecimal.Decimal) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldGTE(FieldTotalAmount, v))
}

// TotalAmountLT applies the LT predicate on the "total_amount" field.
func TotalAmountLT(v alpacadecimal.Decimal) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldLT(FieldTotalAmount, v))
}

// TotalAmountLTE applies the LTE predicate on the "total_amount" field.
func TotalAmountLTE(v alpacadecimal.Decimal) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldLTE(FieldTotalAmount, v))
}

// DueDateEQ applies the EQ predicate on the "due_date" field.
func DueDateEQ(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldEQ(FieldDueDate, v))
}

// DueDateNEQ applies the NEQ predicate on the "due_date" field.
func DueDateNEQ(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldNEQ(FieldDueDate, v))
}

// DueDateIn applies the In predicate on the "due_date" field.
func DueDateIn(vs ...time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldIn(FieldDueDate, vs...))
}

// DueDateNotIn applies the NotIn predicate on the "due_date" field.
func DueDateNotIn(vs ...time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldNotIn(FieldDueDate, vs...))
}

// DueDateGT applies the GT predicate on the "due_date" field.
func DueDateGT(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldGT(FieldDueDate, v))
}

// DueDateGTE applies the GTE predicate on the "due_date" field.
func DueDateGTE(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldGTE(FieldDueDate, v))
}

// DueDateLT applies the LT predicate on the "due_date" field.
func DueDateLT(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldLT(FieldDueDate, v))
}

// DueDateLTE applies the LTE predicate on the "due_date" field.
func DueDateLTE(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldLTE(FieldDueDate, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v invoice.InvoiceStatus) predicate.BillingInvoice {
	vc := v
	return predicate.BillingInvoice(sql.FieldEQ(FieldStatus, vc))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v invoice.InvoiceStatus) predicate.BillingInvoice {
	vc := v
	return predicate.BillingInvoice(sql.FieldNEQ(FieldStatus, vc))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...invoice.InvoiceStatus) predicate.BillingInvoice {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BillingInvoice(sql.FieldIn(FieldStatus, v...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...invoice.InvoiceStatus) predicate.BillingInvoice {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BillingInvoice(sql.FieldNotIn(FieldStatus, v...))
}

// PeriodStartEQ applies the EQ predicate on the "period_start" field.
func PeriodStartEQ(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldEQ(FieldPeriodStart, v))
}

// PeriodStartNEQ applies the NEQ predicate on the "period_start" field.
func PeriodStartNEQ(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldNEQ(FieldPeriodStart, v))
}

// PeriodStartIn applies the In predicate on the "period_start" field.
func PeriodStartIn(vs ...time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldIn(FieldPeriodStart, vs...))
}

// PeriodStartNotIn applies the NotIn predicate on the "period_start" field.
func PeriodStartNotIn(vs ...time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldNotIn(FieldPeriodStart, vs...))
}

// PeriodStartGT applies the GT predicate on the "period_start" field.
func PeriodStartGT(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldGT(FieldPeriodStart, v))
}

// PeriodStartGTE applies the GTE predicate on the "period_start" field.
func PeriodStartGTE(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldGTE(FieldPeriodStart, v))
}

// PeriodStartLT applies the LT predicate on the "period_start" field.
func PeriodStartLT(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldLT(FieldPeriodStart, v))
}

// PeriodStartLTE applies the LTE predicate on the "period_start" field.
func PeriodStartLTE(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldLTE(FieldPeriodStart, v))
}

// PeriodEndEQ applies the EQ predicate on the "period_end" field.
func PeriodEndEQ(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldEQ(FieldPeriodEnd, v))
}

// PeriodEndNEQ applies the NEQ predicate on the "period_end" field.
func PeriodEndNEQ(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldNEQ(FieldPeriodEnd, v))
}

// PeriodEndIn applies the In predicate on the "period_end" field.
func PeriodEndIn(vs ...time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldIn(FieldPeriodEnd, vs...))
}

// PeriodEndNotIn applies the NotIn predicate on the "period_end" field.
func PeriodEndNotIn(vs ...time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldNotIn(FieldPeriodEnd, vs...))
}

// PeriodEndGT applies the GT predicate on the "period_end" field.
func PeriodEndGT(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldGT(FieldPeriodEnd, v))
}

// PeriodEndGTE applies the GTE predicate on the "period_end" field.
func PeriodEndGTE(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldGTE(FieldPeriodEnd, v))
}

// PeriodEndLT applies the LT predicate on the "period_end" field.
func PeriodEndLT(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldLT(FieldPeriodEnd, v))
}

// PeriodEndLTE applies the LTE predicate on the "period_end" field.
func PeriodEndLTE(v time.Time) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.FieldLTE(FieldPeriodEnd, v))
}

// HasBillingProfile applies the HasEdge predicate on the "billing_profile" edge.
func HasBillingProfile() predicate.BillingInvoice {
	return predicate.BillingInvoice(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BillingProfileTable, BillingProfileColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBillingProfileWith applies the HasEdge predicate on the "billing_profile" edge with a given conditions (other predicates).
func HasBillingProfileWith(preds ...predicate.BillingProfile) predicate.BillingInvoice {
	return predicate.BillingInvoice(func(s *sql.Selector) {
		step := newBillingProfileStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBillingInvoiceItems applies the HasEdge predicate on the "billing_invoice_items" edge.
func HasBillingInvoiceItems() predicate.BillingInvoice {
	return predicate.BillingInvoice(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BillingInvoiceItemsTable, BillingInvoiceItemsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBillingInvoiceItemsWith applies the HasEdge predicate on the "billing_invoice_items" edge with a given conditions (other predicates).
func HasBillingInvoiceItemsWith(preds ...predicate.BillingInvoiceItem) predicate.BillingInvoice {
	return predicate.BillingInvoice(func(s *sql.Selector) {
		step := newBillingInvoiceItemsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.BillingInvoice) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BillingInvoice) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.BillingInvoice) predicate.BillingInvoice {
	return predicate.BillingInvoice(sql.NotPredicates(p))
}
