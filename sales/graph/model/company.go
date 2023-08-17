package model

import (
	"context"
	"time"

	"github.com/dailytravel/x/sales/db"
	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
	"github.com/typesense/typesense-go/typesense/api/pointer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Company struct {
	Model       `bson:",inline"`
	UID         primitive.ObjectID `json:"uid" bson:"uid"`
	Type        string             `json:"type" bson:"type"`
	Name        string             `json:"name" bson:"name"`
	Description *string            `json:"description,omitempty" bson:"description,omitempty"`
	Industry    *string            `json:"industry,omitempty" bson:"industry,omitempty"`
	Employees   *int               `json:"employees,omitempty" bson:"employees,omitempty"`
	Revenue     *float64           `json:"revenue,omitempty" bson:"revenue,omitempty"`
	City        *string            `json:"city,omitempty" bson:"city,omitempty"`
	Zip         *string            `json:"zip,omitempty" bson:"zip,omitempty"`
	State       *string            `json:"state,omitempty" bson:"state,omitempty"`
	Country     *string            `json:"country,omitempty" bson:"country,omitempty"`
	Timezone    *string            `json:"timezone,omitempty" bson:"timezone,omitempty"`
	Phone       *string            `json:"phone,omitempty" bson:"phone,omitempty"`
	Website     *string            `json:"website,omitempty" bson:"website,omitempty"`
	Status      string             `json:"status" bson:"status"`
}

func (Company) IsEntity() {}

func (i *Company) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Company
	return bson.Marshal((*t)(i))
}

func (i *Company) Collection() string {
	return "companies"
}

func (i *Company) Sanitize(s string) string {
	return s
}

func (i *Company) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "type", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "deleted_at", Value: 1}}, Options: options.Index()},
	}
}

func (i *Company) Schema() interface{} {
	return &api.CollectionSchema{
		Name: i.Collection(),
		Fields: []api.Field{
			{Name: "uid", Type: "string"},
			{Name: "name", Type: "string", Optional: pointer.True()},
			{Name: "type", Type: "string", Optional: pointer.True()},
			{Name: "status", Type: "string", Facet: pointer.True()},
			{Name: "created_by", Type: "string", Optional: pointer.True()},
			{Name: "updated_by", Type: "string", Optional: pointer.True()},
			{Name: "created_at", Type: "int32"},
			{Name: "updated_at", Type: "int32"},
		},
		DefaultSortingField: pointer.String("created_at"),
		TokenSeparators:     &[]string{"(", ")", "-"},
		EnableNestedFields:  pointer.True(),
	}
}

func (i *Company) Document() map[string]interface{} {
	// followers := i.Followers()

	document := map[string]interface{}{
		"id": i.ID,
	}

	return document
}

func (i *Company) Insert(collection typesense.CollectionInterface) error {
	document := i.Document()

	// Retrieve Typesense collection schema and create it if it doesn't exist
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

func (i *Company) Update(collection typesense.CollectionInterface, documentKey primitive.M, updatedFields primitive.M, removedFields primitive.A) error {
	documentID := documentKey["_id"].(primitive.ObjectID).Hex()

	// Create a map to hold the updated fields
	updatePayload := make(map[string]interface{})

	// Check if 'deleted_at' field is in updatedFields and its value is of type primitive.Timestamp
	_, deletedAtExist := updatedFields["deleted_at"].(primitive.Timestamp)
	if deletedAtExist {
		if err := i.Delete(collection, documentKey); err != nil {
			return err
		}
		return nil
	}

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
		// If the update fails, attempt to retrieve the item from the dataModel
		var item *Company
		if err := db.Database.Collection(i.Collection()).FindOne(context.Background(), documentKey).Decode(&item); err != nil {
			return err
		}

		// Insert the item if it doesn't exist in the collection
		if err := i.Insert(collection); err != nil {
			return err
		}
	}

	return nil
}

func (i *Company) Delete(collection typesense.CollectionInterface, documentKey primitive.M) error {
	id := documentKey["_id"].(primitive.ObjectID).Hex()

	// Delete document from Typesense
	if _, err := collection.Document(id).Delete(); err != nil {
		return err
	}

	return nil
}
