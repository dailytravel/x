package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Inventory struct {
	Model    `bson:",inline"`
	Product  primitive.ObjectID  `json:"product,omitempty" bson:"product,omitempty"`
	Date     primitive.Timestamp `json:"date,omitempty" bson:"date,omitempty"`
	Quantity int                 `json:"quantity,omitempty" bson:"quantity,omitempty"`
}

func (i *Inventory) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Inventory
	return bson.Marshal((*t)(i))
}

func (i *Inventory) Collection() string {
	return "inventories"
}

func (i *Inventory) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "product", Value: 1}, {Key: "date", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
	}
}
