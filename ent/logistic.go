// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/ent/logistic"
)

// Logistic is the model entity for the Logistic schema.
type Logistic struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// TrackingLink holds the value of the "tracking_link" field.
	TrackingLink string `json:"tracking_link,omitempty"`
	// Tasks holds the value of the "tasks" field.
	Tasks *models.TookanMultiTaskResponse `json:"tasks,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LogisticQuery when eager-loading is set.
	Edges LogisticEdges `json:"edges"`
}

// LogisticEdges holds the relations/edges for other nodes in the graph.
type LogisticEdges struct {
	// Order holds the value of the order edge.
	Order []*Order `json:"order,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OrderOrErr returns the Order value or an error if the edge
// was not loaded in eager-loading.
func (e LogisticEdges) OrderOrErr() ([]*Order, error) {
	if e.loadedTypes[0] {
		return e.Order, nil
	}
	return nil, &NotLoadedError{edge: "order"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Logistic) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case logistic.FieldTasks:
			values[i] = new([]byte)
		case logistic.FieldID:
			values[i] = new(sql.NullInt64)
		case logistic.FieldTrackingLink:
			values[i] = new(sql.NullString)
		case logistic.FieldCreatedAt, logistic.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Logistic", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Logistic fields.
func (l *Logistic) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case logistic.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			l.ID = int(value.Int64)
		case logistic.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				l.CreatedAt = value.Time
			}
		case logistic.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				l.UpdatedAt = value.Time
			}
		case logistic.FieldTrackingLink:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tracking_link", values[i])
			} else if value.Valid {
				l.TrackingLink = value.String
			}
		case logistic.FieldTasks:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tasks", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &l.Tasks); err != nil {
					return fmt.Errorf("unmarshal field tasks: %w", err)
				}
			}
		}
	}
	return nil
}

// QueryOrder queries the "order" edge of the Logistic entity.
func (l *Logistic) QueryOrder() *OrderQuery {
	return (&LogisticClient{config: l.config}).QueryOrder(l)
}

// Update returns a builder for updating this Logistic.
// Note that you need to call Logistic.Unwrap() before calling this method if this Logistic
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Logistic) Update() *LogisticUpdateOne {
	return (&LogisticClient{config: l.config}).UpdateOne(l)
}

// Unwrap unwraps the Logistic entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *Logistic) Unwrap() *Logistic {
	_tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: Logistic is not a transactional entity")
	}
	l.config.driver = _tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Logistic) String() string {
	var builder strings.Builder
	builder.WriteString("Logistic(")
	builder.WriteString(fmt.Sprintf("id=%v, ", l.ID))
	builder.WriteString("created_at=")
	builder.WriteString(l.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(l.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("tracking_link=")
	builder.WriteString(l.TrackingLink)
	builder.WriteString(", ")
	builder.WriteString("tasks=")
	builder.WriteString(fmt.Sprintf("%v", l.Tasks))
	builder.WriteByte(')')
	return builder.String()
}

// Logistics is a parsable slice of Logistic.
type Logistics []*Logistic

func (l Logistics) config(cfg config) {
	for _i := range l {
		l[_i].config = cfg
	}
}