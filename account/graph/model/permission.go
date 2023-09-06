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

type Permission struct {
	Model       `bson:",inline"`
	Api         primitive.ObjectID `json:"api" bson:"api"`
	Name        string             `json:"name" bson:"name"`
	Description *string            `json:"description,omitempty" bson:"description,omitempty"`
}

// remove spacing characters
func (i *Permission) Santize(s string) string {
	return html.EscapeString(strings.TrimSpace(s))
}

func (i *Permission) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Permission
	return bson.Marshal((*t)(i))
}

func (i *Permission) Collection() string {
	return "permissions"
}

func (i *Permission) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "api", Value: 1}, {Key: "name", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "created_at", Value: 1}}},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}},
	}
}
