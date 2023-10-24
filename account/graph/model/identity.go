package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Identity struct {
	Model      `bson:",inline"`
	UID        primitive.ObjectID `json:"uid" bson:"uid"`
	Provider   string             `json:"provider" bson:"provider"`
	UserID     string             `json:"user_id" bson:"user_id"`
	Connection string             `json:"connection" bson:"connection"`
	IsSocial   bool               `json:"is_social" bson:"is_social"`
	Status     string             `json:"status"`
}

func (i *Identity) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Identity
	return bson.Marshal((*t)(i))
}

func (i *Identity) Collection() string {
	return "identities"
}

func (i *Identity) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "provider", Value: 1}, {Key: "user_id", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
	}
}
