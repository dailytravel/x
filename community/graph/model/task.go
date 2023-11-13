package model

type Task struct {
	ID string `json:"_id" bson:"id"`
}

func (Task) IsEntity() {}
