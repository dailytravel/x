package model

type Task struct {
	UID      string   `json:"uid" bson:"uid"`
	Assignee string   `json:"assignee,omitempty" bson:"assignee,omitempty"`
	Shares   []string `json:"shares,omitempty" bson:"shares,omitempty"`
}

func (Task) IsEntity() {}
