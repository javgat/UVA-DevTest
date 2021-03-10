// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/rs/cors"

	"uva-devtest/handlers"
	"uva-devtest/restapi/operations"
	"uva-devtest/restapi/operations/auth"
	"uva-devtest/restapi/operations/team"
	"uva-devtest/restapi/operations/user"
)

//go:generate swagger generate server --target ../../BackEnd --name Dev --spec ../swagger.yml --principal interface{}

func configureFlags(api *operations.DevAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.DevAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "Bearer" header is set
	api.BearerHeaderAuth = handlers.BearerAuth

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	api.AuthLoginHandler = auth.LoginHandlerFunc(handlers.Login) // POST /accesstokens

	api.UserPutPasswordHandler = user.PutPasswordHandlerFunc(handlers.PutPassword) // PUT /passwords/{username}

	api.UserRegisterUserHandler = user.RegisterUserHandlerFunc(handlers.RegisterUser) // POST /users
	api.UserGetUsersHandler = user.GetUsersHandlerFunc(handlers.GetUsers)             // GET /users

	api.UserGetUserHandler = user.GetUserHandlerFunc(handlers.GetUser)                            // GET /users/{username}
	api.UserPutUserHandler = user.PutUserHandlerFunc(handlers.PutUser)                            // PUT /users/{username}
	api.UserDeleteUserHandler = user.DeleteUserHandlerFunc(handlers.DeleteUser)                   // DELETE /users/{username}
	api.UserGetTeamsOfUserHandler = user.GetTeamsOfUserHandlerFunc(handlers.GetTeamsOfUser)       // GET /users/{username}/teams
	api.UserAddTeamOfUserHandler = user.AddTeamOfUserHandlerFunc(handlers.AddTeamOfUser)          // PUT /users/{username}/teams/{teamname}
	api.UserDeleteTeamOfUserHandler = user.DeleteTeamOfUserHandlerFunc(handlers.DeleteTeamOfUser) // DELETE /users/{username}/teams/{teamname}
	api.TeamGetUserTeamRoleHandler = team.GetUserTeamRoleHandlerFunc(handlers.GetUserTeamRole)    // GET /users/{username}/teams/{teamname}/role
	api.TeamPutUserTeamRoleHandler = team.PutUserTeamRoleHandlerFunc(handlers.PutUserTeamRole)    // PUT /users/{username}/teams/{teamname}/role

	api.TeamGetTeamsHandler = nil           // GET /teams
	api.TeamPostTeamHandler = nil           // POST /teams
	api.TeamGetTeamHandler = nil            // GET /teams/{teamname}
	api.TeamPutTeamHandler = nil            // PUT /teams/{teamname}
	api.TeamDeleteTeamHandler = nil         // DELETE /teams/{teamname}
	api.UserGetUsersFromTeamHandler = nil   // GET /teams/{teamname}/users
	api.UserAddUserFromTeamHandler = nil    // PUT /teams/{teamname}/users/{username}
	api.UserDeleteUserFromTeamHandler = nil // DELETE /teams/{teamname}/users/{username}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	var c = cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"POST", "PUT", "GET", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Accept-Encoding", "Authorization", "Content-Type",
			"cache-control", "Origin", "X-CSRF-Token"}, //por que el token?
		MaxAge: 300,
		Debug:  true,
	})
	handler = c.Handler(handler)
	return handler
}
