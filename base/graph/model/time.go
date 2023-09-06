package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Time struct {
	Model       `bson:",inline"`
	Parent      *primitive.ObjectID `json:"parent,omitempty" bson:"parent,omitempty"`
	Name        string              `json:"name" bson:"name"`
	Description string              `json:"description,omitempty" bson:"description,omitempty"`
	Start       primitive.DateTime  `json:"start" bson:"start"`
	End         primitive.DateTime  `json:"end" bson:"end"`
}

func (i *Time) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Time
	return bson.Marshal((*t)(i))
}

func (i *Time) Collection() string {
	return "times"
}

func (i *Time) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "start", Value: 1}, {Key: "end", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "created_at", Value: 1}}},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}},
	}
}
