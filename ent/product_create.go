// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SeyramWood/ent/favourite"
	"github.com/SeyramWood/ent/merchant"
	"github.com/SeyramWood/ent/orderdetail"
	"github.com/SeyramWood/ent/product"
	"github.com/SeyramWood/ent/productcategorymajor"
	"github.com/SeyramWood/ent/productcategoryminor"
)

// ProductCreate is the builder for creating a Product entity.
type ProductCreate struct {
	config
	mutation *ProductMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (pc *ProductCreate) SetCreatedAt(t time.Time) *ProductCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *ProductCreate) SetNillableCreatedAt(t *time.Time) *ProductCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *ProductCreate) SetUpdatedAt(t time.Time) *ProductCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *ProductCreate) SetNillableUpdatedAt(t *time.Time) *ProductCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetName sets the "name" field.
func (pc *ProductCreate) SetName(s string) *ProductCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetPrice sets the "price" field.
func (pc *ProductCreate) SetPrice(f float64) *ProductCreate {
	pc.mutation.SetPrice(f)
	return pc
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (pc *ProductCreate) SetNillablePrice(f *float64) *ProductCreate {
	if f != nil {
		pc.SetPrice(*f)
	}
	return pc
}

// SetPromoPrice sets the "promo_price" field.
func (pc *ProductCreate) SetPromoPrice(f float64) *ProductCreate {
	pc.mutation.SetPromoPrice(f)
	return pc
}

// SetNillablePromoPrice sets the "promo_price" field if the given value is not nil.
func (pc *ProductCreate) SetNillablePromoPrice(f *float64) *ProductCreate {
	if f != nil {
		pc.SetPromoPrice(*f)
	}
	return pc
}

// SetWeight sets the "weight" field.
func (pc *ProductCreate) SetWeight(u uint32) *ProductCreate {
	pc.mutation.SetWeight(u)
	return pc
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (pc *ProductCreate) SetNillableWeight(u *uint32) *ProductCreate {
	if u != nil {
		pc.SetWeight(*u)
	}
	return pc
}

// SetQuantity sets the "quantity" field.
func (pc *ProductCreate) SetQuantity(u uint32) *ProductCreate {
	pc.mutation.SetQuantity(u)
	return pc
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (pc *ProductCreate) SetNillableQuantity(u *uint32) *ProductCreate {
	if u != nil {
		pc.SetQuantity(*u)
	}
	return pc
}

// SetUnit sets the "unit" field.
func (pc *ProductCreate) SetUnit(s string) *ProductCreate {
	pc.mutation.SetUnit(s)
	return pc
}

// SetDescription sets the "description" field.
func (pc *ProductCreate) SetDescription(s string) *ProductCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetImage sets the "image" field.
func (pc *ProductCreate) SetImage(s string) *ProductCreate {
	pc.mutation.SetImage(s)
	return pc
}

// AddOrderDetailIDs adds the "order_details" edge to the OrderDetail entity by IDs.
func (pc *ProductCreate) AddOrderDetailIDs(ids ...int) *ProductCreate {
	pc.mutation.AddOrderDetailIDs(ids...)
	return pc
}

// AddOrderDetails adds the "order_details" edges to the OrderDetail entity.
func (pc *ProductCreate) AddOrderDetails(o ...*OrderDetail) *ProductCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return pc.AddOrderDetailIDs(ids...)
}

// AddFavouriteIDs adds the "favourites" edge to the Favourite entity by IDs.
func (pc *ProductCreate) AddFavouriteIDs(ids ...int) *ProductCreate {
	pc.mutation.AddFavouriteIDs(ids...)
	return pc
}

// AddFavourites adds the "favourites" edges to the Favourite entity.
func (pc *ProductCreate) AddFavourites(f ...*Favourite) *ProductCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return pc.AddFavouriteIDs(ids...)
}

// SetMerchantID sets the "merchant" edge to the Merchant entity by ID.
func (pc *ProductCreate) SetMerchantID(id int) *ProductCreate {
	pc.mutation.SetMerchantID(id)
	return pc
}

// SetMerchant sets the "merchant" edge to the Merchant entity.
func (pc *ProductCreate) SetMerchant(m *Merchant) *ProductCreate {
	return pc.SetMerchantID(m.ID)
}

// SetMajorID sets the "major" edge to the ProductCategoryMajor entity by ID.
func (pc *ProductCreate) SetMajorID(id int) *ProductCreate {
	pc.mutation.SetMajorID(id)
	return pc
}

// SetMajor sets the "major" edge to the ProductCategoryMajor entity.
func (pc *ProductCreate) SetMajor(p *ProductCategoryMajor) *ProductCreate {
	return pc.SetMajorID(p.ID)
}

// SetMinorID sets the "minor" edge to the ProductCategoryMinor entity by ID.
func (pc *ProductCreate) SetMinorID(id int) *ProductCreate {
	pc.mutation.SetMinorID(id)
	return pc
}

// SetMinor sets the "minor" edge to the ProductCategoryMinor entity.
func (pc *ProductCreate) SetMinor(p *ProductCategoryMinor) *ProductCreate {
	return pc.SetMinorID(p.ID)
}

// Mutation returns the ProductMutation object of the builder.
func (pc *ProductCreate) Mutation() *ProductMutation {
	return pc.mutation
}

// Save creates the Product in the database.
func (pc *ProductCreate) Save(ctx context.Context) (*Product, error) {
	var (
		err  error
		node *Product
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProductCreate) SaveX(ctx context.Context) *Product {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProductCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProductCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ProductCreate) defaults() {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := product.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := product.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pc.mutation.Price(); !ok {
		v := product.DefaultPrice
		pc.mutation.SetPrice(v)
	}
	if _, ok := pc.mutation.Weight(); !ok {
		v := product.DefaultWeight
		pc.mutation.SetWeight(v)
	}
	if _, ok := pc.mutation.Quantity(); !ok {
		v := product.DefaultQuantity
		pc.mutation.SetQuantity(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProductCreate) check() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Product.created_at"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Product.updated_at"`)}
	}
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Product.name"`)}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := product.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Product.name": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "Product.price"`)}
	}
	if _, ok := pc.mutation.Weight(); !ok {
		return &ValidationError{Name: "weight", err: errors.New(`ent: missing required field "Product.weight"`)}
	}
	if _, ok := pc.mutation.Quantity(); !ok {
		return &ValidationError{Name: "quantity", err: errors.New(`ent: missing required field "Product.quantity"`)}
	}
	if _, ok := pc.mutation.Unit(); !ok {
		return &ValidationError{Name: "unit", err: errors.New(`ent: missing required field "Product.unit"`)}
	}
	if v, ok := pc.mutation.Unit(); ok {
		if err := product.UnitValidator(v); err != nil {
			return &ValidationError{Name: "unit", err: fmt.Errorf(`ent: validator failed for field "Product.unit": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Product.description"`)}
	}
	if v, ok := pc.mutation.Description(); ok {
		if err := product.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Product.description": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Image(); !ok {
		return &ValidationError{Name: "image", err: errors.New(`ent: missing required field "Product.image"`)}
	}
	if v, ok := pc.mutation.Image(); ok {
		if err := product.ImageValidator(v); err != nil {
			return &ValidationError{Name: "image", err: fmt.Errorf(`ent: validator failed for field "Product.image": %w`, err)}
		}
	}
	if _, ok := pc.mutation.MerchantID(); !ok {
		return &ValidationError{Name: "merchant", err: errors.New(`ent: missing required edge "Product.merchant"`)}
	}
	if _, ok := pc.mutation.MajorID(); !ok {
		return &ValidationError{Name: "major", err: errors.New(`ent: missing required edge "Product.major"`)}
	}
	if _, ok := pc.mutation.MinorID(); !ok {
		return &ValidationError{Name: "minor", err: errors.New(`ent: missing required edge "Product.minor"`)}
	}
	return nil
}

func (pc *ProductCreate) sqlSave(ctx context.Context) (*Product, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pc *ProductCreate) createSpec() (*Product, *sqlgraph.CreateSpec) {
	var (
		_node = &Product{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: product.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: product.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: product.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: product.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldName,
		})
		_node.Name = value
	}
	if value, ok := pc.mutation.Price(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: product.FieldPrice,
		})
		_node.Price = value
	}
	if value, ok := pc.mutation.PromoPrice(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: product.FieldPromoPrice,
		})
		_node.PromoPrice = &value
	}
	if value, ok := pc.mutation.Weight(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: product.FieldWeight,
		})
		_node.Weight = value
	}
	if value, ok := pc.mutation.Quantity(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: product.FieldQuantity,
		})
		_node.Quantity = value
	}
	if value, ok := pc.mutation.Unit(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldUnit,
		})
		_node.Unit = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := pc.mutation.Image(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldImage,
		})
		_node.Image = value
	}
	if nodes := pc.mutation.OrderDetailsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.OrderDetailsTable,
			Columns: []string{product.OrderDetailsColumn},
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
	if nodes := pc.mutation.FavouritesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.FavouritesTable,
			Columns: []string{product.FavouritesColumn},
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
	if nodes := pc.mutation.MerchantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.MerchantTable,
			Columns: []string{product.MerchantColumn},
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
		_node.merchant_products = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.MajorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.MajorTable,
			Columns: []string{product.MajorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productcategorymajor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.product_category_major_products = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.MinorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.MinorTable,
			Columns: []string{product.MinorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productcategoryminor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.product_category_minor_products = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProductCreateBulk is the builder for creating many Product entities in bulk.
type ProductCreateBulk struct {
	config
	builders []*ProductCreate
}

// Save creates the Product entities in the database.
func (pcb *ProductCreateBulk) Save(ctx context.Context) ([]*Product, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Product, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProductMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProductCreateBulk) SaveX(ctx context.Context) []*Product {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProductCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProductCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
