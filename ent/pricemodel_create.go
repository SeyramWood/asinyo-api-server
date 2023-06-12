// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SeyramWood/ent/pricemodel"
	"github.com/SeyramWood/ent/product"
)

// PriceModelCreate is the builder for creating a PriceModel entity.
type PriceModelCreate struct {
	config
	mutation *PriceModelMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (pmc *PriceModelCreate) SetCreatedAt(t time.Time) *PriceModelCreate {
	pmc.mutation.SetCreatedAt(t)
	return pmc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pmc *PriceModelCreate) SetNillableCreatedAt(t *time.Time) *PriceModelCreate {
	if t != nil {
		pmc.SetCreatedAt(*t)
	}
	return pmc
}

// SetUpdatedAt sets the "updated_at" field.
func (pmc *PriceModelCreate) SetUpdatedAt(t time.Time) *PriceModelCreate {
	pmc.mutation.SetUpdatedAt(t)
	return pmc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pmc *PriceModelCreate) SetNillableUpdatedAt(t *time.Time) *PriceModelCreate {
	if t != nil {
		pmc.SetUpdatedAt(*t)
	}
	return pmc
}

// SetName sets the "name" field.
func (pmc *PriceModelCreate) SetName(s string) *PriceModelCreate {
	pmc.mutation.SetName(s)
	return pmc
}

// SetInitials sets the "initials" field.
func (pmc *PriceModelCreate) SetInitials(s string) *PriceModelCreate {
	pmc.mutation.SetInitials(s)
	return pmc
}

// SetFormula sets the "formula" field.
func (pmc *PriceModelCreate) SetFormula(s string) *PriceModelCreate {
	pmc.mutation.SetFormula(s)
	return pmc
}

// SetAsinyoFormula sets the "asinyo_formula" field.
func (pmc *PriceModelCreate) SetAsinyoFormula(s string) *PriceModelCreate {
	pmc.mutation.SetAsinyoFormula(s)
	return pmc
}

// SetNillableAsinyoFormula sets the "asinyo_formula" field if the given value is not nil.
func (pmc *PriceModelCreate) SetNillableAsinyoFormula(s *string) *PriceModelCreate {
	if s != nil {
		pmc.SetAsinyoFormula(*s)
	}
	return pmc
}

// AddModelIDs adds the "model" edge to the Product entity by IDs.
func (pmc *PriceModelCreate) AddModelIDs(ids ...int) *PriceModelCreate {
	pmc.mutation.AddModelIDs(ids...)
	return pmc
}

// AddModel adds the "model" edges to the Product entity.
func (pmc *PriceModelCreate) AddModel(p ...*Product) *PriceModelCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pmc.AddModelIDs(ids...)
}

// Mutation returns the PriceModelMutation object of the builder.
func (pmc *PriceModelCreate) Mutation() *PriceModelMutation {
	return pmc.mutation
}

