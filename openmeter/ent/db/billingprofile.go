// Code generated by ent, DO NOT EDIT.

package db

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	dbapp "github.com/openmeterio/openmeter/openmeter/ent/db/app"
	"github.com/openmeterio/openmeter/openmeter/ent/db/billingprofile"
	"github.com/openmeterio/openmeter/openmeter/ent/db/billingworkflowconfig"
	"github.com/openmeterio/openmeter/pkg/models"
)

// BillingProfile is the model entity for the BillingProfile schema.
type BillingProfile struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Namespace holds the value of the "namespace" field.
	Namespace string `json:"namespace,omitempty"`
	// Metadata holds the value of the "metadata" field.
	Metadata map[string]string `json:"metadata,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description *string `json:"description,omitempty"`
	// SupplierAddressCountry holds the value of the "supplier_address_country" field.
	SupplierAddressCountry *models.CountryCode `json:"supplier_address_country,omitempty"`
	// SupplierAddressPostalCode holds the value of the "supplier_address_postal_code" field.
	SupplierAddressPostalCode *string `json:"supplier_address_postal_code,omitempty"`
	// SupplierAddressState holds the value of the "supplier_address_state" field.
	SupplierAddressState *string `json:"supplier_address_state,omitempty"`
	// SupplierAddressCity holds the value of the "supplier_address_city" field.
	SupplierAddressCity *string `json:"supplier_address_city,omitempty"`
	// SupplierAddressLine1 holds the value of the "supplier_address_line1" field.
	SupplierAddressLine1 *string `json:"supplier_address_line1,omitempty"`
	// SupplierAddressLine2 holds the value of the "supplier_address_line2" field.
	SupplierAddressLine2 *string `json:"supplier_address_line2,omitempty"`
	// SupplierAddressPhoneNumber holds the value of the "supplier_address_phone_number" field.
	SupplierAddressPhoneNumber *string `json:"supplier_address_phone_number,omitempty"`
	// TaxAppID holds the value of the "tax_app_id" field.
	TaxAppID string `json:"tax_app_id,omitempty"`
	// InvoicingAppID holds the value of the "invoicing_app_id" field.
	InvoicingAppID string `json:"invoicing_app_id,omitempty"`
	// PaymentAppID holds the value of the "payment_app_id" field.
	PaymentAppID string `json:"payment_app_id,omitempty"`
	// WorkflowConfigID holds the value of the "workflow_config_id" field.
	WorkflowConfigID string `json:"workflow_config_id,omitempty"`
	// Default holds the value of the "default" field.
	Default bool `json:"default,omitempty"`
	// SupplierName holds the value of the "supplier_name" field.
	SupplierName string `json:"supplier_name,omitempty"`
	// SupplierTaxCode holds the value of the "supplier_tax_code" field.
	SupplierTaxCode *string `json:"supplier_tax_code,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BillingProfileQuery when eager-loading is set.
	Edges        BillingProfileEdges `json:"edges"`
	selectValues sql.SelectValues
}

