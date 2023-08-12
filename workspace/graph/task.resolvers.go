package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"
	"time"

	"github.com/dailytravel/x/workspace/graph/model"
)

// CreateTask is the resolver for the createTask field.
func (r *mutationResolver) CreateTask(ctx context.Context, input model.NewTask) (*model.Task, error) {
	panic(fmt.Errorf("not implemented: CreateTask - createTask"))
}

// UpdateTask is the resolver for the updateTask field.
func (r *mutationResolver) UpdateTask(ctx context.Context, id string, input model.UpdateTask) (*model.Task, error) {
	panic(fmt.Errorf("not implemented: UpdateTask - updateTask"))
}

// DeleteTask is the resolver for the deleteTask field.
func (r *mutationResolver) DeleteTask(ctx context.Context, id string) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: DeleteTask - deleteTask"))
}

// DeleteTasks is the resolver for the deleteTasks field.
func (r *mutationResolver) DeleteTasks(ctx context.Context, ids []string) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: DeleteTasks - deleteTasks"))
}

// Task is the resolver for the task field.
func (r *queryResolver) Task(ctx context.Context, id string) (*model.Task, error) {
	panic(fmt.Errorf("not implemented: Task - task"))
}

// Tasks is the resolver for the tasks field.
func (r *queryResolver) Tasks(ctx context.Context, list string) (*model.Tasks, error) {
	panic(fmt.Errorf("not implemented: Tasks - tasks"))
}

// ID is the resolver for the id field.
func (r *taskResolver) ID(ctx context.Context, obj *model.Task) (string, error) {
	return obj.ID.Hex(), nil
}

// User is the resolver for the user field.
func (r *taskResolver) User(ctx context.Context, obj *model.Task) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Parent is the resolver for the parent field.
func (r *taskResolver) Parent(ctx context.Context, obj *model.Task) (*model.Task, error) {
	panic(fmt.Errorf("not implemented: Parent - parent"))
}

// Subtasks is the resolver for the subtasks field.
func (r *taskResolver) Subtasks(ctx context.Context, obj *model.Task) ([]*model.Task, error) {
	panic(fmt.Errorf("not implemented: Subtasks - subtasks"))
}

// List is the resolver for the list field.
func (r *taskResolver) List(ctx context.Context, obj *model.Task) (*model.List, error) {
	panic(fmt.Errorf("not implemented: List - list"))
}

// StartDate is the resolver for the start_date field.
func (r *taskResolver) StartDate(ctx context.Context, obj *model.Task) (*string, error) {
	panic(fmt.Errorf("not implemented: StartDate - start_date"))
}

// DueDate is the resolver for the due_date field.
func (r *taskResolver) DueDate(ctx context.Context, obj *model.Task) (*string, error) {
	panic(fmt.Errorf("not implemented: DueDate - due_date"))
}

// Metadata is the resolver for the metadata field.
func (r *taskResolver) Metadata(ctx context.Context, obj *model.Task) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// CreatedAt is the resolver for the created_at field.
func (r *taskResolver) CreatedAt(ctx context.Context, obj *model.Task) (string, error) {
	panic(fmt.Errorf("not implemented: CreatedAt - created_at"))
}

// UpdatedAt is the resolver for the updated_at field.
func (r *taskResolver) UpdatedAt(ctx context.Context, obj *model.Task) (string, error) {
	return time.Unix(int64(obj.UpdatedAt.T), 0).Format(time.RFC3339), nil
}

// Comments is the resolver for the comments field.
func (r *taskResolver) Comments(ctx context.Context, obj *model.Task) ([]*model.Comment, error) {
	panic(fmt.Errorf("not implemented: Comments - comments"))
}

// Followers is the resolver for the followers field.
func (r *taskResolver) Followers(ctx context.Context, obj *model.Task) ([]*model.Follow, error) {
	panic(fmt.Errorf("not implemented: Followers - followers"))
}

// Reactions is the resolver for the reactions field.
func (r *taskResolver) Reactions(ctx context.Context, obj *model.Task) ([]*model.Reaction, error) {
	panic(fmt.Errorf("not implemented: Reactions - reactions"))
}

// Task returns TaskResolver implementation.
func (r *Resolver) Task() TaskResolver { return &taskResolver{r} }

type taskResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *taskResolver) Owner(ctx context.Context, obj *model.Task) (*model.User, error) {
	panic(fmt.Errorf("not implemented: Owner - owner"))
}
func (r *taskResolver) CreatedBy(ctx context.Context, obj *model.Task) (*model.User, error) {
	return &model.User{
		Model: model.Model{
			ID: obj.CreatedBy,
		},
	}, nil
}
func (r *taskResolver) UpdatedBy(ctx context.Context, obj *model.Task) (*model.User, error) {
	return &model.User{
		Model: model.Model{
			ID: obj.UpdatedBy,
		},
	}, nil
}
