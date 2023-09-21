package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Quote struct {
	Model
	UID         primitive.ObjectID  `json:"uid" bson:"uid"`
	Contact     primitive.ObjectID  `json:"contacts" bson:"contacts"`
	Code        string              `json:"code" bson:"code"`
	Purchase    *string             `json:"purchase,omitempty" bson:"purchase,omitempty"`
	Locale      string              `json:"locale" bson:"locale"`
	Name        string              `json:"name" bson:"name"`
	Description *string             `json:"description,omitempty" bson:"description,omitempty"`
	Template    string              `json:"template" bson:"template"`
	ValidUntil  primitive.Timestamp `json:"validUntil" bson:"valid_until"`
	Terms       string              `json:"terms" bson:"terms"`
	Payment     string              `json:"payment" bson:"payment"`
	Notes       string              `json:"notes,omitempty" bson:"notes,omitempty"`
	Billing     primitive.M         `json:"billing,omitempty" bson:"billing,omitempty"`
	Status      string              `json:"status" bson:"status"`
}

func (i *Quote) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Quote
	return bson.Marshal((*t)(i))
}

func (i *Quote) Collection() string {
	return "quotes"
}

func (i *Quote) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}}},
		{Keys: bson.D{{Key: "contact", Value: 1}}},
		{Keys: bson.D{{Key: "locale", Value: 1}}},
		{Keys: bson.D{{Key: "code", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "purchase", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
		{Keys: bson.D{{Key: "deleted", Value: 1}}},
	}
}
