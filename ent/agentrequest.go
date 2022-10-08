// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/SeyramWood/ent/agentrequest"
)

// AgentRequest is the model entity for the AgentRequest schema.
type AgentRequest struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt               time.Time `json:"updated_at,omitempty"`
	agent_requests          *int
	merchant_store_requests *int
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AgentRequest) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case agentrequest.FieldID:
			values[i] = new(sql.NullInt64)
		case agentrequest.FieldCreatedAt, agentrequest.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case agentrequest.ForeignKeys[0]: // agent_requests
			values[i] = new(sql.NullInt64)
		case agentrequest.ForeignKeys[1]: // merchant_store_requests
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AgentRequest", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AgentRequest fields.
func (ar *AgentRequest) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case agentrequest.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ar.ID = int(value.Int64)
		case agentrequest.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ar.CreatedAt = value.Time
			}
		case agentrequest.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ar.UpdatedAt = value.Time
			}
		case agentrequest.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field agent_requests", value)
			} else if value.Valid {
				ar.agent_requests = new(int)
				*ar.agent_requests = int(value.Int64)
			}
		case agentrequest.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field merchant_store_requests", value)
			} else if value.Valid {
				ar.merchant_store_requests = new(int)
				*ar.merchant_store_requests = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AgentRequest.
// Note that you need to call AgentRequest.Unwrap() before calling this method if this AgentRequest
// was returned from a transaction, and the transaction was committed or rolled back.
func (ar *AgentRequest) Update() *AgentRequestUpdateOne {
	return (&AgentRequestClient{config: ar.config}).UpdateOne(ar)
}

// Unwrap unwraps the AgentRequest entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ar *AgentRequest) Unwrap() *AgentRequest {
	_tx, ok := ar.config.driver.(*txDriver)
	if !ok {
		panic("ent: AgentRequest is not a transactional entity")
	}
	ar.config.driver = _tx.drv
	return ar
}

// String implements the fmt.Stringer.
func (ar *AgentRequest) String() string {
	var builder strings.Builder
	builder.WriteString("AgentRequest(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ar.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ar.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ar.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// AgentRequests is a parsable slice of AgentRequest.
type AgentRequests []*AgentRequest

func (ar AgentRequests) config(cfg config) {
	for _i := range ar {
		ar[_i].config = cfg
	}
}
