package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type List struct {
	Model `bson:",inline"`
	UID   primitive.ObjectID `json:"uid" bson:"uid"`
	Board primitive.ObjectID `json:"board" bson:"board"`
	Name  string             `json:"name" bson:"name"`
	Order int                `json:"order" bson:"order"`
}

func (i *List) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t List
	return bson.Marshal((*t)(i))
}

func (i *List) Collection() string {
	return "lists"
}

func (i *List) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "board", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "name", Value: "text"}}},
		{Keys: bson.D{{Key: "order", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "deleted_at", Value: 1}}, Options: options.Index()},
	}
}
