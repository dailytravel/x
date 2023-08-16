package model

import (
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

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Country
	return bson.Marshal((*t)(i))
}

func (i *Country) Collection() string {
	return "countries"
}

func (i *Country) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "code", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "name", Value: "text"}}},
		{Keys: bson.D{{Key: "dialing", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "continent", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "currency", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
	}
}
