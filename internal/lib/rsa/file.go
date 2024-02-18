package rsa

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func PublicKeyFromFile(path string) (*rsa.PublicKey, error) {
	block, err := readBlockFromFile(path)
	if err != nil {
		return nil, err
	}

	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	switch key := key.(type) {
	case *rsa.PublicKey:
		return key, nil
	default:
		return nil, err
	}
}

func PrivateKeyFromFile(path string) (*rsa.PrivateKey, error) {
	block, err := readBlockFromFile(path)
	if err != nil {
		return nil, err
	}

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func readBlockFromFile(path string) (*pem.Block, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(bytes)
	if block == nil {
		return nil, err
	}

	return block, nil
}
