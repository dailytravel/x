package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"
	"time"

	"github.com/dailytravel/x/account/graph/model"
	"github.com/dailytravel/x/account/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ID is the resolver for the id field.
func (r *integrationResolver) ID(ctx context.Context, obj *model.Integration) (string, error) {
	return obj.ID.Hex(), nil
}

// Metadata is the resolver for the metadata field.
func (r *integrationResolver) Metadata(ctx context.Context, obj *model.Integration) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// CreatedAt is the resolver for the created_at field.
func (r *integrationResolver) CreatedAt(ctx context.Context, obj *model.Integration) (string, error) {
	return time.Unix(int64(obj.CreatedAt.T), 0).Format(time.RFC3339), nil
}

// UpdatedAt is the resolver for the updated_at field.
func (r *integrationResolver) UpdatedAt(ctx context.Context, obj *model.Integration) (string, error) {
	return time.Unix(int64(obj.UpdatedAt.T), 0).Format(time.RFC3339), nil
}

// CreatedBy is the resolver for the created_by field.
func (r *integrationResolver) CreatedBy(ctx context.Context, obj *model.Integration) (*model.User, error) {
	var item *model.User

	if err := r.db.Collection(model.UserCollection).FindOne(ctx, bson.M{"_id": obj.CreatedBy}).Decode(&item); err != nil {
		return nil, err
	}

	return item, nil
}

// UpdatedBy is the resolver for the updated_by field.
func (r *integrationResolver) UpdatedBy(ctx context.Context, obj *model.Integration) (*model.User, error) {
	var item *model.User

	if err := r.db.Collection(model.UserCollection).FindOne(ctx, bson.M{"_id": obj.UpdatedBy}).Decode(&item); err != nil {
		return nil, err
	}

	return item, nil
}

// CreateIntegration is the resolver for the createIntegration field.
func (r *mutationResolver) CreateIntegration(ctx context.Context, input model.NewIntegration) (*model.Integration, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	item := &model.Integration{
		Model: model.Model{
			CreatedBy: uid,
			UpdatedBy: uid,
		},
	}

	_, err = r.db.Collection(item.Collection()).InsertOne(ctx, item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

// UpdateIntegration is the resolver for the updateIntegration field.
func (r *mutationResolver) UpdateIntegration(ctx context.Context, id string, input model.UpdateIntegration) (*model.Integration, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	// Convert the ID string to ObjectID
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Find the item by ID
	item := &model.Integration{}
	filter := bson.M{"_id": _id}
	err = r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(item)
	if err != nil {
		return nil, err
	}

	if input.Status != nil {
		item.Status = *input.Status
	}

	item.UpdatedBy = uid

	if err := r.db.Collection(item.Collection()).FindOneAndUpdate(ctx, filter, item).Decode(item); err != nil {
		return nil, err
	}

	return item, nil
}

// DeleteIntegration is the resolver for the deleteIntegration field.
func (r *mutationResolver) DeleteIntegration(ctx context.Context, id string) (map[string]interface{}, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	res, err := r.db.Collection("integrations").DeleteOne(ctx, bson.M{"_id": _id})
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

// DeleteIntegrations is the resolver for the deleteIntegrations field.
func (r *mutationResolver) DeleteIntegrations(ctx context.Context, ids []string) (map[string]interface{}, error) {
	_ids := make([]primitive.ObjectID, len(ids))
	for i, id := range ids {
		_id, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		_ids[i] = _id
	}
	filter := bson.M{"_id": bson.M{"$in": _ids}}

	res, err := r.db.Collection("integrations").DeleteMany(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("error deleting integrations: %v", err)
	}

	if res.DeletedCount == 0 {
		return nil, fmt.Errorf("integrations not found")
	}

	return map[string]interface{}{
		"success": true,
	}, nil
}

// Integrations is the resolver for the integrations field.
func (r *queryResolver) Integrations(ctx context.Context, args map[string]interface{}) (*model.Integrations, error) {
	var items []*model.Integration
	//find all items
	cur, err := r.db.Collection(model.RoleCollection).Find(ctx, r.model.Query(args), r.model.Options(args))
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var item *model.Integration
		if err := cur.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	//get total count
	count, err := r.db.Collection("integrations").CountDocuments(ctx, r.model.Query(args), nil)
	if err != nil {
		return nil, err
	}

	return &model.Integrations{
		Count: int(count),
		Data:  items,
	}, nil
}

// Integration is the resolver for the integration field.
func (r *queryResolver) Integration(ctx context.Context, id string) (*model.Integration, error) {
	item := &model.Integration{}
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	if err := r.db.Collection(item.Collection()).FindOne(ctx, bson.M{"_id": _id}).Decode(item); err != nil {
		return nil, err
	}

	return item, nil
}

// Integration returns IntegrationResolver implementation.
func (r *Resolver) Integration() IntegrationResolver { return &integrationResolver{r} }

type integrationResolver struct{ *Resolver }
