// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SeyramWood/ent/address"
	"github.com/SeyramWood/ent/admin"
	"github.com/SeyramWood/ent/businesscustomer"
	"github.com/SeyramWood/ent/customer"
	"github.com/SeyramWood/ent/favourite"
	"github.com/SeyramWood/ent/individualcustomer"
	"github.com/SeyramWood/ent/notification"
	"github.com/SeyramWood/ent/order"
	"github.com/SeyramWood/ent/purchaserequest"
)

// CustomerCreate is the builder for creating a Customer entity.
type CustomerCreate struct {
	config
	mutation *CustomerMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (cc *CustomerCreate) SetCreatedAt(t time.Time) *CustomerCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CustomerCreate) SetNillableCreatedAt(t *time.Time) *CustomerCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CustomerCreate) SetUpdatedAt(t time.Time) *CustomerCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CustomerCreate) SetNillableUpdatedAt(t *time.Time) *CustomerCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetUsername sets the "username" field.
func (cc *CustomerCreate) SetUsername(s string) *CustomerCreate {
	cc.mutation.SetUsername(s)
	return cc
}

// SetPassword sets the "password" field.
func (cc *CustomerCreate) SetPassword(b []byte) *CustomerCreate {
	cc.mutation.SetPassword(b)
	return cc
}

// SetType sets the "type" field.
func (cc *CustomerCreate) SetType(s string) *CustomerCreate {
	cc.mutation.SetType(s)
	return cc
}

// SetBusinessID sets the "business" edge to the BusinessCustomer entity by ID.
func (cc *CustomerCreate) SetBusinessID(id int) *CustomerCreate {
	cc.mutation.SetBusinessID(id)
	return cc
}

// SetNillableBusinessID sets the "business" edge to the BusinessCustomer entity by ID if the given value is not nil.
func (cc *CustomerCreate) SetNillableBusinessID(id *int) *CustomerCreate {
	if id != nil {
		cc = cc.SetBusinessID(*id)
	}
	return cc
}

// SetBusiness sets the "business" edge to the BusinessCustomer entity.
func (cc *CustomerCreate) SetBusiness(b *BusinessCustomer) *CustomerCreate {
	return cc.SetBusinessID(b.ID)
}

// SetIndividualID sets the "individual" edge to the IndividualCustomer entity by ID.
func (cc *CustomerCreate) SetIndividualID(id int) *CustomerCreate {
	cc.mutation.SetIndividualID(id)
	return cc
}

// SetNillableIndividualID sets the "individual" edge to the IndividualCustomer entity by ID if the given value is not nil.
func (cc *CustomerCreate) SetNillableIndividualID(id *int) *CustomerCreate {
	if id != nil {
		cc = cc.SetIndividualID(*id)
	}
	return cc
}

// SetIndividual sets the "individual" edge to the IndividualCustomer entity.
func (cc *CustomerCreate) SetIndividual(i *IndividualCustomer) *CustomerCreate {
	return cc.SetIndividualID(i.ID)
}

// AddAddressIDs adds the "addresses" edge to the Address entity by IDs.
func (cc *CustomerCreate) AddAddressIDs(ids ...int) *CustomerCreate {
	cc.mutation.AddAddressIDs(ids...)
	return cc
}

// AddAddresses adds the "addresses" edges to the Address entity.
func (cc *CustomerCreate) AddAddresses(a ...*Address) *CustomerCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cc.AddAddressIDs(ids...)
}

// AddOrderIDs adds the "orders" edge to the Order entity by IDs.
func (cc *CustomerCreate) AddOrderIDs(ids ...int) *CustomerCreate {
	cc.mutation.AddOrderIDs(ids...)
	return cc
}

// AddOrders adds the "orders" edges to the Order entity.
func (cc *CustomerCreate) AddOrders(o ...*Order) *CustomerCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return cc.AddOrderIDs(ids...)
}

// AddFavouriteIDs adds the "favourites" edge to the Favourite entity by IDs.
func (cc *CustomerCreate) AddFavouriteIDs(ids ...int) *CustomerCreate {
	cc.mutation.AddFavouriteIDs(ids...)
	return cc
}

// AddFavourites adds the "favourites" edges to the Favourite entity.
func (cc *CustomerCreate) AddFavourites(f ...*Favourite) *CustomerCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return cc.AddFavouriteIDs(ids...)
}

