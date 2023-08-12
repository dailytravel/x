package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Payroll struct {
	Model    `bson:",inline"`
	Employee primitive.ObjectID `bson:"employee" json:"employee"`
	PayDate  primitive.DateTime `json:"pay_date" bson:"pay_date"`
	Amount   float64            `json:"amount" bson:"amount"`
	Currency string             `json:"currency" bson:"currency"`
	Metadata primitive.M        `json:"metadata,omitempty" bson:"metadata,omitempty"`
	Status   string             `json:"status" bson:"status"`
}

func (i *Payroll) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Payroll
	return bson.Marshal((*t)(i))
}

func (i *Payroll) Collection() string {
	return "payrolls"
}

func (i *Payroll) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "employee", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "pay_date", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "deleted_at", Value: 1}}, Options: options.Index()},
	}
}
