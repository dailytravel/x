package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Notification struct {
	Model      `bson:",inline"`
	UID        primitive.ObjectID  `json:"uid" bson:"uid"`
	Notifiable Notifiable          `json:"notifiable" bson:"notifiable"`
	Locale     string              `json:"locale" bson:"locale"`
	Type       string              `json:"type" bson:"type"`
	Read       primitive.Timestamp `json:"read,omitempty" bson:"read,omitempty"`
}

type Notifiable struct {
	ID   primitive.ObjectID `json:"id" bson:"id"`
	Type string             `json:"type" bson:"type"`
}

func (i *Notification) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Notification
	return bson.Marshal((*t)(i))
}

func (i *Notification) Collection() string {
	return "notifications"
}

func (i *Notification) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "read", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
	}
}
