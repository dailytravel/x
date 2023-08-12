package model

type User struct {
	Model `bson:",inline"`
	Name  string `json:"name"`
}

func (User) IsEntity() {}
