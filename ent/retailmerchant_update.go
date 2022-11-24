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
	"github.com/SeyramWood/ent/merchant"
	"github.com/SeyramWood/ent/predicate"
	"github.com/SeyramWood/ent/retailmerchant"
)

// RetailMerchantUpdate is the builder for updating RetailMerchant entities.
type RetailMerchantUpdate struct {
	config
	hooks    []Hook
	mutation *RetailMerchantMutation
}

// Where appends a list predicates to the RetailMerchantUpdate builder.
func (rmu *RetailMerchantUpdate) Where(ps ...predicate.RetailMerchant) *RetailMerchantUpdate {
	rmu.mutation.Where(ps...)
	return rmu
}

// SetUpdatedAt sets the "updated_at" field.
func (rmu *RetailMerchantUpdate) SetUpdatedAt(t time.Time) *RetailMerchantUpdate {
	rmu.mutation.SetUpdatedAt(t)
	return rmu
}

// SetGhanaCard sets the "ghana_card" field.
func (rmu *RetailMerchantUpdate) SetGhanaCard(s string) *RetailMerchantUpdate {
	rmu.mutation.SetGhanaCard(s)
	return rmu
}

// SetLastName sets the "last_name" field.
func (rmu *RetailMerchantUpdate) SetLastName(s string) *RetailMerchantUpdate {
	rmu.mutation.SetLastName(s)
	return rmu
}

// SetOtherName sets the "other_name" field.
func (rmu *RetailMerchantUpdate) SetOtherName(s string) *RetailMerchantUpdate {
	rmu.mutation.SetOtherName(s)
	return rmu
}

// SetPhone sets the "phone" field.
func (rmu *RetailMerchantUpdate) SetPhone(s string) *RetailMerchantUpdate {
	rmu.mutation.SetPhone(s)
	return rmu
}

// SetOtherPhone sets the "other_phone" field.
func (rmu *RetailMerchantUpdate) SetOtherPhone(s string) *RetailMerchantUpdate {
	rmu.mutation.SetOtherPhone(s)
	return rmu
}

// SetNillableOtherPhone sets the "other_phone" field if the given value is not nil.
func (rmu *RetailMerchantUpdate) SetNillableOtherPhone(s *string) *RetailMerchantUpdate {
	if s != nil {
		rmu.SetOtherPhone(*s)
	}
	return rmu
}

// ClearOtherPhone clears the value of the "other_phone" field.
func (rmu *RetailMerchantUpdate) ClearOtherPhone() *RetailMerchantUpdate {
	rmu.mutation.ClearOtherPhone()
	return rmu
}

// SetMerchantID sets the "merchant" edge to the Merchant entity by ID.
func (rmu *RetailMerchantUpdate) SetMerchantID(id int) *RetailMerchantUpdate {
	rmu.mutation.SetMerchantID(id)
	return rmu
}

// SetMerchant sets the "merchant" edge to the Merchant entity.
func (rmu *RetailMerchantUpdate) SetMerchant(m *Merchant) *RetailMerchantUpdate {
	return rmu.SetMerchantID(m.ID)
}

// Mutation returns the RetailMerchantMutation object of the builder.
func (rmu *RetailMerchantUpdate) Mutation() *RetailMerchantMutation {
	return rmu.mutation
}

