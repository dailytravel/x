package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Balance struct {
	Model     `bson:",inline"`
	UID       primitive.ObjectID   `json:"uid" bson:"uid"`
	Type      string               `json:"type" bson:"type"`
	Points    int                  `json:"points" bson:"points"`
	Notes     *string              `json:"notes,omitempty" bson:"notes,omitempty"`
	ExpiredAt *primitive.Timestamp `json:"expired_at,omitempty" bson:"expired_at,omitempty"`
	Status    string               `json:"status" bson:"status"`
}

func (i *Balance) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Balance
	return bson.Marshal((*t)(i))
}

func (i *Balance) Collection() string {
	return "balances"
}

func (i *Balance) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "expired_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "deleted_at", Value: 1}}, Options: options.Index()},
	}
}
