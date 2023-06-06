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
	"github.com/SeyramWood/ent/customer"
	"github.com/SeyramWood/ent/individualcustomer"
	"github.com/SeyramWood/ent/predicate"
)

// IndividualCustomerUpdate is the builder for updating IndividualCustomer entities.
type IndividualCustomerUpdate struct {
	config
	hooks    []Hook
	mutation *IndividualCustomerMutation
}

// Where appends a list predicates to the IndividualCustomerUpdate builder.
func (icu *IndividualCustomerUpdate) Where(ps ...predicate.IndividualCustomer) *IndividualCustomerUpdate {
	icu.mutation.Where(ps...)
	return icu
}

// SetUpdatedAt sets the "updated_at" field.
func (icu *IndividualCustomerUpdate) SetUpdatedAt(t time.Time) *IndividualCustomerUpdate {
	icu.mutation.SetUpdatedAt(t)
	return icu
}

// SetLastName sets the "last_name" field.
func (icu *IndividualCustomerUpdate) SetLastName(s string) *IndividualCustomerUpdate {
	icu.mutation.SetLastName(s)
	return icu
}

// SetOtherName sets the "other_name" field.
func (icu *IndividualCustomerUpdate) SetOtherName(s string) *IndividualCustomerUpdate {
	icu.mutation.SetOtherName(s)
	return icu
}

// SetPhone sets the "phone" field.
func (icu *IndividualCustomerUpdate) SetPhone(s string) *IndividualCustomerUpdate {
	icu.mutation.SetPhone(s)
	return icu
}

// SetOtherPhone sets the "other_phone" field.
func (icu *IndividualCustomerUpdate) SetOtherPhone(s string) *IndividualCustomerUpdate {
	icu.mutation.SetOtherPhone(s)
	return icu
}

// SetNillableOtherPhone sets the "other_phone" field if the given value is not nil.
func (icu *IndividualCustomerUpdate) SetNillableOtherPhone(s *string) *IndividualCustomerUpdate {
	if s != nil {
		icu.SetOtherPhone(*s)
	}
	return icu
}

// ClearOtherPhone clears the value of the "other_phone" field.
func (icu *IndividualCustomerUpdate) ClearOtherPhone() *IndividualCustomerUpdate {
	icu.mutation.ClearOtherPhone()
	return icu
}

// SetCustomerID sets the "customer" edge to the Customer entity by ID.
func (icu *IndividualCustomerUpdate) SetCustomerID(id int) *IndividualCustomerUpdate {
	icu.mutation.SetCustomerID(id)
	return icu
}

// SetCustomer sets the "customer" edge to the Customer entity.
func (icu *IndividualCustomerUpdate) SetCustomer(c *Customer) *IndividualCustomerUpdate {
	return icu.SetCustomerID(c.ID)
}

// Mutation returns the IndividualCustomerMutation object of the builder.
func (icu *IndividualCustomerUpdate) Mutation() *IndividualCustomerMutation {
	return icu.mutation
}

