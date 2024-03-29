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
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/ent/businesscustomer"
	"github.com/SeyramWood/ent/customer"
	"github.com/SeyramWood/ent/predicate"
)

// BusinessCustomerUpdate is the builder for updating BusinessCustomer entities.
type BusinessCustomerUpdate struct {
	config
	hooks    []Hook
	mutation *BusinessCustomerMutation
}

// Where appends a list predicates to the BusinessCustomerUpdate builder.
func (bcu *BusinessCustomerUpdate) Where(ps ...predicate.BusinessCustomer) *BusinessCustomerUpdate {
	bcu.mutation.Where(ps...)
	return bcu
}

// SetUpdatedAt sets the "updated_at" field.
func (bcu *BusinessCustomerUpdate) SetUpdatedAt(t time.Time) *BusinessCustomerUpdate {
	bcu.mutation.SetUpdatedAt(t)
	return bcu
}

// SetName sets the "name" field.
func (bcu *BusinessCustomerUpdate) SetName(s string) *BusinessCustomerUpdate {
	bcu.mutation.SetName(s)
	return bcu
}

// SetPhone sets the "phone" field.
func (bcu *BusinessCustomerUpdate) SetPhone(s string) *BusinessCustomerUpdate {
	bcu.mutation.SetPhone(s)
	return bcu
}

// SetOtherPhone sets the "other_phone" field.
func (bcu *BusinessCustomerUpdate) SetOtherPhone(s string) *BusinessCustomerUpdate {
	bcu.mutation.SetOtherPhone(s)
	return bcu
}

// SetNillableOtherPhone sets the "other_phone" field if the given value is not nil.
func (bcu *BusinessCustomerUpdate) SetNillableOtherPhone(s *string) *BusinessCustomerUpdate {
	if s != nil {
		bcu.SetOtherPhone(*s)
	}
	return bcu
}

// ClearOtherPhone clears the value of the "other_phone" field.
func (bcu *BusinessCustomerUpdate) ClearOtherPhone() *BusinessCustomerUpdate {
	bcu.mutation.ClearOtherPhone()
	return bcu
}

// SetLogo sets the "logo" field.
func (bcu *BusinessCustomerUpdate) SetLogo(s string) *BusinessCustomerUpdate {
	bcu.mutation.SetLogo(s)
	return bcu
}

// SetNillableLogo sets the "logo" field if the given value is not nil.
func (bcu *BusinessCustomerUpdate) SetNillableLogo(s *string) *BusinessCustomerUpdate {
	if s != nil {
		bcu.SetLogo(*s)
	}
	return bcu
}

// ClearLogo clears the value of the "logo" field.
func (bcu *BusinessCustomerUpdate) ClearLogo() *BusinessCustomerUpdate {
	bcu.mutation.ClearLogo()
	return bcu
}

// SetContact sets the "contact" field.
func (bcu *BusinessCustomerUpdate) SetContact(mcc *models.BusinessCustomerContact) *BusinessCustomerUpdate {
	bcu.mutation.SetContact(mcc)
	return bcu
}

// ClearContact clears the value of the "contact" field.
func (bcu *BusinessCustomerUpdate) ClearContact() *BusinessCustomerUpdate {
	bcu.mutation.ClearContact()
	return bcu
}

// SetCustomerID sets the "customer" edge to the Customer entity by ID.
func (bcu *BusinessCustomerUpdate) SetCustomerID(id int) *BusinessCustomerUpdate {
	bcu.mutation.SetCustomerID(id)
	return bcu
}

// SetCustomer sets the "customer" edge to the Customer entity.
func (bcu *BusinessCustomerUpdate) SetCustomer(c *Customer) *BusinessCustomerUpdate {
	return bcu.SetCustomerID(c.ID)
}

// Mutation returns the BusinessCustomerMutation object of the builder.
func (bcu *BusinessCustomerUpdate) Mutation() *BusinessCustomerMutation {
	return bcu.mutation
}

