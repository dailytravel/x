package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Job struct {
	Model        `bson:",inline"`
	Code         string      `json:"code" bson:"code"`
	Locale       string      `json:"locale" bson:"locale"`
	Type         string      `json:"type" bson:"type"`
	Location     string      `json:"location" bson:"location"`
	Title        primitive.M `json:"title" bson:"title"`
	Description  primitive.M `json:"description,omitempty" bson:"description,omitempty"`
	Requirements primitive.M `json:"requirements,omitempty" bson:"requirements,omitempty"`
	Skills       primitive.M `json:"skills,omitempty" bson:"skills,omitempty"`
	Benefits     primitive.M `json:"benefits,omitempty" bson:"benefits,omitempty"`
	Status       string      `json:"status" bson:"status"`
}

func (i *Job) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Job
	return bson.Marshal((*t)(i))
}

func (i *Job) Collection() string {
	return "Jobs"
}

func (i *Job) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "locale", Value: 1}}},
		{Keys: bson.D{{Key: "type", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
		{Keys: bson.D{{Key: "deleted", Value: 1}}},
	}
}
