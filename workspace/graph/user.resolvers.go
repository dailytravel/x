package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"

	"github.com/dailytravel/x/workspace/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Goals is the resolver for the goals field.
func (r *userResolver) Goals(ctx context.Context, obj *model.User) ([]*model.Goal, error) {
	var items []*model.Goal

	uid, err := primitive.ObjectIDFromHex(obj.ID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"uid": uid}
	//find all items
	cur, err := r.db.Collection("goals").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var item *model.Goal
		if err := cur.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

// Portfolios is the resolver for the portfolios field.
func (r *userResolver) Portfolios(ctx context.Context, obj *model.User) ([]*model.Portfolio, error) {
	var items []*model.Portfolio

	uid, err := primitive.ObjectIDFromHex(obj.ID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"uid": uid}
	//find all items
	cur, err := r.db.Collection("portfolios").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var item *model.Portfolio
		if err := cur.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

// Boards is the resolver for the boards field.
func (r *userResolver) Boards(ctx context.Context, obj *model.User) ([]*model.Board, error) {
	var items []*model.Board

	uid, err := primitive.ObjectIDFromHex(obj.ID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"uid": uid}
	//find all items
	cur, err := r.db.Collection("boards").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var item *model.Board
		if err := cur.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

// Lists is the resolver for the lists field.
func (r *userResolver) Lists(ctx context.Context, obj *model.User) ([]*model.List, error) {
	var items []*model.List

	uid, err := primitive.ObjectIDFromHex(obj.ID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"uid": uid}
	//find all items
	cur, err := r.db.Collection("lists").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var item *model.List
		if err := cur.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

// Tasks is the resolver for the tasks field.
func (r *userResolver) Tasks(ctx context.Context, obj *model.User) ([]*model.Task, error) {
	var items []*model.Task

	uid, err := primitive.ObjectIDFromHex(obj.ID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"uid": uid}
	//find all items
	cur, err := r.db.Collection("tasks").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var item *model.Task
		if err := cur.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
