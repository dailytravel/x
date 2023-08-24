package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"errors"
	"time"

	"github.com/dailytravel/x/hrm/auth"
	"github.com/dailytravel/x/hrm/graph/model"
	"github.com/dailytravel/x/hrm/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ID is the resolver for the id field.
func (r *jobResolver) ID(ctx context.Context, obj *model.Job) (string, error) {
	return obj.ID.Hex(), nil
}

// Title is the resolver for the title field.
func (r *jobResolver) Title(ctx context.Context, obj *model.Job) (string, error) {
	// Get the locale from the context
	locale := auth.Locale(ctx)
	if locale == nil {
		locale = &obj.Locale
	}

	// Try to retrieve the title for the requested locale
	if title, ok := obj.Title[*locale].(string); ok {
		return title, nil
	}

	return obj.Title[obj.Locale].(string), nil
}

// Description is the resolver for the description field.
func (r *jobResolver) Description(ctx context.Context, obj *model.Job) (string, error) {
	// Get the locale from the context
	locale := auth.Locale(ctx)
	if locale == nil {
		locale = &obj.Locale
	}

	// Try to retrieve the description for the requested locale
	if description, ok := obj.Description[*locale].(string); ok {
		return description, nil
	}

	return obj.Description[obj.Locale].(string), nil
}

// Requirements is the resolver for the requirements field.
func (r *jobResolver) Requirements(ctx context.Context, obj *model.Job) (string, error) {
	// Get the locale from the context
	locale := auth.Locale(ctx)
	if locale == nil {
		locale = &obj.Locale
	}

	// Try to retrieve the requirements for the requested locale
	if requirements, ok := obj.Requirements[*locale].(string); ok {
		return requirements, nil
	}

	return obj.Requirements[obj.Locale].(string), nil
}

// Skills is the resolver for the skills field.
func (r *jobResolver) Skills(ctx context.Context, obj *model.Job) (*string, error) {
	// Get the locale from the context
	locale := auth.Locale(ctx)
	if locale == nil {
		locale = &obj.Locale
	}

	// Try to retrieve the skills for the requested locale
	if skills, ok := obj.Skills[*locale].(string); ok {
		return &skills, nil
	}

	return obj.Skills[obj.Locale].(*string), nil
}

// Benefits is the resolver for the benefits field.
func (r *jobResolver) Benefits(ctx context.Context, obj *model.Job) (*string, error) {
	// Get the locale from the context
	locale := auth.Locale(ctx)
	if locale == nil {
		locale = &obj.Locale
	}

	// Try to retrieve the benefits for the requested locale
	if benefits, ok := obj.Benefits[*locale].(string); ok {
		return &benefits, nil
	}

	return obj.Benefits[obj.Locale].(*string), nil
}

// Metadata is the resolver for the metadata field.
func (r *jobResolver) Metadata(ctx context.Context, obj *model.Job) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// CreatedAt is the resolver for the created_at field.
func (r *jobResolver) CreatedAt(ctx context.Context, obj *model.Job) (string, error) {
	return time.Unix(int64(obj.CreatedAt.T), 0).Format(time.RFC3339), nil
}

// UpdatedAt is the resolver for the updated_at field.
func (r *jobResolver) UpdatedAt(ctx context.Context, obj *model.Job) (string, error) {
	return time.Unix(int64(obj.UpdatedAt.T), 0).Format(time.RFC3339), nil
}

// CreatedBy is the resolver for the created_by field.
func (r *jobResolver) CreatedBy(ctx context.Context, obj *model.Job) (*string, error) {
	if obj.CreatedBy == nil {
		return nil, nil
	}

	createdBy := obj.CreatedBy.Hex()

	return &createdBy, nil
}

// UpdatedBy is the resolver for the updated_by field.
func (r *jobResolver) UpdatedBy(ctx context.Context, obj *model.Job) (*string, error) {
	if obj.UpdatedBy == nil {
		return nil, nil
	}

	updatedBy := obj.UpdatedBy.Hex()

	return &updatedBy, nil
}

