package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"

	"github.com/dailytravel/x/configuration/graph/model"
)

// FindPlaceByID is the resolver for the findPlaceByID field.
func (r *entityResolver) FindPlaceByID(ctx context.Context, id string) (*model.Place, error) {
	panic(fmt.Errorf("not implemented: FindPlaceByID - findPlaceByID"))
}

// FindTermByID is the resolver for the findTermByID field.
func (r *entityResolver) FindTermByID(ctx context.Context, id string) (*model.Term, error) {
	panic(fmt.Errorf("not implemented: FindTermByID - findTermByID"))
}

// FindUserByID is the resolver for the findUserByID field.
func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	return &model.User{
		ID: id,
	}, nil
}

// Entity returns EntityResolver implementation.
func (r *Resolver) Entity() EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
