package model

type Assignment struct {
	UID string `json:"uid" bson:"uid"`
}

func (Assignment) IsEntity() {}
