package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"
	"time"

	"github.com/dailytravel/x/marketing/graph/model"
)

// ID is the resolver for the id field.
func (r *linkResolver) ID(ctx context.Context, obj *model.Link) (string, error) {
	return obj.ID.Hex(), nil
}

// Metadata is the resolver for the metadata field.
func (r *linkResolver) Metadata(ctx context.Context, obj *model.Link) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// Created is the resolver for the created field.
func (r *linkResolver) Created(ctx context.Context, obj *model.Link) (string, error) {
	return time.Unix(int64(obj.Created.T), 0).Format(time.RFC3339), nil
}

// Updated is the resolver for the updated field.
func (r *linkResolver) Updated(ctx context.Context, obj *model.Link) (string, error) {
	return time.Unix(int64(obj.Updated.T), 0).Format(time.RFC3339), nil
}

// UID is the resolver for the uid field.
func (r *linkResolver) UID(ctx context.Context, obj *model.Link) (string, error) {
	return obj.UID.Hex(), nil
}

// Tags is the resolver for the tags field.
func (r *linkResolver) Tags(ctx context.Context, obj *model.Link) ([]*string, error) {
	panic(fmt.Errorf("not implemented: Tags - tags"))
}

// CreateLink is the resolver for the createLink field.
func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	panic(fmt.Errorf("not implemented: CreateLink - createLink"))
}

// UpdateLink is the resolver for the updateLink field.
func (r *mutationResolver) UpdateLink(ctx context.Context, id string, input model.UpdateLink) (*model.Link, error) {
	panic(fmt.Errorf("not implemented: UpdateLink - updateLink"))
}

// DeleteLink is the resolver for the deleteLink field.
func (r *mutationResolver) DeleteLink(ctx context.Context, id string) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: DeleteLink - deleteLink"))
}

// DeleteLinks is the resolver for the deleteLinks field.
func (r *mutationResolver) DeleteLinks(ctx context.Context, ids []string) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: DeleteLinks - deleteLinks"))
}

// Links is the resolver for the links field.
func (r *queryResolver) Links(ctx context.Context, filter map[string]interface{}, project map[string]interface{}, sort map[string]interface{}, collation map[string]interface{}, limit *int, skip *int) (*model.Links, error) {
	panic(fmt.Errorf("not implemented: Links - links"))
}

// Link is the resolver for the link field.
func (r *queryResolver) Link(ctx context.Context, id string) (*model.Link, error) {
	panic(fmt.Errorf("not implemented: Link - link"))
}

// Link returns LinkResolver implementation.
func (r *Resolver) Link() LinkResolver { return &linkResolver{r} }

type linkResolver struct{ *Resolver }
