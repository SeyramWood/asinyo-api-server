// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/SeyramWood/ent/customer"
	"github.com/SeyramWood/ent/individualcustomer"
)

// IndividualCustomer is the model entity for the IndividualCustomer schema.
type IndividualCustomer struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// OtherName holds the value of the "other_name" field.
	OtherName string `json:"other_name,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// OtherPhone holds the value of the "other_phone" field.
	OtherPhone string `json:"other_phone,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the IndividualCustomerQuery when eager-loading is set.
	Edges               IndividualCustomerEdges `json:"edges"`
	customer_individual *int
	selectValues        sql.SelectValues
}

// IndividualCustomerEdges holds the relations/edges for other nodes in the graph.
type IndividualCustomerEdges struct {
	// Customer holds the value of the customer edge.
	Customer *Customer `json:"customer,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CustomerOrErr returns the Customer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e IndividualCustomerEdges) CustomerOrErr() (*Customer, error) {
	if e.loadedTypes[0] {
		if e.Customer == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: customer.Label}
		}
		return e.Customer, nil
	}
	return nil, &NotLoadedError{edge: "customer"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*IndividualCustomer) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case individualcustomer.FieldID:
			values[i] = new(sql.NullInt64)
		case individualcustomer.FieldLastName, individualcustomer.FieldOtherName, individualcustomer.FieldPhone, individualcustomer.FieldOtherPhone:
			values[i] = new(sql.NullString)
		case individualcustomer.FieldCreatedAt, individualcustomer.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case individualcustomer.ForeignKeys[0]: // customer_individual
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the IndividualCustomer fields.
func (ic *IndividualCustomer) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case individualcustomer.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ic.ID = int(value.Int64)
		case individualcustomer.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ic.CreatedAt = value.Time
			}
		case individualcustomer.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ic.UpdatedAt = value.Time
			}
		case individualcustomer.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				ic.LastName = value.String
			}
		case individualcustomer.FieldOtherName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field other_name", values[i])
			} else if value.Valid {
				ic.OtherName = value.String
			}
		case individualcustomer.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				ic.Phone = value.String
			}
		case individualcustomer.FieldOtherPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field other_phone", values[i])
			} else if value.Valid {
				ic.OtherPhone = value.String
			}
		case individualcustomer.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field customer_individual", value)
			} else if value.Valid {
				ic.customer_individual = new(int)
				*ic.customer_individual = int(value.Int64)
			}
		default:
			ic.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the IndividualCustomer.
// This includes values selected through modifiers, order, etc.
func (ic *IndividualCustomer) Value(name string) (ent.Value, error) {
	return ic.selectValues.Get(name)
}

// QueryCustomer queries the "customer" edge of the IndividualCustomer entity.
func (ic *IndividualCustomer) QueryCustomer() *CustomerQuery {
	return NewIndividualCustomerClient(ic.config).QueryCustomer(ic)
}

// Update returns a builder for updating this IndividualCustomer.
// Note that you need to call IndividualCustomer.Unwrap() before calling this method if this IndividualCustomer
// was returned from a transaction, and the transaction was committed or rolled back.
func (ic *IndividualCustomer) Update() *IndividualCustomerUpdateOne {
	return NewIndividualCustomerClient(ic.config).UpdateOne(ic)
}

// Unwrap unwraps the IndividualCustomer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ic *IndividualCustomer) Unwrap() *IndividualCustomer {
	_tx, ok := ic.config.driver.(*txDriver)
	if !ok {
		panic("ent: IndividualCustomer is not a transactional entity")
	}
	ic.config.driver = _tx.drv
	return ic
}

// String implements the fmt.Stringer.
func (ic *IndividualCustomer) String() string {
	var builder strings.Builder
	builder.WriteString("IndividualCustomer(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ic.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ic.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ic.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("last_name=")
	builder.WriteString(ic.LastName)
	builder.WriteString(", ")
	builder.WriteString("other_name=")
	builder.WriteString(ic.OtherName)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(ic.Phone)
	builder.WriteString(", ")
	builder.WriteString("other_phone=")
	builder.WriteString(ic.OtherPhone)
	builder.WriteByte(')')
	return builder.String()
}

// IndividualCustomers is a parsable slice of IndividualCustomer.
type IndividualCustomers []*IndividualCustomer
