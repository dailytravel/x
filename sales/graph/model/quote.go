package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Quote struct {
	Model
	UID         primitive.ObjectID  `json:"uid" bson:"uid"`
	Contact     primitive.ObjectID  `json:"contacts" bson:"contacts"`
	Reference   string              `json:"reference" bson:"reference"`
	Purchase    *string             `json:"purchase,omitempty" bson:"purchase,omitempty"`
	Locale      string              `json:"locale" bson:"locale"`
	Name        string              `json:"name" bson:"name"`
	Description *string             `json:"description,omitempty" bson:"description,omitempty"`
	Template    string              `json:"template" bson:"template"`
	ValidUntil  primitive.Timestamp `json:"valid_until" bson:"valid_until"`
	Terms       string              `json:"terms" bson:"terms"`
	Payment     string              `json:"payment" bson:"payment"`
	Notes       string              `json:"notes,omitempty" bson:"notes,omitempty"`
	Billing     primitive.M         `json:"billing,omitempty" bson:"billing,omitempty"`
	Status      string              `json:"status" bson:"status"`
}

func (i *Quote) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Quote
	return bson.Marshal((*t)(i))
}

func (i *Quote) Collection() string {
	return "quotes"
}

func (i *Quote) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "contact", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "locale", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "reference", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "purchase", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "deleted_at", Value: 1}}, Options: options.Index()},
	}
}
