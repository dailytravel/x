package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Organization struct {
	Model       `bson:",inline"`
	UID         *primitive.ObjectID `json:"uid,omitempty" bson:"uid,omitempty"`
	Parent      *primitive.ObjectID `json:"parent,omitempty" bson:"parent,omitempty"`
	Type        string              `json:"type" bson:"type"`
	Name        string              `json:"name" bson:"name"`
	Description *string             `json:"description,omitempty" bson:"description,omitempty"`
	Status      string              `json:"status" bson:"status"`
}

func (Organization) IsEntity() {}

func (i *Organization) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Organization
	return bson.Marshal((*t)(i))
}

func (i *Organization) Collection() string {
	return "organizations"
}

func (i *Organization) Sanitize(s string) string {
	return s
}

func (i *Organization) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "parent", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "type", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "deleted_at", Value: 1}}, Options: options.Index()},
	}
}
