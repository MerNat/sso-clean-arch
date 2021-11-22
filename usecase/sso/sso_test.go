package sso

import (
	"context"
	"testing"

	"github.com/mernat/sso-clean-arch/models"
	"github.com/stretchr/testify/assert"
)

func Test_serviceLayer_RegistrationService(t *testing.T) {
	u := &models.User{
		Name:     "AnyName",
		Email:    "anyemail@gmail.com",
		Password: "mypassword",
	}

	ctx := context.Background()

	ssoRepo := models.NewInmemRepository()
	ssoUseCase := NewService(ssoRepo)

	err := ssoUseCase.RegistrationService(ctx, u)

	assert.Nil(t, err)
}
