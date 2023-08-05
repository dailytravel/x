package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Membership struct {
	Model   `bson:",inline"`
	User    primitive.ObjectID `json:"user,omitempty" bson:"user,omitempty"`
	Tier    primitive.ObjectID `json:"tier,omitempty" bson:"tier,omitempty"`
	Number  string             `json:"number,omitempty" bson:"number,omitempty"`
	Since   primitive.DateTime `json:"since,omitempty" bson:"since,omitempty"`
	Until   primitive.DateTime `json:"until,omitempty" bson:"until,omitempty"`
	Billing primitive.M        `json:"billing,omitempty" bson:"billing,omitempty"`
	Payment primitive.M        `json:"payment,omitempty" bson:"payment,omitempty"`
	Status  string             `json:"status,omitempty" bson:"status,omitempty"`
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
