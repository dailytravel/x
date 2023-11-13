package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Share struct {
	Model      `bson:",inline"`
	UID        primitive.ObjectID `bson:"uid" json:"uid"`
	Object     Object             `json:"object" bson:"object"`
	Permission string             `json:"permission" bson:"permission"`
	Status     string             `json:"status" bson:"status"`
}

func (Share) IsEntity() {}

type Object struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Type string             `json:"type" bson:"type"`
}

func (i *Share) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Share
	return bson.Marshal((*t)(i))
}

func (i *Share) Collection() string {
	return "shares"
}

func (i *Share) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}, {Key: "object._id", Value: 1}, {Key: "object.collection", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
	}
}
