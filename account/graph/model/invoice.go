package model

type Invoice struct {
	UID string `json:"uid" bson:"uid"`
}

func (Invoice) IsEntity() {}