// ClearCustomer clears the "customer" edge to the Customer entity.
func (icu *IndividualCustomerUpdate) ClearCustomer() *IndividualCustomerUpdate {
	icu.mutation.ClearCustomer()
	return icu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (icu *IndividualCustomerUpdate) Save(ctx context.Context) (int, error) {
	icu.defaults()
	return withHooks(ctx, icu.sqlSave, icu.mutation, icu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (icu *IndividualCustomerUpdate) SaveX(ctx context.Context) int {
	affected, err := icu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (icu *IndividualCustomerUpdate) Exec(ctx context.Context) error {
	_, err := icu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icu *IndividualCustomerUpdate) ExecX(ctx context.Context) {
	if err := icu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (icu *IndividualCustomerUpdate) defaults() {
	if _, ok := icu.mutation.UpdatedAt(); !ok {
		v := individualcustomer.UpdateDefaultUpdatedAt()
		icu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (icu *IndividualCustomerUpdate) check() error {
	if v, ok := icu.mutation.LastName(); ok {
		if err := individualcustomer.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "IndividualCustomer.last_name": %w`, err)}
		}
	}
	if v, ok := icu.mutation.OtherName(); ok {
		if err := individualcustomer.OtherNameValidator(v); err != nil {
			return &ValidationError{Name: "other_name", err: fmt.Errorf(`ent: validator failed for field "IndividualCustomer.other_name": %w`, err)}
		}
	}
	if v, ok := icu.mutation.Phone(); ok {
		if err := individualcustomer.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "IndividualCustomer.phone": %w`, err)}
		}
	}
	if _, ok := icu.mutation.CustomerID(); icu.mutation.CustomerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "IndividualCustomer.customer"`)
	}
	return nil
}

func (icu *IndividualCustomerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := icu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(individualcustomer.Table, individualcustomer.Columns, sqlgraph.NewFieldSpec(individualcustomer.FieldID, field.TypeInt))
	if ps := icu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := icu.mutation.UpdatedAt(); ok {
		_spec.SetField(individualcustomer.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := icu.mutation.LastName(); ok {
		_spec.SetField(individualcustomer.FieldLastName, field.TypeString, value)
	}
	if value, ok := icu.mutation.OtherName(); ok {
		_spec.SetField(individualcustomer.FieldOtherName, field.TypeString, value)
	}
	if value, ok := icu.mutation.Phone(); ok {
		_spec.SetField(individualcustomer.FieldPhone, field.TypeString, value)
	}
	if value, ok := icu.mutation.OtherPhone(); ok {
		_spec.SetField(individualcustomer.FieldOtherPhone, field.TypeString, value)
	}
	if icu.mutation.OtherPhoneCleared() {
		_spec.ClearField(individualcustomer.FieldOtherPhone, field.TypeString)
	}
	if icu.mutation.CustomerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := icu.mutation.CustomerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, icu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{individualcustomer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	icu.mutation.done = true
	return n, nil
}

// IndividualCustomerUpdateOne is the builder for updating a single IndividualCustomer entity.
type IndividualCustomerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *IndividualCustomerMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (icuo *IndividualCustomerUpdateOne) SetUpdatedAt(t time.Time) *IndividualCustomerUpdateOne {
	icuo.mutation.SetUpdatedAt(t)
	return icuo
}

// SetLastName sets the "last_name" field.
func (icuo *IndividualCustomerUpdateOne) SetLastName(s string) *IndividualCustomerUpdateOne {
	icuo.mutation.SetLastName(s)
	return icuo
}

// SetOtherName sets the "other_name" field.
func (icuo *IndividualCustomerUpdateOne) SetOtherName(s string) *IndividualCustomerUpdateOne {
	icuo.mutation.SetOtherName(s)
	return icuo
}

// SetPhone sets the "phone" field.
func (icuo *IndividualCustomerUpdateOne) SetPhone(s string) *IndividualCustomerUpdateOne {
	icuo.mutation.SetPhone(s)
	return icuo
}

// SetOtherPhone sets the "other_phone" field.
func (icuo *IndividualCustomerUpdateOne) SetOtherPhone(s string) *IndividualCustomerUpdateOne {
	icuo.mutation.SetOtherPhone(s)
	return icuo
}

// SetNillableOtherPhone sets the "other_phone" field if the given value is not nil.
func (icuo *IndividualCustomerUpdateOne) SetNillableOtherPhone(s *string) *IndividualCustomerUpdateOne {
	if s != nil {
		icuo.SetOtherPhone(*s)
	}
	return icuo
}

// ClearOtherPhone clears the value of the "other_phone" field.
func (icuo *IndividualCustomerUpdateOne) ClearOtherPhone() *IndividualCustomerUpdateOne {
	icuo.mutation.ClearOtherPhone()
	return icuo
}

// SetCustomerID sets the "customer" edge to the Customer entity by ID.
func (icuo *IndividualCustomerUpdateOne) SetCustomerID(id int) *IndividualCustomerUpdateOne {
	icuo.mutation.SetCustomerID(id)
	return icuo
}

// SetCustomer sets the "customer" edge to the Customer entity.
func (icuo *IndividualCustomerUpdateOne) SetCustomer(c *Customer) *IndividualCustomerUpdateOne {
	return icuo.SetCustomerID(c.ID)
}

// Mutation returns the IndividualCustomerMutation object of the builder.
func (icuo *IndividualCustomerUpdateOne) Mutation() *IndividualCustomerMutation {
	return icuo.mutation
}

// ClearCustomer clears the "customer" edge to the Customer entity.
func (icuo *IndividualCustomerUpdateOne) ClearCustomer() *IndividualCustomerUpdateOne {
	icuo.mutation.ClearCustomer()
	return icuo
}

// Where appends a list predicates to the IndividualCustomerUpdate builder.
func (icuo *IndividualCustomerUpdateOne) Where(ps ...predicate.IndividualCustomer) *IndividualCustomerUpdateOne {
	icuo.mutation.Where(ps...)
	return icuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (icuo *IndividualCustomerUpdateOne) Select(field string, fields ...string) *IndividualCustomerUpdateOne {
	icuo.fields = append([]string{field}, fields...)
	return icuo
}

// Save executes the query and returns the updated IndividualCustomer entity.
func (icuo *IndividualCustomerUpdateOne) Save(ctx context.Context) (*IndividualCustomer, error) {
	icuo.defaults()
	return withHooks(ctx, icuo.sqlSave, icuo.mutation, icuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (icuo *IndividualCustomerUpdateOne) SaveX(ctx context.Context) *IndividualCustomer {
	node, err := icuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (icuo *IndividualCustomerUpdateOne) Exec(ctx context.Context) error {
	_, err := icuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icuo *IndividualCustomerUpdateOne) ExecX(ctx context.Context) {
	if err := icuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (icuo *IndividualCustomerUpdateOne) defaults() {
	if _, ok := icuo.mutation.UpdatedAt(); !ok {
		v := individualcustomer.UpdateDefaultUpdatedAt()
		icuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (icuo *IndividualCustomerUpdateOne) check() error {
	if v, ok := icuo.mutation.LastName(); ok {
		if err := individualcustomer.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "IndividualCustomer.last_name": %w`, err)}
		}
	}
	if v, ok := icuo.mutation.OtherName(); ok {
		if err := individualcustomer.OtherNameValidator(v); err != nil {
			return &ValidationError{Name: "other_name", err: fmt.Errorf(`ent: validator failed for field "IndividualCustomer.other_name": %w`, err)}
		}
	}
	if v, ok := icuo.mutation.Phone(); ok {
		if err := individualcustomer.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "IndividualCustomer.phone": %w`, err)}
		}
	}
	if _, ok := icuo.mutation.CustomerID(); icuo.mutation.CustomerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "IndividualCustomer.customer"`)
	}
	return nil
}

