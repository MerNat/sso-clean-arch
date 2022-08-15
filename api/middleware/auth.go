package middleware

import (
	"context"
	"crypto/rsa"
	"errors"
	"net/http"
	"regexp"
	"strings"

	serializer "github.com/mernat/sso-clean-arch/api/json"
	"github.com/mernat/sso-clean-arch/config"
	"github.com/mernat/sso-clean-arch/models"
	"github.com/square/go-jose/v3/jwt"
)

var JwtAuth = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		currentURL := r.URL.Path

		if strings.Contains(currentURL, "/sso/stream") {
			next.ServeHTTP(w, r)
		}

		isSwagger, _ := regexp.MatchString("/docs*", currentURL)
		if isSwagger {
			next.ServeHTTP(w, r)
			return
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type")
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "OPTIONS" {
			return
		}

		rejectURL := []string{
			"/sso/register",
			"/sso/login",
			"/sso/jwks",
			"/sso/broadcast",
		}

		for _, value := range rejectURL {
			if strings.Contains(currentURL, value) {
				next.ServeHTTP(w, r)
				return
			}
		}

		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			serializer.JSON(w, http.StatusBadRequest, &serializer.GenericResponse{
				Success: false,
				Code:    http.StatusBadRequest,
				Message: "missing token",
			})
			return
		}

		splitted := strings.Split(tokenHeader, " ")

		if len(splitted) != 2 {
			serializer.JSON(w, http.StatusBadRequest, &serializer.GenericResponse{
				Success: false,
				Code:    http.StatusBadRequest,
				Message: "Invalid/Malformed Token",
			})
			return
		}

		//TODO We can also add a jwtParameter check here. (If it'll allow for a specific jwtParam.)

		theToken := &splitted[1]

		user, err := GetUserFromToken(*theToken)

		if err != nil {
			serializer.JSON(w, http.StatusBadRequest, &serializer.GenericResponse{
				Success: false,
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}

		ctx := context.WithValue(r.Context(), models.UserContextKey{}, user.Name)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

//GetUserFromToken gets customer_id from token
func GetUserFromToken(tk string) (user *models.UserClaim, err error) {

	parsedJWK := config.Config.JWKS.Keys[0]
	parsedRSAKey, ok := parsedJWK.Key.(*rsa.PublicKey)
	if !ok {
		err = errors.New("type cast failed")
		return
	}

	user, err = ParseJWT(tk, parsedRSAKey)
	if err != nil {
		return
	}
	return
}

func ParseJWT(signedJWT string, pubKey *rsa.PublicKey) (*models.UserClaim, error) {
	token, err := jwt.ParseSigned(signedJWT)
	if err != nil {
		return nil, errors.New("invalid jwt")
	}

	claims := new(models.UserClaim)
	if err := token.Claims(pubKey, claims); err != nil {
		return nil, errors.New("invalid jwt")
	}

	err = claims.Valid()
	if err != nil {
		if err == jwt.ErrExpired {
			return nil, errors.New("invalid jwt")
		}

		return nil, errors.New("invalid jwt")
	}

	return claims, nil
}
