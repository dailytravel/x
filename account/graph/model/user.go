package model

import (
	"html"
	"strings"
	"time"

	"github.com/typesense/typesense-go/typesense/api"
	"github.com/typesense/typesense-go/typesense/api/pointer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Model         `bson:",inline"`
	Name          string               `json:"name" bson:"name"`
	Email         string               `json:"email" bson:"email"`
	Phone         *string              `json:"phone,omitempty" bson:"phone,omitempty"`
	Password      string               `json:"password" bson:"password"`
	Roles         []string             `json:"roles,omitempty" bson:"roles,omitempty"`
	Mfa           *Mfa                 `json:"mfa,omitempty" bson:"mfa,omitempty"`
	Locale        *string              `json:"locale,omitempty" bson:"locale,omitempty"`
	Timezone      *string              `json:"timezone,omitempty" bson:"timezone,omitempty"`
	Picture       *string              `json:"picture,omitempty" bson:"picture,omitempty"`
	LastLogin     *primitive.Timestamp `json:"lastLogin,omitempty" bson:"last_login,omitempty"`
	EmailVerified *bool                `json:"emailVerified,omitempty" bson:"email_verified,omitempty"`
	PhoneVerified *bool                `json:"phoneVerified,omitempty" bson:"phone_verified,omitempty"`
	Status        *string              `json:"status" bson:"status"`
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
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "created_at", Value: 1}}},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}},
		{Keys: bson.D{{Key: "verified_at", Value: 1}}},
		{Keys: bson.D{{Key: "deleted_at", Value: 1}}},
		{Keys: bson.D{{Key: "name", Value: "text"}, {Key: "email", Value: "text"}}, Options: options.Index().SetWeights(bson.M{"name": 2, "email": 1})},
	}
}

func (i *User) Schema() interface{} {
	return &api.CollectionSchema{
		Name: i.Collection(),
		Fields: []api.Field{
			{Name: "name", Type: "string", Optional: pointer.True()},
			{Name: "picture", Type: "string", Optional: pointer.True()},
			{Name: "email", Type: "string"},
			{Name: "phone", Type: "string", Optional: pointer.True()},
			{Name: "status", Type: "string", Facet: pointer.True()},
			{Name: "created_at", Type: "int32"},
			{Name: "updated_at", Type: "int32"},
			{Name: "roles", Type: "string[]"},
		},
		DefaultSortingField: pointer.String("created_at"),
		EnableNestedFields:  pointer.True(),
	}
}

func (i *User) Document() map[string]interface{} {
	document := map[string]interface{}{
		"id":         i.ID.Hex(),
		"name":       i.Name,
		"email":      i.Email,
		"phone":      i.Phone,
		"picture":    i.Picture,
		"status":     i.Status,
		"roles":      i.Roles,
		"created_at": i.CreatedAt.T,
		"updated_at": i.UpdatedAt.T,
	}

	return document
}
