package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Salary struct {
	Model     `bson:",inline"`
	UID       primitive.ObjectID `json:"uid" bson:"uid"`
	Amount    float64            `json:"amount" bson:"amount"`
	Currency  string             `json:"currency" bson:"currency"`
	StartDate primitive.DateTime `json:"start_date" bson:"start_date"`
	EndDate   primitive.DateTime `json:"end_date" bson:"end_date"`
	Status    string             `json:"status" bson:"status"`
}

func (i *Salary) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Salary
	return bson.Marshal((*t)(i))
}

func (i *Salary) Collection() string {
	return "salaries"
}

func (i *Salary) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "start_date", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "end_date", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "deleted_at", Value: 1}}, Options: options.Index()},
	}
}
