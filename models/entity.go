package models

import "github.com/golang-jwt/jwt"

//UserContextKey declares context
type UserContextKey struct{}

type UserClaim struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

type User struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
