package model

type Attendance struct {
	UID string `json:"uid" bson:"uid"`
}

func (Attendance) IsEntity() {}
