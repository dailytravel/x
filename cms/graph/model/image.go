package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Image struct {
	Model   `bson:",inline"`
	Object  Object      `json:"object" bson:"object"`
	Locale  string      `json:"locale" bson:"locale"`
	Title   primitive.M `json:"title" bson:"title"`
	Caption primitive.M `json:"caption" bson:"caption"`
	Type    string      `json:"type" bson:"type"`
	URL     string      `json:"url" bson:"url"`
	Order   int         `json:"order" bson:"order"`
}

type Object struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	Collection string             `json:"collection" bson:"collection"`
}

func (Image) IsEntity() {}

func (i *Image) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Image
	return bson.Marshal((*t)(i))
}

func (i *Image) Collection() string {
	return "files"
}

func (i *Image) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "object._id", Value: 1}, {Key: "object.collection", Value: 1}}},
		{Keys: bson.D{{Key: "type", Value: 1}}},
		{Keys: bson.D{{Key: "locale", Value: 1}}},
		{Keys: bson.D{{Key: "order", Value: 1}}},
		{Keys: bson.D{{Key: "created ", Value: 1}}},
		{Keys: bson.D{{Key: "updated ", Value: 1}}},
	}
}
