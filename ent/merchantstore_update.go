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
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/ent/merchant"
	"github.com/SeyramWood/ent/merchantstore"
	"github.com/SeyramWood/ent/orderdetail"
	"github.com/SeyramWood/ent/predicate"
)

// MerchantStoreUpdate is the builder for updating MerchantStore entities.
type MerchantStoreUpdate struct {
	config
	hooks    []Hook
	mutation *MerchantStoreMutation
}

// Where appends a list predicates to the MerchantStoreUpdate builder.
func (msu *MerchantStoreUpdate) Where(ps ...predicate.MerchantStore) *MerchantStoreUpdate {
	msu.mutation.Where(ps...)
	return msu
}

// SetUpdatedAt sets the "updated_at" field.
func (msu *MerchantStoreUpdate) SetUpdatedAt(t time.Time) *MerchantStoreUpdate {
	msu.mutation.SetUpdatedAt(t)
	return msu
}

// SetName sets the "name" field.
func (msu *MerchantStoreUpdate) SetName(s string) *MerchantStoreUpdate {
	msu.mutation.SetName(s)
	return msu
}

// SetAbout sets the "about" field.
func (msu *MerchantStoreUpdate) SetAbout(s string) *MerchantStoreUpdate {
	msu.mutation.SetAbout(s)
	return msu
}

// SetDescTitle sets the "desc_title" field.
func (msu *MerchantStoreUpdate) SetDescTitle(s string) *MerchantStoreUpdate {
	msu.mutation.SetDescTitle(s)
	return msu
}

// SetDescription sets the "description" field.
func (msu *MerchantStoreUpdate) SetDescription(s string) *MerchantStoreUpdate {
	msu.mutation.SetDescription(s)
	return msu
}

// SetLogo sets the "logo" field.
func (msu *MerchantStoreUpdate) SetLogo(s string) *MerchantStoreUpdate {
	msu.mutation.SetLogo(s)
	return msu
}

// SetImages sets the "images" field.
func (msu *MerchantStoreUpdate) SetImages(s []string) *MerchantStoreUpdate {
	msu.mutation.SetImages(s)
	return msu
}

// ClearImages clears the value of the "images" field.
func (msu *MerchantStoreUpdate) ClearImages() *MerchantStoreUpdate {
	msu.mutation.ClearImages()
	return msu
}

// SetDefaultAccount sets the "default_account" field.
func (msu *MerchantStoreUpdate) SetDefaultAccount(ma merchantstore.DefaultAccount) *MerchantStoreUpdate {
	msu.mutation.SetDefaultAccount(ma)
	return msu
}

// SetNillableDefaultAccount sets the "default_account" field if the given value is not nil.
func (msu *MerchantStoreUpdate) SetNillableDefaultAccount(ma *merchantstore.DefaultAccount) *MerchantStoreUpdate {
	if ma != nil {
		msu.SetDefaultAccount(*ma)
	}
	return msu
}

// ClearDefaultAccount clears the value of the "default_account" field.
func (msu *MerchantStoreUpdate) ClearDefaultAccount() *MerchantStoreUpdate {
	msu.mutation.ClearDefaultAccount()
	return msu
}

// SetBankAccount sets the "bank_account" field.
func (msu *MerchantStoreUpdate) SetBankAccount(mba *models.MerchantBankAccount) *MerchantStoreUpdate {
	msu.mutation.SetBankAccount(mba)
	return msu
}

// ClearBankAccount clears the value of the "bank_account" field.
func (msu *MerchantStoreUpdate) ClearBankAccount() *MerchantStoreUpdate {
	msu.mutation.ClearBankAccount()
	return msu
}

// SetMomoAccount sets the "momo_account" field.
func (msu *MerchantStoreUpdate) SetMomoAccount(mma *models.MerchantMomoAccount) *MerchantStoreUpdate {
	msu.mutation.SetMomoAccount(mma)
	return msu
}

