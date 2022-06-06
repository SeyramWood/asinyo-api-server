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
	"github.com/SeyramWood/ent/address"
	"github.com/SeyramWood/ent/agent"
	"github.com/SeyramWood/ent/customer"
	"github.com/SeyramWood/ent/merchant"
	"github.com/SeyramWood/ent/order"
	"github.com/SeyramWood/ent/predicate"
	"github.com/SeyramWood/ent/product"
)

// OrderUpdate is the builder for updating Order entities.
type OrderUpdate struct {
	config
	hooks    []Hook
	mutation *OrderMutation
}

// Where appends a list predicates to the OrderUpdate builder.
func (ou *OrderUpdate) Where(ps ...predicate.Order) *OrderUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetUpdatedAt sets the "updated_at" field.
func (ou *OrderUpdate) SetUpdatedAt(t time.Time) *OrderUpdate {
	ou.mutation.SetUpdatedAt(t)
	return ou
}

// SetStatus sets the "status" field.
func (ou *OrderUpdate) SetStatus(o order.Status) *OrderUpdate {
	ou.mutation.SetStatus(o)
	return ou
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableStatus(o *order.Status) *OrderUpdate {
	if o != nil {
		ou.SetStatus(*o)
	}
	return ou
}

// SetDeliveredAt sets the "delivered_at" field.
func (ou *OrderUpdate) SetDeliveredAt(t time.Time) *OrderUpdate {
	ou.mutation.SetDeliveredAt(t)
	return ou
}

// SetNillableDeliveredAt sets the "delivered_at" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableDeliveredAt(t *time.Time) *OrderUpdate {
	if t != nil {
		ou.SetDeliveredAt(*t)
	}
	return ou
}

// ClearDeliveredAt clears the value of the "delivered_at" field.
func (ou *OrderUpdate) ClearDeliveredAt() *OrderUpdate {
	ou.mutation.ClearDeliveredAt()
	return ou
}

// SetMerchantID sets the "merchant" edge to the Merchant entity by ID.
func (ou *OrderUpdate) SetMerchantID(id int) *OrderUpdate {
	ou.mutation.SetMerchantID(id)
	return ou
}

// SetNillableMerchantID sets the "merchant" edge to the Merchant entity by ID if the given value is not nil.
func (ou *OrderUpdate) SetNillableMerchantID(id *int) *OrderUpdate {
	if id != nil {
		ou = ou.SetMerchantID(*id)
	}
	return ou
}

// SetMerchant sets the "merchant" edge to the Merchant entity.
func (ou *OrderUpdate) SetMerchant(m *Merchant) *OrderUpdate {
	return ou.SetMerchantID(m.ID)
}

// SetAgentID sets the "agent" edge to the Agent entity by ID.
func (ou *OrderUpdate) SetAgentID(id int) *OrderUpdate {
	ou.mutation.SetAgentID(id)
	return ou
}

// SetNillableAgentID sets the "agent" edge to the Agent entity by ID if the given value is not nil.
func (ou *OrderUpdate) SetNillableAgentID(id *int) *OrderUpdate {
	if id != nil {
		ou = ou.SetAgentID(*id)
	}
	return ou
}

// SetAgent sets the "agent" edge to the Agent entity.
func (ou *OrderUpdate) SetAgent(a *Agent) *OrderUpdate {
	return ou.SetAgentID(a.ID)
}

// SetCustomerID sets the "customer" edge to the Customer entity by ID.
func (ou *OrderUpdate) SetCustomerID(id int) *OrderUpdate {
	ou.mutation.SetCustomerID(id)
	return ou
}

// SetNillableCustomerID sets the "customer" edge to the Customer entity by ID if the given value is not nil.
func (ou *OrderUpdate) SetNillableCustomerID(id *int) *OrderUpdate {
	if id != nil {
		ou = ou.SetCustomerID(*id)
	}
	return ou
}

// SetCustomer sets the "customer" edge to the Customer entity.
func (ou *OrderUpdate) SetCustomer(c *Customer) *OrderUpdate {
	return ou.SetCustomerID(c.ID)
}

// SetAddressID sets the "address" edge to the Address entity by ID.
func (ou *OrderUpdate) SetAddressID(id int) *OrderUpdate {
	ou.mutation.SetAddressID(id)
	return ou
}

