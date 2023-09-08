package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	ClientCollection     = "clients"
	InvitationCollection = "invitations"
	PermissionCollection = "permissions"
	RoleCollection       = "roles"
	KeyCollection        = "keys"
	UserCollection       = "users"
)

type Model struct {
	ID        primitive.ObjectID  `json:"id,omitempty" bson:"_id,omitempty"`
	Metadata  primitive.M         `json:"metadata,omitempty" bson:"metadata,omitempty"`
	Created   primitive.Timestamp `json:"created,omitempty" bson:"created,omitempty"`
	Updated   primitive.Timestamp `json:"updated,omitempty" bson:"updated,omitempty"`
	Deleted   primitive.Timestamp `json:"deleted,omitempty" bson:"deleted,omitempty"`
	Timestamp primitive.Timestamp `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
}
