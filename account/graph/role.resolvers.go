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

// CreateRole is the resolver for the createRole field.
func (r *mutationResolver) CreateRole(ctx context.Context, input model.NewRole) (*model.Role, error) {
	// Convert permission names to primitive.ObjectID values
	permissions := make([]*primitive.ObjectID, len(input.Permissions))
	for i, permissionNamePtr := range input.Permissions {
		permissionName := *permissionNamePtr                           // Dereference the pointer to get the string value
		permissionID, err := primitive.ObjectIDFromHex(permissionName) // Convert permission name to ObjectID if needed
		if err != nil {
			return nil, err
		}
		permissions[i] = &permissionID
	}

	item := &model.Role{
		Name:        input.Name,
		Description: input.Description,
		Permissions: permissions, // Assign the converted permission IDs

	}

	res, err := r.db.Collection(item.Collection()).InsertOne(ctx, item, nil)
	if err != nil {
		return nil, err
	}

	item.ID = res.InsertedID.(primitive.ObjectID)

	return item, nil
}

// UpdateRole is the resolver for the updateRole field.
func (r *mutationResolver) UpdateRole(ctx context.Context, id string, input model.UpdateRole) (*model.Role, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": _id}
	item := &model.Role{}

	if input.Name != nil {
		item.Name = *input.Name
	}

	if input.Description != nil {
		item.Description = input.Description
	}

	if input.Permissions != nil {
		// Convert permission names to primitive.ObjectID values
		permissions := make([]*primitive.ObjectID, len(input.Permissions))
		for i, permissionNamePtr := range input.Permissions {
			permissionName := *permissionNamePtr                           // Dereference the pointer to get the string value
			permissionID, err := primitive.ObjectIDFromHex(permissionName) // Convert permission name to ObjectID if needed
			if err != nil {
				return nil, err
			}
			permissions[i] = &permissionID
		}
		item.Permissions = permissions // Assign the converted permission IDs
	}

	if err := r.db.Collection(item.Collection()).FindOneAndUpdate(ctx, filter, bson.M{"$set": item}).Decode(item); err != nil {
		return nil, err
	}

	return item, nil
}

// DeleteRole is the resolver for the deleteRole field.
func (r *mutationResolver) DeleteRole(ctx context.Context, id string) (map[string]interface{}, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	res, err := r.db.Collection("roles").DeleteOne(ctx, bson.M{"_id": _id})
	if err != nil {
		return nil, fmt.Errorf("error deleting role: %v", err)
	}

	if res.DeletedCount == 0 {
		return nil, fmt.Errorf("role not found")
	}

	return map[string]interface{}{
		"success": true,
	}, nil
}

// DeleteRoles is the resolver for the deleteRoles field.
func (r *mutationResolver) DeleteRoles(ctx context.Context, ids []string) (map[string]interface{}, error) {
	_ids := make([]primitive.ObjectID, len(ids))
	for i, id := range ids {
		_id, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		_ids[i] = _id
	}
	filter := bson.M{"_id": bson.M{"$in": _ids}}

	res, err := r.db.Collection("roles").DeleteMany(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("error deleting roles: %v", err)
	}

	if res.DeletedCount == 0 {
		return nil, fmt.Errorf("roles not found")
	}

	return map[string]interface{}{
		"success": true,
	}, nil
}

// Role is the resolver for the role field.
func (r *queryResolver) Role(ctx context.Context, id string) (*model.Role, error) {
	var item *model.Role
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	if err := r.db.Collection(item.Collection()).FindOne(ctx, bson.M{"_id": _id}).Decode(item); err != nil {
		return nil, err
	}

	return item, nil
}

// Roles is the resolver for the roles field.
func (r *queryResolver) Roles(ctx context.Context, stages map[string]interface{}) (*model.Roles, error) {
	pipeline := bson.A{}

	// Add additional stages to the pipeline
	for key, value := range stages {
		stage := bson.D{{Key: key, Value: value}}
		pipeline = append(pipeline, stage)
	}

	cursor, err := r.db.Collection("roles").Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var items []*model.Role

	if err := cursor.All(ctx, &items); err != nil {
		return nil, err
	}

	return &model.Roles{
		Count: int(cursor.RemainingBatchLength()),
		Data:  items,
	}, nil
}

// ID is the resolver for the id field.
func (r *roleResolver) ID(ctx context.Context, obj *model.Role) (string, error) {
	return obj.ID.Hex(), nil
}

// Permissions is the resolver for the permissions field.
func (r *roleResolver) Permissions(ctx context.Context, obj *model.Role) ([]*model.Permission, error) {
	if obj.Permissions == nil || len(obj.Permissions) == 0 {
		return []*model.Permission{}, nil
	}

	cursor, err := r.db.Collection("permissions").Find(ctx, bson.M{"_id": bson.M{"$in": obj.Permissions}}, nil)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var items []*model.Permission
	for cursor.Next(ctx) {
		var item *model.Permission
		if err := cursor.Decode(&item); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

// Created is the resolver for the created field.
func (r *roleResolver) Created(ctx context.Context, obj *model.Role) (string, error) {
	return time.Unix(int64(obj.Created.T), 0).Format(time.RFC3339), nil
}

// Updated is the resolver for the updated field.
func (r *roleResolver) Updated(ctx context.Context, obj *model.Role) (string, error) {
	return time.Unix(int64(obj.Updated.T), 0).Format(time.RFC3339), nil
}

// Role returns RoleResolver implementation.
func (r *Resolver) Role() RoleResolver { return &roleResolver{r} }

type roleResolver struct{ *Resolver }