func (icuo *IndividualCustomerUpdateOne) sqlSave(ctx context.Context) (_node *IndividualCustomer, err error) {
	if err := icuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(individualcustomer.Table, individualcustomer.Columns, sqlgraph.NewFieldSpec(individualcustomer.FieldID, field.TypeInt))
	id, ok := icuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "IndividualCustomer.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := icuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, individualcustomer.FieldID)
		for _, f := range fields {
			if !individualcustomer.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != individualcustomer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := icuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := icuo.mutation.UpdatedAt(); ok {
		_spec.SetField(individualcustomer.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := icuo.mutation.LastName(); ok {
		_spec.SetField(individualcustomer.FieldLastName, field.TypeString, value)
	}
	if value, ok := icuo.mutation.OtherName(); ok {
		_spec.SetField(individualcustomer.FieldOtherName, field.TypeString, value)
	}
	if value, ok := icuo.mutation.Phone(); ok {
		_spec.SetField(individualcustomer.FieldPhone, field.TypeString, value)
	}
	if value, ok := icuo.mutation.OtherPhone(); ok {
		_spec.SetField(individualcustomer.FieldOtherPhone, field.TypeString, value)
	}
	if icuo.mutation.OtherPhoneCleared() {
		_spec.ClearField(individualcustomer.FieldOtherPhone, field.TypeString)
	}
	if icuo.mutation.CustomerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := icuo.mutation.CustomerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &IndividualCustomer{config: icuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, icuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{individualcustomer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	icuo.mutation.done = true
	return _node, nil
}
