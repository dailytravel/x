package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Log struct {
	Model     `bson:",inline"`
	UID       *primitive.ObjectID `json:"uid,omitempty" bson:"uid,omitempty"`
	URL       string              `json:"url" bson:"url"`
	UTM       primitive.M         `json:"utm,omitempty" bson:"utm,omitempty"`
	UserAgent string              `json:"user_agent" bson:"user_agent"`
	ClientIP  string              `json:"client_ip" bson:"client_ip"`
	Status    string              `json:"status" bson:"status"`
}

func (i *Log) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Log
	return bson.Marshal((*t)(i))
}

func (i *Log) Collection() string {
	return "logs"
}

func (i *Log) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "client_ip", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
	}
}
