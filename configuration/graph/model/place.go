package model

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dailytravel/x/configuration/pkg/database"
	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
	"github.com/typesense/typesense-go/typesense/api/pointer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Place struct {
	Model       `bson:",inline"`
	Parent      *primitive.ObjectID `json:"parent,omitempty" bson:"parent,omitempty"`
	Locale      string              `json:"locale" bson:"locale"`
	Type        string              `json:"type" bson:"type"`
	Slug        string              `json:"slug" bson:"slug"`
	Name        primitive.M         `json:"name" bson:"name"`
	Description primitive.M         `json:"description,omitempty" bson:"description,omitempty"`
	Location    *Location           `json:"location,omitempty" bson:"location,omitempty"`
	Reviewable  *bool               `json:"reviewable" bson:"reviewable"`
	Popular     *bool               `json:"popular" bson:"popular"`
	Order       *int                `json:"order,omitempty" bson:"order,omitempty"`
	Status      string              `json:"status" bson:"status"`
}

func (Place) IsEntity() {}

func (i *Place) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Place
	return bson.Marshal((*t)(i))
}

func (i *Place) Collection() string {
	return "places"
}

func (i *Place) Index() []mongo.IndexModel {
	locales := strings.Split(os.Getenv("LOCALES"), ",") // Fetch the list of locales

	// Dynamically construct the keys for the text index based on the locales
	var keys bson.D
	for _, locale := range locales {
		key := fmt.Sprintf("name.%s", locale)
		keys = append(keys, bson.E{Key: key, Value: "text"})
	}

	// Construct weights for the text index, if needed
	var weights bson.D
	for _, locale := range locales {
		key := fmt.Sprintf("name.%s", locale)
		weights = append(weights, bson.E{Key: key, Value: 1})
	}

	// Return all indices, including the text index and base indices for other fields
	return []mongo.IndexModel{
		// Text index based on name fields for different locales
		{
			Keys:    keys,
			Options: options.Index().SetWeights(weights),
		},
		// Unique index for "name.en", with a filter for non-null values
		{
			Keys:    bson.D{{Key: "name.en", Value: 1}},
			Options: options.Index().SetUnique(true).SetPartialFilterExpression(bson.M{"name.en": bson.M{"$exists": true}}),
		},
		// Additional basic indices for frequently queried fields
		{Keys: bson.D{{Key: "slug", Value: 1}}, Options: options.Index().SetUnique(true).SetSparse(true)},
		{Keys: bson.D{{Key: "parent", Value: 1}}},
		{Keys: bson.D{{Key: "type", Value: 1}}},
		{Keys: bson.D{{Key: "locale", Value: 1}}},
		{Keys: bson.D{{Key: "reviewable", Value: 1}}},
		{Keys: bson.D{{Key: "popular", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
	}
}

func (i *Place) Schema() interface{} {
	schema := &api.CollectionSchema{
		Name: i.Collection(),
		Fields: []api.Field{
			{Name: "type", Type: "string", Facet: pointer.True()},
			{Name: "locale", Type: "string", Facet: pointer.True()},
			{Name: "slug", Type: "string"},
			{Name: "name", Type: "object"},
			{Name: "description", Type: "object", Optional: pointer.True()},
			{Name: "location", Type: "geopoint", Optional: pointer.True()},
			{Name: "reviewable", Type: "bool", Facet: pointer.True()},
			{Name: "popular", Type: "bool", Facet: pointer.True()},
			{Name: "status", Type: "string", Facet: pointer.True()},
			{Name: "created", Type: "string"},
			{Name: "updated", Type: "string"},
		},
		EnableNestedFields: pointer.True(),
	}

	return schema
}

func (i *Place) Document() map[string]interface{} {
	document := map[string]interface{}{
		"id":          i.ID, // Convert ID to string if it's not already
		"parent":      i.Parent,
		"locale":      i.Locale,
		"slug":        i.Slug,
		"type":        i.Type,
		"name":        i.Name,
		"description": i.Description,
		"status":      i.Status,
		"created":     time.Unix(int64(i.Created.T), 0).Format(time.RFC3339),
		"updated":     time.Unix(int64(i.Updated.T), 0).Format(time.RFC3339),
	}

	return document
}

func (i *Place) Insert(collection typesense.CollectionInterface) error {
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

func (i *Place) Update(collection typesense.CollectionInterface, documentKey primitive.M, updatedFields primitive.M, removedFields primitive.A) error {
	documentID := documentKey["_id"].(primitive.ObjectID).Hex()

	// Create a map to hold the updated fields
	updatePayload := make(map[string]interface{})

	for field, value := range updatedFields {
		switch field {
		case "created ", "updated ":
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
		var item *Place
		if err := database.Database.Collection(i.Collection()).FindOne(context.Background(), documentKey).Decode(&item); err != nil {
			return err
		}

		if err := i.Insert(collection); err != nil {
			return err
		}
	}

	return nil
}

func (i *Place) Delete(collection typesense.CollectionInterface, documentKey primitive.M) error {
	id := documentKey["_id"].(primitive.ObjectID).Hex()

	// Delete document from Typesense
	if _, err := collection.Document(id).Delete(); err != nil {
		return err
	}

	return nil
}
