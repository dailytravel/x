package auth

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"

	"github.com/dailytravel/x/account/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		if item.Type == "RSA" && item.Certificate != "" {
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
				Kid: item.ID.Hex(),
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

func CreateJWKSFile(col *mongo.Collection) error {
	var items []*model.Key
	status := []string{"current", "previous", "next"}
	filter := bson.M{"expires_at": bson.M{"$exists": false}, "status": bson.M{"$in": status}}
	opts := options.Find()
	opts.SetSort(bson.M{"updated_at": 1})

	cursor, err := col.Find(context.Background(), filter, opts)
	if err != nil {
		return err
	}
	defer cursor.Close(context.Background())

	if err = cursor.All(context.Background(), &items); err != nil {
		return err
	}

	jwks, err := getJWKS(items)
	if err != nil {
		return err
	}

	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	// Create the desired file path relative to the current working directory
	filePath := filepath.Join(currentDir, "gateway", ".well-known", "jwks.json")

	file, err := json.MarshalIndent(jwks, "", " ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(filePath, file, 0644); err != nil {
		return err
	}

	return nil
}
