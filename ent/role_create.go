// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SeyramWood/ent/admin"
	"github.com/SeyramWood/ent/permission"
	"github.com/SeyramWood/ent/role"
)

// RoleCreate is the builder for creating a Role entity.
type RoleCreate struct {
	config
	mutation *RoleMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (rc *RoleCreate) SetCreatedAt(t time.Time) *RoleCreate {
	rc.mutation.SetCreatedAt(t)
	return rc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rc *RoleCreate) SetNillableCreatedAt(t *time.Time) *RoleCreate {
	if t != nil {
		rc.SetCreatedAt(*t)
	}
	return rc
}

// SetUpdatedAt sets the "updated_at" field.
func (rc *RoleCreate) SetUpdatedAt(t time.Time) *RoleCreate {
	rc.mutation.SetUpdatedAt(t)
	return rc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rc *RoleCreate) SetNillableUpdatedAt(t *time.Time) *RoleCreate {
	if t != nil {
		rc.SetUpdatedAt(*t)
	}
	return rc
}

// SetRole sets the "role" field.
func (rc *RoleCreate) SetRole(s string) *RoleCreate {
	rc.mutation.SetRole(s)
	return rc
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (rc *RoleCreate) SetNillableRole(s *string) *RoleCreate {
	if s != nil {
		rc.SetRole(*s)
	}
	return rc
}

// AddPermissionIDs adds the "permissions" edge to the Permission entity by IDs.
func (rc *RoleCreate) AddPermissionIDs(ids ...int) *RoleCreate {
	rc.mutation.AddPermissionIDs(ids...)
	return rc
}

// AddPermissions adds the "permissions" edges to the Permission entity.
func (rc *RoleCreate) AddPermissions(p ...*Permission) *RoleCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return rc.AddPermissionIDs(ids...)
}

// AddAdminIDs adds the "admin" edge to the Admin entity by IDs.
func (rc *RoleCreate) AddAdminIDs(ids ...int) *RoleCreate {
	rc.mutation.AddAdminIDs(ids...)
	return rc
}

// AddAdmin adds the "admin" edges to the Admin entity.
func (rc *RoleCreate) AddAdmin(a ...*Admin) *RoleCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return rc.AddAdminIDs(ids...)
}

// Mutation returns the RoleMutation object of the builder.
func (rc *RoleCreate) Mutation() *RoleMutation {
	return rc.mutation
}

// Save creates the Role in the database.
func (rc *RoleCreate) Save(ctx context.Context) (*Role, error) {
	var (
		err  error
		node *Role
	)
	rc.defaults()
	if len(rc.hooks) == 0 {
		if err = rc.check(); err != nil {
			return nil, err
		}
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rc.check(); err != nil {
				return nil, err
			}
			rc.mutation = mutation
			if node, err = rc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			if rc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, rc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Role)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from RoleMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RoleCreate) SaveX(ctx context.Context) *Role {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RoleCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RoleCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RoleCreate) defaults() {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		v := role.DefaultCreatedAt()
		rc.mutation.SetCreatedAt(v)
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		v := role.DefaultUpdatedAt()
		rc.mutation.SetUpdatedAt(v)
	}
	if _, ok := rc.mutation.Role(); !ok {
		v := role.DefaultRole
		rc.mutation.SetRole(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RoleCreate) check() error {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Role.created_at"`)}
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Role.updated_at"`)}
	}
	if _, ok := rc.mutation.Role(); !ok {
		return &ValidationError{Name: "role", err: errors.New(`ent: missing required field "Role.role"`)}
	}
	return nil
}

func (rc *RoleCreate) sqlSave(ctx context.Context) (*Role, error) {
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (rc *RoleCreate) createSpec() (*Role, *sqlgraph.CreateSpec) {
	var (
		_node = &Role{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: role.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: role.FieldID,
			},
		}
	)
	if value, ok := rc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: role.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := rc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: role.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := rc.mutation.Role(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: role.FieldRole,
		})
		_node.Role = value
	}
	if nodes := rc.mutation.PermissionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.PermissionsTable,
			Columns: role.PermissionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: permission.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.AdminIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   role.AdminTable,
			Columns: role.AdminPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: admin.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RoleCreateBulk is the builder for creating many Role entities in bulk.
type RoleCreateBulk struct {
	config
	builders []*RoleCreate
}

// Save creates the Role entities in the database.
func (rcb *RoleCreateBulk) Save(ctx context.Context) ([]*Role, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Role, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RoleMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RoleCreateBulk) SaveX(ctx context.Context) []*Role {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RoleCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RoleCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
