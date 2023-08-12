package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Event struct {
	Model       `bson:",inline"`
	Owner       primitive.ObjectID   `bson:"owner,omitempty" json:"owner,omitempty"`
	Type        string               `bson:"type,omitempty" json:"type,omitempty"`
	Locale      string               `json:"locale,omitempty" bson:"locale,omitempty"`
	Title       string               `json:"title,omitempty" bson:"title,omitempty"`
	Description string               `json:"description,omitempty" bson:"description,omitempty"`
	Location    string               `json:"location,omitempty" bson:"location,omitempty"`
	Start       primitive.Timestamp  `json:"start,omitempty" bson:"start,omitempty"`
	End         primitive.Timestamp  `json:"end,omitempty" bson:"end,omitempty"`
	Timezone    string               `json:"timezone,omitempty" bson:"timezone,omitempty"`
	AllDay      bool                 `json:"all_day,omitempty" bson:"all_day,omitempty"`
	Color       string               `json:"color,omitempty" bson:"color,omitempty"`
	ShowAs      string               `json:"show_as,omitempty" bson:"show_as,omitempty"`
	Status      string               `json:"status,omitempty" bson:"status,omitempty"`
	Reminders   []Reminder           `json:"reminders,omitempty" bson:"reminders,omitempty"`
	Attendees   []primitive.ObjectID `json:"attendees,omitempty" bson:"attendees,omitempty"`
	Recurrence  Recurrence           `json:"recurrence,omitempty" bson:"recurrence,omitempty"`
}

type Recurrence struct {
	Frequency  string                `json:"frequency"`
	Interval   int                   `json:"interval"`
	EndDate    primitive.Timestamp   `json:"end_date,omitempty" bson:"end_date,omitempty"`
	Exceptions []primitive.Timestamp `json:"exceptions,omitempty" bson:"exceptions,omitempty"`
}

func (i *Event) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Event
	return bson.Marshal((*t)(i))
}

func (i *Event) Collection() string {
	return "events"
}

func (i *Event) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "owner", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "type", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "locale", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "start", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "end", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "all_day", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "deleted_at", Value: 1}}, Options: options.Index()},
	}
}
