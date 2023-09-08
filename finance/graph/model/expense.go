package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Expense struct {
	Model       `bson:",inline"`
	User        primitive.ObjectID `json:"user" bson:"user"`
	Type        string             `json:"type" bson:"type"`
	Reference   string             `json:"reference" bson:"reference"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Amount      float64            `json:"amount" bson:"amount"`
	Currency    string             `json:"currency" bson:"currency"`
	Date        primitive.DateTime `json:"date" bson:"date"`
	Notes       string             `json:"notes,omitempty" bson:"notes,omitempty"`
	Status      string             `json:"status" bson:"status"`
}

func (i *Expense) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Expense
	return bson.Marshal((*t)(i))
}

func (i *Expense) Collection() string {
	return "expenses"
}

func (i *Expense) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "reference", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "user", Value: 1}}},
		{Keys: bson.D{{Key: "type", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
	}
}
