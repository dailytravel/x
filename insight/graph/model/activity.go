package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Activity struct {
	Model    `bson:",inline"`
	UID      primitive.ObjectID `json:"uid" bson:"uid"`
	Target   primitive.ObjectID `json:"target" bson:"target"`
	Action   string             `json:"action" bson:"action"`
	Metadata primitive.M        `json:"metadata,omitempty" bson:"metadata,omitempty"`
}

type Activitable struct {
	ID   primitive.ObjectID `json:"id" bson:"id"`
	Type string             `json:"type" bson:"type"`
}

func (i *Activity) Collection() string {
	return "activities"
}

func (i *Activity) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Activity
	return bson.Marshal((*t)(i))
}

func (i *Activity) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "target", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "action", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
	}
}
