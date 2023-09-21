package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Cart struct {
	Model   `bson:",inline"`
	UID     primitive.ObjectID `bson:"uid" json:"uid"`
	Expires primitive.DateTime `bson:"expires" json:"expires"`
	Status  string             `json:"status" bson:"status"`
}

func (i *Cart) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Cart
	return bson.Marshal((*t)(i))
}

func (i *Cart) Collection() string {
	return "orders"
}

func (i *Cart) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}}},
		{Keys: bson.D{{Key: "expires", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
	}
}
