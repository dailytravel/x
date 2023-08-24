package model

type Comment struct {
	UID string `json:"uid" bson:"uid"`
}

func (Comment) IsEntity() {}
