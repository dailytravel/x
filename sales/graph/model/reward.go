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
	Tier        primitive.ObjectID `bson:"tier,omitempty" json:"tier,omitempty"`
	Locale      string             `bson:"locale,omitempty" json:"locale,omitempty"`
	Name        primitive.M        `json:"name,omitempty" bson:"name,omitempty"`
	Destination primitive.M        `json:"destination,omitempty" bson:"destination,omitempty"`
	Cost        int64              `json:"cost,omitempty" bson:"cost,omitempty"`
	ExpiresAt   int64              `json:"expires_at,omitempty" bson:"expires_at,omitempty"`
	Status      string             `json:"status,omitempty" bson:"status,omitempty"`
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
