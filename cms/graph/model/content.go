package model

import (
	"context"
	"time"

	"github.com/dailytravel/x/cms/pkg/database"
	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
	"github.com/typesense/typesense-go/typesense/api/pointer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Content struct {
	Model       `bson:",inline"`
	UID         primitive.ObjectID    `json:"uid" bson:"uid"`
	Parent      *primitive.ObjectID   `json:"parent,omitempty" bson:"parent,omitempty"`
	Slug        *string               `json:"slug" bson:"slug"`
	Locale      string                `json:"locale" bson:"locale"`
	Type        string                `json:"type" bson:"type"`
	Title       primitive.M           `json:"title" bson:"title"`
	Summary     primitive.M           `json:"summary" bson:"summary"`
	Body        primitive.M           `json:"body" bson:"body"`
	Terms       []*primitive.ObjectID `json:"terms,omitempty" bson:"terms,omitempty"`
	Status      string                `json:"status" bson:"status"`
	Commentable bool                  `json:"commentable" bson:"commentable"`
	Metadata    primitive.M           `json:"metadata" bson:"metadata"`
}

func (Content) IsEntity() {}

func (i *Content) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Content
	return bson.Marshal((*t)(i))
}

func (i *Content) Collection() string {
	return "contents"
}

func (i *Content) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "parent", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "type", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "slug", Value: 1}}, Options: options.Index().SetUnique(true).SetSparse(true)},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created ", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated ", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "deleted ", Value: 1}}, Options: options.Index()},
	}
}

func (i *Content) Schema() interface{} {
	schema := &api.CollectionSchema{
		Name: i.Collection(),
		Fields: []api.Field{
			{Name: "parent", Type: "string", Facet: pointer.True()},
			{Name: "type", Type: "string", Facet: pointer.True()},
			{Name: "locale", Type: "string", Facet: pointer.True()},
			{Name: "slug", Type: "string"},
			{Name: "title", Type: "object[]"},
			{Name: "summary", Type: "object[]", Optional: pointer.True()},
			{Name: "body", Type: "object[]", Optional: pointer.True()},
			{Name: "status", Type: "string", Facet: pointer.True()},
			{Name: "created ", Type: "string"},
			{Name: "updated ", Type: "string"},
		},
		DefaultSortingField: pointer.String("created "),
		EnableNestedFields:  pointer.True(),
	}

	return schema
}

func (i *Content) Document() map[string]interface{} {
	document := map[string]interface{}{
		"id":       i.ID,
		"uid":      i.UID,
		"locale":   i.Locale,
		"type":     i.Type,
		"title":    i.Title,
		"summary":  i.Summary,
		"body":     i.Body,
		"slug":     i.Slug,
		"created ": time.Unix(int64(i.Created.T), 0).Format(time.RFC3339),
		"updated ": time.Unix(int64(i.Updated.T), 0).Format(time.RFC3339),
	}

	if i.Parent != nil {
		document["parent"] = i.Parent
	}

	return document
}

func (i *Content) Insert(collection typesense.CollectionInterface) error {
	document := i.Document()

	if _, err := collection.Retrieve(); err != nil {
		// Create collection
		if _, err := database.Client.Collections().Create(i.Schema().(*api.CollectionSchema)); err != nil {
			return err
		}
	}

	if _, err := collection.Documents().Create(document); err != nil {
		return err
	}

	return nil
}

func (i *Content) Update(collection typesense.CollectionInterface, documentKey primitive.M, updatedFields primitive.M, removedFields primitive.A) error {
	documentID := documentKey["_id"].(primitive.ObjectID).Hex()

	// Create a map to hold the updated fields
	updatePayload := make(map[string]interface{})

	for field, value := range updatedFields {
		switch field {
		case "created ", "updated ", "last_activity":
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
		var item *Content
		if err := database.Database.Collection(i.Collection()).FindOne(context.Background(), documentKey).Decode(&item); err != nil {
			return err
		}

		if err := i.Insert(collection); err != nil {
			return err
		}
	}

	return nil
}

func (i *Content) Delete(collection typesense.CollectionInterface, documentKey primitive.M) error {
	id := documentKey["_id"].(primitive.ObjectID).Hex()

	// Delete document from Typesense
	if _, err := collection.Document(id).Delete(); err != nil {
		return err
	}

	return nil
}
