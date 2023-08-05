package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Program struct {
	Model       `bson:",inline"`
	Type        string      `bson:"type,optional" json:"type,omitempty"`
	Locale      string      `bson:"locale,omitempty" json:"locale,omitempty"`
	Name        primitive.M `json:"name,omitempty" bson:"name,omitempty"`
	Description primitive.M `bson:"description,omitempty" json:"description,omitempty"`
	Points      int         `json:"points,omitempty" bson:"points,omitempty"`
	Status      string      `json:"status,omitempty" bson:"status,omitempty"`
}

func (i *Program) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Program
	return bson.Marshal((*t)(i))
}

func (i *Program) Collection() string {
	return "programs"
}

func (i *Program) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "type", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "locale", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "deleted_at", Value: 1}}, Options: options.Index()},
	}
}
