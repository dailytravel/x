package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dailytravel/x/account/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateToken is the resolver for the createToken field.
func (r *mutationResolver) CreateToken(ctx context.Context, input model.NewToken) (*model.Token, error) {
	panic(fmt.Errorf("not implemented: CreateToken - createToken"))
}

// UpdateToken is the resolver for the updateToken field.
func (r *mutationResolver) UpdateToken(ctx context.Context, id string, input model.UpdateToken) (*model.Token, error) {
	panic(fmt.Errorf("not implemented: UpdateToken - updateToken"))
}

// DeleteToken is the resolver for the deleteToken field.
func (r *mutationResolver) DeleteToken(ctx context.Context, id string) (*model.Token, error) {
	panic(fmt.Errorf("not implemented: DeleteToken - deleteToken"))
}

// DeleteTokens is the resolver for the deleteTokens field.
func (r *mutationResolver) DeleteTokens(ctx context.Context, ids []string) ([]*model.Token, error) {
	panic(fmt.Errorf("not implemented: DeleteTokens - deleteTokens"))
}

// Tokens is the resolver for the tokens field.
func (r *queryResolver) Tokens(ctx context.Context, args map[string]interface{}) (*model.Tokens, error) {
	panic(fmt.Errorf("not implemented: Tokens - tokens"))
}

// Token is the resolver for the token field.
func (r *queryResolver) Token(ctx context.Context, id string) (*model.Token, error) {
	var item *model.Token

	_id, err := primitive.ObjectIDFromHex(id)
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

// ID is the resolver for the id field.
func (r *tokenResolver) ID(ctx context.Context, obj *model.Token) (string, error) {
	return obj.ID.Hex(), nil
}

// User is the resolver for the user field.
func (r *tokenResolver) User(ctx context.Context, obj *model.Token) (*model.User, error) {
	var item *model.User

	if err := r.db.Collection(item.Collection()).FindOne(ctx, bson.M{"_id": obj.ID}).Decode(&item); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("no document found for filter %v", bson.M{"_id": obj.ID})
		}
		return nil, err
	}

	return item, nil
}

// LastUsedAt is the resolver for the last_used_at field.
func (r *tokenResolver) LastUsedAt(ctx context.Context, obj *model.Token) (string, error) {
	return time.Unix(int64(obj.LastUsedAt.T), 0).Format(time.RFC3339), nil
}

// ExpiresAt is the resolver for the expires_at field.
func (r *tokenResolver) ExpiresAt(ctx context.Context, obj *model.Token) (string, error) {
	return time.Unix(int64(obj.ExpiresAt.T), 0).Format(time.RFC3339), nil
}

// CreatedAt is the resolver for the created_at field.
func (r *tokenResolver) CreatedAt(ctx context.Context, obj *model.Token) (string, error) {
	return time.Unix(int64(obj.CreatedAt.T), 0).Format(time.RFC3339), nil
}

// UpdatedAt is the resolver for the updated_at field.
func (r *tokenResolver) UpdatedAt(ctx context.Context, obj *model.Token) (string, error) {
	return time.Unix(int64(obj.UpdatedAt.T), 0).Format(time.RFC3339), nil
}

// Token returns TokenResolver implementation.
func (r *Resolver) Token() TokenResolver { return &tokenResolver{r} }

type tokenResolver struct{ *Resolver }
