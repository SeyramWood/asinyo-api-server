// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SeyramWood/ent/predicate"
	"github.com/SeyramWood/ent/productcategoryminor"
)

// ProductCategoryMinorDelete is the builder for deleting a ProductCategoryMinor entity.
type ProductCategoryMinorDelete struct {
	config
	hooks    []Hook
	mutation *ProductCategoryMinorMutation
}

// Where appends a list predicates to the ProductCategoryMinorDelete builder.
func (pcmd *ProductCategoryMinorDelete) Where(ps ...predicate.ProductCategoryMinor) *ProductCategoryMinorDelete {
	pcmd.mutation.Where(ps...)
	return pcmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pcmd *ProductCategoryMinorDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pcmd.hooks) == 0 {
		affected, err = pcmd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductCategoryMinorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pcmd.mutation = mutation
			affected, err = pcmd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pcmd.hooks) - 1; i >= 0; i-- {
			if pcmd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pcmd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pcmd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcmd *ProductCategoryMinorDelete) ExecX(ctx context.Context) int {
	n, err := pcmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pcmd *ProductCategoryMinorDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: productcategoryminor.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: productcategoryminor.FieldID,
			},
		},
	}
	if ps := pcmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, pcmd.driver, _spec)
}

// ProductCategoryMinorDeleteOne is the builder for deleting a single ProductCategoryMinor entity.
type ProductCategoryMinorDeleteOne struct {
	pcmd *ProductCategoryMinorDelete
}

// Exec executes the deletion query.
func (pcmdo *ProductCategoryMinorDeleteOne) Exec(ctx context.Context) error {
	n, err := pcmdo.pcmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{productcategoryminor.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pcmdo *ProductCategoryMinorDeleteOne) ExecX(ctx context.Context) {
	pcmdo.pcmd.ExecX(ctx)
}
