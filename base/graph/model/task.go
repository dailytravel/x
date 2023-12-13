package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Task struct {
	Model     `bson:",inline"`
	UID       primitive.ObjectID    `json:"uid" bson:"uid"`
	Assignee  *primitive.ObjectID   `json:"assignee,omitempty" bson:"assignee,omitempty"`
	Parent    *primitive.ObjectID   `json:"parent,omitempty" bson:"parent,omitempty"`
	List      *primitive.ObjectID   `json:"list,omitempty" bson:"list,omitempty"`
	Name      string                `json:"name" bson:"name"`
	Notes     *string               `json:"notes,omitempty" bson:"notes,omitempty"`
	Priority  *string               `json:"priority,omitempty" bson:"priority,omitempty"`
	Start     *primitive.DateTime   `json:"start" bson:"start"`
	End       *primitive.DateTime   `json:"end,omitempty" bson:"end,omitempty"`
	Order     *int                  `json:"order,omitempty" bson:"order,omitempty"`
	Completed bool                  `json:"completed,omitempty" bson:"completed,omitempty"`
	Status    string                `json:"status" bson:"status"`
	Shares    []*primitive.ObjectID `json:"shares,omitempty" bson:"shares,omitempty"`
}

func (Task) IsEntity() {}

func (i *Task) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Task
	return bson.Marshal((*t)(i))
}

func (i *Task) Collection() string {
	return "tasks"
}

func (i *Task) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "name", Value: "text"}, {Key: "notes", Value: "text"}}, Options: options.Index().SetWeights(bson.M{"name": 2, "notes": 1})},
		{Keys: bson.D{{Key: "uid", Value: 1}}},
		{Keys: bson.D{{Key: "parent", Value: 1}}},
		{Keys: bson.D{{Key: "list", Value: 1}}},
		{Keys: bson.D{{Key: "priority", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "start", Value: 1}}},
		{Keys: bson.D{{Key: "end", Value: 1}}},
		{Keys: bson.D{{Key: "assignee", Value: 1}}},
		{Keys: bson.D{{Key: "completed", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
	}
}
