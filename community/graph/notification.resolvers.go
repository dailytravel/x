package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dailytravel/x/community/graph/model"
	"github.com/dailytravel/x/community/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UpdateNotification is the resolver for the updateNotification field.
func (r *mutationResolver) UpdateNotification(ctx context.Context, id string) (*model.Notification, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": _id}

	item := &model.Notification{
		Read: primitive.Timestamp{T: uint32(time.Now().Unix())},
	}

	err = r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(&item)
	if err != nil {
		return nil, err
	}

	// Update the contact
	if _, err := r.db.Collection(item.Collection()).UpdateOne(ctx, filter, bson.M{"$set": item}, nil); err != nil {
		return nil, err
	}

	return item, nil
}

// DeleteNotification is the resolver for the deleteNotification field.
func (r *mutationResolver) DeleteNotification(ctx context.Context, id string) (map[string]interface{}, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	// Convert the ID string to ObjectID
	notificationID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Define the filter to match the notification by ID and the user who is deleting it
	filter := bson.M{
		"_id": notificationID,
		"uid": uid,
	}

	// Perform the delete operation
	result, err := r.db.Collection("notifications").DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	if result.DeletedCount == 0 {
		return map[string]interface{}{"status": "error", "message": "Notification not found or not owned by the user"}, nil
	}

	return map[string]interface{}{"status": "success", "message": "Notification deleted successfully"}, nil
}

// DeleteNotifications is the resolver for the deleteNotifications field.
func (r *mutationResolver) DeleteNotifications(ctx context.Context, ids []string) (map[string]interface{}, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	// Convert the list of ID strings to ObjectIDs
	var notificationIDs []primitive.ObjectID
	for _, id := range ids {
		notificationID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		notificationIDs = append(notificationIDs, notificationID)
	}

	// Define the filter to match the notifications by IDs and the user who is deleting them
	filter := bson.M{
		"_id": bson.M{"$in": notificationIDs},
		"uid": uid,
	}

	// Perform the delete operation
	result, err := r.db.Collection("notifications").DeleteMany(ctx, filter)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"status": "success", "deletedCount": result.DeletedCount}, nil
}

// ID is the resolver for the id field.
func (r *notificationResolver) ID(ctx context.Context, obj *model.Notification) (string, error) {
	return obj.ID.Hex(), nil
}

// UID is the resolver for the uid field.
func (r *notificationResolver) UID(ctx context.Context, obj *model.Notification) (string, error) {
	return obj.ID.Hex(), nil
}

// Read is the resolver for the read field.
func (r *notificationResolver) Read(ctx context.Context, obj *model.Notification) (*string, error) {
	panic(fmt.Errorf("not implemented: Read - read"))
}

// Notifiable is the resolver for the notifiable field.
func (r *notificationResolver) Notifiable(ctx context.Context, obj *model.Notification) (map[string]interface{}, error) {
	var item map[string]interface{}

	err := r.db.Collection(obj.Notifiable.Type).FindOne(ctx, bson.M{"_id": obj.Notifiable.ID}).Decode(&item)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return item, nil
}

// Metadata is the resolver for the metadata field.
func (r *notificationResolver) Metadata(ctx context.Context, obj *model.Notification) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// Created is the resolver for the created field.
func (r *notificationResolver) Created(ctx context.Context, obj *model.Notification) (string, error) {
	return time.Unix(int64(obj.Created.T), 0).Format(time.RFC3339), nil
}

// Updated is the resolver for the updated field.
func (r *notificationResolver) Updated(ctx context.Context, obj *model.Notification) (string, error) {
	return time.Unix(int64(obj.Updated.T), 0).Format(time.RFC3339), nil
}

// Notifications is the resolver for the notifications field.
func (r *queryResolver) Notifications(ctx context.Context, stages map[string]interface{}) (*model.Notifications, error) {
	var items []*model.Notification
	//find all items
	cur, err := r.db.Collection("notifications").Find(ctx, nil)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var item *model.Notification
		if err := cur.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	//get total count
	count, err := r.db.Collection("notifications").CountDocuments(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &model.Notifications{
		Count: int(count),
		Data:  items,
	}, nil
}

// Notification is the resolver for the notification field.
func (r *queryResolver) Notification(ctx context.Context, id string) (*model.Notification, error) {
	var item *model.Notification
	col := r.db.Collection(item.Collection())

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": _id}

	err = col.FindOne(ctx, filter).Decode(&item)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("no document found for filter %v", filter)
		}
		return nil, err
	}

	return item, nil
}

// Notification returns NotificationResolver implementation.
func (r *Resolver) Notification() NotificationResolver { return &notificationResolver{r} }

type notificationResolver struct{ *Resolver }
