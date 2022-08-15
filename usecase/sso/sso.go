package sso

import (
	"context"
	"errors"
	"time"

	"github.com/mernat/sso-clean-arch/config"
	"github.com/mernat/sso-clean-arch/models"
	"github.com/mernat/sso-clean-arch/utils"
	"github.com/square/go-jose/v3"
	"github.com/square/go-jose/v3/jwt"
	"golang.org/x/crypto/bcrypt"
)

type serviceLayer struct {
	repo models.Repository
}

func NewService(repo models.Repository) Service {
	return &serviceLayer{
		repo: repo,
	}
}

func (s *serviceLayer) RegistrationService(ctx context.Context, user *models.User) (err error) {
	//Validate user name and email
	if !utils.IsEmailValid(user.Email){
		err = errors.New("email not valid")
		return
	}
	if !utils.IsUsernameValid(user.Name){
		err = errors.New("username not valid")
		return 
	}
	if !utils.IsPasswordValid(user.Password){
		err = errors.New("password not valid")
		return 
	}
	err = s.repo.CreateUser(ctx, user)
	return
}

func (s *serviceLayer) AuthService(ctx context.Context, user *models.User) (token string, err error) {
	password := user.Password
	if password == "" {
		err = errors.New("invalid credentials")
		return
	}

	err = s.repo.GetUser(ctx, user)

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
