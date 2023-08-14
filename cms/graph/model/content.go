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

type Content struct {
	Model       `bson:",inline"`
	User        primitive.ObjectID   `json:"user" bson:"user"`
	Parent      primitive.ObjectID   `json:"parent,omitempty" bson:"parent,omitempty"`
	Slug        string               `json:"slug" bson:"slug"`
	Locale      string               `json:"locale" bson:"locale"`
	Type        string               `json:"type" bson:"type"`
	Title       primitive.M          `json:"title" bson:"title"`
	Summary     primitive.M          `json:"summary" bson:"summary"`
	Body        primitive.M          `json:"body" bson:"body"`
	Status      string               `json:"status" bson:"status"`
	Commentable bool                 `json:"commentable" bson:"commentable"`
	Metadata    primitive.M          `json:"metadata" bson:"metadata"`
	Attachments []primitive.ObjectID `json:"attachments,omitempty" bson:"attachments,omitempty"`
}

func (Content) IsEntity() {}

func (i *Content) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Content
	return bson.Marshal((*t)(i))
}

func (i *Content) Collection() string {
	return "contents"
}

func (i *Content) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "user", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "parent", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "type", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "slug", Value: 1}}, Options: options.Index().SetUnique(true).SetSparse(true)},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "deleted_at", Value: 1}}, Options: options.Index()},
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
			{Name: "created_at", Type: "string"},
			{Name: "updated_at", Type: "string"},
		},
		DefaultSortingField: pointer.String("created_at"),
		EnableNestedFields:  pointer.True(),
	}

	return schema
}

func (i *Content) Document() map[string]interface{} {
	document := map[string]interface{}{
		"id":         i.ID.Hex(),
		"user":       i.User.Hex(),
		"locale":     i.Locale,
		"type":       i.Type,
		"title":      i.Title,
		"summary":    i.Summary,
		"body":       i.Body,
		"slug":       i.Slug,
		"created_at": time.Unix(int64(i.CreatedAt.T), 0).Format(time.RFC3339),
		"updated_at": time.Unix(int64(i.UpdatedAt.T), 0).Format(time.RFC3339),
	}

	if i.Parent != primitive.NilObjectID {
		document["parent"] = i.Parent.Hex()
	}

	return document
}

func (i *Content) Insert(collection typesense.CollectionInterface) error {
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

func (i *Content) Update(collection typesense.CollectionInterface, documentKey primitive.M, updatedFields primitive.M, removedFields primitive.A) error {
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
		var item *Content
		if err := db.Database.Collection(i.Collection()).FindOne(context.Background(), documentKey).Decode(&item); err != nil {
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

func (i *Content) Setuser(id *string) {
	if id != nil {
		i.User, _ = primitive.ObjectIDFromHex(*id)
	}
}

func (i *Content) SetParent(id *string) {
	if id != nil {
		i.Parent, _ = primitive.ObjectIDFromHex(*id)
	}
}

func (i *Content) SetType(t *string) {
	if t != nil {
		i.Type = *t
	}
}

func (i *Content) SetLocale(l *string) {
	if l != nil {
		i.Locale = *l
	}
}

func (i *Content) SetSlug(s *string) {
	if s != nil {
		i.Slug = *s
	}
}

func (i *Content) SetTitle(locale string, t *string) {
	if t != nil {
		i.Title[locale] = *t
	}
}

func (i *Content) SetSummary(locale string, s *string) {
	if s != nil {
		i.Summary[locale] = *s
	}
}

func (i *Content) SetBody(locale string, b *string) {
	if b != nil {
		i.Body[locale] = *b
	}
}

func (i *Content) SetStatus(s *string) {
	if s != nil {
		i.Status = *s
	}
}

// func (i *Content) SetCommentStatus(s *string) {
// 	if s != nil {
// 		i.CommentStatus = *s
// 	}
// }

// func (i *Content) SetCommentCount(c *int) {
// 	if c != nil {
// 		i.CommentCount = *c
// 	}
// }

func (i *Content) SetMetadata(m map[string]interface{}) {
	if m != nil {
		if i.Metadata == nil {
			i.Metadata = make(map[string]interface{})
		}

		for k, v := range m {
			i.Metadata[k] = v
		}
	}
}
