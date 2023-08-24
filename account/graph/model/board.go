package model

type Board struct {
	UID string `json:"uid" bson:"uid"`
}

func (Board) IsEntity() {}
