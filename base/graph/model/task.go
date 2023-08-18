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
	UID       primitive.ObjectID  `json:"uid" bson:"uid"`
	Parent    *primitive.ObjectID `json:"parent,omitempty" bson:"parent,omitempty"`
	List      *primitive.ObjectID `json:"list" bson:"list"`
	Name      string              `json:"name" bson:"name"`
	Notes     *string             `json:"notes,omitempty" bson:"notes,omitempty"`
	Priority  *string             `json:"priority,omitempty" bson:"priority,omitempty"`
	StartDate *primitive.DateTime `json:"start_date" bson:"start_date"`
	DueDate   *primitive.DateTime `json:"due_date" bson:"due_date"`
	Order     int                 `json:"order" bson:"order"`
	Status    string              `json:"status" bson:"status"`
	Labels    []string            `json:"labels,omitempty" bson:"labels,omitempty"`
}

func (i *Task) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Task
	return bson.Marshal((*t)(i))
}

func (i *Task) Collection() string {
	return "tasks"
}

func (i *Task) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "name", Value: "text"}, {Key: "notes", Value: "text"}}, Options: options.Index().SetWeights(bson.M{"name": 2, "notes": 1})},
		{Keys: bson.D{{Key: "user", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "parent", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "list", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "priority", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "order", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "due_date", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
	}
}
