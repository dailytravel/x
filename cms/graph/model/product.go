package model

type Product struct {
	ID string `json:"id" bson:"_id"`
}

func (Product) IsEntity() {}
