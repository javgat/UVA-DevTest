// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package handlers provides functions that handle http requests
package handlers

import (
	"database/sql"
	"log"
	"uva-devtest/emailHelper"
	"uva-devtest/models"
	"uva-devtest/permissions"
	"uva-devtest/persistence/dao"
	"uva-devtest/persistence/dbconnection"
	"uva-devtest/restapi/operations/published_test"

	"github.com/go-openapi/runtime/middleware"
)

// GetPTests GET /publicPublishedTests. Returns all public published tests.
// Auth: CanVerPTests
func GetPublicPTests(params published_test.GetPublicPublishedTestsParams, u *models.User) middleware.Responder {
	if permissions.CanVerPTests(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var ts []*dao.Test
			ts, err = dao.GetPublicPublishedTests(db, params.Tags, params.LikeTitle, params.Orderby,
				params.Limit, params.Offset)
			if err == nil {
				var mts []*models.Test
				mts, err = dao.ToModelTests(ts)
				if err == nil {
					return published_test.NewGetPublicPublishedTestsOK().WithPayload(mts)
				}
			}
		}
		log.Println("Error en users_handler GetPublicPTests(): ", err)
		return published_test.NewGetPublicPublishedTestsInternalServerError()
	}
	return published_test.NewGetPublicPublishedTestsForbidden()
}

// GetPTest GET /publicPublishedTests/{testid}. Returns a public published tests.
// Auth: CanVerPTests
func GetPublicPTest(params published_test.GetPublicPublishedTestParams, u *models.User) middleware.Responder {
	if permissions.CanVerPTests(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var ts *dao.Test
			ts, err = dao.GetPublicPublishedTest(db, params.Testid)
			if err == nil && ts != nil {
				var mts *models.Test
				mts, err = dao.ToModelTest(ts)
				if err == nil {
					return published_test.NewGetPublicPublishedTestOK().WithPayload(mts)
				}
			}
			if ts == nil {
				return published_test.NewGetPublicPublishedTestGone()
			}
		}
		log.Println("Error en users_handler GetPublicPTest(): ", err)
		return published_test.NewGetPublicPublishedTestInternalServerError()
	}
	return published_test.NewGetPublicPublishedTestForbidden()
}

// GetPTests GET /publishedTests. Returns all published tests.
// Auth: CanAdminPTests
func GetPTests(params published_test.GetPublishedTestsParams, u *models.User) middleware.Responder {
	db, err := dbconnection.ConnectDb()
	if permissions.CanAdminPTests(u) {
		if err == nil {
			var ts []*dao.Test
			ts, err = dao.GetPublishedTests(db)
			if err == nil {
				var mts []*models.Test
				mts, err = dao.ToModelTests(ts)
				if err == nil {
					return published_test.NewGetPublishedTestsOK().WithPayload(mts)
				}
			}
		}
		log.Println("Error en users_handler GetPTests(): ", err)
		return published_test.NewGetPublishedTestsInternalServerError()
	}
	return published_test.NewGetPublishedTestsForbidden()
}

// GetPTest GET /publishedTests/{testid}. Returns a published test.
// Auth: CanAdminPTests, or Test TestAdmin or TestInvited
func GetPTest(params published_test.GetPublishedTestParams, u *models.User) middleware.Responder {
	db, err := dbconnection.ConnectDb()
	if permissions.CanAdminPTests(u) || isTestAdmin(u, params.Testid) || isTestInvited(u, params.Testid) {
		if err == nil {
			var ts *dao.Test
			ts, err = dao.GetPublishedTest(db, params.Testid)
			if err == nil && ts != nil {
				var mts *models.Test
				mts, err = dao.ToModelTest(ts)
				if err == nil {
					return published_test.NewGetPublishedTestOK().WithPayload(mts)
				}
			}
		}
		log.Println("Error en users_handler GetPTest(): ", err)
		return published_test.NewGetPublishedTestInternalServerError()
	}
	return published_test.NewGetPublishedTestForbidden()
}

// PutTest PUT /publishedTests/{testid}. Updates a ptest.
// Auth: TestAdmin or CanAdminTests
func PutPTest(params published_test.PutPublishedTestParams, u *models.User) middleware.Responder {
	if permissions.CanAdminTests(u) || isTestAdmin(u, params.Testid) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.PutPTest(db, params.Testid, params.PTestUpdate)
			if err == nil {
				return published_test.NewPutPublishedTestOK()
			}
		}
		log.Println("Error en published_test_handler PutPublishedTest(): ", err)
		return published_test.NewPutPublishedTestInternalServerError()
	}
	return published_test.NewPutPublishedTestForbidden()
}

