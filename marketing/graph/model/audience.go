package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Audience
	return bson.Marshal((*t)(i))
}

func (i *Audience) Collection() string {
	return "audiences"
}

func (i *Audience) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
	}
}
