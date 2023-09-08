package model

import (
	"time"

	"github.com/typesense/typesense-go/typesense/api"
	"github.com/typesense/typesense-go/typesense/api/pointer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Recipient struct {
	Model   `bson:",inline"`
	UID     primitive.ObjectID   `json:"uid" bson:"uid"`
	Message primitive.ObjectID   `json:"message" bson:"message"`
	Read    *primitive.Timestamp `json:"read,omitempty" bson:"read,omitempty"`
}

func (i *Recipient) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Recipient
	return bson.Marshal((*t)(i))
}

func (i *Recipient) Collection() string {
	return "recipients"
}

func (i *Recipient) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}, {Key: "message", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated", Value: 1}}, Options: options.Index()},
	}
}

func (i *Recipient) Schema() interface{} {
	return &api.CollectionSchema{
		Name: i.Collection(),
		Fields: []api.Field{
			{Name: "user", Type: "string", Facet: pointer.True()},
			{Name: "message", Type: "string", Facet: pointer.True()},
			{Name: "status", Type: "string", Facet: pointer.True()},
			{Name: "read", Type: "int64", Facet: pointer.True()},
			{Name: "created", Type: "int64", Facet: pointer.True()},
			{Name: "updated", Type: "int64", Facet: pointer.True()},
		},
		DefaultSortingField: pointer.String("created"),
	}
}
