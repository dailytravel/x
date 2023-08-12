package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"
	"time"

	"github.com/dailytravel/x/account/graph/model"
)

// ID is the resolver for the id field.
func (r *clientResolver) ID(ctx context.Context, obj *model.Client) (string, error) {
	return obj.ID.Hex(), nil
}

// User is the resolver for the user field.
func (r *clientResolver) User(ctx context.Context, obj *model.Client) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Metadata is the resolver for the metadata field.
func (r *clientResolver) Metadata(ctx context.Context, obj *model.Client) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// LastUsed is the resolver for the last_used field.
func (r *clientResolver) LastUsed(ctx context.Context, obj *model.Client) (*string, error) {
	if obj.LastUsed.IsZero() {
		return nil, nil
	}

	lastUsed := time.Unix(int64(obj.LastUsed.T), 0).Format(time.RFC3339)
	return &lastUsed, nil
}

// CreatedAt is the resolver for the created_at field.
func (r *clientResolver) CreatedAt(ctx context.Context, obj *model.Client) (string, error) {
	return time.Unix(int64(obj.CreatedAt.T), 0).Format(time.RFC3339), nil
}

// UpdatedAt is the resolver for the updated_at field.
func (r *clientResolver) UpdatedAt(ctx context.Context, obj *model.Client) (string, error) {
	return time.Unix(int64(obj.UpdatedAt.T), 0).Format(time.RFC3339), nil
}

// ExpiresAt is the resolver for the expires_at field.
func (r *clientResolver) ExpiresAt(ctx context.Context, obj *model.Client) (*string, error) {
	if obj.ExpiresAt.IsZero() {
		return nil, nil
	}

	expiresAt := time.Unix(int64(obj.ExpiresAt.T), 0).Format(time.RFC3339)
	return &expiresAt, nil
}

// CreateClient is the resolver for the createClient field.
func (r *mutationResolver) CreateClient(ctx context.Context, input model.NewClient) (*model.Client, error) {
	panic(fmt.Errorf("not implemented: CreateClient - createClient"))
}

// UpdateClient is the resolver for the updateClient field.
func (r *mutationResolver) UpdateClient(ctx context.Context, id string, input model.UpdateClient) (*model.Client, error) {
	panic(fmt.Errorf("not implemented: UpdateClient - updateClient"))
}

// DeleteClient is the resolver for the deleteClient field.
func (r *mutationResolver) DeleteClient(ctx context.Context, id string) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: DeleteClient - deleteClient"))
}

// DeleteClients is the resolver for the deleteClients field.
func (r *mutationResolver) DeleteClients(ctx context.Context, ids []string) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: DeleteClients - deleteClients"))
}

// Client is the resolver for the client field.
func (r *queryResolver) Client(ctx context.Context, id string) (*model.Client, error) {
	panic(fmt.Errorf("not implemented: Client - client"))
}

// Clients is the resolver for the clients field.
func (r *queryResolver) Clients(ctx context.Context, args map[string]interface{}) (*model.Clients, error) {
	panic(fmt.Errorf("not implemented: Clients - clients"))
}

// Client returns ClientResolver implementation.
func (r *Resolver) Client() ClientResolver { return &clientResolver{r} }

type clientResolver struct{ *Resolver }
