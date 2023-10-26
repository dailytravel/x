package utils

import (
	"github.com/dailytravel/x/account/graph/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Credential(user model.User, credentialType string, secret string, expires *primitive.Timestamp) *model.Credential {
	return &model.Credential{
		UID:     user.ID,
		Type:    credentialType,
		Secret:  secret,
		Expires: *expires,
		Status:  "ACTIVE",
	}
}
