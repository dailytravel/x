package model

import (
	"context"
	"time"

	"github.com/dailytravel/x/cms/db"
	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
	"github.com/typesense/typesense-go/typesense/api/pointer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type File struct {
	Model       `bson:",inline"`
	Owner       primitive.ObjectID `bson:"owner" json:"owner"`
	Locale      string             `json:"locale" bson:"locale"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Type        string             `json:"type" bson:"type"`
	Size        int64              `json:"size" bson:"size"`
	Provider    string             `json:"provider" bson:"provider"`
	URL         string             `json:"url" bson:"url"`
	Starred     bool               `json:"starred" bson:"starred"`
	Status      string             `json:"status" bson:"status"`
}

func (File) IsEntity() {}

func (i *File) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t File
	return bson.Marshal((*t)(i))
}

func (i *File) Collection() string {
	return "files"
}

func (i *File) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "name", Value: "text"}}, Options: options.Index().SetWeights(bson.D{{Key: "name", Value: 1}})},
		{Keys: bson.D{{Key: "type", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "locale", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "owner", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "starred", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
	}
}

func (i *File) Schema() interface{} {
	return &api.CollectionSchema{
		Name: i.Collection(),
		Fields: []api.Field{
			{Name: "owner", Type: "string", Facet: pointer.True()},
			{Name: "type", Type: "string", Facet: pointer.True(), Optional: pointer.True()},
			{Name: "name", Type: "string"},
			{Name: "description", Type: "string", Optional: pointer.True()},
			{Name: "url", Type: "string"},
			{Name: "size", Type: "int32"},
			{Name: "starred", Type: "bool", Facet: pointer.True()},
			{Name: "status", Type: "string", Facet: pointer.True()},
			{Name: "created_at", Type: "int32"},
			{Name: "updated_at", Type: "int32"},
			{Name: "taxonomies", Type: "string[]", Optional: pointer.True()},
			{Name: "followers", Type: "string[]", Optional: pointer.True()},
		},
		DefaultSortingField: pointer.String("created_at"),
		EnableNestedFields:  pointer.True(),
	}
}

func (i *File) Document() map[string]interface{} {
	document := map[string]interface{}{
		"id":          i.ID,
		"owner":       i.Owner,
		"locale":      i.Locale,
		"name":        i.Name,
		"description": i.Description,
		"type":        i.Type,
		"url":         i.URL,
		"size":        i.Size,
		"starred":     i.Starred,
		"status":      i.Status,
		"created_at":  i.CreatedAt.T,
		"updated_at":  i.UpdatedAt.T,
	}

	// if i.Followers() != nil {
	// 	document["followers"] = i.Followers()
	// }

	// if i.Taxonomies() != nil {
	// 	document["taxonomies"] = i.Taxonomies()
	// }

	return document
}

func (i *File) Insert(collection typesense.CollectionInterface) error {
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

func (i *File) Update(collection typesense.CollectionInterface, documentKey primitive.M, updatedFields primitive.M, removedFields primitive.A) error {
	documentID := documentKey["_id"].(primitive.ObjectID).Hex()

	// Create a map to hold the updated fields
	updatePayload := make(map[string]interface{})

	for field, value := range updatedFields {
		switch field {
		case "created_at", "updated_at", "last_activity":
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
		var item *File
		if err := db.Database.Collection(i.Collection()).FindOne(context.Background(), documentKey).Decode(&item); err != nil {
			return err
		}

		if err := i.Insert(collection); err != nil {
			return err
		}
	}

	return nil
}

func (i *File) Delete(collection typesense.CollectionInterface, documentKey primitive.M) error {
	id := documentKey["_id"].(primitive.ObjectID).Hex()

	// Delete document from Typesense
	if _, err := collection.Document(id).Delete(); err != nil {
		return err
	}

	return nil
}