// GetUsersFromPTest GET /publishedTests/{testid}/users. Returns a published tests users.
// Auth: CanVerPTests
func GetUsersFromPTest(params published_test.GetUsersFromPublishedTestParams, u *models.User) middleware.Responder {
	if permissions.CanVerPTests(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var us []*dao.User
			us, err = dao.GetUsersInvitedPTest(db, params.Testid)
			if err == nil {
				mus := dao.ToModelsUser(us)
				if err == nil {
					return published_test.NewGetUsersFromPublishedTestOK().WithPayload(mus)
				}
			}
		}
		log.Println("Error en users_handler GetUsersFromPTest(): ", err)
		return published_test.NewGetUsersFromPublishedTestInternalServerError()
	}
	return published_test.NewGetUsersFromPublishedTestForbidden()
}

// InviteUserPTest PUT /publishedTests/{testid}/users/{username}. Invites a user to solve the test
// Auth: TestAdmin or CanAdminPTests
func InviteUserPTest(params published_test.InviteUserToPublishedTestParams, u *models.User) middleware.Responder {
	if permissions.CanAdminPTests(u) || isTestAdmin(u, params.Testid) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.InviteUserPTest(db, params.Testid, params.Username)
			if err == nil {
				if params.Message == nil || *params.Message.SendEmail {
					emailHelper.SendEmailUserInvitedToTest(params.Username, params.Testid, params.Message)
				}
				return published_test.NewInviteUserToPublishedTestOK()
			}
		}
		log.Println("Error en users_handler InviteUserPTest(): ", err)
		return published_test.NewInviteUserToPublishedTestInternalServerError()
	}
	return published_test.NewInviteUserToPublishedTestForbidden()
}

// RemoveUserPTest DELETE /publishedTests/{testid}/users/{username}. Withdraws invitation to user to solve test
// Auth: TestAdmin or CanAdminPTests
func RemoveUserPTest(params published_test.RemoveUserToPublishedTestParams, u *models.User) middleware.Responder {
	if permissions.CanAdminPTests(u) || isTestAdmin(u, params.Testid) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.RemoveUserPTest(db, params.Testid, params.Username)
			if err == nil {
				return published_test.NewRemoveUserToPublishedTestOK()
			}
		}
		log.Println("Error en users_handler RemoveUserPTest(): ", err)
		return published_test.NewRemoveUserToPublishedTestInternalServerError()
	}
	return published_test.NewRemoveUserToPublishedTestForbidden()
}

// GetTeamsFromPTest GET /publishedTests/{testid}/teams. Returns a published tests teams.
// Auth: CanVerPTests
func GetTeamsFromPTest(params published_test.GetTeamsFromPublishedTestParams, u *models.User) middleware.Responder {
	if permissions.CanVerPTests(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var ts []*dao.Team
			ts, err = dao.GetTeamsInvitedPTest(db, params.Testid)
			if err == nil {
				mts := dao.ToModelsTeams(ts)
				if err == nil {
					return published_test.NewGetTeamsFromPublishedTestOK().WithPayload(mts)
				}
			}
		}
		log.Println("Error en users_handler GetTeamsFromPTest(): ", err)
		return published_test.NewGetTeamsFromPublishedTestInternalServerError()
	}
	return published_test.NewGetTeamsFromPublishedTestForbidden()
}

func sendEmailInvitedTeamMembers(db *sql.DB, teamname string, testid int64, message *models.Message) {
	users, err := dao.GetUsersFromTeam(db, teamname)
	if err == nil {
		for _, u := range users {
			emailHelper.SendEmailUserTeamInvitedToTest(*u.Username, testid, teamname, message)
		}
	}
}

// InviteTeamPTest PUT /publishedTests/{testid}/teams/{teamname}. Invites a team to solve the test
// Auth: TestAdmin or CanAdminPTests
func InviteTeamPTest(params published_test.InviteTeamToPublishedTestParams, u *models.User) middleware.Responder {
	if permissions.CanAdminPTests(u) || isTestAdmin(u, params.Testid) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.InviteTeamPTest(db, params.Testid, params.Teamname)
			if err == nil {
				if params.Message == nil || *params.Message.SendEmail {
					sendEmailInvitedTeamMembers(db, params.Teamname, params.Testid, params.Message)
				}
				return published_test.NewInviteTeamToPublishedTestOK()
			}
		}
		log.Println("Error en users_handler InviteTeamPTest(): ", err)
		return published_test.NewInviteTeamToPublishedTestInternalServerError()
	}
	return published_test.NewInviteTeamToPublishedTestForbidden()
}

