package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dailytravel/x/community/graph/model"
	"github.com/dailytravel/x/community/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ID is the resolver for the id field.
func (r *conversationResolver) ID(ctx context.Context, obj *model.Conversation) (string, error) {
	return obj.ID.Hex(), nil
}

// Metadata is the resolver for the metadata field.
func (r *conversationResolver) Metadata(ctx context.Context, obj *model.Conversation) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// Message is the resolver for the message field.
func (r *conversationResolver) Message(ctx context.Context, obj *model.Conversation) (*model.Message, error) {
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

// CreatedAt is the resolver for the created_at field.
func (r *conversationResolver) CreatedAt(ctx context.Context, obj *model.Conversation) (string, error) {
	return time.Unix(int64(obj.CreatedAt.T), 0).Format(time.RFC3339), nil
}

// UpdatedAt is the resolver for the updated_at field.
func (r *conversationResolver) UpdatedAt(ctx context.Context, obj *model.Conversation) (string, error) {
	return time.Unix(int64(obj.UpdatedAt.T), 0).Format(time.RFC3339), nil
}

// Messages is the resolver for the messages field.
func (r *conversationResolver) Messages(ctx context.Context, obj *model.Conversation) ([]*model.Message, error) {
	var items []*model.Message

	filter := bson.M{"conversation": obj.ID}

	cur, err := r.db.Collection("messages").Find(ctx, filter, nil)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var item *model.Message
		if err := cur.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

// UID is the resolver for the uid field.
func (r *conversationResolver) UID(ctx context.Context, obj *model.Conversation) (string, error) {
	return obj.UID.Hex(), nil
}

// CreatedBy is the resolver for the created_by field.
func (r *conversationResolver) CreatedBy(ctx context.Context, obj *model.Conversation) (*string, error) {
	if obj.CreatedBy == nil {
		return nil, nil
	}

	createdBy := obj.CreatedBy.Hex()

	return &createdBy, nil
}

// UpdatedBy is the resolver for the updated_by field.
func (r *conversationResolver) UpdatedBy(ctx context.Context, obj *model.Conversation) (*string, error) {
	if obj.UpdatedBy == nil {
		return nil, nil
	}

	updatedBy := obj.UpdatedBy.Hex()

	return &updatedBy, nil
}

// Shares is the resolver for the shares field.
func (r *conversationResolver) Shares(ctx context.Context, obj *model.Conversation) ([]*model.Share, error) {
	panic(fmt.Errorf("not implemented: Shares - shares"))
}

// Comments is the resolver for the comments field.
func (r *conversationResolver) Comments(ctx context.Context, obj *model.Conversation) ([]*model.Comment, error) {
	var items []*model.Comment

	filter := bson.M{"commentable._id": obj.ID, "commentable.type": "conversations"}

	cur, err := r.db.Collection("comments").Find(ctx, filter, nil)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var item *model.Comment
		if err := cur.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

// CreateConversation is the resolver for the createConversation field.
func (r *mutationResolver) CreateConversation(ctx context.Context, input model.NewConversation) (*model.Conversation, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	item := &model.Conversation{
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

// UpdateConversation is the resolver for the updateConversation field.
func (r *mutationResolver) UpdateConversation(ctx context.Context, id string, input model.UpdateConversation) (*model.Conversation, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": _id}

	item := &model.Conversation{}
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

// LeaveConversation is the resolver for the leaveConversation field.
func (r *mutationResolver) LeaveConversation(ctx context.Context, id string) (*model.Conversation, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	// Convert the ID string to ObjectID
	convID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Define the filter to find the conversation by ID and the user leaving the conversation
	filter := bson.M{
		"_id":     convID,
		"members": uid,
	}

	// Define the update to remove the user from the members list
	update := bson.M{
		"$pull": bson.M{
			"members": uid,
		},
	}

	// Perform the update operation
	_, err = r.db.Collection("conversations").UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	// Retrieve the updated conversation
	var updatedConversation model.Conversation
	err = r.db.Collection("conversations").FindOne(ctx, bson.M{"_id": convID}).Decode(&updatedConversation)
	if err != nil {
		return nil, err
	}

	return &updatedConversation, nil
}

// DeleteConversation is the resolver for the deleteConversation field.
func (r *mutationResolver) DeleteConversation(ctx context.Context, id string) (map[string]interface{}, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	// Convert the ID string to ObjectID
	convID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Define the filter to find the conversation by ID and owner
	filter := bson.M{
		"_id":       convID,
		"createdBy": uid,
	}

	// Perform the deletion operation
	result, err := r.db.Collection("conversations").DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	// Check if the conversation was found and deleted
	if result.DeletedCount == 0 {
		return map[string]interface{}{"status": "error", "message": "Conversation not found or not owned by user"}, nil
	}

	// Return success response
	return map[string]interface{}{"status": "success", "message": "Conversation deleted"}, nil
}

// DeleteConversations is the resolver for the deleteConversations field.
func (r *mutationResolver) DeleteConversations(ctx context.Context, ids []string) (map[string]interface{}, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	// Convert the list of ID strings to ObjectIDs
	var objectIDs []primitive.ObjectID
	for _, id := range ids {
		convID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		objectIDs = append(objectIDs, convID)
	}

	// Define the filter to find conversations by IDs and owner
	filter := bson.M{
		"_id":       bson.M{"$in": objectIDs},
		"createdBy": uid,
	}

	// Perform the deletion operation
	result, err := r.db.Collection("conversations").DeleteMany(ctx, filter)
	if err != nil {
		return nil, err
	}

	// Return the number of conversations deleted
	return map[string]interface{}{"status": "success", "deletedCount": result.DeletedCount}, nil
}

// Conversation is the resolver for the conversation field.
func (r *queryResolver) Conversation(ctx context.Context, id string) (*model.Conversation, error) {
	var item *model.Conversation

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

// Conversations is the resolver for the conversations field.
func (r *queryResolver) Conversations(ctx context.Context, args map[string]interface{}) (*model.Conversations, error) {
	var items []*model.Conversation
	//find all items
	cur, err := r.db.Collection("conversations").Find(ctx, utils.Query(args), utils.Options(args))
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var item *model.Conversation
		if err := cur.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	//get total count
	count, err := r.db.Collection("conversations").CountDocuments(ctx, utils.Query(args), nil)
	if err != nil {
		return nil, err
	}

	return &model.Conversations{
		Count: int(count),
		Data:  items,
	}, nil
}

// Conversation returns ConversationResolver implementation.
func (r *Resolver) Conversation() ConversationResolver { return &conversationResolver{r} }

type conversationResolver struct{ *Resolver }
