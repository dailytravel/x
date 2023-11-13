package model

type Task struct {
	UID           string   `json:"uid" bson:"uid"`
	Collaborators []string `json:"collaborators,omitempty" bson:"collaborators,omitempty"`
}

func (Task) IsEntity() {}