// ClearMomoAccount clears the value of the "momo_account" field.
func (msu *MerchantStoreUpdate) ClearMomoAccount() *MerchantStoreUpdate {
	msu.mutation.ClearMomoAccount()
	return msu
}

// SetMerchantType sets the "merchant_type" field.
func (msu *MerchantStoreUpdate) SetMerchantType(s string) *MerchantStoreUpdate {
	msu.mutation.SetMerchantType(s)
	return msu
}

// SetMerchantID sets the "merchant" edge to the Merchant entity by ID.
func (msu *MerchantStoreUpdate) SetMerchantID(id int) *MerchantStoreUpdate {
	msu.mutation.SetMerchantID(id)
	return msu
}

// SetNillableMerchantID sets the "merchant" edge to the Merchant entity by ID if the given value is not nil.
func (msu *MerchantStoreUpdate) SetNillableMerchantID(id *int) *MerchantStoreUpdate {
	if id != nil {
		msu = msu.SetMerchantID(*id)
	}
	return msu
}

// SetMerchant sets the "merchant" edge to the Merchant entity.
func (msu *MerchantStoreUpdate) SetMerchant(m *Merchant) *MerchantStoreUpdate {
	return msu.SetMerchantID(m.ID)
}

// AddOrderIDs adds the "orders" edge to the OrderDetail entity by IDs.
func (msu *MerchantStoreUpdate) AddOrderIDs(ids ...int) *MerchantStoreUpdate {
	msu.mutation.AddOrderIDs(ids...)
	return msu
}

// AddOrders adds the "orders" edges to the OrderDetail entity.
func (msu *MerchantStoreUpdate) AddOrders(o ...*OrderDetail) *MerchantStoreUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return msu.AddOrderIDs(ids...)
}

// Mutation returns the MerchantStoreMutation object of the builder.
func (msu *MerchantStoreUpdate) Mutation() *MerchantStoreMutation {
	return msu.mutation
}

// ClearMerchant clears the "merchant" edge to the Merchant entity.
func (msu *MerchantStoreUpdate) ClearMerchant() *MerchantStoreUpdate {
	msu.mutation.ClearMerchant()
	return msu
}

// ClearOrders clears all "orders" edges to the OrderDetail entity.
func (msu *MerchantStoreUpdate) ClearOrders() *MerchantStoreUpdate {
	msu.mutation.ClearOrders()
	return msu
}

// RemoveOrderIDs removes the "orders" edge to OrderDetail entities by IDs.
func (msu *MerchantStoreUpdate) RemoveOrderIDs(ids ...int) *MerchantStoreUpdate {
	msu.mutation.RemoveOrderIDs(ids...)
	return msu
}

