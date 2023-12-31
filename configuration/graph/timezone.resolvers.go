package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dailytravel/x/configuration/graph/model"
	"github.com/dailytravel/x/configuration/pkg/auth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateTimezone is the resolver for the createTimezone field.
func (r *mutationResolver) CreateTimezone(ctx context.Context, input model.NewTimezone) (*model.Timezone, error) {
	panic(fmt.Errorf("not implemented: CreateTimezone - createTimezone"))
}

// UpdateTimezone is the resolver for the updateTimezone field.
func (r *mutationResolver) UpdateTimezone(ctx context.Context, id string, input model.UpdateTimezone) (*model.Timezone, error) {
	panic(fmt.Errorf("not implemented: UpdateTimezone - updateTimezone"))
}

// ImportTimezones is the resolver for the importTimezones field.
func (r *mutationResolver) ImportTimezones(ctx context.Context, file string) ([]*model.Timezone, error) {
	panic(fmt.Errorf("not implemented: ImportTimezones - importTimezones"))
}

// DeleteTimezone is the resolver for the deleteTimezone field.
func (r *mutationResolver) DeleteTimezone(ctx context.Context, id string) (map[string]interface{}, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	res, err := r.db.Collection("timezones").DeleteOne(ctx, bson.M{"_id": _id})
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

// DeleteTimezones is the resolver for the deleteTimezones field.
func (r *mutationResolver) DeleteTimezones(ctx context.Context, ids []string) (map[string]interface{}, error) {
	var _ids []primitive.ObjectID

	for _, id := range ids {
		_id, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		_ids = append(_ids, _id)
	}

	res, err := r.db.Collection("timezones").DeleteMany(ctx, bson.M{"_id": bson.M{"$in": _ids}})
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

// Timezones is the resolver for the timezones field.
func (r *queryResolver) Timezones(ctx context.Context, stages map[string]interface{}) (*model.Timezones, error) {
	pipeline := bson.A{}

	// Add additional stages to the pipeline
	for key, value := range stages {
		stage := bson.D{{Key: key, Value: value}}
		pipeline = append(pipeline, stage)
	}

	cursor, err := r.db.Collection("timezones").Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var items []*model.Timezone

	if err := cursor.All(ctx, &items); err != nil {
		return nil, err
	}

	return &model.Timezones{
		Count: int(cursor.RemainingBatchLength()),
		Data:  items,
	}, nil
}

// Timezone is the resolver for the timezone field.
func (r *queryResolver) Timezone(ctx context.Context, id string) (*model.Timezone, error) {
	var item *model.Timezone
	col := r.db.Collection(item.Collection())

	filter := bson.M{"_id": id}

	err := col.FindOne(ctx, filter).Decode(&item)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("no document found for filter %v", filter)
		}
		return nil, err
	}

	return item, nil
}

// ID is the resolver for the id field.
func (r *timezoneResolver) ID(ctx context.Context, obj *model.Timezone) (string, error) {
	return obj.ID.Hex(), nil
}

// Name is the resolver for the name field.
func (r *timezoneResolver) Name(ctx context.Context, obj *model.Timezone) (string, error) {
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

// Metadata is the resolver for the metadata field.
func (r *timezoneResolver) Metadata(ctx context.Context, obj *model.Timezone) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// Created is the resolver for the created field.
func (r *timezoneResolver) Created(ctx context.Context, obj *model.Timezone) (string, error) {
	return time.Unix(int64(obj.Created.T), 0).Format(time.RFC3339), nil
}

// Updated is the resolver for the updated field.
func (r *timezoneResolver) Updated(ctx context.Context, obj *model.Timezone) (string, error) {
	return time.Unix(int64(obj.Updated.T), 0).Format(time.RFC3339), nil
}

// Timezone returns TimezoneResolver implementation.
func (r *Resolver) Timezone() TimezoneResolver { return &timezoneResolver{r} }

type timezoneResolver struct{ *Resolver }
