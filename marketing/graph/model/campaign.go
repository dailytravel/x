package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Campaign struct {
	Model    `bson:",inline"`
	Audience *primitive.ObjectID `json:"audience,omitempty" bson:"audience,omitempty"`
	Type     string              `json:"type" bson:"type"`
	Name     string              `json:"name" bson:"name"`
	Metadata primitive.M         `json:"metadata,omitempty" bson:"metadata,omitempty"`
	Status   string              `json:"status" bson:"status"`
}

func (Campaign) IsEntity() {}

func (i *Campaign) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Campaign
	return bson.Marshal((*t)(i))
}

func (i *Campaign) Collection() string {
	return "campaigns"
}

func (i *Campaign) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "audience", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated", Value: 1}}, Options: options.Index()},
	}
}
