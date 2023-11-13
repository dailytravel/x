package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/dailytravel/x/base/graph/model"
	"github.com/dailytravel/x/base/internal/utils"
	"github.com/typesense/typesense-go/typesense/api/pointer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateTask is the resolver for the createTask field.
func (r *mutationResolver) CreateTask(ctx context.Context, input model.NewTask) (*model.Task, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		log.Printf("error getting uid: %v", err)
		return nil, err
	}

	// Create a new task
	item := &model.Task{
		UID:  *uid,
		Name: input.Name,
	}

	if input.Priority != nil {
		item.Priority = input.Priority
	}

	if input.Notes != nil {
		item.Notes = input.Notes
	}

	if input.Status != nil {
		item.Status = *input.Status
	}

	if input.Metadata != nil {
		if item.Metadata == nil {
			item.Metadata = make(map[string]interface{})
		}
		for k, v := range input.Metadata {
			item.Metadata[k] = v
		}
	}

	if input.Parent != nil {
		_id, err := primitive.ObjectIDFromHex(*input.Parent)
		if err != nil {
			return nil, err
		}
		item.Parent = &_id
	}

	if input.Start != nil {
		startAt, err := time.Parse(time.RFC3339, *input.Start)
		if err != nil {
			return nil, fmt.Errorf("invalid Start format: %v", err)
		}
		dt := primitive.NewDateTimeFromTime(startAt)
		item.Start = &dt
	}

	if input.Order != nil {
		item.Order = input.Order
	}

	if input.Metadata != nil {
		if item.Metadata == nil {
			item.Metadata = make(map[string]interface{})
		}

		for k, v := range input.Metadata {
			item.Metadata[k] = v
		}
	}

	res, err := r.db.Collection(item.Collection()).InsertOne(ctx, item)
	if err != nil {
		return nil, err
	}

	item.ID = res.InsertedID.(primitive.ObjectID)

	return item, nil
}

// UpdateTask is the resolver for the updateTask field.
func (r *mutationResolver) UpdateTask(ctx context.Context, id string, input model.UpdateTask) (*model.Task, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Retrieve the existing task using its ID
	item := &model.Task{}
	filter := bson.M{"_id": _id}
	err = r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(item)
	if err != nil {
		return nil, err
	}

	// Update the task fields based on the input
	if input.Name != nil {
		item.Name = *input.Name
	}

	if input.Priority != nil {
		item.Priority = input.Priority
	}

	if input.Notes != nil {
		item.Notes = input.Notes
	}

	if input.Status != nil {
		item.Status = *input.Status
	}

	if input.Parent != nil {
		_id, err := primitive.ObjectIDFromHex(*input.Parent)
		if err != nil {
			return nil, err
		}
		item.Parent = &_id
	}

	if input.Start != nil {
		startAt, err := time.Parse(time.RFC3339, *input.Start)
		if err != nil {
			return nil, fmt.Errorf("invalid Start format: %v", err)
		}
		dt := primitive.NewDateTimeFromTime(startAt)
		item.Start = &dt
	}

	if input.Order != nil {
		item.Order = input.Order
	}

	if input.Metadata != nil {
		if item.Metadata == nil {
			item.Metadata = make(map[string]interface{})
		}

		for k, v := range input.Metadata {
			item.Metadata[k] = v
		}
	}

	// Update the task in the database
	update := bson.M{"$set": item}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updatedTask model.Task
	err = r.db.Collection("tasks").FindOneAndUpdate(ctx, filter, update, opts).Decode(&updatedTask)
	if err != nil {
		return nil, err
	}

	return &updatedTask, nil
}

// DeleteTask is the resolver for the deleteTask field.
func (r *mutationResolver) DeleteTask(ctx context.Context, id string) (map[string]interface{}, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"_id": objectID,
		"$or": []bson.M{
			{"uid": uid},
			{"collaborators": uid},
			{"assignments": bson.M{"$elemMatch": bson.M{"uid": uid}}},
		},
	}

	result, err := r.db.Collection("tasks").DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	if result.DeletedCount == 0 {
		return nil, errors.New("task not found or unauthorized to delete")
	}

	response := map[string]interface{}{"success": "Task deleted successfully"}
	return response, nil
}

// DeleteTasks is the resolver for the deleteTasks field.
func (r *mutationResolver) DeleteTasks(ctx context.Context, ids []string) (map[string]interface{}, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	objectIDs := make([]primitive.ObjectID, len(ids))
	for i, id := range ids {
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		objectIDs[i] = objID
	}

	filter := bson.M{
		"_id": bson.M{"$in": objectIDs},
		"$or": []bson.M{
			{"uid": uid},
			{"collaborators": uid},
			{"assignments": bson.M{"$elemMatch": bson.M{"uid": uid}}},
		},
	}

	result, err := r.db.Collection("tasks").DeleteMany(ctx, filter)
	if err != nil {
		return nil, err
	}

	if result.DeletedCount != int64(len(ids)) {
		return nil, errors.New("tasks not found or unauthorized to delete")
	}

	response := map[string]interface{}{"success": "Tasks deleted successfully"}
	return response, nil
}

