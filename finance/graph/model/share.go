package model

type Share struct {
	ID string `json:"id" bson:"_id"`
}

func (Share) IsEntity() {}
