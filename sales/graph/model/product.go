package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Product struct {
	Model       `bson:",inline"`
	Content     primitive.ObjectID `bson:"content,omitempty" json:"content,omitempty"`
	Locale      string             `json:"locale,omitempty" bson:"locale,omitempty"`
	Type        string             `json:"type,omitempty" bson:"type,omitempty"`
	Sku         string             `json:"sku,omitempty" bson:"sku,omitempty"`
	Name        primitive.M        `json:"name,omitempty" bson:"name,omitempty"`
	Description primitive.M        `json:"description,omitempty" bson:"description,omitempty"`
	Quantity    int                `json:"quantity,omitempty" bson:"quantity,omitempty"`
	Currency    string             `json:"currency,omitempty" bson:"currency,omitempty"`
	Status      string             `json:"status,omitempty" bson:"status,omitempty"`
}

func (i *Product) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Product
	return bson.Marshal((*t)(i))
}

func (i *Product) Collection() string {
	return "products"
}

func (i *Product) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "sku", Value: 1}}, Options: options.Index().SetUnique(true).SetSparse(true)},
		{Keys: bson.D{{Key: "content", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "locale", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
	}
}
