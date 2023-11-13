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
)

// CreateWorkspace is the resolver for the createWorkspace field.
func (r *mutationResolver) CreateWorkspace(ctx context.Context, input model.NewWorkspace) (*model.Workspace, error) {
	panic(fmt.Errorf("not implemented: CreateWorkspace - createWorkspace"))
}

// UpdateWorkspace is the resolver for the updateWorkspace field.
func (r *mutationResolver) UpdateWorkspace(ctx context.Context, id string, input model.UpdateWorkspace) (*model.Workspace, error) {
	panic(fmt.Errorf("not implemented: UpdateWorkspace - updateWorkspace"))
}

// DeleteWorkspace is the resolver for the deleteWorkspace field.
func (r *mutationResolver) DeleteWorkspace(ctx context.Context, id string) (*model.Workspace, error) {
	panic(fmt.Errorf("not implemented: DeleteWorkspace - deleteWorkspace"))
}

// Workspace is the resolver for the workspace field.
func (r *queryResolver) Workspace(ctx context.Context, id string) (*model.Workspace, error) {
	var item *model.Workspace

	if err := r.db.Collection("workspaces").FindOne(ctx, bson.M{"_id": id}).Decode(&item); err != nil {
		return nil, err
	}

	return item, nil
}

// Workspaces is the resolver for the workspaces field.
func (r *queryResolver) Workspaces(ctx context.Context, stages map[string]interface{}) (*model.Workspaces, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	matchStage := bson.M{
		"$match": bson.M{
			"$and": []interface{}{
				bson.M{"status": bson.M{"$ne": "ARCHIVED"}},
				bson.M{"$or": []interface{}{
					bson.M{"uid": uid},
					bson.M{
						"members": bson.M{
							"$elemMatch": bson.M{
								"workspace": "$_id",
								"uid":       uid,
							},
						},
					},
				}},
			},
		},
	}

	pipeline := []bson.M{matchStage}

	for key, value := range stages {
		stage := bson.M{key: value}
		pipeline = append(pipeline, stage)
	}

	cursor, err := r.db.Collection("boards").Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var items []*model.Workspace

	if err := cursor.All(ctx, &items); err != nil {
		return nil, err
	}

	return &model.Workspaces{
		Count: len(items),
		Data:  items,
	}, nil
}

// ID is the resolver for the id field.
func (r *workspaceResolver) ID(ctx context.Context, obj *model.Workspace) (string, error) {
	return obj.ID.Hex(), nil
}

// Owner is the resolver for the owner field.
func (r *workspaceResolver) Owner(ctx context.Context, obj *model.Workspace) (*model.User, error) {
	var item *model.User

	if err := r.db.Collection("users").FindOne(ctx, bson.M{"_id": obj.UID}).Decode(&item); err != nil {
		return nil, err
	}

	return item, nil
}

// Members is the resolver for the members field.
func (r *workspaceResolver) Members(ctx context.Context, obj *model.Workspace) ([]*model.Member, error) {
	var items []*model.Member

	cursor, err := r.db.Collection("members").Find(ctx, bson.M{"workspace": obj.ID})
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &items); err != nil {
		return nil, err
	}

	return items, nil
}

// Metadata is the resolver for the metadata field.
func (r *workspaceResolver) Metadata(ctx context.Context, obj *model.Workspace) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// Created is the resolver for the created field.
func (r *workspaceResolver) Created(ctx context.Context, obj *model.Workspace) (string, error) {
	return time.Unix(int64(obj.Created.T), 0).Format(time.RFC3339), nil
}

// Updated is the resolver for the updated field.
func (r *workspaceResolver) Updated(ctx context.Context, obj *model.Workspace) (string, error) {
	return time.Unix(int64(obj.Updated.T), 0).Format(time.RFC3339), nil
}

// Workspace returns WorkspaceResolver implementation.
func (r *Resolver) Workspace() WorkspaceResolver { return &workspaceResolver{r} }

type workspaceResolver struct{ *Resolver }
