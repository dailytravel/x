package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	Model       `bson:",inline"`
	UID         *primitive.ObjectID `json:"uid,omitempty" bson:"uid,omitempty"`
	Type        string              `json:"type" bson:"type"`
	Name        string              `json:"name" bson:"name"`
	Description *string             `json:"description,omitempty" bson:"description,omitempty"`
	Secret      string              `json:"secret" bson:"secret"`
	Domain      string              `json:"domain" bson:"domain"`
	Redirect    string              `json:"redirect,omitempty" bson:"redirect,omitempty"`
	Provider    string              `json:"provider,omitempty" bson:"provider,omitempty"`
	Status      string              `json:"status" bson:"status"`
}

func (i *Client) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Client
	return bson.Marshal((*t)(i))
}

func (i *Client) Collection() string {
	return "clients"
}

func (i *Client) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "name", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "secret", Value: 1}}},
		{Keys: bson.D{{Key: "domain", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
	}
}
