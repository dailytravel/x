package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Employee struct {
	Model        `bson:",inline"`
	UID          primitive.ObjectID `json:"uid" bson:"uid"`
	Organization primitive.ObjectID `json:"organization" bson:"organization"`
	Position     primitive.ObjectID `json:"position" bson:"position"`
	Reference    string             `json:"reference" bson:"reference"`
	FirstName    string             `json:"first_name" bson:"first_name"`
	LastName     string             `json:"last_name" bson:"last_name"`
	Email        string             `json:"email" bson:"email"`
	Phone        string             `json:"phone" bson:"phone"`
	Address      string             `json:"address" bson:"address"`
	Birthday     primitive.DateTime `json:"birthday" bson:"birthday"`
	HireDate     primitive.DateTime `json:"hire_date" bson:"hire_date"`
	Status       string             `json:"status" bson:"status"`
}

func (i *Employee) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Employee
	return bson.Marshal((*t)(i))
}

func (i *Employee) Collection() string {
	return "employees"
}

func (i *Employee) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "organization", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "position", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "deleted_at", Value: 1}}, Options: options.Index()},
	}
}
