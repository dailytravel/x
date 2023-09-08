package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Balance struct {
	Model   `bson:",inline"`
	UID     primitive.ObjectID   `json:"uid" bson:"uid"`
	Type    string               `json:"type" bson:"type"`
	Credits int                  `json:"credits" bson:"credits"`
	Notes   *string              `json:"notes,omitempty" bson:"notes,omitempty"`
	Expires *primitive.Timestamp `json:"expires,omitempty" bson:"expires,omitempty"`
	Status  string               `json:"status" bson:"status"`
}

func (i *Balance) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Balance
	return bson.Marshal((*t)(i))
}

func (i *Balance) Collection() string {
	return "balances"
}

func (i *Balance) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "expires", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
		{Keys: bson.D{{Key: "deleted", Value: 1}}},
	}
}
