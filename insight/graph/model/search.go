package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Search struct {
	Model     `bson:",inline"`
	UID       *primitive.ObjectID `json:"uid" bson:"uid"`
	Locale    string              `json:"locale" bson:"locale"`
	Keyword   string              `json:"keyword" bson:"keyword"`
	ClientIP  *string             `json:"client_ip,omitempty" bson:"client_ip,omitempty"`
	UserAgent *string             `json:"user_agent,omitempty" bson:"user_agent,omitempty"`
}

func (i *Search) Collection() string {
	return "searches"
}

func (i *Search) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Timestamp.IsZero() {
		i.Timestamp = now
	}

	type t Search
	return bson.Marshal((*t)(i))
}

func (i *Search) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}}},
		{Keys: bson.D{{Key: "locale", Value: 1}}},
		{Keys: bson.D{{Key: "client_ip", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "timestamp", Value: 1}}},
	}
}
