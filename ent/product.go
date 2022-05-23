// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/SeyramWood/ent/product"
)

// Product is the model entity for the Product schema.
type Product struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "Name" field.
	Name string `json:"Name,omitempty"`
	// Price holds the value of the "Price" field.
	Price float64 `json:"Price,omitempty"`
	// PromoPrice holds the value of the "PromoPrice" field.
	PromoPrice float64 `json:"PromoPrice,omitempty"`
	// Description holds the value of the "Description" field.
	Description string `json:"Description,omitempty"`
	// Image holds the value of the "Image" field.
	Image string `json:"Image,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProductQuery when eager-loading is set.
	Edges ProductEdges `json:"edges"`
}

// ProductEdges holds the relations/edges for other nodes in the graph.
type ProductEdges struct {
	// Major holds the value of the major edge.
	Major []*ProductCategoryMajor `json:"major,omitempty"`
	// Minor holds the value of the minor edge.
	Minor []*ProductCategoryMinor `json:"minor,omitempty"`
	// Mechant holds the value of the mechant edge.
	Mechant []*Merchant `json:"mechant,omitempty"`
	// Supplier holds the value of the supplier edge.
	Supplier []*SupplierMerchant `json:"supplier,omitempty"`
	// Retailer holds the value of the retailer edge.
	Retailer []*RetailMerchant `json:"retailer,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// MajorOrErr returns the Major value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) MajorOrErr() ([]*ProductCategoryMajor, error) {
	if e.loadedTypes[0] {
		return e.Major, nil
	}
	return nil, &NotLoadedError{edge: "major"}
}

// MinorOrErr returns the Minor value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) MinorOrErr() ([]*ProductCategoryMinor, error) {
	if e.loadedTypes[1] {
		return e.Minor, nil
	}
	return nil, &NotLoadedError{edge: "minor"}
}

// MechantOrErr returns the Mechant value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) MechantOrErr() ([]*Merchant, error) {
	if e.loadedTypes[2] {
		return e.Mechant, nil
	}
	return nil, &NotLoadedError{edge: "mechant"}
}

// SupplierOrErr returns the Supplier value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) SupplierOrErr() ([]*SupplierMerchant, error) {
	if e.loadedTypes[3] {
		return e.Supplier, nil
	}
	return nil, &NotLoadedError{edge: "supplier"}
}

// RetailerOrErr returns the Retailer value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) RetailerOrErr() ([]*RetailMerchant, error) {
	if e.loadedTypes[4] {
		return e.Retailer, nil
	}
	return nil, &NotLoadedError{edge: "retailer"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Product) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case product.FieldPrice, product.FieldPromoPrice:
			values[i] = new(sql.NullFloat64)
		case product.FieldID:
			values[i] = new(sql.NullInt64)
		case product.FieldName, product.FieldDescription, product.FieldImage:
			values[i] = new(sql.NullString)
		case product.FieldCreatedAt, product.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Product", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Product fields.
func (pr *Product) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case product.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int(value.Int64)
		case product.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pr.CreatedAt = value.Time
			}
		case product.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pr.UpdatedAt = value.Time
			}
		case product.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case product.FieldPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field Price", values[i])
			} else if value.Valid {
				pr.Price = value.Float64
			}
		case product.FieldPromoPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field PromoPrice", values[i])
			} else if value.Valid {
				pr.PromoPrice = value.Float64
			}
		case product.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Description", values[i])
			} else if value.Valid {
				pr.Description = value.String
			}
		case product.FieldImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Image", values[i])
			} else if value.Valid {
				pr.Image = value.String
			}
		}
	}
	return nil
}

// QueryMajor queries the "major" edge of the Product entity.
func (pr *Product) QueryMajor() *ProductCategoryMajorQuery {
	return (&ProductClient{config: pr.config}).QueryMajor(pr)
}

// QueryMinor queries the "minor" edge of the Product entity.
func (pr *Product) QueryMinor() *ProductCategoryMinorQuery {
	return (&ProductClient{config: pr.config}).QueryMinor(pr)
}

// QueryMechant queries the "mechant" edge of the Product entity.
func (pr *Product) QueryMechant() *MerchantQuery {
	return (&ProductClient{config: pr.config}).QueryMechant(pr)
}

// QuerySupplier queries the "supplier" edge of the Product entity.
func (pr *Product) QuerySupplier() *SupplierMerchantQuery {
	return (&ProductClient{config: pr.config}).QuerySupplier(pr)
}

// QueryRetailer queries the "retailer" edge of the Product entity.
func (pr *Product) QueryRetailer() *RetailMerchantQuery {
	return (&ProductClient{config: pr.config}).QueryRetailer(pr)
}

// Update returns a builder for updating this Product.
// Note that you need to call Product.Unwrap() before calling this method if this Product
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Product) Update() *ProductUpdateOne {
	return (&ProductClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the Product entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Product) Unwrap() *Product {
	tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Product is not a transactional entity")
	}
	pr.config.driver = tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Product) String() string {
	var builder strings.Builder
	builder.WriteString("Product(")
	builder.WriteString(fmt.Sprintf("id=%v", pr.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(pr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(pr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", Name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", Price=")
	builder.WriteString(fmt.Sprintf("%v", pr.Price))
	builder.WriteString(", PromoPrice=")
	builder.WriteString(fmt.Sprintf("%v", pr.PromoPrice))
	builder.WriteString(", Description=")
	builder.WriteString(pr.Description)
	builder.WriteString(", Image=")
	builder.WriteString(pr.Image)
	builder.WriteByte(')')
	return builder.String()
}

// Products is a parsable slice of Product.
type Products []*Product

func (pr Products) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}