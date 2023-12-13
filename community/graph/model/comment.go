package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Comment struct {
	Model      `bson:",inline"`
	UID        *primitive.ObjectID  `bson:"uid" json:"uid"`
	Parent     *primitive.ObjectID  `json:"parent,omitempty" bson:"parent,omitempty"`
	Object     Object               `json:"object" bson:"object"`
	Locale     string               `json:"locale,omitempty" bson:"locale,omitempty"`
	Subject    *string              `json:"subject,omitempty" bson:"subject,omitempty"`
	Body       primitive.M          `json:"body,omitempty" bson:"body,omitempty"`
	Rating     int                  `json:"rating,omitempty" bson:"rating,omitempty"`
	Recommends bool                 `json:"recommends,omitempty" bson:"recommends,omitempty"`
	Status     string               `json:"status" bson:"status"`
	Response   *string              `json:"response,omitempty" bson:"response,omitempty"`
	Responded  *primitive.Timestamp `json:"responded,omitempty" bson:"responded,omitempty"`
}

func (Comment) IsEntity() {}

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
		{Keys: bson.D{{Key: "object._id", Value: 1}, {Key: "object.collection", Value: 1}}},
		{Keys: bson.D{{Key: "uid", Value: 1}}},
		{Keys: bson.D{{Key: "parent", Value: 1}}},
		{Keys: bson.D{{Key: "recommends", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
		{Keys: bson.D{{Key: "deleted", Value: 1}}},
	}
}
