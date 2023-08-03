package graph

import (
	"github.com/dailytravel/x/reporting/graph/model"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	model model.Model
	db    *mongo.Database
	redis *redis.Client
}

func NewResolver(db *mongo.Database, rdb *redis.Client) *Resolver {
	return &Resolver{
		db:    db,
		redis: rdb,
		model: model.Model{},
	}
}
