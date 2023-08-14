package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Conversation struct {
	Model       `bson:",inline"`
	User        primitive.ObjectID  `json:"owner" bson:"owner"`
	Type        string              `json:"type" bson:"type"`
	Name        *string             `json:"name,omitempty" bson:"name,omitempty"`
	Description *string             `json:"description,omitempty" bson:"description,omitempty"`
	Labels      []*string           `json:"labels,omitempty" bson:"labels,omitempty"`
	Message     *primitive.ObjectID `json:"message" bson:"message"`
	Status      string              `json:"status" bson:"status"`
}

func (i *Conversation) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Conversation
	return bson.Marshal((*t)(i))
}

func (i *Conversation) Collection() string {
	return "conversations"
}

func (i *Conversation) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "user", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "type", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "message", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "deleted_at", Value: 1}}, Options: options.Index()},
	}
}
