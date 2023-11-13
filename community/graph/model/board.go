package model

type Board struct {
	ID string `json:"_id" bson:"id"`
}

func (Board) IsEntity() {}
