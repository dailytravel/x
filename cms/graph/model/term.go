package model

import (
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

type Term struct {
	Model       `bson:",inline"`
	Parent      primitive.ObjectID `json:"parent,omitempty" bson:"parent,omitempty"`
	Locale      string             `json:"locale" bson:"locale"`
	Name        primitive.M        `json:"name" bson:"name"`
	Description primitive.M        `json:"description,omitempty" bson:"description,omitempty"`
	Type        string             `json:"type" bson:"type"`
	Slug        *string            `json:"slug" bson:"slug"`
	Order       int                `json:"order" bson:"order"`
	Count       int                `json:"count" bson:"count"`
}

func (Term) IsEntity() {}

func (i *Term) Collection() string {
	return "terms"
}

func (i *Term) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Term
	return bson.Marshal((*t)(i))
}

func (i *Term) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "slug", Value: 1}}, Options: options.Index().SetUnique(true).SetSparse(true)},
		{Keys: bson.D{{Key: "type", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_by", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_by", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created ", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated ", Value: 1}}, Options: options.Index()},
	}
}

func (i *Term) Schema() interface{} {
	schema := &api.CollectionSchema{
		Name: i.Collection(),
		Fields: []api.Field{
			{Name: "slug", Type: "string"},
			{Name: "parent", Type: "string", Optional: pointer.True()},
			{Name: "type", Type: "string", Facet: pointer.True()},
			{Name: "locale", Type: "string", Facet: pointer.True()},
			{Name: "name", Type: "object"},
			{Name: "description", Type: "object"},
			{Name: "order", Type: "int32", Optional: pointer.True()},
			{Name: "count", Type: "int32", Optional: pointer.True()},
			{Name: "created ", Type: "int32"},
			{Name: "updated ", Type: "int32"},
		},
		DefaultSortingField: pointer.String("created "),
		EnableNestedFields:  pointer.True(),
	}

	return schema
}

func (i *Term) Document() map[string]interface{} {
	document := map[string]interface{}{
		"id":          i.ID,
		"parent":      i.Parent,
		"locale":      i.Locale,
		"name":        i.Name,
		"description": i.Description,
		"type":        i.Type,
		"slug":        i.Slug,
		"order":       i.Order,
		"count":       i.Count,
		"created ":    i.Created.T,
		"updated ":    i.Updated.T,
	}

	return document
}

func (i *Term) Insert(collection typesense.CollectionInterface) error {
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

func (i *Term) Update(collection typesense.CollectionInterface, documentKey primitive.M, updatedFields primitive.M, removedFields primitive.A) error {
	documentID := documentKey["_id"].(primitive.ObjectID).Hex()

	// Check if 'deleted ' field is in updatedFields and its value is of type primitive.Timestamp
	_, deletedAtExist := updatedFields["deleted "].(primitive.Timestamp)
	if deletedAtExist {
		if err := i.Delete(collection, documentKey); err != nil {
			return err
		}
		return nil
	}

	// Find the document
	_, err := collection.Document(documentID).Retrieve()
	if err != nil {
		// If the document doesn't exist, insert it into the collection
		if err := i.Insert(collection); err != nil {
			return err
		}
	}

	// Create a map to hold the updated fields
	updatePayload := make(map[string]interface{})

	// Loop through updatedFields
	for field, value := range updatedFields {
		switch field {
		case "created ", "updated ", "last_activity":
			if timestamp, ok := value.(primitive.Timestamp); ok {
				updatePayload[field] = timestamp.T
			}
		default:
			updatePayload[field] = value
		}
	}

	// Loop through removedFields
	for _, field := range removedFields {
		updatePayload[field.(string)] = nil
	}

	// Update the document with the updatePayload
	if _, err := collection.Document(documentID).Update(updatePayload); err != nil {
		return err
	}

	return nil
}

func (i *Term) Delete(collection typesense.CollectionInterface, documentKey primitive.M) error {
	id := documentKey["_id"].(primitive.ObjectID).Hex()

	// Delete document from Typesense
	if _, err := collection.Document(id).Delete(); err != nil {
		return err
	}

	return nil
}
