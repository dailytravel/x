package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Voucher struct {
	Model       `bson:",inline"`
	Locale      string             `json:"locale" bson:"locale"`
	Code        string             `json:"code" bson:"code"`
	Type        string             `json:"type" bson:"type"`
	Name        primitive.M        `json:"name" bson:"name"`
	Description primitive.M        `json:"description,omitempty" bson:"description,omitempty"`
	Start       primitive.DateTime `json:"start,omitempty" bson:"start,omitempty"`
	End         primitive.DateTime `json:"end,omitempty" bson:"end,omitempty"`
	Price       float64            `json:"price" bson:"price"`
	Discount    float64            `json:"discount,omitempty" bson:"discount,omitempty"`
	Currency    string             `json:"currency" bson:"currency"`
	Status      string             `json:"status" bson:"status"`
	Package     primitive.ObjectID `json:"package" bson:"package"`
	Supplier    primitive.ObjectID `json:"supplier" bson:"supplier"`
}

func (i *Voucher) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Voucher
	return bson.Marshal((*t)(i))
}

func (i *Voucher) Collection() string {
	return "vouchers"
}

func (i *Voucher) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "sku", Value: 1}}, Options: options.Index().SetUnique(true).SetSparse(true)},
		{Keys: bson.D{{Key: "type", Value: 1}}},
		{Keys: bson.D{{Key: "start", Value: 1}}},
		{Keys: bson.D{{Key: "end", Value: 1}}},
		{Keys: bson.D{{Key: "package", Value: 1}}},
		{Keys: bson.D{{Key: "supplier", Value: 1}}},
		{Keys: bson.D{{Key: "locale", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
	}
}