// ClearCustomer clears the "customer" edge to the Customer entity.
func (bcu *BusinessCustomerUpdate) ClearCustomer() *BusinessCustomerUpdate {
	bcu.mutation.ClearCustomer()
	return bcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bcu *BusinessCustomerUpdate) Save(ctx context.Context) (int, error) {
	bcu.defaults()
	return withHooks(ctx, bcu.sqlSave, bcu.mutation, bcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bcu *BusinessCustomerUpdate) SaveX(ctx context.Context) int {
	affected, err := bcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bcu *BusinessCustomerUpdate) Exec(ctx context.Context) error {
	_, err := bcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcu *BusinessCustomerUpdate) ExecX(ctx context.Context) {
	if err := bcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bcu *BusinessCustomerUpdate) defaults() {
	if _, ok := bcu.mutation.UpdatedAt(); !ok {
		v := businesscustomer.UpdateDefaultUpdatedAt()
		bcu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bcu *BusinessCustomerUpdate) check() error {
	if v, ok := bcu.mutation.Name(); ok {
		if err := businesscustomer.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "BusinessCustomer.name": %w`, err)}
		}
	}
	if v, ok := bcu.mutation.Phone(); ok {
		if err := businesscustomer.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "BusinessCustomer.phone": %w`, err)}
		}
	}
	if _, ok := bcu.mutation.CustomerID(); bcu.mutation.CustomerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "BusinessCustomer.customer"`)
	}
	return nil
}

func (bcu *BusinessCustomerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := bcu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(businesscustomer.Table, businesscustomer.Columns, sqlgraph.NewFieldSpec(businesscustomer.FieldID, field.TypeInt))
	if ps := bcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bcu.mutation.UpdatedAt(); ok {
		_spec.SetField(businesscustomer.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := bcu.mutation.Name(); ok {
		_spec.SetField(businesscustomer.FieldName, field.TypeString, value)
	}
	if value, ok := bcu.mutation.Phone(); ok {
		_spec.SetField(businesscustomer.FieldPhone, field.TypeString, value)
	}
	if value, ok := bcu.mutation.OtherPhone(); ok {
		_spec.SetField(businesscustomer.FieldOtherPhone, field.TypeString, value)
	}
	if bcu.mutation.OtherPhoneCleared() {
		_spec.ClearField(businesscustomer.FieldOtherPhone, field.TypeString)
	}
	if value, ok := bcu.mutation.Logo(); ok {
		_spec.SetField(businesscustomer.FieldLogo, field.TypeString, value)
	}
	if bcu.mutation.LogoCleared() {
		_spec.ClearField(businesscustomer.FieldLogo, field.TypeString)
	}
	if value, ok := bcu.mutation.Contact(); ok {
		_spec.SetField(businesscustomer.FieldContact, field.TypeJSON, value)
	}
	if bcu.mutation.ContactCleared() {
		_spec.ClearField(businesscustomer.FieldContact, field.TypeJSON)
	}
	if bcu.mutation.CustomerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   businesscustomer.CustomerTable,
			Columns: []string{businesscustomer.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcu.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   businesscustomer.CustomerTable,
			Columns: []string{businesscustomer.CustomerColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, bcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{businesscustomer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bcu.mutation.done = true
	return n, nil
}

// BusinessCustomerUpdateOne is the builder for updating a single BusinessCustomer entity.
type BusinessCustomerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BusinessCustomerMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (bcuo *BusinessCustomerUpdateOne) SetUpdatedAt(t time.Time) *BusinessCustomerUpdateOne {
	bcuo.mutation.SetUpdatedAt(t)
	return bcuo
}

// SetName sets the "name" field.
func (bcuo *BusinessCustomerUpdateOne) SetName(s string) *BusinessCustomerUpdateOne {
	bcuo.mutation.SetName(s)
	return bcuo
}

// SetPhone sets the "phone" field.
func (bcuo *BusinessCustomerUpdateOne) SetPhone(s string) *BusinessCustomerUpdateOne {
	bcuo.mutation.SetPhone(s)
	return bcuo
}

// SetOtherPhone sets the "other_phone" field.
func (bcuo *BusinessCustomerUpdateOne) SetOtherPhone(s string) *BusinessCustomerUpdateOne {
	bcuo.mutation.SetOtherPhone(s)
	return bcuo
}

// SetNillableOtherPhone sets the "other_phone" field if the given value is not nil.
func (bcuo *BusinessCustomerUpdateOne) SetNillableOtherPhone(s *string) *BusinessCustomerUpdateOne {
	if s != nil {
		bcuo.SetOtherPhone(*s)
	}
	return bcuo
}

// ClearOtherPhone clears the value of the "other_phone" field.
func (bcuo *BusinessCustomerUpdateOne) ClearOtherPhone() *BusinessCustomerUpdateOne {
	bcuo.mutation.ClearOtherPhone()
	return bcuo
}

// SetLogo sets the "logo" field.
func (bcuo *BusinessCustomerUpdateOne) SetLogo(s string) *BusinessCustomerUpdateOne {
	bcuo.mutation.SetLogo(s)
	return bcuo
}

// SetNillableLogo sets the "logo" field if the given value is not nil.
func (bcuo *BusinessCustomerUpdateOne) SetNillableLogo(s *string) *BusinessCustomerUpdateOne {
	if s != nil {
		bcuo.SetLogo(*s)
	}
	return bcuo
}

// ClearLogo clears the value of the "logo" field.
func (bcuo *BusinessCustomerUpdateOne) ClearLogo() *BusinessCustomerUpdateOne {
	bcuo.mutation.ClearLogo()
	return bcuo
}

// SetContact sets the "contact" field.
func (bcuo *BusinessCustomerUpdateOne) SetContact(mcc *models.BusinessCustomerContact) *BusinessCustomerUpdateOne {
	bcuo.mutation.SetContact(mcc)
	return bcuo
}

// ClearContact clears the value of the "contact" field.
func (bcuo *BusinessCustomerUpdateOne) ClearContact() *BusinessCustomerUpdateOne {
	bcuo.mutation.ClearContact()
	return bcuo
}

// SetCustomerID sets the "customer" edge to the Customer entity by ID.
func (bcuo *BusinessCustomerUpdateOne) SetCustomerID(id int) *BusinessCustomerUpdateOne {
	bcuo.mutation.SetCustomerID(id)
	return bcuo
}

// SetCustomer sets the "customer" edge to the Customer entity.
func (bcuo *BusinessCustomerUpdateOne) SetCustomer(c *Customer) *BusinessCustomerUpdateOne {
	return bcuo.SetCustomerID(c.ID)
}

// Mutation returns the BusinessCustomerMutation object of the builder.
func (bcuo *BusinessCustomerUpdateOne) Mutation() *BusinessCustomerMutation {
	return bcuo.mutation
}

// ClearCustomer clears the "customer" edge to the Customer entity.
func (bcuo *BusinessCustomerUpdateOne) ClearCustomer() *BusinessCustomerUpdateOne {
	bcuo.mutation.ClearCustomer()
	return bcuo
}

// Where appends a list predicates to the BusinessCustomerUpdate builder.
func (bcuo *BusinessCustomerUpdateOne) Where(ps ...predicate.BusinessCustomer) *BusinessCustomerUpdateOne {
	bcuo.mutation.Where(ps...)
	return bcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bcuo *BusinessCustomerUpdateOne) Select(field string, fields ...string) *BusinessCustomerUpdateOne {
	bcuo.fields = append([]string{field}, fields...)
	return bcuo
}

// Save executes the query and returns the updated BusinessCustomer entity.
func (bcuo *BusinessCustomerUpdateOne) Save(ctx context.Context) (*BusinessCustomer, error) {
	bcuo.defaults()
	return withHooks(ctx, bcuo.sqlSave, bcuo.mutation, bcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bcuo *BusinessCustomerUpdateOne) SaveX(ctx context.Context) *BusinessCustomer {
	node, err := bcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bcuo *BusinessCustomerUpdateOne) Exec(ctx context.Context) error {
	_, err := bcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcuo *BusinessCustomerUpdateOne) ExecX(ctx context.Context) {
	if err := bcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bcuo *BusinessCustomerUpdateOne) defaults() {
	if _, ok := bcuo.mutation.UpdatedAt(); !ok {
		v := businesscustomer.UpdateDefaultUpdatedAt()
		bcuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bcuo *BusinessCustomerUpdateOne) check() error {
	if v, ok := bcuo.mutation.Name(); ok {
		if err := businesscustomer.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "BusinessCustomer.name": %w`, err)}
		}
	}
	if v, ok := bcuo.mutation.Phone(); ok {
		if err := businesscustomer.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "BusinessCustomer.phone": %w`, err)}
		}
	}
	if _, ok := bcuo.mutation.CustomerID(); bcuo.mutation.CustomerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "BusinessCustomer.customer"`)
	}
	return nil
}

