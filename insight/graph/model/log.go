package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Log struct {
	Model     `bson:",inline"`
	UID       *primitive.ObjectID `json:"uid,omitempty" bson:"uid,omitempty"`
	URL       string              `json:"url" bson:"url"`
	Referrer  *string             `json:"referrer,omitempty" bson:"referrer,omitempty"`
	Title     *string             `json:"title,omitempty" bson:"title,omitempty"`
	Utm       primitive.M         `json:"utm,omitempty" bson:"utm,omitempty"`
	Status    string              `json:"status" bson:"status"`
	ClientIP  *string             `json:"clientIp,omitempty" bson:"client_ip,omitempty"`
	UserAgent *string             `json:"userAgent,omitempty" bson:"user_agent,omitempty"`
}

func (i *Log) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Timestamp.IsZero() {
		i.Timestamp = now
	}

	type t Log
	return bson.Marshal((*t)(i))
}

func (i *Log) Collection() string {
	return "logs"
}

func (i *Log) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "client_ip", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "timestamp", Value: 1}}},
	}
}
