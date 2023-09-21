package model

type Post struct {
	ID string `json:"id" bson:"_id"`
}

func (Post) IsEntity() {}
