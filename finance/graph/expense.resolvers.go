package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dailytravel/x/finance/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ID is the resolver for the id field.
func (r *expenseResolver) ID(ctx context.Context, obj *model.Expense) (string, error) {
	return obj.ID.Hex(), nil
}

// Metadata is the resolver for the metadata field.
func (r *expenseResolver) Metadata(ctx context.Context, obj *model.Expense) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// Date is the resolver for the date field.
func (r *expenseResolver) Date(ctx context.Context, obj *model.Expense) (string, error) {
	return obj.Date.Time().Format(time.RFC3339), nil
}

// UID is the resolver for the uid field.
func (r *expenseResolver) UID(ctx context.Context, obj *model.Expense) (string, error) {
	return obj.ID.Hex(), nil
}

// Created is the resolver for the created field.
func (r *expenseResolver) Created(ctx context.Context, obj *model.Expense) (string, error) {
	return time.Unix(int64(obj.Created.T), 0).Format(time.RFC3339), nil
}

// Updated is the resolver for the updated field.
func (r *expenseResolver) Updated(ctx context.Context, obj *model.Expense) (string, error) {
	return time.Unix(int64(obj.Updated.T), 0).Format(time.RFC3339), nil
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
func (r *mutationResolver) DeleteExpense(ctx context.Context, id string) (bool, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}

	// Delete the share from the MongoDB collection
	result, err := r.db.Collection("expenses").DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return false, err
	}

	// Check if any document was deleted
	if result.DeletedCount == 0 {
		return false, nil
	}

	return true, nil
}

// DeleteExpenses is the resolver for the deleteExpenses field.
func (r *mutationResolver) DeleteExpenses(ctx context.Context, ids []string) (bool, error) {
	// Convert string IDs to ObjectIDs
	objIDs := make([]primitive.ObjectID, 0, len(ids))
	for _, id := range ids {
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return false, err
		}
		objIDs = append(objIDs, objID)
	}

	// Build the filter to match multiple IDs
	filter := bson.M{"_id": bson.M{"$in": objIDs}}

	// Delete the expenses from the MongoDB collection
	result, err := r.db.Collection("expenses").DeleteMany(ctx, filter)
	if err != nil {
		return false, err
	}

	// Check if any documents were deleted
	if result.DeletedCount == 0 {
		return false, nil
	}

	return true, nil
}

// Expenses is the resolver for the expenses field.
func (r *queryResolver) Expenses(ctx context.Context, stages map[string]interface{}) (*model.Expenses, error) {
	var items []*model.Expense
	//find all items
	cur, err := r.db.Collection("expenses").Find(ctx, nil)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var item *model.Expense
		if err := cur.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	//get total count
	count, err := r.db.Collection("expenses").CountDocuments(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &model.Expenses{
		Count: int(count),
		Data:  items,
	}, nil
}

// Expense is the resolver for the expense field.
func (r *queryResolver) Expense(ctx context.Context, id string) (*model.Expense, error) {
	var item *model.Expense

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": _id}
	if err := r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(&item); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("document not found")
		}
		return nil, err
	}

	return item, nil
}

// Expense returns ExpenseResolver implementation.
func (r *Resolver) Expense() ExpenseResolver { return &expenseResolver{r} }

type expenseResolver struct{ *Resolver }
