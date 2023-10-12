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
	GivenName     *string              `json:"given_name,omitempty" bson:"given_name,omitempty"`
	FamilyName    *string              `json:"family_name,omitempty" bson:"family_name,omitempty"`
	Email         string               `json:"email" bson:"email"`
	Phone         *string              `json:"phone,omitempty" bson:"phone,omitempty"`
	Roles         []*string            `json:"roles,omitempty" bson:"roles,omitempty"`
	Mfa           *Mfa                 `json:"mfa,omitempty" bson:"mfa,omitempty"`
	Locale        *string              `json:"locale,omitempty" bson:"locale,omitempty"`
	Timezone      *string              `json:"timezone,omitempty" bson:"timezone,omitempty"`
	Picture       *string              `json:"picture,omitempty" bson:"picture,omitempty"`
	LastLogin     *primitive.Timestamp `json:"last_login,omitempty" bson:"last_login,omitempty"`
	EmailVerified *bool                `json:"email_verified,omitempty" bson:"email_verified,omitempty"`
	PhoneVerified *bool                `json:"phone_verified,omitempty" bson:"phone_verified,omitempty"`
	Status        *string              `json:"status" bson:"status"`
}

func (User) IsEntity() {}

// remove spacing characters
func (i *User) Santize(s string) string {
	return html.EscapeString(strings.TrimSpace(s))
}

func (i *User) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

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
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
		{Keys: bson.D{{Key: "verified", Value: 1}}},
		{Keys: bson.D{{Key: "deleted", Value: 1}}},
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
			{Name: "created", Type: "int32"},
			{Name: "updated", Type: "int32"},
			{Name: "roles", Type: "string[]"},
		},
		DefaultSortingField: pointer.String("created"),
		EnableNestedFields:  pointer.True(),
	}
}

func (i *User) Document() map[string]interface{} {
	document := map[string]interface{}{
		"id":      i.ID.Hex(),
		"name":    i.Name,
		"email":   i.Email,
		"phone":   i.Phone,
		"picture": i.Picture,
		"status":  i.Status,
		"roles":   i.Roles,
		"created": i.Created.T,
		"updated": i.Updated.T,
	}

	return document
}
