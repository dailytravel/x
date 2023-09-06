package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dailytravel/x/base/graph/model"
	"github.com/dailytravel/x/base/internal/utils"
	"github.com/typesense/typesense-go/typesense/api/pointer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreatePortfolio is the resolver for the createPortfolio field.
func (r *mutationResolver) CreatePortfolio(ctx context.Context, input model.NewPortfolio) (*model.Portfolio, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	item := &model.Portfolio{
		UID: *uid,
		Model: model.Model{
			Metadata:  input.Metadata,
			CreatedBy: uid,
			UpdatedBy: uid,
		},
	}

	res, err := r.db.Collection(item.Collection()).InsertOne(ctx, item)
	if err != nil {
		return nil, err
	}

	item.ID = res.InsertedID.(primitive.ObjectID)

	return item, nil
}

// UpdatePortfolio is the resolver for the updatePortfolio field.
func (r *mutationResolver) UpdatePortfolio(ctx context.Context, id string, input model.UpdatePortfolio) (*model.Portfolio, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Fetch the existing board
	item := &model.Portfolio{}
	filter := bson.M{"_id": _id}
	err = r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(item)
	if err != nil {
		return nil, err
	}

	// Update fields based on input

	item.UpdatedBy = uid

	// Perform the update in the database
	update := bson.M{
		"$set": item,
	}
	_, err = r.db.Collection(item.Collection()).UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return item, nil
}

// DeletePortfolio is the resolver for the deletePortfolio field.
func (r *mutationResolver) DeletePortfolio(ctx context.Context, id string) (map[string]interface{}, error) {
	// Convert the ID string to ObjectID
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Create a filter based on the ID
	filter := bson.M{"_id": _id}

	// Perform the delete operation
	result, err := r.db.Collection("portfolios").DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	// Create and return the response
	response := map[string]interface{}{
		"status":       "success",
		"deletedCount": result.DeletedCount,
	}
	return response, nil
}

// DeletePortfolios is the resolver for the deletePortfolios field.
func (r *mutationResolver) DeletePortfolios(ctx context.Context, ids []string) (map[string]interface{}, error) {
	// Convert the list of ID strings to ObjectIDs
	var objectIDs []primitive.ObjectID
	for _, id := range ids {
		_id, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		objectIDs = append(objectIDs, _id)
	}

	// Define the filter to match the given IDs
	filter := bson.M{"_id": bson.M{"$in": objectIDs}}

	// Perform the delete operation
	result, err := r.db.Collection("portfolios").DeleteMany(ctx, filter)
	if err != nil {
		return nil, err
	}

	// Create and return the response
	response := map[string]interface{}{
		"status":       "success",
		"deletedCount": result.DeletedCount,
	}
	return response, nil
}

// ID is the resolver for the id field.
func (r *portfolioResolver) ID(ctx context.Context, obj *model.Portfolio) (string, error) {
	return obj.ID.Hex(), nil
}

// Boards is the resolver for the boards field.
func (r *portfolioResolver) Boards(ctx context.Context, obj *model.Portfolio) ([]*model.Board, error) {
	var items []*model.Board

	// Create a cursor for the query
	cursor, err := r.db.Collection("boards").Find(ctx, bson.M{"portfolio": obj.ID.Hex()})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Iterate over the cursor and decode documents
	for cursor.Next(ctx) {
		var item model.Board
		if err := cursor.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, &item)
	}

	// Check for cursor errors
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

// Metadata is the resolver for the metadata field.
func (r *portfolioResolver) Metadata(ctx context.Context, obj *model.Portfolio) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// CreatedAt is the resolver for the created_at field.
func (r *portfolioResolver) CreatedAt(ctx context.Context, obj *model.Portfolio) (string, error) {
	return time.Unix(int64(obj.CreatedAt.T), 0).Format(time.RFC3339), nil
}

// UpdatedAt is the resolver for the updated_at field.
func (r *portfolioResolver) UpdatedAt(ctx context.Context, obj *model.Portfolio) (string, error) {
	return time.Unix(int64(obj.UpdatedAt.T), 0).Format(time.RFC3339), nil
}

// UID is the resolver for the uid field.
func (r *portfolioResolver) UID(ctx context.Context, obj *model.Portfolio) (string, error) {
	return obj.UID.Hex(), nil
}

// CreatedBy is the resolver for the created_by field.
func (r *portfolioResolver) CreatedBy(ctx context.Context, obj *model.Portfolio) (*string, error) {
	if obj.CreatedBy == nil {
		return nil, nil
	}

	return pointer.String(obj.CreatedBy.Hex()), nil
}

// UpdatedBy is the resolver for the updated_by field.
func (r *portfolioResolver) UpdatedBy(ctx context.Context, obj *model.Portfolio) (*string, error) {
	if obj.UpdatedBy == nil {
		return nil, nil
	}

	return pointer.String(obj.UpdatedBy.Hex()), nil
}

// Portfolio is the resolver for the portfolio field.
func (r *queryResolver) Portfolio(ctx context.Context, id string) (*model.Portfolio, error) {
	var item *model.Portfolio

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": _id}

	err = r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(item)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("no document found for filter %v", filter)
		}
		return nil, err
	}

	return item, nil
}

// Portfolios is the resolver for the portfolios field.
func (r *queryResolver) Portfolios(ctx context.Context, args map[string]interface{}) (*model.Portfolios, error) {
	var items []*model.Portfolio

	opts := utils.Options(args)
	opts.SetSort(bson.M{"order": 1})
	opts.SetSort(bson.M{"created_at": -1})

	// Build the filter based on the provided arguments
	filter := bson.M{}

	// Add filters based on the arguments, if provided
	if name, ok := args["name"].(string); ok && name != "" {
		filter["name"] = name
	}

	// Create a cursor for the query
	cursor, err := r.db.Collection("portfolios").Find(ctx, utils.Query(args), opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Iterate over the cursor and decode documents
	for cursor.Next(ctx) {
		var item model.Portfolio
		if err := cursor.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, &item)
	}

	// Check for cursor errors
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// You can get the total count using CountDocuments method
	count, err := r.db.Collection("portfolios").CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &model.Portfolios{Data: items, Count: int(count)}, nil
}

// Portfolio returns PortfolioResolver implementation.
func (r *Resolver) Portfolio() PortfolioResolver { return &portfolioResolver{r} }

type portfolioResolver struct{ *Resolver }
