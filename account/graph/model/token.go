package model

import (
	"html"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Token struct {
	Model      `bson:",inline"`
	UID        primitive.ObjectID   `json:"uid" bson:"uid"`
	Name       string               `json:"name" bson:"name"`
	Token      string               `json:"token" bson:"token"`
	Abilities  []*string            `json:"abilities" bson:"abilities"`
	ExpiresAt  *primitive.Timestamp `json:"expires_at,omitempty" bson:"expires_at,omitempty"`
	LastUsedAt *primitive.Timestamp `json:"last_used_at,omitempty" bson:"last_used_at,omitempty"`
	Status     string               `json:"status" bson:"status"`
}

// remove spacing characters
func (i *Token) Santize(s string) string {
	return html.EscapeString(strings.TrimSpace(s))
}

func (i *Token) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Token
	return bson.Marshal((*t)(i))
}

func (i *Token) Collection() string {
	return "tokens"
}

func (i *Token) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "token", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
	}
}
