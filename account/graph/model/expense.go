package model

type Expense struct {
	UID string `json:"uid" bson:"uid"`
}

func (Expense) IsEntity() {}
