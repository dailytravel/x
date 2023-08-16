package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Currency struct {
	Model     `bson:",inline"`
	Code      string      `json:"code" bson:"code"`
	Locale    string      `json:"locale" bson:"locale"`
	Name      primitive.M `json:"name" bson:"name"`
	Rate      float64     `json:"rate" bson:"rate"`
	Symbol    string      `json:"symbol,omitempty" bson:"symbol,omitempty"`
	Precision int         `json:"precision,omitempty" bson:"precision,omitempty"`
	Decimal   string      `json:"decimal,omitempty" bson:"decimal,omitempty"`
	Thousand  string      `json:"thousand,omitempty" bson:"thousand,omitempty"`
	Order     int         `json:"order" bson:"order"`
}

func (i *Currency) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Currency
	return bson.Marshal((*t)(i))
}

func (i *Currency) Collection() string {
	return "currencies"
}

func (i *Currency) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "code", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "locale", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
	}
}
