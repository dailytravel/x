package model

type Task struct {
	UID string `json:"uid" bson:"uid"`
}

func (Task) IsEntity() {}
