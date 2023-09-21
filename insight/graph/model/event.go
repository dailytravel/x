package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Event struct {
	Model     `bson:",inline"`
	Campaign  *primitive.ObjectID `json:"campaign,omitempty" bson:"campaign,omitempty"`
	Sender    string              `json:"sender" bson:"sender"`
	Recipient string              `json:"recipient" bson:"recipient"`
	Subject   string              `json:"subject" bson:"subject"`
	Message   string              `json:"message" bson:"message"`
	Status    string              `json:"status" bson:"status"`
	ClientIP  *string             `json:"clientIp,omitempty" bson:"client_ip,omitempty"`
	UserAgent *string             `json:"userAgent,omitempty" bson:"user_agent,omitempty"`
	Exception *string             `json:"exception,omitempty" bson:"exception,omitempty"`
}

func (i *Event) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Timestamp.IsZero() {
		i.Timestamp = now
	}

	type t Event
	return bson.Marshal((*t)(i))
}

func (i *Event) Collection() string {
	return "events"
}

func (i *Event) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "client_ip", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "timestamp", Value: 1}}},
	}
}
