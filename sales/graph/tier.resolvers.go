package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dailytravel/x/sales/graph/model"
	"github.com/dailytravel/x/sales/pkg/auth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateTier is the resolver for the createTier field.
func (r *mutationResolver) CreateTier(ctx context.Context, input model.NewTier) (*model.Tier, error) {
	// uid, err := utils.UID(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	item := &model.Tier{
		Locale:      input.Locale,
		Name:        bson.M{input.Locale: input.Name},
		Description: bson.M{input.Locale: input.Description},
		Model: model.Model{
			Metadata: input.Metadata,
		},
	}

	// Set the fields from the input
	_, err := r.db.Collection(item.Collection()).InsertOne(ctx, item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

// UpdateTier is the resolver for the updateTier field.
func (r *mutationResolver) UpdateTier(ctx context.Context, id string, input model.UpdateTier) (*model.Tier, error) {
	// uid, err := utils.UID(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	// Convert the ID string to ObjectID
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Create an update document with the fields to be updated
	item := &model.Tier{}
	filter := bson.M{"_id": _id}
	err = r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(item)
	if err != nil {
		return nil, err
	}

	if input.Name != nil {
		item.Name[item.Locale] = *input.Name
	}

	if input.Description != nil {
		item.Description[item.Locale] = *input.Description
	}

	if input.Metadata != nil {
		for k, v := range input.Metadata {
			item.Metadata[k] = v
		}
	}

	// Perform the update in the database
	res, err := r.db.Collection(item.Collection()).UpdateOne(ctx, filter, item)
	if err != nil {
		return nil, err
	}

	// Check if the coupon was actually updated
	if res.ModifiedCount == 0 {
		return nil, fmt.Errorf("no coupon was updated")
	}

	return item, nil
}

// DeleteTier is the resolver for the deleteTier field.
func (r *mutationResolver) DeleteTier(ctx context.Context, id string) (map[string]interface{}, error) {
	// uid, err := utils.UID(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	// Convert the ID string to ObjectID
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Define the filter to match the given ID
	filter := bson.M{"_id": _id}

	// Define the update to mark the record as deleted
	update := bson.M{
		"$set": bson.M{
			"deleted": primitive.Timestamp{T: uint32(time.Now().Unix())},
			"status":  "deleted",
			"updated": primitive.Timestamp{T: uint32(time.Now().Unix())},
		},
	}

	// Perform the update operation in the database
	result, err := r.db.Collection("tiers").UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	if result.ModifiedCount == 0 {
		return nil, fmt.Errorf("tier not found")
	}

	return map[string]interface{}{"status": "success", "deletedCount": result.ModifiedCount}, nil
}

// DeleteTiers is the resolver for the deleteTiers field.
func (r *mutationResolver) DeleteTiers(ctx context.Context, ids []string) (map[string]interface{}, error) {
	// uid, err := utils.UID(ctx)
	// if err != nil {
	// 	return nil, err
	// }

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

	// Define the update to mark records as deleted
	update := bson.M{
		"$set": bson.M{
			"deleted": primitive.Timestamp{T: uint32(time.Now().Unix())},
			"status":  "deleted",
			"updated": primitive.Timestamp{T: uint32(time.Now().Unix())},
		},
	}

	// Perform the update operation in the database
	result, err := r.db.Collection("tiers").UpdateMany(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"status": "success", "deletedCount": result.ModifiedCount}, nil
}

// Tiers is the resolver for the tiers field.
func (r *queryResolver) Tiers(ctx context.Context, stages map[string]interface{}) (*model.Tiers, error) {
	var items []*model.Tier
	//find all items
	cur, err := r.db.Collection("tiers").Find(ctx, nil)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var item *model.Tier
		if err := cur.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	//get total count
	count, err := r.db.Collection("tiers").CountDocuments(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &model.Tiers{
		Count: int(count),
		Data:  items,
	}, nil
}

// Tier is the resolver for the tier field.
func (r *queryResolver) Tier(ctx context.Context, id string) (*model.Tier, error) {
	var item *model.Tier

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

// ID is the resolver for the id field.
func (r *tierResolver) ID(ctx context.Context, obj *model.Tier) (string, error) {
	return obj.ID.Hex(), nil
}

// Name is the resolver for the name field.
func (r *tierResolver) Name(ctx context.Context, obj *model.Tier) (string, error) {
	// Get the locale from the context
	locale := auth.Locale(ctx)
	if locale == nil {
		locale = &obj.Locale
	}

	// Try to retrieve the name for the requested locale
	if name, ok := obj.Name[*locale].(string); ok {
		return name, nil
	}

	return obj.Name[obj.Locale].(string), nil
}

// Description is the resolver for the description field.
func (r *tierResolver) Description(ctx context.Context, obj *model.Tier) (string, error) {
	// Get the locale from the context
	locale := auth.Locale(ctx)
	if locale == nil {
		locale = &obj.Locale
	}

	// Try to retrieve the description for the requested locale
	if description, ok := obj.Description[*locale].(string); ok {
		return description, nil
	}

	return obj.Description[obj.Locale].(string), nil
}

// Benefits is the resolver for the benefits field.
func (r *tierResolver) Benefits(ctx context.Context, obj *model.Tier) ([]*model.Benefit, error) {
	var items []*model.Benefit

	filter := bson.M{"_id": bson.M{"$in": obj.Benefits}}

	cursor, err := r.db.Collection("benefits").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &items); err != nil {
		return nil, err
	}

	return items, nil
}

// Rewards is the resolver for the rewards field.
func (r *tierResolver) Rewards(ctx context.Context, obj *model.Tier) ([]*model.Reward, error) {
	var items []*model.Reward

	filter := bson.M{"_id": bson.M{"$in": obj.Benefits}}

	cursor, err := r.db.Collection("rewards").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &items); err != nil {
		return nil, err
	}

	return items, nil
}

// Metadata is the resolver for the metadata field.
func (r *tierResolver) Metadata(ctx context.Context, obj *model.Tier) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// Created is the resolver for the created field.
func (r *tierResolver) Created(ctx context.Context, obj *model.Tier) (string, error) {
	return time.Unix(int64(obj.Created.T), 0).Format(time.RFC3339), nil
}

// Updated is the resolver for the updated field.
func (r *tierResolver) Updated(ctx context.Context, obj *model.Tier) (string, error) {
	return time.Unix(int64(obj.Updated.T), 0).Format(time.RFC3339), nil
}

// Tier returns TierResolver implementation.
func (r *Resolver) Tier() TierResolver { return &tierResolver{r} }

type tierResolver struct{ *Resolver }