// SetNillableAddressID sets the "address" edge to the Address entity by ID if the given value is not nil.
func (ou *OrderUpdate) SetNillableAddressID(id *int) *OrderUpdate {
	if id != nil {
		ou = ou.SetAddressID(*id)
	}
	return ou
}

// SetAddress sets the "address" edge to the Address entity.
func (ou *OrderUpdate) SetAddress(a *Address) *OrderUpdate {
	return ou.SetAddressID(a.ID)
}

// SetProductID sets the "product" edge to the Product entity by ID.
func (ou *OrderUpdate) SetProductID(id int) *OrderUpdate {
	ou.mutation.SetProductID(id)
	return ou
}

// SetNillableProductID sets the "product" edge to the Product entity by ID if the given value is not nil.
func (ou *OrderUpdate) SetNillableProductID(id *int) *OrderUpdate {
	if id != nil {
		ou = ou.SetProductID(*id)
	}
	return ou
}

// SetProduct sets the "product" edge to the Product entity.
func (ou *OrderUpdate) SetProduct(p *Product) *OrderUpdate {
	return ou.SetProductID(p.ID)
}

// Mutation returns the OrderMutation object of the builder.
func (ou *OrderUpdate) Mutation() *OrderMutation {
	return ou.mutation
}

// ClearMerchant clears the "merchant" edge to the Merchant entity.
func (ou *OrderUpdate) ClearMerchant() *OrderUpdate {
	ou.mutation.ClearMerchant()
	return ou
}

// ClearAgent clears the "agent" edge to the Agent entity.
func (ou *OrderUpdate) ClearAgent() *OrderUpdate {
	ou.mutation.ClearAgent()
	return ou
}

// ClearCustomer clears the "customer" edge to the Customer entity.
func (ou *OrderUpdate) ClearCustomer() *OrderUpdate {
	ou.mutation.ClearCustomer()
	return ou
}

// ClearAddress clears the "address" edge to the Address entity.
func (ou *OrderUpdate) ClearAddress() *OrderUpdate {
	ou.mutation.ClearAddress()
	return ou
}