// BillingProfileEdges holds the relations/edges for other nodes in the graph.
type BillingProfileEdges struct {
	// BillingInvoices holds the value of the billing_invoices edge.
	BillingInvoices []*BillingInvoice `json:"billing_invoices,omitempty"`
	// BillingCustomerOverride holds the value of the billing_customer_override edge.
	BillingCustomerOverride []*BillingCustomerOverride `json:"billing_customer_override,omitempty"`
	// WorkflowConfig holds the value of the workflow_config edge.
	WorkflowConfig *BillingWorkflowConfig `json:"workflow_config,omitempty"`
	// TaxApp holds the value of the tax_app edge.
	TaxApp *App `json:"tax_app,omitempty"`
	// InvoicingApp holds the value of the invoicing_app edge.
	InvoicingApp *App `json:"invoicing_app,omitempty"`
	// PaymentApp holds the value of the payment_app edge.
	PaymentApp *App `json:"payment_app,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// BillingInvoicesOrErr returns the BillingInvoices value or an error if the edge
// was not loaded in eager-loading.
func (e BillingProfileEdges) BillingInvoicesOrErr() ([]*BillingInvoice, error) {
	if e.loadedTypes[0] {
		return e.BillingInvoices, nil
	}
	return nil, &NotLoadedError{edge: "billing_invoices"}
}

// BillingCustomerOverrideOrErr returns the BillingCustomerOverride value or an error if the edge
// was not loaded in eager-loading.
func (e BillingProfileEdges) BillingCustomerOverrideOrErr() ([]*BillingCustomerOverride, error) {
	if e.loadedTypes[1] {
		return e.BillingCustomerOverride, nil
	}
	return nil, &NotLoadedError{edge: "billing_customer_override"}
}

// WorkflowConfigOrErr returns the WorkflowConfig value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BillingProfileEdges) WorkflowConfigOrErr() (*BillingWorkflowConfig, error) {
	if e.WorkflowConfig != nil {
		return e.WorkflowConfig, nil
	} else if e.loadedTypes[2] {
		return nil, &NotFoundError{label: billingworkflowconfig.Label}
	}
	return nil, &NotLoadedError{edge: "workflow_config"}
}

// TaxAppOrErr returns the TaxApp value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BillingProfileEdges) TaxAppOrErr() (*App, error) {
	if e.TaxApp != nil {
		return e.TaxApp, nil
	} else if e.loadedTypes[3] {
		return nil, &NotFoundError{label: dbapp.Label}
	}
	return nil, &NotLoadedError{edge: "tax_app"}
}

// InvoicingAppOrErr returns the InvoicingApp value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BillingProfileEdges) InvoicingAppOrErr() (*App, error) {
	if e.InvoicingApp != nil {
		return e.InvoicingApp, nil
	} else if e.loadedTypes[4] {
		return nil, &NotFoundError{label: dbapp.Label}
	}
	return nil, &NotLoadedError{edge: "invoicing_app"}
}

// PaymentAppOrErr returns the PaymentApp value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BillingProfileEdges) PaymentAppOrErr() (*App, error) {
	if e.PaymentApp != nil {
		return e.PaymentApp, nil
	} else if e.loadedTypes[5] {
		return nil, &NotFoundError{label: dbapp.Label}
	}
	return nil, &NotLoadedError{edge: "payment_app"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BillingProfile) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case billingprofile.FieldMetadata:
			values[i] = new([]byte)
		case billingprofile.FieldDefault:
			values[i] = new(sql.NullBool)
		case billingprofile.FieldID, billingprofile.FieldNamespace, billingprofile.FieldName, billingprofile.FieldDescription, billingprofile.FieldSupplierAddressCountry, billingprofile.FieldSupplierAddressPostalCode, billingprofile.FieldSupplierAddressState, billingprofile.FieldSupplierAddressCity, billingprofile.FieldSupplierAddressLine1, billingprofile.FieldSupplierAddressLine2, billingprofile.FieldSupplierAddressPhoneNumber, billingprofile.FieldTaxAppID, billingprofile.FieldInvoicingAppID, billingprofile.FieldPaymentAppID, billingprofile.FieldWorkflowConfigID, billingprofile.FieldSupplierName, billingprofile.FieldSupplierTaxCode:
			values[i] = new(sql.NullString)
		case billingprofile.FieldCreatedAt, billingprofile.FieldUpdatedAt, billingprofile.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BillingProfile fields.
func (bp *BillingProfile) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case billingprofile.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				bp.ID = value.String
			}
		case billingprofile.FieldNamespace:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field namespace", values[i])
			} else if value.Valid {
				bp.Namespace = value.String
			}
		case billingprofile.FieldMetadata:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field metadata", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &bp.Metadata); err != nil {
					return fmt.Errorf("unmarshal field metadata: %w", err)
				}
			}
		case billingprofile.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				bp.CreatedAt = value.Time
			}
		case billingprofile.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				bp.UpdatedAt = value.Time
			}
		case billingprofile.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				bp.DeletedAt = new(time.Time)
				*bp.DeletedAt = value.Time
			}
		case billingprofile.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				bp.Name = value.String
			}
		case billingprofile.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				bp.Description = new(string)
				*bp.Description = value.String
			}
		case billingprofile.FieldSupplierAddressCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field supplier_address_country", values[i])
			} else if value.Valid {
				bp.SupplierAddressCountry = new(models.CountryCode)
				*bp.SupplierAddressCountry = models.CountryCode(value.String)
			}
		case billingprofile.FieldSupplierAddressPostalCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field supplier_address_postal_code", values[i])
			} else if value.Valid {
				bp.SupplierAddressPostalCode = new(string)
				*bp.SupplierAddressPostalCode = value.String
			}
		case billingprofile.FieldSupplierAddressState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field supplier_address_state", values[i])
			} else if value.Valid {
				bp.SupplierAddressState = new(string)
				*bp.SupplierAddressState = value.String
			}
		case billingprofile.FieldSupplierAddressCity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field supplier_address_city", values[i])
			} else if value.Valid {
				bp.SupplierAddressCity = new(string)
				*bp.SupplierAddressCity = value.String
			}
		case billingprofile.FieldSupplierAddressLine1:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field supplier_address_line1", values[i])
			} else if value.Valid {
				bp.SupplierAddressLine1 = new(string)
				*bp.SupplierAddressLine1 = value.String
			}
		case billingprofile.FieldSupplierAddressLine2:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field supplier_address_line2", values[i])
			} else if value.Valid {
				bp.SupplierAddressLine2 = new(string)
				*bp.SupplierAddressLine2 = value.String
			}
		case billingprofile.FieldSupplierAddressPhoneNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field supplier_address_phone_number", values[i])
			} else if value.Valid {
				bp.SupplierAddressPhoneNumber = new(string)
				*bp.SupplierAddressPhoneNumber = value.String
			}
		case billingprofile.FieldTaxAppID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tax_app_id", values[i])
			} else if value.Valid {
				bp.TaxAppID = value.String
			}
		case billingprofile.FieldInvoicingAppID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field invoicing_app_id", values[i])
			} else if value.Valid {
				bp.InvoicingAppID = value.String
			}
		case billingprofile.FieldPaymentAppID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field payment_app_id", values[i])
			} else if value.Valid {
				bp.PaymentAppID = value.String
			}
		case billingprofile.FieldWorkflowConfigID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field workflow_config_id", values[i])
			} else if value.Valid {
				bp.WorkflowConfigID = value.String
			}
		case billingprofile.FieldDefault:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field default", values[i])
			} else if value.Valid {
				bp.Default = value.Bool
			}
		case billingprofile.FieldSupplierName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field supplier_name", values[i])
			} else if value.Valid {
				bp.SupplierName = value.String
			}
		case billingprofile.FieldSupplierTaxCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field supplier_tax_code", values[i])
			} else if value.Valid {
				bp.SupplierTaxCode = new(string)
				*bp.SupplierTaxCode = value.String
			}
		default:
			bp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the BillingProfile.
// This includes values selected through modifiers, order, etc.
func (bp *BillingProfile) Value(name string) (ent.Value, error) {
	return bp.selectValues.Get(name)
}

// QueryBillingInvoices queries the "billing_invoices" edge of the BillingProfile entity.
func (bp *BillingProfile) QueryBillingInvoices() *BillingInvoiceQuery {
	return NewBillingProfileClient(bp.config).QueryBillingInvoices(bp)
}

// QueryBillingCustomerOverride queries the "billing_customer_override" edge of the BillingProfile entity.
func (bp *BillingProfile) QueryBillingCustomerOverride() *BillingCustomerOverrideQuery {
	return NewBillingProfileClient(bp.config).QueryBillingCustomerOverride(bp)
}

// QueryWorkflowConfig queries the "workflow_config" edge of the BillingProfile entity.
func (bp *BillingProfile) QueryWorkflowConfig() *BillingWorkflowConfigQuery {
	return NewBillingProfileClient(bp.config).QueryWorkflowConfig(bp)
}

// QueryTaxApp queries the "tax_app" edge of the BillingProfile entity.
func (bp *BillingProfile) QueryTaxApp() *AppQuery {
	return NewBillingProfileClient(bp.config).QueryTaxApp(bp)
}

// QueryInvoicingApp queries the "invoicing_app" edge of the BillingProfile entity.
func (bp *BillingProfile) QueryInvoicingApp() *AppQuery {
	return NewBillingProfileClient(bp.config).QueryInvoicingApp(bp)
}

// QueryPaymentApp queries the "payment_app" edge of the BillingProfile entity.
func (bp *BillingProfile) QueryPaymentApp() *AppQuery {
	return NewBillingProfileClient(bp.config).QueryPaymentApp(bp)
}

// Update returns a builder for updating this BillingProfile.
// Note that you need to call BillingProfile.Unwrap() before calling this method if this BillingProfile
// was returned from a transaction, and the transaction was committed or rolled back.
func (bp *BillingProfile) Update() *BillingProfileUpdateOne {
	return NewBillingProfileClient(bp.config).UpdateOne(bp)
}

// Unwrap unwraps the BillingProfile entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bp *BillingProfile) Unwrap() *BillingProfile {
	_tx, ok := bp.config.driver.(*txDriver)
	if !ok {
		panic("db: BillingProfile is not a transactional entity")
	}
	bp.config.driver = _tx.drv
	return bp
}

// String implements the fmt.Stringer.
func (bp *BillingProfile) String() string {
	var builder strings.Builder
	builder.WriteString("BillingProfile(")
	builder.WriteString(fmt.Sprintf("id=%v, ", bp.ID))
	builder.WriteString("namespace=")
	builder.WriteString(bp.Namespace)
	builder.WriteString(", ")
	builder.WriteString("metadata=")
	builder.WriteString(fmt.Sprintf("%v", bp.Metadata))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(bp.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(bp.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := bp.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(bp.Name)
	builder.WriteString(", ")
	if v := bp.Description; v != nil {
		builder.WriteString("description=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := bp.SupplierAddressCountry; v != nil {
		builder.WriteString("supplier_address_country=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := bp.SupplierAddressPostalCode; v != nil {
		builder.WriteString("supplier_address_postal_code=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := bp.SupplierAddressState; v != nil {
		builder.WriteString("supplier_address_state=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := bp.SupplierAddressCity; v != nil {
		builder.WriteString("supplier_address_city=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := bp.SupplierAddressLine1; v != nil {
		builder.WriteString("supplier_address_line1=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := bp.SupplierAddressLine2; v != nil {
		builder.WriteString("supplier_address_line2=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := bp.SupplierAddressPhoneNumber; v != nil {
		builder.WriteString("supplier_address_phone_number=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("tax_app_id=")
	builder.WriteString(bp.TaxAppID)
	builder.WriteString(", ")
	builder.WriteString("invoicing_app_id=")
	builder.WriteString(bp.InvoicingAppID)
	builder.WriteString(", ")
	builder.WriteString("payment_app_id=")
	builder.WriteString(bp.PaymentAppID)
	builder.WriteString(", ")
	builder.WriteString("workflow_config_id=")
	builder.WriteString(bp.WorkflowConfigID)
	builder.WriteString(", ")
	builder.WriteString("default=")
	builder.WriteString(fmt.Sprintf("%v", bp.Default))
	builder.WriteString(", ")
	builder.WriteString("supplier_name=")
	builder.WriteString(bp.SupplierName)
	builder.WriteString(", ")
	if v := bp.SupplierTaxCode; v != nil {
		builder.WriteString("supplier_tax_code=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// BillingProfiles is a parsable slice of BillingProfile.
type BillingProfiles []*BillingProfile
