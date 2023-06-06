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
	"github.com/SeyramWood/ent/order"
	"github.com/SeyramWood/ent/pickupstation"
	"github.com/SeyramWood/ent/predicate"
)

// PickupStationUpdate is the builder for updating PickupStation entities.
type PickupStationUpdate struct {
	config
	hooks    []Hook
	mutation *PickupStationMutation
}

// Where appends a list predicates to the PickupStationUpdate builder.
func (psu *PickupStationUpdate) Where(ps ...predicate.PickupStation) *PickupStationUpdate {
	psu.mutation.Where(ps...)
	return psu
}

// SetUpdatedAt sets the "updated_at" field.
func (psu *PickupStationUpdate) SetUpdatedAt(t time.Time) *PickupStationUpdate {
	psu.mutation.SetUpdatedAt(t)
	return psu
}

// SetRegion sets the "region" field.
func (psu *PickupStationUpdate) SetRegion(s string) *PickupStationUpdate {
	psu.mutation.SetRegion(s)
	return psu
}

// SetCity sets the "city" field.
func (psu *PickupStationUpdate) SetCity(s string) *PickupStationUpdate {
	psu.mutation.SetCity(s)
	return psu
}

// SetName sets the "name" field.
func (psu *PickupStationUpdate) SetName(s string) *PickupStationUpdate {
	psu.mutation.SetName(s)
	return psu
}

// SetAddress sets the "address" field.
func (psu *PickupStationUpdate) SetAddress(s string) *PickupStationUpdate {
	psu.mutation.SetAddress(s)
	return psu
}

// AddOrderIDs adds the "orders" edge to the Order entity by IDs.
func (psu *PickupStationUpdate) AddOrderIDs(ids ...int) *PickupStationUpdate {
	psu.mutation.AddOrderIDs(ids...)
	return psu
}

// AddOrders adds the "orders" edges to the Order entity.
func (psu *PickupStationUpdate) AddOrders(o ...*Order) *PickupStationUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return psu.AddOrderIDs(ids...)
}

// Mutation returns the PickupStationMutation object of the builder.
func (psu *PickupStationUpdate) Mutation() *PickupStationMutation {
	return psu.mutation
}

// ClearOrders clears all "orders" edges to the Order entity.
func (psu *PickupStationUpdate) ClearOrders() *PickupStationUpdate {
	psu.mutation.ClearOrders()
	return psu
}

// RemoveOrderIDs removes the "orders" edge to Order entities by IDs.
func (psu *PickupStationUpdate) RemoveOrderIDs(ids ...int) *PickupStationUpdate {
	psu.mutation.RemoveOrderIDs(ids...)
	return psu
}