// ClearProduct clears the "product" edge to the Product entity.
func (ou *OrderUpdate) ClearProduct() *OrderUpdate {
	ou.mutation.ClearProduct()
	return ou
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OrderUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ou.defaults()
	if len(ou.hooks) == 0 {
		if err = ou.check(); err != nil {
			return 0, err
		}
		affected, err = ou.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ou.check(); err != nil {
				return 0, err
			}
			ou.mutation = mutation
			affected, err = ou.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ou.hooks) - 1; i >= 0; i-- {
			if ou.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ou.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ou.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OrderUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OrderUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OrderUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ou *OrderUpdate) defaults() {
	if _, ok := ou.mutation.UpdatedAt(); !ok {
		v := order.UpdateDefaultUpdatedAt()
		ou.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ou *OrderUpdate) check() error {
	if v, ok := ou.mutation.Status(); ok {
		if err := order.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Order.status": %w`, err)}
		}
	}
	return nil
}

func (ou *OrderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   order.Table,
			Columns: order.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: order.FieldID,
			},
		},
	}
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: order.FieldUpdatedAt,
		})
	}
	if value, ok := ou.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: order.FieldStatus,
		})
	}
	if value, ok := ou.mutation.DeliveredAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: order.FieldDeliveredAt,
		})
	}
	if ou.mutation.DeliveredAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: order.FieldDeliveredAt,
		})
	}
	if ou.mutation.MerchantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.MerchantTable,
			Columns: []string{order.MerchantColumn},
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
	if nodes := ou.mutation.MerchantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.MerchantTable,
			Columns: []string{order.MerchantColumn},
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
	if ou.mutation.AgentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.AgentTable,
			Columns: []string{order.AgentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: agent.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.AgentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.AgentTable,
			Columns: []string{order.AgentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: agent.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ou.mutation.CustomerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.CustomerTable,
			Columns: []string{order.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: customer.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.CustomerTable,
			Columns: []string{order.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: customer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ou.mutation.AddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.AddressTable,
			Columns: []string{order.AddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: address.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.AddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.AddressTable,
			Columns: []string{order.AddressColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ou.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.ProductTable,
			Columns: []string{order.ProductColumn},
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
	if nodes := ou.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.ProductTable,
			Columns: []string{order.ProductColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// OrderUpdateOne is the builder for updating a single Order entity.
type OrderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrderMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ouo *OrderUpdateOne) SetUpdatedAt(t time.Time) *OrderUpdateOne {
	ouo.mutation.SetUpdatedAt(t)
	return ouo
}

// SetStatus sets the "status" field.
func (ouo *OrderUpdateOne) SetStatus(o order.Status) *OrderUpdateOne {
	ouo.mutation.SetStatus(o)
	return ouo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableStatus(o *order.Status) *OrderUpdateOne {
	if o != nil {
		ouo.SetStatus(*o)
	}
	return ouo
}

// SetDeliveredAt sets the "delivered_at" field.
func (ouo *OrderUpdateOne) SetDeliveredAt(t time.Time) *OrderUpdateOne {
	ouo.mutation.SetDeliveredAt(t)
	return ouo
}

// SetNillableDeliveredAt sets the "delivered_at" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableDeliveredAt(t *time.Time) *OrderUpdateOne {
	if t != nil {
		ouo.SetDeliveredAt(*t)
	}
	return ouo
}

// ClearDeliveredAt clears the value of the "delivered_at" field.
func (ouo *OrderUpdateOne) ClearDeliveredAt() *OrderUpdateOne {
	ouo.mutation.ClearDeliveredAt()
	return ouo
}

// SetMerchantID sets the "merchant" edge to the Merchant entity by ID.
func (ouo *OrderUpdateOne) SetMerchantID(id int) *OrderUpdateOne {
	ouo.mutation.SetMerchantID(id)
	return ouo
}

// SetNillableMerchantID sets the "merchant" edge to the Merchant entity by ID if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableMerchantID(id *int) *OrderUpdateOne {
	if id != nil {
		ouo = ouo.SetMerchantID(*id)
	}
	return ouo
}

// SetMerchant sets the "merchant" edge to the Merchant entity.
func (ouo *OrderUpdateOne) SetMerchant(m *Merchant) *OrderUpdateOne {
	return ouo.SetMerchantID(m.ID)
}

// SetAgentID sets the "agent" edge to the Agent entity by ID.
func (ouo *OrderUpdateOne) SetAgentID(id int) *OrderUpdateOne {
	ouo.mutation.SetAgentID(id)
	return ouo
}

// SetNillableAgentID sets the "agent" edge to the Agent entity by ID if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableAgentID(id *int) *OrderUpdateOne {
	if id != nil {
		ouo = ouo.SetAgentID(*id)
	}
	return ouo
}

// SetAgent sets the "agent" edge to the Agent entity.
func (ouo *OrderUpdateOne) SetAgent(a *Agent) *OrderUpdateOne {
	return ouo.SetAgentID(a.ID)
}

// SetCustomerID sets the "customer" edge to the Customer entity by ID.
func (ouo *OrderUpdateOne) SetCustomerID(id int) *OrderUpdateOne {
	ouo.mutation.SetCustomerID(id)
	return ouo
}

// SetNillableCustomerID sets the "customer" edge to the Customer entity by ID if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableCustomerID(id *int) *OrderUpdateOne {
	if id != nil {
		ouo = ouo.SetCustomerID(*id)
	}
	return ouo
}

// SetCustomer sets the "customer" edge to the Customer entity.
func (ouo *OrderUpdateOne) SetCustomer(c *Customer) *OrderUpdateOne {
	return ouo.SetCustomerID(c.ID)
}

// SetAddressID sets the "address" edge to the Address entity by ID.
func (ouo *OrderUpdateOne) SetAddressID(id int) *OrderUpdateOne {
	ouo.mutation.SetAddressID(id)
	return ouo
}

// SetNillableAddressID sets the "address" edge to the Address entity by ID if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableAddressID(id *int) *OrderUpdateOne {
	if id != nil {
		ouo = ouo.SetAddressID(*id)
	}
	return ouo
}

// SetAddress sets the "address" edge to the Address entity.
func (ouo *OrderUpdateOne) SetAddress(a *Address) *OrderUpdateOne {
	return ouo.SetAddressID(a.ID)
}

// SetProductID sets the "product" edge to the Product entity by ID.
func (ouo *OrderUpdateOne) SetProductID(id int) *OrderUpdateOne {
	ouo.mutation.SetProductID(id)
	return ouo
}

// SetNillableProductID sets the "product" edge to the Product entity by ID if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableProductID(id *int) *OrderUpdateOne {
	if id != nil {
		ouo = ouo.SetProductID(*id)
	}
	return ouo
}

// SetProduct sets the "product" edge to the Product entity.
func (ouo *OrderUpdateOne) SetProduct(p *Product) *OrderUpdateOne {
	return ouo.SetProductID(p.ID)
}

// Mutation returns the OrderMutation object of the builder.
func (ouo *OrderUpdateOne) Mutation() *OrderMutation {
	return ouo.mutation
}

// ClearMerchant clears the "merchant" edge to the Merchant entity.
func (ouo *OrderUpdateOne) ClearMerchant() *OrderUpdateOne {
	ouo.mutation.ClearMerchant()
	return ouo
}

// ClearAgent clears the "agent" edge to the Agent entity.
func (ouo *OrderUpdateOne) ClearAgent() *OrderUpdateOne {
	ouo.mutation.ClearAgent()
	return ouo
}

// ClearCustomer clears the "customer" edge to the Customer entity.
func (ouo *OrderUpdateOne) ClearCustomer() *OrderUpdateOne {
	ouo.mutation.ClearCustomer()
	return ouo
}

// ClearAddress clears the "address" edge to the Address entity.
func (ouo *OrderUpdateOne) ClearAddress() *OrderUpdateOne {
	ouo.mutation.ClearAddress()
	return ouo
}

// ClearProduct clears the "product" edge to the Product entity.
func (ouo *OrderUpdateOne) ClearProduct() *OrderUpdateOne {
	ouo.mutation.ClearProduct()
	return ouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OrderUpdateOne) Select(field string, fields ...string) *OrderUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Order entity.
func (ouo *OrderUpdateOne) Save(ctx context.Context) (*Order, error) {
	var (
		err  error
		node *Order
	)
	ouo.defaults()
	if len(ouo.hooks) == 0 {
		if err = ouo.check(); err != nil {
			return nil, err
		}
		node, err = ouo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ouo.check(); err != nil {
				return nil, err
			}
			ouo.mutation = mutation
			node, err = ouo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ouo.hooks) - 1; i >= 0; i-- {
			if ouo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ouo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ouo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OrderUpdateOne) SaveX(ctx context.Context) *Order {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OrderUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OrderUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ouo *OrderUpdateOne) defaults() {
	if _, ok := ouo.mutation.UpdatedAt(); !ok {
		v := order.UpdateDefaultUpdatedAt()
		ouo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ouo *OrderUpdateOne) check() error {
	if v, ok := ouo.mutation.Status(); ok {
		if err := order.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Order.status": %w`, err)}
		}
	}
	return nil
}

func (ouo *OrderUpdateOne) sqlSave(ctx context.Context) (_node *Order, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   order.Table,
			Columns: order.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: order.FieldID,
			},
		},
	}
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Order.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, order.FieldID)
		for _, f := range fields {
			if !order.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != order.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: order.FieldUpdatedAt,
		})
	}
	if value, ok := ouo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: order.FieldStatus,
		})
	}
	if value, ok := ouo.mutation.DeliveredAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: order.FieldDeliveredAt,
		})
	}
	if ouo.mutation.DeliveredAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: order.FieldDeliveredAt,
		})
	}
	if ouo.mutation.MerchantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.MerchantTable,
			Columns: []string{order.MerchantColumn},
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
	if nodes := ouo.mutation.MerchantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.MerchantTable,
			Columns: []string{order.MerchantColumn},
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
	if ouo.mutation.AgentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.AgentTable,
			Columns: []string{order.AgentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: agent.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.AgentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.AgentTable,
			Columns: []string{order.AgentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: agent.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ouo.mutation.CustomerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.CustomerTable,
			Columns: []string{order.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: customer.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.CustomerTable,
			Columns: []string{order.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: customer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ouo.mutation.AddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.AddressTable,
			Columns: []string{order.AddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: address.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.AddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.AddressTable,
			Columns: []string{order.AddressColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ouo.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.ProductTable,
			Columns: []string{order.ProductColumn},
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
	if nodes := ouo.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.ProductTable,
			Columns: []string{order.ProductColumn},
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
	_node = &Order{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
