// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/domain/services"
	"github.com/SeyramWood/ent/agent"
	"github.com/SeyramWood/ent/agentrequest"
	"github.com/SeyramWood/ent/merchant"
	"github.com/SeyramWood/ent/merchantstore"
	"github.com/SeyramWood/ent/order"
	"github.com/SeyramWood/ent/orderdetail"
)

// MerchantStoreCreate is the builder for creating a MerchantStore entity.
type MerchantStoreCreate struct {
	config
	mutation *MerchantStoreMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (msc *MerchantStoreCreate) SetCreatedAt(t time.Time) *MerchantStoreCreate {
	msc.mutation.SetCreatedAt(t)
	return msc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (msc *MerchantStoreCreate) SetNillableCreatedAt(t *time.Time) *MerchantStoreCreate {
	if t != nil {
		msc.SetCreatedAt(*t)
	}
	return msc
}

// SetUpdatedAt sets the "updated_at" field.
func (msc *MerchantStoreCreate) SetUpdatedAt(t time.Time) *MerchantStoreCreate {
	msc.mutation.SetUpdatedAt(t)
	return msc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (msc *MerchantStoreCreate) SetNillableUpdatedAt(t *time.Time) *MerchantStoreCreate {
	if t != nil {
		msc.SetUpdatedAt(*t)
	}
	return msc
}

// SetName sets the "name" field.
func (msc *MerchantStoreCreate) SetName(s string) *MerchantStoreCreate {
	msc.mutation.SetName(s)
	return msc
}

// SetAbout sets the "about" field.
func (msc *MerchantStoreCreate) SetAbout(s string) *MerchantStoreCreate {
	msc.mutation.SetAbout(s)
	return msc
}

// SetSlogan sets the "slogan" field.
func (msc *MerchantStoreCreate) SetSlogan(s string) *MerchantStoreCreate {
	msc.mutation.SetSlogan(s)
	return msc
}

// SetDescription sets the "description" field.
func (msc *MerchantStoreCreate) SetDescription(s string) *MerchantStoreCreate {
	msc.mutation.SetDescription(s)
	return msc
}

// SetLogo sets the "logo" field.
func (msc *MerchantStoreCreate) SetLogo(s string) *MerchantStoreCreate {
	msc.mutation.SetLogo(s)
	return msc
}

// SetImages sets the "images" field.
func (msc *MerchantStoreCreate) SetImages(s []string) *MerchantStoreCreate {
	msc.mutation.SetImages(s)
	return msc
}

// SetDefaultAccount sets the "default_account" field.
func (msc *MerchantStoreCreate) SetDefaultAccount(ma merchantstore.DefaultAccount) *MerchantStoreCreate {
	msc.mutation.SetDefaultAccount(ma)
	return msc
}

// SetNillableDefaultAccount sets the "default_account" field if the given value is not nil.
func (msc *MerchantStoreCreate) SetNillableDefaultAccount(ma *merchantstore.DefaultAccount) *MerchantStoreCreate {
	if ma != nil {
		msc.SetDefaultAccount(*ma)
	}
	return msc
}

// SetBankAccount sets the "bank_account" field.
func (msc *MerchantStoreCreate) SetBankAccount(mba *models.MerchantBankAccount) *MerchantStoreCreate {
	msc.mutation.SetBankAccount(mba)
	return msc
}

// SetMomoAccount sets the "momo_account" field.
func (msc *MerchantStoreCreate) SetMomoAccount(mma *models.MerchantMomoAccount) *MerchantStoreCreate {
	msc.mutation.SetMomoAccount(mma)
	return msc
}

// SetAddress sets the "address" field.
func (msc *MerchantStoreCreate) SetAddress(msa *models.MerchantStoreAddress) *MerchantStoreCreate {
	msc.mutation.SetAddress(msa)
	return msc
}

// SetCoordinate sets the "coordinate" field.
func (msc *MerchantStoreCreate) SetCoordinate(s *services.Coordinate) *MerchantStoreCreate {
	msc.mutation.SetCoordinate(s)
	return msc
}

// SetMerchantType sets the "merchant_type" field.
func (msc *MerchantStoreCreate) SetMerchantType(s string) *MerchantStoreCreate {
	msc.mutation.SetMerchantType(s)
	return msc
}

// SetPermitAgent sets the "permit_agent" field.
func (msc *MerchantStoreCreate) SetPermitAgent(b bool) *MerchantStoreCreate {
	msc.mutation.SetPermitAgent(b)
	return msc
}

// SetNillablePermitAgent sets the "permit_agent" field if the given value is not nil.
func (msc *MerchantStoreCreate) SetNillablePermitAgent(b *bool) *MerchantStoreCreate {
	if b != nil {
		msc.SetPermitAgent(*b)
	}
	return msc
}

// SetMerchantID sets the "merchant" edge to the Merchant entity by ID.
func (msc *MerchantStoreCreate) SetMerchantID(id int) *MerchantStoreCreate {
	msc.mutation.SetMerchantID(id)
	return msc
}

// SetNillableMerchantID sets the "merchant" edge to the Merchant entity by ID if the given value is not nil.
func (msc *MerchantStoreCreate) SetNillableMerchantID(id *int) *MerchantStoreCreate {
	if id != nil {
		msc = msc.SetMerchantID(*id)
	}
	return msc
}

// SetMerchant sets the "merchant" edge to the Merchant entity.
func (msc *MerchantStoreCreate) SetMerchant(m *Merchant) *MerchantStoreCreate {
	return msc.SetMerchantID(m.ID)
}

// SetAgentID sets the "agent" edge to the Agent entity by ID.
func (msc *MerchantStoreCreate) SetAgentID(id int) *MerchantStoreCreate {
	msc.mutation.SetAgentID(id)
	return msc
}

// SetNillableAgentID sets the "agent" edge to the Agent entity by ID if the given value is not nil.
func (msc *MerchantStoreCreate) SetNillableAgentID(id *int) *MerchantStoreCreate {
	if id != nil {
		msc = msc.SetAgentID(*id)
	}
	return msc
}

// SetAgent sets the "agent" edge to the Agent entity.
func (msc *MerchantStoreCreate) SetAgent(a *Agent) *MerchantStoreCreate {
	return msc.SetAgentID(a.ID)
}

// AddRequestIDs adds the "requests" edge to the AgentRequest entity by IDs.
func (msc *MerchantStoreCreate) AddRequestIDs(ids ...int) *MerchantStoreCreate {
	msc.mutation.AddRequestIDs(ids...)
	return msc
}

// AddRequests adds the "requests" edges to the AgentRequest entity.
func (msc *MerchantStoreCreate) AddRequests(a ...*AgentRequest) *MerchantStoreCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return msc.AddRequestIDs(ids...)
}

// AddOrderIDs adds the "orders" edge to the Order entity by IDs.
func (msc *MerchantStoreCreate) AddOrderIDs(ids ...int) *MerchantStoreCreate {
	msc.mutation.AddOrderIDs(ids...)
	return msc
}

// AddOrders adds the "orders" edges to the Order entity.
func (msc *MerchantStoreCreate) AddOrders(o ...*Order) *MerchantStoreCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return msc.AddOrderIDs(ids...)
}

// AddOrderDetailIDs adds the "order_details" edge to the OrderDetail entity by IDs.
func (msc *MerchantStoreCreate) AddOrderDetailIDs(ids ...int) *MerchantStoreCreate {
	msc.mutation.AddOrderDetailIDs(ids...)
	return msc
}

// AddOrderDetails adds the "order_details" edges to the OrderDetail entity.
func (msc *MerchantStoreCreate) AddOrderDetails(o ...*OrderDetail) *MerchantStoreCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return msc.AddOrderDetailIDs(ids...)
}

