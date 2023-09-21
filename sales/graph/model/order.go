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
	UID         primitive.ObjectID `bson:"uid" json:"uid"`
	Locale      string             `json:"locale" bson:"locale"`
	Type        string             `json:"type" bson:"type"`
	Code        string             `json:"code" bson:"code"`
	Coupon      *string            `json:"coupon,omitempty" bson:"coupon,omitempty"`
	Cancellable *bool              `json:"cancellable" bson:"cancellable"`
	Payment     string             `json:"payment" bson:"payment"`
	Status      string             `json:"status" bson:"status"`
}

func (i *Order) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Order
	return bson.Marshal((*t)(i))
}

func (i *Order) Collection() string {
	return "orders"
}

func (i *Order) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "code", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "uid", Value: 1}}},
		{Keys: bson.D{{Key: "type", Value: 1}}},
		{Keys: bson.D{{Key: "locale", Value: 1}}},
		{Keys: bson.D{{Key: "coupon", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "cancellable", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
		{Keys: bson.D{{Key: "deleted", Value: 1}}},
	}
}
