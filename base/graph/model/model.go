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
	CreatedAt primitive.Timestamp `json:"createdAt,omitempty" bson:"created_at,omitempty"`
	UpdatedAt primitive.Timestamp `json:"updatedAt,omitempty" bson:"updated_at,omitempty"`
	DeletedAt primitive.Timestamp `json:"deletedAt,omitempty" bson:"deleted_at,omitempty"`
	CreatedBy *primitive.ObjectID `json:"createdBy,omitempty" bson:"created_by,omitempty"`
	UpdatedBy *primitive.ObjectID `json:"updatedBy,omitempty" bson:"updated_by,omitempty"`
	DeletedBy *primitive.ObjectID `json:"deletedBy,omitempty" bson:"deleted_by,omitempty"`
}
