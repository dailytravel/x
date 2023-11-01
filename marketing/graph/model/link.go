package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Link struct {
	Model       `bson:",inline"`
	Domain      string               `json:"domain" bson:"domain"`
	Code        string               `json:"code" bson:"code"`
	URL         string               `json:"url" bson:"url"`
	Title       *string              `json:"title,omitempty" bson:"title,omitempty"`
	Status      *string              `json:"status,omitempty" bson:"status,omitempty"`
	Engagements int64                `json:"engagements" bson:"engagements"`
	UID         primitive.ObjectID   `json:"uid" bson:"uid"`
	Tags        []primitive.ObjectID `json:"tags" bson:"tags"`
}

func (Link) IsEntity() {}

func (i *Link) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Link
	return bson.Marshal((*t)(i))
}

func (i *Link) Collection() string {
	return "links"
}

func (i *Link) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "code", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "uid", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
	}
}
