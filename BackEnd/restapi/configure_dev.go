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
	"uva-devtest/restapi/operations/answer"
	"uva-devtest/restapi/operations/auth"
	"uva-devtest/restapi/operations/published_test"
	"uva-devtest/restapi/operations/question"
	"uva-devtest/restapi/operations/tag"
	"uva-devtest/restapi/operations/team"
	"uva-devtest/restapi/operations/test"
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

	api.ReAuthCookieAuth = handlers.ReAuth

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	// Cookies y auth

	//TODO

	api.AuthLoginHandler = auth.LoginHandlerFunc(handlers.Login) // POST /accesstokens

	api.AuthReloginHandler = auth.ReloginHandlerFunc(handlers.Relogin) // PUT /accesstokens/{username}

	api.AuthCloseSessionsHandler = auth.CloseSessionsHandlerFunc(handlers.CloseSessions) // DELETE /accesstokens/{username}

	api.AuthLogoutHandler = auth.LogoutHandlerFunc(handlers.Logout) // GET /logout

	// /users

	api.UserRegisterUserHandler = user.RegisterUserHandlerFunc(handlers.RegisterUser) // POST /users
	api.UserGetUsersHandler = user.GetUsersHandlerFunc(handlers.GetUsers)             // GET /users

	api.UserPostEmailUserHandler = user.PostEmailUserHandlerFunc(handlers.PostEmailUser) // POST /emailUsers

	api.UserGetUserHandler = user.GetUserHandlerFunc(handlers.GetUser)             // GET /users/{username}
	api.UserPutUserHandler = user.PutUserHandlerFunc(handlers.PutUser)             // PUT /users/{username}
	api.UserDeleteUserHandler = user.DeleteUserHandlerFunc(handlers.DeleteUser)    // DELETE /users/{username}
	api.UserPutPasswordHandler = user.PutPasswordHandlerFunc(handlers.PutPassword) // PUT /users/{username}/password
	api.UserPutRoleHandler = user.PutRoleHandlerFunc(handlers.PutRole)             // PUT /users/{username}/role

	api.UserRecoverPasswordHandler = user.RecoverPasswordHandlerFunc(handlers.RecoverPassword)       // PUT /users/{username}/recoverPassword
	api.UserPostRecoveryTokenHandler = user.PostRecoveryTokenHandlerFunc(handlers.PostRecoveryToken) // POST /users/{username}/passRecoveryTokens

	api.UserGetTeamsOfUserHandler = user.GetTeamsOfUserHandlerFunc(handlers.GetTeamsOfUser)    // GET /users/{username}/teams
	api.TeamPostTeamHandler = team.PostTeamHandlerFunc(handlers.PostTeam)                      // POST /users/{username}/teams
	api.UserGetTeamFromUserHandler = user.GetTeamFromUserHandlerFunc(handlers.GetTeamFromUser) // GET /users/{username}/teams/{teamname}

	api.UserGetFavoriteEditQuestionsHandler = user.GetFavoriteEditQuestionsHandlerFunc(handlers.GetFavoriteEditQuestions) // GET /users/{username}/favoriteEditQuestions

	api.UserGetFavoriteQuestionsHandler = user.GetFavoriteQuestionsHandlerFunc(handlers.GetFavoriteQuestions)       // GET /users/{username}/favoriteQuestions
	api.UserGetFavoriteQuestionHandler = user.GetFavoriteQuestionHandlerFunc(handlers.GetFavoriteQuestion)          // GET /users/{username}/favoriteQuestions/{questionid}
	api.UserAddQuestionFavoriteHandler = user.AddQuestionFavoriteHandlerFunc(handlers.AddFavoriteQuestion)          // PUT /users/{username}/favoriteQuestions/{questionid}
	api.UserRemoveQuestionFavoriteHandler = user.RemoveQuestionFavoriteHandlerFunc(handlers.RemoveFavoriteQuestion) // DELETE /users/{username}/favoriteQuestions/{questionid}

	api.UserGetAvailableQuestionsOfUserHandler = user.GetAvailableQuestionsOfUserHandlerFunc(handlers.GetAvailableQuestions)             // GET /users/{username}/availableQuestions
	api.UserGetAvailableEditQuestionsOfUserHandler = user.GetAvailableEditQuestionsOfUserHandlerFunc(handlers.GetAvailableEditQuestions) // GET /users/{username}/availableEditQuestions

	api.UserGetSharedQuestionsOfUserHandler = user.GetSharedQuestionsOfUserHandlerFunc(handlers.GetSharedQuestions)  // GET /users/{username}/sharedQuestions
	api.UserGetSharedQuestionFromUserHandler = user.GetSharedQuestionFromUserHandlerFunc(handlers.GetSharedQuestion) // GET /users/{username}/sharedQuestions/{testid}

	api.UserGetEditQuestionsOfUserHandler = user.GetEditQuestionsOfUserHandlerFunc(handlers.GetEditQuestionsOfUser)                   // GET /users/{username}/editQuestions
	api.UserGetPublicEditQuestionsOfUserHandler = user.GetPublicEditQuestionsOfUserHandlerFunc(handlers.GetPublicEditQuestionsOfUser) // GET /users/{username}/publicEditQuestions

	api.UserGetQuestionsOfUserHandler = user.GetQuestionsOfUserHandlerFunc(handlers.GetQuestionsOfUser)  // GET /users/{username}/questions
	api.UserPostQuestionHandler = user.PostQuestionHandlerFunc(handlers.PostQuestionOfUser)              // POST /users/{username}/questions
	api.UserGetQuestionFromUserHandler = user.GetQuestionFromUserHandlerFunc(handlers.GetQuestionOfUser) // GET /users/{username}/questions/{questionid}

	api.UserCopyQuestionHandler = user.CopyQuestionHandlerFunc(handlers.CopyQuestion) // POST /users/{username}/questions/{questionid}/copiedQuestions

	api.UserGetFavoriteEditTestsHandler = user.GetFavoriteEditTestsHandlerFunc(handlers.GetFavoriteEditTests) // GET /users/{username}/favoriteEditTests

	api.UserGetFavoriteTestsHandler = user.GetFavoriteTestsHandlerFunc(handlers.GetFavoriteTests)       // GET /users/{username}/favoriteTests
	api.UserGetFavoriteTestHandler = user.GetFavoriteTestHandlerFunc(handlers.GetFavoriteTest)          // GET /users/{username}/favoriteTests/{testid}
	api.UserAddTestFavoriteHandler = user.AddTestFavoriteHandlerFunc(handlers.AddFavoriteTest)          // PUT /users/{username}/favoriteTests/{testid}
	api.UserRemoveTestFavoriteHandler = user.RemoveTestFavoriteHandlerFunc(handlers.RemoveFavoriteTest) // DELETE /users/{username}/favoriteTests/{testid}

	api.UserGetSharedEditTestsFromUserHandler = user.GetSharedEditTestsFromUserHandlerFunc(handlers.GetSharedEditTests) // GET /users/{username}/sharedEditTests

	api.UserGetSharedPublishedTestsFromUserHandler = user.GetSharedPublishedTestsFromUserHandlerFunc(handlers.GetSharedPublishedTests) // GET /users/{username}/sharedPublishedTests

	api.UserGetSharedTestsFromUserHandler = user.GetSharedTestsFromUserHandlerFunc(handlers.GetSharedTests) // GET /users/{username}/sharedTests
	api.UserGetSharedTestFromUserHandler = user.GetSharedTestFromUserHandlerFunc(handlers.GetSharedTest)    // GET /users/{username}/sharedTests/{testid}

	api.UserGetPublicEditTestsFromUserHandler = user.GetPublicEditTestsFromUserHandlerFunc(handlers.GetPublicETestsFromUser) // GET /users/{username}/publicEditTests

	api.UserGetEditTestsFromUserHandler = user.GetEditTestsFromUserHandlerFunc(handlers.GetEditTestsFromUser)

	api.UserGetTestsFromUserHandler = user.GetTestsFromUserHandlerFunc(handlers.GetTestsFromUser) // GET /users/{username}/tests
	api.UserPostTestHandler = user.PostTestHandlerFunc(handlers.PostTest)                         // POST /users/{username}/tests
	api.UserGetTestFromUserHandler = user.GetTestFromUserHandlerFunc(handlers.GetTestFromUser)    // GET /users/{username}/tests/{testid}

	api.UserCopyTestHandler = user.CopyTestHandlerFunc(handlers.CopyTest) // POST /users/{username}/tests/{testid}/copiedTests

	api.UserGetInvitedTestsByTeamsAndUserHandler = user.GetInvitedTestsByTeamsAndUserHandlerFunc(handlers.GetInvitedTestsByTeamsAndUser) // GET /users/{username}/invitedTestsByTeamsAndUser

	api.UserGetInvitedTestsFromUserHandler = user.GetInvitedTestsFromUserHandlerFunc(handlers.GetInvitedTests) // GET /users/{username}/invitedTests
	api.UserGetInvitedTestFromUserHandler = user.GetInvitedTestFromUserHandlerFunc(handlers.GetInvitedTest)    // GET /users/{username}/invitedTests/{testid}

	api.UserGetPublishedTestsFromUserHandler = user.GetPublishedTestsFromUserHandlerFunc(handlers.GetPublishedTestsFromUser)                   // GET /users/{username}/publishedTests
	api.UserGetPublicPublishedTestsFromUserHandler = user.GetPublicPublishedTestsFromUserHandlerFunc(handlers.GetPublicPublishedTestsFromUser) // GET /users/{username}/publicPublishedTests

	api.UserGetSolvableTestsFromUserHandler = user.GetSolvableTestsFromUserHandlerFunc(handlers.GetSolvableTestsFromUser) // GET /users/{username}/solvableTests
	api.UserGetSolvableTestFromUserHandler = user.GetSolvableTestFromUserHandlerFunc(handlers.GetSolvableTestFromUser)    // GET /users/{username}/solvableTests/{testid}
	api.UserStartAnswerHandler = user.StartAnswerHandlerFunc(handlers.StartAnswer)                                        // POST /users/{username}/solvableTests/{testid}/answers

	api.UserGetOpenAnswersFromUserTestHandler = user.GetOpenAnswersFromUserTestHandlerFunc(handlers.GetOpenAnswersTestUser) // GET /users/{username}/solvableTests/{testid}/openAnswers

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

	api.TeamGetInvitedTestsFromTeamHandler = team.GetInvitedTestsFromTeamHandlerFunc(handlers.GetInvitedTestsFromTeam) // GET /teams/{teamname}/invitedTests
	api.TeamGetInvitedTestFromTeamHandler = team.GetInvitedTestFromTeamHandlerFunc(handlers.GetInvitedTestFromTeam)    // GET /teams/{teamname}/invitedTests/{testid}

	// /questions

	api.QuestionGetAllEditQuestionsHandler = question.GetAllEditQuestionsHandlerFunc(handlers.GetAllEditQuestions) // GET /allEditQuestions
	api.QuestionGetAllQuestionsHandler = question.GetAllQuestionsHandlerFunc(handlers.GetAllQuestions)             // GET /allQuestions
	api.QuestionGetEditQuestionsHandler = question.GetEditQuestionsHandlerFunc(handlers.GetEditQuestions)          // GET /editQuestions

	api.QuestionGetQuestionsHandler = question.GetQuestionsHandlerFunc(handlers.GetQuestions)       // GET /questions
	api.QuestionGetQuestionHandler = question.GetQuestionHandlerFunc(handlers.GetQuestion)          // GET /questions/{questionid}
	api.QuestionPutQuestionHandler = question.PutQuestionHandlerFunc(handlers.PutQuestion)          // PUT /questions/{questionid}
	api.QuestionDeleteQuestionHandler = question.DeleteQuestionHandlerFunc(handlers.DeleteQuestion) // DELETE /questions/{questionid}

	api.QuestionGetOptionsFromQuestionHandler = question.GetOptionsFromQuestionHandlerFunc(handlers.GetOptions) // GET /questions/{questionid}/options
	api.QuestionPostOptionHandler = question.PostOptionHandlerFunc(handlers.PostOption)                         // POST /questions/{questionid}/options

	api.QuestionGetOptionFromQuestionHandler = question.GetOptionFromQuestionHandlerFunc(handlers.GetOption) // GET /questions/{questionid}/options/{optionindice}
	api.QuestionPutOptionHandler = question.PutOptionHandlerFunc(handlers.PutOption)                         // PUT /questions/{questionid}/options/{optionindice}
	api.QuestionDeleteOptionHandler = question.DeleteOptionHandlerFunc(handlers.DeleteOption)                // DELETE /questions/{questionid}/options/{optionindice}

	api.QuestionGetTagsFromQuestionHandler = question.GetTagsFromQuestionHandlerFunc(handlers.GetQuestionTags)       // GET /questions/{questionid}/tags
	api.QuestionGetTagFromQuestionHandler = question.GetTagFromQuestionHandlerFunc(handlers.GetQuestionTag)          // GET /questions/{questionid}/tags/{tag}
	api.QuestionAddTagToQuestionHandler = question.AddTagToQuestionHandlerFunc(handlers.AddQuestionTag)              // PUT /questions/{questionid}/tags/{tag}
	api.QuestionRemoveTagFromQuestionHandler = question.RemoveTagFromQuestionHandlerFunc(handlers.RemoveQuestionTag) // DELETE /questions/{questionid}/tags/{tag}

	api.QuestionGetTeamsFromQuestionHandler = question.GetTeamsFromQuestionHandlerFunc(handlers.GetTeamsFromQuestion) // GET /questions/{questionid}/teams
	api.QuestionAddTeamToQuestionHandler = question.AddTeamToQuestionHandlerFunc(handlers.AddQuestionTeam)            // PUT /questions/{questionid}/teams/{teamname}
	api.QuestionRemoveTeamToQuestionHandler = question.RemoveTeamToQuestionHandlerFunc(handlers.RemoveQuestionTeam)   // DELETE /questions/{questionid}/teams/{teamname}

	// /tests

	api.TestGetPublicEditTestsHandler = test.GetPublicEditTestsHandlerFunc(handlers.GetPublicEditTests) // GET /publicEditTests
	api.TestGetPublicTestsHandler = test.GetPublicTestsHandlerFunc(handlers.GetPublicTests)             // GET /publicTests

	api.TestGetAllEditTestsHandler = test.GetAllEditTestsHandlerFunc(handlers.GetAllEditTests) // GET /editTests

	api.TestGetAllTestsHandler = test.GetAllTestsHandlerFunc(handlers.GetAllTests) // GET /tests
	api.TestGetTestHandler = test.GetTestHandlerFunc(handlers.GetTest)             // GET /tests/{testid}
	api.TestPutTestHandler = test.PutTestHandlerFunc(handlers.PutTest)             // PUT /tests/{testid}
	api.TestDeleteTestHandler = test.DeleteTestHandlerFunc(handlers.DeleteTest)    // DELETE /tests/{testid}

	api.TestGetTagsFromTestHandler = test.GetTagsFromTestHandlerFunc(handlers.GetTagsFromTest)       // GET /tests/{testid}/tags
	api.TestGetTagFromTestHandler = test.GetTagFromTestHandlerFunc(handlers.GetTagFromTest)          // GET /tests/{testid}/tags/{tag}
	api.TestAddTagToTestHandler = test.AddTagToTestHandlerFunc(handlers.AddTagToTest)                // PUT /tests/{testid}/tags/{tag}
	api.TestRemoveTagFromTestHandler = test.RemoveTagFromTestHandlerFunc(handlers.RemoveTagFromTest) // DELETE /tests/{testid}/tags/{tag}

	api.TestGetAdminTeamsFromTestHandler = test.GetAdminTeamsFromTestHandlerFunc(handlers.GetAdminTeamsFromTest) // GET /tests/{testid}/teams
	api.TestAddAdminTeamToTestHandler = test.AddAdminTeamToTestHandlerFunc(handlers.AddAdminTeamToTest)          // PUT /tests/{testid}/teams/{teamname}
	api.TestRemoveAdminTeamToTestHandler = test.RemoveAdminTeamToTestHandlerFunc(handlers.RemoveAdminTeamTest)   // DELETE /tests/{testid}/teams/{teamname}

	api.TestPostPublishedTestHandler = test.PostPublishedTestHandlerFunc(handlers.PublishTest)                       // POST /tests/{testid}/publishedTests
	api.TestGetPublishedTestsFromTestHandler = test.GetPublishedTestsFromTestHandlerFunc(handlers.GetPTestsFromTest) // GET /tests/{testid}/publishedTests

	api.TestGetQuestionsFromTestHandler = test.GetQuestionsFromTestHandlerFunc(handlers.GetQuestionsFromTest)   // GET /tests/{testid}/questions
	api.TestGetQuestionFromTestHandler = test.GetQuestionFromTestHandlerFunc(handlers.GetQuestionFromTest)      // GET /tests/{testid}/questions/{questionid}
	api.TestAddQuestionToTestHandler = test.AddQuestionToTestHandlerFunc(handlers.AddQuestionToTest)            // PUT /tests/{testid}/questions/{questionid}
	api.TestRemoveQuestionFromTestHandler = test.RemoveQuestionFromTestHandlerFunc(handlers.RemoveQuestionTest) // DELETE /tests/{testid}/questions/{questionid}

	// /publishedTests

	api.PublishedTestGetPublicPublishedTestsHandler = published_test.GetPublicPublishedTestsHandlerFunc(handlers.GetPublicPTests) // GET /publishedTests
	api.PublishedTestGetPublicPublishedTestHandler = published_test.GetPublicPublishedTestHandlerFunc(handlers.GetPublicPTest)    // GET /publishedTests/{testid}

	api.PublishedTestGetPublishedTestsHandler = published_test.GetPublishedTestsHandlerFunc(handlers.GetPTests) // GET /publishedTests
	api.PublishedTestGetPublishedTestHandler = published_test.GetPublishedTestHandlerFunc(handlers.GetPTest)    // GET /publishedTests/{testid}

	api.PublishedTestGetUsersFromPublishedTestHandler = published_test.GetUsersFromPublishedTestHandlerFunc(handlers.GetUsersFromPTest) // GET /publishedTests/{testid}/users

	api.PublishedTestInviteUserToPublishedTestHandler = published_test.InviteUserToPublishedTestHandlerFunc(handlers.InviteUserPTest) // PUT /publishedTests/{testid}/users/{username}
	api.PublishedTestRemoveUserToPublishedTestHandler = published_test.RemoveUserToPublishedTestHandlerFunc(handlers.RemoveUserPTest) // DELETE /publishedTests/{testid}/users/{username}

	api.PublishedTestGetTeamsFromPublishedTestHandler = published_test.GetTeamsFromPublishedTestHandlerFunc(handlers.GetTeamsFromPTest) // GET /publishedTests/{testid}/teams
	api.PublishedTestInviteTeamToPublishedTestHandler = published_test.InviteTeamToPublishedTestHandlerFunc(handlers.InviteTeamPTest)   // PUT /publishedTests/{testid}/teams/{teamname}
	api.PublishedTestRemoveTeamToPublishedTestHandler = published_test.RemoveTeamToPublishedTestHandlerFunc(handlers.RemoveTeamPTest)   // DELETE /publishedTests/{testid}/teams/{teamname}

	api.PublishedTestGetTagsFromPublishedTestHandler = published_test.GetTagsFromPublishedTestHandlerFunc(handlers.GetTagsFromPTest) // GET /publishedTests/{testid}/tags
	api.PublishedTestGetTagFromPublishedTestHandler = published_test.GetTagFromPublishedTestHandlerFunc(handlers.GetTagFromPTest)    // GET /publishedTests/{testid}/tags/{tag}

	api.PublishedTestGetQuestionsFromPublishedTestsHandler = published_test.GetQuestionsFromPublishedTestsHandlerFunc(handlers.GetQuestionsPTest) // GET /publishedTests/{testid}/questions
	api.PublishedTestGetQuestionFromPublishedTestsHandler = published_test.GetQuestionFromPublishedTestsHandlerFunc(handlers.GetQuestionPTest)    // GET /publishedTests/{testid}/questions/{questionid}

	api.PublishedTestGetOptionsFromPublishedQuestionHandler = published_test.GetOptionsFromPublishedQuestionHandlerFunc(handlers.GetOptionsPQuestion) // GET /publishedTests/{testid}/questions/{questionid}/options

	api.PublishedTestGetTagsFromPublishedQuestionHandler = published_test.GetTagsFromPublishedQuestionHandlerFunc(handlers.GetTagsPQuestion) // GET /publishedTests/{testid}/questions/{questionid}/tags

	api.PublishedTestGetAnswersFromPublishedTestsHandler = published_test.GetAnswersFromPublishedTestsHandlerFunc(handlers.GetAnswersPTest) // GET /publishedTests/{testid}/answers

	api.PublishedTestGetQuestionAnswersFromPublishedTestQuestionHandler =
		published_test.GetQuestionAnswersFromPublishedTestQuestionHandlerFunc(handlers.GetQuestionAnswersPTest) // GET /publishedTests/{testid}/questions/{questionid}/qanswers

	// /answers

	api.AnswerGetAnswersHandler = answer.GetAnswersHandlerFunc(handlers.GetAnswers)       // GET /answers
	api.AnswerGetAnswerHandler = answer.GetAnswerHandlerFunc(handlers.GetAnswer)          // GET /answers/{answerid}
	api.AnswerFinishAnswerHandler = answer.FinishAnswerHandlerFunc(handlers.FinishAnswer) // PUT /answers/{answerid}

	api.AnswerGetQuestionAnswersFromAnswerHandler = answer.GetQuestionAnswersFromAnswerHandlerFunc(handlers.GetQuestionAnswers)       // GET /answers/{answerid}/qanswers
	api.AnswerPostQuestionAnswerHandler = answer.PostQuestionAnswerHandlerFunc(handlers.PostQuestionAnswer)                           // POST /answers/{answerid}/qanswers
	api.AnswerGetQuestionAnswerFromAnswerHandler = answer.GetQuestionAnswerFromAnswerHandlerFunc(handlers.GetQuestionAnswer)          // GET /answers/{answerid}/qanswers/{questionid}
	api.AnswerPutQuestionAnswerFromAnswerHandler = answer.PutQuestionAnswerFromAnswerHandlerFunc(handlers.PutQuestionAnswer)          // PUT /answers/{answerid}/qanswers/{questionid}
	api.AnswerDeleteQuestionAnswerFromAnswerHandler = answer.DeleteQuestionAnswerFromAnswerHandlerFunc(handlers.DeleteQuestionAnswer) // DELETE /answers/{answerid}/qanswers/{questionid}

	api.AnswerPutReviewHandler = answer.PutReviewHandlerFunc(handlers.PutReview) // PUT /answers/{answerid}/qanswers/{questionid}/review

	api.AnswerGetQuestionsFromAnswerHandler = answer.GetQuestionsFromAnswerHandlerFunc(handlers.GetQuestionsFromAnswer) // GET /answers/{answerid}/questions

	api.AnswerGetQuestionAnswersFromAnswerAndQuestionHandler = answer.GetQuestionAnswersFromAnswerAndQuestionHandlerFunc(handlers.GetQAnswerFromAnswerAndQuestion) // GET /answers/{answerid}/questions/{questionid}/qanswers

	// /tags

	api.TagGetTagsHandler = tag.GetTagsHandlerFunc(handlers.GetTags)                                                 // GET /tags
	api.TagGetTagHandler = tag.GetTagHandlerFunc(handlers.GetTag)                                                    // GET /tags/{tag}
	api.TagGetQuestionsFromTagHandler = tag.GetQuestionsFromTagHandlerFunc(handlers.GetQuestionsFromTag)             // GET /tags/{tag}/questions
	api.TagGetEditQuestionsFromTagHandler = tag.GetEditQuestionsFromTagHandlerFunc(handlers.GetEditQuestionsFromTag) // GET /tags/{tag}/editQuestions
	api.TagGetTestsFromTagHandler = tag.GetTestsFromTagHandlerFunc(handlers.GetTestsFromTag)                         // GET /tags/{tag}/tests
	api.TagGetEditTestsFromTagHandler = tag.GetEditTestsFromTagHandlerFunc(handlers.GetEditTestsFromTag)             // GET /tags/{tag}/editTests

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
			"cache-control", "Origin", "Cache-Control", "Pragma", "Expires", "If-Modified-Since"},
		MaxAge: 300,
		Debug:  true,
	})
	handler = c.Handler(handler)
	return handler
}
