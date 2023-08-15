package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Follow struct {
	Model      `bson:",inline"`
	UID        primitive.ObjectID `bson:"uid" json:"uid"`
	Followable Followable         `json:"followable" bson:"followable"`
	Role       string             `json:"role" bson:"role"`
	Muted      bool               `json:"muted" bson:"muted"`
	Status     string             `json:"status" bson:"status"`
}

func (Follow) IsEntity() {}

type Followable struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Type string             `json:"type" bson:"type"`
}

func (i *Follow) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Follow
	return bson.Marshal((*t)(i))
}

func (i *Follow) Collection() string {
	return "follows"
}

func (i *Follow) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}, {Key: "object._id", Value: 1}, {Key: "object.type", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
	}
}
