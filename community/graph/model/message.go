package model

import (
	"context"
	"time"

	"github.com/dailytravel/x/community/db"
	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
	"github.com/typesense/typesense-go/typesense/api/pointer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Message struct {
	Model        `bson:",inline"`
	Parent       primitive.ObjectID  `json:"parent,omitempty" bson:"parent,omitempty"`
	Conversation primitive.ObjectID  `json:"conversation" bson:"conversation"`
	UID          primitive.ObjectID  `json:"uid" bson:"uid"`
	Locale       string              `json:"locale" bson:"locale"`
	Type         string              `json:"type" bson:"type"`
	Subject      *string             `json:"subject,omitempty" bson:"subject,omitempty"`
	Body         *string             `json:"body,omitempty" bson:"body,omitempty"`
	Status       string              `json:"status" bson:"status"`
	ScheduleAt   primitive.Timestamp `json:"schedule_at,omitempty" bson:"schedule_at,omitempty"`
}

func (i *Message) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Message
	return bson.Marshal((*t)(i))
}

func (i *Message) Collection() string {
	return "messages"
}

func (i *Message) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "parent", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "conversation", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "deleted_at", Value: 1}}, Options: options.Index()},
	}
}

func (i *Message) Schema() interface{} {
	return &api.CollectionSchema{
		Name: i.Collection(),
		Fields: []api.Field{
			{Name: "uid", Type: "string"},
			{Name: "parent", Type: "string"},
			{Name: "conversation", Type: "string"},
			{Name: "locale", Type: "string", Facet: pointer.True()},
			{Name: "type", Type: "string", Facet: pointer.True()},
			{Name: "subject", Type: "string"},
			{Name: "body", Type: "object"},
			{Name: "status", Type: "string", Facet: pointer.True()},
			{Name: "created_at", Type: "int64", Facet: pointer.True()},
			{Name: "updated_at", Type: "int64", Facet: pointer.True()},
			{Name: "recipients", Type: "object[]", Optional: pointer.True()},
		},
		DefaultSortingField: pointer.String("created_at"),
		EnableNestedFields:  pointer.True(),
	}
}

func (i *Message) Document() map[string]interface{} {
	document := map[string]interface{}{
		"id":           i.ID,
		"parent":       i.Parent,
		"uid":          i.UID,
		"conversation": i.Conversation,
		"locale":       i.Locale,
		"type":         i.Type,
		"subject":      i.Subject,
		"body":         i.Body,
		"status":       i.Status,
		// "recipients":   i.Recipients,
		"created_at": time.Unix(int64(i.CreatedAt.T), 0).Format(time.RFC3339),
		"updated_at": time.Unix(int64(i.UpdatedAt.T), 0).Format(time.RFC3339),
	}

	return document
}

func (i *Message) Insert(collection typesense.CollectionInterface) error {
	document := i.Document()

	if _, err := collection.Retrieve(); err != nil {
		// Create collection
		if _, err := db.Client.Collections().Create(i.Schema().(*api.CollectionSchema)); err != nil {
			return err
		}
	}
	if _, err := collection.Documents().Create(document); err != nil {
		return err
	}

	return nil
}

func (i *Message) Update(collection typesense.CollectionInterface, documentKey primitive.M, updatedFields primitive.M, removedFields primitive.A) error {
	documentID := documentKey["_id"].(primitive.ObjectID).Hex()

	// Create a map to hold the updated fields
	updatePayload := make(map[string]interface{})

	for field, value := range updatedFields {
		switch field {
		case "created_at", "updated_at", "schedule_at":
			timestamp := value.(primitive.Timestamp)
			updatePayload[field] = timestamp.T
		default:
			updatePayload[field] = value
		}
	}

	for _, field := range removedFields {
		updatePayload[field.(string)] = nil
	}

	if _, err := collection.Document(documentID).Update(updatePayload); err != nil {
		var item *Message
		if err := db.Database.Collection(i.Collection()).FindOne(context.Background(), documentKey).Decode(&item); err != nil {
			return err
		}

		if err := i.Insert(collection); err != nil {
			return err
		}
	}

	return nil
}

func (i *Message) Delete(collection typesense.CollectionInterface, documentKey primitive.M) error {
	id := documentKey["_id"].(primitive.ObjectID).Hex()

	// Delete document from Typesense
	if _, err := collection.Document(id).Delete(); err != nil {
		return err
	}

	return nil
}
