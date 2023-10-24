package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Activity struct {
	Model     `bson:",inline"`
	UID       primitive.ObjectID `json:"uid" bson:"uid"`
	Object    Object             `json:"object" bson:"object"`
	Action    string             `json:"action" bson:"action"`
	ClientIP  *string            `json:"client_ip,omitempty" bson:"client_ip,omitempty"`
	UserAgent *string            `json:"user_agent,omitempty" bson:"user_agent,omitempty"`
}

type Object struct {
	ID         primitive.ObjectID `json:"id" bson:"id"`
	Collection string             `json:"collection" bson:"collection"`
}

func (i *Activity) Collection() string {
	return "activities"
}

func (i *Activity) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Timestamp.IsZero() {
		i.Timestamp = now
	}

	type t Activity
	return bson.Marshal((*t)(i))
}

func (i *Activity) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}}},
		{Keys: bson.D{{Key: "object._id", Value: 1}, {Key: "object.collection", Value: 1}}},
		{Keys: bson.D{{Key: "action", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "timestamp", Value: 1}}},
	}
}
