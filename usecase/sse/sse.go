package sse

import (
	"github.com/mernat/sso-clean-arch/models"
)

type serviceLayer struct {
	repo models.BrokerRepo
}

func NewBrokerService(repo models.BrokerRepo) Service {
	return &serviceLayer{
		repo: repo,
	}
}

func (s *serviceLayer) GetBroker() *models.Broker {
	return s.repo.GetBroker()
}