// RemoveTeamPTest DELETE /publishedTests/{testid}/teams/{teamname}. Withdraws invitation to team to solve test
// Auth: TestAdmin or CanAdminPTests
func RemoveTeamPTest(params published_test.RemoveTeamToPublishedTestParams, u *models.User) middleware.Responder {
	if permissions.CanAdminPTests(u) || isTestAdmin(u, params.Testid) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.RemoveTeamPTest(db, params.Testid, params.Teamname)
			if err == nil {
				return published_test.NewRemoveTeamToPublishedTestOK()
			}
		}
		log.Println("Error en users_handler RemoveTeamPTest(): ", err)
		return published_test.NewRemoveTeamToPublishedTestInternalServerError()
	}
	return published_test.NewRemoveTeamToPublishedTestForbidden()
}

func isTestInvited(u *models.User, testid int64) bool {
	if u == nil {
		return false
	}
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var t *dao.Test
		t, err = dao.GetSolvableTestFromUser(db, *u.Username, testid)
		if err == nil {
			if t != nil {
				return true
			}
		}
	}
	log.Print("Error en isTestInvited: ", err)
	return false
}

func isTestOpenByUser(u *models.User, testid int64) (bool, error) {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var oa []*dao.Answer
		oa, err = dao.GetOpenAnswersFromUserTest(db, *u.Username, testid)
		if err == nil {
			isTestStarted := (oa != nil) && len(oa) > 0
			return isTestStarted, nil
		}
	}
	return false, err
}

func isTestOpenByUserAuth(u *models.User, testid int64) bool {
	if u == nil {
		return false
	}
	b, e := isTestOpenByUser(u, testid)
	if e == nil {
		return b
	}
	return false
}

func isTestAnsweredByUser(u *models.User, testid int64) (bool, error) {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var t *dao.Test
		t, err = dao.GetATestFromUser(db, *u.Username, testid)
		if err == nil {
			isTestAnswered := t != nil
			return isTestAnswered, nil
		}
	}
	return false, err
}

func isTestAnsweredByUserAuth(u *models.User, testid int64) bool {
	if u == nil {
		return false
	}
	b, e := isTestAnsweredByUser(u, testid)
	if e == nil {
		return b
	}
	return false
}

func hasAnswerVisible(u *models.User, testid int64) (bool, error) {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var oa []*dao.Answer
		oa, err = dao.GetVisibleAnswersFromUserTest(db, *u.Username, testid)
		if err == nil {
			isTestVisible := (oa != nil) && len(oa) > 0
			return isTestVisible, nil
		}
	}
	return false, err
}

func hasAnswerVisibleAuth(u *models.User, tid int64) bool {
	if u == nil {
		return false
	}
	b, e := hasAnswerVisible(u, tid)
	if e == nil {
		return b
	}
	return false
}

// GetQuestionsPTest GET /publishedTests/{testid}/questions
// Auth: (TestAdmin or CanAdminPTests) OR ((CanVerPTests AND accesoPublico OR TestInvited) AND (TestOpenByThem||hasAnswerVisible))
func GetQuestionsPTest(params published_test.GetQuestionsFromPublishedTestsParams, u *models.User) middleware.Responder {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var ts *dao.Test
		ts, err = dao.GetPublishedTest(db, params.Testid)
		if err == nil && ts != nil {
			if !(permissions.CanAdminPTests(u) || isTestAdmin(u, params.Testid)) {
				if !((isTestOpenByUserAuth(u, params.Testid) || hasAnswerVisibleAuth(u, params.Testid)) &&
					(permissions.CanVerPTests(u) && (*ts.AccesoPublico || isTestInvited(u, params.Testid)))) {
					return published_test.NewGetQuestionsFromPublishedTestsForbidden()
				}
			}
			var qs []*dao.Question
			qs, err = dao.GetQuestionsFromTest(db, params.Testid)
			if err == nil {
				var mqs []*models.Question
				mqs, err = dao.ToModelQuestions(qs)
				if err == nil {
					return published_test.NewGetQuestionsFromPublishedTestsOK().WithPayload(mqs)
				}
			}
		}
	}
	return published_test.NewGetQuestionsFromPublishedTestsInternalServerError()
}

