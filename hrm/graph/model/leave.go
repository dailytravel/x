package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Leave struct {
	Model  `bson:",inline"`
	UID    primitive.ObjectID `json:"uid" bson:"uid"`
	Type   string             `json:"type" bson:"type"`
	Start  primitive.DateTime `json:"start" bson:"start"`
	End    primitive.DateTime `json:"end" bson:"end"`
	Reason string             `json:"reason" bson:"reason"`
	Status string             `json:"status" bson:"status"`
}

func (i *Leave) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Leave
	return bson.Marshal((*t)(i))
}

func (i *Leave) Collection() string {
	return "leaves"
}

func (i *Leave) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "employee", Value: 1}}},
		{Keys: bson.D{{Key: "type", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "start", Value: 1}}},
		{Keys: bson.D{{Key: "end", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
		{Keys: bson.D{{Key: "deleted", Value: 1}}},
	}
}
