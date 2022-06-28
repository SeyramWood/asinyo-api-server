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
	"github.com/SeyramWood/ent/merchantstore"
	"github.com/SeyramWood/ent/order"
	"github.com/SeyramWood/ent/orderdetail"
	"github.com/SeyramWood/ent/predicate"
	"github.com/SeyramWood/ent/product"
)

// OrderDetailUpdate is the builder for updating OrderDetail entities.
type OrderDetailUpdate struct {
	config
	hooks    []Hook
	mutation *OrderDetailMutation
}

// Where appends a list predicates to the OrderDetailUpdate builder.
func (odu *OrderDetailUpdate) Where(ps ...predicate.OrderDetail) *OrderDetailUpdate {
	odu.mutation.Where(ps...)
	return odu
}

// SetUpdatedAt sets the "updated_at" field.
func (odu *OrderDetailUpdate) SetUpdatedAt(t time.Time) *OrderDetailUpdate {
	odu.mutation.SetUpdatedAt(t)
	return odu
}

// SetPrice sets the "price" field.
func (odu *OrderDetailUpdate) SetPrice(f float64) *OrderDetailUpdate {
	odu.mutation.ResetPrice()
	odu.mutation.SetPrice(f)
	return odu
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (odu *OrderDetailUpdate) SetNillablePrice(f *float64) *OrderDetailUpdate {
	if f != nil {
		odu.SetPrice(*f)
	}
	return odu
}

// AddPrice adds f to the "price" field.
func (odu *OrderDetailUpdate) AddPrice(f float64) *OrderDetailUpdate {
	odu.mutation.AddPrice(f)
	return odu
}

// SetPromoPrice sets the "promo_price" field.
func (odu *OrderDetailUpdate) SetPromoPrice(f float64) *OrderDetailUpdate {
	odu.mutation.ResetPromoPrice()
	odu.mutation.SetPromoPrice(f)
	return odu
}

// SetNillablePromoPrice sets the "promo_price" field if the given value is not nil.
func (odu *OrderDetailUpdate) SetNillablePromoPrice(f *float64) *OrderDetailUpdate {
	if f != nil {
		odu.SetPromoPrice(*f)
	}
	return odu
}

// AddPromoPrice adds f to the "promo_price" field.
func (odu *OrderDetailUpdate) AddPromoPrice(f float64) *OrderDetailUpdate {
	odu.mutation.AddPromoPrice(f)
	return odu
}

// SetAmount sets the "amount" field.
func (odu *OrderDetailUpdate) SetAmount(f float64) *OrderDetailUpdate {
	odu.mutation.ResetAmount()
	odu.mutation.SetAmount(f)
	return odu
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (odu *OrderDetailUpdate) SetNillableAmount(f *float64) *OrderDetailUpdate {
	if f != nil {
		odu.SetAmount(*f)
	}
	return odu
}

// AddAmount adds f to the "amount" field.
func (odu *OrderDetailUpdate) AddAmount(f float64) *OrderDetailUpdate {
	odu.mutation.AddAmount(f)
	return odu
}

// SetQuantity sets the "quantity" field.
func (odu *OrderDetailUpdate) SetQuantity(i int) *OrderDetailUpdate {
	odu.mutation.ResetQuantity()
	odu.mutation.SetQuantity(i)
	return odu
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (odu *OrderDetailUpdate) SetNillableQuantity(i *int) *OrderDetailUpdate {
	if i != nil {
		odu.SetQuantity(*i)
	}
	return odu
}

// AddQuantity adds i to the "quantity" field.
func (odu *OrderDetailUpdate) AddQuantity(i int) *OrderDetailUpdate {
	odu.mutation.AddQuantity(i)
	return odu
}

// SetOrderID sets the "Order" edge to the Order entity by ID.
func (odu *OrderDetailUpdate) SetOrderID(id int) *OrderDetailUpdate {
	odu.mutation.SetOrderID(id)
	return odu
}

// SetOrder sets the "Order" edge to the Order entity.
func (odu *OrderDetailUpdate) SetOrder(o *Order) *OrderDetailUpdate {
	return odu.SetOrderID(o.ID)
}

// SetProductID sets the "product" edge to the Product entity by ID.
func (odu *OrderDetailUpdate) SetProductID(id int) *OrderDetailUpdate {
	odu.mutation.SetProductID(id)
	return odu
}

// SetProduct sets the "product" edge to the Product entity.
func (odu *OrderDetailUpdate) SetProduct(p *Product) *OrderDetailUpdate {
	return odu.SetProductID(p.ID)
}

// SetStoreID sets the "store" edge to the MerchantStore entity by ID.
func (odu *OrderDetailUpdate) SetStoreID(id int) *OrderDetailUpdate {
	odu.mutation.SetStoreID(id)
	return odu
}

// SetStore sets the "store" edge to the MerchantStore entity.
func (odu *OrderDetailUpdate) SetStore(m *MerchantStore) *OrderDetailUpdate {
	return odu.SetStoreID(m.ID)
}

// Mutation returns the OrderDetailMutation object of the builder.
func (odu *OrderDetailUpdate) Mutation() *OrderDetailMutation {
	return odu.mutation
}

// ClearOrder clears the "Order" edge to the Order entity.
func (odu *OrderDetailUpdate) ClearOrder() *OrderDetailUpdate {
	odu.mutation.ClearOrder()
	return odu
}

// ClearProduct clears the "product" edge to the Product entity.
func (odu *OrderDetailUpdate) ClearProduct() *OrderDetailUpdate {
	odu.mutation.ClearProduct()
	return odu
}

// ClearStore clears the "store" edge to the MerchantStore entity.
func (odu *OrderDetailUpdate) ClearStore() *OrderDetailUpdate {
	odu.mutation.ClearStore()
	return odu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (odu *OrderDetailUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	odu.defaults()
	if len(odu.hooks) == 0 {
		if err = odu.check(); err != nil {
			return 0, err
		}
		affected, err = odu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderDetailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = odu.check(); err != nil {
				return 0, err
			}
			odu.mutation = mutation
			affected, err = odu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(odu.hooks) - 1; i >= 0; i-- {
			if odu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = odu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, odu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (odu *OrderDetailUpdate) SaveX(ctx context.Context) int {
	affected, err := odu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (odu *OrderDetailUpdate) Exec(ctx context.Context) error {
	_, err := odu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (odu *OrderDetailUpdate) ExecX(ctx context.Context) {
	if err := odu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (odu *OrderDetailUpdate) defaults() {
	if _, ok := odu.mutation.UpdatedAt(); !ok {
		v := orderdetail.UpdateDefaultUpdatedAt()
		odu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (odu *OrderDetailUpdate) check() error {
	if _, ok := odu.mutation.OrderID(); odu.mutation.OrderCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "OrderDetail.Order"`)
	}
	if _, ok := odu.mutation.ProductID(); odu.mutation.ProductCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "OrderDetail.product"`)
	}
	if _, ok := odu.mutation.StoreID(); odu.mutation.StoreCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "OrderDetail.store"`)
	}
	return nil
}

func (odu *OrderDetailUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderdetail.Table,
			Columns: orderdetail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: orderdetail.FieldID,
			},
		},
	}
	if ps := odu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := odu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: orderdetail.FieldUpdatedAt,
		})
	}
	if value, ok := odu.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: orderdetail.FieldPrice,
		})
	}
	if value, ok := odu.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: orderdetail.FieldPrice,
		})
	}
	if value, ok := odu.mutation.PromoPrice(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: orderdetail.FieldPromoPrice,
		})
	}
	if value, ok := odu.mutation.AddedPromoPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: orderdetail.FieldPromoPrice,
		})
	}
	if value, ok := odu.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: orderdetail.FieldAmount,
		})
	}
	if value, ok := odu.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: orderdetail.FieldAmount,
		})
	}
	if value, ok := odu.mutation.Quantity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderdetail.FieldQuantity,
		})
	}
	if value, ok := odu.mutation.AddedQuantity(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderdetail.FieldQuantity,
		})
	}
	if odu.mutation.OrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderdetail.OrderTable,
			Columns: []string{orderdetail.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: order.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := odu.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderdetail.OrderTable,
			Columns: []string{orderdetail.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: order.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if odu.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderdetail.ProductTable,
			Columns: []string{orderdetail.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := odu.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderdetail.ProductTable,
			Columns: []string{orderdetail.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if odu.mutation.StoreCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderdetail.StoreTable,
			Columns: []string{orderdetail.StoreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: merchantstore.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := odu.mutation.StoreIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderdetail.StoreTable,
			Columns: []string{orderdetail.StoreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: merchantstore.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, odu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderdetail.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// OrderDetailUpdateOne is the builder for updating a single OrderDetail entity.
type OrderDetailUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrderDetailMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (oduo *OrderDetailUpdateOne) SetUpdatedAt(t time.Time) *OrderDetailUpdateOne {
	oduo.mutation.SetUpdatedAt(t)
	return oduo
}

// SetPrice sets the "price" field.
func (oduo *OrderDetailUpdateOne) SetPrice(f float64) *OrderDetailUpdateOne {
	oduo.mutation.ResetPrice()
	oduo.mutation.SetPrice(f)
	return oduo
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (oduo *OrderDetailUpdateOne) SetNillablePrice(f *float64) *OrderDetailUpdateOne {
	if f != nil {
		oduo.SetPrice(*f)
	}
	return oduo
}

// AddPrice adds f to the "price" field.
func (oduo *OrderDetailUpdateOne) AddPrice(f float64) *OrderDetailUpdateOne {
	oduo.mutation.AddPrice(f)
	return oduo
}

// SetPromoPrice sets the "promo_price" field.
func (oduo *OrderDetailUpdateOne) SetPromoPrice(f float64) *OrderDetailUpdateOne {
	oduo.mutation.ResetPromoPrice()
	oduo.mutation.SetPromoPrice(f)
	return oduo
}

// SetNillablePromoPrice sets the "promo_price" field if the given value is not nil.
func (oduo *OrderDetailUpdateOne) SetNillablePromoPrice(f *float64) *OrderDetailUpdateOne {
	if f != nil {
		oduo.SetPromoPrice(*f)
	}
	return oduo
}

// AddPromoPrice adds f to the "promo_price" field.
func (oduo *OrderDetailUpdateOne) AddPromoPrice(f float64) *OrderDetailUpdateOne {
	oduo.mutation.AddPromoPrice(f)
	return oduo
}

// SetAmount sets the "amount" field.
func (oduo *OrderDetailUpdateOne) SetAmount(f float64) *OrderDetailUpdateOne {
	oduo.mutation.ResetAmount()
	oduo.mutation.SetAmount(f)
	return oduo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (oduo *OrderDetailUpdateOne) SetNillableAmount(f *float64) *OrderDetailUpdateOne {
	if f != nil {
		oduo.SetAmount(*f)
	}
	return oduo
}

// AddAmount adds f to the "amount" field.
func (oduo *OrderDetailUpdateOne) AddAmount(f float64) *OrderDetailUpdateOne {
	oduo.mutation.AddAmount(f)
	return oduo
}

// SetQuantity sets the "quantity" field.
func (oduo *OrderDetailUpdateOne) SetQuantity(i int) *OrderDetailUpdateOne {
	oduo.mutation.ResetQuantity()
	oduo.mutation.SetQuantity(i)
	return oduo
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (oduo *OrderDetailUpdateOne) SetNillableQuantity(i *int) *OrderDetailUpdateOne {
	if i != nil {
		oduo.SetQuantity(*i)
	}
	return oduo
}

// AddQuantity adds i to the "quantity" field.
func (oduo *OrderDetailUpdateOne) AddQuantity(i int) *OrderDetailUpdateOne {
	oduo.mutation.AddQuantity(i)
	return oduo
}

// SetOrderID sets the "Order" edge to the Order entity by ID.
func (oduo *OrderDetailUpdateOne) SetOrderID(id int) *OrderDetailUpdateOne {
	oduo.mutation.SetOrderID(id)
	return oduo
}

// SetOrder sets the "Order" edge to the Order entity.
func (oduo *OrderDetailUpdateOne) SetOrder(o *Order) *OrderDetailUpdateOne {
	return oduo.SetOrderID(o.ID)
}

// SetProductID sets the "product" edge to the Product entity by ID.
func (oduo *OrderDetailUpdateOne) SetProductID(id int) *OrderDetailUpdateOne {
	oduo.mutation.SetProductID(id)
	return oduo
}

// SetProduct sets the "product" edge to the Product entity.
func (oduo *OrderDetailUpdateOne) SetProduct(p *Product) *OrderDetailUpdateOne {
	return oduo.SetProductID(p.ID)
}

// SetStoreID sets the "store" edge to the MerchantStore entity by ID.
func (oduo *OrderDetailUpdateOne) SetStoreID(id int) *OrderDetailUpdateOne {
	oduo.mutation.SetStoreID(id)
	return oduo
}

// SetStore sets the "store" edge to the MerchantStore entity.
func (oduo *OrderDetailUpdateOne) SetStore(m *MerchantStore) *OrderDetailUpdateOne {
	return oduo.SetStoreID(m.ID)
}

// Mutation returns the OrderDetailMutation object of the builder.
func (oduo *OrderDetailUpdateOne) Mutation() *OrderDetailMutation {
	return oduo.mutation
}

// ClearOrder clears the "Order" edge to the Order entity.
func (oduo *OrderDetailUpdateOne) ClearOrder() *OrderDetailUpdateOne {
	oduo.mutation.ClearOrder()
	return oduo
}

// ClearProduct clears the "product" edge to the Product entity.
func (oduo *OrderDetailUpdateOne) ClearProduct() *OrderDetailUpdateOne {
	oduo.mutation.ClearProduct()
	return oduo
}

// ClearStore clears the "store" edge to the MerchantStore entity.
func (oduo *OrderDetailUpdateOne) ClearStore() *OrderDetailUpdateOne {
	oduo.mutation.ClearStore()
	return oduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (oduo *OrderDetailUpdateOne) Select(field string, fields ...string) *OrderDetailUpdateOne {
	oduo.fields = append([]string{field}, fields...)
	return oduo
}

// Save executes the query and returns the updated OrderDetail entity.
func (oduo *OrderDetailUpdateOne) Save(ctx context.Context) (*OrderDetail, error) {
	var (
		err  error
		node *OrderDetail
	)
	oduo.defaults()
	if len(oduo.hooks) == 0 {
		if err = oduo.check(); err != nil {
			return nil, err
		}
		node, err = oduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderDetailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = oduo.check(); err != nil {
				return nil, err
			}
			oduo.mutation = mutation
			node, err = oduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(oduo.hooks) - 1; i >= 0; i-- {
			if oduo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = oduo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oduo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (oduo *OrderDetailUpdateOne) SaveX(ctx context.Context) *OrderDetail {
	node, err := oduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (oduo *OrderDetailUpdateOne) Exec(ctx context.Context) error {
	_, err := oduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oduo *OrderDetailUpdateOne) ExecX(ctx context.Context) {
	if err := oduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oduo *OrderDetailUpdateOne) defaults() {
	if _, ok := oduo.mutation.UpdatedAt(); !ok {
		v := orderdetail.UpdateDefaultUpdatedAt()
		oduo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oduo *OrderDetailUpdateOne) check() error {
	if _, ok := oduo.mutation.OrderID(); oduo.mutation.OrderCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "OrderDetail.Order"`)
	}
	if _, ok := oduo.mutation.ProductID(); oduo.mutation.ProductCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "OrderDetail.product"`)
	}
	if _, ok := oduo.mutation.StoreID(); oduo.mutation.StoreCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "OrderDetail.store"`)
	}
	return nil
}

func (oduo *OrderDetailUpdateOne) sqlSave(ctx context.Context) (_node *OrderDetail, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderdetail.Table,
			Columns: orderdetail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: orderdetail.FieldID,
			},
		},
	}
	id, ok := oduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OrderDetail.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := oduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderdetail.FieldID)
		for _, f := range fields {
			if !orderdetail.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != orderdetail.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := oduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oduo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: orderdetail.FieldUpdatedAt,
		})
	}
	if value, ok := oduo.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: orderdetail.FieldPrice,
		})
	}
	if value, ok := oduo.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: orderdetail.FieldPrice,
		})
	}
	if value, ok := oduo.mutation.PromoPrice(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: orderdetail.FieldPromoPrice,
		})
	}
	if value, ok := oduo.mutation.AddedPromoPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: orderdetail.FieldPromoPrice,
		})
	}
	if value, ok := oduo.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: orderdetail.FieldAmount,
		})
	}
	if value, ok := oduo.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: orderdetail.FieldAmount,
		})
	}
	if value, ok := oduo.mutation.Quantity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderdetail.FieldQuantity,
		})
	}
	if value, ok := oduo.mutation.AddedQuantity(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderdetail.FieldQuantity,
		})
	}
	if oduo.mutation.OrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderdetail.OrderTable,
			Columns: []string{orderdetail.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: order.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oduo.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderdetail.OrderTable,
			Columns: []string{orderdetail.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: order.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if oduo.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderdetail.ProductTable,
			Columns: []string{orderdetail.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oduo.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderdetail.ProductTable,
			Columns: []string{orderdetail.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if oduo.mutation.StoreCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderdetail.StoreTable,
			Columns: []string{orderdetail.StoreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: merchantstore.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oduo.mutation.StoreIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderdetail.StoreTable,
			Columns: []string{orderdetail.StoreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: merchantstore.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &OrderDetail{config: oduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, oduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderdetail.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}