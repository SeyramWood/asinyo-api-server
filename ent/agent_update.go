// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SeyramWood/ent/agent"
	"github.com/SeyramWood/ent/predicate"
)

// AgentUpdate is the builder for updating Agent entities.
type AgentUpdate struct {
	config
	hooks    []Hook
	mutation *AgentMutation
}

// Where appends a list predicates to the AgentUpdate builder.
func (au *AgentUpdate) Where(ps ...predicate.Agent) *AgentUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AgentUpdate) SetUpdatedAt(t time.Time) *AgentUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetUsername sets the "username" field.
func (au *AgentUpdate) SetUsername(s string) *AgentUpdate {
	au.mutation.SetUsername(s)
	return au
}

// SetPassword sets the "password" field.
func (au *AgentUpdate) SetPassword(b []byte) *AgentUpdate {
	au.mutation.SetPassword(b)
	return au
}

// SetGhanaCard sets the "ghana_card" field.
func (au *AgentUpdate) SetGhanaCard(s string) *AgentUpdate {
	au.mutation.SetGhanaCard(s)
	return au
}

// SetLastName sets the "last_name" field.
func (au *AgentUpdate) SetLastName(s string) *AgentUpdate {
	au.mutation.SetLastName(s)
	return au
}

// SetOtherName sets the "other_name" field.
func (au *AgentUpdate) SetOtherName(s string) *AgentUpdate {
	au.mutation.SetOtherName(s)
	return au
}

// SetPhone sets the "phone" field.
func (au *AgentUpdate) SetPhone(s string) *AgentUpdate {
	au.mutation.SetPhone(s)
	return au
}

// SetOtherPhone sets the "other_phone" field.
func (au *AgentUpdate) SetOtherPhone(s string) *AgentUpdate {
	au.mutation.SetOtherPhone(s)
	return au
}

// SetNillableOtherPhone sets the "other_phone" field if the given value is not nil.
func (au *AgentUpdate) SetNillableOtherPhone(s *string) *AgentUpdate {
	if s != nil {
		au.SetOtherPhone(*s)
	}
	return au
}

// ClearOtherPhone clears the value of the "other_phone" field.
func (au *AgentUpdate) ClearOtherPhone() *AgentUpdate {
	au.mutation.ClearOtherPhone()
	return au
}

// SetAddress sets the "address" field.
func (au *AgentUpdate) SetAddress(s string) *AgentUpdate {
	au.mutation.SetAddress(s)
	return au
}

// SetDigitalAddress sets the "digital_address" field.
func (au *AgentUpdate) SetDigitalAddress(s string) *AgentUpdate {
	au.mutation.SetDigitalAddress(s)
	return au
}

// Mutation returns the AgentMutation object of the builder.
func (au *AgentUpdate) Mutation() *AgentMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AgentUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	au.defaults()
	if len(au.hooks) == 0 {
		if err = au.check(); err != nil {
			return 0, err
		}
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AgentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = au.check(); err != nil {
				return 0, err
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			if au.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AgentUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AgentUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AgentUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AgentUpdate) defaults() {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		v := agent.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AgentUpdate) check() error {
	if v, ok := au.mutation.Username(); ok {
		if err := agent.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "Agent.username": %w`, err)}
		}
	}
	if v, ok := au.mutation.Password(); ok {
		if err := agent.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Agent.password": %w`, err)}
		}
	}
	if v, ok := au.mutation.GhanaCard(); ok {
		if err := agent.GhanaCardValidator(v); err != nil {
			return &ValidationError{Name: "ghana_card", err: fmt.Errorf(`ent: validator failed for field "Agent.ghana_card": %w`, err)}
		}
	}
	if v, ok := au.mutation.LastName(); ok {
		if err := agent.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "Agent.last_name": %w`, err)}
		}
	}
	if v, ok := au.mutation.OtherName(); ok {
		if err := agent.OtherNameValidator(v); err != nil {
			return &ValidationError{Name: "other_name", err: fmt.Errorf(`ent: validator failed for field "Agent.other_name": %w`, err)}
		}
	}
	if v, ok := au.mutation.Phone(); ok {
		if err := agent.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "Agent.phone": %w`, err)}
		}
	}
	if v, ok := au.mutation.Address(); ok {
		if err := agent.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`ent: validator failed for field "Agent.address": %w`, err)}
		}
	}
	if v, ok := au.mutation.DigitalAddress(); ok {
		if err := agent.DigitalAddressValidator(v); err != nil {
			return &ValidationError{Name: "digital_address", err: fmt.Errorf(`ent: validator failed for field "Agent.digital_address": %w`, err)}
		}
	}
	return nil
}

