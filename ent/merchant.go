// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/SeyramWood/ent/merchant"
	"github.com/SeyramWood/ent/merchantstore"
	"github.com/SeyramWood/ent/retailmerchant"
	"github.com/SeyramWood/ent/suppliermerchant"
)

// Merchant is the model entity for the Merchant schema.
type Merchant struct {
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
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MerchantQuery when eager-loading is set.
	Edges MerchantEdges `json:"edges"`
}

// MerchantEdges holds the relations/edges for other nodes in the graph.
type MerchantEdges struct {
	// Supplier holds the value of the supplier edge.
	Supplier *SupplierMerchant `json:"supplier,omitempty"`
	// Retailer holds the value of the retailer edge.
	Retailer *RetailMerchant `json:"retailer,omitempty"`
	// Store holds the value of the store edge.
	Store *MerchantStore `json:"store,omitempty"`
	// Products holds the value of the products edge.
	Products []*Product `json:"products,omitempty"`
	// Addresses holds the value of the addresses edge.
	Addresses []*Address `json:"addresses,omitempty"`
	// Orders holds the value of the orders edge.
	Orders []*Order `json:"orders,omitempty"`
	// Baskets holds the value of the baskets edge.
	Baskets []*Basket `json:"baskets,omitempty"`
	// Favourites holds the value of the favourites edge.
	Favourites []*Favourite `json:"favourites,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [8]bool
}

// SupplierOrErr returns the Supplier value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MerchantEdges) SupplierOrErr() (*SupplierMerchant, error) {
	if e.loadedTypes[0] {
		if e.Supplier == nil {
			// The edge supplier was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: suppliermerchant.Label}
		}
		return e.Supplier, nil
	}
	return nil, &NotLoadedError{edge: "supplier"}
}

// RetailerOrErr returns the Retailer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MerchantEdges) RetailerOrErr() (*RetailMerchant, error) {
	if e.loadedTypes[1] {
		if e.Retailer == nil {
			// The edge retailer was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: retailmerchant.Label}
		}
		return e.Retailer, nil
	}
	return nil, &NotLoadedError{edge: "retailer"}
}

// StoreOrErr returns the Store value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MerchantEdges) StoreOrErr() (*MerchantStore, error) {
	if e.loadedTypes[2] {
		if e.Store == nil {
			// The edge store was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: merchantstore.Label}
		}
		return e.Store, nil
	}
	return nil, &NotLoadedError{edge: "store"}
}

// ProductsOrErr returns the Products value or an error if the edge
// was not loaded in eager-loading.
func (e MerchantEdges) ProductsOrErr() ([]*Product, error) {
	if e.loadedTypes[3] {
		return e.Products, nil
	}
	return nil, &NotLoadedError{edge: "products"}
}

// AddressesOrErr returns the Addresses value or an error if the edge
// was not loaded in eager-loading.
func (e MerchantEdges) AddressesOrErr() ([]*Address, error) {
	if e.loadedTypes[4] {
		return e.Addresses, nil
	}
	return nil, &NotLoadedError{edge: "addresses"}
}

// OrdersOrErr returns the Orders value or an error if the edge
// was not loaded in eager-loading.
func (e MerchantEdges) OrdersOrErr() ([]*Order, error) {
	if e.loadedTypes[5] {
		return e.Orders, nil
	}
	return nil, &NotLoadedError{edge: "orders"}
}

// BasketsOrErr returns the Baskets value or an error if the edge
// was not loaded in eager-loading.
func (e MerchantEdges) BasketsOrErr() ([]*Basket, error) {
	if e.loadedTypes[6] {
		return e.Baskets, nil
	}
	return nil, &NotLoadedError{edge: "baskets"}
}

// FavouritesOrErr returns the Favourites value or an error if the edge
// was not loaded in eager-loading.
func (e MerchantEdges) FavouritesOrErr() ([]*Favourite, error) {
	if e.loadedTypes[7] {
		return e.Favourites, nil
	}
	return nil, &NotLoadedError{edge: "favourites"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Merchant) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case merchant.FieldPassword:
			values[i] = new([]byte)
		case merchant.FieldID:
			values[i] = new(sql.NullInt64)
		case merchant.FieldUsername, merchant.FieldType:
			values[i] = new(sql.NullString)
		case merchant.FieldCreatedAt, merchant.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Merchant", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Merchant fields.
func (m *Merchant) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case merchant.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case merchant.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				m.CreatedAt = value.Time
			}
		case merchant.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				m.UpdatedAt = value.Time
			}
		case merchant.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				m.Username = value.String
			}
		case merchant.FieldPassword:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value != nil {
				m.Password = *value
			}
		case merchant.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				m.Type = value.String
			}
		}
	}
	return nil
}

// QuerySupplier queries the "supplier" edge of the Merchant entity.
func (m *Merchant) QuerySupplier() *SupplierMerchantQuery {
	return (&MerchantClient{config: m.config}).QuerySupplier(m)
}

// QueryRetailer queries the "retailer" edge of the Merchant entity.
func (m *Merchant) QueryRetailer() *RetailMerchantQuery {
	return (&MerchantClient{config: m.config}).QueryRetailer(m)
}

// QueryStore queries the "store" edge of the Merchant entity.
func (m *Merchant) QueryStore() *MerchantStoreQuery {
	return (&MerchantClient{config: m.config}).QueryStore(m)
}

// QueryProducts queries the "products" edge of the Merchant entity.
func (m *Merchant) QueryProducts() *ProductQuery {
	return (&MerchantClient{config: m.config}).QueryProducts(m)
}

// QueryAddresses queries the "addresses" edge of the Merchant entity.
func (m *Merchant) QueryAddresses() *AddressQuery {
	return (&MerchantClient{config: m.config}).QueryAddresses(m)
}

// QueryOrders queries the "orders" edge of the Merchant entity.
func (m *Merchant) QueryOrders() *OrderQuery {
	return (&MerchantClient{config: m.config}).QueryOrders(m)
}

// QueryBaskets queries the "baskets" edge of the Merchant entity.
func (m *Merchant) QueryBaskets() *BasketQuery {
	return (&MerchantClient{config: m.config}).QueryBaskets(m)
}

// QueryFavourites queries the "favourites" edge of the Merchant entity.
func (m *Merchant) QueryFavourites() *FavouriteQuery {
	return (&MerchantClient{config: m.config}).QueryFavourites(m)
}

// Update returns a builder for updating this Merchant.
// Note that you need to call Merchant.Unwrap() before calling this method if this Merchant
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Merchant) Update() *MerchantUpdateOne {
	return (&MerchantClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the Merchant entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Merchant) Unwrap() *Merchant {
	tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Merchant is not a transactional entity")
	}
	m.config.driver = tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Merchant) String() string {
	var builder strings.Builder
	builder.WriteString("Merchant(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(m.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(m.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", username=")
	builder.WriteString(m.Username)
	builder.WriteString(", password=<sensitive>")
	builder.WriteString(", type=")
	builder.WriteString(m.Type)
	builder.WriteByte(')')
	return builder.String()
}

// Merchants is a parsable slice of Merchant.
type Merchants []*Merchant

func (m Merchants) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
