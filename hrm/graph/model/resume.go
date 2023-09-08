package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Resume struct {
	Model          `bson:",inline"`
	UID            primitive.ObjectID `json:"uid" bson:"uid"`
	Title          string             `json:"title" bson:"title"`
	Summary        string             `json:"summary" bson:"summary"`
	Experience     []Experience       `json:"experience" bson:"experience"`
	Education      []Education        `json:"education" bson:"education"`
	Skills         []Skill            `json:"skills" bson:"skills"`
	Certifications []Certification    `json:"certifications" bson:"certifications"`
	Languages      []Language         `json:"languages" bson:"languages"`
	References     []Reference        `json:"references" bson:"references"`
	Status         string             `json:"status" bson:"status"`
}

func (i *Resume) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Resume
	return bson.Marshal((*t)(i))
}

func (i *Resume) Collection() string {
	return "resumes"
}

func (i *Resume) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "deleted", Value: 1}}, Options: options.Index()},
	}
}
