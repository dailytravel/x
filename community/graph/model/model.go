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
	CreatedAt primitive.Timestamp `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt primitive.Timestamp `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt primitive.Timestamp `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
	CreatedBy *primitive.ObjectID `json:"created_by,omitempty" bson:"created_by,omitempty"`
	UpdatedBy *primitive.ObjectID `json:"updated_by,omitempty" bson:"updated_by,omitempty"`
	DeletedBy *primitive.ObjectID `json:"deleted_by,omitempty" bson:"deleted_by,omitempty"`
}
