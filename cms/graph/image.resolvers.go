package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"

	"github.com/dailytravel/x/cms/graph/model"
	"github.com/dailytravel/x/cms/pkg/auth"
	"go.mongodb.org/mongo-driver/bson"
)

// ID is the resolver for the id field.
func (r *imageResolver) ID(ctx context.Context, obj *model.Image) (string, error) {
	return obj.ID.Hex(), nil
}

// Object is the resolver for the object field.
func (r *imageResolver) Object(ctx context.Context, obj *model.Image) (map[string]interface{}, error) {
	var object map[string]interface{}
	if err := r.db.Collection(obj.Object.Collection).FindOne(ctx, bson.M{"_id": obj.Object.ID}).Decode(&object); err != nil {
		return nil, err
	}

	return object, nil
}

// Title is the resolver for the title field.
func (r *imageResolver) Title(ctx context.Context, obj *model.Image) (*string, error) {
	// Get the locale from the context
	locale := auth.Locale(ctx)
	if locale == nil {
		locale = &obj.Locale
	}

	// Try to retrieve the title for the requested locale
	if title, ok := obj.Title[*locale].(string); ok {
		return &title, nil
	}

	return obj.Title[obj.Locale].(*string), nil
}

// Caption is the resolver for the caption field.
func (r *imageResolver) Caption(ctx context.Context, obj *model.Image) (*string, error) {
	// Get the locale from the context
	locale := auth.Locale(ctx)
	if locale == nil {
		locale = &obj.Locale
	}

	// Try to retrieve the caption for the requested locale
	if caption, ok := obj.Caption[*locale].(string); ok {
		return &caption, nil
	}

	return obj.Caption[obj.Locale].(*string), nil
}

// Metadata is the resolver for the metadata field.
func (r *imageResolver) Metadata(ctx context.Context, obj *model.Image) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// Image returns ImageResolver implementation.
func (r *Resolver) Image() ImageResolver { return &imageResolver{r} }

type imageResolver struct{ *Resolver }