// CreateJob is the resolver for the createJob field.
func (r *mutationResolver) CreateJob(ctx context.Context, input model.NewJob) (*model.Job, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	item := &model.Job{
		Locale:       input.Locale,
		Title:        bson.M{input.Locale: input.Title},
		Description:  bson.M{input.Locale: input.Description},
		Requirements: bson.M{input.Locale: input.Requirements},
		Skills:       bson.M{input.Locale: input.Skills},
		Benefits:     bson.M{input.Locale: input.Benefits},
		Status:       input.Status,
		Model: model.Model{
			CreatedBy: uid,
			UpdatedBy: uid,
		},
	}

	// Insert the new organization
	if _, err := r.db.Collection(item.Collection()).InsertOne(ctx, item, nil); err != nil {
		return nil, err
	}

	return item, nil
}

// UpdateJob is the resolver for the updateJob field.
func (r *mutationResolver) UpdateJob(ctx context.Context, id string, input model.UpdateJob) (*model.Job, error) {
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
	item := &model.Job{}
	filter := bson.M{"_id": _id}
	err = r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(item)
	if err != nil {
		return nil, err
	}

	// Update the position's fields
	if input.Title != nil {
		item.Title[*input.Locale] = *input.Title
	}

	if input.Description != nil {
		item.Description[*input.Locale] = *input.Description
	}

	if input.Requirements != nil {
		item.Requirements[*input.Locale] = *input.Requirements
	}

	if input.Skills != nil {
		item.Skills[*input.Locale] = *input.Skills
	}

	if input.Benefits != nil {
		item.Benefits[*input.Locale] = *input.Benefits
	}

	if input.Status != nil {
		item.Status = *input.Status
	}

	item.UpdatedBy = uid

	// Update the position in the database
	if err := r.db.Collection(item.Collection()).FindOneAndUpdate(ctx, filter, item).Decode(item); err != nil {
		return nil, err
	}

	return item, nil
}

// DeleteJob is the resolver for the deleteJob field.
func (r *mutationResolver) DeleteJob(ctx context.Context, id string) (map[string]interface{}, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	// Convert the ID string to ObjectID
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Find the position by ID
	position := &model.Job{}
	filter := bson.M{"_id": _id}
	err = r.db.Collection(position.Collection()).FindOne(ctx, filter).Decode(position)
	if err != nil {
		return nil, err
	}

	// Define the update to mark the position as deleted
	update := bson.M{
		"$set": bson.M{
			"deleted_at": primitive.Timestamp{T: uint32(time.Now().Unix())},
			"deleted_by": uid,
			"status":     "deleted",
			"updated_by": uid,
			"updated_at": primitive.Timestamp{T: uint32(time.Now().Unix())},
		},
	}

	// Perform the update operation in the database
	_, err = r.db.Collection(position.Collection()).UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"status": "success"}, nil
}

// DeleteJobs is the resolver for the deleteJobs field.
func (r *mutationResolver) DeleteJobs(ctx context.Context, ids []string) (map[string]interface{}, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	// Convert the list of ID strings to ObjectIDs
	var objectIDs []primitive.ObjectID
	for _, id := range ids {
		_id, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		objectIDs = append(objectIDs, _id)
	}

	// Define the filter to match the given IDs
	filter := bson.M{"_id": bson.M{"$in": objectIDs}}

	// Define the update to mark positions as deleted
	update := bson.M{
		"$set": bson.M{
			"deleted_at": primitive.Timestamp{T: uint32(time.Now().Unix())},
			"deleted_by": uid,
			"status":     "deleted",
			"updated_by": uid,
			"updated_at": primitive.Timestamp{T: uint32(time.Now().Unix())},
		},
	}

	// Perform the update operation in the database
	result, err := r.db.Collection("positions").UpdateMany(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"status": "success", "deletedCount": result.ModifiedCount}, nil
}

// Job is the resolver for the job field.
func (r *queryResolver) Job(ctx context.Context, id string) (*model.Job, error) {
	var item *model.Job

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": _id}

	if err := r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(&item); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return item, nil
}

// Jobs is the resolver for the jobs field.
func (r *queryResolver) Jobs(ctx context.Context, args map[string]interface{}) (map[string]interface{}, error) {
	res, err := r.ts.Collection("jobs").Documents().Search(utils.Params(args))
	if err != nil {
		return nil, err
	}

	// Convert struct to map
	results, err := utils.StructToMap(res)
	if err != nil {
		return nil, err
	}

	return results, nil
}

// Job returns JobResolver implementation.
func (r *Resolver) Job() JobResolver { return &jobResolver{r} }

type jobResolver struct{ *Resolver }
