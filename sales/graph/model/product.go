package model

import (
	"context"
	"time"

	"github.com/dailytravel/x/sales/pkg/database"
	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
	"github.com/typesense/typesense-go/typesense/api/pointer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Product struct {
	Model       `bson:",inline"`
	UID         primitive.ObjectID    `json:"uid" bson:"uid"`
	Locale      string                `json:"locale" bson:"locale"`
	Type        string                `json:"type" bson:"type"`
	Name        primitive.M           `json:"name" bson:"name"`
	Description primitive.M           `json:"description" bson:"description"`
	Duration    int32                 `json:"duration" bson:"duration"`
	Notes       primitive.M           `json:"notes" bson:"notes"`
	Location    primitive.ObjectID    `json:"location" bson:"location"`
	Tips        primitive.M           `json:"tips" bson:"tips"`
	Highlights  primitive.M           `json:"highlights" bson:"highlights"`
	Expectation primitive.M           `json:"expectation" bson:"expectation"`
	Faqs        primitive.M           `json:"faqs" bson:"faqs"`
	Reviews     int32                 `json:"reviews" bson:"reviews"`
	Booked      int32                 `json:"booked" bson:"booked"`
	Rating      float32               `json:"rating" bson:"rating"`
	Reviewable  bool                  `json:"reviewable" bson:"reviewable"`
	Price       float64               `json:"price" bson:"price"`
	Discount    float64               `json:"discount,omitempty" bson:"discount,omitempty"`
	Currency    string                `json:"currency" bson:"currency"`
	Status      string                `json:"status" bson:"status"`
	Terms       []*primitive.ObjectID `json:"terms,omitempty" bson:"terms,omitempty"`
}

func (Product) IsEntity() {}

func (i *Product) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Product
	return bson.Marshal((*t)(i))
}

func (i *Product) Collection() string {
	return "products"
}

func (i *Product) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}}},
		{Keys: bson.D{{Key: "type", Value: 1}}},
		{Keys: bson.D{{Key: "duration", Value: 1}}},
		{Keys: bson.D{{Key: "rating", Value: 1}}},
		{Keys: bson.D{{Key: "booked", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "created ", Value: 1}}},
		{Keys: bson.D{{Key: "updated ", Value: 1}}},
	}
}

func (i *Product) Schema() interface{} {
	schema := &api.CollectionSchema{
		Name: i.Collection(),
		Fields: []api.Field{
			{Name: "type", Type: "string", Facet: pointer.True()},
			{Name: "locale", Type: "string", Facet: pointer.True()},
			{Name: "duration", Type: "int32", Facet: pointer.True()},
			{Name: "name", Type: "object[]"},
			{Name: "description", Type: "object[]", Optional: pointer.True()},
			{Name: "images", Type: "object[]", Optional: pointer.True()},
			{Name: "terms", Type: "string[]", Optional: pointer.True()},
			{Name: "status", Type: "string", Facet: pointer.True()},
			{Name: "created ", Type: "string"},
			{Name: "updated ", Type: "string"},
		},
		DefaultSortingField: pointer.String("created "),
		EnableNestedFields:  pointer.True(),
	}

	return schema
}

func (i *Product) Document() map[string]interface{} {
	document := map[string]interface{}{
		"id":          i.ID,
		"uid":         i.UID,
		"locale":      i.Locale,
		"type":        i.Type,
		"name":        i.Name,
		"description": i.Description,
		"created ":    time.Unix(int64(i.Created.T), 0).Format(time.RFC3339),
		"updated ":    time.Unix(int64(i.Updated.T), 0).Format(time.RFC3339),
	}

	return document
}

func (i *Product) Insert(collection typesense.CollectionInterface) error {
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

func (i *Product) Update(collection typesense.CollectionInterface, documentKey primitive.M, updatedFields primitive.M, removedFields primitive.A) error {
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
		var item *Product
		if err := database.Database.Collection(i.Collection()).FindOne(context.Background(), documentKey).Decode(&item); err != nil {
			return err
		}

		if err := i.Insert(collection); err != nil {
			return err
		}
	}

	return nil
}

func (i *Product) Delete(collection typesense.CollectionInterface, documentKey primitive.M) error {
	id := documentKey["_id"].(primitive.ObjectID).Hex()

	// Delete document from Typesense
	if _, err := collection.Document(id).Delete(); err != nil {
		return err
	}

	return nil
}
