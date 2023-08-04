package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"

	"github.com/dailytravel/x/service/graph/model"
)

// FindBoardByID is the resolver for the findBoardByID field.
func (r *entityResolver) FindBoardByID(ctx context.Context, id string) (*model.Board, error) {
	panic(fmt.Errorf("not implemented: FindBoardByID - findBoardByID"))
}

// Entity returns EntityResolver implementation.
func (r *Resolver) Entity() EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
