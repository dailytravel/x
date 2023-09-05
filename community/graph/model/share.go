package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Share struct {
	Model      `bson:",inline"`
	UID        primitive.ObjectID `bson:"uid" json:"uid"`
	Shareable  Shareable          `json:"shareable" bson:"shareable"`
	Permission string             `json:"permission" bson:"permission"`
	Status     string             `json:"status" bson:"status"`
}

func (Share) IsEntity() {}

type Shareable struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Type string             `json:"type" bson:"type"`
}

func (i *Share) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Share
	return bson.Marshal((*t)(i))
}

func (i *Share) Collection() string {
	return "shares"
}

func (i *Share) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}, {Key: "shareable._id", Value: 1}, {Key: "shareable.type", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_by", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_by", Value: 1}}, Options: options.Index()},
	}
}
