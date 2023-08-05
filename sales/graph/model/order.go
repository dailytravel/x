package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Order struct {
	Model       `bson:",inline"`
	Owner       primitive.ObjectID  `bson:"owner,omitempty" json:"owner,omitempty"`
	Type        string              `json:"type,omitempty" bson:"type,omitempty"`
	Order       string              `json:"locale,omitempty" bson:"locale,omitempty"`
	Reference   string              `json:"reference,omitempty" bson:"reference,omitempty"`
	Amount      float64             `json:"amount,omitempty" bson:"amount,omitempty"`
	Currency    string              `json:"currency,omitempty" bson:"currency,omitempty"`
	Coupon      string              `json:"coupon,omitempty" bson:"coupon,omitempty"`
	Cancellable bool                `json:"cancellable,omitempty" bson:"cancellable,omitempty"`
	CancelledAt primitive.Timestamp `json:"cancelled_at,omitempty" bson:"cancelled_at,omitempty"`
	Status      string              `json:"status,omitempty" bson:"status,omitempty"`
}

func (i *Order) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Order
	return bson.Marshal((*t)(i))
}

func (i *Order) Collection() string {
	return "orders"
}

func (i *Order) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "reference", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "owner", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "type", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "locale", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "coupon", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "cancellable", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "cancelled_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "deleted_at", Value: 1}}, Options: options.Index()},
	}
}
