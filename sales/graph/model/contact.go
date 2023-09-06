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
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Contact struct {
	Model        `bson:",inline"`
	UID          primitive.ObjectID   `json:"uid" bson:"uid"`
	Company      *primitive.ObjectID  `json:"company,omitempty" bson:"company,omitempty"`
	FirstName    *string              `json:"firstName,omitempty" bson:"first_name,omitempty"`
	LastName     *string              `json:"lastName,omitempty" bson:"last_name,omitempty"`
	Birthday     primitive.DateTime   `json:"birthday,omitempty" bson:"birthday,omitempty"`
	Gender       *string              `json:"gender,omitempty" bson:"gender,omitempty"`
	JobTitle     *string              `json:"jobTitle,omitempty" bson:"job_title,omitempty"`
	Email        *string              `json:"email,omitempty" bson:"email,omitempty"`
	Phone        *string              `json:"phone,omitempty" bson:"phone,omitempty"`
	Picture      *string              `json:"picture,omitempty" bson:"picture,omitempty"`
	Street       *string              `json:"street,omitempty" bson:"street,omitempty"`
	City         *string              `json:"city,omitempty" bson:"city,omitempty"`
	Zip          *string              `json:"zip,omitempty" bson:"zip,omitempty"`
	State        *string              `json:"state,omitempty" bson:"state,omitempty"`
	Country      *string              `json:"country,omitempty" bson:"country,omitempty"`
	Website      *string              `json:"website,omitempty" bson:"website,omitempty"`
	Source       string               `json:"source,omitempty" bson:"source,omitempty"`
	Timezone     *string              `json:"timezone,omitempty" bson:"timezone,omitempty"`
	Language     *string              `json:"language,omitempty" bson:"language,omitempty"`
	Rating       *int                 `json:"rating,omitempty" bson:"rating,omitempty"`
	Subscribed   bool                 `json:"subscribed,omitempty" bson:"subscribed,omitempty"`
	Notes        *string              `json:"notes,omitempty" bson:"notes,omitempty"`
	Status       string               `json:"status" bson:"status"`
	LastActivity *primitive.Timestamp `json:"lastActivity,omitempty" bson:"last_activity,omitempty"`
	Labels       []*string            `json:"labels,omitempty" bson:"labels,omitempty"`
}

func (Contact) IsEntity() {}

func (i *Contact) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	// If CreatedAt is zero, then set it to the current timestamp.
	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Contact
	return bson.Marshal((*t)(i))
}

func (i *Contact) Collection() string {
	return "contacts"
}

func (i *Contact) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		// Basic properties indices
		{Keys: bson.D{{Key: "uid", Value: 1}}},
		{Keys: bson.D{{Key: "email", Value: 1}}},
		{Keys: bson.D{{Key: "phone", Value: 1}}},
		{Keys: bson.D{{Key: "type", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},

		// Additional properties indices
		{Keys: bson.D{{Key: "company", Value: 1}}},
		{Keys: bson.D{{Key: "source", Value: 1}}},
		{Keys: bson.D{{Key: "birthday", Value: 1}}},
		{Keys: bson.D{{Key: "labels", Value: 1}}},
		{Keys: bson.D{{Key: "last_activity", Value: 1}}},

		// Name search index with weighting
		{Keys: bson.D{{Key: "first_name", Value: "text"}, {Key: "last_name", Value: "text"}},
			Options: options.Index().SetWeights(bson.M{"first_name": 2, "last_name": 1})},

		// Timestamps indices
		{Keys: bson.D{{Key: "created_at", Value: 1}}},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}},
		{Keys: bson.D{{Key: "deleted_at", Value: 1}}},
	}
}

