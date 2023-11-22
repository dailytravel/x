package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Coupon struct {
	Model       `bson:",inline"`
	UID         *primitive.ObjectID   `json:"uid" bson:"uid"`
	Code        string                `json:"code" bson:"code"`
	Locale      string                `json:"locale" bson:"locale"`
	Name        primitive.M           `json:"name" bson:"name"`
	Description primitive.M           `json:"description" bson:"description"`
	Type        string                `json:"type" bson:"type"`
	Amount      float64               `json:"amount" bson:"amount"`
	MaxUses     *int                  `json:"max_uses,omitempty" bson:"max_uses,omitempty"`
	MaxDiscount *float64              `json:"max_discount,omitempty" bson:"max_discount,omitempty"`
	MinPurchase *float64              `json:"min_purchase,omitempty" bson:"min_purchase,omitempty"`
	Currency    string                `json:"currency" bson:"currency"`
	Uses        *int                  `json:"uses,omitempty" bson:"uses,omitempty"`
	Starts      *primitive.Timestamp  `json:"starts,omitempty" bson:"starts,omitempty"`
	Expires     *primitive.Timestamp  `json:"expires,omitempty" bson:"expires,omitempty"`
	Status      string                `json:"status" bson:"status"`
	Products    []*primitive.ObjectID `json:"products,omitempty" bson:"products,omitempty"`
}

func (i *Coupon) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Coupon
	return bson.Marshal((*t)(i))
}

func (i *Coupon) Collection() string {
	return "coupons"
}

func (i *Coupon) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "code", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "uid", Value: 1}}},
		{Keys: bson.D{{Key: "type", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "starts", Value: 1}}},
		{Keys: bson.D{{Key: "expires", Value: 1}}},
		{Keys: bson.D{{Key: "products", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
	}
}
