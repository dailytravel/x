package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type List struct {
	Model `bson:",inline"`
	UID   primitive.ObjectID   `json:"uid" bson:"uid"`
	Board primitive.ObjectID   `json:"board" bson:"board"`
	Name  string               `json:"name" bson:"name"`
	Order int                  `json:"order" bson:"order"`
	Tasks []primitive.ObjectID `json:"tasks,omitempty" bson:"tasks,omitempty"`
}

func (i *List) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t List
	return bson.Marshal((*t)(i))
}

func (i *List) Collection() string {
	return "lists"
}

func (i *List) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "board", Value: 1}}},
		{Keys: bson.D{{Key: "tasks", Value: 1}}},
		{Keys: bson.D{{Key: "name", Value: "text"}}},
		{Keys: bson.D{{Key: "order", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
		{Keys: bson.D{{Key: "deleted", Value: 1}}},
	}
}
