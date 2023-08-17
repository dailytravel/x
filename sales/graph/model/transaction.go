package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Transaction struct {
	Model       `bson:",inline"`
	UID         primitive.ObjectID `json:"uid" bson:"uid"`
	Type        string             `json:"type" bson:"type"`
	Status      string             `json:"status" bson:"status"`
	Amount      float64            `json:"amount" bson:"amount"`
	Currency    string             `json:"currency" bson:"currency"`
	Date        primitive.DateTime `json:"date" bson:"date"`
	Description *string            `json:"description,omitempty" bson:"description,omitempty"`
}

func (i *Transaction) Collection() string {
	return "transactions"
}

func (i *Transaction) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Transaction
	return bson.Marshal((*t)(i))
}

func (i *Transaction) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "type", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "date", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
	}
}
