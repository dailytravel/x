package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Item struct {
	Model       `bson:",inline"`
	UID         primitive.ObjectID `bson:"uid" json:"uid"`
	Package     primitive.ObjectID `bson:"package" json:"package"`
	Locale      string             `json:"locale" bson:"locale"`
	Code        string             `json:"code" bson:"code"`
	Type        string             `json:"type" bson:"type"`
	Name        primitive.M        `json:"name" bson:"name"`
	Description primitive.M        `json:"description,omitempty" bson:"description,omitempty"`
	Start       primitive.DateTime `json:"start,omitempty" bson:"start,omitempty"`
	End         primitive.DateTime `json:"end,omitempty" bson:"end,omitempty"`
	Price       float64            `json:"price" bson:"price"`
	Discount    *float64           `json:"discount,omitempty" bson:"discount,omitempty"`
	Currency    string             `json:"currency" bson:"currency"`
}

func (i *Item) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Item
	return bson.Marshal((*t)(i))
}

func (i *Item) Collection() string {
	return "items"
}

func (i *Item) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}}},
		{Keys: bson.D{{Key: "package", Value: 1}}},
		{Keys: bson.D{{Key: "locale", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
	}
}