// ClearMerchant clears the "merchant" edge to the Merchant entity.
func (rmu *RetailMerchantUpdate) ClearMerchant() *RetailMerchantUpdate {
	rmu.mutation.ClearMerchant()
	return rmu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rmu *RetailMerchantUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	rmu.defaults()
	if len(rmu.hooks) == 0 {
		if err = rmu.check(); err != nil {
			return 0, err
		}
		affected, err = rmu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RetailMerchantMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rmu.check(); err != nil {
				return 0, err
			}
			rmu.mutation = mutation
			affected, err = rmu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rmu.hooks) - 1; i >= 0; i-- {
			if rmu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rmu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rmu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (rmu *RetailMerchantUpdate) SaveX(ctx context.Context) int {
	affected, err := rmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rmu *RetailMerchantUpdate) Exec(ctx context.Context) error {
	_, err := rmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rmu *RetailMerchantUpdate) ExecX(ctx context.Context) {
	if err := rmu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rmu *RetailMerchantUpdate) defaults() {
	if _, ok := rmu.mutation.UpdatedAt(); !ok {
		v := retailmerchant.UpdateDefaultUpdatedAt()
		rmu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rmu *RetailMerchantUpdate) check() error {
	if v, ok := rmu.mutation.GhanaCard(); ok {
		if err := retailmerchant.GhanaCardValidator(v); err != nil {
			return &ValidationError{Name: "ghana_card", err: fmt.Errorf(`ent: validator failed for field "RetailMerchant.ghana_card": %w`, err)}
		}
	}
	if v, ok := rmu.mutation.LastName(); ok {
		if err := retailmerchant.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "RetailMerchant.last_name": %w`, err)}
		}
	}
	if v, ok := rmu.mutation.OtherName(); ok {
		if err := retailmerchant.OtherNameValidator(v); err != nil {
			return &ValidationError{Name: "other_name", err: fmt.Errorf(`ent: validator failed for field "RetailMerchant.other_name": %w`, err)}
		}
	}
	if v, ok := rmu.mutation.Phone(); ok {
		if err := retailmerchant.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "RetailMerchant.phone": %w`, err)}
		}
	}
	if _, ok := rmu.mutation.MerchantID(); rmu.mutation.MerchantCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "RetailMerchant.merchant"`)
	}
	return nil
}

func (rmu *RetailMerchantUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   retailmerchant.Table,
			Columns: retailmerchant.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: retailmerchant.FieldID,
			},
		},
	}
	if ps := rmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rmu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: retailmerchant.FieldUpdatedAt,
		})
	}
	if value, ok := rmu.mutation.GhanaCard(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: retailmerchant.FieldGhanaCard,
		})
	}
	if value, ok := rmu.mutation.LastName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: retailmerchant.FieldLastName,
		})
	}
	if value, ok := rmu.mutation.OtherName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: retailmerchant.FieldOtherName,
		})
	}
	if value, ok := rmu.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: retailmerchant.FieldPhone,
		})
	}
	if value, ok := rmu.mutation.OtherPhone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: retailmerchant.FieldOtherPhone,
		})
	}
	if rmu.mutation.OtherPhoneCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: retailmerchant.FieldOtherPhone,
		})
	}
	if rmu.mutation.MerchantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   retailmerchant.MerchantTable,
			Columns: []string{retailmerchant.MerchantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: merchant.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rmu.mutation.MerchantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   retailmerchant.MerchantTable,
			Columns: []string{retailmerchant.MerchantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: merchant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{retailmerchant.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// RetailMerchantUpdateOne is the builder for updating a single RetailMerchant entity.
type RetailMerchantUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RetailMerchantMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (rmuo *RetailMerchantUpdateOne) SetUpdatedAt(t time.Time) *RetailMerchantUpdateOne {
	rmuo.mutation.SetUpdatedAt(t)
	return rmuo
}

// SetGhanaCard sets the "ghana_card" field.
func (rmuo *RetailMerchantUpdateOne) SetGhanaCard(s string) *RetailMerchantUpdateOne {
	rmuo.mutation.SetGhanaCard(s)
	return rmuo
}

// SetLastName sets the "last_name" field.
func (rmuo *RetailMerchantUpdateOne) SetLastName(s string) *RetailMerchantUpdateOne {
	rmuo.mutation.SetLastName(s)
	return rmuo
}

// SetOtherName sets the "other_name" field.
func (rmuo *RetailMerchantUpdateOne) SetOtherName(s string) *RetailMerchantUpdateOne {
	rmuo.mutation.SetOtherName(s)
	return rmuo
}

// SetPhone sets the "phone" field.
func (rmuo *RetailMerchantUpdateOne) SetPhone(s string) *RetailMerchantUpdateOne {
	rmuo.mutation.SetPhone(s)
	return rmuo
}

// SetOtherPhone sets the "other_phone" field.
func (rmuo *RetailMerchantUpdateOne) SetOtherPhone(s string) *RetailMerchantUpdateOne {
	rmuo.mutation.SetOtherPhone(s)
	return rmuo
}

// SetNillableOtherPhone sets the "other_phone" field if the given value is not nil.
func (rmuo *RetailMerchantUpdateOne) SetNillableOtherPhone(s *string) *RetailMerchantUpdateOne {
	if s != nil {
		rmuo.SetOtherPhone(*s)
	}
	return rmuo
}

// ClearOtherPhone clears the value of the "other_phone" field.
func (rmuo *RetailMerchantUpdateOne) ClearOtherPhone() *RetailMerchantUpdateOne {
	rmuo.mutation.ClearOtherPhone()
	return rmuo
}

// SetMerchantID sets the "merchant" edge to the Merchant entity by ID.
func (rmuo *RetailMerchantUpdateOne) SetMerchantID(id int) *RetailMerchantUpdateOne {
	rmuo.mutation.SetMerchantID(id)
	return rmuo
}

// SetMerchant sets the "merchant" edge to the Merchant entity.
func (rmuo *RetailMerchantUpdateOne) SetMerchant(m *Merchant) *RetailMerchantUpdateOne {
	return rmuo.SetMerchantID(m.ID)
}

// Mutation returns the RetailMerchantMutation object of the builder.
func (rmuo *RetailMerchantUpdateOne) Mutation() *RetailMerchantMutation {
	return rmuo.mutation
}

// ClearMerchant clears the "merchant" edge to the Merchant entity.
func (rmuo *RetailMerchantUpdateOne) ClearMerchant() *RetailMerchantUpdateOne {
	rmuo.mutation.ClearMerchant()
	return rmuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rmuo *RetailMerchantUpdateOne) Select(field string, fields ...string) *RetailMerchantUpdateOne {
	rmuo.fields = append([]string{field}, fields...)
	return rmuo
}

// Save executes the query and returns the updated RetailMerchant entity.
func (rmuo *RetailMerchantUpdateOne) Save(ctx context.Context) (*RetailMerchant, error) {
	var (
		err  error
		node *RetailMerchant
	)
	rmuo.defaults()
	if len(rmuo.hooks) == 0 {
		if err = rmuo.check(); err != nil {
			return nil, err
		}
		node, err = rmuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RetailMerchantMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rmuo.check(); err != nil {
				return nil, err
			}
			rmuo.mutation = mutation
			node, err = rmuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rmuo.hooks) - 1; i >= 0; i-- {
			if rmuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rmuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, rmuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*RetailMerchant)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from RetailMerchantMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (rmuo *RetailMerchantUpdateOne) SaveX(ctx context.Context) *RetailMerchant {
	node, err := rmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rmuo *RetailMerchantUpdateOne) Exec(ctx context.Context) error {
	_, err := rmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rmuo *RetailMerchantUpdateOne) ExecX(ctx context.Context) {
	if err := rmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rmuo *RetailMerchantUpdateOne) defaults() {
	if _, ok := rmuo.mutation.UpdatedAt(); !ok {
		v := retailmerchant.UpdateDefaultUpdatedAt()
		rmuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rmuo *RetailMerchantUpdateOne) check() error {
	if v, ok := rmuo.mutation.GhanaCard(); ok {
		if err := retailmerchant.GhanaCardValidator(v); err != nil {
			return &ValidationError{Name: "ghana_card", err: fmt.Errorf(`ent: validator failed for field "RetailMerchant.ghana_card": %w`, err)}
		}
	}
	if v, ok := rmuo.mutation.LastName(); ok {
		if err := retailmerchant.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "RetailMerchant.last_name": %w`, err)}
		}
	}
	if v, ok := rmuo.mutation.OtherName(); ok {
		if err := retailmerchant.OtherNameValidator(v); err != nil {
			return &ValidationError{Name: "other_name", err: fmt.Errorf(`ent: validator failed for field "RetailMerchant.other_name": %w`, err)}
		}
	}
	if v, ok := rmuo.mutation.Phone(); ok {
		if err := retailmerchant.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "RetailMerchant.phone": %w`, err)}
		}
	}
	if _, ok := rmuo.mutation.MerchantID(); rmuo.mutation.MerchantCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "RetailMerchant.merchant"`)
	}
	return nil
}

