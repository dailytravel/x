package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"errors"
	"time"

	"github.com/dailytravel/x/community/graph/model"
	"github.com/dailytravel/x/community/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateRecipient is the resolver for the createRecipient field.
func (r *mutationResolver) CreateRecipient(ctx context.Context, input model.NewRecipient) (*model.Recipient, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	item := &model.Recipient{
		Model: model.Model{
			CreatedBy: uid,
			UpdatedBy: uid,
		},
	}

	// Insert the new organization
	if _, err := r.db.Collection(item.Collection()).InsertOne(ctx, item, nil); err != nil {
		return nil, err
	}

	return item, nil
}

// UpdateRecipient is the resolver for the updateRecipient field.
func (r *mutationResolver) UpdateRecipient(ctx context.Context, id string, input model.UpdateRecipient) (*model.Recipient, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": _id}

	item := &model.Recipient{}
	err = r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(&item)
	if err != nil {
		return nil, err
	}

	item.UpdatedBy = uid

	// Update the contact
	if _, err := r.db.Collection(item.Collection()).UpdateOne(ctx, filter, bson.M{"$set": item}, nil); err != nil {
		return nil, err
	}

	return item, nil
}

// DeleteRecipient is the resolver for the deleteRecipient field.
func (r *mutationResolver) DeleteRecipient(ctx context.Context, id string) (map[string]interface{}, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	// Convert the ID string to ObjectID
	recipientID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Define the filter to match the recipient by ID and the user who is deleting it
	filter := bson.M{
		"_id":    recipientID,
		"userId": uid,
	}

	// Perform the delete operation
	result, err := r.db.Collection("recipients").DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	if result.DeletedCount == 0 {
		return map[string]interface{}{"status": "failure", "message": "Recipient not found or not owned by user"}, nil
	}

	return map[string]interface{}{"status": "success", "message": "Recipient deleted successfully"}, nil
}

// DeleteRecipients is the resolver for the deleteRecipients field.
func (r *mutationResolver) DeleteRecipients(ctx context.Context, ids []string) (map[string]interface{}, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	// Convert the list of ID strings to ObjectIDs
	var recipientIDs []primitive.ObjectID
	for _, id := range ids {
		recipientID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		recipientIDs = append(recipientIDs, recipientID)
	}

	// Define the filter to match the recipients by IDs and the user who is deleting them
	filter := bson.M{
		"_id":    bson.M{"$in": recipientIDs},
		"userId": uid,
	}

	// Perform the delete operation
	result, err := r.db.Collection("recipients").DeleteMany(ctx, filter)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"status": "success", "deletedCount": result.DeletedCount}, nil
}

// Recipient is the resolver for the recipient field.
func (r *queryResolver) Recipient(ctx context.Context, id string) (*model.Recipient, error) {
	var item *model.Recipient

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": _id}

	if err := r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(&item); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return item, nil
}

// Recipients is the resolver for the recipients field.
func (r *queryResolver) Recipients(ctx context.Context, args map[string]interface{}) (*model.Recipients, error) {
	var items []*model.Recipient
	//find all items
	cur, err := r.db.Collection("recipients").Find(ctx, utils.Query(args), utils.Options(args))
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var item *model.Recipient
		if err := cur.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	//get total count
	count, err := r.db.Collection("recipients").CountDocuments(ctx, utils.Query(args), nil)
	if err != nil {
		return nil, err
	}

	return &model.Recipients{
		Count: int(count),
		Data:  items,
	}, nil
}

// ID is the resolver for the id field.
func (r *recipientResolver) ID(ctx context.Context, obj *model.Recipient) (string, error) {
	return obj.ID.Hex(), nil
}

// UID is the resolver for the uid field.
func (r *recipientResolver) UID(ctx context.Context, obj *model.Recipient) (string, error) {
	return obj.UID.Hex(), nil
}

// Message is the resolver for the message field.
func (r *recipientResolver) Message(ctx context.Context, obj *model.Recipient) (*model.Message, error) {
	var item *model.Message

	filter := bson.M{"_id": obj.Message}

	if err := r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(&item); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return item, nil
}

// ReadAt is the resolver for the read_at field.
func (r *recipientResolver) ReadAt(ctx context.Context, obj *model.Recipient) (*string, error) {
	if obj.ReadAt == nil {
		return nil, nil
	}

	readAt := time.Unix(int64(obj.ReadAt.T), 0).Format(time.RFC3339)
	return &readAt, nil
}

// CreatedAt is the resolver for the created_at field.
func (r *recipientResolver) CreatedAt(ctx context.Context, obj *model.Recipient) (string, error) {
	return time.Unix(int64(obj.CreatedAt.T), 0).Format(time.RFC3339), nil
}

// UpdatedAt is the resolver for the updated_at field.
func (r *recipientResolver) UpdatedAt(ctx context.Context, obj *model.Recipient) (string, error) {
	return time.Unix(int64(obj.UpdatedAt.T), 0).Format(time.RFC3339), nil
}

// Recipient returns RecipientResolver implementation.
func (r *Resolver) Recipient() RecipientResolver { return &recipientResolver{r} }

type recipientResolver struct{ *Resolver }
