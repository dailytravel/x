package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"
	"time"

	"github.com/dailytravel/x/community/graph/model"
	"github.com/dailytravel/x/community/internal/utils"
	"github.com/typesense/typesense-go/typesense/api/pointer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateShare is the resolver for the createShare field.
func (r *mutationResolver) CreateShare(ctx context.Context, input model.ShareInput) (*model.Share, error) {
	panic(fmt.Errorf("not implemented: CreateShare - createShare"))
}

// UpdateShare is the resolver for the updateShare field.
func (r *mutationResolver) UpdateShare(ctx context.Context, id string, input model.ShareUpdateInput) (*model.Share, error) {
	panic(fmt.Errorf("not implemented: UpdateShare - updateShare"))
}

// DeleteShare is the resolver for the deleteShare field.
func (r *mutationResolver) DeleteShare(ctx context.Context, id string) (bool, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}

	// Delete the share from the MongoDB collection
	result, err := r.db.Collection("shares").DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return false, err
	}

	// Check if any document was deleted
	if result.DeletedCount == 0 {
		return false, nil
	}

	return true, nil
}

// DeleteShares is the resolver for the deleteShares field.
func (r *mutationResolver) DeleteShares(ctx context.Context, ids []string) (bool, error) {
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

	// Delete the shares from the MongoDB collection
	result, err := r.db.Collection("shares").DeleteMany(ctx, filter)
	if err != nil {
		return false, err
	}

	// Check if any documents were deleted
	if result.DeletedCount == 0 {
		return false, nil
	}

	return true, nil
}

// Share is the resolver for the share field.
func (r *queryResolver) Share(ctx context.Context, id string) (*model.Share, error) {
	// Parse string ID to MongoDB ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %v", err)
	}

	// Query the MongoDB collection
	var share *model.Share
	err = r.db.Collection("shares").FindOne(ctx, bson.M{"_id": objID}).Decode(&share)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("Share with ID %s not found", id)
		}
		return nil, err
	}

	return share, nil
}

// Shares is the resolver for the shares field.
func (r *queryResolver) Shares(ctx context.Context, args map[string]interface{}) (*model.Shares, error) {
	var items []*model.Share

	// Ensure type and id exist in args map
	shareableType, typeOk := args["type"].(string)
	shareableID, idOk := args["id"].(primitive.ObjectID)

	if !typeOk || !idOk {
		return nil, fmt.Errorf("missing or invalid type or id arguments")
	}

	// Filter by shareable
	filter := bson.M{
		"shareable": bson.M{
			"type": shareableType,
			"_id":  shareableID,
		},
	}

	// Query the MongoDB collection
	cursor, err := r.db.Collection("shares").Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Decode cursor into items slice
	err = cursor.All(ctx, &items)
	if err != nil {
		return nil, err
	}

	// Count
	count, err := r.db.Collection("shares").CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &model.Shares{Data: items, Count: int(count)}, nil
}

// ID is the resolver for the id field.
func (r *shareResolver) ID(ctx context.Context, obj *model.Share) (string, error) {
	return obj.ID.Hex(), nil
}

// UID is the resolver for the uid field.
func (r *shareResolver) UID(ctx context.Context, obj *model.Share) (string, error) {
	return obj.UID.Hex(), nil
}

// Shareable is the resolver for the shareable field.
func (r *shareResolver) Shareable(ctx context.Context, obj *model.Share) (map[string]interface{}, error) {
	results, err := utils.StructToMap(obj.Shareable)
	if err != nil {
		return nil, err
	}

	return results, nil
}

// Metadata is the resolver for the metadata field.
func (r *shareResolver) Metadata(ctx context.Context, obj *model.Share) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// CreatedAt is the resolver for the created_at field.
func (r *shareResolver) CreatedAt(ctx context.Context, obj *model.Share) (string, error) {
	return time.Unix(int64(obj.CreatedAt.T), 0).Format(time.RFC3339), nil
}

// UpdatedAt is the resolver for the updated_at field.
func (r *shareResolver) UpdatedAt(ctx context.Context, obj *model.Share) (string, error) {
	return time.Unix(int64(obj.UpdatedAt.T), 0).Format(time.RFC3339), nil
}

// CreatedBy is the resolver for the created_by field.
func (r *shareResolver) CreatedBy(ctx context.Context, obj *model.Share) (*string, error) {
	if obj.CreatedBy == nil {
		return nil, nil
	}

	return pointer.String(obj.CreatedBy.Hex()), nil
}

// UpdatedBy is the resolver for the updated_by field.
func (r *shareResolver) UpdatedBy(ctx context.Context, obj *model.Share) (*string, error) {
	if obj.UpdatedBy == nil {
		return nil, nil
	}

	return pointer.String(obj.UpdatedBy.Hex()), nil
}

// Share returns ShareResolver implementation.
func (r *Resolver) Share() ShareResolver { return &shareResolver{r} }

type shareResolver struct{ *Resolver }
