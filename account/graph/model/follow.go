package model

type Follow struct {
	UID string `json:"uid" bson:"uid"`
}

func (Follow) IsEntity() {}
