// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/SeyramWood/ent/customer"
)

// Customer is the model entity for the Customer schema.
type Customer struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Password holds the value of the "password" field.
	Password []byte `json:"-"`
	// FirstName holds the value of the "first_name" field.
	FirstName string `json:"first_name,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// OtherPhone holds the value of the "other_phone" field.
	OtherPhone *string `json:"other_phone,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Customer) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case customer.FieldPassword:
			values[i] = new([]byte)
		case customer.FieldID:
			values[i] = new(sql.NullInt64)
		case customer.FieldUsername, customer.FieldFirstName, customer.FieldLastName, customer.FieldPhone, customer.FieldOtherPhone:
			values[i] = new(sql.NullString)
		case customer.FieldCreatedAt, customer.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Customer", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Customer fields.
func (c *Customer) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case customer.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case customer.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case customer.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case customer.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				c.Username = value.String
			}
		case customer.FieldPassword:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value != nil {
				c.Password = *value
			}
		case customer.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				c.FirstName = value.String
			}
		case customer.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				c.LastName = value.String
			}
		case customer.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				c.Phone = value.String
			}
		case customer.FieldOtherPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field other_phone", values[i])
			} else if value.Valid {
				c.OtherPhone = new(string)
				*c.OtherPhone = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Customer.
// Note that you need to call Customer.Unwrap() before calling this method if this Customer
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Customer) Update() *CustomerUpdateOne {
	return (&CustomerClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Customer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Customer) Unwrap() *Customer {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Customer is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Customer) String() string {
	var builder strings.Builder
	builder.WriteString("Customer(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", username=")
	builder.WriteString(c.Username)
	builder.WriteString(", password=<sensitive>")
	builder.WriteString(", first_name=")
	builder.WriteString(c.FirstName)
	builder.WriteString(", last_name=")
	builder.WriteString(c.LastName)
	builder.WriteString(", phone=")
	builder.WriteString(c.Phone)
	if v := c.OtherPhone; v != nil {
		builder.WriteString(", other_phone=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Customers is a parsable slice of Customer.
type Customers []*Customer

func (c Customers) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}