package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	Model      `bson:",inline"`
	UID        primitive.ObjectID `json:"uid" bson:"uid"`
	Position   string             `json:"position" bson:"position"`
	Resume     string             `json:"resume" bson:"resume"`
	Interview  primitive.DateTime `json:"interview" bson:"interview"`
	Feedback   *string            `json:"feedback" bson:"feedback"`
	ReferredBy *string            `json:"referredBy" bson:"referred_by"`
	Notes      string             `json:"notes" bson:"notes"`
	Status     string             `json:"status" bson:"status"`
}

func (i *Application) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Application
	return bson.Marshal((*t)(i))
}

func (i *Application) Collection() string {
	return "applications"
}

func (i *Application) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}}},
		{Keys: bson.D{{Key: "job", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
		{Keys: bson.D{{Key: "deleted", Value: 1}}},
	}
}
