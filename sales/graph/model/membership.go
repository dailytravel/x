package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Membership struct {
	Model   `bson:",inline"`
	UID     primitive.ObjectID `json:"user" bson:"user"`
	Tier    primitive.ObjectID `json:"tier" bson:"tier"`
	Number  string             `json:"number" bson:"number"`
	Since   string             `json:"since" bson:"since"`
	Until   string             `json:"until" bson:"until"`
	Billing primitive.M        `json:"billing,omitempty" bson:"billing,omitempty"`
	Payment primitive.M        `json:"payment,omitempty" bson:"payment,omitempty"`
	Status  string             `json:"status" bson:"status"`
}

func (Membership) IsEntity() {}

func (i *Membership) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Membership
	return bson.Marshal((*t)(i))
}

func (i *Membership) Collection() string {
	return "memberships"
}

func (i *Membership) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "tier", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "number", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "deleted_at", Value: 1}}, Options: options.Index()},
	}
}
