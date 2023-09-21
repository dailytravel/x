package model

type Package struct {
	ID string `json:"id" bson:"_id"`
}

func (Package) IsEntity() {}
