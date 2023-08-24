package model

type Goal struct {
	UID string `json:"uid" bson:"uid"`
}

func (Goal) IsEntity() {}
