package model

import (
	"fmt"
	"os"
	"strings"
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

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Currency
	return bson.Marshal((*t)(i))
}

func (i *Currency) Collection() string {
	return "currencies"
}

func (i *Currency) Index() []mongo.IndexModel {
	locales := strings.Split(os.Getenv("LOCALES"), ",") // example set of locales

	// Dynamically construct the keys for the index based on the provided locales
	var keys bson.D
	for _, locale := range locales {
		key := fmt.Sprintf("name.%s", locale)
		keys = append(keys, bson.E{Key: key, Value: "text"})
	}

	// Construct weights, if needed
	var weights bson.D
	for _, locale := range locales {
		key := fmt.Sprintf("name.%s", locale)
		weights = append(weights, bson.E{Key: key, Value: 1})
	}

	// Base indices for other fields
	return []mongo.IndexModel{
		{Keys: keys, Options: options.Index().SetWeights(weights)},
		{Keys: bson.D{{Key: "code", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "locale", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
	}
}
