package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Organization struct {
	Model          `bson:",inline"`
	Owner          primitive.ObjectID  `json:"owner,omitempty" bson:"owner,omitempty"`
	Parent         primitive.ObjectID  `json:"parent,omitempty" bson:"parent,omitempty"`
	Type           string              `json:"type,omitempty" bson:"type,omitempty"`
	Name           string              `json:"name,omitempty" bson:"name,omitempty"`
	Description    string              `json:"description,omitempty" bson:"description,omitempty"`
	Industry       string              `json:"industry,omitempty" bson:"industry,omitempty"`
	Employees      int                 `json:"employees,omitempty" bson:"employees,omitempty"`
	Revenue        float64             `json:"revenue,omitempty" bson:"revenue,omitempty"`
	City           string              `json:"city,omitempty" bson:"city,omitempty"`
	Zip            string              `json:"zip,omitempty" bson:"zip,omitempty"`
	State          string              `json:"state,omitempty" bson:"state,omitempty"`
	Country        string              `json:"country,omitempty" bson:"country,omitempty"`
	Timezone       string              `json:"timezone,omitempty" bson:"timezone,omitempty"`
	Phone          string              `json:"phone,omitempty" bson:"phone,omitempty"`
	Website        string              `json:"website,omitempty" bson:"website,omitempty"`
	IsOrganization bool                `json:"is_organization,omitempty" bson:"is_organization,omitempty"`
	Status         string              `json:"status,omitempty" bson:"status,omitempty"`
	LastActivity   primitive.Timestamp `json:"last_activity,omitempty" bson:"last_activity,omitempty"`
}

func (Organization) IsEntity() {}

func (i *Organization) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Organization
	return bson.Marshal((*t)(i))
}

func (i *Organization) Collection() string {
	return "organizations"
}

func (i *Organization) Sanitize(s string) string {
	return s
}

func (i *Organization) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "owner", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "parent", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "type", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "deleted_at", Value: 1}}, Options: options.Index()},
	}
}
