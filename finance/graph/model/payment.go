package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Payment struct {
	Model    `bson:",inline"`
	UID      primitive.ObjectID `json:"uid" bson:"uid"`
	Invoice  primitive.ObjectID `json:"invoice" bson:"invoice"`
	Method   string             `json:"method" bson:"method"`
	Amount   float64            `json:"amount" bson:"amount"`
	Currency string             `json:"currency" bson:"currency"`
	Date     primitive.DateTime `json:"date" bson:"date"`
	Status   string             `json:"status" bson:"status"`
}

func (i *Payment) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Payment
	return bson.Marshal((*t)(i))
}

func (i *Payment) Collection() string {
	return "payments"
}

func (i *Payment) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}}},
		{Keys: bson.D{{Key: "invoice", Value: 1}}},
		{Keys: bson.D{{Key: "date", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
	}
}
