package auth

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"time"

	"github.com/dailytravel/x/account/graph/model"
	"github.com/golang-jwt/jwt/v4"
)

var (
	token *jwt.Token
	err   error
)

// CreateToken takes some claims and a private key (either rsa or ec) and returns a signed json web token
func CreateToken(claims map[string]interface{}, kid string, key interface{}) (string, error) {
	switch k := key.(type) {
	case *rsa.PrivateKey:
		{
			token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims(claims))
			token.Header["kid"] = kid
			return token.SignedString(k)
		}
	case *ecdsa.PrivateKey:
		{
			token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims(claims))
			token.Header["kid"] = kid
			return token.SignedString(k)
		}
	case []byte:
		{
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(claims))
			token.Header["kid"] = kid
			return token.SignedString(k)
		}
	}
	return "", errors.New("invalid private key")
}

func Payload(u *model.User, k *model.Key, clientID string, expiresIn int) (*model.Payload, error) {
	privBlock, _ := pem.Decode([]byte(k.Certificate))
	if privBlock == nil {
		return nil, fmt.Errorf("error decoding private key PEM block")
	}
	privKey, err := x509.ParsePKCS1PrivateKey(privBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("error parsing private key: %v", err)
	}

	//generate token
	access_token, err := CreateToken(jwt.MapClaims(map[string]interface{}{
		"sub":      u.ID.Hex(),
		"email":    u.Email,
		"locale":   u.Locale,
		"timezone": u.Timezone,
		"name":     u.Name,
		"roles":    u.Roles,
		"azp":      clientID,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Minute * time.Duration(8600)).Unix(),
	}), k.Kid, privKey)
	if err != nil {
		return nil, fmt.Errorf("token not found")
	}

	refresh_token, err := CreateToken(jwt.MapClaims(map[string]interface{}{
		"sub": u.ID.Hex(),
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24 * 90).Unix(),
	}), k.Kid, privKey)
	if err != nil {
		return nil, fmt.Errorf("token not found")
	}

	return &model.Payload{
		AccessToken:  access_token,
		RefreshToken: refresh_token,
		TokenType:    "Bearer",
		ExpiresIn:    expiresIn,
	}, nil
}
