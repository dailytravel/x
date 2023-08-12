package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Goal struct {
	Model        `bson:",inline"`
	Owner        primitive.ObjectID  `bson:"owner" json:"owner"`
	Parent       primitive.ObjectID  `json:"parent,omitempty" bson:"parent,omitempty"`
	Organization primitive.ObjectID  `json:"organization" bson:"organization"`
	Time         primitive.ObjectID  `json:"time" bson:"time"`
	Name         string              `json:"name" bson:"name"`
	Notes        string              `json:"notes,omitempty" bson:"notes,omitempty"`
	StartDate    primitive.Timestamp `json:"start_date" bson:"start_date"`
	DueDate      primitive.Timestamp `json:"due_date" bson:"due_date"`
	Status       string              `json:"status" bson:"status"`
}

func (i *Goal) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Goal
	return bson.Marshal((*t)(i))
}

func (i *Goal) Collection() string {
	return "goals"
}

func (i *Goal) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "owner", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "parent", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "organization", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "time", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "deleted_at", Value: 1}}, Options: options.Index()},
	}
}
