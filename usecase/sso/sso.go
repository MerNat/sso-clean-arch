package sso

import (
	"errors"
	"time"

	"github.com/mernat/sso-clean-arch/config"
	"github.com/mernat/sso-clean-arch/models"
	"github.com/square/go-jose/v3"
	"github.com/square/go-jose/v3/jwt"
	"golang.org/x/crypto/bcrypt"
)

type serviceLayer struct {
	service Service
}

func NewService(service Service) Service {
	return &serviceLayer{
		service: service,
	}
}

func (s *serviceLayer) RegistrationService(user models.User) (err error) {

	err = user.CreateUser()

	return
}

func (s *serviceLayer) AuthService(user models.User) (token string, err error) {
	password := user.Password
	if password == "" {
		err = errors.New("invalid credentials")
		return
	}

	err = user.GetUser()

	if err != nil {
		err = errors.New("email not found")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		err = errors.New("invalid credentials")
		return
	}

	//Generate the token

	tk := &models.UserClaim{Name: user.Name, Email: user.Email}
	tk.ExpiresAt = time.Now().AddDate(0, 0, 30).Unix() //Expire after 30 days
	tk.Issuer = "SSO-Issuer"

	signKey := jose.SigningKey{
		Algorithm: jose.RS256,
		Key:       config.Config.PrivateKey,
	}

	signer, err := jose.NewSigner(signKey, nil)
	if err != nil {
		return
	}

	token, err = jwt.Signed(signer).
		Claims(tk).
		CompactSerialize()
	return
}