// RemoveOrders removes "orders" edges to OrderDetail entities.
func (msu *MerchantStoreUpdate) RemoveOrders(o ...*OrderDetail) *MerchantStoreUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return msu.RemoveOrderIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (msu *MerchantStoreUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	msu.defaults()
	if len(msu.hooks) == 0 {
		if err = msu.check(); err != nil {
			return 0, err
		}
		affected, err = msu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MerchantStoreMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = msu.check(); err != nil {
				return 0, err
			}
			msu.mutation = mutation
			affected, err = msu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(msu.hooks) - 1; i >= 0; i-- {
			if msu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = msu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, msu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (msu *MerchantStoreUpdate) SaveX(ctx context.Context) int {
	affected, err := msu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (msu *MerchantStoreUpdate) Exec(ctx context.Context) error {
	_, err := msu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (msu *MerchantStoreUpdate) ExecX(ctx context.Context) {
	if err := msu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (msu *MerchantStoreUpdate) defaults() {
	if _, ok := msu.mutation.UpdatedAt(); !ok {
		v := merchantstore.UpdateDefaultUpdatedAt()
		msu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (msu *MerchantStoreUpdate) check() error {
	if v, ok := msu.mutation.Name(); ok {
		if err := merchantstore.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "MerchantStore.name": %w`, err)}
		}
	}
	if v, ok := msu.mutation.About(); ok {
		if err := merchantstore.AboutValidator(v); err != nil {
			return &ValidationError{Name: "about", err: fmt.Errorf(`ent: validator failed for field "MerchantStore.about": %w`, err)}
		}
	}
	if v, ok := msu.mutation.DescTitle(); ok {
		if err := merchantstore.DescTitleValidator(v); err != nil {
			return &ValidationError{Name: "desc_title", err: fmt.Errorf(`ent: validator failed for field "MerchantStore.desc_title": %w`, err)}
		}
	}
	if v, ok := msu.mutation.Description(); ok {
		if err := merchantstore.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "MerchantStore.description": %w`, err)}
		}
	}
	if v, ok := msu.mutation.Logo(); ok {
		if err := merchantstore.LogoValidator(v); err != nil {
			return &ValidationError{Name: "logo", err: fmt.Errorf(`ent: validator failed for field "MerchantStore.logo": %w`, err)}
		}
	}
	if v, ok := msu.mutation.DefaultAccount(); ok {
		if err := merchantstore.DefaultAccountValidator(v); err != nil {
			return &ValidationError{Name: "default_account", err: fmt.Errorf(`ent: validator failed for field "MerchantStore.default_account": %w`, err)}
		}
	}
	if v, ok := msu.mutation.MerchantType(); ok {
		if err := merchantstore.MerchantTypeValidator(v); err != nil {
			return &ValidationError{Name: "merchant_type", err: fmt.Errorf(`ent: validator failed for field "MerchantStore.merchant_type": %w`, err)}
		}
	}
	return nil
}

func (msu *MerchantStoreUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   merchantstore.Table,
			Columns: merchantstore.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: merchantstore.FieldID,
			},
		},
	}
	if ps := msu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := msu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: merchantstore.FieldUpdatedAt,
		})
	}
	if value, ok := msu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchantstore.FieldName,
		})
	}
	if value, ok := msu.mutation.About(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchantstore.FieldAbout,
		})
	}
	if value, ok := msu.mutation.DescTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchantstore.FieldDescTitle,
		})
	}
	if value, ok := msu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchantstore.FieldDescription,
		})
	}
	if value, ok := msu.mutation.Logo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchantstore.FieldLogo,
		})
	}
	if value, ok := msu.mutation.Images(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: merchantstore.FieldImages,
		})
	}
	if msu.mutation.ImagesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: merchantstore.FieldImages,
		})
	}
	if value, ok := msu.mutation.DefaultAccount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: merchantstore.FieldDefaultAccount,
		})
	}
	if msu.mutation.DefaultAccountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: merchantstore.FieldDefaultAccount,
		})
	}
	if value, ok := msu.mutation.BankAccount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: merchantstore.FieldBankAccount,
		})
	}
	if msu.mutation.BankAccountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: merchantstore.FieldBankAccount,
		})
	}
	if value, ok := msu.mutation.MomoAccount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: merchantstore.FieldMomoAccount,
		})
	}
	if msu.mutation.MomoAccountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: merchantstore.FieldMomoAccount,
		})
	}
	if value, ok := msu.mutation.MerchantType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchantstore.FieldMerchantType,
		})
	}
	if msu.mutation.MerchantCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := msu.mutation.MerchantIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if msu.mutation.OrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   merchantstore.OrdersTable,
			Columns: []string{merchantstore.OrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderdetail.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := msu.mutation.RemovedOrdersIDs(); len(nodes) > 0 && !msu.mutation.OrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   merchantstore.OrdersTable,
			Columns: []string{merchantstore.OrdersColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := msu.mutation.OrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   merchantstore.OrdersTable,
			Columns: []string{merchantstore.OrdersColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, msu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{merchantstore.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// MerchantStoreUpdateOne is the builder for updating a single MerchantStore entity.
type MerchantStoreUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MerchantStoreMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (msuo *MerchantStoreUpdateOne) SetUpdatedAt(t time.Time) *MerchantStoreUpdateOne {
	msuo.mutation.SetUpdatedAt(t)
	return msuo
}

// SetName sets the "name" field.
func (msuo *MerchantStoreUpdateOne) SetName(s string) *MerchantStoreUpdateOne {
	msuo.mutation.SetName(s)
	return msuo
}

// SetAbout sets the "about" field.
func (msuo *MerchantStoreUpdateOne) SetAbout(s string) *MerchantStoreUpdateOne {
	msuo.mutation.SetAbout(s)
	return msuo
}

// SetDescTitle sets the "desc_title" field.
func (msuo *MerchantStoreUpdateOne) SetDescTitle(s string) *MerchantStoreUpdateOne {
	msuo.mutation.SetDescTitle(s)
	return msuo
}

// SetDescription sets the "description" field.
func (msuo *MerchantStoreUpdateOne) SetDescription(s string) *MerchantStoreUpdateOne {
	msuo.mutation.SetDescription(s)
	return msuo
}

// SetLogo sets the "logo" field.
func (msuo *MerchantStoreUpdateOne) SetLogo(s string) *MerchantStoreUpdateOne {
	msuo.mutation.SetLogo(s)
	return msuo
}

// SetImages sets the "images" field.
func (msuo *MerchantStoreUpdateOne) SetImages(s []string) *MerchantStoreUpdateOne {
	msuo.mutation.SetImages(s)
	return msuo
}

// ClearImages clears the value of the "images" field.
func (msuo *MerchantStoreUpdateOne) ClearImages() *MerchantStoreUpdateOne {
	msuo.mutation.ClearImages()
	return msuo
}

// SetDefaultAccount sets the "default_account" field.
func (msuo *MerchantStoreUpdateOne) SetDefaultAccount(ma merchantstore.DefaultAccount) *MerchantStoreUpdateOne {
	msuo.mutation.SetDefaultAccount(ma)
	return msuo
}

// SetNillableDefaultAccount sets the "default_account" field if the given value is not nil.
func (msuo *MerchantStoreUpdateOne) SetNillableDefaultAccount(ma *merchantstore.DefaultAccount) *MerchantStoreUpdateOne {
	if ma != nil {
		msuo.SetDefaultAccount(*ma)
	}
	return msuo
}

// ClearDefaultAccount clears the value of the "default_account" field.
func (msuo *MerchantStoreUpdateOne) ClearDefaultAccount() *MerchantStoreUpdateOne {
	msuo.mutation.ClearDefaultAccount()
	return msuo
}

// SetBankAccount sets the "bank_account" field.
func (msuo *MerchantStoreUpdateOne) SetBankAccount(mba *models.MerchantBankAccount) *MerchantStoreUpdateOne {
	msuo.mutation.SetBankAccount(mba)
	return msuo
}

// ClearBankAccount clears the value of the "bank_account" field.
func (msuo *MerchantStoreUpdateOne) ClearBankAccount() *MerchantStoreUpdateOne {
	msuo.mutation.ClearBankAccount()
	return msuo
}

// SetMomoAccount sets the "momo_account" field.
func (msuo *MerchantStoreUpdateOne) SetMomoAccount(mma *models.MerchantMomoAccount) *MerchantStoreUpdateOne {
	msuo.mutation.SetMomoAccount(mma)
	return msuo
}

// ClearMomoAccount clears the value of the "momo_account" field.
func (msuo *MerchantStoreUpdateOne) ClearMomoAccount() *MerchantStoreUpdateOne {
	msuo.mutation.ClearMomoAccount()
	return msuo
}

// SetMerchantType sets the "merchant_type" field.
func (msuo *MerchantStoreUpdateOne) SetMerchantType(s string) *MerchantStoreUpdateOne {
	msuo.mutation.SetMerchantType(s)
	return msuo
}

// SetMerchantID sets the "merchant" edge to the Merchant entity by ID.
func (msuo *MerchantStoreUpdateOne) SetMerchantID(id int) *MerchantStoreUpdateOne {
	msuo.mutation.SetMerchantID(id)
	return msuo
}

// SetNillableMerchantID sets the "merchant" edge to the Merchant entity by ID if the given value is not nil.
func (msuo *MerchantStoreUpdateOne) SetNillableMerchantID(id *int) *MerchantStoreUpdateOne {
	if id != nil {
		msuo = msuo.SetMerchantID(*id)
	}
	return msuo
}

// SetMerchant sets the "merchant" edge to the Merchant entity.
func (msuo *MerchantStoreUpdateOne) SetMerchant(m *Merchant) *MerchantStoreUpdateOne {
	return msuo.SetMerchantID(m.ID)
}

// AddOrderIDs adds the "orders" edge to the OrderDetail entity by IDs.
func (msuo *MerchantStoreUpdateOne) AddOrderIDs(ids ...int) *MerchantStoreUpdateOne {
	msuo.mutation.AddOrderIDs(ids...)
	return msuo
}

// AddOrders adds the "orders" edges to the OrderDetail entity.
func (msuo *MerchantStoreUpdateOne) AddOrders(o ...*OrderDetail) *MerchantStoreUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return msuo.AddOrderIDs(ids...)
}

// Mutation returns the MerchantStoreMutation object of the builder.
func (msuo *MerchantStoreUpdateOne) Mutation() *MerchantStoreMutation {
	return msuo.mutation
}

// ClearMerchant clears the "merchant" edge to the Merchant entity.
func (msuo *MerchantStoreUpdateOne) ClearMerchant() *MerchantStoreUpdateOne {
	msuo.mutation.ClearMerchant()
	return msuo
}

// ClearOrders clears all "orders" edges to the OrderDetail entity.
func (msuo *MerchantStoreUpdateOne) ClearOrders() *MerchantStoreUpdateOne {
	msuo.mutation.ClearOrders()
	return msuo
}

// RemoveOrderIDs removes the "orders" edge to OrderDetail entities by IDs.
func (msuo *MerchantStoreUpdateOne) RemoveOrderIDs(ids ...int) *MerchantStoreUpdateOne {
	msuo.mutation.RemoveOrderIDs(ids...)
	return msuo
}

// RemoveOrders removes "orders" edges to OrderDetail entities.
func (msuo *MerchantStoreUpdateOne) RemoveOrders(o ...*OrderDetail) *MerchantStoreUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return msuo.RemoveOrderIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (msuo *MerchantStoreUpdateOne) Select(field string, fields ...string) *MerchantStoreUpdateOne {
	msuo.fields = append([]string{field}, fields...)
	return msuo
}

// Save executes the query and returns the updated MerchantStore entity.
func (msuo *MerchantStoreUpdateOne) Save(ctx context.Context) (*MerchantStore, error) {
	var (
		err  error
		node *MerchantStore
	)
	msuo.defaults()
	if len(msuo.hooks) == 0 {
		if err = msuo.check(); err != nil {
			return nil, err
		}
		node, err = msuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MerchantStoreMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = msuo.check(); err != nil {
				return nil, err
			}
			msuo.mutation = mutation
			node, err = msuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(msuo.hooks) - 1; i >= 0; i-- {
			if msuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = msuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, msuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (msuo *MerchantStoreUpdateOne) SaveX(ctx context.Context) *MerchantStore {
	node, err := msuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (msuo *MerchantStoreUpdateOne) Exec(ctx context.Context) error {
	_, err := msuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (msuo *MerchantStoreUpdateOne) ExecX(ctx context.Context) {
	if err := msuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (msuo *MerchantStoreUpdateOne) defaults() {
	if _, ok := msuo.mutation.UpdatedAt(); !ok {
		v := merchantstore.UpdateDefaultUpdatedAt()
		msuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (msuo *MerchantStoreUpdateOne) check() error {
	if v, ok := msuo.mutation.Name(); ok {
		if err := merchantstore.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "MerchantStore.name": %w`, err)}
		}
	}
	if v, ok := msuo.mutation.About(); ok {
		if err := merchantstore.AboutValidator(v); err != nil {
			return &ValidationError{Name: "about", err: fmt.Errorf(`ent: validator failed for field "MerchantStore.about": %w`, err)}
		}
	}
	if v, ok := msuo.mutation.DescTitle(); ok {
		if err := merchantstore.DescTitleValidator(v); err != nil {
			return &ValidationError{Name: "desc_title", err: fmt.Errorf(`ent: validator failed for field "MerchantStore.desc_title": %w`, err)}
		}
	}
	if v, ok := msuo.mutation.Description(); ok {
		if err := merchantstore.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "MerchantStore.description": %w`, err)}
		}
	}
	if v, ok := msuo.mutation.Logo(); ok {
		if err := merchantstore.LogoValidator(v); err != nil {
			return &ValidationError{Name: "logo", err: fmt.Errorf(`ent: validator failed for field "MerchantStore.logo": %w`, err)}
		}
	}
	if v, ok := msuo.mutation.DefaultAccount(); ok {
		if err := merchantstore.DefaultAccountValidator(v); err != nil {
			return &ValidationError{Name: "default_account", err: fmt.Errorf(`ent: validator failed for field "MerchantStore.default_account": %w`, err)}
		}
	}
	if v, ok := msuo.mutation.MerchantType(); ok {
		if err := merchantstore.MerchantTypeValidator(v); err != nil {
			return &ValidationError{Name: "merchant_type", err: fmt.Errorf(`ent: validator failed for field "MerchantStore.merchant_type": %w`, err)}
		}
	}
	return nil
}

func (msuo *MerchantStoreUpdateOne) sqlSave(ctx context.Context) (_node *MerchantStore, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   merchantstore.Table,
			Columns: merchantstore.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: merchantstore.FieldID,
			},
		},
	}
	id, ok := msuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MerchantStore.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := msuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, merchantstore.FieldID)
		for _, f := range fields {
			if !merchantstore.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != merchantstore.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := msuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := msuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: merchantstore.FieldUpdatedAt,
		})
	}
	if value, ok := msuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchantstore.FieldName,
		})
	}
	if value, ok := msuo.mutation.About(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchantstore.FieldAbout,
		})
	}
	if value, ok := msuo.mutation.DescTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchantstore.FieldDescTitle,
		})
	}
	if value, ok := msuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchantstore.FieldDescription,
		})
	}
	if value, ok := msuo.mutation.Logo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchantstore.FieldLogo,
		})
	}
	if value, ok := msuo.mutation.Images(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: merchantstore.FieldImages,
		})
	}
	if msuo.mutation.ImagesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: merchantstore.FieldImages,
		})
	}
	if value, ok := msuo.mutation.DefaultAccount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: merchantstore.FieldDefaultAccount,
		})
	}
	if msuo.mutation.DefaultAccountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: merchantstore.FieldDefaultAccount,
		})
	}
	if value, ok := msuo.mutation.BankAccount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: merchantstore.FieldBankAccount,
		})
	}
	if msuo.mutation.BankAccountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: merchantstore.FieldBankAccount,
		})
	}
	if value, ok := msuo.mutation.MomoAccount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: merchantstore.FieldMomoAccount,
		})
	}
	if msuo.mutation.MomoAccountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: merchantstore.FieldMomoAccount,
		})
	}
	if value, ok := msuo.mutation.MerchantType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchantstore.FieldMerchantType,
		})
	}
	if msuo.mutation.MerchantCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := msuo.mutation.MerchantIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if msuo.mutation.OrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   merchantstore.OrdersTable,
			Columns: []string{merchantstore.OrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderdetail.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := msuo.mutation.RemovedOrdersIDs(); len(nodes) > 0 && !msuo.mutation.OrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   merchantstore.OrdersTable,
			Columns: []string{merchantstore.OrdersColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := msuo.mutation.OrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   merchantstore.OrdersTable,
			Columns: []string{merchantstore.OrdersColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &MerchantStore{config: msuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, msuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{merchantstore.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}