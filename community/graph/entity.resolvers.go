package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"

	"github.com/dailytravel/x/community/graph/model"
)

// FindCommentByID is the resolver for the findCommentByID field.
func (r *entityResolver) FindCommentByID(ctx context.Context, id string) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented: FindCommentByID - findCommentByID"))
}

// FindContactByID is the resolver for the findContactByID field.
func (r *entityResolver) FindContactByID(ctx context.Context, id string) (*model.Contact, error) {
	panic(fmt.Errorf("not implemented: FindContactByID - findContactByID"))
}

// FindContentByID is the resolver for the findContentByID field.
func (r *entityResolver) FindContentByID(ctx context.Context, id string) (*model.Content, error) {
	panic(fmt.Errorf("not implemented: FindContentByID - findContentByID"))
}

// FindDealByID is the resolver for the findDealByID field.
func (r *entityResolver) FindDealByID(ctx context.Context, id string) (*model.Deal, error) {
	panic(fmt.Errorf("not implemented: FindDealByID - findDealByID"))
}

// FindExpenseByID is the resolver for the findExpenseByID field.
func (r *entityResolver) FindExpenseByID(ctx context.Context, id string) (*model.Expense, error) {
	panic(fmt.Errorf("not implemented: FindExpenseByID - findExpenseByID"))
}

// FindFileByID is the resolver for the findFileByID field.
func (r *entityResolver) FindFileByID(ctx context.Context, id string) (*model.File, error) {
	panic(fmt.Errorf("not implemented: FindFileByID - findFileByID"))
}

// FindFollowByID is the resolver for the findFollowByID field.
func (r *entityResolver) FindFollowByID(ctx context.Context, id string) (*model.Follow, error) {
	panic(fmt.Errorf("not implemented: FindFollowByID - findFollowByID"))
}

// FindQuoteByID is the resolver for the findQuoteByID field.
func (r *entityResolver) FindQuoteByID(ctx context.Context, id string) (*model.Quote, error) {
	panic(fmt.Errorf("not implemented: FindQuoteByID - findQuoteByID"))
}

// FindReactionByID is the resolver for the findReactionByID field.
func (r *entityResolver) FindReactionByID(ctx context.Context, id string) (*model.Reaction, error) {
	panic(fmt.Errorf("not implemented: FindReactionByID - findReactionByID"))
}

// FindTaskByID is the resolver for the findTaskByID field.
func (r *entityResolver) FindTaskByID(ctx context.Context, id string) (*model.Task, error) {
	panic(fmt.Errorf("not implemented: FindTaskByID - findTaskByID"))
}

// FindUserByID is the resolver for the findUserByID field.
func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: FindUserByID - findUserByID"))
}

// Entity returns EntityResolver implementation.
func (r *Resolver) Entity() EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
