package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Activity struct {
	Model       `bson:",inline"`
	User        primitive.ObjectID `json:"user" bson:"user"`
	Activitable Activitable        `json:"activitable" bson:"activitable"`
	Action      string             `json:"action" bson:"action"`
	Metadata    primitive.M        `json:"metadata,omitempty" bson:"metadata,omitempty"`
	Status      string             `json:"status" bson:"status"`
}

type Activitable struct {
	ID   primitive.ObjectID `json:"id" bson:"id"`
	Type string             `json:"type" bson:"type"`
}

func (i *Activity) Collection() string {
	return "activities"
}

func (i *Activity) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Activity
	return bson.Marshal((*t)(i))
}
