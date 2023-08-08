package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Leave struct {
	Model     `bson:",inline"`
	Owner     primitive.ObjectID `json:"owner" bson:"owner"`
	Type      string             `json:"type" bson:"type"`
	StartDate primitive.DateTime `json:"start_date" bson:"start_date"`
	EndDate   primitive.DateTime `json:"end_date" bson:"end_date"`
	Reason    string             `json:"reason" bson:"reason"`
	Status    string             `json:"status" bson:"status"`
}

func (i *Leave) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Leave
	return bson.Marshal((*t)(i))
}

func (i *Leave) Collection() string {
	return "leaves"
}

func (i *Leave) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "owner", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "type", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "start_date", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "end_date", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "deleted_at", Value: 1}}, Options: options.Index()},
	}
}
