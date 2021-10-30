package rest

import (
	"net/http"

	"github.com/mernat/sso-clean-arch/api/middleware"
	sso "github.com/mernat/sso-clean-arch/api/rest/sso"
	_ "github.com/mernat/sso-clean-arch/docs"
	ssoUseCase "github.com/mernat/sso-clean-arch/usecase/sso"

	"github.com/gorilla/mux"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title SSO Clean Arch Service
// @version 1
// @description SSO Clean Arch Service Documentation
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email mernat777@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host http://0.0.0.0:8181
// @BasePath /
// @schemes http
func ServeAPI(endpoint string) error {
	var ssoInterface ssoUseCase.Service
	ssoService := ssoUseCase.NewService(ssoInterface)
	ssoHandler := sso.NewSSOServiceHandler(ssoService)

	r := mux.NewRouter()
	ssoRouter := r.PathPrefix("/api/v1").Subrouter()
	ssoRouter.Methods("POST").Path("/sso/register").HandlerFunc(ssoHandler.RegistrationHandler)
	ssoRouter.Methods("POST").Path("/sso/login").HandlerFunc(ssoHandler.LoginHandler)
	//Not applicable on r.w.
	ssoRouter.Methods("GET").Path("/sso/verify").HandlerFunc(ssoHandler.VerifyToken)
	//Sharing JWKS with microservices endpoint
	ssoRouter.Methods("GET").Path("/sso/jwks").HandlerFunc(ssoHandler.GetJwks)

	ssoRouter.PathPrefix("/docs/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://0.0.0.0:8181/api/v1/docs/doc.json"),
	))

	ssoRouter.Use(middleware.JwtAuth)

	return http.ListenAndServe(endpoint, r)
}