package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Payment struct {
	Model    `bson:",inline"`
	Amount   float64            `json:"amount" bson:"amount"`
	Currency string             `json:"currency" bson:"currency"`
	Date     primitive.DateTime `json:"date" bson:"date"`
	Status   string             `json:"status" bson:"status"`
	User     primitive.ObjectID `json:"user" bson:"user"`
	Invoice  primitive.ObjectID `json:"invoice" bson:"invoice"`
	Metadata primitive.M        `json:"metadata,omitempty" bson:"metadata,omitempty"`
}

func (i *Payment) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Payment
	return bson.Marshal((*t)(i))
}

func (i *Payment) Collection() string {
	return "payments"
}

func (i *Payment) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "user", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "invoice", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "date", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
	}
}
