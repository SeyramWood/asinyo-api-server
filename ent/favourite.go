// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/SeyramWood/ent/agent"
	"github.com/SeyramWood/ent/customer"
	"github.com/SeyramWood/ent/favourite"
	"github.com/SeyramWood/ent/merchant"
	"github.com/SeyramWood/ent/product"
)

// Favourite is the model entity for the Favourite schema.
type Favourite struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FavouriteQuery when eager-loading is set.
	Edges               FavouriteEdges `json:"edges"`
	agent_favourites    *int
	customer_favourites *int
	merchant_favourites *int
	product_favourites  *int
	selectValues        sql.SelectValues
}

// FavouriteEdges holds the relations/edges for other nodes in the graph.
type FavouriteEdges struct {
	// Merchant holds the value of the merchant edge.
	Merchant *Merchant `json:"merchant,omitempty"`
	// Agent holds the value of the agent edge.
	Agent *Agent `json:"agent,omitempty"`
	// Customer holds the value of the customer edge.
	Customer *Customer `json:"customer,omitempty"`
	// Product holds the value of the product edge.
	Product *Product `json:"product,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// MerchantOrErr returns the Merchant value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FavouriteEdges) MerchantOrErr() (*Merchant, error) {
	if e.loadedTypes[0] {
		if e.Merchant == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: merchant.Label}
		}
		return e.Merchant, nil
	}
	return nil, &NotLoadedError{edge: "merchant"}
}

// AgentOrErr returns the Agent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FavouriteEdges) AgentOrErr() (*Agent, error) {
	if e.loadedTypes[1] {
		if e.Agent == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: agent.Label}
		}
		return e.Agent, nil
	}
	return nil, &NotLoadedError{edge: "agent"}
}

// CustomerOrErr returns the Customer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FavouriteEdges) CustomerOrErr() (*Customer, error) {
	if e.loadedTypes[2] {
		if e.Customer == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: customer.Label}
		}
		return e.Customer, nil
	}
	return nil, &NotLoadedError{edge: "customer"}
}

// ProductOrErr returns the Product value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FavouriteEdges) ProductOrErr() (*Product, error) {
	if e.loadedTypes[3] {
		if e.Product == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: product.Label}
		}
		return e.Product, nil
	}
	return nil, &NotLoadedError{edge: "product"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Favourite) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case favourite.FieldID:
			values[i] = new(sql.NullInt64)
		case favourite.FieldCreatedAt, favourite.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case favourite.ForeignKeys[0]: // agent_favourites
			values[i] = new(sql.NullInt64)
		case favourite.ForeignKeys[1]: // customer_favourites
			values[i] = new(sql.NullInt64)
		case favourite.ForeignKeys[2]: // merchant_favourites
			values[i] = new(sql.NullInt64)
		case favourite.ForeignKeys[3]: // product_favourites
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Favourite fields.
func (f *Favourite) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case favourite.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			f.ID = int(value.Int64)
		case favourite.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				f.CreatedAt = value.Time
			}
		case favourite.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				f.UpdatedAt = value.Time
			}
		case favourite.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field agent_favourites", value)
			} else if value.Valid {
				f.agent_favourites = new(int)
				*f.agent_favourites = int(value.Int64)
			}
		case favourite.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field customer_favourites", value)
			} else if value.Valid {
				f.customer_favourites = new(int)
				*f.customer_favourites = int(value.Int64)
			}
		case favourite.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field merchant_favourites", value)
			} else if value.Valid {
				f.merchant_favourites = new(int)
				*f.merchant_favourites = int(value.Int64)
			}
		case favourite.ForeignKeys[3]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field product_favourites", value)
			} else if value.Valid {
				f.product_favourites = new(int)
				*f.product_favourites = int(value.Int64)
			}
		default:
			f.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Favourite.
// This includes values selected through modifiers, order, etc.
func (f *Favourite) Value(name string) (ent.Value, error) {
	return f.selectValues.Get(name)
}

// QueryMerchant queries the "merchant" edge of the Favourite entity.
func (f *Favourite) QueryMerchant() *MerchantQuery {
	return NewFavouriteClient(f.config).QueryMerchant(f)
}

// QueryAgent queries the "agent" edge of the Favourite entity.
func (f *Favourite) QueryAgent() *AgentQuery {
	return NewFavouriteClient(f.config).QueryAgent(f)
}

// QueryCustomer queries the "customer" edge of the Favourite entity.
func (f *Favourite) QueryCustomer() *CustomerQuery {
	return NewFavouriteClient(f.config).QueryCustomer(f)
}

// QueryProduct queries the "product" edge of the Favourite entity.
func (f *Favourite) QueryProduct() *ProductQuery {
	return NewFavouriteClient(f.config).QueryProduct(f)
}

// Update returns a builder for updating this Favourite.
// Note that you need to call Favourite.Unwrap() before calling this method if this Favourite
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Favourite) Update() *FavouriteUpdateOne {
	return NewFavouriteClient(f.config).UpdateOne(f)
}

// Unwrap unwraps the Favourite entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *Favourite) Unwrap() *Favourite {
	_tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Favourite is not a transactional entity")
	}
	f.config.driver = _tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Favourite) String() string {
	var builder strings.Builder
	builder.WriteString("Favourite(")
	builder.WriteString(fmt.Sprintf("id=%v, ", f.ID))
	builder.WriteString("created_at=")
	builder.WriteString(f.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(f.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Favourites is a parsable slice of Favourite.
type Favourites []*Favourite
