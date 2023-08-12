package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Invoice struct {
	Model     `bson:",inline"`
	User      primitive.ObjectID  `json:"user" bson:"user"`
	Reference string              `json:"reference" bson:"reference"`
	Template  string              `json:"template" bson:"template"`
	Amount    float64             `json:"amount" bson:"amount"`
	Currency  string              `json:"currency" bson:"currency"`
	DueDate   primitive.Timestamp `json:"due_date" bson:"due_date"`
	Billing   primitive.M         `json:"billing" bson:"billing"`
	Notes     string              `json:"notes,omitempty" bson:"notes,omitempty"`
	Status    string              `json:"status" bson:"status"`
}

func (Invoice) IsEntity() {}

func (i *Invoice) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Invoice
	return bson.Marshal((*t)(i))
}

func (i *Invoice) Collection() string {
	return "invoices"
}

func (i *Invoice) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "reference", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "deleted_at", Value: 1}}, Options: options.Index()},
	}
}
