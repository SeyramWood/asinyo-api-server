// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SeyramWood/ent/predicate"
	"github.com/SeyramWood/ent/pricemodel"
	"github.com/SeyramWood/ent/product"
)

// PriceModelUpdate is the builder for updating PriceModel entities.
type PriceModelUpdate struct {
	config
	hooks    []Hook
	mutation *PriceModelMutation
}

// Where appends a list predicates to the PriceModelUpdate builder.
func (pmu *PriceModelUpdate) Where(ps ...predicate.PriceModel) *PriceModelUpdate {
	pmu.mutation.Where(ps...)
	return pmu
}

// SetUpdatedAt sets the "updated_at" field.
func (pmu *PriceModelUpdate) SetUpdatedAt(t time.Time) *PriceModelUpdate {
	pmu.mutation.SetUpdatedAt(t)
	return pmu
}

// SetName sets the "name" field.
func (pmu *PriceModelUpdate) SetName(s string) *PriceModelUpdate {
	pmu.mutation.SetName(s)
	return pmu
}

// SetInitials sets the "initials" field.
func (pmu *PriceModelUpdate) SetInitials(s string) *PriceModelUpdate {
	pmu.mutation.SetInitials(s)
	return pmu
}

// SetFormula sets the "formula" field.
func (pmu *PriceModelUpdate) SetFormula(s string) *PriceModelUpdate {
	pmu.mutation.SetFormula(s)
	return pmu
}

// SetAsinyoFormula sets the "asinyo_formula" field.
func (pmu *PriceModelUpdate) SetAsinyoFormula(s string) *PriceModelUpdate {
	pmu.mutation.SetAsinyoFormula(s)
	return pmu
}

// SetNillableAsinyoFormula sets the "asinyo_formula" field if the given value is not nil.
func (pmu *PriceModelUpdate) SetNillableAsinyoFormula(s *string) *PriceModelUpdate {
	if s != nil {
		pmu.SetAsinyoFormula(*s)
	}
	return pmu
}

// ClearAsinyoFormula clears the value of the "asinyo_formula" field.
func (pmu *PriceModelUpdate) ClearAsinyoFormula() *PriceModelUpdate {
	pmu.mutation.ClearAsinyoFormula()
	return pmu
}

// AddModelIDs adds the "model" edge to the Product entity by IDs.
func (pmu *PriceModelUpdate) AddModelIDs(ids ...int) *PriceModelUpdate {
	pmu.mutation.AddModelIDs(ids...)
	return pmu
}

// AddModel adds the "model" edges to the Product entity.
func (pmu *PriceModelUpdate) AddModel(p ...*Product) *PriceModelUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pmu.AddModelIDs(ids...)
}

// Mutation returns the PriceModelMutation object of the builder.
func (pmu *PriceModelUpdate) Mutation() *PriceModelMutation {
	return pmu.mutation
}

// ClearModel clears all "model" edges to the Product entity.
func (pmu *PriceModelUpdate) ClearModel() *PriceModelUpdate {
	pmu.mutation.ClearModel()
	return pmu
}

// RemoveModelIDs removes the "model" edge to Product entities by IDs.
func (pmu *PriceModelUpdate) RemoveModelIDs(ids ...int) *PriceModelUpdate {
	pmu.mutation.RemoveModelIDs(ids...)
	return pmu
}

