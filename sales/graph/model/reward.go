package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Reward struct {
	Model       `bson:",inline"`
	Tier        primitive.ObjectID  `bson:"tier" json:"tier"`
	Locale      string              `bson:"locale" json:"locale"`
	Name        primitive.M         `json:"name" bson:"name"`
	Description primitive.M         `json:"description,omitempty" bson:"description,omitempty"`
	Cost        int64               `json:"cost" bson:"cost"`
	Expires     primitive.Timestamp `json:"expires,omitempty" bson:"expires,omitempty"`
	Status      string              `json:"status" bson:"status"`
}

func (i *Reward) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Reward
	return bson.Marshal((*t)(i))
}

func (i *Reward) Collection() string {
	return "rewards"
}

func (i *Reward) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "program", Value: 1}}},
		{Keys: bson.D{{Key: "tier", Value: 1}}},
		{Keys: bson.D{{Key: "locale", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "expires", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
		{Keys: bson.D{{Key: "deleted", Value: 1}}},
	}
}
