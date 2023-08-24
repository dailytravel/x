package model

type Portfolio struct {
	UID string `json:"uid" bson:"uid"`
}

func (Portfolio) IsEntity() {}
