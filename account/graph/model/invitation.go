package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Invitation struct {
	Model     `bson:",inline"`
	Sender    primitive.ObjectID `json:"sender" bson:"sender"`
	Recipient string             `json:"recipient" bson:"recipient"`
	Roles     []string           `json:"roles" bson:"roles"`
	Status    string             `json:"status"`
	Metadata  primitive.M        `json:"metadata,omitempty"`
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
		{Keys: bson.D{{Key: "sender", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "recipient", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
	}
}
