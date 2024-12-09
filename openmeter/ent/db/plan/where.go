// Code generated by ent, DO NOT EDIT.

package plan

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/openmeterio/openmeter/openmeter/ent/db/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Plan {
	return predicate.Plan(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Plan {
	return predicate.Plan(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Plan {
	return predicate.Plan(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Plan {
	return predicate.Plan(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Plan {
	return predicate.Plan(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Plan {
	return predicate.Plan(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Plan {
	return predicate.Plan(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Plan {
	return predicate.Plan(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Plan {
	return predicate.Plan(sql.FieldContainsFold(FieldID, id))
}

// Namespace applies equality check predicate on the "namespace" field. It's identical to NamespaceEQ.
func Namespace(v string) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldNamespace, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldDeletedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldDescription, v))
}

// Key applies equality check predicate on the "key" field. It's identical to KeyEQ.
func Key(v string) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldKey, v))
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v int) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldVersion, v))
}

// Currency applies equality check predicate on the "currency" field. It's identical to CurrencyEQ.
func Currency(v string) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldCurrency, v))
}

// EffectiveFrom applies equality check predicate on the "effective_from" field. It's identical to EffectiveFromEQ.
func EffectiveFrom(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldEffectiveFrom, v))
}

// EffectiveTo applies equality check predicate on the "effective_to" field. It's identical to EffectiveToEQ.
func EffectiveTo(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldEffectiveTo, v))
}

// NamespaceEQ applies the EQ predicate on the "namespace" field.
func NamespaceEQ(v string) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldNamespace, v))
}

// NamespaceNEQ applies the NEQ predicate on the "namespace" field.
func NamespaceNEQ(v string) predicate.Plan {
	return predicate.Plan(sql.FieldNEQ(FieldNamespace, v))
}

// NamespaceIn applies the In predicate on the "namespace" field.
func NamespaceIn(vs ...string) predicate.Plan {
	return predicate.Plan(sql.FieldIn(FieldNamespace, vs...))
}

// NamespaceNotIn applies the NotIn predicate on the "namespace" field.
func NamespaceNotIn(vs ...string) predicate.Plan {
	return predicate.Plan(sql.FieldNotIn(FieldNamespace, vs...))
}

// NamespaceGT applies the GT predicate on the "namespace" field.
func NamespaceGT(v string) predicate.Plan {
	return predicate.Plan(sql.FieldGT(FieldNamespace, v))
}

// NamespaceGTE applies the GTE predicate on the "namespace" field.
func NamespaceGTE(v string) predicate.Plan {
	return predicate.Plan(sql.FieldGTE(FieldNamespace, v))
}

// NamespaceLT applies the LT predicate on the "namespace" field.
func NamespaceLT(v string) predicate.Plan {
	return predicate.Plan(sql.FieldLT(FieldNamespace, v))
}

// NamespaceLTE applies the LTE predicate on the "namespace" field.
func NamespaceLTE(v string) predicate.Plan {
	return predicate.Plan(sql.FieldLTE(FieldNamespace, v))
}

// NamespaceContains applies the Contains predicate on the "namespace" field.
func NamespaceContains(v string) predicate.Plan {
	return predicate.Plan(sql.FieldContains(FieldNamespace, v))
}

// NamespaceHasPrefix applies the HasPrefix predicate on the "namespace" field.
func NamespaceHasPrefix(v string) predicate.Plan {
	return predicate.Plan(sql.FieldHasPrefix(FieldNamespace, v))
}

// NamespaceHasSuffix applies the HasSuffix predicate on the "namespace" field.
func NamespaceHasSuffix(v string) predicate.Plan {
	return predicate.Plan(sql.FieldHasSuffix(FieldNamespace, v))
}

// NamespaceEqualFold applies the EqualFold predicate on the "namespace" field.
func NamespaceEqualFold(v string) predicate.Plan {
	return predicate.Plan(sql.FieldEqualFold(FieldNamespace, v))
}

// NamespaceContainsFold applies the ContainsFold predicate on the "namespace" field.
func NamespaceContainsFold(v string) predicate.Plan {
	return predicate.Plan(sql.FieldContainsFold(FieldNamespace, v))
}

