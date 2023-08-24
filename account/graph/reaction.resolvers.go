package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"errors"
	"fmt"

	"github.com/dailytravel/x/account/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// User is the resolver for the user field.
func (r *reactionResolver) User(ctx context.Context, obj *model.Reaction) (*model.User, error) {
	var item *model.User

	_id, err := primitive.ObjectIDFromHex(obj.UID)
	if err != nil {
		return nil, err
	}

	if err := r.db.Collection(item.Collection()).FindOne(ctx, bson.M{"_id": _id}).Decode(&item); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("no document found for filter %v", bson.M{"_id": _id})
		}
		return nil, err
	}

	return item, nil
}

// Reaction returns ReactionResolver implementation.
func (r *Resolver) Reaction() ReactionResolver { return &reactionResolver{r} }

type reactionResolver struct{ *Resolver }
