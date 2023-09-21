package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Transaction struct {
	Model       `bson:",inline"`
	UID         primitive.ObjectID  `json:"uid" bson:"uid"`
	Wallet      *primitive.ObjectID `json:"wallet" bson:"wallet"`
	Card        *primitive.ObjectID `json:"card" bson:"card"`
	Type        string              `json:"type" bson:"type"`
	Status      string              `json:"status" bson:"status"`
	Amount      float64             `json:"amount" bson:"amount"`
	Currency    string              `json:"currency" bson:"currency"`
	Date        primitive.DateTime  `json:"date" bson:"date"`
	Description *string             `json:"description,omitempty" bson:"description,omitempty"`
}

func (i *Transaction) Collection() string {
	return "transactions"
}

func (i *Transaction) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Transaction
	return bson.Marshal((*t)(i))
}

func (i *Transaction) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}}},
		{Keys: bson.D{{Key: "type", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "date", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
	}
}
