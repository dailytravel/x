package model

type Reaction struct {
	UID string `json:"uid" bson:"uid"`
}

func (Reaction) IsEntity() {}
