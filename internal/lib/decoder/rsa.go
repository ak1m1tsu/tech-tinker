package decoder

import (
	"crypto/rsa"
	"encoding/base64"
	"github.com/golang-jwt/jwt/v4"
)

// DecodeRSAPublicKey decodes a base64 encoded RSA public key string into a *rsa.PublicKey struct.
func DecodeRSAPublicKey(data string) (*rsa.PublicKey, error) {
	decodedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	return jwt.ParseRSAPublicKeyFromPEM(decodedData)
}

// DecodeRSAPrivateKey decodes a base64 encoded RSA private key string into a *rsa.PrivateKey struct.
func DecodeRSAPrivateKey(data string) (*rsa.PrivateKey, error) {
	decodedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	return jwt.ParseRSAPrivateKeyFromPEM(decodedData)
}
