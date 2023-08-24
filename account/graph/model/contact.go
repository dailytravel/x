package model

type Contact struct {
	UID string `json:"uid" bson:"uid"`
}

func (Contact) IsEntity() {}
