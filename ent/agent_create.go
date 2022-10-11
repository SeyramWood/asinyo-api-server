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
	"github.com/SeyramWood/ent/address"
	"github.com/SeyramWood/ent/agent"
	"github.com/SeyramWood/ent/agentrequest"
	"github.com/SeyramWood/ent/favourite"
	"github.com/SeyramWood/ent/merchantstore"
	"github.com/SeyramWood/ent/order"
)

// AgentCreate is the builder for creating a Agent entity.
type AgentCreate struct {
	config
	mutation *AgentMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ac *AgentCreate) SetCreatedAt(t time.Time) *AgentCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AgentCreate) SetNillableCreatedAt(t *time.Time) *AgentCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AgentCreate) SetUpdatedAt(t time.Time) *AgentCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AgentCreate) SetNillableUpdatedAt(t *time.Time) *AgentCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetUsername sets the "username" field.
func (ac *AgentCreate) SetUsername(s string) *AgentCreate {
	ac.mutation.SetUsername(s)
	return ac
}

// SetPassword sets the "password" field.
func (ac *AgentCreate) SetPassword(b []byte) *AgentCreate {
	ac.mutation.SetPassword(b)
	return ac
}

// SetGhanaCard sets the "ghana_card" field.
func (ac *AgentCreate) SetGhanaCard(s string) *AgentCreate {
	ac.mutation.SetGhanaCard(s)
	return ac
}

// SetLastName sets the "last_name" field.
func (ac *AgentCreate) SetLastName(s string) *AgentCreate {
	ac.mutation.SetLastName(s)
	return ac
}

// SetOtherName sets the "other_name" field.
func (ac *AgentCreate) SetOtherName(s string) *AgentCreate {
	ac.mutation.SetOtherName(s)
	return ac
}

// SetPhone sets the "phone" field.
func (ac *AgentCreate) SetPhone(s string) *AgentCreate {
	ac.mutation.SetPhone(s)
	return ac
}

// SetOtherPhone sets the "other_phone" field.
func (ac *AgentCreate) SetOtherPhone(s string) *AgentCreate {
	ac.mutation.SetOtherPhone(s)
	return ac
}

// SetNillableOtherPhone sets the "other_phone" field if the given value is not nil.
func (ac *AgentCreate) SetNillableOtherPhone(s *string) *AgentCreate {
	if s != nil {
		ac.SetOtherPhone(*s)
	}
	return ac
}

// SetAddress sets the "address" field.
func (ac *AgentCreate) SetAddress(s string) *AgentCreate {
	ac.mutation.SetAddress(s)
	return ac
}

// SetDigitalAddress sets the "digital_address" field.
func (ac *AgentCreate) SetDigitalAddress(s string) *AgentCreate {
	ac.mutation.SetDigitalAddress(s)
	return ac
}

// SetRegion sets the "region" field.
func (ac *AgentCreate) SetRegion(s string) *AgentCreate {
	ac.mutation.SetRegion(s)
	return ac
}

<<<<<<< HEAD
=======
// SetNillableRegion sets the "region" field if the given value is not nil.
func (ac *AgentCreate) SetNillableRegion(s *string) *AgentCreate {
	if s != nil {
		ac.SetRegion(*s)
	}
	return ac
}

>>>>>>> dev
// SetDistrict sets the "district" field.
func (ac *AgentCreate) SetDistrict(s string) *AgentCreate {
	ac.mutation.SetDistrict(s)
	return ac
}

<<<<<<< HEAD
=======
// SetNillableDistrict sets the "district" field if the given value is not nil.
func (ac *AgentCreate) SetNillableDistrict(s *string) *AgentCreate {
	if s != nil {
		ac.SetDistrict(*s)
	}
	return ac
}

>>>>>>> dev
// SetCity sets the "city" field.
func (ac *AgentCreate) SetCity(s string) *AgentCreate {
	ac.mutation.SetCity(s)
	return ac
}

<<<<<<< HEAD
=======
// SetNillableCity sets the "city" field if the given value is not nil.
func (ac *AgentCreate) SetNillableCity(s *string) *AgentCreate {
	if s != nil {
		ac.SetCity(*s)
	}
	return ac
}

