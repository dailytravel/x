package migrations

import (
	"context"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"log"

	"github.com/dailytravel/x/account/auth"
	"github.com/dailytravel/x/account/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Key struct {
	Database *mongo.Database
	Model    *model.Key
}

// Create mongo collection
func (m *Key) Migrate() error {
	col := m.Database.Collection(m.Model.Collection())
	indexes, err := col.Indexes().List(context.Background())
	if err != nil {
		return err
	}

	indexNames := make(map[string]bool)
	for indexes.Next(context.Background()) {
		var index bson.M
		if err := indexes.Decode(&index); err != nil {
			return err
		}

		indexNames[index["name"].(string)] = true
	}

	for _, index := range m.Model.Index() {
		keys := index.Keys
		if keys != nil {
			indexName := ""
			for _, key := range keys.(bson.D) {
				indexName = key.Key
				break
			}

			if !indexNames[indexName] {
				if _, err := col.Indexes().CreateOne(context.Background(), index); err != nil {
					return err
				}
			}
		}
	}

	// check if certificate exists
	filter := bson.D{}
	if err := col.FindOne(context.Background(), filter).Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			// create current and next certificate
			for _, status := range []string{"current", "next"} {

				privateKey, err := auth.GenerateRSAKeyPair(2048)
				if err != nil {
					fmt.Println("Error generating RSA key pair:", err)
					return err
				}

				privateKeyPEM := pem.EncodeToMemory(&pem.Block{
					Type:  "RSA PRIVATE KEY",
					Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
				})

				publicKey := &privateKey.PublicKey

				fingerprint, err := auth.CalculateFingerprint(publicKey)
				if err != nil {
					fmt.Println("Error calculating fingerprint:", err)
					return err
				}
				fmt.Printf("Fingerprint: %x\n", fingerprint)

				thumbprintSHA256, err := auth.CalculateThumbprint(publicKey, sha256.New())
				if err != nil {
					fmt.Println("Error calculating thumbprint:", err)
					return err
				}

				// Convert byte slices to hexadecimal strings
				fingerprintHex := hex.EncodeToString(fingerprint)
				thumbprintSHA256Hex := hex.EncodeToString(thumbprintSHA256)

				cert := &model.Key{
					Name:        status,
					Provider:    "local",
					Certificate: string(privateKeyPEM),
					Fingerprint: fingerprintHex,
					Thumbprint:  thumbprintSHA256Hex,
					Type:        "RSA",
					Status:      status,
				}

				if _, err := col.InsertOne(context.Background(), cert); err != nil {
					log.Println("Error inserting certificate:", err)
					return err
				}
			}
		} else {
			return err
		}
	}

	// create JWKS file
	if err := auth.CreateJWKSFile(col); err != nil {
		fmt.Println("Error creating JWKS file:", err)
		log.Panic(err)
	}

	return nil
}
