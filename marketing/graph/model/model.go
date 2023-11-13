package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model struct {
	ID        primitive.ObjectID  `json:"id,omitempty" bson:"_id,omitempty"`
	Workspace primitive.ObjectID  `json:"workspace,omitempty" bson:"workspace,omitempty"`
	Metadata  primitive.M         `json:"metadata,omitempty" bson:"metadata,omitempty"`
	Created   primitive.Timestamp `json:"created,omitempty" bson:"created,omitempty"`
	Updated   primitive.Timestamp `json:"updated,omitempty" bson:"updated,omitempty"`
	Deleted   primitive.Timestamp `json:"deleted,omitempty" bson:"deleted,omitempty"`
	Timestamp primitive.Timestamp `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
}
