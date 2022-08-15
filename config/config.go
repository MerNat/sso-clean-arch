package config

import (
	"crypto/rsa"
	"encoding/json"
	"log"
	"os"

	"github.com/mernat/sso-clean-arch/utils"
	"github.com/square/go-jose/v3"
)

var (
	Config ServiceConfig
)

type ServiceConfig struct {
	RestfulEndpoint *string `json:"restfulapi_endpoint"`
	JWKS            *jose.JSONWebKeySet
	PrivateKey      *rsa.PrivateKey
	Name            *string
	Version         *string
	Environment     *string
	Development     *bool
}

func InitConfiguration(filename string) (*ServiceConfig, error) {
	Config = ServiceConfig{}

	file, err := os.Open(filename)
	if err != nil {
		panic("Configuration file not found.")
	}

	if err := json.NewDecoder(file).Decode(&Config); err != nil {
		panic("Couldn't decode config values to struct.")
	}

	Config.PrivateKey, err = utils.NewRSAKey()

	if err != nil {
		log.Fatalf("failed to generate rsa key: %v", err)
	}

	jwk := utils.NewJSONWebKey(&Config.PrivateKey.PublicKey)

	Config.JWKS = &jose.JSONWebKeySet{Keys: []jose.JSONWebKey{*jwk}}
	return &Config, nil
}
