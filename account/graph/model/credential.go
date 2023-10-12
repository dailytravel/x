package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Credential struct {
	Model
	UID     primitive.ObjectID  `json:"uid" bson:"uid"`
	Type    string              `json:"type" bson:"type"`
	Secret  string              `json:"secret" bson:"secret"`
	Expires primitive.Timestamp `json:"expires" bson:"expires"`
	Status  string              `json:"status"`
}

func (i *Credential) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Credential
	return bson.Marshal((*t)(i))
}

func (i *Credential) Collection() string {
	return "credentials"
}

func (i *Credential) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
	}
}
