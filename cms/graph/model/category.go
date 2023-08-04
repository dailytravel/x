package model

import (
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

type Category struct {
	Model       `bson:",inline"`
	Parent      primitive.ObjectID `json:"parent,omitempty" bson:"parent,omitempty"`
	Locale      string             `json:"locale" bson:"locale"`
	Name        primitive.M        `json:"name" bson:"name"`
	Description primitive.M        `json:"description,omitempty" bson:"description,omitempty"`
	Taxonomy    string             `json:"taxonomy" bson:"taxonomy"`
	Slug        *string            `json:"slug" bson:"slug"`
	Order       int                `json:"order" bson:"order"`
	Count       int                `json:"count" bson:"count"`
}

func (Category) IsEntity() {}

func (i *Category) Collection() string {
	return "categories"
}

func (i *Category) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Category
	return bson.Marshal((*t)(i))
}

func (i *Category) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "slug", Value: 1}}, Options: options.Index().SetUnique(true).SetSparse(true)},
		{Keys: bson.D{{Key: "taxonomy", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_by", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_by", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
	}
}

func (i *Category) Schema() interface{} {
	schema := &api.CollectionSchema{
		Name: i.Collection(),
		Fields: []api.Field{
			{Name: "slug", Type: "string"},
			{Name: "parent", Type: "string", Optional: pointer.True()},
			{Name: "taxonomy", Type: "string", Facet: pointer.True()},
			{Name: "locale", Type: "string", Facet: pointer.True()},
			{Name: "name", Type: "object"},
			{Name: "description", Type: "object"},
			{Name: "order", Type: "int32", Optional: pointer.True()},
			{Name: "count", Type: "int32", Optional: pointer.True()},
			{Name: "created_at", Type: "int32"},
			{Name: "updated_at", Type: "int32"},
		},
		DefaultSortingField: pointer.String("created_at"),
		EnableNestedFields:  pointer.True(),
	}

	return schema
}

func (i *Category) Document() map[string]interface{} {
	document := map[string]interface{}{
		"id":          i.ID,
		"parent":      i.Parent,
		"locale":      i.Locale,
		"name":        i.Name,
		"description": i.Description,
		"taxonomy":    i.Taxonomy,
		"slug":        i.Slug,
		"order":       i.Order,
		"count":       i.Count,
		"created_at":  i.CreatedAt.T,
		"updated_at":  i.UpdatedAt.T,
	}

	return document
}

func (i *Category) Insert(collection typesense.CollectionInterface) error {
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

func (i *Category) Update(collection typesense.CollectionInterface, documentKey primitive.M, updatedFields primitive.M, removedFields primitive.A) error {
	documentID := documentKey["_id"].(primitive.ObjectID).Hex()

	// Check if 'deleted_at' field is in updatedFields and its value is of type primitive.Timestamp
	_, deletedAtExist := updatedFields["deleted_at"].(primitive.Timestamp)
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
		case "created_at", "updated_at", "last_activity":
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

func (i *Category) Delete(collection typesense.CollectionInterface, documentKey primitive.M) error {
	id := documentKey["_id"].(primitive.ObjectID).Hex()

	// Delete document from Typesense
	if _, err := collection.Document(id).Delete(); err != nil {
		return err
	}

	return nil
}
