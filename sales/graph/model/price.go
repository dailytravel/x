package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Price struct {
	Model     `bson:",inline"`
	Product   primitive.ObjectID   `json:"product" bson:"product"`
	StartDate *primitive.Timestamp `json:"start_date,omitempty" bson:"start_date,omitempty"`
	EndDate   *primitive.Timestamp `json:"end_date,omitempty" bson:"end_date,omitempty"`
	Regular   float64              `json:"regular,omitempty" bson:"regular,omitempty"`
	Sale      float64              `json:"sale,omitempty" bson:"sale,omitempty"`
}

func (i *Price) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Price
	return bson.Marshal((*t)(i))
}

func (i *Price) Collection() string {
	return "prices"
}

func (i *Price) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "product", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "start_date", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "end_date", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
	}
}
