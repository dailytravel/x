package model

type Task struct {
	UID      string   `json:"uid" bson:"uid"`
	Assignee string   `json:"assignee,omitempty" bson:"assignee,omitempty"`
	Members  []string `json:"members,omitempty" bson:"members,omitempty"`
}

func (Task) IsEntity() {}
