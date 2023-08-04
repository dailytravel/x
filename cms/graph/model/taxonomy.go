package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Taxonomy struct {
	Model         `bson:",inline"`
	Category      primitive.ObjectID `json:"category" bson:"category"`
	Taxonomizable Taxonomizable      `json:"taxonomizable" bson:"taxonomizable"`
	Sticky        bool               `json:"sticky,omitempty" bson:"sticky,omitempty"`
}

type Taxonomizable struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Type string             `json:"type" bson:"type"`
}

func (i *Taxonomy) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Taxonomy
	return bson.Marshal((*t)(i))
}

func (i *Taxonomy) Collection() string {
	return "taxonomies"
}

func (i *Taxonomy) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "taxonomizable._id", Value: 1}, {Key: "taxonomizable.type", Value: 1}, {Key: "category", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "sticky", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
	}
}
