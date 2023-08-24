package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"time"

	"github.com/dailytravel/x/account/auth"
	"github.com/dailytravel/x/account/graph/model"
	"github.com/dailytravel/x/account/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ID is the resolver for the id field.
func (r *keyResolver) ID(ctx context.Context, obj *model.Key) (string, error) {
	return obj.ID.Hex(), nil
}

// ExpiresAt is the resolver for the expires_at field.
func (r *keyResolver) ExpiresAt(ctx context.Context, obj *model.Key) (*string, error) {
	if obj.ExpiresAt != nil {
		expiresAtStr := time.Unix(int64(obj.ExpiresAt.T), 0).Format(time.RFC3339)
		return &expiresAtStr, nil
	}
	return nil, nil
}

// CreatedAt is the resolver for the created_at field.
func (r *keyResolver) CreatedAt(ctx context.Context, obj *model.Key) (string, error) {
	return time.Unix(int64(obj.CreatedAt.T), 0).Format(time.RFC3339), nil
}

// UpdatedAt is the resolver for the updated_at field.
func (r *keyResolver) UpdatedAt(ctx context.Context, obj *model.Key) (string, error) {
	return time.Unix(int64(obj.UpdatedAt.T), 0).Format(time.RFC3339), nil
}

// Owner is the resolver for the owner field.
func (r *keyResolver) User(ctx context.Context, obj *model.Key) (*model.User, error) {
	var item *model.User
	col := r.db.Collection(item.Collection())

	filter := bson.M{"_id": obj.UID}
	options := options.FindOne().SetProjection(bson.M{"_id": 1, "name": 1, "email": 1, "photos": 1})

	err := col.FindOne(ctx, filter, options).Decode(&item)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return item, nil
}

// RevokeKey is the resolver for the revokeKey field.
func (r *mutationResolver) RevokeKey(ctx context.Context) (*model.Key, error) {
	col := r.db.Collection("keys")

	// Revoke current and previous keys by setting their status to "archived" and "revoked_at" to the current time
	filter := bson.M{"status": bson.M{"$in": []string{"current", "previous"}}}
	update := bson.M{"$set": bson.M{"status": "archived", "revoked_at": time.Now().Unix()}}
	_, err := col.UpdateMany(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	// Update the next key to become the current key
	_, err = col.UpdateOne(ctx, bson.M{"status": "next"}, bson.M{"$set": bson.M{"status": "current"}})
	if err != nil {
		return nil, err
	}

	// Generate a new key for the "next" status
	privateKey, err := auth.GenerateRSAKeyPair(2048)
	if err != nil {
		return nil, err
	}

	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	publicKey := &privateKey.PublicKey

	fingerprint, err := auth.CalculateFingerprint(publicKey)
	if err != nil {
		return nil, err
	}

	thumbprintSHA256, err := auth.CalculateThumbprint(publicKey, sha256.New())
	if err != nil {
		return nil, err
	}

	// Convert byte slices to hexadecimal strings
	fingerprintHex := hex.EncodeToString(fingerprint)
	thumbprintSHA256Hex := hex.EncodeToString(thumbprintSHA256)

	// Create the new "next" key
	cert := &model.Key{
		Name:        "next",
		Provider:    "local",
		Certificate: string(privateKeyPEM),
		Fingerprint: fingerprintHex,
		Thumbprint:  thumbprintSHA256Hex,
		Type:        "RSA",
		Status:      "next",
	}

	// Insert the new key into the database
	_, err = col.InsertOne(ctx, cert)
	if err != nil {
		return nil, err
	}

	// Update the JWKS file with the new key
	if err := auth.CreateJWKSFile(col); err != nil {
		return nil, err
	}

	// Return the new key information (optional)
	return cert, nil
}

// RotateKey is the resolver for the rotateKey field.
func (r *mutationResolver) RotateKey(ctx context.Context) (*model.Key, error) {
	col := r.db.Collection("keys")

	// Revoke the "previous" key by setting its status to "archived" and "revoked_at" to the current time
	filter := bson.M{"status": "previous"}
	update := bson.M{"$set": bson.M{"status": "archived", "revoked_at": time.Now().Unix()}}
	_, err := col.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	// Update the "current" key to become the "previous" key
	filter = bson.M{"status": "current"}
	update = bson.M{"$set": bson.M{"status": "previous"}}
	_, err = col.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	// Update the "next" key to become the "current" key
	filter = bson.M{"status": "next"}
	update = bson.M{"$set": bson.M{"status": "current"}}
	_, err = col.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	// Generate a new key for the "next" status
	privateKey, err := auth.GenerateRSAKeyPair(2048)
	if err != nil {
		return nil, err
	}

	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	publicKey := &privateKey.PublicKey

	fingerprint, err := auth.CalculateFingerprint(publicKey)
	if err != nil {
		return nil, err
	}

	thumbprintSHA256, err := auth.CalculateThumbprint(publicKey, sha256.New())
	if err != nil {
		return nil, err
	}

	// Convert byte slices to hexadecimal strings
	fingerprintHex := hex.EncodeToString(fingerprint)
	thumbprintSHA256Hex := hex.EncodeToString(thumbprintSHA256)

	// Create the new "next" key
	cert := &model.Key{
		Name:        "next",
		Provider:    "local",
		Certificate: string(privateKeyPEM),
		Fingerprint: fingerprintHex,
		Thumbprint:  thumbprintSHA256Hex,
		Type:        "RSA",
		Status:      "next",
	}

	// Insert the new key into the database
	_, err = col.InsertOne(ctx, cert)
	if err != nil {
		return nil, err
	}

	// Update the JWKS file with the new key
	if err := auth.CreateJWKSFile(col); err != nil {
		return nil, err
	}

	// Return the new key information (optional)
	return cert, nil
}

// Keys is the resolver for the keys field.
func (r *queryResolver) Keys(ctx context.Context, args map[string]interface{}) (*model.Keys, error) {
	var items []*model.Key
	//find all items
	cur, err := r.db.Collection(model.KeyCollection).Find(ctx, utils.Query(args), utils.Options(args))
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var item *model.Key
		if err := cur.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	//get total count
	count, err := r.db.Collection(model.KeyCollection).CountDocuments(ctx, utils.Query(args), nil)
	if err != nil {
		return nil, err
	}

	return &model.Keys{
		Count: int(count),
		Data:  items,
	}, nil
}

// Key is the resolver for the key field.
func (r *queryResolver) Key(ctx context.Context, id string) (*model.Key, error) {
	var item *model.Key
	col := r.db.Collection(item.Collection())
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	if err := col.FindOne(ctx, bson.M{"_id": _id}).Decode(&item); err != nil {
		return nil, nil
	}

	return item, nil
}

// Key returns KeyResolver implementation.
func (r *Resolver) Key() KeyResolver { return &keyResolver{r} }

type keyResolver struct{ *Resolver }
