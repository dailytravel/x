package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Wishlist struct {
	Model   `bson:",inline"`
	UID     primitive.ObjectID `json:"uid" bson:"uid"`
	Content primitive.ObjectID `json:"content" bson:"content"`
	Status  string             `json:"status" bson:"status"`
}

func (i *Wishlist) Collection() string {
	return "wishlists"
}

func (i *Wishlist) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Wishlist
	return bson.Marshal((*t)(i))
}

func (i *Wishlist) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}, {Key: "product", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
	}
}
