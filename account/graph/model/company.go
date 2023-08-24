package model

type Company struct {
	UID string `json:"uid" bson:"uid"`
}

func (Company) IsEntity() {}
