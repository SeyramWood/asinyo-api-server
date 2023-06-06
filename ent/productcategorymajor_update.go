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
	"github.com/SeyramWood/ent/product"
	"github.com/SeyramWood/ent/productcategorymajor"
	"github.com/SeyramWood/ent/productcategoryminor"
)

// ProductCategoryMajorUpdate is the builder for updating ProductCategoryMajor entities.
type ProductCategoryMajorUpdate struct {
	config
	hooks    []Hook
	mutation *ProductCategoryMajorMutation
}

// Where appends a list predicates to the ProductCategoryMajorUpdate builder.
func (pcmu *ProductCategoryMajorUpdate) Where(ps ...predicate.ProductCategoryMajor) *ProductCategoryMajorUpdate {
	pcmu.mutation.Where(ps...)
	return pcmu
}

// SetUpdatedAt sets the "updated_at" field.
func (pcmu *ProductCategoryMajorUpdate) SetUpdatedAt(t time.Time) *ProductCategoryMajorUpdate {
	pcmu.mutation.SetUpdatedAt(t)
	return pcmu
}

// SetCategory sets the "category" field.
func (pcmu *ProductCategoryMajorUpdate) SetCategory(s string) *ProductCategoryMajorUpdate {
	pcmu.mutation.SetCategory(s)
	return pcmu
}

// SetSlug sets the "slug" field.
func (pcmu *ProductCategoryMajorUpdate) SetSlug(s string) *ProductCategoryMajorUpdate {
	pcmu.mutation.SetSlug(s)
	return pcmu
}

// AddMinorIDs adds the "minors" edge to the ProductCategoryMinor entity by IDs.
func (pcmu *ProductCategoryMajorUpdate) AddMinorIDs(ids ...int) *ProductCategoryMajorUpdate {
	pcmu.mutation.AddMinorIDs(ids...)
	return pcmu
}

// AddMinors adds the "minors" edges to the ProductCategoryMinor entity.
func (pcmu *ProductCategoryMajorUpdate) AddMinors(p ...*ProductCategoryMinor) *ProductCategoryMajorUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pcmu.AddMinorIDs(ids...)
}

// AddProductIDs adds the "products" edge to the Product entity by IDs.
func (pcmu *ProductCategoryMajorUpdate) AddProductIDs(ids ...int) *ProductCategoryMajorUpdate {
	pcmu.mutation.AddProductIDs(ids...)
	return pcmu
}

// AddProducts adds the "products" edges to the Product entity.
func (pcmu *ProductCategoryMajorUpdate) AddProducts(p ...*Product) *ProductCategoryMajorUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pcmu.AddProductIDs(ids...)
}

// Mutation returns the ProductCategoryMajorMutation object of the builder.
func (pcmu *ProductCategoryMajorUpdate) Mutation() *ProductCategoryMajorMutation {
	return pcmu.mutation
}

// ClearMinors clears all "minors" edges to the ProductCategoryMinor entity.
func (pcmu *ProductCategoryMajorUpdate) ClearMinors() *ProductCategoryMajorUpdate {
	pcmu.mutation.ClearMinors()
	return pcmu
}

// RemoveMinorIDs removes the "minors" edge to ProductCategoryMinor entities by IDs.
func (pcmu *ProductCategoryMajorUpdate) RemoveMinorIDs(ids ...int) *ProductCategoryMajorUpdate {
	pcmu.mutation.RemoveMinorIDs(ids...)
	return pcmu
}

// RemoveMinors removes "minors" edges to ProductCategoryMinor entities.
func (pcmu *ProductCategoryMajorUpdate) RemoveMinors(p ...*ProductCategoryMinor) *ProductCategoryMajorUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pcmu.RemoveMinorIDs(ids...)
}

// ClearProducts clears all "products" edges to the Product entity.
func (pcmu *ProductCategoryMajorUpdate) ClearProducts() *ProductCategoryMajorUpdate {
	pcmu.mutation.ClearProducts()
	return pcmu
}