// RemoveOrders removes "orders" edges to Order entities.
func (psu *PickupStationUpdate) RemoveOrders(o ...*Order) *PickupStationUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return psu.RemoveOrderIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (psu *PickupStationUpdate) Save(ctx context.Context) (int, error) {
	psu.defaults()
	return withHooks(ctx, psu.sqlSave, psu.mutation, psu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (psu *PickupStationUpdate) SaveX(ctx context.Context) int {
	affected, err := psu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (psu *PickupStationUpdate) Exec(ctx context.Context) error {
	_, err := psu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psu *PickupStationUpdate) ExecX(ctx context.Context) {
	if err := psu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (psu *PickupStationUpdate) defaults() {
	if _, ok := psu.mutation.UpdatedAt(); !ok {
		v := pickupstation.UpdateDefaultUpdatedAt()
		psu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (psu *PickupStationUpdate) check() error {
	if v, ok := psu.mutation.Region(); ok {
		if err := pickupstation.RegionValidator(v); err != nil {
			return &ValidationError{Name: "region", err: fmt.Errorf(`ent: validator failed for field "PickupStation.region": %w`, err)}
		}
	}
	if v, ok := psu.mutation.City(); ok {
		if err := pickupstation.CityValidator(v); err != nil {
			return &ValidationError{Name: "city", err: fmt.Errorf(`ent: validator failed for field "PickupStation.city": %w`, err)}
		}
	}
	if v, ok := psu.mutation.Name(); ok {
		if err := pickupstation.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "PickupStation.name": %w`, err)}
		}
	}
	if v, ok := psu.mutation.Address(); ok {
		if err := pickupstation.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`ent: validator failed for field "PickupStation.address": %w`, err)}
		}
	}
	return nil
}

func (psu *PickupStationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := psu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(pickupstation.Table, pickupstation.Columns, sqlgraph.NewFieldSpec(pickupstation.FieldID, field.TypeInt))
	if ps := psu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := psu.mutation.UpdatedAt(); ok {
		_spec.SetField(pickupstation.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := psu.mutation.Region(); ok {
		_spec.SetField(pickupstation.FieldRegion, field.TypeString, value)
	}
	if value, ok := psu.mutation.City(); ok {
		_spec.SetField(pickupstation.FieldCity, field.TypeString, value)
	}
	if value, ok := psu.mutation.Name(); ok {
		_spec.SetField(pickupstation.FieldName, field.TypeString, value)
	}
	if value, ok := psu.mutation.Address(); ok {
		_spec.SetField(pickupstation.FieldAddress, field.TypeString, value)
	}
	if psu.mutation.OrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pickupstation.OrdersTable,
			Columns: []string{pickupstation.OrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psu.mutation.RemovedOrdersIDs(); len(nodes) > 0 && !psu.mutation.OrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pickupstation.OrdersTable,
			Columns: []string{pickupstation.OrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psu.mutation.OrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pickupstation.OrdersTable,
			Columns: []string{pickupstation.OrdersColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, psu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pickupstation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	psu.mutation.done = true
	return n, nil
}

// PickupStationUpdateOne is the builder for updating a single PickupStation entity.
type PickupStationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PickupStationMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (psuo *PickupStationUpdateOne) SetUpdatedAt(t time.Time) *PickupStationUpdateOne {
	psuo.mutation.SetUpdatedAt(t)
	return psuo
}

// SetRegion sets the "region" field.
func (psuo *PickupStationUpdateOne) SetRegion(s string) *PickupStationUpdateOne {
	psuo.mutation.SetRegion(s)
	return psuo
}

// SetCity sets the "city" field.
func (psuo *PickupStationUpdateOne) SetCity(s string) *PickupStationUpdateOne {
	psuo.mutation.SetCity(s)
	return psuo
}

// SetName sets the "name" field.
func (psuo *PickupStationUpdateOne) SetName(s string) *PickupStationUpdateOne {
	psuo.mutation.SetName(s)
	return psuo
}

// SetAddress sets the "address" field.
func (psuo *PickupStationUpdateOne) SetAddress(s string) *PickupStationUpdateOne {
	psuo.mutation.SetAddress(s)
	return psuo
}

// AddOrderIDs adds the "orders" edge to the Order entity by IDs.
func (psuo *PickupStationUpdateOne) AddOrderIDs(ids ...int) *PickupStationUpdateOne {
	psuo.mutation.AddOrderIDs(ids...)
	return psuo
}

// AddOrders adds the "orders" edges to the Order entity.
func (psuo *PickupStationUpdateOne) AddOrders(o ...*Order) *PickupStationUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return psuo.AddOrderIDs(ids...)
}

// Mutation returns the PickupStationMutation object of the builder.
func (psuo *PickupStationUpdateOne) Mutation() *PickupStationMutation {
	return psuo.mutation
}

// ClearOrders clears all "orders" edges to the Order entity.
func (psuo *PickupStationUpdateOne) ClearOrders() *PickupStationUpdateOne {
	psuo.mutation.ClearOrders()
	return psuo
}

// RemoveOrderIDs removes the "orders" edge to Order entities by IDs.
func (psuo *PickupStationUpdateOne) RemoveOrderIDs(ids ...int) *PickupStationUpdateOne {
	psuo.mutation.RemoveOrderIDs(ids...)
	return psuo
}

// RemoveOrders removes "orders" edges to Order entities.
func (psuo *PickupStationUpdateOne) RemoveOrders(o ...*Order) *PickupStationUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return psuo.RemoveOrderIDs(ids...)
}

// Where appends a list predicates to the PickupStationUpdate builder.
func (psuo *PickupStationUpdateOne) Where(ps ...predicate.PickupStation) *PickupStationUpdateOne {
	psuo.mutation.Where(ps...)
	return psuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (psuo *PickupStationUpdateOne) Select(field string, fields ...string) *PickupStationUpdateOne {
	psuo.fields = append([]string{field}, fields...)
	return psuo
}

// Save executes the query and returns the updated PickupStation entity.
func (psuo *PickupStationUpdateOne) Save(ctx context.Context) (*PickupStation, error) {
	psuo.defaults()
	return withHooks(ctx, psuo.sqlSave, psuo.mutation, psuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (psuo *PickupStationUpdateOne) SaveX(ctx context.Context) *PickupStation {
	node, err := psuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (psuo *PickupStationUpdateOne) Exec(ctx context.Context) error {
	_, err := psuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psuo *PickupStationUpdateOne) ExecX(ctx context.Context) {
	if err := psuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (psuo *PickupStationUpdateOne) defaults() {
	if _, ok := psuo.mutation.UpdatedAt(); !ok {
		v := pickupstation.UpdateDefaultUpdatedAt()
		psuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (psuo *PickupStationUpdateOne) check() error {
	if v, ok := psuo.mutation.Region(); ok {
		if err := pickupstation.RegionValidator(v); err != nil {
			return &ValidationError{Name: "region", err: fmt.Errorf(`ent: validator failed for field "PickupStation.region": %w`, err)}
		}
	}
	if v, ok := psuo.mutation.City(); ok {
		if err := pickupstation.CityValidator(v); err != nil {
			return &ValidationError{Name: "city", err: fmt.Errorf(`ent: validator failed for field "PickupStation.city": %w`, err)}
		}
	}
	if v, ok := psuo.mutation.Name(); ok {
		if err := pickupstation.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "PickupStation.name": %w`, err)}
		}
	}
	if v, ok := psuo.mutation.Address(); ok {
		if err := pickupstation.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`ent: validator failed for field "PickupStation.address": %w`, err)}
		}
	}
	return nil
}

