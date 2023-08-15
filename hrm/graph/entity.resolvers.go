package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"

	"github.com/dailytravel/x/hrm/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FindOrganizationByID is the resolver for the findOrganizationByID field.
func (r *entityResolver) FindOrganizationByID(ctx context.Context, id string) (*model.Organization, error) {
	var item *model.Organization

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	if err := r.db.Collection(item.Collection()).FindOne(ctx, bson.M{"_id": _id}).Decode(&item); err != nil {
		return nil, err
	}

	return item, nil
}

// FindUserByID is the resolver for the findUserByID field.
func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: FindUserByID - findUserByID"))
}

// Entity returns EntityResolver implementation.
func (r *Resolver) Entity() EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *entityResolver) FindReactionByID(ctx context.Context, id string) (*model.Reaction, error) {
	var item *model.Reaction

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	if err := r.db.Collection(item.Collection()).FindOne(ctx, bson.M{"_id": _id}).Decode(&item); err != nil {
		return nil, err
	}

	return item, nil
}
