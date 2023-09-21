package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Wallet struct {
	Model    `bson:",inline"`
	UID      primitive.ObjectID `json:"uid" bson:"uid"`
	Balance  float64            `json:"balance" bson:"balance"`
	Currency string             `json:"currency" bson:"currency"`
	Status   string             `json:"status" bson:"status"`
}

func (i *Wallet) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Wallet
	return bson.Marshal((*t)(i))
}

func (i *Wallet) Collection() string {
	return "wallets"
}

func (i *Wallet) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
	}
}