>>>>>>> dev
// SetDefaultAccount sets the "default_account" field.
func (ac *AgentCreate) SetDefaultAccount(aa agent.DefaultAccount) *AgentCreate {
	ac.mutation.SetDefaultAccount(aa)
	return ac
}

// SetNillableDefaultAccount sets the "default_account" field if the given value is not nil.
func (ac *AgentCreate) SetNillableDefaultAccount(aa *agent.DefaultAccount) *AgentCreate {
	if aa != nil {
		ac.SetDefaultAccount(*aa)
	}
	return ac
}

// SetBankAccount sets the "bank_account" field.
func (ac *AgentCreate) SetBankAccount(mba *models.MerchantBankAccount) *AgentCreate {
	ac.mutation.SetBankAccount(mba)
	return ac
}

// SetMomoAccount sets the "momo_account" field.
func (ac *AgentCreate) SetMomoAccount(mma *models.MerchantMomoAccount) *AgentCreate {
	ac.mutation.SetMomoAccount(mma)
	return ac
}

// SetVerified sets the "verified" field.
func (ac *AgentCreate) SetVerified(b bool) *AgentCreate {
	ac.mutation.SetVerified(b)
	return ac
}

// SetNillableVerified sets the "verified" field if the given value is not nil.
func (ac *AgentCreate) SetNillableVerified(b *bool) *AgentCreate {
	if b != nil {
		ac.SetVerified(*b)
	}
	return ac
}

// SetCompliance sets the "compliance" field.
func (ac *AgentCreate) SetCompliance(mcm *models.AgentComplianceModel) *AgentCreate {
	ac.mutation.SetCompliance(mcm)
	return ac
}

// AddAddressIDs adds the "addresses" edge to the Address entity by IDs.
func (ac *AgentCreate) AddAddressIDs(ids ...int) *AgentCreate {
	ac.mutation.AddAddressIDs(ids...)
	return ac
}

// AddAddresses adds the "addresses" edges to the Address entity.
func (ac *AgentCreate) AddAddresses(a ...*Address) *AgentCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ac.AddAddressIDs(ids...)
}

// AddOrderIDs adds the "orders" edge to the Order entity by IDs.
func (ac *AgentCreate) AddOrderIDs(ids ...int) *AgentCreate {
	ac.mutation.AddOrderIDs(ids...)
	return ac
}

// AddOrders adds the "orders" edges to the Order entity.
func (ac *AgentCreate) AddOrders(o ...*Order) *AgentCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ac.AddOrderIDs(ids...)
}

// AddFavouriteIDs adds the "favourites" edge to the Favourite entity by IDs.
func (ac *AgentCreate) AddFavouriteIDs(ids ...int) *AgentCreate {
	ac.mutation.AddFavouriteIDs(ids...)
	return ac
}

// AddFavourites adds the "favourites" edges to the Favourite entity.
func (ac *AgentCreate) AddFavourites(f ...*Favourite) *AgentCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ac.AddFavouriteIDs(ids...)
}

// AddStoreIDs adds the "store" edge to the MerchantStore entity by IDs.
func (ac *AgentCreate) AddStoreIDs(ids ...int) *AgentCreate {
	ac.mutation.AddStoreIDs(ids...)
	return ac
}

// AddStore adds the "store" edges to the MerchantStore entity.
func (ac *AgentCreate) AddStore(m ...*MerchantStore) *AgentCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return ac.AddStoreIDs(ids...)
}

// AddRequestIDs adds the "requests" edge to the AgentRequest entity by IDs.
func (ac *AgentCreate) AddRequestIDs(ids ...int) *AgentCreate {
	ac.mutation.AddRequestIDs(ids...)
	return ac
}

// AddRequests adds the "requests" edges to the AgentRequest entity.
func (ac *AgentCreate) AddRequests(a ...*AgentRequest) *AgentCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ac.AddRequestIDs(ids...)
}

// Mutation returns the AgentMutation object of the builder.
func (ac *AgentCreate) Mutation() *AgentMutation {
	return ac.mutation
}

