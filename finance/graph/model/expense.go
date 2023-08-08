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
	Owner       primitive.ObjectID `json:"owner" bson:"owner"`
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

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Expense
	return bson.Marshal((*t)(i))
}

func (i *Expense) Collection() string {
	return "expenses"
}

func (i *Expense) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "reference", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "owner", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "type", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_by", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_by", Value: 1}}, Options: options.Index()},
	}
}
