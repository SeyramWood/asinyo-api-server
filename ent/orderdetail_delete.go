// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SeyramWood/ent/orderdetail"
	"github.com/SeyramWood/ent/predicate"
)

// OrderDetailDelete is the builder for deleting a OrderDetail entity.
type OrderDetailDelete struct {
	config
	hooks    []Hook
	mutation *OrderDetailMutation
}

// Where appends a list predicates to the OrderDetailDelete builder.
func (odd *OrderDetailDelete) Where(ps ...predicate.OrderDetail) *OrderDetailDelete {
	odd.mutation.Where(ps...)
	return odd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (odd *OrderDetailDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, odd.sqlExec, odd.mutation, odd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (odd *OrderDetailDelete) ExecX(ctx context.Context) int {
	n, err := odd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (odd *OrderDetailDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(orderdetail.Table, sqlgraph.NewFieldSpec(orderdetail.FieldID, field.TypeInt))
	if ps := odd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, odd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	odd.mutation.done = true
	return affected, err
}

// OrderDetailDeleteOne is the builder for deleting a single OrderDetail entity.
type OrderDetailDeleteOne struct {
	odd *OrderDetailDelete
}

// Where appends a list predicates to the OrderDetailDelete builder.
func (oddo *OrderDetailDeleteOne) Where(ps ...predicate.OrderDetail) *OrderDetailDeleteOne {
	oddo.odd.mutation.Where(ps...)
	return oddo
}

// Exec executes the deletion query.
func (oddo *OrderDetailDeleteOne) Exec(ctx context.Context) error {
	n, err := oddo.odd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{orderdetail.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (oddo *OrderDetailDeleteOne) ExecX(ctx context.Context) {
	if err := oddo.Exec(ctx); err != nil {
		panic(err)
	}
}