// Save creates the Agent in the database.
func (ac *AgentCreate) Save(ctx context.Context) (*Agent, error) {
	var (
		err  error
		node *Agent
	)
	ac.defaults()
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AgentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ac.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Agent)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AgentMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AgentCreate) SaveX(ctx context.Context) *Agent {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AgentCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AgentCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AgentCreate) defaults() {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := agent.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := agent.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
	if _, ok := ac.mutation.Verified(); !ok {
		v := agent.DefaultVerified
		ac.mutation.SetVerified(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AgentCreate) check() error {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Agent.created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Agent.updated_at"`)}
	}
	if _, ok := ac.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "Agent.username"`)}
	}
	if v, ok := ac.mutation.Username(); ok {
		if err := agent.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "Agent.username": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "Agent.password"`)}
	}
	if v, ok := ac.mutation.Password(); ok {
		if err := agent.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Agent.password": %w`, err)}
		}
	}
	if _, ok := ac.mutation.GhanaCard(); !ok {
		return &ValidationError{Name: "ghana_card", err: errors.New(`ent: missing required field "Agent.ghana_card"`)}
	}
	if v, ok := ac.mutation.GhanaCard(); ok {
		if err := agent.GhanaCardValidator(v); err != nil {
			return &ValidationError{Name: "ghana_card", err: fmt.Errorf(`ent: validator failed for field "Agent.ghana_card": %w`, err)}
		}
	}
	if _, ok := ac.mutation.LastName(); !ok {
		return &ValidationError{Name: "last_name", err: errors.New(`ent: missing required field "Agent.last_name"`)}
	}
	if v, ok := ac.mutation.LastName(); ok {
		if err := agent.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "Agent.last_name": %w`, err)}
		}
	}
	if _, ok := ac.mutation.OtherName(); !ok {
		return &ValidationError{Name: "other_name", err: errors.New(`ent: missing required field "Agent.other_name"`)}
	}
	if v, ok := ac.mutation.OtherName(); ok {
		if err := agent.OtherNameValidator(v); err != nil {
			return &ValidationError{Name: "other_name", err: fmt.Errorf(`ent: validator failed for field "Agent.other_name": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Phone(); !ok {
		return &ValidationError{Name: "phone", err: errors.New(`ent: missing required field "Agent.phone"`)}
	}
	if v, ok := ac.mutation.Phone(); ok {
		if err := agent.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "Agent.phone": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "Agent.address"`)}
	}
	if v, ok := ac.mutation.Address(); ok {
		if err := agent.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`ent: validator failed for field "Agent.address": %w`, err)}
		}
	}
	if _, ok := ac.mutation.DigitalAddress(); !ok {
		return &ValidationError{Name: "digital_address", err: errors.New(`ent: missing required field "Agent.digital_address"`)}
	}
	if v, ok := ac.mutation.DigitalAddress(); ok {
		if err := agent.DigitalAddressValidator(v); err != nil {
			return &ValidationError{Name: "digital_address", err: fmt.Errorf(`ent: validator failed for field "Agent.digital_address": %w`, err)}
		}
	}
<<<<<<< HEAD
	if _, ok := ac.mutation.Region(); !ok {
		return &ValidationError{Name: "region", err: errors.New(`ent: missing required field "Agent.region"`)}
	}
	if v, ok := ac.mutation.Region(); ok {
		if err := agent.RegionValidator(v); err != nil {
			return &ValidationError{Name: "region", err: fmt.Errorf(`ent: validator failed for field "Agent.region": %w`, err)}
		}
	}
	if _, ok := ac.mutation.District(); !ok {
		return &ValidationError{Name: "district", err: errors.New(`ent: missing required field "Agent.district"`)}
	}
	if v, ok := ac.mutation.District(); ok {
		if err := agent.DistrictValidator(v); err != nil {
			return &ValidationError{Name: "district", err: fmt.Errorf(`ent: validator failed for field "Agent.district": %w`, err)}
		}
	}
	if _, ok := ac.mutation.City(); !ok {
		return &ValidationError{Name: "city", err: errors.New(`ent: missing required field "Agent.city"`)}
	}
	if v, ok := ac.mutation.City(); ok {
		if err := agent.CityValidator(v); err != nil {
			return &ValidationError{Name: "city", err: fmt.Errorf(`ent: validator failed for field "Agent.city": %w`, err)}
		}
	}
=======
>>>>>>> dev
	if v, ok := ac.mutation.DefaultAccount(); ok {
		if err := agent.DefaultAccountValidator(v); err != nil {
			return &ValidationError{Name: "default_account", err: fmt.Errorf(`ent: validator failed for field "Agent.default_account": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Verified(); !ok {
		return &ValidationError{Name: "verified", err: errors.New(`ent: missing required field "Agent.verified"`)}
	}
	return nil
}

func (ac *AgentCreate) sqlSave(ctx context.Context) (*Agent, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ac *AgentCreate) createSpec() (*Agent, *sqlgraph.CreateSpec) {
	var (
		_node = &Agent{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: agent.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: agent.FieldID,
			},
		}
	)
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: agent.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: agent.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.Username(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agent.FieldUsername,
		})
		_node.Username = value
	}
	if value, ok := ac.mutation.Password(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: agent.FieldPassword,
		})
		_node.Password = value
	}
	if value, ok := ac.mutation.GhanaCard(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agent.FieldGhanaCard,
		})
		_node.GhanaCard = value
	}
	if value, ok := ac.mutation.LastName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agent.FieldLastName,
		})
		_node.LastName = value
	}
	if value, ok := ac.mutation.OtherName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agent.FieldOtherName,
		})
		_node.OtherName = value
	}
	if value, ok := ac.mutation.Phone(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agent.FieldPhone,
		})
		_node.Phone = value
	}
	if value, ok := ac.mutation.OtherPhone(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agent.FieldOtherPhone,
		})
		_node.OtherPhone = &value
	}
	if value, ok := ac.mutation.Address(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agent.FieldAddress,
		})
		_node.Address = value
	}
	if value, ok := ac.mutation.DigitalAddress(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agent.FieldDigitalAddress,
		})
		_node.DigitalAddress = value
	}
	if value, ok := ac.mutation.Region(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agent.FieldRegion,
		})
		_node.Region = &value
	}
	if value, ok := ac.mutation.District(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agent.FieldDistrict,
		})
		_node.District = &value
	}
	if value, ok := ac.mutation.City(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agent.FieldCity,
		})
		_node.City = &value
	}
	if value, ok := ac.mutation.DefaultAccount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: agent.FieldDefaultAccount,
		})
		_node.DefaultAccount = value
	}
	if value, ok := ac.mutation.BankAccount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: agent.FieldBankAccount,
		})
		_node.BankAccount = value
	}
	if value, ok := ac.mutation.MomoAccount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: agent.FieldMomoAccount,
		})
		_node.MomoAccount = value
	}
	if value, ok := ac.mutation.Verified(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: agent.FieldVerified,
		})
		_node.Verified = value
	}
	if value, ok := ac.mutation.Compliance(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: agent.FieldCompliance,
		})
		_node.Compliance = value
	}
	if nodes := ac.mutation.AddressesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   agent.AddressesTable,
			Columns: []string{agent.AddressesColumn},
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
	if nodes := ac.mutation.OrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   agent.OrdersTable,
			Columns: []string{agent.OrdersColumn},
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
	if nodes := ac.mutation.FavouritesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   agent.FavouritesTable,
			Columns: []string{agent.FavouritesColumn},
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
	if nodes := ac.mutation.StoreIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   agent.StoreTable,
			Columns: []string{agent.StoreColumn},
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
	if nodes := ac.mutation.RequestsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   agent.RequestsTable,
			Columns: []string{agent.RequestsColumn},
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
	return _node, _spec
}

// AgentCreateBulk is the builder for creating many Agent entities in bulk.
type AgentCreateBulk struct {
	config
	builders []*AgentCreate
}

// Save creates the Agent entities in the database.
func (acb *AgentCreateBulk) Save(ctx context.Context) ([]*Agent, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Agent, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AgentMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AgentCreateBulk) SaveX(ctx context.Context) []*Agent {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AgentCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AgentCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
