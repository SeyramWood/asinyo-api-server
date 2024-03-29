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
	"github.com/SeyramWood/ent/agent"
	"github.com/SeyramWood/ent/customer"
	"github.com/SeyramWood/ent/logistic"
	"github.com/SeyramWood/ent/merchant"
	"github.com/SeyramWood/ent/merchantstore"
	"github.com/SeyramWood/ent/order"
	"github.com/SeyramWood/ent/orderdetail"
	"github.com/SeyramWood/ent/pickupstation"
	"github.com/SeyramWood/ent/purchaserequest"
)

// OrderCreate is the builder for creating a Order entity.
type OrderCreate struct {
	config
	mutation *OrderMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (oc *OrderCreate) SetCreatedAt(t time.Time) *OrderCreate {
	oc.mutation.SetCreatedAt(t)
	return oc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (oc *OrderCreate) SetNillableCreatedAt(t *time.Time) *OrderCreate {
	if t != nil {
		oc.SetCreatedAt(*t)
	}
	return oc
}

// SetUpdatedAt sets the "updated_at" field.
func (oc *OrderCreate) SetUpdatedAt(t time.Time) *OrderCreate {
	oc.mutation.SetUpdatedAt(t)
	return oc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (oc *OrderCreate) SetNillableUpdatedAt(t *time.Time) *OrderCreate {
	if t != nil {
		oc.SetUpdatedAt(*t)
	}
	return oc
}

// SetOrderNumber sets the "order_number" field.
func (oc *OrderCreate) SetOrderNumber(s string) *OrderCreate {
	oc.mutation.SetOrderNumber(s)
	return oc
}

// SetCurrency sets the "currency" field.
func (oc *OrderCreate) SetCurrency(s string) *OrderCreate {
	oc.mutation.SetCurrency(s)
	return oc
}

// SetAmount sets the "amount" field.
func (oc *OrderCreate) SetAmount(f float64) *OrderCreate {
	oc.mutation.SetAmount(f)
	return oc
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (oc *OrderCreate) SetNillableAmount(f *float64) *OrderCreate {
	if f != nil {
		oc.SetAmount(*f)
	}
	return oc
}

// SetDeliveryFee sets the "delivery_fee" field.
func (oc *OrderCreate) SetDeliveryFee(f float64) *OrderCreate {
	oc.mutation.SetDeliveryFee(f)
	return oc
}

// SetNillableDeliveryFee sets the "delivery_fee" field if the given value is not nil.
func (oc *OrderCreate) SetNillableDeliveryFee(f *float64) *OrderCreate {
	if f != nil {
		oc.SetDeliveryFee(*f)
	}
	return oc
}

// SetReference sets the "reference" field.
func (oc *OrderCreate) SetReference(s string) *OrderCreate {
	oc.mutation.SetReference(s)
	return oc
}

// SetNillableReference sets the "reference" field if the given value is not nil.
func (oc *OrderCreate) SetNillableReference(s *string) *OrderCreate {
	if s != nil {
		oc.SetReference(*s)
	}
	return oc
}

// SetChannel sets the "channel" field.
func (oc *OrderCreate) SetChannel(s string) *OrderCreate {
	oc.mutation.SetChannel(s)
	return oc
}

// SetNillableChannel sets the "channel" field if the given value is not nil.
func (oc *OrderCreate) SetNillableChannel(s *string) *OrderCreate {
	if s != nil {
		oc.SetChannel(*s)
	}
	return oc
}

// SetPaidAt sets the "paid_at" field.
func (oc *OrderCreate) SetPaidAt(s string) *OrderCreate {
	oc.mutation.SetPaidAt(s)
	return oc
}

// SetNillablePaidAt sets the "paid_at" field if the given value is not nil.
func (oc *OrderCreate) SetNillablePaidAt(s *string) *OrderCreate {
	if s != nil {
		oc.SetPaidAt(*s)
	}
	return oc
}

// SetDeliveryMethod sets the "delivery_method" field.
func (oc *OrderCreate) SetDeliveryMethod(om order.DeliveryMethod) *OrderCreate {
	oc.mutation.SetDeliveryMethod(om)
	return oc
}

// SetPaymentMethod sets the "payment_method" field.
func (oc *OrderCreate) SetPaymentMethod(om order.PaymentMethod) *OrderCreate {
	oc.mutation.SetPaymentMethod(om)
	return oc
}

// SetNillablePaymentMethod sets the "payment_method" field if the given value is not nil.
func (oc *OrderCreate) SetNillablePaymentMethod(om *order.PaymentMethod) *OrderCreate {
	if om != nil {
		oc.SetPaymentMethod(*om)
	}
	return oc
}

// SetStatus sets the "status" field.
func (oc *OrderCreate) SetStatus(o order.Status) *OrderCreate {
	oc.mutation.SetStatus(o)
	return oc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (oc *OrderCreate) SetNillableStatus(o *order.Status) *OrderCreate {
	if o != nil {
		oc.SetStatus(*o)
	}
	return oc
}

// SetCustomerApproval sets the "customer_approval" field.
func (oc *OrderCreate) SetCustomerApproval(oa order.CustomerApproval) *OrderCreate {
	oc.mutation.SetCustomerApproval(oa)
	return oc
}

// SetNillableCustomerApproval sets the "customer_approval" field if the given value is not nil.
func (oc *OrderCreate) SetNillableCustomerApproval(oa *order.CustomerApproval) *OrderCreate {
	if oa != nil {
		oc.SetCustomerApproval(*oa)
	}
	return oc
}

// SetStoreTasksCreated sets the "store_tasks_created" field.
func (oc *OrderCreate) SetStoreTasksCreated(i []int) *OrderCreate {
	oc.mutation.SetStoreTasksCreated(i)
	return oc
}

// SetDeliveredAt sets the "delivered_at" field.
func (oc *OrderCreate) SetDeliveredAt(t time.Time) *OrderCreate {
	oc.mutation.SetDeliveredAt(t)
	return oc
}

// SetNillableDeliveredAt sets the "delivered_at" field if the given value is not nil.
func (oc *OrderCreate) SetNillableDeliveredAt(t *time.Time) *OrderCreate {
	if t != nil {
		oc.SetDeliveredAt(*t)
	}
	return oc
}

// AddDetailIDs adds the "details" edge to the OrderDetail entity by IDs.
func (oc *OrderCreate) AddDetailIDs(ids ...int) *OrderCreate {
	oc.mutation.AddDetailIDs(ids...)
	return oc
}

// AddDetails adds the "details" edges to the OrderDetail entity.
func (oc *OrderCreate) AddDetails(o ...*OrderDetail) *OrderCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oc.AddDetailIDs(ids...)
}

// SetLogisticID sets the "logistic" edge to the Logistic entity by ID.
func (oc *OrderCreate) SetLogisticID(id int) *OrderCreate {
	oc.mutation.SetLogisticID(id)
	return oc
}

// SetNillableLogisticID sets the "logistic" edge to the Logistic entity by ID if the given value is not nil.
func (oc *OrderCreate) SetNillableLogisticID(id *int) *OrderCreate {
	if id != nil {
		oc = oc.SetLogisticID(*id)
	}
	return oc
}

// SetLogistic sets the "logistic" edge to the Logistic entity.
func (oc *OrderCreate) SetLogistic(l *Logistic) *OrderCreate {
	return oc.SetLogisticID(l.ID)
}

// SetMerchantID sets the "merchant" edge to the Merchant entity by ID.
func (oc *OrderCreate) SetMerchantID(id int) *OrderCreate {
	oc.mutation.SetMerchantID(id)
	return oc
}

// SetNillableMerchantID sets the "merchant" edge to the Merchant entity by ID if the given value is not nil.
func (oc *OrderCreate) SetNillableMerchantID(id *int) *OrderCreate {
	if id != nil {
		oc = oc.SetMerchantID(*id)
	}
	return oc
}

// SetMerchant sets the "merchant" edge to the Merchant entity.
func (oc *OrderCreate) SetMerchant(m *Merchant) *OrderCreate {
	return oc.SetMerchantID(m.ID)
}

// SetAgentID sets the "agent" edge to the Agent entity by ID.
func (oc *OrderCreate) SetAgentID(id int) *OrderCreate {
	oc.mutation.SetAgentID(id)
	return oc
}

// SetNillableAgentID sets the "agent" edge to the Agent entity by ID if the given value is not nil.
func (oc *OrderCreate) SetNillableAgentID(id *int) *OrderCreate {
	if id != nil {
		oc = oc.SetAgentID(*id)
	}
	return oc
}

// SetAgent sets the "agent" edge to the Agent entity.
func (oc *OrderCreate) SetAgent(a *Agent) *OrderCreate {
	return oc.SetAgentID(a.ID)
}

// SetCustomerID sets the "customer" edge to the Customer entity by ID.
func (oc *OrderCreate) SetCustomerID(id int) *OrderCreate {
	oc.mutation.SetCustomerID(id)
	return oc
}

// SetNillableCustomerID sets the "customer" edge to the Customer entity by ID if the given value is not nil.
func (oc *OrderCreate) SetNillableCustomerID(id *int) *OrderCreate {
	if id != nil {
		oc = oc.SetCustomerID(*id)
	}
	return oc
}

// SetCustomer sets the "customer" edge to the Customer entity.
func (oc *OrderCreate) SetCustomer(c *Customer) *OrderCreate {
	return oc.SetCustomerID(c.ID)
}

// SetAddressID sets the "address" edge to the Address entity by ID.
func (oc *OrderCreate) SetAddressID(id int) *OrderCreate {
	oc.mutation.SetAddressID(id)
	return oc
}

// SetNillableAddressID sets the "address" edge to the Address entity by ID if the given value is not nil.
func (oc *OrderCreate) SetNillableAddressID(id *int) *OrderCreate {
	if id != nil {
		oc = oc.SetAddressID(*id)
	}
	return oc
}

// SetAddress sets the "address" edge to the Address entity.
func (oc *OrderCreate) SetAddress(a *Address) *OrderCreate {
	return oc.SetAddressID(a.ID)
}

// SetPickupID sets the "pickup" edge to the PickupStation entity by ID.
func (oc *OrderCreate) SetPickupID(id int) *OrderCreate {
	oc.mutation.SetPickupID(id)
	return oc
}

// SetNillablePickupID sets the "pickup" edge to the PickupStation entity by ID if the given value is not nil.
func (oc *OrderCreate) SetNillablePickupID(id *int) *OrderCreate {
	if id != nil {
		oc = oc.SetPickupID(*id)
	}
	return oc
}

// SetPickup sets the "pickup" edge to the PickupStation entity.
func (oc *OrderCreate) SetPickup(p *PickupStation) *OrderCreate {
	return oc.SetPickupID(p.ID)
}

// AddStoreIDs adds the "stores" edge to the MerchantStore entity by IDs.
func (oc *OrderCreate) AddStoreIDs(ids ...int) *OrderCreate {
	oc.mutation.AddStoreIDs(ids...)
	return oc
}

// AddStores adds the "stores" edges to the MerchantStore entity.
func (oc *OrderCreate) AddStores(m ...*MerchantStore) *OrderCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return oc.AddStoreIDs(ids...)
}

// SetPurchaseRequestID sets the "purchase_request" edge to the PurchaseRequest entity by ID.
func (oc *OrderCreate) SetPurchaseRequestID(id int) *OrderCreate {
	oc.mutation.SetPurchaseRequestID(id)
	return oc
}

// SetNillablePurchaseRequestID sets the "purchase_request" edge to the PurchaseRequest entity by ID if the given value is not nil.
func (oc *OrderCreate) SetNillablePurchaseRequestID(id *int) *OrderCreate {
	if id != nil {
		oc = oc.SetPurchaseRequestID(*id)
	}
	return oc
}

// SetPurchaseRequest sets the "purchase_request" edge to the PurchaseRequest entity.
func (oc *OrderCreate) SetPurchaseRequest(p *PurchaseRequest) *OrderCreate {
	return oc.SetPurchaseRequestID(p.ID)
}

// Mutation returns the OrderMutation object of the builder.
func (oc *OrderCreate) Mutation() *OrderMutation {
	return oc.mutation
}

// Save creates the Order in the database.
func (oc *OrderCreate) Save(ctx context.Context) (*Order, error) {
	oc.defaults()
	return withHooks(ctx, oc.sqlSave, oc.mutation, oc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OrderCreate) SaveX(ctx context.Context) *Order {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *OrderCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *OrderCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oc *OrderCreate) defaults() {
	if _, ok := oc.mutation.CreatedAt(); !ok {
		v := order.DefaultCreatedAt()
		oc.mutation.SetCreatedAt(v)
	}
	if _, ok := oc.mutation.UpdatedAt(); !ok {
		v := order.DefaultUpdatedAt()
		oc.mutation.SetUpdatedAt(v)
	}
	if _, ok := oc.mutation.Amount(); !ok {
		v := order.DefaultAmount
		oc.mutation.SetAmount(v)
	}
	if _, ok := oc.mutation.DeliveryFee(); !ok {
		v := order.DefaultDeliveryFee
		oc.mutation.SetDeliveryFee(v)
	}
	if _, ok := oc.mutation.PaymentMethod(); !ok {
		v := order.DefaultPaymentMethod
		oc.mutation.SetPaymentMethod(v)
	}
	if _, ok := oc.mutation.Status(); !ok {
		v := order.DefaultStatus
		oc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oc *OrderCreate) check() error {
	if _, ok := oc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Order.created_at"`)}
	}
	if _, ok := oc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Order.updated_at"`)}
	}
	if _, ok := oc.mutation.OrderNumber(); !ok {
		return &ValidationError{Name: "order_number", err: errors.New(`ent: missing required field "Order.order_number"`)}
	}
	if v, ok := oc.mutation.OrderNumber(); ok {
		if err := order.OrderNumberValidator(v); err != nil {
			return &ValidationError{Name: "order_number", err: fmt.Errorf(`ent: validator failed for field "Order.order_number": %w`, err)}
		}
	}
	if _, ok := oc.mutation.Currency(); !ok {
		return &ValidationError{Name: "currency", err: errors.New(`ent: missing required field "Order.currency"`)}
	}
	if v, ok := oc.mutation.Currency(); ok {
		if err := order.CurrencyValidator(v); err != nil {
			return &ValidationError{Name: "currency", err: fmt.Errorf(`ent: validator failed for field "Order.currency": %w`, err)}
		}
	}
	if _, ok := oc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "Order.amount"`)}
	}
	if _, ok := oc.mutation.DeliveryFee(); !ok {
		return &ValidationError{Name: "delivery_fee", err: errors.New(`ent: missing required field "Order.delivery_fee"`)}
	}
	if _, ok := oc.mutation.DeliveryMethod(); !ok {
		return &ValidationError{Name: "delivery_method", err: errors.New(`ent: missing required field "Order.delivery_method"`)}
	}
	if v, ok := oc.mutation.DeliveryMethod(); ok {
		if err := order.DeliveryMethodValidator(v); err != nil {
			return &ValidationError{Name: "delivery_method", err: fmt.Errorf(`ent: validator failed for field "Order.delivery_method": %w`, err)}
		}
	}
	if _, ok := oc.mutation.PaymentMethod(); !ok {
		return &ValidationError{Name: "payment_method", err: errors.New(`ent: missing required field "Order.payment_method"`)}
	}
	if v, ok := oc.mutation.PaymentMethod(); ok {
		if err := order.PaymentMethodValidator(v); err != nil {
			return &ValidationError{Name: "payment_method", err: fmt.Errorf(`ent: validator failed for field "Order.payment_method": %w`, err)}
		}
	}
	if _, ok := oc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Order.status"`)}
	}
	if v, ok := oc.mutation.Status(); ok {
		if err := order.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Order.status": %w`, err)}
		}
	}
	if v, ok := oc.mutation.CustomerApproval(); ok {
		if err := order.CustomerApprovalValidator(v); err != nil {
			return &ValidationError{Name: "customer_approval", err: fmt.Errorf(`ent: validator failed for field "Order.customer_approval": %w`, err)}
		}
	}
	return nil
}

