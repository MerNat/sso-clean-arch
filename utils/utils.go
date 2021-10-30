package utils

import (
	"crypto/rand"
	"crypto/rsa"

	"github.com/square/go-jose/v3"
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(plaintext string) (cryptedtext string) {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(plaintext), bcrypt.DefaultCost)
	cryptedtext = string(hashPassword)
	return
}

func NewRSAKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}

func NewJSONWebKey(pubKey *rsa.PublicKey) *jose.JSONWebKey {
	return &jose.JSONWebKey{
		Key:       pubKey,
		Algorithm: "RS256",
		Use:       "sig",
	}
}