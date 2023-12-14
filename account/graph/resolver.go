//go:generate go run github.com/99designs/gqlgen generate
package graph

import (
	"context"
	"fmt"

	"github.com/dailytravel/x/account/graph/model"
	"github.com/dailytravel/x/account/pkg/auth"
	"github.com/go-redis/redis/v8"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/typesense/typesense-go/typesense"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	model model.Model
	db    *mongo.Database
	redis *redis.Client
	ts    *typesense.Client
}

func NewResolver(db *mongo.Database, rdb *redis.Client, ts *typesense.Client) *Resolver {
	return &Resolver{
		db:    db,
		redis: rdb,
		ts:    ts,
		model: model.Model{},
	}
}

func (r *Resolver) getClientByID(ctx context.Context, id primitive.ObjectID) (*model.Client, error) {
	var c *model.Client
	err := r.db.Collection("clients").FindOne(ctx, bson.M{"_id": id}, nil).Decode(&c)
	if err != nil {
		return nil, fmt.Errorf("client not found")
	}
	return c, nil
}

func (r *Resolver) getCurrentKey(ctx context.Context) (*model.Key, error) {
	var k *model.Key

	err := r.db.Collection("keys").FindOne(ctx, bson.M{"status": "current"}, nil).Decode(&k)
	if err != nil {
		return nil, fmt.Errorf("key not found")
	}

	return k, nil
}

func (r *Resolver) getAPIByIdentifier(ctx context.Context, identifier string) (*model.Api, error) {
	var a *model.Api
	err := r.db.Collection("apis").FindOne(ctx, bson.M{"identifier": identifier}, nil).Decode(&a)
	if err != nil {
		return nil, fmt.Errorf("api not found")
	}
	return a, nil
}

func (r *Resolver) getUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var u *model.User
	err := r.db.Collection(u.Collection()).FindOne(ctx, bson.M{"email": email}).Decode(&u)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, fmt.Errorf("error checking for existing user: %v", err)
		}
	}
	return u, nil
}

func (r *Resolver) getCredentialPassword(ctx context.Context, uid primitive.ObjectID) (*model.Credential, error) {
	var c *model.Credential
	err := r.db.Collection(c.Collection()).FindOne(ctx, bson.M{"uid": uid, "type": "PASSWORD"}, nil).Decode(&c)
	if err != nil {
		return nil, fmt.Errorf("credential not found")
	}
	return c, nil
}

func (r *Resolver) insertUser(ctx context.Context, i *model.User) error {
	res, err := r.db.Collection(i.Collection()).InsertOne(ctx, i)
	if err != nil {
		return fmt.Errorf("failed to insert user: %v", err)
	}
	i.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *Resolver) insertCredentials(ctx context.Context, docs []interface{}) error {
	if _, err := r.db.Collection("credentials").InsertMany(ctx, docs); err != nil {
		return fmt.Errorf("failed to insert credentials: %v", err)
	}
	return nil
}

func (r *Resolver) generateTokens(ctx context.Context, claims jwt.MapClaims) (*string, error) {
	k, err := r.getCurrentKey(ctx)
	if err != nil {
		return nil, err
	}

	token, err := auth.Token(claims, *k)

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (r *Resolver) insertToken(ctx context.Context, document *model.Token) (*model.Token, error) {
	res, err := r.db.Collection(document.Collection()).InsertOne(ctx, document)
	if err != nil {
		return nil, fmt.Errorf("failed to insert token: %v", err)
	}

	document.ID = res.InsertedID.(primitive.ObjectID)

	return document, nil
}
