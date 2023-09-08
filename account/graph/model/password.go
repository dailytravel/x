package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Password struct {
	Model `bson:",inline"`
	UID   primitive.ObjectID `json:"uid" bson:"uid"`
	Hash  string             `json:"hash" bson:"hash"`
}

func (i *Password) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Password
	return bson.Marshal((*t)(i))
}

func (i *Password) Collection() string {
	return "passwords"
}

func (i *Password) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
	}
}
