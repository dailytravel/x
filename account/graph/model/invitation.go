package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Invitation struct {
	Model    `bson:",inline"`
	UID      primitive.ObjectID  `json:"uid" bson:"uid"`
	Email    string              `json:"email" bson:"email"`
	Roles    []*string           `json:"roles,omitempty" bson:"roles,omitempty"`
	Status   string              `json:"status" bson:"status"`
	Expires  primitive.Timestamp `json:"expires" bson:"expires"`
	Metadata primitive.M         `json:"metadata,omitempty"`
}

func (i *Invitation) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Invitation
	return bson.Marshal((*t)(i))
}

func (i *Invitation) Collection() string {
	return "invitations"
}

func (i *Invitation) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}}},
		{Keys: bson.D{{Key: "email", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "expires", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
	}
}
