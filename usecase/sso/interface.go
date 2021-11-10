package sso

import "github.com/mernat/sso-clean-arch/models"

type Service interface {
	RegistrationService(user *models.User) (err error)
	AuthService(user *models.User) (token string, err error)
}
