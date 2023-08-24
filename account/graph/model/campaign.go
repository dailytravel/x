package model

type Campaign struct {
	UID string `json:"uid" bson:"uid"`
}

func (Campaign) IsEntity() {}