// GetQuestionsPTest GET /publishedTests/{testid}/questions/{questionid}
// Auth: (TestAdmin or CanAdminPTests) OR ((CanVerPTests AND accesoPublico OR TestInvited) AND (TestOpenByThem||hasAnswerVisible))
func GetQuestionPTest(params published_test.GetQuestionFromPublishedTestsParams, u *models.User) middleware.Responder {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var ts *dao.Test
		ts, err = dao.GetPublishedTest(db, params.Testid)
		if err == nil && ts != nil {
			if !(permissions.CanAdminPTests(u) || isTestAdmin(u, params.Testid)) {
				if !((isTestOpenByUserAuth(u, params.Testid) || hasAnswerVisibleAuth(u, params.Testid)) && (permissions.CanVerPTests(u) &&
					(*ts.AccesoPublico || isTestInvited(u, params.Testid)))) {
					return published_test.NewGetQuestionFromPublishedTestsForbidden()
				}
			}
			var qs *dao.Question
			qs, err = dao.GetQuestionFromTest(db, params.Testid, params.Questionid)
			if err == nil && qs != nil {
				var mqs *models.Question
				mqs, err = dao.ToModelQuestion(qs)
				if err == nil {
					return published_test.NewGetQuestionFromPublishedTestsOK().WithPayload(mqs)
				}
			}
			return published_test.NewGetQuestionFromPublishedTestsGone()
		}
	}
	return published_test.NewGetQuestionFromPublishedTestsInternalServerError()
}

// GetAnswersPTest GET /publishedTests/{testid}/answers
// Auth: CanVerAnswers if accesoPublico, else: TestAdmin or CanAdminPTests
func GetAnswersPTest(params published_test.GetAnswersFromPublishedTestsParams, u *models.User) middleware.Responder {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var ts *dao.Test
		ts, err = dao.GetPublishedTest(db, params.Testid)
		if err == nil && ts != nil {
			if (!*ts.AccesoPublico && !(permissions.CanAdminPTests(u) || isTestAdmin(u, params.Testid))) ||
				!permissions.CanVerAnswers(u) {
				return published_test.NewGetAnswersFromPublishedTestsForbidden()
			}
			var as []*dao.Answer
			as, err = dao.GetAnswersFromPTest(db, params.Testid)
			if err == nil {
				var mas []*models.Answer
				mas, err = dao.ToModelAnswers(as)
				if err == nil {
					return published_test.NewGetAnswersFromPublishedTestsOK().WithPayload(mas)
				}
			}
		}
	}
	return published_test.NewGetAnswersFromPublishedTestsInternalServerError()
}

// GetCAnswersPTest GET /publishedTests/{testid}/correctedAnswers
// Auth: CanVerAnswers if accesoPublico, else: TestAdmin or CanAdminPTests
func GetCAnswersPTest(params published_test.GetCorrectedAnswersFromPublishedTestsParams, u *models.User) middleware.Responder {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var ts *dao.Test
		ts, err = dao.GetPublishedTest(db, params.Testid)
		if err == nil && ts != nil {
			if (!*ts.AccesoPublico && !(permissions.CanAdminPTests(u) || isTestAdmin(u, params.Testid))) ||
				!permissions.CanVerAnswers(u) {
				return published_test.NewGetCorrectedAnswersFromPublishedTestsForbidden()
			}
			var as []*dao.Answer
			as, err = dao.GetCorrectedAnswersFromPTest(db, params.Testid, params.OrderByAnswer)
			if err == nil {
				var mas []*models.Answer
				mas, err = dao.ToModelAnswers(as)
				if err == nil {
					return published_test.NewGetCorrectedAnswersFromPublishedTestsOK().WithPayload(mas)
				}
			}
		}
	}
	return published_test.NewGetCorrectedAnswersFromPublishedTestsInternalServerError()
}

// GetUCAnswersPTest GET /publishedTests/{testid}/correctedAnswers
// Auth: CanVerAnswers if accesoPublico, else: TestAdmin or CanAdminPTests
func GetUCAnswersPTest(params published_test.GetUncorrectedAnswersFromPublishedTestsParams, u *models.User) middleware.Responder {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var ts *dao.Test
		ts, err = dao.GetPublishedTest(db, params.Testid)
		if err == nil && ts != nil {
			if (!*ts.AccesoPublico && !(permissions.CanAdminPTests(u) || isTestAdmin(u, params.Testid))) ||
				!permissions.CanVerAnswers(u) {
				return published_test.NewGetUncorrectedAnswersFromPublishedTestsForbidden()
			}
			var as []*dao.Answer
			as, err = dao.GetUncorrectedAnswersFromPTest(db, params.Testid, params.OrderByAnswer)
			if err == nil {
				var mas []*models.Answer
				mas, err = dao.ToModelAnswers(as)
				if err == nil {
					return published_test.NewGetUncorrectedAnswersFromPublishedTestsOK().WithPayload(mas)
				}
			}
		}
	}
	return published_test.NewGetUncorrectedAnswersFromPublishedTestsInternalServerError()
}

