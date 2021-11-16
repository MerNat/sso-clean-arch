package sso

import (
	"testing"

	"github.com/mernat/sso-clean-arch/models"
	"github.com/stretchr/testify/assert"
)

func Test_serviceLayer_RegistrationService(t *testing.T) {
	u := &models.User{
		Name: "AnyName",
		Email: "anyemail@gmail.com",
		Password: "mypassword",
	}

	ssoRepo := models.NewInmemRepository()
	ssoUseCase := NewService(ssoRepo)

	err := ssoUseCase.RegistrationService(u)

	assert.Nil(t, err)
}
