package model

type File struct {
	UID string `json:"uid" bson:"uid"`
}

func (File) IsEntity() {}
