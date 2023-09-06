package auth

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/api/idtoken"
)

func GoogleLogin(c context.Context, idToken string) (string, error) {
	payload, err := idtoken.Validate(c, idToken, os.Getenv("GOOGLE_CLIENT_ID"))
	if err != nil {
		return "", err
	}
	if payload.Audience != os.Getenv("GOOGLE_CLIENT_ID") {
		return "", fmt.Errorf("error validating audience")
	}
	if payload.Issuer != "accounts.google.com" && payload.Issuer != "https://accounts.google.com" {
		return "", fmt.Errorf("error validating issuer")
	}
	return payload.Claims["email"].(string), nil
}
