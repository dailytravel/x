package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Point struct {
	Model     `bson:",inline"`
	UID       primitive.ObjectID  `json:"uid" bson:"uid"`
	Target    Target              `json:"target" bson:"target"`
	Points    int                 `json:"points" bson:"points"`
	Type      string              `json:"type" bson:"type"`
	Metadata  primitive.M         `json:"metadata,omitempty" bson:"metadata,omitempty"`
	ExpiresAt primitive.Timestamp `json:"expires_at" bson:"expires_at"`
}

type Target struct {
	ID   primitive.ObjectID `json:"id" bson:"id"`
	Type PointType          `json:"type" bson:"type"`
}

func (i *Point) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Point
	return bson.Marshal((*t)(i))
}

func (i *Point) Collection() string {
	return "points"
}

func (i *Point) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "target._id", Value: 1}, {Key: "target.type", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "deleted_at", Value: 1}}, Options: options.Index()},
	}
}
