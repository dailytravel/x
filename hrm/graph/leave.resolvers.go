package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"
	"time"

	"github.com/dailytravel/x/hrm/graph/model"
	"github.com/dailytravel/x/hrm/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ID is the resolver for the id field.
func (r *leaveResolver) ID(ctx context.Context, obj *model.Leave) (string, error) {
	return obj.ID.Hex(), nil
}

// Start is the resolver for the start field.
func (r *leaveResolver) Start(ctx context.Context, obj *model.Leave) (string, error) {
	panic(fmt.Errorf("not implemented: Start - start"))
}

// End is the resolver for the end field.
func (r *leaveResolver) End(ctx context.Context, obj *model.Leave) (*string, error) {
	panic(fmt.Errorf("not implemented: End - end"))
}

// Metadata is the resolver for the metadata field.
func (r *leaveResolver) Metadata(ctx context.Context, obj *model.Leave) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// Created is the resolver for the created field.
func (r *leaveResolver) Created(ctx context.Context, obj *model.Leave) (string, error) {
	return time.Unix(int64(obj.Created.T), 0).Format(time.RFC3339), nil
}

// Updated is the resolver for the updated field.
func (r *leaveResolver) Updated(ctx context.Context, obj *model.Leave) (string, error) {
	return time.Unix(int64(obj.Updated.T), 0).Format(time.RFC3339), nil
}

// UID is the resolver for the uid field.
func (r *leaveResolver) UID(ctx context.Context, obj *model.Leave) (string, error) {
	panic(fmt.Errorf("not implemented: UID - uid"))
}

// CreateLeave is the resolver for the createLeave field.
func (r *mutationResolver) CreateLeave(ctx context.Context, input model.NewLeave) (*model.Leave, error) {
	item := &model.Leave{
		Type:   input.Type,
		Reason: input.Reason,
		Status: *input.Status,
	}

	//convert string to primitive.DateTime
	Start, err := time.Parse(time.RFC3339, input.Start)
	if err != nil {
		return nil, err
	}

	item.Start = primitive.NewDateTimeFromTime(Start)

	//convert string to primitive.DateTime
	if input.End != nil {
		End, err := time.Parse(time.RFC3339, *input.End)
		if err != nil {
			return nil, err
		}
		item.End = primitive.NewDateTimeFromTime(End)
	}

	// Set the fields from the input
	_, err = r.db.Collection(item.Collection()).InsertOne(ctx, item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

// UpdateLeave is the resolver for the updateLeave field.
func (r *mutationResolver) UpdateLeave(ctx context.Context, id string, input model.UpdateLeave) (*model.Leave, error) {
	// Convert the ID string to ObjectID
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Find the leave by ID
	item := &model.Leave{}
	filter := bson.M{"_id": _id}
	err = r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(item)
	if err != nil {
		return nil, err
	}

	// Update the leave fields based on input
	if input.Type != nil {
		item.Type = *input.Type
	}

	if input.Reason != nil {
		item.Reason = *input.Reason
	}

	if input.Status != nil {
		item.Status = *input.Status
	}

	// Update the leave in the database
	_, err = r.db.Collection(item.Collection()).UpdateOne(ctx, filter, bson.M{"$set": item})
	if err != nil {
		return nil, err
	}

	return item, nil
}

// DeleteLeave is the resolver for the deleteLeave field.
func (r *mutationResolver) DeleteLeave(ctx context.Context, id string) (map[string]interface{}, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	// Convert the ID string to ObjectID
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Define the filter to match the leave by ID
	filter := bson.M{"_id": _id}

	// Set the fields to mark the leave as deleted
	update := bson.M{
		"$set": bson.M{
			"deleted_at": primitive.Timestamp{T: uint32(time.Now().Unix())},
			"deleted_by": uid,
			"status":     "deleted",
			"updated_by": uid,
			"updated_at": primitive.Timestamp{T: uint32(time.Now().Unix())},
		},
	}

	// Perform the update and retrieve the modified count
	opts := options.Update().SetUpsert(false)
	result, err := r.db.Collection("leaves").UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"success":      true,
		"deletedCount": result.ModifiedCount,
	}, nil
}

// DeleteLeaves is the resolver for the deleteLeaves field.
func (r *mutationResolver) DeleteLeaves(ctx context.Context, ids []string) (map[string]interface{}, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	// Convert the string IDs to ObjectIDs
	var objectIDs []primitive.ObjectID
	for _, id := range ids {
		_id, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		objectIDs = append(objectIDs, _id)
	}

	// Define the filter to match leaves by IDs
	filter := bson.M{"_id": bson.M{"$in": objectIDs}}

	// Set the fields to mark the leaves as deleted
	update := bson.M{
		"$set": bson.M{
			"deleted_at": primitive.Timestamp{T: uint32(time.Now().Unix())},
			"deleted_by": uid,
			"status":     "deleted",
			"updated_by": uid,
			"updated_at": primitive.Timestamp{T: uint32(time.Now().Unix())},
		},
	}

	// Perform the update and retrieve the modified count
	opts := options.Update().SetUpsert(false)
	result, err := r.db.Collection("leaves").UpdateMany(ctx, filter, update, opts)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"success":      true,
		"deletedCount": result.ModifiedCount,
	}, nil
}

// Leaves is the resolver for the leaves field.
func (r *queryResolver) Leaves(ctx context.Context, filter map[string]interface{}, project map[string]interface{}, sort map[string]interface{}, collation map[string]interface{}, limit *int, skip *int) (*model.Leaves, error) {
	var items []*model.Leave
	//find all items
	cur, err := r.db.Collection("leaves").Find(ctx, nil)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var item *model.Leave
		if err := cur.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	//get total count
	count, err := r.db.Collection("leaves").CountDocuments(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &model.Leaves{
		Count: int(count),
		Data:  items,
	}, nil
}

// Leave is the resolver for the leave field.
func (r *queryResolver) Leave(ctx context.Context, id string) (*model.Leave, error) {
	var item *model.Leave
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	if err := r.db.Collection(item.Collection()).FindOne(ctx, bson.M{"_id": _id}).Decode(&item); err != nil {
		return nil, err
	}

	return item, nil
}

// Leave returns LeaveResolver implementation.
func (r *Resolver) Leave() LeaveResolver { return &leaveResolver{r} }

type leaveResolver struct{ *Resolver }
