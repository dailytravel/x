package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Session struct {
	Model   `bson:",inline"`
	UID     primitive.ObjectID  `json:"uid" bson:"uid"`
	Token   string              `json:"token" bson:"token"`
	Expires primitive.Timestamp `json:"expires" bson:"expires"`
	Status  string              `json:"status" bson:"status"`
}

func (i *Session) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Session
	return bson.Marshal((*t)(i))
}

func (i *Session) Collection() string {
	return "sessions"
}

func (i *Session) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}}},
		{Keys: bson.D{{Key: "token", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "expires", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
	}
}
