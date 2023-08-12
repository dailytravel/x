package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"
	"time"

	"github.com/dailytravel/x/sales/graph/model"
)

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input model.NewOrder) (*model.Order, error) {
	panic(fmt.Errorf("not implemented: CreateOrder - createOrder"))
}

// UpdateOrder is the resolver for the updateOrder field.
func (r *mutationResolver) UpdateOrder(ctx context.Context, id string, input model.UpdateOrder) (*model.Order, error) {
	panic(fmt.Errorf("not implemented: UpdateOrder - updateOrder"))
}

// DeleteOrder is the resolver for the deleteOrder field.
func (r *mutationResolver) DeleteOrder(ctx context.Context, id string) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: DeleteOrder - deleteOrder"))
}

// DeleteOrders is the resolver for the deleteOrders field.
func (r *mutationResolver) DeleteOrders(ctx context.Context, ids []string) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: DeleteOrders - deleteOrders"))
}

// ID is the resolver for the id field.
func (r *orderResolver) ID(ctx context.Context, obj *model.Order) (string, error) {
	return obj.ID.Hex(), nil
}

// User is the resolver for the user field.
func (r *orderResolver) User(ctx context.Context, obj *model.Order) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Total is the resolver for the total field.
func (r *orderResolver) Total(ctx context.Context, obj *model.Order) (float64, error) {
	panic(fmt.Errorf("not implemented: Total - total"))
}

// CancelledAt is the resolver for the cancelled_at field.
func (r *orderResolver) CancelledAt(ctx context.Context, obj *model.Order) (*int, error) {
	panic(fmt.Errorf("not implemented: CancelledAt - cancelled_at"))
}

// CreatedAt is the resolver for the created_at field.
func (r *orderResolver) CreatedAt(ctx context.Context, obj *model.Order) (string, error) {
	return time.Unix(int64(obj.CreatedAt.T), 0).Format(time.RFC3339), nil
}

// UpdatedAt is the resolver for the updated_at field.
func (r *orderResolver) UpdatedAt(ctx context.Context, obj *model.Order) (string, error) {
	return time.Unix(int64(obj.UpdatedAt.T), 0).Format(time.RFC3339), nil
}

// Order is the resolver for the order field.
func (r *queryResolver) Order(ctx context.Context, id string) (*model.Order, error) {
	panic(fmt.Errorf("not implemented: Order - order"))
}

// Orders is the resolver for the orders field.
func (r *queryResolver) Orders(ctx context.Context, args map[string]interface{}) (*model.Orders, error) {
	panic(fmt.Errorf("not implemented: Orders - orders"))
}

// Order returns OrderResolver implementation.
func (r *Resolver) Order() OrderResolver { return &orderResolver{r} }

type orderResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *orderResolver) Owner(ctx context.Context, obj *model.Order) (*model.User, error) {
	panic(fmt.Errorf("not implemented: Owner - owner"))
}
