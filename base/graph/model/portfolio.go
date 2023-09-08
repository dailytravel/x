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
	UID         primitive.ObjectID `json:"uid" bson:"uid"`
	Name        string             `json:"name" bson:"name"`
	Color       string             `json:"color" bson:"color"`
	Description *string            `json:"description,omitempty" bson:"description,omitempty"`
	Status      string             `json:"status" bson:"status"`
}

func (i *Portfolio) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Portfolio
	return bson.Marshal((*t)(i))
}

func (i *Portfolio) Collection() string {
	return "portfolios"
}

func (i *Portfolio) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "deleted", Value: 1}}, Options: options.Index()},
	}
}
