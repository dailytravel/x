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
	"github.com/typesense/typesense-go/typesense/api/pointer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateLocation is the resolver for the createLocation field.
func (r *mutationResolver) CreateLocation(ctx context.Context, input model.NewLocation) (*model.Place, error) {
	item := &model.Place{
		Type:   input.Type,
		Locale: input.Locale,
		Name: map[string]interface{}{
			input.Locale: input.Name,
		},
		Model: model.Model{
			Metadata: input.Metadata,
		},
	}

	if input.Parent != nil {
		_parent, err := primitive.ObjectIDFromHex(*input.Parent)
		if err != nil {
			return nil, err
		}
		item.Parent = &_parent
	}

	if input.Description != nil {
		if item.Description == nil {
			item.Description = make(map[string]interface{})
		}

		item.Description[item.Locale] = input.Description
	}

	if input.Location != nil {
		item.Location = &model.Location{
			Lat: input.Location.Lat,
			Lng: input.Location.Lng,
		}
	}

	//insert the item into the database
	res, err := r.db.Collection(item.Collection()).InsertOne(ctx, item)
	if err != nil {
		return nil, fmt.Errorf("error creating log: %v", err)
	}

	item.ID = res.InsertedID.(primitive.ObjectID)

	return item, nil
}

// UpdateLocation is the resolver for the updateLocation field.
func (r *mutationResolver) UpdateLocation(ctx context.Context, id string, input model.UpdateLocation) (*model.Place, error) {
	var item *model.Place

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	//get item
	if err := r.db.Collection(item.Collection()).FindOne(ctx, bson.M{"_id": _id}, nil).Decode(&item); err != nil {
		return nil, err
	}

	//update item
	if input.Name != nil {
		item.Name[item.Locale] = *input.Name
	}

	if input.Description != nil {
		item.Description[item.Locale] = *input.Description
	}

	if input.Metadata != nil {
		if item.Metadata == nil {
			item.Metadata = make(map[string]interface{})
		}
		for key, value := range input.Metadata {
			item.Metadata[key] = value
		}
	}

	if input.Location != nil {
		item.Location = &model.Location{
			Lat: input.Location.Lat,
			Lng: input.Location.Lng,
		}
	}

	//update item
	if _, err := r.db.Collection(item.Collection()).UpdateOne(ctx, bson.M{"_id": _id}, bson.M{"$set": item}); err != nil {
		return nil, err
	}

	return item, nil
}

// DeleteLocation is the resolver for the deleteLocation field.
func (r *mutationResolver) DeleteLocation(ctx context.Context, id string) (*bool, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	res, err := r.db.Collection("places").DeleteOne(ctx, bson.M{"_id": _id})
	if err != nil {
		return nil, fmt.Errorf("error deleting log: %v", err)
	}

	if res.DeletedCount == 0 {
		return nil, fmt.Errorf("log not found")
	}

	return pointer.True(), nil
}

// DeleteLocations is the resolver for the deleteLocations field.
func (r *mutationResolver) DeleteLocations(ctx context.Context, ids []string) (*bool, error) {
	var _ids []primitive.ObjectID

	for _, id := range ids {
		_id, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		_ids = append(_ids, _id)
	}

	res, err := r.db.Collection("places").DeleteMany(ctx, bson.M{"_id": bson.M{"$in": _ids}})
	if err != nil {
		return nil, fmt.Errorf("error deleting log: %v", err)
	}

	if res.DeletedCount == 0 {
		return nil, fmt.Errorf("place not found")
	}

	return pointer.True(), nil
}

// ID is the resolver for the id field.
func (r *placeResolver) ID(ctx context.Context, obj *model.Place) (string, error) {
	return obj.ID.Hex(), nil
}

// Parent is the resolver for the parent field.
func (r *placeResolver) Parent(ctx context.Context, obj *model.Place) (*model.Place, error) {
	var item *model.Place

	filter := bson.M{"parent": obj.Parent}
	options := options.FindOne().SetProjection(bson.M{"_id": 1, "name": 1, "description": 1})

	err := r.db.Collection(item.Collection()).FindOne(ctx, filter, options).Decode(&item)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return item, nil
}

// Name is the resolver for the name field.
func (r *placeResolver) Name(ctx context.Context, obj *model.Place) (string, error) {
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
func (r *placeResolver) Description(ctx context.Context, obj *model.Place) (*string, error) {
	// Get the locale from the context
	locale := auth.Locale(ctx)
	if locale == nil {
		locale = &obj.Locale
	}

	// Try to retrieve the description for the requested locale
	description, ok := obj.Description[*locale].(string)
	if !ok {
		// If the description for the requested locale is not found, fall back to the default locale
		defaultDescription, defaultOK := obj.Description[obj.Locale].(string)
		if !defaultOK {
			return nil, nil
		}
		description = defaultDescription
	}

	return &description, nil
}

// Metadata is the resolver for the metadata field.
func (r *placeResolver) Metadata(ctx context.Context, obj *model.Place) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// Created is the resolver for the created field.
func (r *placeResolver) Created(ctx context.Context, obj *model.Place) (string, error) {
	return time.Unix(int64(obj.Created.T), 0).Format(time.RFC3339), nil
}

// Updated is the resolver for the updated field.
func (r *placeResolver) Updated(ctx context.Context, obj *model.Place) (string, error) {
	return time.Unix(int64(obj.Updated.T), 0).Format(time.RFC3339), nil
}

// Place is the resolver for the place field.
func (r *queryResolver) Place(ctx context.Context, id string) (*model.Place, error) {
	var item *model.Place
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

// Places is the resolver for the places field.
func (r *queryResolver) Places(ctx context.Context, stages map[string]interface{}) (*model.Places, error) {
	pipeline := bson.A{}

	// Add additional stages to the pipeline
	for key, value := range stages {
		stage := bson.D{{Key: key, Value: value}}
		pipeline = append(pipeline, stage)
	}

	cursor, err := r.db.Collection("places").Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var items []*model.Place

	if err := cursor.All(ctx, &items); err != nil {
		return nil, err
	}

	return &model.Places{
		Count: int(cursor.RemainingBatchLength()),
		Data:  items,
	}, nil
}

// Place returns PlaceResolver implementation.
func (r *Resolver) Place() PlaceResolver { return &placeResolver{r} }

type placeResolver struct{ *Resolver }
