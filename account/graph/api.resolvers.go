package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"
	"time"

	"github.com/dailytravel/x/account/graph/model"
	"github.com/dailytravel/x/account/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ID is the resolver for the id field.
func (r *apiResolver) ID(ctx context.Context, obj *model.Api) (string, error) {
	return obj.ID.Hex(), nil
}

// Metadata is the resolver for the metadata field.
func (r *apiResolver) Metadata(ctx context.Context, obj *model.Api) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// Created is the resolver for the created field.
func (r *apiResolver) Created(ctx context.Context, obj *model.Api) (string, error) {
	return time.Unix(int64(obj.Created.T), 0).Format(time.RFC3339), nil
}

// Updated is the resolver for the updated field.
func (r *apiResolver) Updated(ctx context.Context, obj *model.Api) (string, error) {
	return time.Unix(int64(obj.Updated.T), 0).Format(time.RFC3339), nil
}

// CreateAPI is the resolver for the createApi field.
func (r *mutationResolver) CreateAPI(ctx context.Context, input model.NewAPI) (*model.Api, error) {
	// uid, err := utils.UID(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	item := &model.Api{}

	_, err := r.db.Collection(item.Collection()).InsertOne(ctx, item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

// UpdateAPI is the resolver for the updateApi field.
func (r *mutationResolver) UpdateAPI(ctx context.Context, id string, input model.UpdateAPI) (*model.Api, error) {
	// uid, err := utils.UID(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	// Convert the ID string to ObjectID
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Find the item by ID
	item := &model.Api{}
	filter := bson.M{"_id": _id}
	err = r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(item)
	if err != nil {
		return nil, err
	}

	if err := r.db.Collection(item.Collection()).FindOneAndUpdate(ctx, filter, item).Decode(item); err != nil {
		return nil, err
	}

	return item, nil
}

// DeleteAPI is the resolver for the deleteApi field.
func (r *mutationResolver) DeleteAPI(ctx context.Context, id string) (map[string]interface{}, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	res, err := r.db.Collection("apis").DeleteOne(ctx, bson.M{"_id": _id})
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

// DeleteApis is the resolver for the deleteApis field.
func (r *mutationResolver) DeleteApis(ctx context.Context, ids []string) (map[string]interface{}, error) {
	_ids := make([]primitive.ObjectID, len(ids))
	for i, id := range ids {
		_id, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		_ids[i] = _id
	}
	filter := bson.M{"_id": bson.M{"$in": _ids}}

	res, err := r.db.Collection("apis").DeleteMany(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("error deleting apis: %v", err)
	}

	if res.DeletedCount == 0 {
		return nil, fmt.Errorf("apis not found")
	}

	return map[string]interface{}{
		"success": true,
	}, nil
}

// API is the resolver for the api field.
func (r *queryResolver) API(ctx context.Context, id string) (*model.Api, error) {
	item := &model.Api{}
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	if err := r.db.Collection(item.Collection()).FindOne(ctx, bson.M{"_id": _id}).Decode(item); err != nil {
		return nil, err
	}

	return item, nil
}

// Apis is the resolver for the apis field.
func (r *queryResolver) Apis(ctx context.Context, filter map[string]interface{}, project map[string]interface{}, sort map[string]interface{}, collation map[string]interface{}, limit *int, skip *int) (*model.Apis, error) {
	var items []*model.Api

	// Convert map to bson.M which is a type alias for map[string]interface{}
	_filter := utils.Filter(filter)
	opts := utils.Sort(sort)

	if project != nil {
		opts.SetProjection(project)
	}
	if limit != nil {
		opts.SetLimit(int64(*limit))
	}
	if skip != nil {
		opts.SetSkip(int64(*skip))
	}

	cursor, err := r.db.Collection("apis").Find(ctx, _filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &items); err != nil {
		return nil, err
	}

	//get total count
	count, err := r.db.Collection("apis").CountDocuments(ctx, _filter, nil)
	if err != nil {
		return nil, err
	}

	return &model.Apis{
		Count: int(count),
		Data:  items,
	}, nil
}

// Api returns ApiResolver implementation.
func (r *Resolver) Api() ApiResolver { return &apiResolver{r} }

type apiResolver struct{ *Resolver }
