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
	ReadAt     primitive.Timestamp `json:"read_at,omitempty" bson:"read_at,omitempty"`
}

type Notifiable struct {
	ID   primitive.ObjectID `json:"id" bson:"id"`
	Type string             `json:"type" bson:"type"`
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
		{Keys: bson.D{{Key: "uid", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "read_at", Value: 1}}},
		{Keys: bson.D{{Key: "created_at", Value: 1}}},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}},
	}
}
