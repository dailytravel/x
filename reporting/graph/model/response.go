package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Response struct {
	Model     `bson:",inline"`
	Campaign  primitive.ObjectID `json:"campaign" bson:"campaign"`
	Reference string             `json:"reference" bson:"reference"`
	Type      string             `json:"type" bson:"type"`
	UserAgent string             `json:"user_agent" bson:"user_agent"`
}

func (Response) IsEntity() {}

func (i *Response) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Response
	return bson.Marshal((*t)(i))
}

func (i *Response) Collection() string {
	return "responses"
}

func (i *Response) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "campaign", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "reference", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "type", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
	}
}