// AddNotificationIDs adds the "notifications" edge to the Notification entity by IDs.
func (cc *CustomerCreate) AddNotificationIDs(ids ...int) *CustomerCreate {
	cc.mutation.AddNotificationIDs(ids...)
	return cc
}

// AddNotifications adds the "notifications" edges to the Notification entity.
func (cc *CustomerCreate) AddNotifications(n ...*Notification) *CustomerCreate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return cc.AddNotificationIDs(ids...)
}

// AddPurchaseRequestIDs adds the "purchase_request" edge to the PurchaseRequest entity by IDs.
func (cc *CustomerCreate) AddPurchaseRequestIDs(ids ...int) *CustomerCreate {
	cc.mutation.AddPurchaseRequestIDs(ids...)
	return cc
}

// AddPurchaseRequest adds the "purchase_request" edges to the PurchaseRequest entity.
func (cc *CustomerCreate) AddPurchaseRequest(p ...*PurchaseRequest) *CustomerCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cc.AddPurchaseRequestIDs(ids...)
}

// SetAdminID sets the "admin" edge to the Admin entity by ID.
func (cc *CustomerCreate) SetAdminID(id int) *CustomerCreate {
	cc.mutation.SetAdminID(id)
	return cc
}

// SetNillableAdminID sets the "admin" edge to the Admin entity by ID if the given value is not nil.
func (cc *CustomerCreate) SetNillableAdminID(id *int) *CustomerCreate {
	if id != nil {
		cc = cc.SetAdminID(*id)
	}
	return cc
}

// SetAdmin sets the "admin" edge to the Admin entity.
func (cc *CustomerCreate) SetAdmin(a *Admin) *CustomerCreate {
	return cc.SetAdminID(a.ID)
}

// Mutation returns the CustomerMutation object of the builder.
func (cc *CustomerCreate) Mutation() *CustomerMutation {
	return cc.mutation
}

// Save creates the Customer in the database.
func (cc *CustomerCreate) Save(ctx context.Context) (*Customer, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CustomerCreate) SaveX(ctx context.Context) *Customer {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CustomerCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CustomerCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CustomerCreate) defaults() {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := customer.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := customer.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CustomerCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Customer.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Customer.updated_at"`)}
	}
	if _, ok := cc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "Customer.username"`)}
	}
	if v, ok := cc.mutation.Username(); ok {
		if err := customer.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "Customer.username": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "Customer.password"`)}
	}
	if v, ok := cc.mutation.Password(); ok {
		if err := customer.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Customer.password": %w`, err)}
		}
	}
	if _, ok := cc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Customer.type"`)}
	}
	if v, ok := cc.mutation.GetType(); ok {
		if err := customer.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Customer.type": %w`, err)}
		}
	}
	return nil
}

func (cc *CustomerCreate) sqlSave(ctx context.Context) (*Customer, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CustomerCreate) createSpec() (*Customer, *sqlgraph.CreateSpec) {
	var (
		_node = &Customer{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(customer.Table, sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt))
	)
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(customer.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(customer.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.Username(); ok {
		_spec.SetField(customer.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := cc.mutation.Password(); ok {
		_spec.SetField(customer.FieldPassword, field.TypeBytes, value)
		_node.Password = value
	}
	if value, ok := cc.mutation.GetType(); ok {
		_spec.SetField(customer.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if nodes := cc.mutation.BusinessIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   customer.BusinessTable,
			Columns: []string{customer.BusinessColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(businesscustomer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.IndividualIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   customer.IndividualTable,
			Columns: []string{customer.IndividualColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(individualcustomer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.AddressesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.AddressesTable,
			Columns: []string{customer.AddressesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.OrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.OrdersTable,
			Columns: []string{customer.OrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.FavouritesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.FavouritesTable,
			Columns: []string{customer.FavouritesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(favourite.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.NotificationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   customer.NotificationsTable,
			Columns: customer.NotificationsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notification.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.PurchaseRequestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.PurchaseRequestTable,
			Columns: []string{customer.PurchaseRequestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(purchaserequest.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.AdminIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   customer.AdminTable,
			Columns: []string{customer.AdminColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(admin.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.admin_customers = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CustomerCreateBulk is the builder for creating many Customer entities in bulk.
type CustomerCreateBulk struct {
	config
	builders []*CustomerCreate
}

// Save creates the Customer entities in the database.
func (ccb *CustomerCreateBulk) Save(ctx context.Context) ([]*Customer, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Customer, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CustomerMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CustomerCreateBulk) SaveX(ctx context.Context) []*Customer {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CustomerCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CustomerCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
