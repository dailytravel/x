package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"

	"github.com/dailytravel/x/finance/graph/model"
)

// ID is the resolver for the id field.
func (r *expenseResolver) ID(ctx context.Context, obj *model.Expense) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// Followers is the resolver for the followers field.
func (r *expenseResolver) Followers(ctx context.Context, obj *model.Expense) ([]*model.Follow, error) {
	panic(fmt.Errorf("not implemented: Followers - followers"))
}

// Comments is the resolver for the comments field.
func (r *expenseResolver) Comments(ctx context.Context, obj *model.Expense) ([]*model.Comment, error) {
	panic(fmt.Errorf("not implemented: Comments - comments"))
}

// Metadata is the resolver for the metadata field.
func (r *expenseResolver) Metadata(ctx context.Context, obj *model.Expense) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// Date is the resolver for the date field.
func (r *expenseResolver) Date(ctx context.Context, obj *model.Expense) (string, error) {
	panic(fmt.Errorf("not implemented: Date - date"))
}

// CreatedAt is the resolver for the created_at field.
func (r *expenseResolver) CreatedAt(ctx context.Context, obj *model.Expense) (string, error) {
	panic(fmt.Errorf("not implemented: CreatedAt - created_at"))
}

// UpdatedAt is the resolver for the updated_at field.
func (r *expenseResolver) UpdatedAt(ctx context.Context, obj *model.Expense) (string, error) {
	panic(fmt.Errorf("not implemented: UpdatedAt - updated_at"))
}

// UID is the resolver for the uid field.
func (r *expenseResolver) UID(ctx context.Context, obj *model.Expense) (string, error) {
	return obj.ID.Hex(), nil
}

// CreatedBy is the resolver for the created_by field.
func (r *expenseResolver) CreatedBy(ctx context.Context, obj *model.Expense) (*string, error) {
	panic(fmt.Errorf("not implemented: CreatedBy - created_by"))
}

// UpdatedBy is the resolver for the updated_by field.
func (r *expenseResolver) UpdatedBy(ctx context.Context, obj *model.Expense) (*string, error) {
	panic(fmt.Errorf("not implemented: UpdatedBy - updated_by"))
}

// CreateExpense is the resolver for the createExpense field.
func (r *mutationResolver) CreateExpense(ctx context.Context, input model.NewExpense) (*model.Expense, error) {
	panic(fmt.Errorf("not implemented: CreateExpense - createExpense"))
}

// UpdateExpense is the resolver for the updateExpense field.
func (r *mutationResolver) UpdateExpense(ctx context.Context, id string, input model.UpdateExpense) (*model.Expense, error) {
	panic(fmt.Errorf("not implemented: UpdateExpense - updateExpense"))
}

// DeleteExpense is the resolver for the deleteExpense field.
func (r *mutationResolver) DeleteExpense(ctx context.Context, id string) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: DeleteExpense - deleteExpense"))
}

// DeleteExpenses is the resolver for the deleteExpenses field.
func (r *mutationResolver) DeleteExpenses(ctx context.Context, ids []string) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: DeleteExpenses - deleteExpenses"))
}

// Expenses is the resolver for the expenses field.
func (r *queryResolver) Expenses(ctx context.Context, args map[string]interface{}) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: Expenses - expenses"))
}

// Expense is the resolver for the expense field.
func (r *queryResolver) Expense(ctx context.Context, id string) (*model.Expense, error) {
	panic(fmt.Errorf("not implemented: Expense - expense"))
}

// Expense returns ExpenseResolver implementation.
func (r *Resolver) Expense() ExpenseResolver { return &expenseResolver{r} }

type expenseResolver struct{ *Resolver }