// GetQuestionAnswersPTest GET /publishedTests/{testid}/questions/{questionid}/qanswers
// Auth: CanVerAnswers if accesoPublico, else: TestAdmin or CanAdminPTests
func GetQuestionAnswersPTest(params published_test.GetQuestionAnswersFromPublishedTestQuestionParams, u *models.User) middleware.Responder {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var ts *dao.Test
		ts, err = dao.GetPublishedTest(db, params.Testid)
		if err == nil && ts != nil {
			if (!*ts.AccesoPublico && !(permissions.CanAdminPTests(u) || isTestAdmin(u, params.Testid))) ||
				!permissions.CanVerAnswers(u) {
				return published_test.NewGetQuestionAnswersFromPublishedTestQuestionForbidden()
			}
			var as []*dao.QuestionAnswer
			as, err = dao.GetQuestionAnswersFromPTestQuestion(db, params.Testid, params.Questionid, params.LikeUsername)
			if err == nil {
				mas := dao.ToModelQuestionAnswers(as)
				if err == nil {
					return published_test.NewGetQuestionAnswersFromPublishedTestQuestionOK().WithPayload(mas)
				}
			}
		}
	}
	return published_test.NewGetQuestionAnswersFromPublishedTestQuestionInternalServerError()
}

// GetOptionsPQuestion GET /publishedTests/{testid}/questions/{questionid}/options
// Auth: CanVerPTests if accesoPublico, else: TestInvited, TestAdmin or CanAdminPTests
// Info: If HasAnswerVisible -> Will know which options are correct. If not, all options as non correct
func GetOptionsPQuestion(params published_test.GetOptionsFromPublishedQuestionParams, u *models.User) middleware.Responder {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var ts *dao.Test
		ts, err = dao.GetPublishedTest(db, params.Testid)
		if err == nil && ts != nil {
			if !*ts.AccesoPublico {
				if !(permissions.CanAdminPTests(u) || isTestAdmin(u, params.Testid) || isTestInvited(u, params.Testid)) {
					return published_test.NewGetOptionsFromPublishedQuestionForbidden()
				}
			} else if !permissions.CanVerPTests(u) {
				return published_test.NewGetOptionsFromPublishedQuestionForbidden()
			}
			var qs *dao.Question
			qs, err = dao.GetQuestionFromTest(db, params.Testid, params.Questionid)
			if err == nil && qs != nil {
				var os []*dao.Option
				os, err = dao.GetOptionsQuestion(db, params.Questionid)
				if err == nil {
					var mos []*models.Option
					if hasAnswerVisibleAuth(u, params.Testid) {
						mos = dao.ToModelOptions(os)
					} else {
						mos = dao.ToModelOptionsNoCorrect(os)
					}
					return published_test.NewGetOptionsFromPublishedQuestionOK().WithPayload(mos)
				}
			}
			return published_test.NewGetOptionsFromPublishedQuestionGone()
		}
	}
	return published_test.NewGetOptionsFromPublishedQuestionInternalServerError()
}

// GetOptionsPQuestion GET /publishedTests/{testid}/questions/{questionid}/tags
// Auth: CanVerPTests if accesoPublico, else: TestInvited, TestAdmin or CanAdminPTests
func GetTagsPQuestion(params published_test.GetTagsFromPublishedQuestionParams, u *models.User) middleware.Responder {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var ts *dao.Test
		ts, err = dao.GetPublishedTest(db, params.Testid)
		if err == nil && ts != nil {
			if !*ts.AccesoPublico {
				if !(permissions.CanAdminPTests(u) || isTestAdmin(u, params.Testid) || isTestInvited(u, params.Testid)) {
					return published_test.NewGetTagsFromPublishedQuestionForbidden()
				}
			} else if !permissions.CanVerPTests(u) {
				return published_test.NewGetTagsFromPublishedQuestionForbidden()
			}
			var qs *dao.Question
			qs, err = dao.GetQuestionFromTest(db, params.Testid, params.Questionid)
			if err == nil && qs != nil {
				var ts []*dao.Tag
				ts, err = dao.GetQuestionTags(db, params.Questionid)
				if err == nil {
					mts := dao.ToModelTags(ts)
					return published_test.NewGetTagsFromPublishedQuestionOK().WithPayload(mts)
				}
			}
			return published_test.NewGetTagsFromPublishedQuestionGone()
		}
	}
	return published_test.NewGetTagsFromPublishedQuestionInternalServerError()
}

