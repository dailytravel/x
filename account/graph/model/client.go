package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	Model       `bson:",inline"`
	UID         primitive.ObjectID  `json:"uid" bson:"uid"`
	Type        string              `json:"type" bson:"type"`
	Name        string              `json:"name" bson:"name"`
	Description string              `json:"description,omitempty" bson:"description,omitempty"`
	Secret      string              `json:"secret" bson:"secret"`
	Domains     []string            `json:"domains,omitempty" bson:"domains,omitempty"`
	Redirect    string              `json:"redirect,omitempty" bson:"redirect,omitempty"`
	Revoked     bool                `json:"revoked,omitempty" bson:"revoked,omitempty"`
	Provider    string              `json:"provider,omitempty" bson:"provider,omitempty"`
	Permissions []string            `json:"permissions,omitempty" bson:"permissions,omitempty"`
	LastUsed    primitive.Timestamp `json:"last_used,omitempty" bson:"last_used,omitempty"`
	ExpiresAt   primitive.Timestamp `json:"expires_at,omitempty" bson:"expires_at,omitempty"`
	Status      string              `json:"status" bson:"status"`
}

func (i *Client) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Client
	return bson.Marshal((*t)(i))
}

func (i *Client) Collection() string {
	return "clients"
}

func (i *Client) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
	}
}
