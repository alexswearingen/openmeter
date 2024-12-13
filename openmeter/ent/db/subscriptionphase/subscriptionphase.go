// Code generated by ent, DO NOT EDIT.

package subscriptionphase

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the subscriptionphase type in the database.
	Label = "subscription_phase"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNamespace holds the string denoting the namespace field in the database.
	FieldNamespace = "namespace"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldMetadata holds the string denoting the metadata field in the database.
	FieldMetadata = "metadata"
	// FieldSubscriptionID holds the string denoting the subscription_id field in the database.
	FieldSubscriptionID = "subscription_id"
	// FieldKey holds the string denoting the key field in the database.
	FieldKey = "key"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldActiveFrom holds the string denoting the active_from field in the database.
	FieldActiveFrom = "active_from"
	// EdgeSubscription holds the string denoting the subscription edge name in mutations.
	EdgeSubscription = "subscription"
	// EdgeItems holds the string denoting the items edge name in mutations.
	EdgeItems = "items"
	// EdgeBillingLines holds the string denoting the billing_lines edge name in mutations.
	EdgeBillingLines = "billing_lines"
	// Table holds the table name of the subscriptionphase in the database.
	Table = "subscription_phases"
	// SubscriptionTable is the table that holds the subscription relation/edge.
	SubscriptionTable = "subscription_phases"
	// SubscriptionInverseTable is the table name for the Subscription entity.
	// It exists in this package in order to avoid circular dependency with the "subscription" package.
	SubscriptionInverseTable = "subscriptions"
	// SubscriptionColumn is the table column denoting the subscription relation/edge.
	SubscriptionColumn = "subscription_id"
	// ItemsTable is the table that holds the items relation/edge.
	ItemsTable = "subscription_items"
	// ItemsInverseTable is the table name for the SubscriptionItem entity.
	// It exists in this package in order to avoid circular dependency with the "subscriptionitem" package.
	ItemsInverseTable = "subscription_items"
	// ItemsColumn is the table column denoting the items relation/edge.
	ItemsColumn = "phase_id"
	// BillingLinesTable is the table that holds the billing_lines relation/edge.
	BillingLinesTable = "billing_invoice_lines"
	// BillingLinesInverseTable is the table name for the BillingInvoiceLine entity.
	// It exists in this package in order to avoid circular dependency with the "billinginvoiceline" package.
	BillingLinesInverseTable = "billing_invoice_lines"
	// BillingLinesColumn is the table column denoting the billing_lines relation/edge.
	BillingLinesColumn = "subscription_phase_id"
)

// Columns holds all SQL columns for subscriptionphase fields.
var Columns = []string{
	FieldID,
	FieldNamespace,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldMetadata,
	FieldSubscriptionID,
	FieldKey,
	FieldName,
	FieldDescription,
	FieldActiveFrom,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NamespaceValidator is a validator for the "namespace" field. It is called by the builders before save.
	NamespaceValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// SubscriptionIDValidator is a validator for the "subscription_id" field. It is called by the builders before save.
	SubscriptionIDValidator func(string) error
	// KeyValidator is a validator for the "key" field. It is called by the builders before save.
	KeyValidator func(string) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// OrderOption defines the ordering options for the SubscriptionPhase queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByNamespace orders the results by the namespace field.
func ByNamespace(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNamespace, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// BySubscriptionID orders the results by the subscription_id field.
func BySubscriptionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubscriptionID, opts...).ToFunc()
}

// ByKey orders the results by the key field.
func ByKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKey, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByActiveFrom orders the results by the active_from field.
func ByActiveFrom(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldActiveFrom, opts...).ToFunc()
}

// BySubscriptionField orders the results by subscription field.
func BySubscriptionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubscriptionStep(), sql.OrderByField(field, opts...))
	}
}

// ByItemsCount orders the results by items count.
func ByItemsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newItemsStep(), opts...)
	}
}

// ByItems orders the results by items terms.
func ByItems(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newItemsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByBillingLinesCount orders the results by billing_lines count.
func ByBillingLinesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newBillingLinesStep(), opts...)
	}
}

// ByBillingLines orders the results by billing_lines terms.
func ByBillingLines(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBillingLinesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newSubscriptionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SubscriptionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SubscriptionTable, SubscriptionColumn),
	)
}
func newItemsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ItemsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ItemsTable, ItemsColumn),
	)
}
func newBillingLinesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BillingLinesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, BillingLinesTable, BillingLinesColumn),
	)
}
