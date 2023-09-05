package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"

	"github.com/dailytravel/x/account/graph/model"
)

// User is the resolver for the user field.
func (r *shareResolver) User(ctx context.Context, obj *model.Share) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Share returns ShareResolver implementation.
func (r *Resolver) Share() ShareResolver { return &shareResolver{r} }

type shareResolver struct{ *Resolver }