// RemoveModel removes "model" edges to Product entities.
func (pmu *PriceModelUpdate) RemoveModel(p ...*Product) *PriceModelUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pmu.RemoveModelIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pmu *PriceModelUpdate) Save(ctx context.Context) (int, error) {
	pmu.defaults()
	return withHooks(ctx, pmu.sqlSave, pmu.mutation, pmu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pmu *PriceModelUpdate) SaveX(ctx context.Context) int {
	affected, err := pmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pmu *PriceModelUpdate) Exec(ctx context.Context) error {
	_, err := pmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pmu *PriceModelUpdate) ExecX(ctx context.Context) {
	if err := pmu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pmu *PriceModelUpdate) defaults() {
	if _, ok := pmu.mutation.UpdatedAt(); !ok {
		v := pricemodel.UpdateDefaultUpdatedAt()
		pmu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pmu *PriceModelUpdate) check() error {
	if v, ok := pmu.mutation.Name(); ok {
		if err := pricemodel.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "PriceModel.name": %w`, err)}
		}
	}
	if v, ok := pmu.mutation.Initials(); ok {
		if err := pricemodel.InitialsValidator(v); err != nil {
			return &ValidationError{Name: "initials", err: fmt.Errorf(`ent: validator failed for field "PriceModel.initials": %w`, err)}
		}
	}
	if v, ok := pmu.mutation.Formula(); ok {
		if err := pricemodel.FormulaValidator(v); err != nil {
			return &ValidationError{Name: "formula", err: fmt.Errorf(`ent: validator failed for field "PriceModel.formula": %w`, err)}
		}
	}
	return nil
}

func (pmu *PriceModelUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pmu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(pricemodel.Table, pricemodel.Columns, sqlgraph.NewFieldSpec(pricemodel.FieldID, field.TypeInt))
	if ps := pmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pmu.mutation.UpdatedAt(); ok {
		_spec.SetField(pricemodel.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pmu.mutation.Name(); ok {
		_spec.SetField(pricemodel.FieldName, field.TypeString, value)
	}
	if value, ok := pmu.mutation.Initials(); ok {
		_spec.SetField(pricemodel.FieldInitials, field.TypeString, value)
	}
	if value, ok := pmu.mutation.Formula(); ok {
		_spec.SetField(pricemodel.FieldFormula, field.TypeString, value)
	}
	if value, ok := pmu.mutation.AsinyoFormula(); ok {
		_spec.SetField(pricemodel.FieldAsinyoFormula, field.TypeString, value)
	}
	if pmu.mutation.AsinyoFormulaCleared() {
		_spec.ClearField(pricemodel.FieldAsinyoFormula, field.TypeString)
	}
	if pmu.mutation.ModelCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pmu.mutation.RemovedModelIDs(); len(nodes) > 0 && !pmu.mutation.ModelCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pmu.mutation.ModelIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pricemodel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pmu.mutation.done = true
	return n, nil
}

// PriceModelUpdateOne is the builder for updating a single PriceModel entity.
type PriceModelUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PriceModelMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (pmuo *PriceModelUpdateOne) SetUpdatedAt(t time.Time) *PriceModelUpdateOne {
	pmuo.mutation.SetUpdatedAt(t)
	return pmuo
}

// SetName sets the "name" field.
func (pmuo *PriceModelUpdateOne) SetName(s string) *PriceModelUpdateOne {
	pmuo.mutation.SetName(s)
	return pmuo
}

// SetInitials sets the "initials" field.
func (pmuo *PriceModelUpdateOne) SetInitials(s string) *PriceModelUpdateOne {
	pmuo.mutation.SetInitials(s)
	return pmuo
}

// SetFormula sets the "formula" field.
func (pmuo *PriceModelUpdateOne) SetFormula(s string) *PriceModelUpdateOne {
	pmuo.mutation.SetFormula(s)
	return pmuo
}

// SetAsinyoFormula sets the "asinyo_formula" field.
func (pmuo *PriceModelUpdateOne) SetAsinyoFormula(s string) *PriceModelUpdateOne {
	pmuo.mutation.SetAsinyoFormula(s)
	return pmuo
}

// SetNillableAsinyoFormula sets the "asinyo_formula" field if the given value is not nil.
func (pmuo *PriceModelUpdateOne) SetNillableAsinyoFormula(s *string) *PriceModelUpdateOne {
	if s != nil {
		pmuo.SetAsinyoFormula(*s)
	}
	return pmuo
}

// ClearAsinyoFormula clears the value of the "asinyo_formula" field.
func (pmuo *PriceModelUpdateOne) ClearAsinyoFormula() *PriceModelUpdateOne {
	pmuo.mutation.ClearAsinyoFormula()
	return pmuo
}

// AddModelIDs adds the "model" edge to the Product entity by IDs.
func (pmuo *PriceModelUpdateOne) AddModelIDs(ids ...int) *PriceModelUpdateOne {
	pmuo.mutation.AddModelIDs(ids...)
	return pmuo
}

// AddModel adds the "model" edges to the Product entity.
func (pmuo *PriceModelUpdateOne) AddModel(p ...*Product) *PriceModelUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pmuo.AddModelIDs(ids...)
}

// Mutation returns the PriceModelMutation object of the builder.
func (pmuo *PriceModelUpdateOne) Mutation() *PriceModelMutation {
	return pmuo.mutation
}

// ClearModel clears all "model" edges to the Product entity.
func (pmuo *PriceModelUpdateOne) ClearModel() *PriceModelUpdateOne {
	pmuo.mutation.ClearModel()
	return pmuo
}

// RemoveModelIDs removes the "model" edge to Product entities by IDs.
func (pmuo *PriceModelUpdateOne) RemoveModelIDs(ids ...int) *PriceModelUpdateOne {
	pmuo.mutation.RemoveModelIDs(ids...)
	return pmuo
}

// RemoveModel removes "model" edges to Product entities.
func (pmuo *PriceModelUpdateOne) RemoveModel(p ...*Product) *PriceModelUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pmuo.RemoveModelIDs(ids...)
}

// Where appends a list predicates to the PriceModelUpdate builder.
func (pmuo *PriceModelUpdateOne) Where(ps ...predicate.PriceModel) *PriceModelUpdateOne {
	pmuo.mutation.Where(ps...)
	return pmuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pmuo *PriceModelUpdateOne) Select(field string, fields ...string) *PriceModelUpdateOne {
	pmuo.fields = append([]string{field}, fields...)
	return pmuo
}

// Save executes the query and returns the updated PriceModel entity.
func (pmuo *PriceModelUpdateOne) Save(ctx context.Context) (*PriceModel, error) {
	pmuo.defaults()
	return withHooks(ctx, pmuo.sqlSave, pmuo.mutation, pmuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pmuo *PriceModelUpdateOne) SaveX(ctx context.Context) *PriceModel {
	node, err := pmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pmuo *PriceModelUpdateOne) Exec(ctx context.Context) error {
	_, err := pmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pmuo *PriceModelUpdateOne) ExecX(ctx context.Context) {
	if err := pmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pmuo *PriceModelUpdateOne) defaults() {
	if _, ok := pmuo.mutation.UpdatedAt(); !ok {
		v := pricemodel.UpdateDefaultUpdatedAt()
		pmuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pmuo *PriceModelUpdateOne) check() error {
	if v, ok := pmuo.mutation.Name(); ok {
		if err := pricemodel.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "PriceModel.name": %w`, err)}
		}
	}
	if v, ok := pmuo.mutation.Initials(); ok {
		if err := pricemodel.InitialsValidator(v); err != nil {
			return &ValidationError{Name: "initials", err: fmt.Errorf(`ent: validator failed for field "PriceModel.initials": %w`, err)}
		}
	}
	if v, ok := pmuo.mutation.Formula(); ok {
		if err := pricemodel.FormulaValidator(v); err != nil {
			return &ValidationError{Name: "formula", err: fmt.Errorf(`ent: validator failed for field "PriceModel.formula": %w`, err)}
		}
	}
	return nil
}

