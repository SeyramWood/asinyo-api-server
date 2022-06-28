// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SeyramWood/ent/merchantstore"
	"github.com/SeyramWood/ent/order"
	"github.com/SeyramWood/ent/orderdetail"
	"github.com/SeyramWood/ent/product"
)

// OrderDetailCreate is the builder for creating a OrderDetail entity.
type OrderDetailCreate struct {
	config
	mutation *OrderDetailMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (odc *OrderDetailCreate) SetCreatedAt(t time.Time) *OrderDetailCreate {
	odc.mutation.SetCreatedAt(t)
	return odc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (odc *OrderDetailCreate) SetNillableCreatedAt(t *time.Time) *OrderDetailCreate {
	if t != nil {
		odc.SetCreatedAt(*t)
	}
	return odc
}

// SetUpdatedAt sets the "updated_at" field.
func (odc *OrderDetailCreate) SetUpdatedAt(t time.Time) *OrderDetailCreate {
	odc.mutation.SetUpdatedAt(t)
	return odc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (odc *OrderDetailCreate) SetNillableUpdatedAt(t *time.Time) *OrderDetailCreate {
	if t != nil {
		odc.SetUpdatedAt(*t)
	}
	return odc
}

// SetPrice sets the "price" field.
func (odc *OrderDetailCreate) SetPrice(f float64) *OrderDetailCreate {
	odc.mutation.SetPrice(f)
	return odc
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (odc *OrderDetailCreate) SetNillablePrice(f *float64) *OrderDetailCreate {
	if f != nil {
		odc.SetPrice(*f)
	}
	return odc
}

// SetPromoPrice sets the "promo_price" field.
func (odc *OrderDetailCreate) SetPromoPrice(f float64) *OrderDetailCreate {
	odc.mutation.SetPromoPrice(f)
	return odc
}

// SetNillablePromoPrice sets the "promo_price" field if the given value is not nil.
func (odc *OrderDetailCreate) SetNillablePromoPrice(f *float64) *OrderDetailCreate {
	if f != nil {
		odc.SetPromoPrice(*f)
	}
	return odc
}

// SetAmount sets the "amount" field.
func (odc *OrderDetailCreate) SetAmount(f float64) *OrderDetailCreate {
	odc.mutation.SetAmount(f)
	return odc
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (odc *OrderDetailCreate) SetNillableAmount(f *float64) *OrderDetailCreate {
	if f != nil {
		odc.SetAmount(*f)
	}
	return odc
}

// SetQuantity sets the "quantity" field.
func (odc *OrderDetailCreate) SetQuantity(i int) *OrderDetailCreate {
	odc.mutation.SetQuantity(i)
	return odc
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (odc *OrderDetailCreate) SetNillableQuantity(i *int) *OrderDetailCreate {
	if i != nil {
		odc.SetQuantity(*i)
	}
	return odc
}

// SetOrderID sets the "Order" edge to the Order entity by ID.
func (odc *OrderDetailCreate) SetOrderID(id int) *OrderDetailCreate {
	odc.mutation.SetOrderID(id)
	return odc
}

// SetOrder sets the "Order" edge to the Order entity.
func (odc *OrderDetailCreate) SetOrder(o *Order) *OrderDetailCreate {
	return odc.SetOrderID(o.ID)
}

// SetProductID sets the "product" edge to the Product entity by ID.
func (odc *OrderDetailCreate) SetProductID(id int) *OrderDetailCreate {
	odc.mutation.SetProductID(id)
	return odc
}

// SetProduct sets the "product" edge to the Product entity.
func (odc *OrderDetailCreate) SetProduct(p *Product) *OrderDetailCreate {
	return odc.SetProductID(p.ID)
}

// SetStoreID sets the "store" edge to the MerchantStore entity by ID.
func (odc *OrderDetailCreate) SetStoreID(id int) *OrderDetailCreate {
	odc.mutation.SetStoreID(id)
	return odc
}

// SetStore sets the "store" edge to the MerchantStore entity.
func (odc *OrderDetailCreate) SetStore(m *MerchantStore) *OrderDetailCreate {
	return odc.SetStoreID(m.ID)
}

// Mutation returns the OrderDetailMutation object of the builder.
func (odc *OrderDetailCreate) Mutation() *OrderDetailMutation {
	return odc.mutation
}

// Save creates the OrderDetail in the database.
func (odc *OrderDetailCreate) Save(ctx context.Context) (*OrderDetail, error) {
	var (
		err  error
		node *OrderDetail
	)
	odc.defaults()
	if len(odc.hooks) == 0 {
		if err = odc.check(); err != nil {
			return nil, err
		}
		node, err = odc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderDetailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = odc.check(); err != nil {
				return nil, err
			}
			odc.mutation = mutation
			if node, err = odc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(odc.hooks) - 1; i >= 0; i-- {
			if odc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = odc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, odc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (odc *OrderDetailCreate) SaveX(ctx context.Context) *OrderDetail {
	v, err := odc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (odc *OrderDetailCreate) Exec(ctx context.Context) error {
	_, err := odc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (odc *OrderDetailCreate) ExecX(ctx context.Context) {
	if err := odc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (odc *OrderDetailCreate) defaults() {
	if _, ok := odc.mutation.CreatedAt(); !ok {
		v := orderdetail.DefaultCreatedAt()
		odc.mutation.SetCreatedAt(v)
	}
	if _, ok := odc.mutation.UpdatedAt(); !ok {
		v := orderdetail.DefaultUpdatedAt()
		odc.mutation.SetUpdatedAt(v)
	}
	if _, ok := odc.mutation.Price(); !ok {
		v := orderdetail.DefaultPrice
		odc.mutation.SetPrice(v)
	}
	if _, ok := odc.mutation.PromoPrice(); !ok {
		v := orderdetail.DefaultPromoPrice
		odc.mutation.SetPromoPrice(v)
	}
	if _, ok := odc.mutation.Amount(); !ok {
		v := orderdetail.DefaultAmount
		odc.mutation.SetAmount(v)
	}
	if _, ok := odc.mutation.Quantity(); !ok {
		v := orderdetail.DefaultQuantity
		odc.mutation.SetQuantity(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (odc *OrderDetailCreate) check() error {
	if _, ok := odc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "OrderDetail.created_at"`)}
	}
	if _, ok := odc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "OrderDetail.updated_at"`)}
	}
	if _, ok := odc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "OrderDetail.price"`)}
	}
	if _, ok := odc.mutation.PromoPrice(); !ok {
		return &ValidationError{Name: "promo_price", err: errors.New(`ent: missing required field "OrderDetail.promo_price"`)}
	}
	if _, ok := odc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "OrderDetail.amount"`)}
	}
	if _, ok := odc.mutation.Quantity(); !ok {
		return &ValidationError{Name: "quantity", err: errors.New(`ent: missing required field "OrderDetail.quantity"`)}
	}
	if _, ok := odc.mutation.OrderID(); !ok {
		return &ValidationError{Name: "Order", err: errors.New(`ent: missing required edge "OrderDetail.Order"`)}
	}
	if _, ok := odc.mutation.ProductID(); !ok {
		return &ValidationError{Name: "product", err: errors.New(`ent: missing required edge "OrderDetail.product"`)}
	}
	if _, ok := odc.mutation.StoreID(); !ok {
		return &ValidationError{Name: "store", err: errors.New(`ent: missing required edge "OrderDetail.store"`)}
	}
	return nil
}

func (odc *OrderDetailCreate) sqlSave(ctx context.Context) (*OrderDetail, error) {
	_node, _spec := odc.createSpec()
	if err := sqlgraph.CreateNode(ctx, odc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (odc *OrderDetailCreate) createSpec() (*OrderDetail, *sqlgraph.CreateSpec) {
	var (
		_node = &OrderDetail{config: odc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: orderdetail.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: orderdetail.FieldID,
			},
		}
	)
	if value, ok := odc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: orderdetail.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := odc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: orderdetail.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := odc.mutation.Price(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: orderdetail.FieldPrice,
		})
		_node.Price = value
	}
	if value, ok := odc.mutation.PromoPrice(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: orderdetail.FieldPromoPrice,
		})
		_node.PromoPrice = value
	}
	if value, ok := odc.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: orderdetail.FieldAmount,
		})
		_node.Amount = value
	}
	if value, ok := odc.mutation.Quantity(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderdetail.FieldQuantity,
		})
		_node.Quantity = value
	}
	if nodes := odc.mutation.OrderIDs(); len(nodes) > 0 {
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
		_node.order_details = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := odc.mutation.ProductIDs(); len(nodes) > 0 {
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
		_node.product_orders = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := odc.mutation.StoreIDs(); len(nodes) > 0 {
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
		_node.merchant_store_orders = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrderDetailCreateBulk is the builder for creating many OrderDetail entities in bulk.
type OrderDetailCreateBulk struct {
	config
	builders []*OrderDetailCreate
}

// Save creates the OrderDetail entities in the database.
func (odcb *OrderDetailCreateBulk) Save(ctx context.Context) ([]*OrderDetail, error) {
	specs := make([]*sqlgraph.CreateSpec, len(odcb.builders))
	nodes := make([]*OrderDetail, len(odcb.builders))
	mutators := make([]Mutator, len(odcb.builders))
	for i := range odcb.builders {
		func(i int, root context.Context) {
			builder := odcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderDetailMutation)
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
					_, err = mutators[i+1].Mutate(root, odcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, odcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, odcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (odcb *OrderDetailCreateBulk) SaveX(ctx context.Context) []*OrderDetail {
	v, err := odcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (odcb *OrderDetailCreateBulk) Exec(ctx context.Context) error {
	_, err := odcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (odcb *OrderDetailCreateBulk) ExecX(ctx context.Context) {
	if err := odcb.Exec(ctx); err != nil {
		panic(err)
	}
}