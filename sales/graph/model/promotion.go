package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Promotion struct {
	Model       `bson:",inline"`
	Type        string      `bson:"type,optional" json:"type,omitempty"`
	Locale      string      `bson:"locale,omitempty" json:"locale,omitempty"`
	Name        primitive.M `json:"name,omitempty" bson:"name,omitempty"`
	Description primitive.M `bson:"description,omitempty" json:"description,omitempty"`
	Credits     int         `json:"credits,omitempty" bson:"credits,omitempty"`
	Status      string      `json:"status,omitempty" bson:"status,omitempty"`
}

func (i *Promotion) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Promotion
	return bson.Marshal((*t)(i))
}

func (i *Promotion) Collection() string {
	return "promotions"
}

func (i *Promotion) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "type", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "locale", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
		{Keys: bson.D{{Key: "deleted", Value: 1}}},
	}
}
