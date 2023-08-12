package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Quote struct {
	Model
	User        primitive.ObjectID  `json:"user,omitempty" bson:"user,omitempty"`
	Contact     primitive.ObjectID  `json:"contacts,omitempty" bson:"contacts,omitempty"`
	Reference   string              `json:"reference,omitempty" bson:"reference,omitempty"`
	Purchase    string              `json:"purchase,omitempty" bson:"purchase,omitempty"`
	Locale      string              `json:"locale,omitempty" bson:"locale,omitempty"`
	Name        string              `json:"name,omitempty" bson:"name,omitempty"`
	Description string              `json:"description,omitempty" bson:"description,omitempty"`
	Template    string              `json:"template,omitempty" bson:"template,omitempty"`
	ValidUntil  primitive.Timestamp `json:"valid_until,omitempty" bson:"valid_until,omitempty"`
	Amount      float64             `json:"amount,omitempty" bson:"amount,omitempty"`
	Currency    string              `json:"currency,omitempty" bson:"currency,omitempty"`
	Terms       string              `json:"terms,omitempty" bson:"terms,omitempty"`
	Payment     string              `json:"payment,omitempty" bson:"payment,omitempty"`
	Notes       string              `json:"notes,omitempty" bson:"notes,omitempty"`
	Billing     primitive.M         `json:"billing,omitempty" bson:"billing,omitempty"`
	Status      string              `json:"status,omitempty" bson:"status,omitempty"`
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
