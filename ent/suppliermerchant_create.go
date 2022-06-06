// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SeyramWood/ent/merchant"
	"github.com/SeyramWood/ent/suppliermerchant"
)

// SupplierMerchantCreate is the builder for creating a SupplierMerchant entity.
type SupplierMerchantCreate struct {
	config
	mutation *SupplierMerchantMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (smc *SupplierMerchantCreate) SetCreatedAt(t time.Time) *SupplierMerchantCreate {
	smc.mutation.SetCreatedAt(t)
	return smc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (smc *SupplierMerchantCreate) SetNillableCreatedAt(t *time.Time) *SupplierMerchantCreate {
	if t != nil {
		smc.SetCreatedAt(*t)
	}
	return smc
}

// SetUpdatedAt sets the "updated_at" field.
func (smc *SupplierMerchantCreate) SetUpdatedAt(t time.Time) *SupplierMerchantCreate {
	smc.mutation.SetUpdatedAt(t)
	return smc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (smc *SupplierMerchantCreate) SetNillableUpdatedAt(t *time.Time) *SupplierMerchantCreate {
	if t != nil {
		smc.SetUpdatedAt(*t)
	}
	return smc
}

// SetGhanaCard sets the "ghana_card" field.
func (smc *SupplierMerchantCreate) SetGhanaCard(s string) *SupplierMerchantCreate {
	smc.mutation.SetGhanaCard(s)
	return smc
}

// SetLastName sets the "last_name" field.
func (smc *SupplierMerchantCreate) SetLastName(s string) *SupplierMerchantCreate {
	smc.mutation.SetLastName(s)
	return smc
}

// SetOtherName sets the "other_name" field.
func (smc *SupplierMerchantCreate) SetOtherName(s string) *SupplierMerchantCreate {
	smc.mutation.SetOtherName(s)
	return smc
}

// SetPhone sets the "phone" field.
func (smc *SupplierMerchantCreate) SetPhone(s string) *SupplierMerchantCreate {
	smc.mutation.SetPhone(s)
	return smc
}

// SetOtherPhone sets the "other_phone" field.
func (smc *SupplierMerchantCreate) SetOtherPhone(s string) *SupplierMerchantCreate {
	smc.mutation.SetOtherPhone(s)
	return smc
}

// SetNillableOtherPhone sets the "other_phone" field if the given value is not nil.
func (smc *SupplierMerchantCreate) SetNillableOtherPhone(s *string) *SupplierMerchantCreate {
	if s != nil {
		smc.SetOtherPhone(*s)
	}
	return smc
}

// SetAddress sets the "address" field.
func (smc *SupplierMerchantCreate) SetAddress(s string) *SupplierMerchantCreate {
	smc.mutation.SetAddress(s)
	return smc
}

// SetDigitalAddress sets the "digital_address" field.
func (smc *SupplierMerchantCreate) SetDigitalAddress(s string) *SupplierMerchantCreate {
	smc.mutation.SetDigitalAddress(s)
	return smc
}

// SetMerchantID sets the "merchant" edge to the Merchant entity by ID.
func (smc *SupplierMerchantCreate) SetMerchantID(id int) *SupplierMerchantCreate {
	smc.mutation.SetMerchantID(id)
	return smc
}

// SetMerchant sets the "merchant" edge to the Merchant entity.
func (smc *SupplierMerchantCreate) SetMerchant(m *Merchant) *SupplierMerchantCreate {
	return smc.SetMerchantID(m.ID)
}

// Mutation returns the SupplierMerchantMutation object of the builder.
func (smc *SupplierMerchantCreate) Mutation() *SupplierMerchantMutation {
	return smc.mutation
}

// Save creates the SupplierMerchant in the database.
func (smc *SupplierMerchantCreate) Save(ctx context.Context) (*SupplierMerchant, error) {
	var (
		err  error
		node *SupplierMerchant
	)
	smc.defaults()
	if len(smc.hooks) == 0 {
		if err = smc.check(); err != nil {
			return nil, err
		}
		node, err = smc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SupplierMerchantMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = smc.check(); err != nil {
				return nil, err
			}
			smc.mutation = mutation
			if node, err = smc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(smc.hooks) - 1; i >= 0; i-- {
			if smc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = smc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, smc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (smc *SupplierMerchantCreate) SaveX(ctx context.Context) *SupplierMerchant {
	v, err := smc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (smc *SupplierMerchantCreate) Exec(ctx context.Context) error {
	_, err := smc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smc *SupplierMerchantCreate) ExecX(ctx context.Context) {
	if err := smc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (smc *SupplierMerchantCreate) defaults() {
	if _, ok := smc.mutation.CreatedAt(); !ok {
		v := suppliermerchant.DefaultCreatedAt()
		smc.mutation.SetCreatedAt(v)
	}
	if _, ok := smc.mutation.UpdatedAt(); !ok {
		v := suppliermerchant.DefaultUpdatedAt()
		smc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (smc *SupplierMerchantCreate) check() error {
	if _, ok := smc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "SupplierMerchant.created_at"`)}
	}
	if _, ok := smc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "SupplierMerchant.updated_at"`)}
	}
	if _, ok := smc.mutation.GhanaCard(); !ok {
		return &ValidationError{Name: "ghana_card", err: errors.New(`ent: missing required field "SupplierMerchant.ghana_card"`)}
	}
	if v, ok := smc.mutation.GhanaCard(); ok {
		if err := suppliermerchant.GhanaCardValidator(v); err != nil {
			return &ValidationError{Name: "ghana_card", err: fmt.Errorf(`ent: validator failed for field "SupplierMerchant.ghana_card": %w`, err)}
		}
	}
	if _, ok := smc.mutation.LastName(); !ok {
		return &ValidationError{Name: "last_name", err: errors.New(`ent: missing required field "SupplierMerchant.last_name"`)}
	}
	if v, ok := smc.mutation.LastName(); ok {
		if err := suppliermerchant.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "SupplierMerchant.last_name": %w`, err)}
		}
	}
	if _, ok := smc.mutation.OtherName(); !ok {
		return &ValidationError{Name: "other_name", err: errors.New(`ent: missing required field "SupplierMerchant.other_name"`)}
	}
	if v, ok := smc.mutation.OtherName(); ok {
		if err := suppliermerchant.OtherNameValidator(v); err != nil {
			return &ValidationError{Name: "other_name", err: fmt.Errorf(`ent: validator failed for field "SupplierMerchant.other_name": %w`, err)}
		}
	}
	if _, ok := smc.mutation.Phone(); !ok {
		return &ValidationError{Name: "phone", err: errors.New(`ent: missing required field "SupplierMerchant.phone"`)}
	}
	if v, ok := smc.mutation.Phone(); ok {
		if err := suppliermerchant.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "SupplierMerchant.phone": %w`, err)}
		}
	}
	if _, ok := smc.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "SupplierMerchant.address"`)}
	}
	if v, ok := smc.mutation.Address(); ok {
		if err := suppliermerchant.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`ent: validator failed for field "SupplierMerchant.address": %w`, err)}
		}
	}
	if _, ok := smc.mutation.DigitalAddress(); !ok {
		return &ValidationError{Name: "digital_address", err: errors.New(`ent: missing required field "SupplierMerchant.digital_address"`)}
	}
	if v, ok := smc.mutation.DigitalAddress(); ok {
		if err := suppliermerchant.DigitalAddressValidator(v); err != nil {
			return &ValidationError{Name: "digital_address", err: fmt.Errorf(`ent: validator failed for field "SupplierMerchant.digital_address": %w`, err)}
		}
	}
	if _, ok := smc.mutation.MerchantID(); !ok {
		return &ValidationError{Name: "merchant", err: errors.New(`ent: missing required edge "SupplierMerchant.merchant"`)}
	}
	return nil
}

func (smc *SupplierMerchantCreate) sqlSave(ctx context.Context) (*SupplierMerchant, error) {
	_node, _spec := smc.createSpec()
	if err := sqlgraph.CreateNode(ctx, smc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (smc *SupplierMerchantCreate) createSpec() (*SupplierMerchant, *sqlgraph.CreateSpec) {
	var (
		_node = &SupplierMerchant{config: smc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: suppliermerchant.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: suppliermerchant.FieldID,
			},
		}
	)
	if value, ok := smc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: suppliermerchant.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := smc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: suppliermerchant.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := smc.mutation.GhanaCard(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: suppliermerchant.FieldGhanaCard,
		})
		_node.GhanaCard = value
	}
	if value, ok := smc.mutation.LastName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: suppliermerchant.FieldLastName,
		})
		_node.LastName = value
	}
	if value, ok := smc.mutation.OtherName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: suppliermerchant.FieldOtherName,
		})
		_node.OtherName = value
	}
	if value, ok := smc.mutation.Phone(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: suppliermerchant.FieldPhone,
		})
		_node.Phone = value
	}
	if value, ok := smc.mutation.OtherPhone(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: suppliermerchant.FieldOtherPhone,
		})
		_node.OtherPhone = &value
	}
	if value, ok := smc.mutation.Address(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: suppliermerchant.FieldAddress,
		})
		_node.Address = value
	}
	if value, ok := smc.mutation.DigitalAddress(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: suppliermerchant.FieldDigitalAddress,
		})
		_node.DigitalAddress = value
	}
	if nodes := smc.mutation.MerchantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   suppliermerchant.MerchantTable,
			Columns: []string{suppliermerchant.MerchantColumn},
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
		_node.merchant_supplier = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SupplierMerchantCreateBulk is the builder for creating many SupplierMerchant entities in bulk.
type SupplierMerchantCreateBulk struct {
	config
	builders []*SupplierMerchantCreate
}

// Save creates the SupplierMerchant entities in the database.
func (smcb *SupplierMerchantCreateBulk) Save(ctx context.Context) ([]*SupplierMerchant, error) {
	specs := make([]*sqlgraph.CreateSpec, len(smcb.builders))
	nodes := make([]*SupplierMerchant, len(smcb.builders))
	mutators := make([]Mutator, len(smcb.builders))
	for i := range smcb.builders {
		func(i int, root context.Context) {
			builder := smcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SupplierMerchantMutation)
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
					_, err = mutators[i+1].Mutate(root, smcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, smcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, smcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (smcb *SupplierMerchantCreateBulk) SaveX(ctx context.Context) []*SupplierMerchant {
	v, err := smcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (smcb *SupplierMerchantCreateBulk) Exec(ctx context.Context) error {
	_, err := smcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smcb *SupplierMerchantCreateBulk) ExecX(ctx context.Context) {
	if err := smcb.Exec(ctx); err != nil {
		panic(err)
	}
}
