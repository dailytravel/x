package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"
	"time"

	"github.com/dailytravel/x/account/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreatePermission is the resolver for the createPermission field.
func (r *mutationResolver) CreatePermission(ctx context.Context, input model.NewPermission) (*model.Permission, error) {
	// uid, err := utils.UID(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	item := &model.Permission{
		Name:        input.Name,
		Description: input.Description,
	}

	res, err := r.db.Collection(item.Collection()).InsertOne(ctx, item, nil)
	if err != nil {
		return nil, err
	}

	item.ID = res.InsertedID.(primitive.ObjectID)

	return item, nil
}

// UpdatePermission is the resolver for the updatePermission field.
func (r *mutationResolver) UpdatePermission(ctx context.Context, id string, input model.UpdatePermission) (*model.Permission, error) {
	// uid, err := utils.UID(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": _id}
	item := &model.Permission{}

	if input.Name != nil {
		item.Name = *input.Name
	}

	if input.Description != nil {
		item.Description = input.Description
	}

	if err := r.db.Collection(item.Collection()).FindOneAndUpdate(ctx, filter, bson.M{"$set": item}).Decode(item); err != nil {
		return nil, err
	}

	return item, nil
}

// DeletePermission is the resolver for the deletePermission field.
func (r *mutationResolver) DeletePermission(ctx context.Context, id string) (map[string]interface{}, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	res, err := r.db.Collection("permissions").DeleteOne(ctx, bson.M{"_id": _id})
	if err != nil {
		return nil, err
	}

	if res.DeletedCount == 0 {
		return nil, fmt.Errorf("not found")
	}

	return map[string]interface{}{
		"success": true,
	}, nil
}

// DeletePermissions is the resolver for the deletePermissions field.
func (r *mutationResolver) DeletePermissions(ctx context.Context, ids []string) (map[string]interface{}, error) {
	_ids := make([]primitive.ObjectID, len(ids))
	for i, id := range ids {
		_id, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		_ids[i] = _id
	}
	filter := bson.M{"_id": bson.M{"$in": ids}}

	res, err := r.db.Collection("permissions").DeleteMany(ctx, filter)
	if err != nil {
		return nil, err
	}

	if res.DeletedCount == 0 {
		return nil, fmt.Errorf("not found")
	}

	return map[string]interface{}{
		"success": true,
	}, nil
}

// ID is the resolver for the id field.
func (r *permissionResolver) ID(ctx context.Context, obj *model.Permission) (string, error) {
	return obj.ID.Hex(), nil
}

// Created is the resolver for the created field.
func (r *permissionResolver) Created(ctx context.Context, obj *model.Permission) (string, error) {
	return time.Unix(int64(obj.Created.T), 0).Format(time.RFC3339), nil
}

// Updated is the resolver for the updated field.
func (r *permissionResolver) Updated(ctx context.Context, obj *model.Permission) (string, error) {
	return time.Unix(int64(obj.Updated.T), 0).Format(time.RFC3339), nil
}

// API is the resolver for the api field.
func (r *permissionResolver) API(ctx context.Context, obj *model.Permission) (*model.Api, error) {
	panic(fmt.Errorf("not implemented: API - api"))
}

// Permissions is the resolver for the permissions field.
func (r *queryResolver) Permissions(ctx context.Context, stages map[string]interface{}) (*model.Permissions, error) {
	pipeline := bson.A{}

	// Add additional stages to the pipeline
	for key, value := range stages {
		stage := bson.D{{Key: key, Value: value}}
		pipeline = append(pipeline, stage)
	}

	cursor, err := r.db.Collection("permissions").Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var items []*model.Permission

	if err := cursor.All(ctx, &items); err != nil {
		return nil, err
	}

	return &model.Permissions{
		Count: int(cursor.RemainingBatchLength()),
		Data:  items,
	}, nil
}

// Permission is the resolver for the permission field.
func (r *queryResolver) Permission(ctx context.Context, id string) (*model.Permission, error) {
	var item *model.Permission
	col := r.db.Collection(item.Collection())
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	if err := col.FindOne(ctx, bson.M{"_id": _id}).Decode(&item); err != nil {
		return nil, err
	}

	return item, nil
}

// Permission returns PermissionResolver implementation.
func (r *Resolver) Permission() PermissionResolver { return &permissionResolver{r} }

type permissionResolver struct{ *Resolver }
