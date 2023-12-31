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

type Integration struct {
	Model       `bson:",inline"`
	Type        string  `json:"type" bson:"type"`
	Name        string  `json:"name" bson:"name"`
	Description *string `json:"description,omitempty" bson:"description,omitempty"`
	Status      string  `json:"status" bson:"status"`
}

// remove spacing characters
func (i *Integration) Santize(s string) string {
	return html.EscapeString(strings.TrimSpace(s))
}

func (i *Integration) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Integration
	return bson.Marshal((*t)(i))
}

func (i *Integration) Collection() string {
	return "integrations"
}

func (i *Integration) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "name", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
	}
}
