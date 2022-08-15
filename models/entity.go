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

type Broker struct {

	/*
	Events are pushed to this channel by the main events-gathering routine
	*/
	Notifier chan []byte

	// New client connections
	NewClients chan chan []byte

	// Closed client connections
	ClosingClients chan chan []byte

	// Client connections registry
	Clients map[chan []byte]bool
}

type MessageBroker struct {
	Message *string `json:"message"`
}