// Mutation returns the MerchantStoreMutation object of the builder.
func (msc *MerchantStoreCreate) Mutation() *MerchantStoreMutation {
	return msc.mutation
}

// Save creates the MerchantStore in the database.
func (msc *MerchantStoreCreate) Save(ctx context.Context) (*MerchantStore, error) {
	var (
		err  error
		node *MerchantStore
	)
	msc.defaults()
	if len(msc.hooks) == 0 {
		if err = msc.check(); err != nil {
			return nil, err
		}
		node, err = msc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MerchantStoreMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = msc.check(); err != nil {
				return nil, err
			}
			msc.mutation = mutation
			if node, err = msc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(msc.hooks) - 1; i >= 0; i-- {
			if msc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = msc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, msc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*MerchantStore)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from MerchantStoreMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (msc *MerchantStoreCreate) SaveX(ctx context.Context) *MerchantStore {
	v, err := msc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (msc *MerchantStoreCreate) Exec(ctx context.Context) error {
	_, err := msc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (msc *MerchantStoreCreate) ExecX(ctx context.Context) {
	if err := msc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (msc *MerchantStoreCreate) defaults() {
	if _, ok := msc.mutation.CreatedAt(); !ok {
		v := merchantstore.DefaultCreatedAt()
		msc.mutation.SetCreatedAt(v)
	}
	if _, ok := msc.mutation.UpdatedAt(); !ok {
		v := merchantstore.DefaultUpdatedAt()
		msc.mutation.SetUpdatedAt(v)
	}
	if _, ok := msc.mutation.PermitAgent(); !ok {
		v := merchantstore.DefaultPermitAgent
		msc.mutation.SetPermitAgent(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (msc *MerchantStoreCreate) check() error {
	if _, ok := msc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "MerchantStore.created_at"`)}
	}
	if _, ok := msc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "MerchantStore.updated_at"`)}
	}
	if _, ok := msc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "MerchantStore.name"`)}
	}
	if v, ok := msc.mutation.Name(); ok {
		if err := merchantstore.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "MerchantStore.name": %w`, err)}
		}
	}
	if _, ok := msc.mutation.About(); !ok {
		return &ValidationError{Name: "about", err: errors.New(`ent: missing required field "MerchantStore.about"`)}
	}
	if v, ok := msc.mutation.About(); ok {
		if err := merchantstore.AboutValidator(v); err != nil {
			return &ValidationError{Name: "about", err: fmt.Errorf(`ent: validator failed for field "MerchantStore.about": %w`, err)}
		}
	}
	if _, ok := msc.mutation.Slogan(); !ok {
		return &ValidationError{Name: "slogan", err: errors.New(`ent: missing required field "MerchantStore.slogan"`)}
	}
	if v, ok := msc.mutation.Slogan(); ok {
		if err := merchantstore.SloganValidator(v); err != nil {
			return &ValidationError{Name: "slogan", err: fmt.Errorf(`ent: validator failed for field "MerchantStore.slogan": %w`, err)}
		}
	}
	if _, ok := msc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "MerchantStore.description"`)}
	}
	if v, ok := msc.mutation.Description(); ok {
		if err := merchantstore.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "MerchantStore.description": %w`, err)}
		}
	}
	if _, ok := msc.mutation.Logo(); !ok {
		return &ValidationError{Name: "logo", err: errors.New(`ent: missing required field "MerchantStore.logo"`)}
	}
	if v, ok := msc.mutation.Logo(); ok {
		if err := merchantstore.LogoValidator(v); err != nil {
			return &ValidationError{Name: "logo", err: fmt.Errorf(`ent: validator failed for field "MerchantStore.logo": %w`, err)}
		}
	}
	if v, ok := msc.mutation.DefaultAccount(); ok {
		if err := merchantstore.DefaultAccountValidator(v); err != nil {
			return &ValidationError{Name: "default_account", err: fmt.Errorf(`ent: validator failed for field "MerchantStore.default_account": %w`, err)}
		}
	}
	if _, ok := msc.mutation.MerchantType(); !ok {
		return &ValidationError{Name: "merchant_type", err: errors.New(`ent: missing required field "MerchantStore.merchant_type"`)}
	}
	if v, ok := msc.mutation.MerchantType(); ok {
		if err := merchantstore.MerchantTypeValidator(v); err != nil {
			return &ValidationError{Name: "merchant_type", err: fmt.Errorf(`ent: validator failed for field "MerchantStore.merchant_type": %w`, err)}
		}
	}
	if _, ok := msc.mutation.PermitAgent(); !ok {
		return &ValidationError{Name: "permit_agent", err: errors.New(`ent: missing required field "MerchantStore.permit_agent"`)}
	}
	return nil
}

