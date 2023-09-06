package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"hash"
)

func GenerateRSAKeyPair(bits int) (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func CalculateFingerprint(pubKey *rsa.PublicKey) ([]byte, error) {
	// Serialize the public key to DER format
	pubKeyBytes, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		return nil, err
	}

	// Calculate the SHA-1 hash of the DER-encoded public key
	sha1Hash := sha1.New()
	_, _ = sha1Hash.Write(pubKeyBytes)
	fingerprint := sha1Hash.Sum(nil)

	return fingerprint, nil
}

func CalculateThumbprint(pubKey *rsa.PublicKey, hashFunc hash.Hash) ([]byte, error) {
	// Serialize the public key to DER format
	pubKeyBytes, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		return nil, err
	}

	// Calculate the hash of the DER-encoded public key using the specified hash function
	_, _ = hashFunc.Write(pubKeyBytes)
	thumbprint := hashFunc.Sum(nil)

	return thumbprint, nil
}
