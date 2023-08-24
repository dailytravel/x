package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Application struct {
	Model  `bson:",inline"`
	UID    primitive.ObjectID `json:"uid" bson:"uid"`
	Job    primitive.ObjectID `json:"job" bson:"job"`
	Notes  string             `json:"notes" bson:"notes"`
	Status string             `json:"status" bson:"status"`
}

func (i *Application) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Application
	return bson.Marshal((*t)(i))
}

func (i *Application) Collection() string {
	return "applications"
}

func (i *Application) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "job", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "deleted_at", Value: 1}}, Options: options.Index()},
	}
}
