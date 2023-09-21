package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Attendance struct {
	Model   `bson:",inline"`
	UID     primitive.ObjectID   `json:"uid" bson:"uid"`
	Date    primitive.DateTime   `json:"date" bson:"date"`
	TimeIn  *primitive.Timestamp `json:"time_in,omitempty" bson:"time_in,omitempty"`
	TimeOut *primitive.Timestamp `json:"time_out,omitempty" bson:"time_out,omitempty"`
	Notes   string               `json:"notes,omitempty" bson:"notes,omitempty"`
	Status  string               `json:"status" bson:"status"`
}

func (i *Attendance) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Attendance
	return bson.Marshal((*t)(i))
}

func (i *Attendance) Collection() string {
	return "attendances"
}

func (i *Attendance) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}}},
		{Keys: bson.D{{Key: "date", Value: 1}}},
		{Keys: bson.D{{Key: "time_in", Value: 1}}},
		{Keys: bson.D{{Key: "time_out", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
	}
}
