package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Assignment struct {
	Model  `bson:",inline"`
	UID    primitive.ObjectID `json:"uid" bson:"uid"`
	Task   primitive.ObjectID `json:"task,omitempty" bson:"task,omitempty"`
	Status string             `json:"status" bson:"status"`
}

func (Assignment) IsEntity() {}

func (i *Assignment) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Assignment
	return bson.Marshal((*t)(i))
}

func (i *Assignment) Collection() string {
	return "assignments"
}

func (i *Assignment) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}}},
		{Keys: bson.D{{Key: "task", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
		{Keys: bson.D{{Key: "deleted", Value: 1}}},
	}
}
