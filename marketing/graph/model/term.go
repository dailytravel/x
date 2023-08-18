package model

type Term struct {
	ID string `json:"id" bson:"_id"`
}

func (Term) IsEntity() {}
