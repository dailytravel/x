package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Invitation struct {
	Model     `bson:",inline"`
	UID       primitive.ObjectID  `json:"uid" bson:"uid"`
	Email     string              `json:"email" bson:"email"`
	Roles     []string            `json:"roles" bson:"roles"`
	Status    string              `json:"status" bson:"status"`
	ExpiresAt primitive.Timestamp `json:"expiresAt" bson:"expires_at"`
	Metadata  primitive.M         `json:"metadata,omitempty"`
}

func (i *Invitation) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

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
		{Keys: bson.D{{Key: "expires_at", Value: 1}}},
		{Keys: bson.D{{Key: "created_at", Value: 1}}},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}},
	}
}
