package model

type Link struct {
	UID string `json:"uid" bson:"uid"`
}

func (Link) IsEntity() {}