// RemoveProductIDs removes the "products" edge to Product entities by IDs.
func (pcmu *ProductCategoryMajorUpdate) RemoveProductIDs(ids ...int) *ProductCategoryMajorUpdate {
	pcmu.mutation.RemoveProductIDs(ids...)
	return pcmu
}

// RemoveProducts removes "products" edges to Product entities.
func (pcmu *ProductCategoryMajorUpdate) RemoveProducts(p ...*Product) *ProductCategoryMajorUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pcmu.RemoveProductIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pcmu *ProductCategoryMajorUpdate) Save(ctx context.Context) (int, error) {
	pcmu.defaults()
	return withHooks(ctx, pcmu.sqlSave, pcmu.mutation, pcmu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pcmu *ProductCategoryMajorUpdate) SaveX(ctx context.Context) int {
	affected, err := pcmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pcmu *ProductCategoryMajorUpdate) Exec(ctx context.Context) error {
	_, err := pcmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcmu *ProductCategoryMajorUpdate) ExecX(ctx context.Context) {
	if err := pcmu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pcmu *ProductCategoryMajorUpdate) defaults() {
	if _, ok := pcmu.mutation.UpdatedAt(); !ok {
		v := productcategorymajor.UpdateDefaultUpdatedAt()
		pcmu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pcmu *ProductCategoryMajorUpdate) check() error {
	if v, ok := pcmu.mutation.Category(); ok {
		if err := productcategorymajor.CategoryValidator(v); err != nil {
			return &ValidationError{Name: "category", err: fmt.Errorf(`ent: validator failed for field "ProductCategoryMajor.category": %w`, err)}
		}
	}
	if v, ok := pcmu.mutation.Slug(); ok {
		if err := productcategorymajor.SlugValidator(v); err != nil {
			return &ValidationError{Name: "slug", err: fmt.Errorf(`ent: validator failed for field "ProductCategoryMajor.slug": %w`, err)}
		}
	}
	return nil
}

func (pcmu *ProductCategoryMajorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pcmu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(productcategorymajor.Table, productcategorymajor.Columns, sqlgraph.NewFieldSpec(productcategorymajor.FieldID, field.TypeInt))
	if ps := pcmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pcmu.mutation.UpdatedAt(); ok {
		_spec.SetField(productcategorymajor.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pcmu.mutation.Category(); ok {
		_spec.SetField(productcategorymajor.FieldCategory, field.TypeString, value)
	}
	if value, ok := pcmu.mutation.Slug(); ok {
		_spec.SetField(productcategorymajor.FieldSlug, field.TypeString, value)
	}
	if pcmu.mutation.MinorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productcategorymajor.MinorsTable,
			Columns: []string{productcategorymajor.MinorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(productcategoryminor.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcmu.mutation.RemovedMinorsIDs(); len(nodes) > 0 && !pcmu.mutation.MinorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productcategorymajor.MinorsTable,
			Columns: []string{productcategorymajor.MinorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(productcategoryminor.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcmu.mutation.MinorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productcategorymajor.MinorsTable,
			Columns: []string{productcategorymajor.MinorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(productcategoryminor.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pcmu.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productcategorymajor.ProductsTable,
			Columns: []string{productcategorymajor.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcmu.mutation.RemovedProductsIDs(); len(nodes) > 0 && !pcmu.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productcategorymajor.ProductsTable,
			Columns: []string{productcategorymajor.ProductsColumn},
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
	if nodes := pcmu.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productcategorymajor.ProductsTable,
			Columns: []string{productcategorymajor.ProductsColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, pcmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{productcategorymajor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pcmu.mutation.done = true
	return n, nil
}

// ProductCategoryMajorUpdateOne is the builder for updating a single ProductCategoryMajor entity.
type ProductCategoryMajorUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProductCategoryMajorMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (pcmuo *ProductCategoryMajorUpdateOne) SetUpdatedAt(t time.Time) *ProductCategoryMajorUpdateOne {
	pcmuo.mutation.SetUpdatedAt(t)
	return pcmuo
}

// SetCategory sets the "category" field.
func (pcmuo *ProductCategoryMajorUpdateOne) SetCategory(s string) *ProductCategoryMajorUpdateOne {
	pcmuo.mutation.SetCategory(s)
	return pcmuo
}

// SetSlug sets the "slug" field.
func (pcmuo *ProductCategoryMajorUpdateOne) SetSlug(s string) *ProductCategoryMajorUpdateOne {
	pcmuo.mutation.SetSlug(s)
	return pcmuo
}

// AddMinorIDs adds the "minors" edge to the ProductCategoryMinor entity by IDs.
func (pcmuo *ProductCategoryMajorUpdateOne) AddMinorIDs(ids ...int) *ProductCategoryMajorUpdateOne {
	pcmuo.mutation.AddMinorIDs(ids...)
	return pcmuo
}

// AddMinors adds the "minors" edges to the ProductCategoryMinor entity.
func (pcmuo *ProductCategoryMajorUpdateOne) AddMinors(p ...*ProductCategoryMinor) *ProductCategoryMajorUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pcmuo.AddMinorIDs(ids...)
}

// AddProductIDs adds the "products" edge to the Product entity by IDs.
func (pcmuo *ProductCategoryMajorUpdateOne) AddProductIDs(ids ...int) *ProductCategoryMajorUpdateOne {
	pcmuo.mutation.AddProductIDs(ids...)
	return pcmuo
}

// AddProducts adds the "products" edges to the Product entity.
func (pcmuo *ProductCategoryMajorUpdateOne) AddProducts(p ...*Product) *ProductCategoryMajorUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pcmuo.AddProductIDs(ids...)
}

// Mutation returns the ProductCategoryMajorMutation object of the builder.
func (pcmuo *ProductCategoryMajorUpdateOne) Mutation() *ProductCategoryMajorMutation {
	return pcmuo.mutation
}

// ClearMinors clears all "minors" edges to the ProductCategoryMinor entity.
func (pcmuo *ProductCategoryMajorUpdateOne) ClearMinors() *ProductCategoryMajorUpdateOne {
	pcmuo.mutation.ClearMinors()
	return pcmuo
}

// RemoveMinorIDs removes the "minors" edge to ProductCategoryMinor entities by IDs.
func (pcmuo *ProductCategoryMajorUpdateOne) RemoveMinorIDs(ids ...int) *ProductCategoryMajorUpdateOne {
	pcmuo.mutation.RemoveMinorIDs(ids...)
	return pcmuo
}

// RemoveMinors removes "minors" edges to ProductCategoryMinor entities.
func (pcmuo *ProductCategoryMajorUpdateOne) RemoveMinors(p ...*ProductCategoryMinor) *ProductCategoryMajorUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pcmuo.RemoveMinorIDs(ids...)
}

// ClearProducts clears all "products" edges to the Product entity.
func (pcmuo *ProductCategoryMajorUpdateOne) ClearProducts() *ProductCategoryMajorUpdateOne {
	pcmuo.mutation.ClearProducts()
	return pcmuo
}

// RemoveProductIDs removes the "products" edge to Product entities by IDs.
func (pcmuo *ProductCategoryMajorUpdateOne) RemoveProductIDs(ids ...int) *ProductCategoryMajorUpdateOne {
	pcmuo.mutation.RemoveProductIDs(ids...)
	return pcmuo
}

// RemoveProducts removes "products" edges to Product entities.
func (pcmuo *ProductCategoryMajorUpdateOne) RemoveProducts(p ...*Product) *ProductCategoryMajorUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pcmuo.RemoveProductIDs(ids...)
}

// Where appends a list predicates to the ProductCategoryMajorUpdate builder.
func (pcmuo *ProductCategoryMajorUpdateOne) Where(ps ...predicate.ProductCategoryMajor) *ProductCategoryMajorUpdateOne {
	pcmuo.mutation.Where(ps...)
	return pcmuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pcmuo *ProductCategoryMajorUpdateOne) Select(field string, fields ...string) *ProductCategoryMajorUpdateOne {
	pcmuo.fields = append([]string{field}, fields...)
	return pcmuo
}

// Save executes the query and returns the updated ProductCategoryMajor entity.
func (pcmuo *ProductCategoryMajorUpdateOne) Save(ctx context.Context) (*ProductCategoryMajor, error) {
	pcmuo.defaults()
	return withHooks(ctx, pcmuo.sqlSave, pcmuo.mutation, pcmuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pcmuo *ProductCategoryMajorUpdateOne) SaveX(ctx context.Context) *ProductCategoryMajor {
	node, err := pcmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pcmuo *ProductCategoryMajorUpdateOne) Exec(ctx context.Context) error {
	_, err := pcmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcmuo *ProductCategoryMajorUpdateOne) ExecX(ctx context.Context) {
	if err := pcmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pcmuo *ProductCategoryMajorUpdateOne) defaults() {
	if _, ok := pcmuo.mutation.UpdatedAt(); !ok {
		v := productcategorymajor.UpdateDefaultUpdatedAt()
		pcmuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pcmuo *ProductCategoryMajorUpdateOne) check() error {
	if v, ok := pcmuo.mutation.Category(); ok {
		if err := productcategorymajor.CategoryValidator(v); err != nil {
			return &ValidationError{Name: "category", err: fmt.Errorf(`ent: validator failed for field "ProductCategoryMajor.category": %w`, err)}
		}
	}
	if v, ok := pcmuo.mutation.Slug(); ok {
		if err := productcategorymajor.SlugValidator(v); err != nil {
			return &ValidationError{Name: "slug", err: fmt.Errorf(`ent: validator failed for field "ProductCategoryMajor.slug": %w`, err)}
		}
	}
	return nil
}

func (pcmuo *ProductCategoryMajorUpdateOne) sqlSave(ctx context.Context) (_node *ProductCategoryMajor, err error) {
	if err := pcmuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(productcategorymajor.Table, productcategorymajor.Columns, sqlgraph.NewFieldSpec(productcategorymajor.FieldID, field.TypeInt))
	id, ok := pcmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ProductCategoryMajor.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pcmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, productcategorymajor.FieldID)
		for _, f := range fields {
			if !productcategorymajor.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != productcategorymajor.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pcmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pcmuo.mutation.UpdatedAt(); ok {
		_spec.SetField(productcategorymajor.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pcmuo.mutation.Category(); ok {
		_spec.SetField(productcategorymajor.FieldCategory, field.TypeString, value)
	}
	if value, ok := pcmuo.mutation.Slug(); ok {
		_spec.SetField(productcategorymajor.FieldSlug, field.TypeString, value)
	}
	if pcmuo.mutation.MinorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productcategorymajor.MinorsTable,
			Columns: []string{productcategorymajor.MinorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(productcategoryminor.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcmuo.mutation.RemovedMinorsIDs(); len(nodes) > 0 && !pcmuo.mutation.MinorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productcategorymajor.MinorsTable,
			Columns: []string{productcategorymajor.MinorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(productcategoryminor.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcmuo.mutation.MinorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productcategorymajor.MinorsTable,
			Columns: []string{productcategorymajor.MinorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(productcategoryminor.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pcmuo.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productcategorymajor.ProductsTable,
			Columns: []string{productcategorymajor.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcmuo.mutation.RemovedProductsIDs(); len(nodes) > 0 && !pcmuo.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productcategorymajor.ProductsTable,
			Columns: []string{productcategorymajor.ProductsColumn},
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
	if nodes := pcmuo.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productcategorymajor.ProductsTable,
			Columns: []string{productcategorymajor.ProductsColumn},
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
	_node = &ProductCategoryMajor{config: pcmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pcmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{productcategorymajor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pcmuo.mutation.done = true
	return _node, nil
}
