package model

type Place struct {
	ID string `json:"id" bson:"_id"`
}

func (Place) IsEntity() {}
