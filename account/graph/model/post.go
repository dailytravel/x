package model

type Post struct {
	UID string `json:"uid" bson:"uid"`
}

func (Post) IsEntity() {}
