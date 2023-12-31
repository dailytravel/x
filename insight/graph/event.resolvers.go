package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"
	"time"

	"github.com/dailytravel/x/insight/graph/model"
)

// ID is the resolver for the id field.
func (r *eventResolver) ID(ctx context.Context, obj *model.Event) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// Campaign is the resolver for the campaign field.
func (r *eventResolver) Campaign(ctx context.Context, obj *model.Event) (string, error) {
	panic(fmt.Errorf("not implemented: Campaign - campaign"))
}

// Status is the resolver for the status field.
func (r *eventResolver) Status(ctx context.Context, obj *model.Event) (model.MailStatus, error) {
	panic(fmt.Errorf("not implemented: Status - status"))
}

// Timestamp is the resolver for the timestamp field.
func (r *eventResolver) Timestamp(ctx context.Context, obj *model.Event) (string, error) {
	return time.Unix(int64(obj.Timestamp.T), 0).Format(time.RFC3339), nil
}

// Metadata is the resolver for the metadata field.
func (r *eventResolver) Metadata(ctx context.Context, obj *model.Event) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// Event returns EventResolver implementation.
func (r *Resolver) Event() EventResolver { return &eventResolver{r} }

type eventResolver struct{ *Resolver }