// GET /publishedTests/{testid}/tags
// Auth: CanVerPTests if accesoPublico, else: TestInvited, TestAdmin or CanAdminPTests
func GetTagsFromPTest(params published_test.GetTagsFromPublishedTestParams, u *models.User) middleware.Responder {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var ts *dao.Test
		ts, err = dao.GetPublishedTest(db, params.Testid)
		if err == nil && ts != nil {
			if !*ts.AccesoPublico {
				if !(permissions.CanAdminPTests(u) || isTestAdmin(u, params.Testid) || isTestInvited(u, params.Testid)) {
					return published_test.NewGetTagsFromPublishedTestForbidden()
				}
			} else if !permissions.CanVerPTests(u) {
				return published_test.NewGetTagsFromPublishedTestForbidden()
			}
			var tags []*dao.Tag
			tags, err = dao.GetTestTags(db, params.Testid)
			if err == nil {
				mts := dao.ToModelTags(tags)
				return published_test.NewGetTagsFromPublishedTestOK().WithPayload(mts)
			}
			return published_test.NewGetTagsFromPublishedTestGone()
		}
	}
	return published_test.NewGetTagsFromPublishedTestInternalServerError()
}

// GET /publishedTests/{testid}/tags/{tag}
// Auth: CanVerPTests if accesoPublico, else: TestInvited, TestAdmin or CanAdminPTests
func GetTagFromPTest(params published_test.GetTagFromPublishedTestParams, u *models.User) middleware.Responder {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var ts *dao.Test
		ts, err = dao.GetPublishedTest(db, params.Testid)
		if err == nil && ts != nil {
			if !*ts.AccesoPublico {
				if !(permissions.CanAdminPTests(u) || isTestAdmin(u, params.Testid) || isTestInvited(u, params.Testid)) {
					return published_test.NewGetTagFromPublishedTestForbidden()
				}
			} else if !permissions.CanVerPTests(u) {
				return published_test.NewGetTagFromPublishedTestForbidden()
			}
			var tag *dao.Tag
			tag, err = dao.GetTestTag(db, params.Testid, params.Tag)
			if err == nil {
				mt := dao.ToModelTag(tag)
				return published_test.NewGetTagFromPublishedTestOK().WithPayload(mt)
			}
			return published_test.NewGetTagFromPublishedTestGone()
		}
	}
	return published_test.NewGetTagFromPublishedTestInternalServerError()
}

// GetVisiblePruebas GET /publishedTests/{testid}/questions/{questionid}/pruebas
// Auth: CanVerPTests if accesoPublico, else: TestInvited, TestAdmin or CanAdminPTests
func GetVisiblePruebas(params published_test.GetVisiblePruebasFromQuestionTestParams, u *models.User) middleware.Responder {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var ts *dao.Test
		ts, err = dao.GetPublishedTest(db, params.Testid)
		if err == nil && ts != nil {
			if !*ts.AccesoPublico {
				if !(permissions.CanAdminPTests(u) || isTestAdmin(u, params.Testid) || isTestInvited(u, params.Testid)) {
					return published_test.NewGetVisiblePruebasFromQuestionTestForbidden()
				}
			} else if !permissions.CanVerPTests(u) {
				return published_test.NewGetVisiblePruebasFromQuestionTestForbidden()
			}
			var qs *dao.Question
			qs, err = dao.GetQuestionFromTest(db, params.Testid, params.Questionid)
			if err == nil && qs != nil {
				var ps []*dao.Prueba
				ps, err = dao.GetVisiblePruebasQuestion(db, params.Questionid)
				if err == nil {
					var mps []*models.Prueba
					mps = dao.ToModelPruebas(ps)
					return published_test.NewGetVisiblePruebasFromQuestionTestOK().WithPayload(mps)
				}
			}
			return published_test.NewGetVisiblePruebasFromQuestionTestGone()
		}
	}
	return published_test.NewGetVisiblePruebasFromQuestionTestInternalServerError()
}
