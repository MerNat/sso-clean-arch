package sso

import (
	"context"

	"github.com/mernat/sso-clean-arch/models"
)

type Service interface {
	RegistrationService(ctx context.Context, user *models.User) (err error)
	AuthService(ctx context.Context, user *models.User) (token string, err error)
}
