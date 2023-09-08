package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"errors"
	"fmt"

	"github.com/dailytravel/x/insight/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ID is the resolver for the id field.
func (r *logResolver) ID(ctx context.Context, obj *model.Log) (string, error) {
	return obj.ID.Hex(), nil
}

// UID is the resolver for the uid field.
func (r *logResolver) UID(ctx context.Context, obj *model.Log) (*string, error) {
	if obj.UID == nil {
		return nil, nil
	}

	uid := obj.UID.Hex()
	return &uid, nil
}

// Utm is the resolver for the utm field.
func (r *logResolver) Utm(ctx context.Context, obj *model.Log) (map[string]interface{}, error) {
	return obj.Utm, nil
}

// Metadata is the resolver for the metadata field.
func (r *logResolver) Metadata(ctx context.Context, obj *model.Log) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// Timestamp is the resolver for the timestamp field.
func (r *logResolver) Timestamp(ctx context.Context, obj *model.Log) (string, error) {
	panic(fmt.Errorf("not implemented: Timestamp - timestamp"))
}

// CreateLog is the resolver for the createLog field.
func (r *mutationResolver) CreateLog(ctx context.Context, input model.NewLog) (*model.Log, error) {
	panic(fmt.Errorf("not implemented: CreateLog - createLog"))
}

// UpdateLog is the resolver for the updateLog field.
func (r *mutationResolver) UpdateLog(ctx context.Context, id string, input model.UpdateLog) (*model.Log, error) {
	panic(fmt.Errorf("not implemented: UpdateLog - updateLog"))
}

// DeleteLog is the resolver for the deleteLog field.
func (r *mutationResolver) DeleteLog(ctx context.Context, id string) (map[string]interface{}, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	res, err := r.db.Collection("logs").DeleteOne(ctx, bson.M{"_id": _id})
	if err != nil {
		return nil, fmt.Errorf("error deleting log: %v", err)
	}

	if res.DeletedCount == 0 {
		return nil, fmt.Errorf("log not found")
	}

	return map[string]interface{}{
		"status": "success",
	}, nil
}

// DeleteLogs is the resolver for the deleteLogs field.
func (r *mutationResolver) DeleteLogs(ctx context.Context, ids []*string) (map[string]interface{}, error) {
	var _ids []primitive.ObjectID

	for _, id := range ids {
		_id, err := primitive.ObjectIDFromHex(*id)
		if err != nil {
			return nil, err
		}
		_ids = append(_ids, _id)
	}

	res, err := r.db.Collection("logs").DeleteMany(ctx, bson.M{"_id": bson.M{"$in": _ids}})
	if err != nil {
		return nil, fmt.Errorf("error deleting log: %v", err)
	}

	if res.DeletedCount == 0 {
		return nil, fmt.Errorf("log not found")
	}

	return map[string]interface{}{
		"status": "success",
	}, nil
}

// Log is the resolver for the log field.
func (r *queryResolver) Log(ctx context.Context, id string) (*model.Log, error) {
	var item *model.Log
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": _id}

	if err := r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(&item); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("no document found for filter %v", filter)
		}
		return nil, err
	}

	return item, nil
}

// Logs is the resolver for the logs field.
func (r *queryResolver) Logs(ctx context.Context, args map[string]interface{}) (*model.Logs, error) {
	var items []*model.Log
	//find all items
	cur, err := r.db.Collection("logs").Find(ctx, nil)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var item *model.Log
		if err := cur.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	//get total count
	count, err := r.db.Collection("locales").CountDocuments(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &model.Logs{
		Count: int(count),
		Data:  items,
	}, nil
}

// Log returns LogResolver implementation.
func (r *Resolver) Log() LogResolver { return &logResolver{r} }

type logResolver struct{ *Resolver }
