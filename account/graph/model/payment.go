package model

type Payment struct {
	UID string `json:"uid" bson:"uid"`
}

func (Payment) IsEntity() {}
