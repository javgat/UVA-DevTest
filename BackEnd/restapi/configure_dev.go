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
	"uva-devtest/restapi/operations/question"
	"uva-devtest/restapi/operations/team"
	"uva-devtest/restapi/operations/user"
)

//go:generate swagger generate server --target ../../BackEnd --name Dev --spec ../swagger.yml --principal models.User

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

	// Applies when the "Cookie" header is set
	api.BearerCookieAuth = handlers.BearerAuth

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	// Cookies y auth

	//TODO

	api.AuthLoginHandler = auth.LoginHandlerFunc(handlers.Login) // POST /accesstokens

	// /users

	api.UserRegisterUserHandler = user.RegisterUserHandlerFunc(handlers.RegisterUser) // POST /users
	api.UserGetUsersHandler = user.GetUsersHandlerFunc(handlers.GetUsers)             // GET /users

	api.UserGetUserHandler = user.GetUserHandlerFunc(handlers.GetUser)             // GET /users/{username}
	api.UserPutUserHandler = user.PutUserHandlerFunc(handlers.PutUser)             // PUT /users/{username}
	api.UserDeleteUserHandler = user.DeleteUserHandlerFunc(handlers.DeleteUser)    // DELETE /users/{username}
	api.UserPutPasswordHandler = user.PutPasswordHandlerFunc(handlers.PutPassword) // PUT /users/{username}/password
	api.UserPutRoleHandler = user.PutRoleHandlerFunc(handlers.PutRole)             // PUT /users/{username}/role

	api.UserGetTeamsOfUserHandler = user.GetTeamsOfUserHandlerFunc(handlers.GetTeamsOfUser)    // GET /users/{username}/teams
	api.TeamPostTeamHandler = team.PostTeamHandlerFunc(handlers.PostTeam)                      // POST /users/{username}/teams
	api.UserGetTeamFromUserHandler = user.GetTeamFromUserHandlerFunc(handlers.GetTeamFromUser) // GET /users/{username}/teams/{teamname}

	api.UserGetQuestionsOfUserHandler = user.GetQuestionsOfUserHandlerFunc(handlers.GetQuestionsOfUser)  // GET /users/{username}/questions
	api.UserPostQuestionHandler = user.PostQuestionHandlerFunc(handlers.PostQuestionOfUser)              // POST /users/{username}/questions
	api.UserGetQuestionFromUserHandler = user.GetQuestionFromUserHandlerFunc(handlers.GetQuestionOfUser) // GET /users/{username}/questions/{questionid}

	api.UserGetTestsFromUserHandler = user.GetTestsFromUserHandlerFunc(handlers.GetTestsFromUser) // GET /users/{username}/tests
	api.UserPostTestHandler = user.PostTestHandlerFunc(handlers.PostTest)                         // POST /users/{username}/tests
	api.UserGetTestFromUserHandler = user.GetTestFromUserHandlerFunc(handlers.GetTestFromUser)    // GET /users/{username}/tests/{testid}

	api.UserGetPublishedTestsFromUserHandler = user.GetPublishedTestsFromUserHandlerFunc(handlers.GetPTestsFromUser) // GET /users/{username}/publishedTests
	api.UserGetPublishedTestFromUserHandler = user.GetPublishedTestFromUserHandlerFunc(handlers.GetPTestFromUser)    // GET /users/{username}/publishedTests/{testid}
	api.UserStartAnswerHandler = user.StartAnswerHandlerFunc(handlers.StartAnswer)                                   // POST /users/{username}/publishedTests/{testid}/answers

	api.UserGetAnsweredTestsFromUserHandler = user.GetAnsweredTestsFromUserHandlerFunc(handlers.GetATestsFromUser) // GET /users/{username}/answeredTests
	api.UserGetAnsweredTestFromUserHandler = user.GetAnsweredTestFromUserHandlerFunc(handlers.GetATestFromUser)    // GET /users/{username}/answeredTests/{testid}
	api.UserGetAnswersFromUserAnsweredTestHandler =
		user.GetAnswersFromUserAnsweredTestHandlerFunc(handlers.GetAnswersFromUserATest) // GET /users/{username}/answeredTests/{testid}/answers

	api.UserGetAnswersFromUserHandler = user.GetAnswersFromUserHandlerFunc(handlers.GetAnswersFromUser) // GET /users/{username}/answers
	api.UserGetAnswerFromUserHandler = user.GetAnswerFromUserHandlerFunc(handlers.GetAnswerFromUser)    // GET /users/{username}/answers/{answerid}

	// /teams

	api.TeamGetTeamsHandler = team.GetTeamsHandlerFunc(handlers.GetTeams) // GET /teams

	api.TeamGetTeamHandler = team.GetTeamHandlerFunc(handlers.GetTeam)          // GET /teams/{teamname}
	api.TeamPutTeamHandler = team.PutTeamHandlerFunc(handlers.PutTeam)          // PUT /teams/{teamname}
	api.TeamDeleteTeamHandler = team.DeleteTeamHandlerFunc(handlers.DeleteTeam) // DELETE /teams/{teamname}

	api.TeamGetAdminsHandler = team.GetAdminsHandlerFunc(handlers.GetAdminsFromTeam) // GET /teams/{teamname}/admins
	api.TeamGetAdminHandler = team.GetAdminHandlerFunc(handlers.GetAdminFromTeam)    // GET /teams/{teamname}/admins/{username}
	api.TeamAddAdminHandler = team.AddAdminHandlerFunc(handlers.AddAdminToTeam)      // PUT /teams/{teamname}/admins/{username}

	api.TeamGetMembersHandler = team.GetMembersHandlerFunc(handlers.GetMembersFromTeam) // GET /teams/{teamname}/members
	api.TeamGetMemberHandler = team.GetMemberHandlerFunc(handlers.GetMemberFromTeam)    // GET /teams/{teamname}/members/{username}
	api.TeamAddMemberHandler = team.AddMemberHandlerFunc(handlers.AddMemberToTeam)      // PUT /teams/{teamname}/members/{username}

	api.TeamGetUsersFromTeamHandler = team.GetUsersFromTeamHandlerFunc(handlers.GetUsersFromTeam)       // GET /teams/{teamname}/users
	api.TeamGetUserFromTeamHandler = team.GetUserFromTeamHandlerFunc(handlers.GetUserFromTeam)          // GET /teams/{teamname}/users/{username}
	api.TeamDeleteUserFromTeamHandler = team.DeleteUserFromTeamHandlerFunc(handlers.DeleteUserFromTeam) // DELETE /teams/{teamname}/users/{username}

	api.TeamGetQuestionsFromTeamHandler = team.GetQuestionsFromTeamHandlerFunc(handlers.GetQuestionsFromTeam) // GET /teams/{teamname}/questions
	api.TeamGetQuestionFromTeamHandler = team.GetQuestionFromTeamHandlerFunc(handlers.GetQuestionFromTeam)    // GET /teams/{teamname}/questions/{questionid}

	api.TeamGetTestsFromTeamHandler = team.GetTestsFromTeamHandlerFunc(handlers.GetTestsFromTeam) // GET /teams/{teamname}/tests
	api.TeamGetTestFromTeamHandler = team.GetTestFromTeamHandlerFunc(handlers.GetTestFromTeam)    // GET /teams/{teamname}/tests/{testid}

	api.TeamGetPublishedTestsFromTeamHandler = team.GetPublishedTestsFromTeamHandlerFunc(handlers.GetPTestsFromTeam) // GET /teams/{teamname}/publishedTests
	api.TeamGetPublishedTestFromTeamHandler = team.GetPublishedTestFromTeamHandlerFunc(handlers.GetPTestFromTeam)    // GET /teams/{teamname}/publishedTests/{testid}

	// /questions

	api.QuestionGetQuestionsHandler = question.GetQuestionsHandlerFunc(handlers.GetQuestions)       // GET /questions
	api.QuestionGetQuestionHandler = question.GetQuestionHandlerFunc(handlers.GetQuestion)          // GET /questions/{questionid}
	api.QuestionPutQuestionHandler = question.PutQuestionHandlerFunc(handlers.PutQuestion)          // PUT /questions/{questionid}
	api.QuestionDeleteQuestionHandler = question.DeleteQuestionHandlerFunc(handlers.DeleteQuestion) // DELETE /questions/{questionid}

	api.QuestionGetTagsFromQuestionHandler = question.GetTagsFromQuestionHandlerFunc(handlers.GetQuestionTags)       // GET /questions/{questionid}/tags
	api.QuestionGetTagFromQuestionHandler = question.GetTagFromQuestionHandlerFunc(handlers.GetQuestionTag)          // GET /questions/{questionid}/tags/{tag}
	api.QuestionAddTagToQuestionHandler = question.AddTagToQuestionHandlerFunc(handlers.AddQuestionTag)              // PUT /questions/{questionid}/tags/{tag}
	api.QuestionRemoveTagFromQuestionHandler = question.RemoveTagFromQuestionHandlerFunc(handlers.RemoveQuestionTag) // DELETE /questions/{questionid}/tags/{tag}

	api.QuestionGetTeamsFromQuestionHandler = question.GetTeamsFromQuestionHandlerFunc(handlers.GetTeamsFromQuestion) // GET /questions/{questionid}/teams
	api.QuestionAddTeamToQuestionHandler = question.AddTeamToQuestionHandlerFunc(handlers.AddQuestionTeam)            // PUT /questions/{questionid}/teams/{teamname}
	api.QuestionRemoveTeamToQuestionHandler = question.RemoveTeamToQuestionHandlerFunc(handlers.RemoveQuestionTeam)   // DELETE /questions/{questionid}/teams/{teamname}

	// TODO tests...

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
		//AllowedOrigins:   []string{"*"},
		AllowedOrigins:   []string{"https://localhost:*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"POST", "PUT", "GET", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Accept-Encoding", "Authorization", "Content-Type",
			"cache-control", "Origin"},
		MaxAge: 300,
		Debug:  true,
	})
	handler = c.Handler(handler)
	return handler
}