func (bcuo *BusinessCustomerUpdateOne) sqlSave(ctx context.Context) (_node *BusinessCustomer, err error) {
	if err := bcuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(businesscustomer.Table, businesscustomer.Columns, sqlgraph.NewFieldSpec(businesscustomer.FieldID, field.TypeInt))
	id, ok := bcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "BusinessCustomer.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, businesscustomer.FieldID)
		for _, f := range fields {
			if !businesscustomer.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != businesscustomer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bcuo.mutation.UpdatedAt(); ok {
		_spec.SetField(businesscustomer.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := bcuo.mutation.Name(); ok {
		_spec.SetField(businesscustomer.FieldName, field.TypeString, value)
	}
	if value, ok := bcuo.mutation.Phone(); ok {
		_spec.SetField(businesscustomer.FieldPhone, field.TypeString, value)
	}
	if value, ok := bcuo.mutation.OtherPhone(); ok {
		_spec.SetField(businesscustomer.FieldOtherPhone, field.TypeString, value)
	}
	if bcuo.mutation.OtherPhoneCleared() {
		_spec.ClearField(businesscustomer.FieldOtherPhone, field.TypeString)
	}
	if value, ok := bcuo.mutation.Logo(); ok {
		_spec.SetField(businesscustomer.FieldLogo, field.TypeString, value)
	}
	if bcuo.mutation.LogoCleared() {
		_spec.ClearField(businesscustomer.FieldLogo, field.TypeString)
	}
	if value, ok := bcuo.mutation.Contact(); ok {
		_spec.SetField(businesscustomer.FieldContact, field.TypeJSON, value)
	}
	if bcuo.mutation.ContactCleared() {
		_spec.ClearField(businesscustomer.FieldContact, field.TypeJSON)
	}
	if bcuo.mutation.CustomerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   businesscustomer.CustomerTable,
			Columns: []string{businesscustomer.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcuo.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   businesscustomer.CustomerTable,
			Columns: []string{businesscustomer.CustomerColumn},
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
	_node = &BusinessCustomer{config: bcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{businesscustomer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	bcuo.mutation.done = true
	return _node, nil
}
