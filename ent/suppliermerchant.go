// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/SeyramWood/ent/merchant"
	"github.com/SeyramWood/ent/suppliermerchant"
)

// SupplierMerchant is the model entity for the SupplierMerchant schema.
type SupplierMerchant struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// GhanaCard holds the value of the "ghana_card" field.
	GhanaCard string `json:"ghana_card,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// OtherName holds the value of the "other_name" field.
	OtherName string `json:"other_name,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// OtherPhone holds the value of the "other_phone" field.
	OtherPhone *string `json:"other_phone,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SupplierMerchantQuery when eager-loading is set.
	Edges             SupplierMerchantEdges `json:"edges"`
	merchant_supplier *int
	selectValues      sql.SelectValues
}

// SupplierMerchantEdges holds the relations/edges for other nodes in the graph.
type SupplierMerchantEdges struct {
	// Merchant holds the value of the merchant edge.
	Merchant *Merchant `json:"merchant,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// MerchantOrErr returns the Merchant value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SupplierMerchantEdges) MerchantOrErr() (*Merchant, error) {
	if e.loadedTypes[0] {
		if e.Merchant == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: merchant.Label}
		}
		return e.Merchant, nil
	}
	return nil, &NotLoadedError{edge: "merchant"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SupplierMerchant) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case suppliermerchant.FieldID:
			values[i] = new(sql.NullInt64)
		case suppliermerchant.FieldGhanaCard, suppliermerchant.FieldLastName, suppliermerchant.FieldOtherName, suppliermerchant.FieldPhone, suppliermerchant.FieldOtherPhone:
			values[i] = new(sql.NullString)
		case suppliermerchant.FieldCreatedAt, suppliermerchant.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case suppliermerchant.ForeignKeys[0]: // merchant_supplier
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SupplierMerchant fields.
func (sm *SupplierMerchant) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case suppliermerchant.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sm.ID = int(value.Int64)
		case suppliermerchant.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sm.CreatedAt = value.Time
			}
		case suppliermerchant.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sm.UpdatedAt = value.Time
			}
		case suppliermerchant.FieldGhanaCard:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ghana_card", values[i])
			} else if value.Valid {
				sm.GhanaCard = value.String
			}
		case suppliermerchant.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				sm.LastName = value.String
			}
		case suppliermerchant.FieldOtherName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field other_name", values[i])
			} else if value.Valid {
				sm.OtherName = value.String
			}
		case suppliermerchant.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				sm.Phone = value.String
			}
		case suppliermerchant.FieldOtherPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field other_phone", values[i])
			} else if value.Valid {
				sm.OtherPhone = new(string)
				*sm.OtherPhone = value.String
			}
		case suppliermerchant.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field merchant_supplier", value)
			} else if value.Valid {
				sm.merchant_supplier = new(int)
				*sm.merchant_supplier = int(value.Int64)
			}
		default:
			sm.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SupplierMerchant.
// This includes values selected through modifiers, order, etc.
func (sm *SupplierMerchant) Value(name string) (ent.Value, error) {
	return sm.selectValues.Get(name)
}

// QueryMerchant queries the "merchant" edge of the SupplierMerchant entity.
func (sm *SupplierMerchant) QueryMerchant() *MerchantQuery {
	return NewSupplierMerchantClient(sm.config).QueryMerchant(sm)
}

// Update returns a builder for updating this SupplierMerchant.
// Note that you need to call SupplierMerchant.Unwrap() before calling this method if this SupplierMerchant
// was returned from a transaction, and the transaction was committed or rolled back.
func (sm *SupplierMerchant) Update() *SupplierMerchantUpdateOne {
	return NewSupplierMerchantClient(sm.config).UpdateOne(sm)
}

// Unwrap unwraps the SupplierMerchant entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sm *SupplierMerchant) Unwrap() *SupplierMerchant {
	_tx, ok := sm.config.driver.(*txDriver)
	if !ok {
		panic("ent: SupplierMerchant is not a transactional entity")
	}
	sm.config.driver = _tx.drv
	return sm
}

// String implements the fmt.Stringer.
func (sm *SupplierMerchant) String() string {
	var builder strings.Builder
	builder.WriteString("SupplierMerchant(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sm.ID))
	builder.WriteString("created_at=")
	builder.WriteString(sm.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(sm.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("ghana_card=")
	builder.WriteString(sm.GhanaCard)
	builder.WriteString(", ")
	builder.WriteString("last_name=")
	builder.WriteString(sm.LastName)
	builder.WriteString(", ")
	builder.WriteString("other_name=")
	builder.WriteString(sm.OtherName)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(sm.Phone)
	builder.WriteString(", ")
	if v := sm.OtherPhone; v != nil {
		builder.WriteString("other_phone=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// SupplierMerchants is a parsable slice of SupplierMerchant.
type SupplierMerchants []*SupplierMerchant
