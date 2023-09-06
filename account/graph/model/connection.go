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

type Connection struct {
	Model       `bson:",inline"`
	Name        string             `json:"name" bson:"name"`
	Description *string            `json:"description,omitempty" bson:"description,omitempty"`
	Type        string             `json:"type" bson:"type"`
	Client      primitive.ObjectID `json:"client" bson:"client"`
	Status      string             `json:"status" bson:"status"`
}

// remove spacing characters
func (i *Connection) Santize(s string) string {
	return html.EscapeString(strings.TrimSpace(s))
}

func (i *Connection) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Connection
	return bson.Marshal((*t)(i))
}

func (i *Connection) Collection() string {
	return "connections"
}

func (i *Connection) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "name", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "type", Value: 1}}},
		{Keys: bson.D{{Key: "client", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "created_at", Value: 1}}},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}},
	}
}
