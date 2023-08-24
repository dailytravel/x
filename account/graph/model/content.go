package model

type Content struct {
	UID string `json:"uid" bson:"uid"`
}

func (Content) IsEntity() {}
