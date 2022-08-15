package sse

import "github.com/mernat/sso-clean-arch/models"

type Service interface {
	GetBroker() *models.Broker
}
