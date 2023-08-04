package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Notification struct {
	Model      `bson:",inline"`
	User       primitive.ObjectID  `json:"user" bson:"user"`
	Notifiable Object              `json:"notifiable" bson:"notifiable"`
	Locale     string              `json:"locale" bson:"locale"`
	Type       string              `json:"type" bson:"type"`
	ReadAt     primitive.Timestamp `json:"read_at,omitempty" bson:"read_at,omitempty"`
}

func (i *Notification) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Notification
	return bson.Marshal((*t)(i))
}

func (i *Notification) Collection() string {
	return "notifications"
}

func (i *Notification) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "read_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
	}
}
