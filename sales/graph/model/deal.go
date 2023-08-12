package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Deal struct {
	Model       `bson:",inline"`
	User        primitive.ObjectID `json:"user,omitempty" bson:"user,omitempty"`
	Contact     primitive.ObjectID `json:"contact,omitempty" bson:"contact,omitempty"`
	Type        string             `json:"type" bson:"type"`
	Name        string             `json:"name" bson:"name"`
	Amount      *float64           `json:"amount,omitempty" bson:"amount,omitempty"`
	Currency    *string            `json:"currency,omitempty" bson:"currency,omitempty"`
	CloseDate   *string            `json:"close_date,omitempty" bson:"close_date,omitempty"`
	Pipeline    *Pipeline          `json:"pipeline,omitempty" bson:"pipeline,omitempty"`
	Stage       string             `json:"stage,omitempty" bson:"stage,omitempty"`
	Priority    string             `json:"priority,omitempty" bson:"priority,omitempty"`
	Source      string             `json:"source,omitempty" bson:"source,omitempty"`
	LoseReason  *string            `json:"lose_reason,omitempty" bson:"lose_reason,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Order       *int               `json:"order,omitempty" bson:"order,omitempty"`
	Metadata    primitive.M        `json:"metadata,omitempty" bson:"metadata,omitempty"`
	Status      string             `json:"status" bson:"status"`
}

func (i *Deal) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Deal
	return bson.Marshal((*t)(i))
}

func (i *Deal) Collection() string {
	return "Deals"
}

func (i *Deal) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "user", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "contact", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "type", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "close_date", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "pipeline", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "stage", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "priority", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "source", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "deleted_at", Value: 1}}, Options: options.Index()},
	}
}
