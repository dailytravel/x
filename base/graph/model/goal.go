package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Goal struct {
	Model        `bson:",inline"`
	UID          primitive.ObjectID  `bson:"uid" json:"uid"`
	Parent       *primitive.ObjectID `json:"parent,omitempty" bson:"parent,omitempty"`
	Organization *primitive.ObjectID `json:"organization,omitempty" bson:"organization,omitempty"`
	Time         primitive.ObjectID  `json:"time" bson:"time"`
	Name         string              `json:"name" bson:"name"`
	Notes        *string             `json:"notes,omitempty" bson:"notes,omitempty"`
	Start        *primitive.DateTime `json:"start,omitempty" bson:"start,omitempty"`
	End          *primitive.DateTime `json:"end,omitempty" bson:"end,omitempty"`
	Status       string              `json:"status" bson:"status"`
}

func (i *Goal) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Goal
	return bson.Marshal((*t)(i))
}

func (i *Goal) Collection() string {
	return "goals"
}

func (i *Goal) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "user", Value: 1}}},
		{Keys: bson.D{{Key: "parent", Value: 1}}},
		{Keys: bson.D{{Key: "organization", Value: 1}}},
		{Keys: bson.D{{Key: "time", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "start", Value: 1}}},
		{Keys: bson.D{{Key: "end", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
		{Keys: bson.D{{Key: "deleted", Value: 1}}},
	}
}
