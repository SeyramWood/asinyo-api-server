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
	"github.com/SeyramWood/ent/logistic"
	"github.com/SeyramWood/ent/order"
	"github.com/SeyramWood/ent/predicate"
)

// LogisticUpdate is the builder for updating Logistic entities.
type LogisticUpdate struct {
	config
	hooks    []Hook
	mutation *LogisticMutation
}

// Where appends a list predicates to the LogisticUpdate builder.
func (lu *LogisticUpdate) Where(ps ...predicate.Logistic) *LogisticUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetUpdatedAt sets the "updated_at" field.
func (lu *LogisticUpdate) SetUpdatedAt(t time.Time) *LogisticUpdate {
	lu.mutation.SetUpdatedAt(t)
	return lu
}

// SetType sets the "type" field.
func (lu *LogisticUpdate) SetType(s string) *LogisticUpdate {
	lu.mutation.SetType(s)
	return lu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (lu *LogisticUpdate) SetNillableType(s *string) *LogisticUpdate {
	if s != nil {
		lu.SetType(*s)
	}
	return lu
}

// SetTask sets the "task" field.
func (lu *LogisticUpdate) SetTask(s *struct {
	Data interface{} "json:\"data\""
}) *LogisticUpdate {
	lu.mutation.SetTask(s)
	return lu
}

// ClearTask clears the value of the "task" field.
func (lu *LogisticUpdate) ClearTask() *LogisticUpdate {
	lu.mutation.ClearTask()
	return lu
}

// SetOrderID sets the "order" edge to the Order entity by ID.
func (lu *LogisticUpdate) SetOrderID(id int) *LogisticUpdate {
	lu.mutation.SetOrderID(id)
	return lu
}

// SetNillableOrderID sets the "order" edge to the Order entity by ID if the given value is not nil.
func (lu *LogisticUpdate) SetNillableOrderID(id *int) *LogisticUpdate {
	if id != nil {
		lu = lu.SetOrderID(*id)
	}
	return lu
}

// SetOrder sets the "order" edge to the Order entity.
func (lu *LogisticUpdate) SetOrder(o *Order) *LogisticUpdate {
	return lu.SetOrderID(o.ID)
}

// Mutation returns the LogisticMutation object of the builder.
func (lu *LogisticUpdate) Mutation() *LogisticMutation {
	return lu.mutation
}

// ClearOrder clears the "order" edge to the Order entity.
func (lu *LogisticUpdate) ClearOrder() *LogisticUpdate {
	lu.mutation.ClearOrder()
	return lu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LogisticUpdate) Save(ctx context.Context) (int, error) {
	lu.defaults()
	return withHooks(ctx, lu.sqlSave, lu.mutation, lu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LogisticUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LogisticUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LogisticUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lu *LogisticUpdate) defaults() {
	if _, ok := lu.mutation.UpdatedAt(); !ok {
		v := logistic.UpdateDefaultUpdatedAt()
		lu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lu *LogisticUpdate) check() error {
	if v, ok := lu.mutation.GetType(); ok {
		if err := logistic.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Logistic.type": %w`, err)}
		}
	}
	return nil
}

func (lu *LogisticUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := lu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(logistic.Table, logistic.Columns, sqlgraph.NewFieldSpec(logistic.FieldID, field.TypeInt))
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.UpdatedAt(); ok {
		_spec.SetField(logistic.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := lu.mutation.GetType(); ok {
		_spec.SetField(logistic.FieldType, field.TypeString, value)
	}
	if value, ok := lu.mutation.Task(); ok {
		_spec.SetField(logistic.FieldTask, field.TypeJSON, value)
	}
	if lu.mutation.TaskCleared() {
		_spec.ClearField(logistic.FieldTask, field.TypeJSON)
	}
	if lu.mutation.OrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   logistic.OrderTable,
			Columns: []string{logistic.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   logistic.OrderTable,
			Columns: []string{logistic.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{logistic.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lu.mutation.done = true
	return n, nil
}

// LogisticUpdateOne is the builder for updating a single Logistic entity.
type LogisticUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LogisticMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (luo *LogisticUpdateOne) SetUpdatedAt(t time.Time) *LogisticUpdateOne {
	luo.mutation.SetUpdatedAt(t)
	return luo
}

// SetType sets the "type" field.
func (luo *LogisticUpdateOne) SetType(s string) *LogisticUpdateOne {
	luo.mutation.SetType(s)
	return luo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (luo *LogisticUpdateOne) SetNillableType(s *string) *LogisticUpdateOne {
	if s != nil {
		luo.SetType(*s)
	}
	return luo
}

// SetTask sets the "task" field.
func (luo *LogisticUpdateOne) SetTask(s *struct {
	Data interface{} "json:\"data\""
}) *LogisticUpdateOne {
	luo.mutation.SetTask(s)
	return luo
}

// ClearTask clears the value of the "task" field.
func (luo *LogisticUpdateOne) ClearTask() *LogisticUpdateOne {
	luo.mutation.ClearTask()
	return luo
}

// SetOrderID sets the "order" edge to the Order entity by ID.
func (luo *LogisticUpdateOne) SetOrderID(id int) *LogisticUpdateOne {
	luo.mutation.SetOrderID(id)
	return luo
}

// SetNillableOrderID sets the "order" edge to the Order entity by ID if the given value is not nil.
func (luo *LogisticUpdateOne) SetNillableOrderID(id *int) *LogisticUpdateOne {
	if id != nil {
		luo = luo.SetOrderID(*id)
	}
	return luo
}

// SetOrder sets the "order" edge to the Order entity.
func (luo *LogisticUpdateOne) SetOrder(o *Order) *LogisticUpdateOne {
	return luo.SetOrderID(o.ID)
}

// Mutation returns the LogisticMutation object of the builder.
func (luo *LogisticUpdateOne) Mutation() *LogisticMutation {
	return luo.mutation
}

// ClearOrder clears the "order" edge to the Order entity.
func (luo *LogisticUpdateOne) ClearOrder() *LogisticUpdateOne {
	luo.mutation.ClearOrder()
	return luo
}

// Where appends a list predicates to the LogisticUpdate builder.
func (luo *LogisticUpdateOne) Where(ps ...predicate.Logistic) *LogisticUpdateOne {
	luo.mutation.Where(ps...)
	return luo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LogisticUpdateOne) Select(field string, fields ...string) *LogisticUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Logistic entity.
func (luo *LogisticUpdateOne) Save(ctx context.Context) (*Logistic, error) {
	luo.defaults()
	return withHooks(ctx, luo.sqlSave, luo.mutation, luo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LogisticUpdateOne) SaveX(ctx context.Context) *Logistic {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LogisticUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LogisticUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (luo *LogisticUpdateOne) defaults() {
	if _, ok := luo.mutation.UpdatedAt(); !ok {
		v := logistic.UpdateDefaultUpdatedAt()
		luo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (luo *LogisticUpdateOne) check() error {
	if v, ok := luo.mutation.GetType(); ok {
		if err := logistic.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Logistic.type": %w`, err)}
		}
	}
	return nil
}

func (luo *LogisticUpdateOne) sqlSave(ctx context.Context) (_node *Logistic, err error) {
	if err := luo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(logistic.Table, logistic.Columns, sqlgraph.NewFieldSpec(logistic.FieldID, field.TypeInt))
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Logistic.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, logistic.FieldID)
		for _, f := range fields {
			if !logistic.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != logistic.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.UpdatedAt(); ok {
		_spec.SetField(logistic.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := luo.mutation.GetType(); ok {
		_spec.SetField(logistic.FieldType, field.TypeString, value)
	}
	if value, ok := luo.mutation.Task(); ok {
		_spec.SetField(logistic.FieldTask, field.TypeJSON, value)
	}
	if luo.mutation.TaskCleared() {
		_spec.ClearField(logistic.FieldTask, field.TypeJSON)
	}
	if luo.mutation.OrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   logistic.OrderTable,
			Columns: []string{logistic.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   logistic.OrderTable,
			Columns: []string{logistic.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Logistic{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{logistic.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	luo.mutation.done = true
	return _node, nil
}
