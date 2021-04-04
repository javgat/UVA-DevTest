// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package handlers provides functions that handle http requests
package handlers

import (
	"log"
	"uva-devtest/models"
	"uva-devtest/persistence/dao"
	"uva-devtest/persistence/dbconnection"
	"uva-devtest/restapi/operations/published_test"

	"github.com/go-openapi/runtime/middleware"
)

// GetPTests GET /publishedTests. Returns all published tests.
// Auth: ALL
func GetPTests(params published_test.GetPublishedTestsParams, u *models.User) middleware.Responder {
	db, err := dbconnection.ConnectDb()
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

// GetPTest GET /publishedTests/{testid}. Returns a published tests.
// Auth: ALL
func GetPTest(params published_test.GetPublishedTestParams, u *models.User) middleware.Responder {
	db, err := dbconnection.ConnectDb()
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

// GetUsersFromPTest GET /publishedTests/{testid}/users. Returns a published tests users.
// Auth: ALL
func GetUsersFromPTest(params published_test.GetUsersFromPublishedTestParams, u *models.User) middleware.Responder {
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

// InviteUserPTest PUT /publishedTests/{testid}/users/{username}. Invites a user to solve the test
// Auth: TestAdmin or Admin
func InviteUserPTest(params published_test.InviteUserToPublishedTestParams, u *models.User) middleware.Responder {
	if isAdmin(u) || isTestAdmin(u, params.Testid) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.InviteUserPTest(db, params.Testid, params.Username)
			if err == nil {
				return published_test.NewInviteUserToPublishedTestOK()
			}
		}
		log.Println("Error en users_handler InviteUserPTest(): ", err)
		return published_test.NewInviteUserToPublishedTestInternalServerError()
	}
	return published_test.NewInviteUserToPublishedTestForbidden()
}

// RemoveUserPTest DELETE /publishedTests/{testid}/users/{username}. Withdraws invitation to user to solve test
// Auth: TestAdmin or Admin
func RemoveUserPTest(params published_test.RemoveUserToPublishedTestParams, u *models.User) middleware.Responder {
	if isAdmin(u) || isTestAdmin(u, params.Testid) {
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
// Auth: ALL
func GetTeamsFromPTest(params published_test.GetTeamsFromPublishedTestParams, u *models.User) middleware.Responder {
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

// InviteTeamPTest PUT /publishedTests/{testid}/teams/{teamname}. Invites a team to solve the test
// Auth: TestAdmin or Admin
func InviteTeamPTest(params published_test.InviteTeamToPublishedTestParams, u *models.User) middleware.Responder {
	if isAdmin(u) || isTestAdmin(u, params.Testid) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.InviteTeamPTest(db, params.Testid, params.Teamname)
			if err == nil {
				return published_test.NewInviteTeamToPublishedTestOK()
			}
		}
		log.Println("Error en users_handler InviteTeamPTest(): ", err)
		return published_test.NewInviteTeamToPublishedTestInternalServerError()
	}
	return published_test.NewInviteTeamToPublishedTestForbidden()
}

// RemoveTeamPTest DELETE /publishedTests/{testid}/teams/{teamname}. Withdraws invitation to team to solve test
// Auth: TestAdmin or Admin
func RemoveTeamPTest(params published_test.RemoveTeamToPublishedTestParams, u *models.User) middleware.Responder {
	if isAdmin(u) || isTestAdmin(u, params.Testid) {
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
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var t *dao.Test
		t, err = dao.GetPTestFromUser(db, *u.Username, testid)
		if err == nil {
			if t != nil {
				return true
			}
			var tms []*dao.Team
			tms, err = dao.GetTeamsUsername(db, *u.Username)
			if err == nil {
				for _, tm := range tms {
					t, err = dao.GetPTestFromTeam(db, *tm.Teamname, testid)
					if err != nil {
						log.Print("Error en isTestInvited: ", err)
						return false
					}
					if t != nil {
						return true
					}
				}
				return false
			}
		}
	}
	log.Print("Error en isTestInvited: ", err)
	return false
}

// GetQuestionsPTest GET /publishedTests/{testid}/questions
// Auth: ALL if accesoPublico, else: TestInvited, TestAdmin or Admin
func GetQuestionsPTest(params published_test.GetQuestionsFromPublishedTestsParams, u *models.User) middleware.Responder {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var ts *dao.Test
		ts, err = dao.GetPublishedTest(db, params.Testid)
		if err == nil && ts != nil {
			if !*ts.AccesoPublico {
				if !(isAdmin(u) || isTestAdmin(u, params.Testid) || isTestInvited(u, params.Testid)) {
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
// Auth: ALL if accesoPublico, else: TestInvited, TestAdmin or Admin
func GetQuestionPTest(params published_test.GetQuestionFromPublishedTestsParams, u *models.User) middleware.Responder {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var ts *dao.Test
		ts, err = dao.GetPublishedTest(db, params.Testid)
		if err == nil && ts != nil {
			if !*ts.AccesoPublico {
				if !(isAdmin(u) || isTestAdmin(u, params.Testid) || isTestInvited(u, params.Testid)) {
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
// Auth: Teacher or Admin if accesoPublico, else: TestAdmin or Admin
func GetAnswersPTest(params published_test.GetAnswersFromPublishedTestsParams, u *models.User) middleware.Responder {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var ts *dao.Test
		ts, err = dao.GetPublishedTest(db, params.Testid)
		if err == nil && ts != nil {
			if (!*ts.AccesoPublico && !(isAdmin(u) || isTestAdmin(u, params.Testid))) ||
				!isTeacherOrAdmin(u) {
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

// GetQuestionAnswersPTest GET /publishedTests/{testid}/questions/{questionid}/qanswers
// Auth: Teacher or Admin if accesoPublico, else: TestAdmin or Admin
func GetQuestionAnswersPTest(params published_test.GetQuestionAnswersFromPublishedTestQuestionParams, u *models.User) middleware.Responder {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var ts *dao.Test
		ts, err = dao.GetPublishedTest(db, params.Testid)
		if err == nil && ts != nil {
			if (!*ts.AccesoPublico && !(isAdmin(u) || isTestAdmin(u, params.Testid))) ||
				!isTeacherOrAdmin(u) {
				return published_test.NewGetQuestionAnswersFromPublishedTestQuestionForbidden()
			}
			var as []*dao.QuestionAnswer
			as, err = dao.GetQuestionAnswersFromPTestQuestion(db, params.Testid, params.Questionid)
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
// Auth: ALL if accesoPublico, else: TestInvited, TestAdmin or Admin
func GetOptionsPQuestion(params published_test.GetOptionsFromPublishedQuestionParams, u *models.User) middleware.Responder {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var ts *dao.Test
		ts, err = dao.GetPublishedTest(db, params.Testid)
		if err == nil && ts != nil {
			if !*ts.AccesoPublico {
				if !(isAdmin(u) || isTestAdmin(u, params.Testid) || isTestInvited(u, params.Testid)) {
					return published_test.NewGetOptionsFromPublishedQuestionForbidden()
				}
			}
			var qs *dao.Question
			qs, err = dao.GetQuestionFromTest(db, params.Testid, params.Questionid)
			if err == nil && qs != nil {
				var os []*dao.Option
				os, err = dao.GetOptionsQuestion(db, params.Questionid)
				if err == nil {
					mos := dao.ToModelOptionsNoCorrect(os)
					return published_test.NewGetOptionsFromPublishedQuestionOK().WithPayload(mos)
				}
			}
			return published_test.NewGetOptionsFromPublishedQuestionGone()
		}
	}
	return published_test.NewGetOptionsFromPublishedQuestionInternalServerError()
}

// GetOptionsPQuestion GET /publishedTests/{testid}/questions/{questionid}/tags
// Auth: ALL if accesoPublico, else: TestInvited, TestAdmin or Admin
func GetTagsPQuestion(params published_test.GetTagsFromPublishedQuestionParams, u *models.User) middleware.Responder {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var ts *dao.Test
		ts, err = dao.GetPublishedTest(db, params.Testid)
		if err == nil && ts != nil {
			if !*ts.AccesoPublico {
				if !(isAdmin(u) || isTestAdmin(u, params.Testid) || isTestInvited(u, params.Testid)) {
					return published_test.NewGetTagsFromPublishedQuestionForbidden()
				}
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