func (msc *MerchantStoreCreate) sqlSave(ctx context.Context) (*MerchantStore, error) {
	_node, _spec := msc.createSpec()
	if err := sqlgraph.CreateNode(ctx, msc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (msc *MerchantStoreCreate) createSpec() (*MerchantStore, *sqlgraph.CreateSpec) {
	var (
		_node = &MerchantStore{config: msc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: merchantstore.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: merchantstore.FieldID,
			},
		}
	)
	if value, ok := msc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: merchantstore.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := msc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: merchantstore.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := msc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchantstore.FieldName,
		})
		_node.Name = value
	}
	if value, ok := msc.mutation.About(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchantstore.FieldAbout,
		})
		_node.About = value
	}
	if value, ok := msc.mutation.Slogan(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchantstore.FieldSlogan,
		})
		_node.Slogan = value
	}
	if value, ok := msc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchantstore.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := msc.mutation.Logo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchantstore.FieldLogo,
		})
		_node.Logo = value
	}
	if value, ok := msc.mutation.Images(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: merchantstore.FieldImages,
		})
		_node.Images = value
	}
	if value, ok := msc.mutation.DefaultAccount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: merchantstore.FieldDefaultAccount,
		})
		_node.DefaultAccount = value
	}
	if value, ok := msc.mutation.BankAccount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: merchantstore.FieldBankAccount,
		})
		_node.BankAccount = value
	}
	if value, ok := msc.mutation.MomoAccount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: merchantstore.FieldMomoAccount,
		})
		_node.MomoAccount = value
	}
	if value, ok := msc.mutation.Address(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: merchantstore.FieldAddress,
		})
		_node.Address = value
	}
	if value, ok := msc.mutation.Coordinate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: merchantstore.FieldCoordinate,
		})
		_node.Coordinate = value
	}
	if value, ok := msc.mutation.MerchantType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchantstore.FieldMerchantType,
		})
		_node.MerchantType = value
	}
	if value, ok := msc.mutation.PermitAgent(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: merchantstore.FieldPermitAgent,
		})
		_node.PermitAgent = value
	}
	if nodes := msc.mutation.MerchantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   merchantstore.MerchantTable,
			Columns: []string{merchantstore.MerchantColumn},
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
		_node.merchant_store = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := msc.mutation.AgentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   merchantstore.AgentTable,
			Columns: []string{merchantstore.AgentColumn},
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
		_node.agent_store = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := msc.mutation.RequestsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   merchantstore.RequestsTable,
			Columns: []string{merchantstore.RequestsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: agentrequest.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := msc.mutation.OrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   merchantstore.OrdersTable,
			Columns: merchantstore.OrdersPrimaryKey,
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
	if nodes := msc.mutation.OrderDetailsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   merchantstore.OrderDetailsTable,
			Columns: []string{merchantstore.OrderDetailsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderdetail.FieldID,
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

// MerchantStoreCreateBulk is the builder for creating many MerchantStore entities in bulk.
type MerchantStoreCreateBulk struct {
	config
	builders []*MerchantStoreCreate
}

// Save creates the MerchantStore entities in the database.
func (mscb *MerchantStoreCreateBulk) Save(ctx context.Context) ([]*MerchantStore, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mscb.builders))
	nodes := make([]*MerchantStore, len(mscb.builders))
	mutators := make([]Mutator, len(mscb.builders))
	for i := range mscb.builders {
		func(i int, root context.Context) {
			builder := mscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MerchantStoreMutation)
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
					_, err = mutators[i+1].Mutate(root, mscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mscb *MerchantStoreCreateBulk) SaveX(ctx context.Context) []*MerchantStore {
	v, err := mscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mscb *MerchantStoreCreateBulk) Exec(ctx context.Context) error {
	_, err := mscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mscb *MerchantStoreCreateBulk) ExecX(ctx context.Context) {
	if err := mscb.Exec(ctx); err != nil {
		panic(err)
	}
}
