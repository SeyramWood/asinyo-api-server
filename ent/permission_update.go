// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SeyramWood/ent/permission"
	"github.com/SeyramWood/ent/predicate"
	"github.com/SeyramWood/ent/role"
)

// PermissionUpdate is the builder for updating Permission entities.
type PermissionUpdate struct {
	config
	hooks    []Hook
	mutation *PermissionMutation
}

// Where appends a list predicates to the PermissionUpdate builder.
func (pu *PermissionUpdate) Where(ps ...predicate.Permission) *PermissionUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PermissionUpdate) SetUpdatedAt(t time.Time) *PermissionUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetPermission sets the "permission" field.
func (pu *PermissionUpdate) SetPermission(s string) *PermissionUpdate {
	pu.mutation.SetPermission(s)
	return pu
}

// SetNillablePermission sets the "permission" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillablePermission(s *string) *PermissionUpdate {
	if s != nil {
		pu.SetPermission(*s)
	}
	return pu
}

// SetSlug sets the "slug" field.
func (pu *PermissionUpdate) SetSlug(s string) *PermissionUpdate {
	pu.mutation.SetSlug(s)
	return pu
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableSlug(s *string) *PermissionUpdate {
	if s != nil {
		pu.SetSlug(*s)
	}
	return pu
}

// AddRoleIDs adds the "role" edge to the Role entity by IDs.
func (pu *PermissionUpdate) AddRoleIDs(ids ...int) *PermissionUpdate {
	pu.mutation.AddRoleIDs(ids...)
	return pu
}

// AddRole adds the "role" edges to the Role entity.
func (pu *PermissionUpdate) AddRole(r ...*Role) *PermissionUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pu.AddRoleIDs(ids...)
}

// Mutation returns the PermissionMutation object of the builder.
func (pu *PermissionUpdate) Mutation() *PermissionMutation {
	return pu.mutation
}

// ClearRole clears all "role" edges to the Role entity.
func (pu *PermissionUpdate) ClearRole() *PermissionUpdate {
	pu.mutation.ClearRole()
	return pu
}

// RemoveRoleIDs removes the "role" edge to Role entities by IDs.
func (pu *PermissionUpdate) RemoveRoleIDs(ids ...int) *PermissionUpdate {
	pu.mutation.RemoveRoleIDs(ids...)
	return pu
}

// RemoveRole removes "role" edges to Role entities.
func (pu *PermissionUpdate) RemoveRole(r ...*Role) *PermissionUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pu.RemoveRoleIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PermissionUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PermissionUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PermissionUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PermissionUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PermissionUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := permission.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

func (pu *PermissionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(permission.Table, permission.Columns, sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(permission.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.Permission(); ok {
		_spec.SetField(permission.FieldPermission, field.TypeString, value)
	}
	if value, ok := pu.mutation.Slug(); ok {
		_spec.SetField(permission.FieldSlug, field.TypeString, value)
	}
	if pu.mutation.RoleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   permission.RoleTable,
			Columns: permission.RolePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedRoleIDs(); len(nodes) > 0 && !pu.mutation.RoleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   permission.RoleTable,
			Columns: permission.RolePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   permission.RoleTable,
			Columns: permission.RolePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{permission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PermissionUpdateOne is the builder for updating a single Permission entity.
type PermissionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PermissionMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PermissionUpdateOne) SetUpdatedAt(t time.Time) *PermissionUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetPermission sets the "permission" field.
func (puo *PermissionUpdateOne) SetPermission(s string) *PermissionUpdateOne {
	puo.mutation.SetPermission(s)
	return puo
}

// SetNillablePermission sets the "permission" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillablePermission(s *string) *PermissionUpdateOne {
	if s != nil {
		puo.SetPermission(*s)
	}
	return puo
}

// SetSlug sets the "slug" field.
func (puo *PermissionUpdateOne) SetSlug(s string) *PermissionUpdateOne {
	puo.mutation.SetSlug(s)
	return puo
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableSlug(s *string) *PermissionUpdateOne {
	if s != nil {
		puo.SetSlug(*s)
	}
	return puo
}

// AddRoleIDs adds the "role" edge to the Role entity by IDs.
func (puo *PermissionUpdateOne) AddRoleIDs(ids ...int) *PermissionUpdateOne {
	puo.mutation.AddRoleIDs(ids...)
	return puo
}

// AddRole adds the "role" edges to the Role entity.
func (puo *PermissionUpdateOne) AddRole(r ...*Role) *PermissionUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return puo.AddRoleIDs(ids...)
}

// Mutation returns the PermissionMutation object of the builder.
func (puo *PermissionUpdateOne) Mutation() *PermissionMutation {
	return puo.mutation
}

// ClearRole clears all "role" edges to the Role entity.
func (puo *PermissionUpdateOne) ClearRole() *PermissionUpdateOne {
	puo.mutation.ClearRole()
	return puo
}

// RemoveRoleIDs removes the "role" edge to Role entities by IDs.
func (puo *PermissionUpdateOne) RemoveRoleIDs(ids ...int) *PermissionUpdateOne {
	puo.mutation.RemoveRoleIDs(ids...)
	return puo
}

// RemoveRole removes "role" edges to Role entities.
func (puo *PermissionUpdateOne) RemoveRole(r ...*Role) *PermissionUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return puo.RemoveRoleIDs(ids...)
}

// Where appends a list predicates to the PermissionUpdate builder.
func (puo *PermissionUpdateOne) Where(ps ...predicate.Permission) *PermissionUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PermissionUpdateOne) Select(field string, fields ...string) *PermissionUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Permission entity.
func (puo *PermissionUpdateOne) Save(ctx context.Context) (*Permission, error) {
	puo.defaults()
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PermissionUpdateOne) SaveX(ctx context.Context) *Permission {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PermissionUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PermissionUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PermissionUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := permission.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

func (puo *PermissionUpdateOne) sqlSave(ctx context.Context) (_node *Permission, err error) {
	_spec := sqlgraph.NewUpdateSpec(permission.Table, permission.Columns, sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Permission.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, permission.FieldID)
		for _, f := range fields {
			if !permission.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != permission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(permission.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.Permission(); ok {
		_spec.SetField(permission.FieldPermission, field.TypeString, value)
	}
	if value, ok := puo.mutation.Slug(); ok {
		_spec.SetField(permission.FieldSlug, field.TypeString, value)
	}
	if puo.mutation.RoleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   permission.RoleTable,
			Columns: permission.RolePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedRoleIDs(); len(nodes) > 0 && !puo.mutation.RoleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   permission.RoleTable,
			Columns: permission.RolePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   permission.RoleTable,
			Columns: permission.RolePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Permission{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{permission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
