package model

import (
	"time"

	"github.com/typesense/typesense-go/typesense/api"
	"github.com/typesense/typesense-go/typesense/api/pointer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Log struct {
	Model     `bson:",inline"`
	User      primitive.ObjectID  `bson:"user" json:"user"`
	Message   string              `json:"message" bson:"message"`
	Status    string              `json:"status" bson:"status"`
	Method    string              `json:"method" bson:"method"`
	Latency   int64               `json:"latency" bson:"latency"`
	Path      string              `json:"path" bson:"path"`
	URL       string              `json:"url" bson:"url"`
	UserAgent string              `json:"user_agent" bson:"user_agent"`
	ClientIP  string              `json:"client_ip" bson:"client_ip"`
	Timestamp primitive.Timestamp `json:"timestamp" bson:"timestamp"`
}

func (i *Log) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Log
	return bson.Marshal((*t)(i))
}

func (i *Log) Collection() string {
	return "logs"
}

func (i *Log) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "method", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "timestamp", Value: 1}}, Options: options.Index()},
	}
}

func (i *Log) Schema() interface{} {
	return &api.CollectionSchema{
		Name: i.Collection(),
		Fields: []api.Field{
			{Name: "message", Type: "string"},
			{Name: "method", Type: "string", Facet: pointer.True()},
			{Name: "client_ip", Type: "string", Facet: pointer.True()},
			{Name: "path", Type: "string", Facet: pointer.True()},
			{Name: "status", Type: "int32", Facet: pointer.True()},
			{Name: "timestamp", Type: "int64", Facet: pointer.True()},
		},
		DefaultSortingField: pointer.String("timestamp"),
	}
}
