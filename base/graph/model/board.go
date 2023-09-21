package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Board struct {
	Model        `bson:",inline"`
	UID          primitive.ObjectID  `json:"uid" bson:"uid"`
	Organization *primitive.ObjectID `json:"organization,omitempty" bson:"organization,omitempty"`
	Portfolio    *primitive.ObjectID `json:"portfolio,omitempty" bson:"portfolio,omitempty"`
	Type         string              `json:"type" bson:"type"`
	Title        string              `json:"title" bson:"title"`
	Description  *string             `json:"description,omitempty" bson:"description,omitempty"`
	IsTemplate   bool                `json:"is_template,omitempty" bson:"is_template,omitempty"`
	End          *primitive.DateTime `json:"end,omitempty" bson:"end,omitempty"`
	Starred      bool                `json:"starred,omitempty" bson:"starred,omitempty"`
	Background   string              `json:"background,omitempty" bson:"background,omitempty"`
	Order        int                 `json:"order" bson:"order"`
	Status       string              `json:"status" bson:"status"`
}

func (Board) IsEntity() {}

func (i *Board) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Board
	return bson.Marshal((*t)(i))
}

func (i *Board) Collection() string {
	return "boards"
}

func (i *Board) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}}},
		{Keys: bson.D{{Key: "organization", Value: 1}}},
		{Keys: bson.D{{Key: "type", Value: 1}}},
		{Keys: bson.D{{Key: "starred", Value: 1}}},
		{Keys: bson.D{{Key: "order", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "is_template", Value: 1}}},
		{Keys: bson.D{{Key: "end", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
		{Keys: bson.D{{Key: "deleted", Value: 1}}},
	}
}
