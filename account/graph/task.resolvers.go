package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"

	"github.com/dailytravel/x/account/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// User is the resolver for the user field.
func (r *taskResolver) User(ctx context.Context, obj *model.Task) (*model.User, error) {
	_id, err := primitive.ObjectIDFromHex(obj.UID)
	if err != nil {
		return nil, fmt.Errorf("failed to convert UID to ObjectID: %w", err)
	}

	var item *model.User

	filter := bson.M{"_id": _id}
	if err := r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(&item); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // Return nil if no user is found, rather than an error.
		}
		return nil, fmt.Errorf("failed to fetch user from database: %w", err)
	}

	return item, nil
}

// Lead is the resolver for the lead field.
func (r *taskResolver) Lead(ctx context.Context, obj *model.Task) (*model.User, error) {
	_id, err := primitive.ObjectIDFromHex(obj.Assignee)
	if err != nil {
		return nil, nil
	}

	var item *model.User

	filter := bson.M{"_id": _id}
	if err := r.db.Collection("users").FindOne(ctx, filter).Decode(&item); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // Return nil if no user is found, rather than an error.
		}
		return nil, nil
	}

	return item, nil
}

// Followers is the resolver for the followers field.
func (r *taskResolver) Followers(ctx context.Context, obj *model.Task) ([]*model.User, error) {
	var items []*model.User
	var ids []primitive.ObjectID

	if len(obj.Shares) == 0 {
		return nil, nil // No IDs, return empty result
	}

	for _, id := range obj.Shares {
		_id, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, nil
		}
		ids = append(ids, _id)
	}

	if len(ids) == 0 {
		return nil, nil // No IDs, return empty result
	}

	filter := bson.M{"_id": bson.M{"$in": ids}}
	cursor, err := r.db.Collection("users").Find(ctx, filter)
	if err != nil {
		return nil, nil
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &items); err != nil {
		return nil, nil
	}

	return items, nil
}

// Task returns TaskResolver implementation.
func (r *Resolver) Task() TaskResolver { return &taskResolver{r} }

type taskResolver struct{ *Resolver }
