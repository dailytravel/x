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
	User        primitive.ObjectID `json:"user" bson:"user"`
	Parent      primitive.ObjectID `json:"parent,omitempty" bson:"parent,omitempty"`
	Commentable Commentable        `json:"commentable" bson:"commentable"`
	Locale      string             `json:"locale,omitempty" bson:"locale,omitempty"`
	Body        primitive.M        `json:"body,omitempty" bson:"body,omitempty"`
	Rating      int                `json:"rating,omitempty" bson:"rating,omitempty"`
	Status      string             `json:"status" bson:"status"`
}

func (Comment) IsEntity() {}

type Commentable struct {
	ID   primitive.ObjectID `json:"id" bson:"id"`
	Type string             `json:"type" bson:"type"`
}

func (i *Comment) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Comment
	return bson.Marshal((*t)(i))
}

func (i *Comment) Collection() string {
	return "comments"
}

func (i *Comment) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "user", Value: 1}, {Key: "commentable._id", Value: 1}, {Key: "commentable.type", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "parent", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "deleted_at", Value: 1}}, Options: options.Index()},
	}
}
