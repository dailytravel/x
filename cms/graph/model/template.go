package model

type Template struct {
	ID string `json:"id" bson:"_id"`
}

func (Template) IsEntity() {}