func (pmuo *PriceModelUpdateOne) sqlSave(ctx context.Context) (_node *PriceModel, err error) {
	if err := pmuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(pricemodel.Table, pricemodel.Columns, sqlgraph.NewFieldSpec(pricemodel.FieldID, field.TypeInt))
	id, ok := pmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PriceModel.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pricemodel.FieldID)
		for _, f := range fields {
			if !pricemodel.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != pricemodel.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pmuo.mutation.UpdatedAt(); ok {
		_spec.SetField(pricemodel.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pmuo.mutation.Name(); ok {
		_spec.SetField(pricemodel.FieldName, field.TypeString, value)
	}
	if value, ok := pmuo.mutation.Initials(); ok {
		_spec.SetField(pricemodel.FieldInitials, field.TypeString, value)
	}
	if value, ok := pmuo.mutation.Formula(); ok {
		_spec.SetField(pricemodel.FieldFormula, field.TypeString, value)
	}
	if value, ok := pmuo.mutation.AsinyoFormula(); ok {
		_spec.SetField(pricemodel.FieldAsinyoFormula, field.TypeString, value)
	}
	if pmuo.mutation.AsinyoFormulaCleared() {
		_spec.ClearField(pricemodel.FieldAsinyoFormula, field.TypeString)
	}
	if pmuo.mutation.ModelCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pmuo.mutation.RemovedModelIDs(); len(nodes) > 0 && !pmuo.mutation.ModelCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pmuo.mutation.ModelIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PriceModel{config: pmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pricemodel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pmuo.mutation.done = true
	return _node, nil
}
