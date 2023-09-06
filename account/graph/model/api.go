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

type Api struct {
	Model       `bson:",inline"`
	Name        string  `json:"name" bson:"name"`
	Description *string `json:"description,omitempty" bson:"description,omitempty"`
	Identifier  string  `json:"identifier" bson:"identifier"`
	Algorithm   string  `json:"algorithm" bson:"algorithm"`
	Expiration  int64   `json:"expiration" bson:"expiration"`
	Status      string  `json:"status" bson:"status"`
}

// remove spacing characters
func (i *Api) Santize(s string) string {
	return html.EscapeString(strings.TrimSpace(s))
}

func (i *Api) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Api
	return bson.Marshal((*t)(i))
}

func (i *Api) Collection() string {
	return "apis"
}

func (i *Api) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "identifier", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "created_at", Value: 1}}},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}},
	}
}
