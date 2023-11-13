package model

type Post struct {
	ID string `bson:"_id" json:"id"`
}

func (Post) IsEntity() {}
