package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Audience struct {
	Model       `bson:",inline"`
	Name        string                `json:"name" bson:"name"`
	Description *string               `json:"description,omitempty" bson:"description,omitempty"`
	Segments    []*primitive.ObjectID `json:"segments,omitempty" bson:"segments,omitempty"`
	Metadata    primitive.M           `json:"metadata,omitempty" bson:"metadata,omitempty"`
	Status      string                `json:"status" bson:"status"`
}

func (i *Audience) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Audience
	return bson.Marshal((*t)(i))
}

func (i *Audience) Collection() string {
	return "audiences"
}

func (i *Audience) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
	}
}
