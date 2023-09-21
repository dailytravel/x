package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Conversation struct {
	Model       `bson:",inline"`
	Type        string              `json:"type" bson:"type"`
	Name        *string             `json:"name,omitempty" bson:"name,omitempty"`
	Description *string             `json:"description,omitempty" bson:"description,omitempty"`
	Labels      []*string           `json:"labels,omitempty" bson:"labels,omitempty"`
	Message     *primitive.ObjectID `json:"message" bson:"message"`
	Status      string              `json:"status" bson:"status"`
}

func (i *Conversation) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Conversation
	return bson.Marshal((*t)(i))
}

func (i *Conversation) Collection() string {
	return "conversations"
}

func (i *Conversation) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "type", Value: 1}}},
		{Keys: bson.D{{Key: "message", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
		{Keys: bson.D{{Key: "deleted", Value: 1}}},
	}
}
