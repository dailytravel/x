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
	Code        string               `json:"code,omitempty" bson:"code,omitempty"`
	Locale      string               `json:"locale,omitempty" bson:"locale,omitempty"`
	Description string               `json:"description,omitempty" bson:"description,omitempty"`
	Type        string               `json:"type,omitempty" bson:"type,omitempty"`
	Amount      float64              `json:"amount,omitempty" bson:"amount,omitempty"`
	MaxUses     int                  `json:"max_uses,omitempty" bson:"max_uses,omitempty"`
	MaxDiscount float64              `json:"max_discount,omitempty" bson:"max_discount,omitempty"`
	MinPurchase float64              `json:"min_purchase,omitempty" bson:"min_purchase,omitempty"`
	Currency    string               `json:"currency,omitempty" bson:"currency,omitempty"`
	Products    []primitive.ObjectID `json:"products,omitempty" bson:"products,omitempty"`
	Uses        int                  `json:"uses,omitempty" bson:"uses,omitempty"`
	StartDate   primitive.Timestamp  `json:"start_date,omitempty" bson:"start_date,omitempty"`
	EndDate     primitive.Timestamp  `json:"end_date,omitempty" bson:"end_date,omitempty"`
	Status      string               `json:"status,omitempty" bson:"status,omitempty"`
}

func (i *Coupon) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Coupon
	return bson.Marshal((*t)(i))
}

func (i *Coupon) Collection() string {
	return "coupons"
}

func (i *Coupon) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "code", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "type", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "start_date", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "end_date", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "deleted_at", Value: 1}}, Options: options.Index()},
	}
}
