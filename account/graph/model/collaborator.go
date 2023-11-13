package model

type Collaborator struct {
	UID string `json:"uid" bson:"uid"`
}

func (Collaborator) IsEntity() {}
