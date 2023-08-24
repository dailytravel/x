package model

type Coupon struct {
	UID string `json:"uid" bson:"uid"`
}

func (Coupon) IsEntity() {}
