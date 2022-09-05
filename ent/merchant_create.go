// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SeyramWood/ent/address"
	"github.com/SeyramWood/ent/favourite"
	"github.com/SeyramWood/ent/merchant"
	"github.com/SeyramWood/ent/merchantstore"
	"github.com/SeyramWood/ent/order"
	"github.com/SeyramWood/ent/product"
	"github.com/SeyramWood/ent/retailmerchant"
	"github.com/SeyramWood/ent/suppliermerchant"
)

// MerchantCreate is the builder for creating a Merchant entity.
type MerchantCreate struct {
	config
	mutation *MerchantMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (mc *MerchantCreate) SetCreatedAt(t time.Time) *MerchantCreate {
	mc.mutation.SetCreatedAt(t)
	return mc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mc *MerchantCreate) SetNillableCreatedAt(t *time.Time) *MerchantCreate {
	if t != nil {
		mc.SetCreatedAt(*t)
	}
	return mc
}

// SetUpdatedAt sets the "updated_at" field.
func (mc *MerchantCreate) SetUpdatedAt(t time.Time) *MerchantCreate {
	mc.mutation.SetUpdatedAt(t)
	return mc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mc *MerchantCreate) SetNillableUpdatedAt(t *time.Time) *MerchantCreate {
	if t != nil {
		mc.SetUpdatedAt(*t)
	}
	return mc
}

// SetUsername sets the "username" field.
func (mc *MerchantCreate) SetUsername(s string) *MerchantCreate {
	mc.mutation.SetUsername(s)
	return mc
}

// SetPassword sets the "password" field.
func (mc *MerchantCreate) SetPassword(b []byte) *MerchantCreate {
	mc.mutation.SetPassword(b)
	return mc
}

// SetType sets the "type" field.
func (mc *MerchantCreate) SetType(s string) *MerchantCreate {
	mc.mutation.SetType(s)
	return mc
}

// SetOtp sets the "otp" field.
func (mc *MerchantCreate) SetOtp(m merchant.Otp) *MerchantCreate {
	mc.mutation.SetOtp(m)
	return mc
}

// SetNillableOtp sets the "otp" field if the given value is not nil.
func (mc *MerchantCreate) SetNillableOtp(m *merchant.Otp) *MerchantCreate {
	if m != nil {
		mc.SetOtp(*m)
	}
	return mc
}

// SetSupplierID sets the "supplier" edge to the SupplierMerchant entity by ID.
func (mc *MerchantCreate) SetSupplierID(id int) *MerchantCreate {
	mc.mutation.SetSupplierID(id)
	return mc
}

// SetNillableSupplierID sets the "supplier" edge to the SupplierMerchant entity by ID if the given value is not nil.
func (mc *MerchantCreate) SetNillableSupplierID(id *int) *MerchantCreate {
	if id != nil {
		mc = mc.SetSupplierID(*id)
	}
	return mc
}

// SetSupplier sets the "supplier" edge to the SupplierMerchant entity.
func (mc *MerchantCreate) SetSupplier(s *SupplierMerchant) *MerchantCreate {
	return mc.SetSupplierID(s.ID)
}

// SetRetailerID sets the "retailer" edge to the RetailMerchant entity by ID.
func (mc *MerchantCreate) SetRetailerID(id int) *MerchantCreate {
	mc.mutation.SetRetailerID(id)
	return mc
}

// SetNillableRetailerID sets the "retailer" edge to the RetailMerchant entity by ID if the given value is not nil.
func (mc *MerchantCreate) SetNillableRetailerID(id *int) *MerchantCreate {
	if id != nil {
		mc = mc.SetRetailerID(*id)
	}
	return mc
}

// SetRetailer sets the "retailer" edge to the RetailMerchant entity.
func (mc *MerchantCreate) SetRetailer(r *RetailMerchant) *MerchantCreate {
	return mc.SetRetailerID(r.ID)
}

// SetStoreID sets the "store" edge to the MerchantStore entity by ID.
func (mc *MerchantCreate) SetStoreID(id int) *MerchantCreate {
	mc.mutation.SetStoreID(id)
	return mc
}

// SetNillableStoreID sets the "store" edge to the MerchantStore entity by ID if the given value is not nil.
func (mc *MerchantCreate) SetNillableStoreID(id *int) *MerchantCreate {
	if id != nil {
		mc = mc.SetStoreID(*id)
	}
	return mc
}

// SetStore sets the "store" edge to the MerchantStore entity.
func (mc *MerchantCreate) SetStore(m *MerchantStore) *MerchantCreate {
	return mc.SetStoreID(m.ID)
}

// AddProductIDs adds the "products" edge to the Product entity by IDs.
func (mc *MerchantCreate) AddProductIDs(ids ...int) *MerchantCreate {
	mc.mutation.AddProductIDs(ids...)
	return mc
}

// AddProducts adds the "products" edges to the Product entity.
func (mc *MerchantCreate) AddProducts(p ...*Product) *MerchantCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mc.AddProductIDs(ids...)
}

