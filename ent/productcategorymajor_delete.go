// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SeyramWood/ent/predicate"
	"github.com/SeyramWood/ent/productcategorymajor"
)

// ProductCategoryMajorDelete is the builder for deleting a ProductCategoryMajor entity.
type ProductCategoryMajorDelete struct {
	config
	hooks    []Hook
	mutation *ProductCategoryMajorMutation
}

// Where appends a list predicates to the ProductCategoryMajorDelete builder.
func (pcmd *ProductCategoryMajorDelete) Where(ps ...predicate.ProductCategoryMajor) *ProductCategoryMajorDelete {
	pcmd.mutation.Where(ps...)
	return pcmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pcmd *ProductCategoryMajorDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pcmd.sqlExec, pcmd.mutation, pcmd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pcmd *ProductCategoryMajorDelete) ExecX(ctx context.Context) int {
	n, err := pcmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pcmd *ProductCategoryMajorDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(productcategorymajor.Table, sqlgraph.NewFieldSpec(productcategorymajor.FieldID, field.TypeInt))
	if ps := pcmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pcmd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pcmd.mutation.done = true
	return affected, err
}

// ProductCategoryMajorDeleteOne is the builder for deleting a single ProductCategoryMajor entity.
type ProductCategoryMajorDeleteOne struct {
	pcmd *ProductCategoryMajorDelete
}

// Where appends a list predicates to the ProductCategoryMajorDelete builder.
func (pcmdo *ProductCategoryMajorDeleteOne) Where(ps ...predicate.ProductCategoryMajor) *ProductCategoryMajorDeleteOne {
	pcmdo.pcmd.mutation.Where(ps...)
	return pcmdo
}

// Exec executes the deletion query.
func (pcmdo *ProductCategoryMajorDeleteOne) Exec(ctx context.Context) error {
	n, err := pcmdo.pcmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{productcategorymajor.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pcmdo *ProductCategoryMajorDeleteOne) ExecX(ctx context.Context) {
	if err := pcmdo.Exec(ctx); err != nil {
		panic(err)
	}
}
