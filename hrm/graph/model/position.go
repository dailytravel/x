package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Position struct {
	Model        `bson:",inline"`
	Organization primitive.ObjectID `json:"organization" bson:"organization"`
	Locale       string             `json:"locale" bson:"locale"`
	Title        primitive.M        `json:"title" bson:"title"`
	Description  primitive.M        `json:"description,omitempty" bson:"description,omitempty"`
	Salary       string             `json:"salary" bson:"salary"`
	Status       string             `json:"status" bson:"status"`
}

func (i *Position) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Position
	return bson.Marshal((*t)(i))
}

func (i *Position) Collection() string {
	return "positions"
}

func (i *Position) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "organization", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "locale", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "deleted_at", Value: 1}}, Options: options.Index()},
	}
}
