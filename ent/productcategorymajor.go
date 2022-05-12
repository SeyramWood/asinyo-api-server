// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/SeyramWood/ent/productcategorymajor"
)

// ProductCategoryMajor is the model entity for the ProductCategoryMajor schema.
type ProductCategoryMajor struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Category holds the value of the "category" field.
	Category string `json:"category,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProductCategoryMajorQuery when eager-loading is set.
	Edges ProductCategoryMajorEdges `json:"edges"`
}

// ProductCategoryMajorEdges holds the relations/edges for other nodes in the graph.
type ProductCategoryMajorEdges struct {
	// Minors holds the value of the minors edge.
	Minors []*ProductCategoryMinor `json:"minors,omitempty"`
	// Products holds the value of the products edge.
	Products []*Product `json:"products,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// MinorsOrErr returns the Minors value or an error if the edge
// was not loaded in eager-loading.
func (e ProductCategoryMajorEdges) MinorsOrErr() ([]*ProductCategoryMinor, error) {
	if e.loadedTypes[0] {
		return e.Minors, nil
	}
	return nil, &NotLoadedError{edge: "minors"}
}

// ProductsOrErr returns the Products value or an error if the edge
// was not loaded in eager-loading.
func (e ProductCategoryMajorEdges) ProductsOrErr() ([]*Product, error) {
	if e.loadedTypes[1] {
		return e.Products, nil
	}
	return nil, &NotLoadedError{edge: "products"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProductCategoryMajor) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case productcategorymajor.FieldID:
			values[i] = new(sql.NullInt64)
		case productcategorymajor.FieldCategory:
			values[i] = new(sql.NullString)
		case productcategorymajor.FieldCreatedAt, productcategorymajor.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ProductCategoryMajor", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProductCategoryMajor fields.
func (pcm *ProductCategoryMajor) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case productcategorymajor.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pcm.ID = int(value.Int64)
		case productcategorymajor.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pcm.CreatedAt = value.Time
			}
		case productcategorymajor.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pcm.UpdatedAt = value.Time
			}
		case productcategorymajor.FieldCategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field category", values[i])
			} else if value.Valid {
				pcm.Category = value.String
			}
		}
	}
	return nil
}

// QueryMinors queries the "minors" edge of the ProductCategoryMajor entity.
func (pcm *ProductCategoryMajor) QueryMinors() *ProductCategoryMinorQuery {
	return (&ProductCategoryMajorClient{config: pcm.config}).QueryMinors(pcm)
}

// QueryProducts queries the "products" edge of the ProductCategoryMajor entity.
func (pcm *ProductCategoryMajor) QueryProducts() *ProductQuery {
	return (&ProductCategoryMajorClient{config: pcm.config}).QueryProducts(pcm)
}

// Update returns a builder for updating this ProductCategoryMajor.
// Note that you need to call ProductCategoryMajor.Unwrap() before calling this method if this ProductCategoryMajor
// was returned from a transaction, and the transaction was committed or rolled back.
func (pcm *ProductCategoryMajor) Update() *ProductCategoryMajorUpdateOne {
	return (&ProductCategoryMajorClient{config: pcm.config}).UpdateOne(pcm)
}

// Unwrap unwraps the ProductCategoryMajor entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pcm *ProductCategoryMajor) Unwrap() *ProductCategoryMajor {
	tx, ok := pcm.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProductCategoryMajor is not a transactional entity")
	}
	pcm.config.driver = tx.drv
	return pcm
}

// String implements the fmt.Stringer.
func (pcm *ProductCategoryMajor) String() string {
	var builder strings.Builder
	builder.WriteString("ProductCategoryMajor(")
	builder.WriteString(fmt.Sprintf("id=%v", pcm.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(pcm.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(pcm.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", category=")
	builder.WriteString(pcm.Category)
	builder.WriteByte(')')
	return builder.String()
}

// ProductCategoryMajors is a parsable slice of ProductCategoryMajor.
type ProductCategoryMajors []*ProductCategoryMajor

func (pcm ProductCategoryMajors) config(cfg config) {
	for _i := range pcm {
		pcm[_i].config = cfg
	}
}
