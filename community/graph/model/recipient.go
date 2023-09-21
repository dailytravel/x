package model

import (
	"time"

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
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
	}
}
