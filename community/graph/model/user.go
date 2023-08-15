package model

type User struct {
	ID string `json:"id" bson:"_id"`
}

func (User) IsEntity() {}
