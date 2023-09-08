package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Comment struct {
	Model       `bson:",inline"`
	UID         *primitive.ObjectID `bson:"uid" json:"uid"`
	Parent      *primitive.ObjectID `json:"parent,omitempty" bson:"parent,omitempty"`
	Commentable Commentable         `json:"commentable" bson:"commentable"`
	Name        *string             `json:"name,omitempty" bson:"name,omitempty"`
	Email       *string             `json:"email,omitempty" bson:"email,omitempty"`
	Locale      string              `json:"locale,omitempty" bson:"locale,omitempty"`
	Body        primitive.M         `json:"body,omitempty" bson:"body,omitempty"`
	Rating      int                 `json:"rating,omitempty" bson:"rating,omitempty"`
	Status      string              `json:"status" bson:"status"`
}

func (Comment) IsEntity() {}

type Commentable struct {
	ID   primitive.ObjectID `json:"id" bson:"id"`
	Type string             `json:"type" bson:"type"`
}

func (i *Comment) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Comment
	return bson.Marshal((*t)(i))
}

func (i *Comment) Collection() string {
	return "comments"
}

func (i *Comment) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "commentable._id", Value: 1}, {Key: "commentable.type", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "uid", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "parent", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "deleted", Value: 1}}, Options: options.Index()},
	}
}