func (rmuo *RetailMerchantUpdateOne) sqlSave(ctx context.Context) (_node *RetailMerchant, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   retailmerchant.Table,
			Columns: retailmerchant.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: retailmerchant.FieldID,
			},
		},
	}
	id, ok := rmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "RetailMerchant.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, retailmerchant.FieldID)
		for _, f := range fields {
			if !retailmerchant.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != retailmerchant.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rmuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: retailmerchant.FieldUpdatedAt,
		})
	}
	if value, ok := rmuo.mutation.GhanaCard(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: retailmerchant.FieldGhanaCard,
		})
	}
	if value, ok := rmuo.mutation.LastName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: retailmerchant.FieldLastName,
		})
	}
	if value, ok := rmuo.mutation.OtherName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: retailmerchant.FieldOtherName,
		})
	}
	if value, ok := rmuo.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: retailmerchant.FieldPhone,
		})
	}
	if value, ok := rmuo.mutation.OtherPhone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: retailmerchant.FieldOtherPhone,
		})
	}
	if rmuo.mutation.OtherPhoneCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: retailmerchant.FieldOtherPhone,
		})
	}
	if rmuo.mutation.MerchantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   retailmerchant.MerchantTable,
			Columns: []string{retailmerchant.MerchantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: merchant.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rmuo.mutation.MerchantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   retailmerchant.MerchantTable,
			Columns: []string{retailmerchant.MerchantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: merchant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &RetailMerchant{config: rmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{retailmerchant.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
