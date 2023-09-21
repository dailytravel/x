package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Package struct {
	Model       `bson:",inline"`
	Product     primitive.ObjectID   `json:"product" bson:"product"`
	Locale      string               `json:"locale" bson:"locale"`
	Name        primitive.M          `json:"name" bson:"name"`
	Description primitive.M          `json:"description" bson:"description"`
	Details     []primitive.ObjectID `json:"details" bson:"details"`
	Includes    []primitive.ObjectID `json:"includes" bson:"includes"`
	Excludes    []primitive.ObjectID `json:"excludes" bson:"excludes"`
	Price       float64              `json:"price" bson:"price"`
	Discount    float64              `json:"discount,omitempty" bson:"discount,omitempty"`
	Currency    string               `json:"currency" bson:"currency"`
	Status      string               `json:"status" bson:"status"`
}

func (i *Package) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Package
	return bson.Marshal((*t)(i))
}

func (i *Package) Collection() string {
	return "packages"
}

func (i *Package) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "product", Value: 1}}},
		{Keys: bson.D{{Key: "locale", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
	}
}
