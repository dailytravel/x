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

type Country struct {
	Model     `bson:",inline"`
	Code      string      `json:"code" bson:"code"`
	Locale    string      `json:"locale" bson:"locale"`
	Name      primitive.M `json:"name" bson:"name"`
	Dialing   string      `json:"dialing,omitempty" bson:"dialing,omitempty"`
	Continent string      `json:"continent,omitempty" bson:"continent,omitempty"`
	Currency  string      `json:"currency,omitempty" bson:"currency,omitempty"`
	Languages []string    `json:"languages,omitempty" bson:"languages,omitempty"`
	Capital   string      `json:"capital,omitempty" bson:"capital,omitempty"`
	Flag      string      `json:"flag,omitempty" bson:"flag,omitempty"`
}

func (i *Country) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Country
	return bson.Marshal((*t)(i))
}

func (i *Country) Collection() string {
	return "countries"
}

func (i *Country) Index() []mongo.IndexModel {
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
		{Keys: bson.D{{Key: "dialing", Value: 1}}},
		{Keys: bson.D{{Key: "continent", Value: 1}}},
		{Keys: bson.D{{Key: "currency", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
	}
}
