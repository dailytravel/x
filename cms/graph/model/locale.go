package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Locale struct {
	Model      `bson:",inline"`
	Code       string      `json:"code" bson:"code"`
	Locale     string      `json:"locale" bson:"locale"`
	Name       primitive.M `json:"name" bson:"name"`
	Order      int         `json:"order" bson:"order"`
	Rtl        bool        `json:"rtl,omitempty" bson:"rtl,omitempty"`
	DateFormat string      `json:"date_format,omitempty" bson:"date_format,omitempty"`
	TimeFormat string      `json:"time_format,omitempty" bson:"time_format,omitempty"`
	WeekStart  int         `json:"week_start,omitempty" bson:"week_start,omitempty"`
}

func (i *Locale) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Locale
	return bson.Marshal((*t)(i))
}

func (i *Locale) Collection() string {
	return "locales"
}

func (i *Locale) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "code", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "locale", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
	}
}
