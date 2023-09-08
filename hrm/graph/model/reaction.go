package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Reaction struct {
	Model     `bson:",inline"`
	UID       primitive.ObjectID `json:"uid" bson:"uid"`
	Action    string             `json:"action" bson:"action"`
	Reactable Reactable          `json:"reactable" bson:"reactable"`
}

func (Reaction) IsEntity() {}

type Reactable struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Type string             `json:"type" bson:"type"`
}

func (i *Reaction) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Reaction
	return bson.Marshal((*t)(i))
}

func (i *Reaction) Collection() string {
	return "reactions"
}

func (i *Reaction) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}, {Key: "reactable._id", Value: 1}, {Key: "reactable.type", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "created", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated", Value: 1}}, Options: options.Index()},
	}
}
