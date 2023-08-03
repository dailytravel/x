package model

import (
	"html"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Model         `bson:",inline"`
	Locale        *string              `json:"locale,omitempty" bson:"locale,omitempty"`
	Name          string               `json:"name" bson:"name"`
	Email         string               `json:"email" bson:"email"`
	Phone         *string              `json:"phone,omitempty" bson:"phone,omitempty"`
	Password      string               `json:"password" bson:"password"`
	Roles         []*string            `json:"roles,omitempty" bson:"roles,omitempty"`
	Mfa           *Mfa                 `json:"mfa,omitempty" bson:"mfa,omitempty"`
	Timezone      *string              `json:"timezone,omitempty" bson:"timezone,omitempty"`
	LoginAttempts *int                 `json:"login_attempts,omitempty" bson:"login_attempts,omitempty"`
	LastLogin     *int                 `json:"last_login,omitempty" bson:"last_login,omitempty"`
	LastIP        *string              `json:"last_ip,omitempty" bson:"last_ip,omitempty"`
	VerifiedAt    *primitive.Timestamp `json:"verified_at,omitempty" bson:"verified_at,omitempty"`
	LastActivity  *primitive.Timestamp `json:"last_activity,omitempty" bson:"last_activity,omitempty"`
	Identities    []*Identity          `json:"identities,omitempty" bson:"identities,omitempty"`
	Metadata      primitive.M          `json:"metadata,omitempty" bson:"metadata,omitempty"`
	Status        *string              `json:"status,omitempty" bson:"status,omitempty"`
}

func (User) IsEntity() {}

// remove spacing characters
func (i *User) Santize(s string) string {
	return html.EscapeString(strings.TrimSpace(s))
}

func (i *User) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t User
	return bson.Marshal((*t)(i))
}

func (i *User) Collection() string {
	return "users"
}

func (i *User) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "email", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "phone", Value: 1}}, Options: options.Index().SetUnique(true).SetSparse(true)},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "verified_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "deleted_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "name", Value: "text"}, {Key: "email", Value: "text"}}, Options: options.Index().SetWeights(bson.M{"name": 2, "email": 1})},
	}
}