// Save creates the PriceModel in the database.
func (pmc *PriceModelCreate) Save(ctx context.Context) (*PriceModel, error) {
	pmc.defaults()
	return withHooks(ctx, pmc.sqlSave, pmc.mutation, pmc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pmc *PriceModelCreate) SaveX(ctx context.Context) *PriceModel {
	v, err := pmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pmc *PriceModelCreate) Exec(ctx context.Context) error {
	_, err := pmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pmc *PriceModelCreate) ExecX(ctx context.Context) {
	if err := pmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pmc *PriceModelCreate) defaults() {
	if _, ok := pmc.mutation.CreatedAt(); !ok {
		v := pricemodel.DefaultCreatedAt()
		pmc.mutation.SetCreatedAt(v)
	}
	if _, ok := pmc.mutation.UpdatedAt(); !ok {
		v := pricemodel.DefaultUpdatedAt()
		pmc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pmc.mutation.AsinyoFormula(); !ok {
		v := pricemodel.DefaultAsinyoFormula
		pmc.mutation.SetAsinyoFormula(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pmc *PriceModelCreate) check() error {
	if _, ok := pmc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "PriceModel.created_at"`)}
	}
	if _, ok := pmc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "PriceModel.updated_at"`)}
	}
	if _, ok := pmc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "PriceModel.name"`)}
	}
	if v, ok := pmc.mutation.Name(); ok {
		if err := pricemodel.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "PriceModel.name": %w`, err)}
		}
	}
	if _, ok := pmc.mutation.Initials(); !ok {
		return &ValidationError{Name: "initials", err: errors.New(`ent: missing required field "PriceModel.initials"`)}
	}
	if v, ok := pmc.mutation.Initials(); ok {
		if err := pricemodel.InitialsValidator(v); err != nil {
			return &ValidationError{Name: "initials", err: fmt.Errorf(`ent: validator failed for field "PriceModel.initials": %w`, err)}
		}
	}
	if _, ok := pmc.mutation.Formula(); !ok {
		return &ValidationError{Name: "formula", err: errors.New(`ent: missing required field "PriceModel.formula"`)}
	}
	if v, ok := pmc.mutation.Formula(); ok {
		if err := pricemodel.FormulaValidator(v); err != nil {
			return &ValidationError{Name: "formula", err: fmt.Errorf(`ent: validator failed for field "PriceModel.formula": %w`, err)}
		}
	}
	return nil
}

func (pmc *PriceModelCreate) sqlSave(ctx context.Context) (*PriceModel, error) {
	if err := pmc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pmc.mutation.id = &_node.ID
	pmc.mutation.done = true
	return _node, nil
}

func (pmc *PriceModelCreate) createSpec() (*PriceModel, *sqlgraph.CreateSpec) {
	var (
		_node = &PriceModel{config: pmc.config}
		_spec = sqlgraph.NewCreateSpec(pricemodel.Table, sqlgraph.NewFieldSpec(pricemodel.FieldID, field.TypeInt))
	)
	if value, ok := pmc.mutation.CreatedAt(); ok {
		_spec.SetField(pricemodel.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pmc.mutation.UpdatedAt(); ok {
		_spec.SetField(pricemodel.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pmc.mutation.Name(); ok {
		_spec.SetField(pricemodel.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pmc.mutation.Initials(); ok {
		_spec.SetField(pricemodel.FieldInitials, field.TypeString, value)
		_node.Initials = value
	}
	if value, ok := pmc.mutation.Formula(); ok {
		_spec.SetField(pricemodel.FieldFormula, field.TypeString, value)
		_node.Formula = value
	}
	if value, ok := pmc.mutation.AsinyoFormula(); ok {
		_spec.SetField(pricemodel.FieldAsinyoFormula, field.TypeString, value)
		_node.AsinyoFormula = value
	}
	if nodes := pmc.mutation.ModelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pricemodel.ModelTable,
			Columns: []string{pricemodel.ModelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PriceModelCreateBulk is the builder for creating many PriceModel entities in bulk.
type PriceModelCreateBulk struct {
	config
	builders []*PriceModelCreate
}

// Save creates the PriceModel entities in the database.
func (pmcb *PriceModelCreateBulk) Save(ctx context.Context) ([]*PriceModel, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pmcb.builders))
	nodes := make([]*PriceModel, len(pmcb.builders))
	mutators := make([]Mutator, len(pmcb.builders))
	for i := range pmcb.builders {
		func(i int, root context.Context) {
			builder := pmcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PriceModelMutation)
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
					_, err = mutators[i+1].Mutate(root, pmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pmcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pmcb *PriceModelCreateBulk) SaveX(ctx context.Context) []*PriceModel {
	v, err := pmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pmcb *PriceModelCreateBulk) Exec(ctx context.Context) error {
	_, err := pmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pmcb *PriceModelCreateBulk) ExecX(ctx context.Context) {
	if err := pmcb.Exec(ctx); err != nil {
		panic(err)
	}
}