func (psuo *PickupStationUpdateOne) sqlSave(ctx context.Context) (_node *PickupStation, err error) {
	if err := psuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(pickupstation.Table, pickupstation.Columns, sqlgraph.NewFieldSpec(pickupstation.FieldID, field.TypeInt))
	id, ok := psuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PickupStation.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := psuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pickupstation.FieldID)
		for _, f := range fields {
			if !pickupstation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != pickupstation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := psuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := psuo.mutation.UpdatedAt(); ok {
		_spec.SetField(pickupstation.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := psuo.mutation.Region(); ok {
		_spec.SetField(pickupstation.FieldRegion, field.TypeString, value)
	}
	if value, ok := psuo.mutation.City(); ok {
		_spec.SetField(pickupstation.FieldCity, field.TypeString, value)
	}
	if value, ok := psuo.mutation.Name(); ok {
		_spec.SetField(pickupstation.FieldName, field.TypeString, value)
	}
	if value, ok := psuo.mutation.Address(); ok {
		_spec.SetField(pickupstation.FieldAddress, field.TypeString, value)
	}
	if psuo.mutation.OrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pickupstation.OrdersTable,
			Columns: []string{pickupstation.OrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psuo.mutation.RemovedOrdersIDs(); len(nodes) > 0 && !psuo.mutation.OrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pickupstation.OrdersTable,
			Columns: []string{pickupstation.OrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psuo.mutation.OrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pickupstation.OrdersTable,
			Columns: []string{pickupstation.OrdersColumn},
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
	_node = &PickupStation{config: psuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, psuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pickupstation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	psuo.mutation.done = true
	return _node, nil
}
