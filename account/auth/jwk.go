package auth

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"math/big"

	"github.com/dailytravel/x/account/graph/model"
)

// JWKSKey represents a JSON Web Key (JWK).
type JWKSKey struct {
	Kid     string   `json:"kid"`
	Kty     string   `json:"kty"`
	Use     string   `json:"use"`
	Alg     string   `json:"alg"`
	N       string   `json:"n"`
	E       string   `json:"e"`
	X5c     []string `json:"x5c,omitempty"`
	X5t     string   `json:"x5t,omitempty"`
	X5tS256 string   `json:"x5t#S256,omitempty"`
}

func getJWKS(items []*model.Key) (map[string]interface{}, error) {
	jwks := make(map[string]interface{})
	keys := make([]JWKSKey, 0)

	for _, item := range items {
		if item.Type == "rsa" && item.Certificate != "" {
			privBlock, _ := pem.Decode([]byte(item.Certificate))
			if privBlock == nil {
				return nil, fmt.Errorf("error decoding private key PEM block")
			}
			privKey, err := x509.ParsePKCS1PrivateKey(privBlock.Bytes)
			if err != nil {
				return nil, fmt.Errorf("error parsing private key: %v", err)
			}

			pubKey := privKey.Public()

			jwksKey := JWKSKey{
				Kid: item.Kid,
				Kty: "RSA",
				Use: "sig",   // Use "sig" for signatures, "enc" for encryption
				Alg: "RS256", // The algorithm used with this key (e.g., RS256, RS384, RS512)
				N:   base64.RawURLEncoding.EncodeToString(pubKey.(*rsa.PublicKey).N.Bytes()),
				E:   base64.RawURLEncoding.EncodeToString(big.NewInt(int64(pubKey.(*rsa.PublicKey).E)).Bytes()),
				X5c: []string{base64.StdEncoding.EncodeToString(privBlock.Bytes)},
			}

			keys = append(keys, jwksKey)
		}
	}

	// Add the keys to the JWKS map
	jwks["keys"] = keys

	return jwks, nil
}
