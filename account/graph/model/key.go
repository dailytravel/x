package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Key struct {
	Model       `bson:",inline"`
	Owner       *primitive.ObjectID  `json:"owner,omitempty" bson:"owner,omitempty"`
	Name        string               `json:"name" bson:"name"`
	Type        string               `json:"type" bson:"type"`
	Provider    string               `json:"provider" bson:"provider"`
	Kid         string               `json:"kid" bson:"kid"`
	Certificate string               `json:"certificate" bson:"certificate"`
	Fingerprint string               `json:"fingerprint" bson:"fingerprint"`
	Thumbprint  string               `json:"thumbprint" bson:"thumbprint"`
	ExpiresAt   *primitive.Timestamp `json:"expires_at,omitempty" bson:"expires_at,omitempty"`
	Status      string               `json:"status" bson:"status"`
}

func (i *Key) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Key
	return bson.Marshal((*t)(i))
}

func (i *Key) Collection() string {
	return KeyCollection
}

func (i *Key) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "provider", Value: 1}, {Key: "kid", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "type", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
	}
}
