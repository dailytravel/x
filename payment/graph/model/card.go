package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Card struct {
	Model       `bson:",inline"`
	UID         primitive.ObjectID  `json:"uid" bson:"uid"`
	Wallet      *primitive.ObjectID `json:"wallet" bson:"wallet"`
	Name        string              `json:"name" bson:"name"`
	Number      string              `json:"number" bson:"number"`
	ExpMonth    int                 `json:"exp_month" bson:"exp_month"`
	ExpYear     int                 `json:"exp_year" bson:"exp_year"`
	Cvv         string              `json:"cvv" bson:"cvv"`
	Last4       string              `json:"last4" bson:"last4"`
	Brand       string              `json:"brand" bson:"brand"`
	Country     string              `json:"country" bson:"country"`
	Funding     string              `json:"funding" bson:"funding"`
	Fingerprint string              `json:"fingerprint" bson:"fingerprint"`
	Billing     primitive.M         `json:"billing" bson:"billing"`
	Status      string              `json:"status" bson:"status"`
}

func (i *Card) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Card
	return bson.Marshal((*t)(i))
}

func (i *Card) Collection() string {
	return "cards"
}

func (i *Card) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}}},
		{Keys: bson.D{{Key: "number", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "name", Value: 1}}},
		{Keys: bson.D{{Key: "last4", Value: 1}}},
		{Keys: bson.D{{Key: "brand", Value: 1}}},
		{Keys: bson.D{{Key: "country", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
	}
}
