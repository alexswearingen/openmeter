// Code generated by ent, DO NOT EDIT.

package billinginvoice

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/openmeterio/openmeter/openmeter/billing/invoice"
	"github.com/openmeterio/openmeter/openmeter/billing/provider"
)

const (
	// Label holds the string label denoting the billinginvoice type in the database.
	Label = "billing_invoice"
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
	// FieldKey holds the string denoting the key field in the database.
	FieldKey = "key"
	// FieldCustomerID holds the string denoting the customer_id field in the database.
	FieldCustomerID = "customer_id"
	// FieldBillingProfileID holds the string denoting the billing_profile_id field in the database.
	FieldBillingProfileID = "billing_profile_id"
	// FieldVoidedAt holds the string denoting the voided_at field in the database.
	FieldVoidedAt = "voided_at"
	// FieldCurrency holds the string denoting the currency field in the database.
	FieldCurrency = "currency"
	// FieldTotalAmount holds the string denoting the total_amount field in the database.
	FieldTotalAmount = "total_amount"
	// FieldDueDate holds the string denoting the due_date field in the database.
	FieldDueDate = "due_date"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldProviderConfig holds the string denoting the provider_config field in the database.
	FieldProviderConfig = "provider_config"
	// FieldBillingConfig holds the string denoting the billing_config field in the database.
	FieldBillingConfig = "billing_config"
	// FieldProviderReference holds the string denoting the provider_reference field in the database.
	FieldProviderReference = "provider_reference"
	// FieldPeriodStart holds the string denoting the period_start field in the database.
	FieldPeriodStart = "period_start"
	// FieldPeriodEnd holds the string denoting the period_end field in the database.
	FieldPeriodEnd = "period_end"
	// EdgeBillingProfile holds the string denoting the billing_profile edge name in mutations.
	EdgeBillingProfile = "billing_profile"
	// EdgeBillingInvoiceItems holds the string denoting the billing_invoice_items edge name in mutations.
	EdgeBillingInvoiceItems = "billing_invoice_items"
	// Table holds the table name of the billinginvoice in the database.
	Table = "billing_invoices"
	// BillingProfileTable is the table that holds the billing_profile relation/edge.
	BillingProfileTable = "billing_invoices"
	// BillingProfileInverseTable is the table name for the BillingProfile entity.
	// It exists in this package in order to avoid circular dependency with the "billingprofile" package.
	BillingProfileInverseTable = "billing_profiles"
	// BillingProfileColumn is the table column denoting the billing_profile relation/edge.
	BillingProfileColumn = "billing_profile_id"
	// BillingInvoiceItemsTable is the table that holds the billing_invoice_items relation/edge.
	BillingInvoiceItemsTable = "billing_invoice_items"
	// BillingInvoiceItemsInverseTable is the table name for the BillingInvoiceItem entity.
	// It exists in this package in order to avoid circular dependency with the "billinginvoiceitem" package.
	BillingInvoiceItemsInverseTable = "billing_invoice_items"
	// BillingInvoiceItemsColumn is the table column denoting the billing_invoice_items relation/edge.
	BillingInvoiceItemsColumn = "invoice_id"
)

// Columns holds all SQL columns for billinginvoice fields.
var Columns = []string{
	FieldID,
	FieldNamespace,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldMetadata,
	FieldKey,
	FieldCustomerID,
	FieldBillingProfileID,
	FieldVoidedAt,
	FieldCurrency,
	FieldTotalAmount,
	FieldDueDate,
	FieldStatus,
	FieldProviderConfig,
	FieldBillingConfig,
	FieldProviderReference,
	FieldPeriodStart,
	FieldPeriodEnd,
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
	// KeyValidator is a validator for the "key" field. It is called by the builders before save.
	KeyValidator func(string) error
	// CustomerIDValidator is a validator for the "customer_id" field. It is called by the builders before save.
	CustomerIDValidator func(string) error
	// BillingProfileIDValidator is a validator for the "billing_profile_id" field. It is called by the builders before save.
	BillingProfileIDValidator func(string) error
	// CurrencyValidator is a validator for the "currency" field. It is called by the builders before save.
	CurrencyValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
	// ValueScanner of all BillingInvoice fields.
	ValueScanner struct {
		ProviderConfig    field.TypeValueScanner[provider.Configuration]
		ProviderReference field.TypeValueScanner[provider.Reference]
	}
)

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s invoice.InvoiceStatus) error {
	switch s {
	case "created", "draft", "draft_sync", "draft_sync_failed", "issuing", "issued", "issuing_failed", "manual_approval_needed":
		return nil
	default:
		return fmt.Errorf("billinginvoice: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the BillingInvoice queries.
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

// ByKey orders the results by the key field.
func ByKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKey, opts...).ToFunc()
}

// ByCustomerID orders the results by the customer_id field.
func ByCustomerID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCustomerID, opts...).ToFunc()
}

// ByBillingProfileID orders the results by the billing_profile_id field.
func ByBillingProfileID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBillingProfileID, opts...).ToFunc()
}

// ByVoidedAt orders the results by the voided_at field.
func ByVoidedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVoidedAt, opts...).ToFunc()
}

// ByCurrency orders the results by the currency field.
func ByCurrency(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCurrency, opts...).ToFunc()
}

// ByTotalAmount orders the results by the total_amount field.
func ByTotalAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalAmount, opts...).ToFunc()
}

// ByDueDate orders the results by the due_date field.
func ByDueDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDueDate, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByProviderConfig orders the results by the provider_config field.
func ByProviderConfig(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProviderConfig, opts...).ToFunc()
}

// ByProviderReference orders the results by the provider_reference field.
func ByProviderReference(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProviderReference, opts...).ToFunc()
}

// ByPeriodStart orders the results by the period_start field.
func ByPeriodStart(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPeriodStart, opts...).ToFunc()
}

// ByPeriodEnd orders the results by the period_end field.
func ByPeriodEnd(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPeriodEnd, opts...).ToFunc()
}

// ByBillingProfileField orders the results by billing_profile field.
func ByBillingProfileField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBillingProfileStep(), sql.OrderByField(field, opts...))
	}
}

// ByBillingInvoiceItemsCount orders the results by billing_invoice_items count.
func ByBillingInvoiceItemsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newBillingInvoiceItemsStep(), opts...)
	}
}

// ByBillingInvoiceItems orders the results by billing_invoice_items terms.
func ByBillingInvoiceItems(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBillingInvoiceItemsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newBillingProfileStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BillingProfileInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, BillingProfileTable, BillingProfileColumn),
	)
}
func newBillingInvoiceItemsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BillingInvoiceItemsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, BillingInvoiceItemsTable, BillingInvoiceItemsColumn),
	)
}