func (au *AgentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   agent.Table,
			Columns: agent.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: agent.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: agent.FieldUpdatedAt,
		})
	}
	if value, ok := au.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agent.FieldUsername,
		})
	}
	if value, ok := au.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: agent.FieldPassword,
		})
	}
	if value, ok := au.mutation.GhanaCard(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agent.FieldGhanaCard,
		})
	}
	if value, ok := au.mutation.LastName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agent.FieldLastName,
		})
	}
	if value, ok := au.mutation.OtherName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agent.FieldOtherName,
		})
	}
	if value, ok := au.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agent.FieldPhone,
		})
	}
	if value, ok := au.mutation.OtherPhone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agent.FieldOtherPhone,
		})
	}
	if au.mutation.OtherPhoneCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: agent.FieldOtherPhone,
		})
	}
	if value, ok := au.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agent.FieldAddress,
		})
	}
	if value, ok := au.mutation.DigitalAddress(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agent.FieldDigitalAddress,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{agent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AgentUpdateOne is the builder for updating a single Agent entity.
type AgentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AgentMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AgentUpdateOne) SetUpdatedAt(t time.Time) *AgentUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetUsername sets the "username" field.
func (auo *AgentUpdateOne) SetUsername(s string) *AgentUpdateOne {
	auo.mutation.SetUsername(s)
	return auo
}

// SetPassword sets the "password" field.
func (auo *AgentUpdateOne) SetPassword(b []byte) *AgentUpdateOne {
	auo.mutation.SetPassword(b)
	return auo
}

// SetGhanaCard sets the "ghana_card" field.
func (auo *AgentUpdateOne) SetGhanaCard(s string) *AgentUpdateOne {
	auo.mutation.SetGhanaCard(s)
	return auo
}

// SetLastName sets the "last_name" field.
func (auo *AgentUpdateOne) SetLastName(s string) *AgentUpdateOne {
	auo.mutation.SetLastName(s)
	return auo
}

// SetOtherName sets the "other_name" field.
func (auo *AgentUpdateOne) SetOtherName(s string) *AgentUpdateOne {
	auo.mutation.SetOtherName(s)
	return auo
}

// SetPhone sets the "phone" field.
func (auo *AgentUpdateOne) SetPhone(s string) *AgentUpdateOne {
	auo.mutation.SetPhone(s)
	return auo
}

// SetOtherPhone sets the "other_phone" field.
func (auo *AgentUpdateOne) SetOtherPhone(s string) *AgentUpdateOne {
	auo.mutation.SetOtherPhone(s)
	return auo
}

// SetNillableOtherPhone sets the "other_phone" field if the given value is not nil.
func (auo *AgentUpdateOne) SetNillableOtherPhone(s *string) *AgentUpdateOne {
	if s != nil {
		auo.SetOtherPhone(*s)
	}
	return auo
}

// ClearOtherPhone clears the value of the "other_phone" field.
func (auo *AgentUpdateOne) ClearOtherPhone() *AgentUpdateOne {
	auo.mutation.ClearOtherPhone()
	return auo
}

// SetAddress sets the "address" field.
func (auo *AgentUpdateOne) SetAddress(s string) *AgentUpdateOne {
	auo.mutation.SetAddress(s)
	return auo
}

// SetDigitalAddress sets the "digital_address" field.
func (auo *AgentUpdateOne) SetDigitalAddress(s string) *AgentUpdateOne {
	auo.mutation.SetDigitalAddress(s)
	return auo
}

// Mutation returns the AgentMutation object of the builder.
func (auo *AgentUpdateOne) Mutation() *AgentMutation {
	return auo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AgentUpdateOne) Select(field string, fields ...string) *AgentUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Agent entity.
func (auo *AgentUpdateOne) Save(ctx context.Context) (*Agent, error) {
	var (
		err  error
		node *Agent
	)
	auo.defaults()
	if len(auo.hooks) == 0 {
		if err = auo.check(); err != nil {
			return nil, err
		}
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AgentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = auo.check(); err != nil {
				return nil, err
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			if auo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AgentUpdateOne) SaveX(ctx context.Context) *Agent {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AgentUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AgentUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AgentUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		v := agent.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AgentUpdateOne) check() error {
	if v, ok := auo.mutation.Username(); ok {
		if err := agent.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "Agent.username": %w`, err)}
		}
	}
	if v, ok := auo.mutation.Password(); ok {
		if err := agent.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Agent.password": %w`, err)}
		}
	}
	if v, ok := auo.mutation.GhanaCard(); ok {
		if err := agent.GhanaCardValidator(v); err != nil {
			return &ValidationError{Name: "ghana_card", err: fmt.Errorf(`ent: validator failed for field "Agent.ghana_card": %w`, err)}
		}
	}
	if v, ok := auo.mutation.LastName(); ok {
		if err := agent.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "Agent.last_name": %w`, err)}
		}
	}
	if v, ok := auo.mutation.OtherName(); ok {
		if err := agent.OtherNameValidator(v); err != nil {
			return &ValidationError{Name: "other_name", err: fmt.Errorf(`ent: validator failed for field "Agent.other_name": %w`, err)}
		}
	}
	if v, ok := auo.mutation.Phone(); ok {
		if err := agent.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "Agent.phone": %w`, err)}
		}
	}
	if v, ok := auo.mutation.Address(); ok {
		if err := agent.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`ent: validator failed for field "Agent.address": %w`, err)}
		}
	}
	if v, ok := auo.mutation.DigitalAddress(); ok {
		if err := agent.DigitalAddressValidator(v); err != nil {
			return &ValidationError{Name: "digital_address", err: fmt.Errorf(`ent: validator failed for field "Agent.digital_address": %w`, err)}
		}
	}
	return nil
}

func (auo *AgentUpdateOne) sqlSave(ctx context.Context) (_node *Agent, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   agent.Table,
			Columns: agent.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: agent.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Agent.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, agent.FieldID)
		for _, f := range fields {
			if !agent.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != agent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: agent.FieldUpdatedAt,
		})
	}
	if value, ok := auo.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agent.FieldUsername,
		})
	}
	if value, ok := auo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: agent.FieldPassword,
		})
	}
	if value, ok := auo.mutation.GhanaCard(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agent.FieldGhanaCard,
		})
	}
	if value, ok := auo.mutation.LastName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agent.FieldLastName,
		})
	}
	if value, ok := auo.mutation.OtherName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agent.FieldOtherName,
		})
	}
	if value, ok := auo.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agent.FieldPhone,
		})
	}
	if value, ok := auo.mutation.OtherPhone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agent.FieldOtherPhone,
		})
	}
	if auo.mutation.OtherPhoneCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: agent.FieldOtherPhone,
		})
	}
	if value, ok := auo.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agent.FieldAddress,
		})
	}
	if value, ok := auo.mutation.DigitalAddress(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agent.FieldDigitalAddress,
		})
	}
	_node = &Agent{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{agent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
