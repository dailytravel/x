package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Reward struct {
	Model       `bson:",inline"`
	Tier        primitive.ObjectID `bson:"tier" json:"tier"`
	Locale      string             `bson:"locale" json:"locale"`
	Name        primitive.M        `json:"name" bson:"name"`
	Description primitive.M        `json:"description,omitempty" bson:"description,omitempty"`
	Cost        int64              `json:"cost" bson:"cost"`
	ExpiresAt   int64              `json:"expires_at,omitempty" bson:"expires_at,omitempty"`
	Status      string             `json:"status" bson:"status"`
}

func (i *Reward) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Reward
	return bson.Marshal((*t)(i))
}

func (i *Reward) Collection() string {
	return "rewards"
}

func (i *Reward) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "program", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "tier", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "locale", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "expires_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "deleted_at", Value: 1}}, Options: options.Index()},
	}
}
