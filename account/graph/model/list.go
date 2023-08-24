package model

type List struct {
	UID string `json:"uid" bson:"uid"`
}

func (List) IsEntity() {}