// MetadataIsNil applies the IsNil predicate on the "metadata" field.
func MetadataIsNil() predicate.Plan {
	return predicate.Plan(sql.FieldIsNull(FieldMetadata))
}

// MetadataNotNil applies the NotNil predicate on the "metadata" field.
func MetadataNotNil() predicate.Plan {
	return predicate.Plan(sql.FieldNotNull(FieldMetadata))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Plan {
	return predicate.Plan(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Plan {
	return predicate.Plan(sql.FieldNotNull(FieldDeletedAt))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Plan {
	return predicate.Plan(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Plan {
	return predicate.Plan(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Plan {
	return predicate.Plan(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Plan {
	return predicate.Plan(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Plan {
	return predicate.Plan(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Plan {
	return predicate.Plan(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Plan {
	return predicate.Plan(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Plan {
	return predicate.Plan(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Plan {
	return predicate.Plan(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Plan {
	return predicate.Plan(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Plan {
	return predicate.Plan(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Plan {
	return predicate.Plan(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Plan {
	return predicate.Plan(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Plan {
	return predicate.Plan(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Plan {
	return predicate.Plan(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Plan {
	return predicate.Plan(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Plan {
	return predicate.Plan(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Plan {
	return predicate.Plan(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Plan {
	return predicate.Plan(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Plan {
	return predicate.Plan(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Plan {
	return predicate.Plan(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Plan {
	return predicate.Plan(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Plan {
	return predicate.Plan(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Plan {
	return predicate.Plan(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Plan {
	return predicate.Plan(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Plan {
	return predicate.Plan(sql.FieldContainsFold(FieldDescription, v))
}

// KeyEQ applies the EQ predicate on the "key" field.
func KeyEQ(v string) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldKey, v))
}

// KeyNEQ applies the NEQ predicate on the "key" field.
func KeyNEQ(v string) predicate.Plan {
	return predicate.Plan(sql.FieldNEQ(FieldKey, v))
}

// KeyIn applies the In predicate on the "key" field.
func KeyIn(vs ...string) predicate.Plan {
	return predicate.Plan(sql.FieldIn(FieldKey, vs...))
}

// KeyNotIn applies the NotIn predicate on the "key" field.
func KeyNotIn(vs ...string) predicate.Plan {
	return predicate.Plan(sql.FieldNotIn(FieldKey, vs...))
}

// KeyGT applies the GT predicate on the "key" field.
func KeyGT(v string) predicate.Plan {
	return predicate.Plan(sql.FieldGT(FieldKey, v))
}

// KeyGTE applies the GTE predicate on the "key" field.
func KeyGTE(v string) predicate.Plan {
	return predicate.Plan(sql.FieldGTE(FieldKey, v))
}

// KeyLT applies the LT predicate on the "key" field.
func KeyLT(v string) predicate.Plan {
	return predicate.Plan(sql.FieldLT(FieldKey, v))
}

// KeyLTE applies the LTE predicate on the "key" field.
func KeyLTE(v string) predicate.Plan {
	return predicate.Plan(sql.FieldLTE(FieldKey, v))
}

// KeyContains applies the Contains predicate on the "key" field.
func KeyContains(v string) predicate.Plan {
	return predicate.Plan(sql.FieldContains(FieldKey, v))
}

// KeyHasPrefix applies the HasPrefix predicate on the "key" field.
func KeyHasPrefix(v string) predicate.Plan {
	return predicate.Plan(sql.FieldHasPrefix(FieldKey, v))
}

// KeyHasSuffix applies the HasSuffix predicate on the "key" field.
func KeyHasSuffix(v string) predicate.Plan {
	return predicate.Plan(sql.FieldHasSuffix(FieldKey, v))
}

// KeyEqualFold applies the EqualFold predicate on the "key" field.
func KeyEqualFold(v string) predicate.Plan {
	return predicate.Plan(sql.FieldEqualFold(FieldKey, v))
}

// KeyContainsFold applies the ContainsFold predicate on the "key" field.
func KeyContainsFold(v string) predicate.Plan {
	return predicate.Plan(sql.FieldContainsFold(FieldKey, v))
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v int) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldVersion, v))
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v int) predicate.Plan {
	return predicate.Plan(sql.FieldNEQ(FieldVersion, v))
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...int) predicate.Plan {
	return predicate.Plan(sql.FieldIn(FieldVersion, vs...))
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...int) predicate.Plan {
	return predicate.Plan(sql.FieldNotIn(FieldVersion, vs...))
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v int) predicate.Plan {
	return predicate.Plan(sql.FieldGT(FieldVersion, v))
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v int) predicate.Plan {
	return predicate.Plan(sql.FieldGTE(FieldVersion, v))
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v int) predicate.Plan {
	return predicate.Plan(sql.FieldLT(FieldVersion, v))
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v int) predicate.Plan {
	return predicate.Plan(sql.FieldLTE(FieldVersion, v))
}

// CurrencyEQ applies the EQ predicate on the "currency" field.
func CurrencyEQ(v string) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldCurrency, v))
}

// CurrencyNEQ applies the NEQ predicate on the "currency" field.
func CurrencyNEQ(v string) predicate.Plan {
	return predicate.Plan(sql.FieldNEQ(FieldCurrency, v))
}

// CurrencyIn applies the In predicate on the "currency" field.
func CurrencyIn(vs ...string) predicate.Plan {
	return predicate.Plan(sql.FieldIn(FieldCurrency, vs...))
}

// CurrencyNotIn applies the NotIn predicate on the "currency" field.
func CurrencyNotIn(vs ...string) predicate.Plan {
	return predicate.Plan(sql.FieldNotIn(FieldCurrency, vs...))
}

// CurrencyGT applies the GT predicate on the "currency" field.
func CurrencyGT(v string) predicate.Plan {
	return predicate.Plan(sql.FieldGT(FieldCurrency, v))
}

// CurrencyGTE applies the GTE predicate on the "currency" field.
func CurrencyGTE(v string) predicate.Plan {
	return predicate.Plan(sql.FieldGTE(FieldCurrency, v))
}

// CurrencyLT applies the LT predicate on the "currency" field.
func CurrencyLT(v string) predicate.Plan {
	return predicate.Plan(sql.FieldLT(FieldCurrency, v))
}

// CurrencyLTE applies the LTE predicate on the "currency" field.
func CurrencyLTE(v string) predicate.Plan {
	return predicate.Plan(sql.FieldLTE(FieldCurrency, v))
}

// CurrencyContains applies the Contains predicate on the "currency" field.
func CurrencyContains(v string) predicate.Plan {
	return predicate.Plan(sql.FieldContains(FieldCurrency, v))
}

// CurrencyHasPrefix applies the HasPrefix predicate on the "currency" field.
func CurrencyHasPrefix(v string) predicate.Plan {
	return predicate.Plan(sql.FieldHasPrefix(FieldCurrency, v))
}

// CurrencyHasSuffix applies the HasSuffix predicate on the "currency" field.
func CurrencyHasSuffix(v string) predicate.Plan {
	return predicate.Plan(sql.FieldHasSuffix(FieldCurrency, v))
}

// CurrencyEqualFold applies the EqualFold predicate on the "currency" field.
func CurrencyEqualFold(v string) predicate.Plan {
	return predicate.Plan(sql.FieldEqualFold(FieldCurrency, v))
}

// CurrencyContainsFold applies the ContainsFold predicate on the "currency" field.
func CurrencyContainsFold(v string) predicate.Plan {
	return predicate.Plan(sql.FieldContainsFold(FieldCurrency, v))
}

// EffectiveFromEQ applies the EQ predicate on the "effective_from" field.
func EffectiveFromEQ(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldEffectiveFrom, v))
}

// EffectiveFromNEQ applies the NEQ predicate on the "effective_from" field.
func EffectiveFromNEQ(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldNEQ(FieldEffectiveFrom, v))
}

// EffectiveFromIn applies the In predicate on the "effective_from" field.
func EffectiveFromIn(vs ...time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldIn(FieldEffectiveFrom, vs...))
}

// EffectiveFromNotIn applies the NotIn predicate on the "effective_from" field.
func EffectiveFromNotIn(vs ...time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldNotIn(FieldEffectiveFrom, vs...))
}

// EffectiveFromGT applies the GT predicate on the "effective_from" field.
func EffectiveFromGT(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldGT(FieldEffectiveFrom, v))
}

// EffectiveFromGTE applies the GTE predicate on the "effective_from" field.
func EffectiveFromGTE(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldGTE(FieldEffectiveFrom, v))
}

// EffectiveFromLT applies the LT predicate on the "effective_from" field.
func EffectiveFromLT(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldLT(FieldEffectiveFrom, v))
}

// EffectiveFromLTE applies the LTE predicate on the "effective_from" field.
func EffectiveFromLTE(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldLTE(FieldEffectiveFrom, v))
}

// EffectiveFromIsNil applies the IsNil predicate on the "effective_from" field.
func EffectiveFromIsNil() predicate.Plan {
	return predicate.Plan(sql.FieldIsNull(FieldEffectiveFrom))
}

// EffectiveFromNotNil applies the NotNil predicate on the "effective_from" field.
func EffectiveFromNotNil() predicate.Plan {
	return predicate.Plan(sql.FieldNotNull(FieldEffectiveFrom))
}

// EffectiveToEQ applies the EQ predicate on the "effective_to" field.
func EffectiveToEQ(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldEQ(FieldEffectiveTo, v))
}

// EffectiveToNEQ applies the NEQ predicate on the "effective_to" field.
func EffectiveToNEQ(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldNEQ(FieldEffectiveTo, v))
}

// EffectiveToIn applies the In predicate on the "effective_to" field.
func EffectiveToIn(vs ...time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldIn(FieldEffectiveTo, vs...))
}

// EffectiveToNotIn applies the NotIn predicate on the "effective_to" field.
func EffectiveToNotIn(vs ...time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldNotIn(FieldEffectiveTo, vs...))
}

// EffectiveToGT applies the GT predicate on the "effective_to" field.
func EffectiveToGT(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldGT(FieldEffectiveTo, v))
}

// EffectiveToGTE applies the GTE predicate on the "effective_to" field.
func EffectiveToGTE(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldGTE(FieldEffectiveTo, v))
}

// EffectiveToLT applies the LT predicate on the "effective_to" field.
func EffectiveToLT(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldLT(FieldEffectiveTo, v))
}

// EffectiveToLTE applies the LTE predicate on the "effective_to" field.
func EffectiveToLTE(v time.Time) predicate.Plan {
	return predicate.Plan(sql.FieldLTE(FieldEffectiveTo, v))
}

// EffectiveToIsNil applies the IsNil predicate on the "effective_to" field.
func EffectiveToIsNil() predicate.Plan {
	return predicate.Plan(sql.FieldIsNull(FieldEffectiveTo))
}

// EffectiveToNotNil applies the NotNil predicate on the "effective_to" field.
func EffectiveToNotNil() predicate.Plan {
	return predicate.Plan(sql.FieldNotNull(FieldEffectiveTo))
}

// HasPhases applies the HasEdge predicate on the "phases" edge.
func HasPhases() predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PhasesTable, PhasesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPhasesWith applies the HasEdge predicate on the "phases" edge with a given conditions (other predicates).
func HasPhasesWith(preds ...predicate.PlanPhase) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		step := newPhasesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSubscriptions applies the HasEdge predicate on the "subscriptions" edge.
func HasSubscriptions() predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SubscriptionsTable, SubscriptionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubscriptionsWith applies the HasEdge predicate on the "subscriptions" edge with a given conditions (other predicates).
func HasSubscriptionsWith(preds ...predicate.Subscription) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		step := newSubscriptionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Plan) predicate.Plan {
	return predicate.Plan(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Plan) predicate.Plan {
	return predicate.Plan(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Plan) predicate.Plan {
	return predicate.Plan(sql.NotPredicates(p))
}
