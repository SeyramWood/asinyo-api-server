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
	"github.com/SeyramWood/ent/agentrequest"
	"github.com/SeyramWood/ent/predicate"
)

// AgentRequestUpdate is the builder for updating AgentRequest entities.
type AgentRequestUpdate struct {
	config
	hooks    []Hook
	mutation *AgentRequestMutation
}

// Where appends a list predicates to the AgentRequestUpdate builder.
func (aru *AgentRequestUpdate) Where(ps ...predicate.AgentRequest) *AgentRequestUpdate {
	aru.mutation.Where(ps...)
	return aru
}

// SetUpdatedAt sets the "updated_at" field.
func (aru *AgentRequestUpdate) SetUpdatedAt(t time.Time) *AgentRequestUpdate {
	aru.mutation.SetUpdatedAt(t)
	return aru
}

// Mutation returns the AgentRequestMutation object of the builder.
func (aru *AgentRequestUpdate) Mutation() *AgentRequestMutation {
	return aru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aru *AgentRequestUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	aru.defaults()
	if len(aru.hooks) == 0 {
		affected, err = aru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AgentRequestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aru.mutation = mutation
			affected, err = aru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(aru.hooks) - 1; i >= 0; i-- {
			if aru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (aru *AgentRequestUpdate) SaveX(ctx context.Context) int {
	affected, err := aru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aru *AgentRequestUpdate) Exec(ctx context.Context) error {
	_, err := aru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aru *AgentRequestUpdate) ExecX(ctx context.Context) {
	if err := aru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aru *AgentRequestUpdate) defaults() {
	if _, ok := aru.mutation.UpdatedAt(); !ok {
		v := agentrequest.UpdateDefaultUpdatedAt()
		aru.mutation.SetUpdatedAt(v)
	}
}

func (aru *AgentRequestUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   agentrequest.Table,
			Columns: agentrequest.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: agentrequest.FieldID,
			},
		},
	}
	if ps := aru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aru.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: agentrequest.FieldUpdatedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, aru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{agentrequest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// AgentRequestUpdateOne is the builder for updating a single AgentRequest entity.
type AgentRequestUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AgentRequestMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (aruo *AgentRequestUpdateOne) SetUpdatedAt(t time.Time) *AgentRequestUpdateOne {
	aruo.mutation.SetUpdatedAt(t)
	return aruo
}

// Mutation returns the AgentRequestMutation object of the builder.
func (aruo *AgentRequestUpdateOne) Mutation() *AgentRequestMutation {
	return aruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aruo *AgentRequestUpdateOne) Select(field string, fields ...string) *AgentRequestUpdateOne {
	aruo.fields = append([]string{field}, fields...)
	return aruo
}

// Save executes the query and returns the updated AgentRequest entity.
func (aruo *AgentRequestUpdateOne) Save(ctx context.Context) (*AgentRequest, error) {
	var (
		err  error
		node *AgentRequest
	)
	aruo.defaults()
	if len(aruo.hooks) == 0 {
		node, err = aruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AgentRequestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aruo.mutation = mutation
			node, err = aruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(aruo.hooks) - 1; i >= 0; i-- {
			if aruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aruo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, aruo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AgentRequest)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AgentRequestMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (aruo *AgentRequestUpdateOne) SaveX(ctx context.Context) *AgentRequest {
	node, err := aruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aruo *AgentRequestUpdateOne) Exec(ctx context.Context) error {
	_, err := aruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aruo *AgentRequestUpdateOne) ExecX(ctx context.Context) {
	if err := aruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aruo *AgentRequestUpdateOne) defaults() {
	if _, ok := aruo.mutation.UpdatedAt(); !ok {
		v := agentrequest.UpdateDefaultUpdatedAt()
		aruo.mutation.SetUpdatedAt(v)
	}
}

func (aruo *AgentRequestUpdateOne) sqlSave(ctx context.Context) (_node *AgentRequest, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   agentrequest.Table,
			Columns: agentrequest.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: agentrequest.FieldID,
			},
		},
	}
	id, ok := aruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AgentRequest.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, agentrequest.FieldID)
		for _, f := range fields {
			if !agentrequest.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != agentrequest.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aruo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: agentrequest.FieldUpdatedAt,
		})
	}
	_node = &AgentRequest{config: aruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{agentrequest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}