package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Variant struct {
	Model       `bson:",inline"`
	Package     primitive.ObjectID `json:"package" bson:"package"`
	Locale      string             `json:"locale" bson:"locale"`
	Sku         string             `json:"sku,omitempty" bson:"sku,omitempty"`
	Name        primitive.M        `json:"name" bson:"name"`
	Description primitive.M        `json:"description,omitempty" bson:"description,omitempty"`
	Quantity    int                `json:"quantity" bson:"quantity"`
	Price       float64            `json:"price" bson:"price"`
	Discount    float64            `json:"discount,omitempty" bson:"discount,omitempty"`
	Currency    string             `json:"currency" bson:"currency"`
	StartDate   primitive.DateTime `json:"start_date,omitempty" bson:"start_date,omitempty"`
	EndDate     primitive.DateTime `json:"end_date,omitempty" bson:"end_date,omitempty"`
	Status      string             `json:"status" bson:"status"`
}

func (i *Variant) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Variant
	return bson.Marshal((*t)(i))
}

func (i *Variant) Collection() string {
	return "variants"
}

func (i *Variant) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "sku", Value: 1}}, Options: options.Index().SetUnique(true).SetSparse(true)},
		{Keys: bson.D{{Key: "content", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "locale", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
	}
}
