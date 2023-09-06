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
	Description primitive.M           `json:"description" bson:"description"`
	Type        string                `json:"type" bson:"type"`
	Amount      float64               `json:"amount" bson:"amount"`
	MaxUses     *int                  `json:"maxUses,omitempty" bson:"max_uses,omitempty"`
	MaxDiscount *float64              `json:"maxDiscount,omitempty" bson:"max_discount,omitempty"`
	MinPurchase *float64              `json:"minPurchase,omitempty" bson:"min_purchase,omitempty"`
	Currency    string                `json:"currency" bson:"currency"`
	Products    []*primitive.ObjectID `json:"products,omitempty" bson:"products,omitempty"`
	Uses        *int                  `json:"uses,omitempty" bson:"uses,omitempty"`
	Start       primitive.DateTime    `json:"start" bson:"start"`
	End         primitive.DateTime    `json:"end" bson:"end"`
	Status      string                `json:"status" bson:"status"`
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
		{Keys: bson.D{{Key: "uid", Value: 1}}},
		{Keys: bson.D{{Key: "type", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "start", Value: 1}}},
		{Keys: bson.D{{Key: "end", Value: 1}}},
		{Keys: bson.D{{Key: "created_at", Value: 1}}},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}},
		{Keys: bson.D{{Key: "deleted_at", Value: 1}}},
	}
}
