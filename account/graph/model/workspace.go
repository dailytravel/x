package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Workspace struct {
	Model       `bson:",inline"`
	UID         *primitive.ObjectID `json:"uid,omitempty" bson:"uid,omitempty"`
	Name        string              `json:"name" bson:"name"`
	Description *string             `json:"description,omitempty" bson:"description,omitempty"`
	Status      string              `json:"status" bson:"status"`
}

func (i *Workspace) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Workspace
	return bson.Marshal((*t)(i))
}

func (i *Workspace) Collection() string {
	return "workspaces"
}

func (i *Workspace) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "name", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "uid", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
	}
}
