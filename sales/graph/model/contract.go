package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Contract struct {
	Model       `bson:",inline"`
	User        primitive.ObjectID `json:"user" bson:"user"`
	Contact     primitive.ObjectID `json:"contact" bson:"contact"`
	Reference   string             `json:"reference" bson:"reference"`
	Description string             `json:"description" bson:"description"`
	Amount      float64            `json:"amount" bson:"amount"`
	Currency    string             `json:"currency" bson:"currency"`
	StartDate   primitive.DateTime `json:"start_date" bson:"start_date"`
	EndDate     primitive.DateTime `json:"end_date" bson:"end_date"`
	AutoRenew   bool               `json:"auto_renew" bson:"auto_renew"`
	Metadata    primitive.M        `json:"metadata,omitempty" bson:"metadata,omitempty"`
	Status      string             `json:"status" bson:"status"`
}

func (i *Contract) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Contract
	return bson.Marshal((*t)(i))
}

func (i *Contract) Collection() string {
	return "contracts"
}

func (i *Contract) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "reference", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "user", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "contact", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "deleted_at", Value: 1}}, Options: options.Index()},
	}
}
