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

// ID is the resolver for the id field.
func (r *currencyResolver) ID(ctx context.Context, obj *model.Currency) (string, error) {
	return obj.ID.Hex(), nil
}

// Name is the resolver for the name field.
func (r *currencyResolver) Name(ctx context.Context, obj *model.Currency) (string, error) {
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
func (r *currencyResolver) Metadata(ctx context.Context, obj *model.Currency) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// Created is the resolver for the created field.
func (r *currencyResolver) Created(ctx context.Context, obj *model.Currency) (string, error) {
	return time.Unix(int64(obj.Created.T), 0).Format(time.RFC3339), nil
}

// Updated is the resolver for the updated field.
func (r *currencyResolver) Updated(ctx context.Context, obj *model.Currency) (string, error) {
	return time.Unix(int64(obj.Updated.T), 0).Format(time.RFC3339), nil
}

// CreateCurrency is the resolver for the createCurrency field.
func (r *mutationResolver) CreateCurrency(ctx context.Context, input model.NewCurrency) (*model.Currency, error) {
	var item *model.Currency

	doc := &model.Currency{
		Code: input.Code,
		Name: map[string]interface{}{
			input.Locale: input.Name,
		},
	}

	//insert item
	if _, err := r.db.Collection(item.Collection()).InsertOne(ctx, doc, nil); err != nil {
		return nil, err
	}

	return item, nil
}

// UpdateCurrency is the resolver for the updateCurrency field.
func (r *mutationResolver) UpdateCurrency(ctx context.Context, id string, input *model.UpdateCurrency) (*model.Currency, error) {
	var item *model.Currency

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	//get item
	if err := r.db.Collection(item.Collection()).FindOne(ctx, bson.M{"_id": _id}, nil).Decode(&item); err != nil {
		return nil, err
	}

	//update item
	if _, err := r.db.Collection(item.Collection()).UpdateOne(ctx, bson.M{"_id": _id}, bson.M{"$set": input}); err != nil {
		return nil, err
	}

	return item, nil
}

// DeleteCurrency is the resolver for the deleteCurrency field.
func (r *mutationResolver) DeleteCurrency(ctx context.Context, id string) (map[string]interface{}, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	res, err := r.db.Collection("currencies").DeleteOne(ctx, bson.M{"_id": _id})
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

// DeleteCurrencies is the resolver for the deleteCurrencies field.
func (r *mutationResolver) DeleteCurrencies(ctx context.Context, ids []string) (map[string]interface{}, error) {
	var _ids []primitive.ObjectID

	for _, id := range ids {
		_id, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		_ids = append(_ids, _id)
	}

	res, err := r.db.Collection("currencies").DeleteMany(ctx, bson.M{"_id": bson.M{"$in": _ids}})
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

// Currency is the resolver for the currency field.
func (r *queryResolver) Currency(ctx context.Context, code string) (*model.Currency, error) {
	var item *model.Currency
	col := r.db.Collection(item.Collection())

	filter := bson.M{"code": code}

	err := col.FindOne(ctx, filter).Decode(&item)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("no document found for filter %v", filter)
		}
		return nil, err
	}

	return item, nil
}

// Currencies is the resolver for the currencies field.
func (r *queryResolver) Currencies(ctx context.Context, stages map[string]interface{}) (*model.Currencies, error) {
	pipeline := bson.A{}

	// Add additional stages to the pipeline
	for key, value := range stages {
		stage := bson.D{{Key: key, Value: value}}
		pipeline = append(pipeline, stage)
	}

	cursor, err := r.db.Collection("currencies").Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var items []*model.Currency

	if err := cursor.All(ctx, &items); err != nil {
		return nil, err
	}

	return &model.Currencies{
		Count: int(cursor.RemainingBatchLength()),
		Data:  items,
	}, nil
}

// Currency returns CurrencyResolver implementation.
func (r *Resolver) Currency() CurrencyResolver { return &currencyResolver{r} }

type currencyResolver struct{ *Resolver }
