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

type Role struct {
	Model       `bson:",inline"`
	Name        string                `json:"name" bson:"name"`
	Description *string               `json:"description,omitempty" bson:"description,omitempty"`
	Permissions []*primitive.ObjectID `json:"permissions" bson:"permissions"`
}

// remove spacing characters
func (i *Role) Santize(s string) string {
	return html.EscapeString(strings.TrimSpace(s))
}

func (i *Role) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Role
	return bson.Marshal((*t)(i))
}

func (i *Role) Collection() string {
	return "roles"
}

func (i *Role) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "name", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
	}
}
