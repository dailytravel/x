package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Benefit struct {
	Model       `bson:",inline"`
	Locale      string      `json:"locale" bson:"locale"`
	Description primitive.M `json:"description" bson:"description"`
	Status      string      `json:"status" bson:"status"`
}

func (i *Benefit) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Benefit
	return bson.Marshal((*t)(i))
}

func (i *Benefit) Collection() string {
	return "benefits"
}

func (i *Benefit) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "locale", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
		{Keys: bson.D{{Key: "deleted", Value: 1}}},
	}
}
