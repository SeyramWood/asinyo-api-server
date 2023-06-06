// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SeyramWood/ent/customer"
	"github.com/SeyramWood/ent/individualcustomer"
)

// IndividualCustomerCreate is the builder for creating a IndividualCustomer entity.
type IndividualCustomerCreate struct {
	config
	mutation *IndividualCustomerMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (icc *IndividualCustomerCreate) SetCreatedAt(t time.Time) *IndividualCustomerCreate {
	icc.mutation.SetCreatedAt(t)
	return icc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (icc *IndividualCustomerCreate) SetNillableCreatedAt(t *time.Time) *IndividualCustomerCreate {
	if t != nil {
		icc.SetCreatedAt(*t)
	}
	return icc
}

// SetUpdatedAt sets the "updated_at" field.
func (icc *IndividualCustomerCreate) SetUpdatedAt(t time.Time) *IndividualCustomerCreate {
	icc.mutation.SetUpdatedAt(t)
	return icc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (icc *IndividualCustomerCreate) SetNillableUpdatedAt(t *time.Time) *IndividualCustomerCreate {
	if t != nil {
		icc.SetUpdatedAt(*t)
	}
	return icc
}

// SetLastName sets the "last_name" field.
func (icc *IndividualCustomerCreate) SetLastName(s string) *IndividualCustomerCreate {
	icc.mutation.SetLastName(s)
	return icc
}

// SetOtherName sets the "other_name" field.
func (icc *IndividualCustomerCreate) SetOtherName(s string) *IndividualCustomerCreate {
	icc.mutation.SetOtherName(s)
	return icc
}

// SetPhone sets the "phone" field.
func (icc *IndividualCustomerCreate) SetPhone(s string) *IndividualCustomerCreate {
	icc.mutation.SetPhone(s)
	return icc
}

// SetOtherPhone sets the "other_phone" field.
func (icc *IndividualCustomerCreate) SetOtherPhone(s string) *IndividualCustomerCreate {
	icc.mutation.SetOtherPhone(s)
	return icc
}

// SetNillableOtherPhone sets the "other_phone" field if the given value is not nil.
func (icc *IndividualCustomerCreate) SetNillableOtherPhone(s *string) *IndividualCustomerCreate {
	if s != nil {
		icc.SetOtherPhone(*s)
	}
	return icc
}

// SetCustomerID sets the "customer" edge to the Customer entity by ID.
func (icc *IndividualCustomerCreate) SetCustomerID(id int) *IndividualCustomerCreate {
	icc.mutation.SetCustomerID(id)
	return icc
}

// SetCustomer sets the "customer" edge to the Customer entity.
func (icc *IndividualCustomerCreate) SetCustomer(c *Customer) *IndividualCustomerCreate {
	return icc.SetCustomerID(c.ID)
}

// Mutation returns the IndividualCustomerMutation object of the builder.
func (icc *IndividualCustomerCreate) Mutation() *IndividualCustomerMutation {
	return icc.mutation
}

// Save creates the IndividualCustomer in the database.
func (icc *IndividualCustomerCreate) Save(ctx context.Context) (*IndividualCustomer, error) {
	icc.defaults()
	return withHooks(ctx, icc.sqlSave, icc.mutation, icc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (icc *IndividualCustomerCreate) SaveX(ctx context.Context) *IndividualCustomer {
	v, err := icc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icc *IndividualCustomerCreate) Exec(ctx context.Context) error {
	_, err := icc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icc *IndividualCustomerCreate) ExecX(ctx context.Context) {
	if err := icc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (icc *IndividualCustomerCreate) defaults() {
	if _, ok := icc.mutation.CreatedAt(); !ok {
		v := individualcustomer.DefaultCreatedAt()
		icc.mutation.SetCreatedAt(v)
	}
	if _, ok := icc.mutation.UpdatedAt(); !ok {
		v := individualcustomer.DefaultUpdatedAt()
		icc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (icc *IndividualCustomerCreate) check() error {
	if _, ok := icc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "IndividualCustomer.created_at"`)}
	}
	if _, ok := icc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "IndividualCustomer.updated_at"`)}
	}
	if _, ok := icc.mutation.LastName(); !ok {
		return &ValidationError{Name: "last_name", err: errors.New(`ent: missing required field "IndividualCustomer.last_name"`)}
	}
	if v, ok := icc.mutation.LastName(); ok {
		if err := individualcustomer.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "IndividualCustomer.last_name": %w`, err)}
		}
	}
	if _, ok := icc.mutation.OtherName(); !ok {
		return &ValidationError{Name: "other_name", err: errors.New(`ent: missing required field "IndividualCustomer.other_name"`)}
	}
	if v, ok := icc.mutation.OtherName(); ok {
		if err := individualcustomer.OtherNameValidator(v); err != nil {
			return &ValidationError{Name: "other_name", err: fmt.Errorf(`ent: validator failed for field "IndividualCustomer.other_name": %w`, err)}
		}
	}
	if _, ok := icc.mutation.Phone(); !ok {
		return &ValidationError{Name: "phone", err: errors.New(`ent: missing required field "IndividualCustomer.phone"`)}
	}
	if v, ok := icc.mutation.Phone(); ok {
		if err := individualcustomer.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "IndividualCustomer.phone": %w`, err)}
		}
	}
	if _, ok := icc.mutation.CustomerID(); !ok {
		return &ValidationError{Name: "customer", err: errors.New(`ent: missing required edge "IndividualCustomer.customer"`)}
	}
	return nil
}

func (icc *IndividualCustomerCreate) sqlSave(ctx context.Context) (*IndividualCustomer, error) {
	if err := icc.check(); err != nil {
		return nil, err
	}
	_node, _spec := icc.createSpec()
	if err := sqlgraph.CreateNode(ctx, icc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	icc.mutation.id = &_node.ID
	icc.mutation.done = true
	return _node, nil
}

func (icc *IndividualCustomerCreate) createSpec() (*IndividualCustomer, *sqlgraph.CreateSpec) {
	var (
		_node = &IndividualCustomer{config: icc.config}
		_spec = sqlgraph.NewCreateSpec(individualcustomer.Table, sqlgraph.NewFieldSpec(individualcustomer.FieldID, field.TypeInt))
	)
	if value, ok := icc.mutation.CreatedAt(); ok {
		_spec.SetField(individualcustomer.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := icc.mutation.UpdatedAt(); ok {
		_spec.SetField(individualcustomer.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := icc.mutation.LastName(); ok {
		_spec.SetField(individualcustomer.FieldLastName, field.TypeString, value)
		_node.LastName = value
	}
	if value, ok := icc.mutation.OtherName(); ok {
		_spec.SetField(individualcustomer.FieldOtherName, field.TypeString, value)
		_node.OtherName = value
	}
	if value, ok := icc.mutation.Phone(); ok {
		_spec.SetField(individualcustomer.FieldPhone, field.TypeString, value)
		_node.Phone = value
	}
	if value, ok := icc.mutation.OtherPhone(); ok {
		_spec.SetField(individualcustomer.FieldOtherPhone, field.TypeString, value)
		_node.OtherPhone = value
	}
	if nodes := icc.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   individualcustomer.CustomerTable,
			Columns: []string{individualcustomer.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.customer_individual = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// IndividualCustomerCreateBulk is the builder for creating many IndividualCustomer entities in bulk.
type IndividualCustomerCreateBulk struct {
	config
	builders []*IndividualCustomerCreate
}

// Save creates the IndividualCustomer entities in the database.
func (iccb *IndividualCustomerCreateBulk) Save(ctx context.Context) ([]*IndividualCustomer, error) {
	specs := make([]*sqlgraph.CreateSpec, len(iccb.builders))
	nodes := make([]*IndividualCustomer, len(iccb.builders))
	mutators := make([]Mutator, len(iccb.builders))
	for i := range iccb.builders {
		func(i int, root context.Context) {
			builder := iccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*IndividualCustomerMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, iccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, iccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, iccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (iccb *IndividualCustomerCreateBulk) SaveX(ctx context.Context) []*IndividualCustomer {
	v, err := iccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (iccb *IndividualCustomerCreateBulk) Exec(ctx context.Context) error {
	_, err := iccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iccb *IndividualCustomerCreateBulk) ExecX(ctx context.Context) {
	if err := iccb.Exec(ctx); err != nil {
		panic(err)
	}
}