// Task is the resolver for the task field.
func (r *queryResolver) Task(ctx context.Context, id string) (*model.Task, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"_id": objectID,
		"$or": []bson.M{
			{"UID": uid},
			{"collaborators": uid},
			{"assignments": bson.M{"$elemMatch": bson.M{"uid": uid}}},
		},
	}

	var task model.Task
	err = r.db.Collection("tasks").FindOne(ctx, filter).Decode(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

// Tasks is the resolver for the tasks field.
func (r *queryResolver) Tasks(ctx context.Context, stages map[string]interface{}) (*model.Tasks, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	matchStage := bson.M{
		"$match": bson.M{
			"$and": []interface{}{
				bson.M{"status": bson.M{"$ne": "ARCHIVED"}},
				bson.M{"$or": bson.A{
					bson.M{"uid": uid},
					bson.M{"collaborators": uid},
					bson.M{
						"assignments.task": "$_id",
						"assignments.uid":  uid,
					},
				}},
			},
		},
	}

	pipeline := bson.A{matchStage}

	for key, value := range stages {
		stage := bson.D{{Key: key, Value: value}}
		pipeline = append(pipeline, stage)
	}

	cursor, err := r.db.Collection("tasks").Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var items []*model.Task
	if err = cursor.All(ctx, &items); err != nil {
		return nil, err
	}

	return &model.Tasks{
		Count: len(items),
		Data:  items,
	}, nil
}

// ID is the resolver for the id field.
func (r *taskResolver) ID(ctx context.Context, obj *model.Task) (string, error) {
	return obj.ID.Hex(), nil
}

// Parent is the resolver for the parent field.
func (r *taskResolver) Parent(ctx context.Context, obj *model.Task) (*model.Task, error) {
	var item *model.Task

	filter := bson.M{"_id": obj.Parent}
	options := options.FindOne().SetProjection(bson.M{"_id": 1, "name": 1})

	err := r.db.Collection(item.Collection()).FindOne(ctx, filter, options).Decode(&item)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return item, nil
}

// Subtasks is the resolver for the subtasks field.
func (r *taskResolver) Subtasks(ctx context.Context, obj *model.Task) ([]*model.Task, error) {
	var items []*model.Task

	filter := bson.M{"parent": obj.ID}
	opts := options.Find().SetSort(bson.M{"order": 1})

	cur, err := r.db.Collection("tasks").Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var item model.Task
		if err := cur.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, &item)
	}

	return items, nil
}

// Start is the resolver for the start field.
func (r *taskResolver) Start(ctx context.Context, obj *model.Task) (*string, error) {
	if obj.Start == nil {
		return nil, nil
	}

	return pointer.String(obj.Start.Time().Format(time.RFC3339)), nil
}

// End is the resolver for the end field.
func (r *taskResolver) End(ctx context.Context, obj *model.Task) (*string, error) {
	if obj.End == nil {
		return nil, nil
	}

	return pointer.String(obj.End.Time().Format(time.RFC3339)), nil
}

// Metadata is the resolver for the metadata field.
func (r *taskResolver) Metadata(ctx context.Context, obj *model.Task) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// UID is the resolver for the uid field.
func (r *taskResolver) UID(ctx context.Context, obj *model.Task) (string, error) {
	return obj.UID.Hex(), nil
}

// Assignee is the resolver for the assignee field.
func (r *taskResolver) Assignee(ctx context.Context, obj *model.Task) (*string, error) {
	if obj.Assignee == nil {
		return nil, nil
	}

	return pointer.String(obj.Assignee.Hex()), nil
}

// Members is the resolver for the members field.
func (r *taskResolver) Members(ctx context.Context, obj *model.Task) ([]string, error) {
	var members []string

	for _, member := range obj.Members {
		members = append(members, member.Hex())
	}

	return members, nil
}

// Created is the resolver for the created field.
func (r *taskResolver) Created(ctx context.Context, obj *model.Task) (string, error) {
	return time.Unix(int64(obj.Created.T), 0).Format(time.RFC3339), nil
}

// Updated is the resolver for the updated field.
func (r *taskResolver) Updated(ctx context.Context, obj *model.Task) (string, error) {
	return time.Unix(int64(obj.Updated.T), 0).Format(time.RFC3339), nil
}

// Task returns TaskResolver implementation.
func (r *Resolver) Task() TaskResolver { return &taskResolver{r} }

type taskResolver struct{ *Resolver }
