package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dailytravel/x/payment/graph/model"
	"github.com/dailytravel/x/payment/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ID is the resolver for the id field.
func (r *cardResolver) ID(ctx context.Context, obj *model.Card) (string, error) {
	return obj.ID.Hex(), nil
}

// Billing is the resolver for the billing field.
func (r *cardResolver) Billing(ctx context.Context, obj *model.Card) (map[string]interface{}, error) {
	return obj.Billing, nil
}

// Metadata is the resolver for the metadata field.
func (r *cardResolver) Metadata(ctx context.Context, obj *model.Card) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// Created is the resolver for the created field.
func (r *cardResolver) Created(ctx context.Context, obj *model.Card) (string, error) {
	return time.Unix(int64(obj.Created.T), 0).Format(time.RFC3339), nil
}

// Updated is the resolver for the updated field.
func (r *cardResolver) Updated(ctx context.Context, obj *model.Card) (string, error) {
	return time.Unix(int64(obj.Updated.T), 0).Format(time.RFC3339), nil
}

// UID is the resolver for the uid field.
func (r *cardResolver) UID(ctx context.Context, obj *model.Card) (string, error) {
	return obj.UID.Hex(), nil
}

// Wallet is the resolver for the wallet field.
func (r *cardResolver) Wallet(ctx context.Context, obj *model.Card) (*model.Wallet, error) {
	var item *model.Wallet
	if err := r.db.Collection("wallets").FindOne(ctx, bson.M{"wallet": obj.Wallet}).Decode(&item); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return item, nil
}

// CreateCard is the resolver for the createCard field.
func (r *mutationResolver) CreateCard(ctx context.Context, input model.NewCardInput) (*model.Card, error) {
	key := []byte(os.Getenv("ENCRYPTION_KEY"))
	if len(key) != 32 { // AES-256 requires a 32-byte key
		return nil, fmt.Errorf("invalid encryption key length")
	}

	number, err := utils.Encrypt([]byte(input.Number), key)
	if err != nil {
		return nil, err
	}

	// Get $1 to test card
	// https://stripe.com/docs/testing#cards

	item := &model.Card{
		Name:     input.Name,
		Number:   string(number),
		ExpMonth: input.ExpMonth,
		ExpYear:  input.ExpYear,
		Cvv:      input.Cvv,
	}

	item.Billing = make(map[string]interface{})
	for k, v := range input.Billing {
		item.Billing[k] = v
	}

	item.Metadata = make(map[string]interface{})
	for k, v := range input.Metadata {
		item.Metadata[k] = v
	}

	// It might be risky to search by plaintext number, even if you're searching encrypted values.
	// Consider using other unique identifiers instead of the actual number.
	if err := r.db.Collection("cards").FindOne(ctx, bson.M{"number": input.Number}).Decode(&item); err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err
		}
	}

	return item, nil
}

// UpdateCard is the resolver for the updateCard field.
func (r *mutationResolver) UpdateCard(ctx context.Context, id string, input model.UpdateCardInput) (*model.Card, error) {
	// Convert the string ID to an ObjectID.
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %v", err)
	}

	// First, find the item with the given ID.
	var item *model.Card
	if err := r.db.Collection("cards").FindOne(ctx, bson.M{"_id": objectID}).Decode(&item); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("card with ID %s not found", id)
		}
		return nil, fmt.Errorf("error finding card: %v", err)
	}

	// Create an update command based on the input.
	update := bson.M{}

	if input.Name != nil {
		update["name"] = *input.Name
	}

	if input.ExpMonth != nil {
		update["exp_month"] = *input.ExpMonth
	}

	if input.ExpYear != nil {
		update["exp_year"] = *input.ExpYear
	}

	if input.Cvv != nil {
		update["cvv"] = *input.Cvv
	}

	if input.Billing != nil {
		update["billing"] = input.Billing
	}

	// Get $1 to test card
	// https://stripe.com/docs/testing#cards

	// Update the card in the database.
	_, err = r.db.Collection("cards").UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		return nil, fmt.Errorf("error updating card: %v", err)
	}

	return item, nil
}

// DeleteCard is the resolver for the deleteCard field.
func (r *mutationResolver) DeleteCard(ctx context.Context, id string) (*model.Card, error) {
	// Convert the string ID to an ObjectID.
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %v", err)
	}

	// First, find the item with the given ID.
	var item *model.Card
	err = r.db.Collection("cards").FindOne(ctx, bson.M{"_id": objectID}).Decode(&item)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("card with ID %s not found", id)
		}
		return nil, fmt.Errorf("error finding card: %v", err)
	}

	// Now, delete the card.
	_, err = r.db.Collection("cards").DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return nil, fmt.Errorf("error deleting card: %v", err)
	}

	return item, nil
}

// Card is the resolver for the card field.
func (r *queryResolver) Card(ctx context.Context, id string) (*model.Card, error) {
	var item *model.Card

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	if err := r.db.Collection("cards").FindOne(ctx, bson.M{"_id": _id}).Decode(&item); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return item, nil
}

// Cards is the resolver for the cards field.
func (r *queryResolver) Cards(ctx context.Context, stages map[string]interface{}) (*model.Cards, error) {
	pipeline := bson.A{}

	// Add additional stages to the pipeline
	for key, value := range stages {
		stage := bson.D{{Key: key, Value: value}}
		pipeline = append(pipeline, stage)
	}

	cursor, err := r.db.Collection("cards").Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var items []*model.Card

	if err := cursor.All(ctx, &items); err != nil {
		return nil, err
	}

	return &model.Cards{
		Count: int(cursor.RemainingBatchLength()),
		Data:  items,
	}, nil
}

// Card returns CardResolver implementation.
func (r *Resolver) Card() CardResolver { return &cardResolver{r} }

type cardResolver struct{ *Resolver }
