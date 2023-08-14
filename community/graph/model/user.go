package model

type User struct {
	Model `bson:",inline"`
}

func (User) IsEntity() {}