func (oc *OrderCreate) sqlSave(ctx context.Context) (*Order, error) {
	if err := oc.check(); err != nil {
		return nil, err
	}
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	oc.mutation.id = &_node.ID
	oc.mutation.done = true
	return _node, nil
}

func (oc *OrderCreate) createSpec() (*Order, *sqlgraph.CreateSpec) {
	var (
		_node = &Order{config: oc.config}
		_spec = sqlgraph.NewCreateSpec(order.Table, sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt))
	)
	if value, ok := oc.mutation.CreatedAt(); ok {
		_spec.SetField(order.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := oc.mutation.UpdatedAt(); ok {
		_spec.SetField(order.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := oc.mutation.OrderNumber(); ok {
		_spec.SetField(order.FieldOrderNumber, field.TypeString, value)
		_node.OrderNumber = value
	}
	if value, ok := oc.mutation.Currency(); ok {
		_spec.SetField(order.FieldCurrency, field.TypeString, value)
		_node.Currency = value
	}
	if value, ok := oc.mutation.Amount(); ok {
		_spec.SetField(order.FieldAmount, field.TypeFloat64, value)
		_node.Amount = value
	}
	if value, ok := oc.mutation.DeliveryFee(); ok {
		_spec.SetField(order.FieldDeliveryFee, field.TypeFloat64, value)
		_node.DeliveryFee = value
	}
	if value, ok := oc.mutation.Reference(); ok {
		_spec.SetField(order.FieldReference, field.TypeString, value)
		_node.Reference = &value
	}
	if value, ok := oc.mutation.Channel(); ok {
		_spec.SetField(order.FieldChannel, field.TypeString, value)
		_node.Channel = &value
	}
	if value, ok := oc.mutation.PaidAt(); ok {
		_spec.SetField(order.FieldPaidAt, field.TypeString, value)
		_node.PaidAt = &value
	}
	if value, ok := oc.mutation.DeliveryMethod(); ok {
		_spec.SetField(order.FieldDeliveryMethod, field.TypeEnum, value)
		_node.DeliveryMethod = value
	}
	if value, ok := oc.mutation.PaymentMethod(); ok {
		_spec.SetField(order.FieldPaymentMethod, field.TypeEnum, value)
		_node.PaymentMethod = value
	}
	if value, ok := oc.mutation.Status(); ok {
		_spec.SetField(order.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := oc.mutation.CustomerApproval(); ok {
		_spec.SetField(order.FieldCustomerApproval, field.TypeEnum, value)
		_node.CustomerApproval = value
	}
	if value, ok := oc.mutation.StoreTasksCreated(); ok {
		_spec.SetField(order.FieldStoreTasksCreated, field.TypeJSON, value)
		_node.StoreTasksCreated = value
	}
	if value, ok := oc.mutation.DeliveredAt(); ok {
		_spec.SetField(order.FieldDeliveredAt, field.TypeTime, value)
		_node.DeliveredAt = &value
	}
	if nodes := oc.mutation.DetailsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.DetailsTable,
			Columns: []string{order.DetailsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderdetail.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.LogisticIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   order.LogisticTable,
			Columns: []string{order.LogisticColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(logistic.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.MerchantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.MerchantTable,
			Columns: []string{order.MerchantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(merchant.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.merchant_orders = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.AgentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.AgentTable,
			Columns: []string{order.AgentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(agent.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.agent_orders = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.CustomerTable,
			Columns: []string{order.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.customer_orders = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.AddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.AddressTable,
			Columns: []string{order.AddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.address_orders = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.PickupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.PickupTable,
			Columns: []string{order.PickupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pickupstation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.pickup_station_orders = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.StoresIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   order.StoresTable,
			Columns: order.StoresPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(merchantstore.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.PurchaseRequestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.PurchaseRequestTable,
			Columns: []string{order.PurchaseRequestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(purchaserequest.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.purchase_request_order = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrderCreateBulk is the builder for creating many Order entities in bulk.
type OrderCreateBulk struct {
	config
	builders []*OrderCreate
}

// Save creates the Order entities in the database.
func (ocb *OrderCreateBulk) Save(ctx context.Context) ([]*Order, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Order, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderMutation)
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
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *OrderCreateBulk) SaveX(ctx context.Context) []*Order {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *OrderCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *OrderCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}
