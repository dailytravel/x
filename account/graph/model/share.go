package model

type Share struct {
	UID string `json:"uid" bson:"uid"`
}

func (Share) IsEntity() {}
