// Code generated by ent, DO NOT EDIT.

package app

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	appentitybase "github.com/openmeterio/openmeter/openmeter/app/entity/base"
	"github.com/openmeterio/openmeter/openmeter/ent/db/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.App {
	return predicate.App(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.App {
	return predicate.App(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.App {
	return predicate.App(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.App {
	return predicate.App(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.App {
	return predicate.App(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.App {
	return predicate.App(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.App {
	return predicate.App(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.App {
	return predicate.App(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.App {
	return predicate.App(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.App {
	return predicate.App(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.App {
	return predicate.App(sql.FieldContainsFold(FieldID, id))
}

// Namespace applies equality check predicate on the "namespace" field. It's identical to NamespaceEQ.
func Namespace(v string) predicate.App {
	return predicate.App(sql.FieldEQ(FieldNamespace, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.App {
	return predicate.App(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.App {
	return predicate.App(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.App {
	return predicate.App(sql.FieldEQ(FieldDeletedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.App {
	return predicate.App(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.App {
	return predicate.App(sql.FieldEQ(FieldDescription, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v appentitybase.AppType) predicate.App {
	vc := string(v)
	return predicate.App(sql.FieldEQ(FieldType, vc))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v appentitybase.AppStatus) predicate.App {
	vc := string(v)
	return predicate.App(sql.FieldEQ(FieldStatus, vc))
}

// IsDefault applies equality check predicate on the "is_default" field. It's identical to IsDefaultEQ.
func IsDefault(v bool) predicate.App {
	return predicate.App(sql.FieldEQ(FieldIsDefault, v))
}

// NamespaceEQ applies the EQ predicate on the "namespace" field.
func NamespaceEQ(v string) predicate.App {
	return predicate.App(sql.FieldEQ(FieldNamespace, v))
}

// NamespaceNEQ applies the NEQ predicate on the "namespace" field.
func NamespaceNEQ(v string) predicate.App {
	return predicate.App(sql.FieldNEQ(FieldNamespace, v))
}

// NamespaceIn applies the In predicate on the "namespace" field.
func NamespaceIn(vs ...string) predicate.App {
	return predicate.App(sql.FieldIn(FieldNamespace, vs...))
}

// NamespaceNotIn applies the NotIn predicate on the "namespace" field.
func NamespaceNotIn(vs ...string) predicate.App {
	return predicate.App(sql.FieldNotIn(FieldNamespace, vs...))
}

// NamespaceGT applies the GT predicate on the "namespace" field.
func NamespaceGT(v string) predicate.App {
	return predicate.App(sql.FieldGT(FieldNamespace, v))
}

// NamespaceGTE applies the GTE predicate on the "namespace" field.
func NamespaceGTE(v string) predicate.App {
	return predicate.App(sql.FieldGTE(FieldNamespace, v))
}

// NamespaceLT applies the LT predicate on the "namespace" field.
func NamespaceLT(v string) predicate.App {
	return predicate.App(sql.FieldLT(FieldNamespace, v))
}

// NamespaceLTE applies the LTE predicate on the "namespace" field.
func NamespaceLTE(v string) predicate.App {
	return predicate.App(sql.FieldLTE(FieldNamespace, v))
}

// NamespaceContains applies the Contains predicate on the "namespace" field.
func NamespaceContains(v string) predicate.App {
	return predicate.App(sql.FieldContains(FieldNamespace, v))
}

// NamespaceHasPrefix applies the HasPrefix predicate on the "namespace" field.
func NamespaceHasPrefix(v string) predicate.App {
	return predicate.App(sql.FieldHasPrefix(FieldNamespace, v))
}

// NamespaceHasSuffix applies the HasSuffix predicate on the "namespace" field.
func NamespaceHasSuffix(v string) predicate.App {
	return predicate.App(sql.FieldHasSuffix(FieldNamespace, v))
}

// NamespaceEqualFold applies the EqualFold predicate on the "namespace" field.
func NamespaceEqualFold(v string) predicate.App {
	return predicate.App(sql.FieldEqualFold(FieldNamespace, v))
}

// NamespaceContainsFold applies the ContainsFold predicate on the "namespace" field.
func NamespaceContainsFold(v string) predicate.App {
	return predicate.App(sql.FieldContainsFold(FieldNamespace, v))
}

// MetadataIsNil applies the IsNil predicate on the "metadata" field.
func MetadataIsNil() predicate.App {
	return predicate.App(sql.FieldIsNull(FieldMetadata))
}

// MetadataNotNil applies the NotNil predicate on the "metadata" field.
func MetadataNotNil() predicate.App {
	return predicate.App(sql.FieldNotNull(FieldMetadata))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.App {
	return predicate.App(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.App {
	return predicate.App(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.App {
	return predicate.App(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.App {
	return predicate.App(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.App {
	return predicate.App(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.App {
	return predicate.App(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.App {
	return predicate.App(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.App {
	return predicate.App(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.App {
	return predicate.App(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.App {
	return predicate.App(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.App {
	return predicate.App(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.App {
	return predicate.App(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.App {
	return predicate.App(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.App {
	return predicate.App(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.App {
	return predicate.App(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.App {
	return predicate.App(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.App {
	return predicate.App(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.App {
	return predicate.App(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.App {
	return predicate.App(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.App {
	return predicate.App(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.App {
	return predicate.App(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.App {
	return predicate.App(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.App {
	return predicate.App(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.App {
	return predicate.App(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.App {
	return predicate.App(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.App {
	return predicate.App(sql.FieldNotNull(FieldDeletedAt))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.App {
	return predicate.App(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.App {
	return predicate.App(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.App {
	return predicate.App(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.App {
	return predicate.App(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.App {
	return predicate.App(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.App {
	return predicate.App(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.App {
	return predicate.App(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.App {
	return predicate.App(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.App {
	return predicate.App(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.App {
	return predicate.App(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.App {
	return predicate.App(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.App {
	return predicate.App(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.App {
	return predicate.App(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.App {
	return predicate.App(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.App {
	return predicate.App(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.App {
	return predicate.App(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.App {
	return predicate.App(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.App {
	return predicate.App(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.App {
	return predicate.App(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.App {
	return predicate.App(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.App {
	return predicate.App(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.App {
	return predicate.App(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.App {
	return predicate.App(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.App {
	return predicate.App(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.App {
	return predicate.App(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.App {
	return predicate.App(sql.FieldContainsFold(FieldDescription, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v appentitybase.AppType) predicate.App {
	vc := string(v)
	return predicate.App(sql.FieldEQ(FieldType, vc))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v appentitybase.AppType) predicate.App {
	vc := string(v)
	return predicate.App(sql.FieldNEQ(FieldType, vc))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...appentitybase.AppType) predicate.App {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.App(sql.FieldIn(FieldType, v...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...appentitybase.AppType) predicate.App {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.App(sql.FieldNotIn(FieldType, v...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v appentitybase.AppType) predicate.App {
	vc := string(v)
	return predicate.App(sql.FieldGT(FieldType, vc))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v appentitybase.AppType) predicate.App {
	vc := string(v)
	return predicate.App(sql.FieldGTE(FieldType, vc))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v appentitybase.AppType) predicate.App {
	vc := string(v)
	return predicate.App(sql.FieldLT(FieldType, vc))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v appentitybase.AppType) predicate.App {
	vc := string(v)
	return predicate.App(sql.FieldLTE(FieldType, vc))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v appentitybase.AppType) predicate.App {
	vc := string(v)
	return predicate.App(sql.FieldContains(FieldType, vc))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v appentitybase.AppType) predicate.App {
	vc := string(v)
	return predicate.App(sql.FieldHasPrefix(FieldType, vc))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v appentitybase.AppType) predicate.App {
	vc := string(v)
	return predicate.App(sql.FieldHasSuffix(FieldType, vc))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v appentitybase.AppType) predicate.App {
	vc := string(v)
	return predicate.App(sql.FieldEqualFold(FieldType, vc))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v appentitybase.AppType) predicate.App {
	vc := string(v)
	return predicate.App(sql.FieldContainsFold(FieldType, vc))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v appentitybase.AppStatus) predicate.App {
	vc := string(v)
	return predicate.App(sql.FieldEQ(FieldStatus, vc))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v appentitybase.AppStatus) predicate.App {
	vc := string(v)
	return predicate.App(sql.FieldNEQ(FieldStatus, vc))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...appentitybase.AppStatus) predicate.App {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.App(sql.FieldIn(FieldStatus, v...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...appentitybase.AppStatus) predicate.App {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.App(sql.FieldNotIn(FieldStatus, v...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v appentitybase.AppStatus) predicate.App {
	vc := string(v)
	return predicate.App(sql.FieldGT(FieldStatus, vc))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v appentitybase.AppStatus) predicate.App {
	vc := string(v)
	return predicate.App(sql.FieldGTE(FieldStatus, vc))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v appentitybase.AppStatus) predicate.App {
	vc := string(v)
	return predicate.App(sql.FieldLT(FieldStatus, vc))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v appentitybase.AppStatus) predicate.App {
	vc := string(v)
	return predicate.App(sql.FieldLTE(FieldStatus, vc))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v appentitybase.AppStatus) predicate.App {
	vc := string(v)
	return predicate.App(sql.FieldContains(FieldStatus, vc))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v appentitybase.AppStatus) predicate.App {
	vc := string(v)
	return predicate.App(sql.FieldHasPrefix(FieldStatus, vc))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v appentitybase.AppStatus) predicate.App {
	vc := string(v)
	return predicate.App(sql.FieldHasSuffix(FieldStatus, vc))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v appentitybase.AppStatus) predicate.App {
	vc := string(v)
	return predicate.App(sql.FieldEqualFold(FieldStatus, vc))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v appentitybase.AppStatus) predicate.App {
	vc := string(v)
	return predicate.App(sql.FieldContainsFold(FieldStatus, vc))
}

// IsDefaultEQ applies the EQ predicate on the "is_default" field.
func IsDefaultEQ(v bool) predicate.App {
	return predicate.App(sql.FieldEQ(FieldIsDefault, v))
}

// IsDefaultNEQ applies the NEQ predicate on the "is_default" field.
func IsDefaultNEQ(v bool) predicate.App {
	return predicate.App(sql.FieldNEQ(FieldIsDefault, v))
}

// HasCustomerApps applies the HasEdge predicate on the "customer_apps" edge.
func HasCustomerApps() predicate.App {
	return predicate.App(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CustomerAppsTable, CustomerAppsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCustomerAppsWith applies the HasEdge predicate on the "customer_apps" edge with a given conditions (other predicates).
func HasCustomerAppsWith(preds ...predicate.AppCustomer) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		step := newCustomerAppsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.App) predicate.App {
	return predicate.App(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.App) predicate.App {
	return predicate.App(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.App) predicate.App {
	return predicate.App(sql.NotPredicates(p))
}
