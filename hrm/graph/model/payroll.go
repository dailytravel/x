package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Payroll struct {
	Model    `bson:",inline"`
	UID      primitive.ObjectID `bson:"uid" json:"uid"`
	Date     primitive.DateTime `json:"date" bson:"date"`
	Amount   float64            `json:"amount" bson:"amount"`
	Currency string             `json:"currency" bson:"currency"`
	Metadata primitive.M        `json:"metadata,omitempty" bson:"metadata,omitempty"`
	Status   string             `json:"status" bson:"status"`
}

func (i *Payroll) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.Created.IsZero() {
		i.Created = now
	}

	i.Updated = now

	type t Payroll
	return bson.Marshal((*t)(i))
}

func (i *Payroll) Collection() string {
	return "payrolls"
}

func (i *Payroll) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}}},
		{Keys: bson.D{{Key: "date", Value: 1}}},
		{Keys: bson.D{{Key: "created", Value: 1}}},
		{Keys: bson.D{{Key: "updated", Value: 1}}},
		{Keys: bson.D{{Key: "deleted", Value: 1}}},
	}
}
