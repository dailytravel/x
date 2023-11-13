package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Relationship struct {
	Model  `bson:",inline"`
	Object Object `json:"object" bson:"object"`
	Target Object `json:"target" bson:"target"`
	Order  *int   `json:"order,omitempty" bson:"order,omitempty"`
}

type Object struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Type string             `json:"type" bson:"type"`
}

func (i *Relationship) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Relationship
	return bson.Marshal((*t)(i))
}

func (i *Relationship) Collection() string {
	return "relationships"
}

func (i *Relationship) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "object._id", Value: 1}, {Key: "object.type", Value: 1}, {Key: "target._id", Value: 1}, {Key: "target.type", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
	}
}
