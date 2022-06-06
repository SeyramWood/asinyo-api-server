// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/SeyramWood/ent"
)

// The AddressFunc type is an adapter to allow the use of ordinary
// function as Address mutator.
type AddressFunc func(context.Context, *ent.AddressMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AddressFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AddressMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AddressMutation", m)
	}
	return f(ctx, mv)
}

// The AdminFunc type is an adapter to allow the use of ordinary
// function as Admin mutator.
type AdminFunc func(context.Context, *ent.AdminMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AdminFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AdminMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AdminMutation", m)
	}
	return f(ctx, mv)
}

// The AgentFunc type is an adapter to allow the use of ordinary
// function as Agent mutator.
type AgentFunc func(context.Context, *ent.AgentMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AgentFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AgentMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AgentMutation", m)
	}
	return f(ctx, mv)
}

// The BasketFunc type is an adapter to allow the use of ordinary
// function as Basket mutator.
type BasketFunc func(context.Context, *ent.BasketMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f BasketFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.BasketMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.BasketMutation", m)
	}
	return f(ctx, mv)
}

// The CustomerFunc type is an adapter to allow the use of ordinary
// function as Customer mutator.
type CustomerFunc func(context.Context, *ent.CustomerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CustomerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CustomerMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CustomerMutation", m)
	}
	return f(ctx, mv)
}

// The FavouriteFunc type is an adapter to allow the use of ordinary
// function as Favourite mutator.
type FavouriteFunc func(context.Context, *ent.FavouriteMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FavouriteFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.FavouriteMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FavouriteMutation", m)
	}
	return f(ctx, mv)
}

// The MerchantFunc type is an adapter to allow the use of ordinary
// function as Merchant mutator.
type MerchantFunc func(context.Context, *ent.MerchantMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MerchantFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.MerchantMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MerchantMutation", m)
	}
	return f(ctx, mv)
}

// The MerchantStoreFunc type is an adapter to allow the use of ordinary
// function as MerchantStore mutator.
type MerchantStoreFunc func(context.Context, *ent.MerchantStoreMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MerchantStoreFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.MerchantStoreMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MerchantStoreMutation", m)
	}
	return f(ctx, mv)
}

// The OrderFunc type is an adapter to allow the use of ordinary
// function as Order mutator.
type OrderFunc func(context.Context, *ent.OrderMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f OrderFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.OrderMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.OrderMutation", m)
	}
	return f(ctx, mv)
}

// The ProductFunc type is an adapter to allow the use of ordinary
// function as Product mutator.
type ProductFunc func(context.Context, *ent.ProductMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProductFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ProductMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProductMutation", m)
	}
	return f(ctx, mv)
}

// The ProductCategoryMajorFunc type is an adapter to allow the use of ordinary
// function as ProductCategoryMajor mutator.
type ProductCategoryMajorFunc func(context.Context, *ent.ProductCategoryMajorMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProductCategoryMajorFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ProductCategoryMajorMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProductCategoryMajorMutation", m)
	}
	return f(ctx, mv)
}

// The ProductCategoryMinorFunc type is an adapter to allow the use of ordinary
// function as ProductCategoryMinor mutator.
type ProductCategoryMinorFunc func(context.Context, *ent.ProductCategoryMinorMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProductCategoryMinorFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ProductCategoryMinorMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProductCategoryMinorMutation", m)
	}
	return f(ctx, mv)
}

// The RetailMerchantFunc type is an adapter to allow the use of ordinary
// function as RetailMerchant mutator.
type RetailMerchantFunc func(context.Context, *ent.RetailMerchantMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RetailMerchantFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.RetailMerchantMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RetailMerchantMutation", m)
	}
	return f(ctx, mv)
}

// The SupplierMerchantFunc type is an adapter to allow the use of ordinary
// function as SupplierMerchant mutator.
type SupplierMerchantFunc func(context.Context, *ent.SupplierMerchantMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SupplierMerchantFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.SupplierMerchantMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SupplierMerchantMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
//
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
//
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
//
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
//
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
