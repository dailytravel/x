package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"

	"github.com/dailytravel/x/sales/graph/model"
)

// ID is the resolver for the id field.
func (r *inventoryResolver) ID(ctx context.Context, obj *model.Inventory) (string, error) {
	return obj.ID.Hex(), nil
}

// Product is the resolver for the product field.
func (r *inventoryResolver) Product(ctx context.Context, obj *model.Inventory) (*model.Product, error) {
	panic(fmt.Errorf("not implemented: Product - product"))
}

// Date is the resolver for the date field.
func (r *inventoryResolver) Date(ctx context.Context, obj *model.Inventory) (int, error) {
	panic(fmt.Errorf("not implemented: Date - date"))
}

// CreatedBy is the resolver for the created_by field.
func (r *inventoryResolver) CreatedBy(ctx context.Context, obj *model.Inventory) (*string, error) {
	panic(fmt.Errorf("not implemented: CreatedBy - created_by"))
}

// UpdatedBy is the resolver for the updated_by field.
func (r *inventoryResolver) UpdatedBy(ctx context.Context, obj *model.Inventory) (*string, error) {
	panic(fmt.Errorf("not implemented: UpdatedBy - updated_by"))
}

// CreateInventory is the resolver for the createInventory field.
func (r *mutationResolver) CreateInventory(ctx context.Context, input model.NewInventory) (*model.Inventory, error) {
	panic(fmt.Errorf("not implemented: CreateInventory - createInventory"))
}

// UpdateInventory is the resolver for the updateInventory field.
func (r *mutationResolver) UpdateInventory(ctx context.Context, id string, input model.UpdateInventory) (*model.Inventory, error) {
	panic(fmt.Errorf("not implemented: UpdateInventory - updateInventory"))
}

// DeleteInventory is the resolver for the deleteInventory field.
func (r *mutationResolver) DeleteInventory(ctx context.Context, id string) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: DeleteInventory - deleteInventory"))
}

// DeleteInventories is the resolver for the deleteInventories field.
func (r *mutationResolver) DeleteInventories(ctx context.Context, ids []string) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: DeleteInventories - deleteInventories"))
}

// Inventories is the resolver for the inventories field.
func (r *queryResolver) Inventories(ctx context.Context, product *string, date int) (*model.Inventories, error) {
	panic(fmt.Errorf("not implemented: Inventories - inventories"))
}

// Inventory is the resolver for the inventory field.
func (r *queryResolver) Inventory(ctx context.Context, id string) (*model.Inventory, error) {
	panic(fmt.Errorf("not implemented: Inventory - inventory"))
}

// Inventory returns InventoryResolver implementation.
func (r *Resolver) Inventory() InventoryResolver { return &inventoryResolver{r} }

type inventoryResolver struct{ *Resolver }
