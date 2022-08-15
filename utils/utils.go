package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"regexp"

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

func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func IsUsernameValid(e string) bool {
	userRegix := regexp.MustCompile(`^[A-Za-z]{4,}$`)
	return userRegix.MatchString(e)
}

func IsPasswordValid(e string) bool {
	passRegix := regexp.MustCompile(`^[A-Za-z0-9!?#]{8,}$`)
	return passRegix.MatchString(e)
}