// AddAddressIDs adds the "addresses" edge to the Address entity by IDs.
func (mc *MerchantCreate) AddAddressIDs(ids ...int) *MerchantCreate {
	mc.mutation.AddAddressIDs(ids...)
	return mc
}

// AddAddresses adds the "addresses" edges to the Address entity.
func (mc *MerchantCreate) AddAddresses(a ...*Address) *MerchantCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return mc.AddAddressIDs(ids...)
}

// AddOrderIDs adds the "orders" edge to the Order entity by IDs.
func (mc *MerchantCreate) AddOrderIDs(ids ...int) *MerchantCreate {
	mc.mutation.AddOrderIDs(ids...)
	return mc
}

// AddOrders adds the "orders" edges to the Order entity.
func (mc *MerchantCreate) AddOrders(o ...*Order) *MerchantCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return mc.AddOrderIDs(ids...)
}

// AddFavouriteIDs adds the "favourites" edge to the Favourite entity by IDs.
func (mc *MerchantCreate) AddFavouriteIDs(ids ...int) *MerchantCreate {
	mc.mutation.AddFavouriteIDs(ids...)
	return mc
}

// AddFavourites adds the "favourites" edges to the Favourite entity.
func (mc *MerchantCreate) AddFavourites(f ...*Favourite) *MerchantCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return mc.AddFavouriteIDs(ids...)
}

// Mutation returns the MerchantMutation object of the builder.
func (mc *MerchantCreate) Mutation() *MerchantMutation {
	return mc.mutation
}

// Save creates the Merchant in the database.
func (mc *MerchantCreate) Save(ctx context.Context) (*Merchant, error) {
	var (
		err  error
		node *Merchant
	)
	mc.defaults()
	if len(mc.hooks) == 0 {
		if err = mc.check(); err != nil {
			return nil, err
		}
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MerchantMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mc.check(); err != nil {
				return nil, err
			}
			mc.mutation = mutation
			if node, err = mc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			if mc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MerchantCreate) SaveX(ctx context.Context) *Merchant {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MerchantCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MerchantCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MerchantCreate) defaults() {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		v := merchant.DefaultCreatedAt()
		mc.mutation.SetCreatedAt(v)
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		v := merchant.DefaultUpdatedAt()
		mc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MerchantCreate) check() error {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Merchant.created_at"`)}
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Merchant.updated_at"`)}
	}
	if _, ok := mc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "Merchant.username"`)}
	}
	if v, ok := mc.mutation.Username(); ok {
		if err := merchant.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "Merchant.username": %w`, err)}
		}
	}
	if _, ok := mc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "Merchant.password"`)}
	}
	if v, ok := mc.mutation.Password(); ok {
		if err := merchant.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Merchant.password": %w`, err)}
		}
	}
	if _, ok := mc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Merchant.type"`)}
	}
	if v, ok := mc.mutation.GetType(); ok {
		if err := merchant.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Merchant.type": %w`, err)}
		}
	}
	if v, ok := mc.mutation.Otp(); ok {
		if err := merchant.OtpValidator(v); err != nil {
			return &ValidationError{Name: "otp", err: fmt.Errorf(`ent: validator failed for field "Merchant.otp": %w`, err)}
		}
	}
	return nil
}

func (mc *MerchantCreate) sqlSave(ctx context.Context) (*Merchant, error) {
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (mc *MerchantCreate) createSpec() (*Merchant, *sqlgraph.CreateSpec) {
	var (
		_node = &Merchant{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: merchant.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: merchant.FieldID,
			},
		}
	)
	if value, ok := mc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: merchant.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := mc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: merchant.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := mc.mutation.Username(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchant.FieldUsername,
		})
		_node.Username = value
	}
	if value, ok := mc.mutation.Password(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: merchant.FieldPassword,
		})
		_node.Password = value
	}
	if value, ok := mc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchant.FieldType,
		})
		_node.Type = value
	}
	if value, ok := mc.mutation.Otp(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: merchant.FieldOtp,
		})
		_node.Otp = value
	}
	if nodes := mc.mutation.SupplierIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   merchant.SupplierTable,
			Columns: []string{merchant.SupplierColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: suppliermerchant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.RetailerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   merchant.RetailerTable,
			Columns: []string{merchant.RetailerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: retailmerchant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.StoreIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   merchant.StoreTable,
			Columns: []string{merchant.StoreColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   merchant.ProductsTable,
			Columns: []string{merchant.ProductsColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.AddressesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   merchant.AddressesTable,
			Columns: []string{merchant.AddressesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: address.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.OrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   merchant.OrdersTable,
			Columns: []string{merchant.OrdersColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.FavouritesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   merchant.FavouritesTable,
			Columns: []string{merchant.FavouritesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: favourite.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MerchantCreateBulk is the builder for creating many Merchant entities in bulk.
type MerchantCreateBulk struct {
	config
	builders []*MerchantCreate
}

// Save creates the Merchant entities in the database.
func (mcb *MerchantCreateBulk) Save(ctx context.Context) ([]*Merchant, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Merchant, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MerchantMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MerchantCreateBulk) SaveX(ctx context.Context) []*Merchant {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MerchantCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MerchantCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
