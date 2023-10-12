package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"

	"github.com/dailytravel/x/account/graph/model"
)

// ID is the resolver for the id field.
func (r *tokenResolver) ID(ctx context.Context, obj *model.Token) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// User is the resolver for the user field.
func (r *tokenResolver) User(ctx context.Context, obj *model.Token) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// ExpiresAt is the resolver for the expires_at field.
func (r *tokenResolver) ExpiresAt(ctx context.Context, obj *model.Token) (string, error) {
	panic(fmt.Errorf("not implemented: ExpiresAt - expires_at"))
}

// CreatedAt is the resolver for the created_at field.
func (r *tokenResolver) CreatedAt(ctx context.Context, obj *model.Token) (string, error) {
	panic(fmt.Errorf("not implemented: CreatedAt - created_at"))
}

// Token returns TokenResolver implementation.
func (r *Resolver) Token() TokenResolver { return &tokenResolver{r} }

type tokenResolver struct{ *Resolver }
