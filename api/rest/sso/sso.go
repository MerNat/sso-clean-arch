package sso

import (
	"encoding/json"
	"net/http"

	serializer "github.com/mernat/sso-clean-arch/api/json"
	"github.com/mernat/sso-clean-arch/config"

	"github.com/mernat/sso-clean-arch/models"
	ssoUseCase "github.com/mernat/sso-clean-arch/usecase/sso"
)

type eventServiceHandler struct {
	service ssoUseCase.Service
}

func NewSSOServiceHandler(s ssoUseCase.Service) *eventServiceHandler {
	return &eventServiceHandler{
		service: s,
	}
}

// RegistrationHandler godoc
// @Summary Create a user
// @Description Creating a User
// @Tags SSO
// @Accept  json
// @Produce  json
// @Param user body models.User true "Users Data"
// @Success 200 {object} json.GenericResponse{success=bool,code=int,message=string} "ok"
// @Router /sso/register [post]
func (f *eventServiceHandler) RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)

	defer r.Body.Close()

	if err != nil {
		serializer.JSON(w, http.StatusBadRequest, &serializer.GenericResponse{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: "cant parse request",
		})
		return
	}

	err = f.service.RegistrationService(user)

	if err != nil {
		serializer.JSON(w, http.StatusBadRequest, &serializer.GenericResponse{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	serializer.JSON(w, http.StatusOK, &serializer.GenericResponse{
		Success: true,
		Code:    http.StatusOK,
		Message: "success",
	})
}

func (f *eventServiceHandler) VerifyToken(w http.ResponseWriter, r *http.Request) {
	userName := r.Context().Value(models.UserContextKey{}).(string)
	userFromToken := &models.User{Name: userName}
	serializer.JSON(w, http.StatusOK, &serializer.GenericResponse{
		Success: true,
		Code:    http.StatusOK,
		Data:    userFromToken,
	})

}

// LoginHandler godoc
// @Summary Login to service
// @Description Login
// @Tags SSO
// @Accept  json
// @Produce  json
// @Param user body models.User{email=string,password=string} true "Users Data"
// @Success 200 {object} json.GenericResponse{success=bool,code=int,message=string} "ok"
// @Router /sso/login [post]
func (f *eventServiceHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)

	defer r.Body.Close()

	if err != nil {
		serializer.JSON(w, http.StatusBadRequest, &serializer.GenericResponse{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: "cant parse request",
		})
		return
	}

	token, err := f.service.AuthService(user)

	if err != nil {
		serializer.JSON(w, http.StatusBadRequest, &serializer.GenericResponse{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	serializer.JSON(w, http.StatusOK, &serializer.GenericResponse{
		Success: true,
		Data:    token,
		Code:    http.StatusOK,
	})
}

func (f *eventServiceHandler) GetJwks(w http.ResponseWriter, r *http.Request) {
	serializer.JSON(w, http.StatusOK, &serializer.GenericResponse{
		Success: true,
		Data:    config.Config.JWKS.Keys,
		Code:    http.StatusOK,
	})
}
