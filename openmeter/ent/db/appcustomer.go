// Code generated by ent, DO NOT EDIT.

package db

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/openmeterio/openmeter/openmeter/ent/db/app"
	"github.com/openmeterio/openmeter/openmeter/ent/db/appcustomer"
	"github.com/openmeterio/openmeter/openmeter/ent/db/customer"
)

// AppCustomer is the model entity for the AppCustomer schema.
type AppCustomer struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Namespace holds the value of the "namespace" field.
	Namespace string `json:"namespace,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID string `json:"app_id,omitempty"`
	// CustomerID holds the value of the "customer_id" field.
	CustomerID string `json:"customer_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AppCustomerQuery when eager-loading is set.
	Edges        AppCustomerEdges `json:"edges"`
	selectValues sql.SelectValues
}

// AppCustomerEdges holds the relations/edges for other nodes in the graph.
type AppCustomerEdges struct {
	// App holds the value of the app edge.
	App *App `json:"app,omitempty"`
	// Customer holds the value of the customer edge.
	Customer *Customer `json:"customer,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// AppOrErr returns the App value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AppCustomerEdges) AppOrErr() (*App, error) {
	if e.App != nil {
		return e.App, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: app.Label}
	}
	return nil, &NotLoadedError{edge: "app"}
}

// CustomerOrErr returns the Customer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AppCustomerEdges) CustomerOrErr() (*Customer, error) {
	if e.Customer != nil {
		return e.Customer, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: customer.Label}
	}
	return nil, &NotLoadedError{edge: "customer"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AppCustomer) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case appcustomer.FieldID:
			values[i] = new(sql.NullInt64)
		case appcustomer.FieldNamespace, appcustomer.FieldAppID, appcustomer.FieldCustomerID:
			values[i] = new(sql.NullString)
		case appcustomer.FieldCreatedAt, appcustomer.FieldUpdatedAt, appcustomer.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AppCustomer fields.
func (ac *AppCustomer) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case appcustomer.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ac.ID = int(value.Int64)
		case appcustomer.FieldNamespace:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field namespace", values[i])
			} else if value.Valid {
				ac.Namespace = value.String
			}
		case appcustomer.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ac.CreatedAt = value.Time
			}
		case appcustomer.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ac.UpdatedAt = value.Time
			}
		case appcustomer.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ac.DeletedAt = new(time.Time)
				*ac.DeletedAt = value.Time
			}
		case appcustomer.FieldAppID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value.Valid {
				ac.AppID = value.String
			}
		case appcustomer.FieldCustomerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field customer_id", values[i])
			} else if value.Valid {
				ac.CustomerID = value.String
			}
		default:
			ac.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the AppCustomer.
// This includes values selected through modifiers, order, etc.
func (ac *AppCustomer) Value(name string) (ent.Value, error) {
	return ac.selectValues.Get(name)
}

// QueryApp queries the "app" edge of the AppCustomer entity.
func (ac *AppCustomer) QueryApp() *AppQuery {
	return NewAppCustomerClient(ac.config).QueryApp(ac)
}

// QueryCustomer queries the "customer" edge of the AppCustomer entity.
func (ac *AppCustomer) QueryCustomer() *CustomerQuery {
	return NewAppCustomerClient(ac.config).QueryCustomer(ac)
}

// Update returns a builder for updating this AppCustomer.
// Note that you need to call AppCustomer.Unwrap() before calling this method if this AppCustomer
// was returned from a transaction, and the transaction was committed or rolled back.
func (ac *AppCustomer) Update() *AppCustomerUpdateOne {
	return NewAppCustomerClient(ac.config).UpdateOne(ac)
}

// Unwrap unwraps the AppCustomer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ac *AppCustomer) Unwrap() *AppCustomer {
	_tx, ok := ac.config.driver.(*txDriver)
	if !ok {
		panic("db: AppCustomer is not a transactional entity")
	}
	ac.config.driver = _tx.drv
	return ac
}

// String implements the fmt.Stringer.
func (ac *AppCustomer) String() string {
	var builder strings.Builder
	builder.WriteString("AppCustomer(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ac.ID))
	builder.WriteString("namespace=")
	builder.WriteString(ac.Namespace)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ac.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ac.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := ac.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(ac.AppID)
	builder.WriteString(", ")
	builder.WriteString("customer_id=")
	builder.WriteString(ac.CustomerID)
	builder.WriteByte(')')
	return builder.String()
}

// AppCustomers is a parsable slice of AppCustomer.
type AppCustomers []*AppCustomer