func (i *Contact) Schema() interface{} {
	return &api.CollectionSchema{
		Name: i.Collection(),
		Fields: []api.Field{
			{Name: "uid", Type: "string"},
			{Name: "company", Type: "string", Optional: pointer.True()},
			{Name: "first_name", Type: "string", Optional: pointer.True()},
			{Name: "last_name", Type: "string", Optional: pointer.True()},
			{Name: "job_title", Type: "string", Optional: pointer.True()},
			{Name: "gender", Type: "string", Optional: pointer.True()},
			{Name: "country", Type: "string", Optional: pointer.True(), Facet: pointer.True()},
			{Name: "source", Type: "string", Optional: pointer.True()},
			{Name: "status", Type: "string", Facet: pointer.True()},
			{Name: "rating", Type: "int32", Optional: pointer.True(), Facet: pointer.True()},
			{Name: "created_at", Type: "string"},
			{Name: "updated_at", Type: "string"},
			{Name: "birthday", Type: "string", Facet: pointer.True()},
			{Name: "email", Type: "string", Optional: pointer.True()},
			{Name: "phone", Type: "string", Optional: pointer.True()},
			{Name: "labels", Type: "string[]", Optional: pointer.True()},
		},
		TokenSeparators:    &[]string{"(", ")", "-"},
		EnableNestedFields: pointer.True(),
	}
}

func (i *Contact) Document() map[string]interface{} {
	document := map[string]interface{}{
		"id":         i.ID,
		"uid":        i.UID,
		"company":    i.Company,
		"first_name": i.FirstName,
		"last_name":  i.LastName,
		"job_title":  i.JobTitle,
		"gender":     i.Gender,
		"source":     i.Source,
		"country":    i.Country,
		"email":      i.Email,
		"phone":      i.Phone,
		"metadata":   i.Metadata,
		"status":     i.Status,
		"created_at": time.Unix(int64(i.CreatedAt.T), 0).Format(time.RFC3339),
		"updated_at": time.Unix(int64(i.UpdatedAt.T), 0).Format(time.RFC3339),
		"labels":     i.Labels,
	}

	if i.Rating != nil {
		document["rating"] = int32(*i.Rating)
	}

	if !i.Birthday.Time().IsZero() {
		document["birthday"] = i.Birthday.Time().Format(time.RFC3339)
	}

	return document
}

func (i *Contact) Insert(collection typesense.CollectionInterface) error {
	document := i.Document()

	// Retrieve Typesense collection schema and create it if it doesn't exist
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

func (i *Contact) Update(collection typesense.CollectionInterface, documentKey primitive.M, updatedFields primitive.M, removedFields primitive.A) error {
	documentID := documentKey["_id"].(primitive.ObjectID).Hex()

	// Check for 'deleted_at' field
	if _, exists := updatedFields["deleted_at"].(primitive.Timestamp); exists {
		return i.Delete(collection, documentKey)
	}

	// Prepare the update payload
	updatePayload := make(map[string]interface{})

	// Populate the update payload with updated fields
	for field, value := range updatedFields {
		switch field {
		case "created_at", "updated_at":
			if timestamp, ok := value.(primitive.Timestamp); ok {
				updatePayload[field] = time.Unix(int64(timestamp.T), 0).Format(time.RFC3339)
			}
		case "birthday":
			if timestamp, ok := value.(primitive.DateTime); ok {
				updatePayload[field] = timestamp.Time().Format(time.RFC3339)
			}
		default:
			updatePayload[field] = value
		}
	}

	// Set removed fields to nil in the update payload
	for _, field := range removedFields {
		updatePayload[field.(string)] = nil
	}

	// Attempt to update the document
	_, err := collection.Document(documentID).Update(updatePayload)
	if err != nil {
		// If update fails, try to fetch the item from MongoDB and re-insert into Typesense
		var item Contact
		if err := database.Database.Collection(i.Collection()).FindOne(context.Background(), documentKey).Decode(&item); err != nil {
			return err
		}
		return i.Insert(collection)
	}

	return nil
}

func (i *Contact) Delete(collection typesense.CollectionInterface, documentKey primitive.M) error {
	id := documentKey["_id"].(primitive.ObjectID).Hex()

	// Delete document from Typesense
	if _, err := collection.Document(id).Delete(); err != nil {
		return err
	}

	return nil
}
