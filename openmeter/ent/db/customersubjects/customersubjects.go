// Code generated by ent, DO NOT EDIT.

package customersubjects

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the customersubjects type in the database.
	Label = "customer_subjects"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNamespace holds the string denoting the namespace field in the database.
	FieldNamespace = "namespace"
	// FieldCustomerID holds the string denoting the customer_id field in the database.
	FieldCustomerID = "customer_id"
	// FieldSubjectKey holds the string denoting the subject_key field in the database.
	FieldSubjectKey = "subject_key"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// EdgeCustomer holds the string denoting the customer edge name in mutations.
	EdgeCustomer = "customer"
	// Table holds the table name of the customersubjects in the database.
	Table = "customer_subjects"
	// CustomerTable is the table that holds the customer relation/edge.
	CustomerTable = "customer_subjects"
	// CustomerInverseTable is the table name for the Customer entity.
	// It exists in this package in order to avoid circular dependency with the "customer" package.
	CustomerInverseTable = "customers"
	// CustomerColumn is the table column denoting the customer relation/edge.
	CustomerColumn = "customer_id"
)

// Columns holds all SQL columns for customersubjects fields.
var Columns = []string{
	FieldID,
	FieldNamespace,
	FieldCustomerID,
	FieldSubjectKey,
	FieldCreatedAt,
	FieldDeletedAt,
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
	// CustomerIDValidator is a validator for the "customer_id" field. It is called by the builders before save.
	CustomerIDValidator func(string) error
	// SubjectKeyValidator is a validator for the "subject_key" field. It is called by the builders before save.
	SubjectKeyValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// OrderOption defines the ordering options for the CustomerSubjects queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByNamespace orders the results by the namespace field.
func ByNamespace(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNamespace, opts...).ToFunc()
}

// ByCustomerID orders the results by the customer_id field.
func ByCustomerID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCustomerID, opts...).ToFunc()
}

// BySubjectKey orders the results by the subject_key field.
func BySubjectKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubjectKey, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByCustomerField orders the results by customer field.
func ByCustomerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCustomerStep(), sql.OrderByField(field, opts...))
	}
}
func newCustomerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CustomerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CustomerTable, CustomerColumn),
	)
}
