package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Tier struct {
	Model       `bson:",inline"`
	Locale      string               `bson:"locale,omitempty" json:"locale,omitempty"`
	Name        primitive.M          `json:"name,omitempty" bson:"name,omitempty"`
	Description primitive.M          `json:"description,omitempty" bson:"description,omitempty"`
	Benefits    []primitive.ObjectID `json:"benefits,omitempty" bson:"benefits,omitempty"`
	Cost        int64                `json:"cost,omitempty" bson:"cost,omitempty"`
	Status      string               `json:"status,omitempty" bson:"status,omitempty"`
}

func (i *Tier) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Tier
	return bson.Marshal((*t)(i))
}

func (i *Tier) Collection() string {
	return "tiers"
}

func (i *Tier) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "locale", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
	}
}
