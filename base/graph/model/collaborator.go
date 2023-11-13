package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Collaborator struct {
	Model      `bson:",inline"`
	UID        primitive.ObjectID `json:"uid" bson:"uid"`
	Board      primitive.ObjectID `json:"board,omitempty" bson:"board,omitempty"`
	Permission string             `json:"permission" bson:"permission"`
	Status     string             `json:"status" bson:"status"`
}

func (Collaborator) IsEntity() {}

func (i *Collaborator) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Collaborator
	return bson.Marshal((*t)(i))
}

func (i *Collaborator) Collection() string {
	return "collaborators"
}

func (i *Collaborator) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}}},
		{Keys: bson.D{{Key: "board", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
		{Keys: bson.D{{Key: "deleted", Value: 1}}},
	}
}
