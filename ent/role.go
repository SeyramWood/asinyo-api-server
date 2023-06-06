// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/SeyramWood/ent/role"
)

// Role is the model entity for the Role schema.
type Role struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Role holds the value of the "role" field.
	Role string `json:"role,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoleQuery when eager-loading is set.
	Edges        RoleEdges `json:"edges"`
	selectValues sql.SelectValues
}

// RoleEdges holds the relations/edges for other nodes in the graph.
type RoleEdges struct {
	// Permissions holds the value of the permissions edge.
	Permissions []*Permission `json:"permissions,omitempty"`
	// Admin holds the value of the admin edge.
	Admin []*Admin `json:"admin,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// PermissionsOrErr returns the Permissions value or an error if the edge
// was not loaded in eager-loading.
func (e RoleEdges) PermissionsOrErr() ([]*Permission, error) {
	if e.loadedTypes[0] {
		return e.Permissions, nil
	}
	return nil, &NotLoadedError{edge: "permissions"}
}

// AdminOrErr returns the Admin value or an error if the edge
// was not loaded in eager-loading.
func (e RoleEdges) AdminOrErr() ([]*Admin, error) {
	if e.loadedTypes[1] {
		return e.Admin, nil
	}
	return nil, &NotLoadedError{edge: "admin"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Role) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case role.FieldID:
			values[i] = new(sql.NullInt64)
		case role.FieldRole:
			values[i] = new(sql.NullString)
		case role.FieldCreatedAt, role.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Role fields.
func (r *Role) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case role.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case role.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		case role.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				r.UpdatedAt = value.Time
			}
		case role.FieldRole:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				r.Role = value.String
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Role.
// This includes values selected through modifiers, order, etc.
func (r *Role) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// QueryPermissions queries the "permissions" edge of the Role entity.
func (r *Role) QueryPermissions() *PermissionQuery {
	return NewRoleClient(r.config).QueryPermissions(r)
}

// QueryAdmin queries the "admin" edge of the Role entity.
func (r *Role) QueryAdmin() *AdminQuery {
	return NewRoleClient(r.config).QueryAdmin(r)
}

// Update returns a builder for updating this Role.
// Note that you need to call Role.Unwrap() before calling this method if this Role
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Role) Update() *RoleUpdateOne {
	return NewRoleClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Role entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Role) Unwrap() *Role {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Role is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Role) String() string {
	var builder strings.Builder
	builder.WriteString("Role(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("created_at=")
	builder.WriteString(r.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(r.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("role=")
	builder.WriteString(r.Role)
	builder.WriteByte(')')
	return builder.String()
}

// Roles is a parsable slice of Role.
type Roles []*Role
