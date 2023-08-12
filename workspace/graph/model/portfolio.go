package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Portfolio struct {
	Model       `bson:",inline"`
	Name        string              `json:"name" bson:"name"`
	Color       string              `json:"color" bson:"color"`
	Description *string             `json:"description,omitempty" bson:"description,omitempty"`
	Owner       primitive.ObjectID  `json:"owner" bson:"owner"`
	Metadata    primitive.M         `json:"metadata,omitempty" bson:"metadata,omitempty"`
	Status      string              `json:"status" bson:"status"`
	CreatedAt   primitive.Timestamp `json:"created_at" bson:"created_at"`
	UpdatedAt   primitive.Timestamp `json:"updated_at" bson:"updated_at"`
}

func (i *Portfolio) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Portfolio
	return bson.Marshal((*t)(i))
}

func (i *Portfolio) Collection() string {
	return "portfolios"
}

func (i *Portfolio) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "owner", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "deleted_at", Value: 1}}, Options: options.Index()},
	}
}
