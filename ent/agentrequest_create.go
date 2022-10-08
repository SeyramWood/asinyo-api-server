// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SeyramWood/ent/agentrequest"
)

// AgentRequestCreate is the builder for creating a AgentRequest entity.
type AgentRequestCreate struct {
	config
	mutation *AgentRequestMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (arc *AgentRequestCreate) SetCreatedAt(t time.Time) *AgentRequestCreate {
	arc.mutation.SetCreatedAt(t)
	return arc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (arc *AgentRequestCreate) SetNillableCreatedAt(t *time.Time) *AgentRequestCreate {
	if t != nil {
		arc.SetCreatedAt(*t)
	}
	return arc
}

// SetUpdatedAt sets the "updated_at" field.
func (arc *AgentRequestCreate) SetUpdatedAt(t time.Time) *AgentRequestCreate {
	arc.mutation.SetUpdatedAt(t)
	return arc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (arc *AgentRequestCreate) SetNillableUpdatedAt(t *time.Time) *AgentRequestCreate {
	if t != nil {
		arc.SetUpdatedAt(*t)
	}
	return arc
}

// Mutation returns the AgentRequestMutation object of the builder.
func (arc *AgentRequestCreate) Mutation() *AgentRequestMutation {
	return arc.mutation
}

// Save creates the AgentRequest in the database.
func (arc *AgentRequestCreate) Save(ctx context.Context) (*AgentRequest, error) {
	var (
		err  error
		node *AgentRequest
	)
	arc.defaults()
	if len(arc.hooks) == 0 {
		if err = arc.check(); err != nil {
			return nil, err
		}
		node, err = arc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AgentRequestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = arc.check(); err != nil {
				return nil, err
			}
			arc.mutation = mutation
			if node, err = arc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(arc.hooks) - 1; i >= 0; i-- {
			if arc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = arc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, arc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (arc *AgentRequestCreate) SaveX(ctx context.Context) *AgentRequest {
	v, err := arc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (arc *AgentRequestCreate) Exec(ctx context.Context) error {
	_, err := arc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (arc *AgentRequestCreate) ExecX(ctx context.Context) {
	if err := arc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (arc *AgentRequestCreate) defaults() {
	if _, ok := arc.mutation.CreatedAt(); !ok {
		v := agentrequest.DefaultCreatedAt()
		arc.mutation.SetCreatedAt(v)
	}
	if _, ok := arc.mutation.UpdatedAt(); !ok {
		v := agentrequest.DefaultUpdatedAt()
		arc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (arc *AgentRequestCreate) check() error {
	if _, ok := arc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AgentRequest.created_at"`)}
	}
	if _, ok := arc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "AgentRequest.updated_at"`)}
	}
	return nil
}

func (arc *AgentRequestCreate) sqlSave(ctx context.Context) (*AgentRequest, error) {
	_node, _spec := arc.createSpec()
	if err := sqlgraph.CreateNode(ctx, arc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (arc *AgentRequestCreate) createSpec() (*AgentRequest, *sqlgraph.CreateSpec) {
	var (
		_node = &AgentRequest{config: arc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: agentrequest.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: agentrequest.FieldID,
			},
		}
	)
	if value, ok := arc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: agentrequest.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := arc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: agentrequest.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// AgentRequestCreateBulk is the builder for creating many AgentRequest entities in bulk.
type AgentRequestCreateBulk struct {
	config
	builders []*AgentRequestCreate
}

// Save creates the AgentRequest entities in the database.
func (arcb *AgentRequestCreateBulk) Save(ctx context.Context) ([]*AgentRequest, error) {
	specs := make([]*sqlgraph.CreateSpec, len(arcb.builders))
	nodes := make([]*AgentRequest, len(arcb.builders))
	mutators := make([]Mutator, len(arcb.builders))
	for i := range arcb.builders {
		func(i int, root context.Context) {
			builder := arcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AgentRequestMutation)
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
					_, err = mutators[i+1].Mutate(root, arcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, arcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, arcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (arcb *AgentRequestCreateBulk) SaveX(ctx context.Context) []*AgentRequest {
	v, err := arcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (arcb *AgentRequestCreateBulk) Exec(ctx context.Context) error {
	_, err := arcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (arcb *AgentRequestCreateBulk) ExecX(ctx context.Context) {
	if err := arcb.Exec(ctx); err != nil {
		panic(err)
	}
}
