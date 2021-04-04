// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package handlers provides functions that handle http requests
package handlers

import (
	"database/sql"
	"errors"
	"log"
	"uva-devtest/models"
	"uva-devtest/persistence/dao"
	"uva-devtest/persistence/dbconnection"
	"uva-devtest/restapi/operations/test"

	"github.com/go-openapi/runtime/middleware"
)

// GetTests GET /tests. Returns all tests.
// Auth: Teacher or Admin
func GetTests(params test.GetTestsParams, u *models.User) middleware.Responder {
	if isTeacherOrAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var ts []*dao.Test
			ts, err = dao.GetTests(db)
			if err == nil {
				var mts []*models.Test
				mts, err = dao.ToModelTests(ts)
				if err == nil {
					return test.NewGetTestsOK().WithPayload(mts)
				}
			}
		}
		log.Println("Error en users_handler GetTests(): ", err)
		return test.NewGetTestsInternalServerError()
	}
	return test.NewGetTestsForbidden()
}

// GetTest GET /tests/{testid}. Returns a test.
// Auth: Teacher or Admin
func GetTest(params test.GetTestParams, u *models.User) middleware.Responder {
	if isTeacherOrAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var ts *dao.Test
			ts, err = dao.GetTest(db, params.Testid)
			if err == nil && ts != nil {
				var mts *models.Test
				mts, err = dao.ToModelTest(ts)
				if err == nil {
					return test.NewGetTestOK().WithPayload(mts)
				}
			}
			log.Println("Error en users_handler GetTest(): ", err)
			return test.NewGetTestGone()
		}
		return test.NewGetTestInternalServerError()
	}
	return test.NewGetTestForbidden()
}

func isTestAdmin(u *models.User, testid int64) bool {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var t *dao.Test
		t, err = dao.GetTestFromUser(db, *u.Username, testid)
		if t != nil && err == nil {
			return true
		}
		var ts []*dao.Team
		ts, err = dao.GetTeamsUsername(db, *u.Username)
		if err == nil {
			for _, itemCopy := range ts {
				t, err = dao.GetTestFromTeam(db, *itemCopy.Teamname, testid)
				if t != nil && err == nil {
					return true
				}
			}
		}
	}
	return false
}

// PutTest PUT /tests/{testid}. Updates a test.
// Auth: TestAdmin or Admin
// Req: !Test.editable -> en SQL
func PutTest(params test.PutTestParams, u *models.User) middleware.Responder {
	if isAdmin(u) || isTestAdmin(u, params.Testid) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.PutTest(db, params.Testid, params.Test)
			if err == nil {
				return test.NewPutTestOK()
			}
		}
		log.Println("Error en users_handler PutTest(): ", err)
		return test.NewPutTestInternalServerError()
	}
	return test.NewPutTestForbidden()
}

// DeleteTest DELETE /tests/{testid}. Deletes a test.
// Auth: TestAdmin or Admin
func DeleteTest(params test.DeleteTestParams, u *models.User) middleware.Responder {
	if isAdmin(u) || isTestAdmin(u, params.Testid) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.DeleteTest(db, params.Testid)
			if err == nil {
				return test.NewDeleteTestOK()
			}
		}
		log.Println("Error en users_handler DeleteTest(): ", err)
		return test.NewDeleteTestInternalServerError()
	}
	return test.NewDeleteTestForbidden()
}

// GetTeamsFromTest GET /tests/{testid}/teams. Get teams that admin a test
// Auth: Teacher or Admin
func GetTeamsFromTest(params test.GetTeamsFromTestParams, u *models.User) middleware.Responder {
	if isTeacherOrAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var ts []*dao.Team
			ts, err = dao.GetTeamsFromTest(db, params.Testid)
			var mts []*models.Team
			if err == nil {
				mts = dao.ToModelsTeams(ts)
				return test.NewGetTeamsFromTestOK().WithPayload(mts)
			}
		}
		log.Println("Error en users_handler GetTeamsFromTest(): ", err)
		return test.NewGetTeamsFromTestInternalServerError()
	}
	return test.NewGetTeamsFromTestForbidden()
}

// AddTeamToTest PUT /tests/{testid}/teams/{teamname}. Adds team to admin a test
// Auth: TestAdmin or Admin
func AddTeamToTest(params test.AddTeamToTestParams, u *models.User) middleware.Responder {
	if isAdmin(u) || isTestAdmin(u, params.Testid) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.AddTeamToTest(db, params.Testid, params.Teamname)
			if err == nil {
				return test.NewAddTeamToTestOK()
			}
		}
		log.Println("Error en users_handler AddTeamToTest(): ", err)
		return test.NewAddTeamToTestInternalServerError()
	}
	return test.NewAddTeamToTestForbidden()
}

// RemoveTeamTest DELETE /tests/{testid}/teams/{teamname}. Removes team from admin a test
// Auth: TestAdmin or Admin
func RemoveTeamTest(params test.RemoveTeamToTestParams, u *models.User) middleware.Responder {
	if isAdmin(u) || isTestAdmin(u, params.Testid) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.RemoveTeamFromTest(db, params.Testid, params.Teamname)
			if err == nil {
				return test.NewRemoveTeamToTestOK()
			}
		}
		log.Println("Error en users_handler RemoveTeamTest(): ", err)
		return test.NewRemoveTeamToTestInternalServerError()
	}
	return test.NewRemoveTeamToTestForbidden()
}

func cloneQuestions(db *sql.DB, mqs []*models.Question, newMTest *models.Test, oldDTest *dao.Test) error {
	bfalse := false
	for _, question := range mqs {
		origqid := question.ID
		question.Editable = &bfalse
		qp, err := dao.PostQuestion(db, question, *newMTest.Username)
		if err == nil {
			if qp != nil {
				newqid := qp.ID
				var vF *int64
				vF, err = dao.GetValorFinal(db, origqid, oldDTest.ID)
				if err == nil {
					if vF != nil {
						err = dao.AddQuestionTest(db, newqid, newMTest.ID, *vF)
						if err == nil {
							var tags []*dao.Tag
							tags, err = dao.GetQuestionTags(db, origqid)
							if err == nil {
								for _, tag := range tags {
									err = dao.AddQuestionTag(db, newqid, *tag.Tag)
									if err != nil {
										return err
									}
								}
								var opciones []*dao.Option
								opciones, err = dao.GetOptionsQuestion(db, origqid)
								for _, opc := range opciones {
									_, err = dao.PostOption(db, newqid, dao.ToModelOption(opc))
									if err != nil {
										return err
									}
								}
							}
						}
					} else {
						err = errors.New("valor final no se pudo obtener")
					}
				}
			} else {
				err = errors.New("valor final no se pudo obtener")
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// PublishTest POST /tests/{testid}/publishedTests. Copies test and questions to published version
// Auth: TestAdmin or Admin
func PublishTest(params test.PostPublishedTestParams, u *models.User) middleware.Responder {
	if isAdmin(u) || isTestAdmin(u, params.Testid) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var oldDaoTest *dao.Test
			var newModelTest *models.Test
			oldDaoTest, err = dao.GetTest(db, params.Testid)
			if err == nil && oldDaoTest != nil {
				var qs []*dao.Question
				qs, err = dao.GetQuestionsFromTest(db, params.Testid)
				if err == nil {
					var mqs []*models.Question
					mqs, err = dao.ToModelQuestions(qs)
					if err == nil {
						newModelTest, err = dao.ToModelTest(oldDaoTest)
						if err == nil {
							bfalse := false
							newModelTest.Editable = &bfalse
							newModelTest, err = dao.PostTest(db, *newModelTest.Username, newModelTest)
							if err == nil {
								err = cloneQuestions(db, mqs, newModelTest, oldDaoTest)
								if err == nil {
									teams, err := dao.GetTeamsFromTest(db, params.Testid)
									if err == nil {
										for _, team := range teams {
											if err == nil {
												err = dao.AddTeamToTest(db, newModelTest.ID, *team.Teamname)
											}
											if err != nil {
												log.Println("Error en users_handler PublishTest(): ", err)
												return test.NewPostPublishedTestInternalServerError()
											}
										}
									}
									return test.NewPostPublishedTestCreated().WithPayload(newModelTest)
								}
							}
						}
					}
				}
			}
		}
		log.Println("Error en users_handler PublishTest(): ", err)
		return test.NewPostPublishedTestInternalServerError()
	}
	return test.NewPostPublishedTestForbidden()
}

// GetQuestionsFromTest GET /tests/{testid}/questions. Get questions from test
// Auth: Teacher or Admin
func GetQuestionsFromTest(params test.GetQuestionsFromTestParams, u *models.User) middleware.Responder {
	if isTeacherOrAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var qs []*dao.Question
			qs, err = dao.GetQuestionsFromTest(db, params.Testid)
			var mqs []*models.Question
			if err == nil {
				mqs, err = dao.ToModelQuestions(qs)
				if err == nil {
					return test.NewGetQuestionsFromTestOK().WithPayload(mqs)
				}
			}
		}
		log.Println("Error en users_handler GetQuestionsFromTest(): ", err)
		return test.NewGetQuestionsFromTestInternalServerError()
	}
	return test.NewGetQuestionsFromTestForbidden()
}

// GetQuestionFromTest GET /tests/{testid}/questions/{questionid}. Get question from test
// Auth: Teacher or Admin
func GetQuestionFromTest(params test.GetQuestionFromTestParams, u *models.User) middleware.Responder {
	if isTeacherOrAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var qs *dao.Question
			qs, err = dao.GetQuestionFromTest(db, params.Testid, params.Questionid)
			var mqs *models.Question
			if err == nil && qs != nil {
				mqs, err = dao.ToModelQuestion(qs)
				if err == nil {
					return test.NewGetQuestionFromTestOK().WithPayload(mqs)
				}
			}
		}
		log.Println("Error en users_handler GetQuestionFromTest(): ", err)
		return test.NewGetQuestionFromTestInternalServerError()
	}
	return test.NewGetQuestionFromTestForbidden()
}

func testEditable(testid int64) bool {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		t, err := dao.GetTest(db, testid)
		if err == nil && t != nil {
			return *t.Editable
		}
	}
	return false
}

// AddQuestionToTest PUT /tests/{testid}/questions/{questionid}. Add question to test
// Auth: TestAdmin or Admin
// Req: Test.editable
func AddQuestionToTest(params test.AddQuestionToTestParams, u *models.User) middleware.Responder {
	if testEditable(params.Testid) && (isAdmin(u) || isTestAdmin(u, params.Testid)) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.AddQuestionTest(db, params.Testid, params.Questionid, *params.ValorFinal.ValorFinal)
			if err == nil {
				return test.NewAddQuestionToTestOK()
			}
			log.Println("Error en users_handler AddQuestionToTest(): ", err)
			return test.NewAddQuestionToTestGone()
		}
	}
	return test.NewAddQuestionToTestInternalServerError()
}

// RemoveQuestionTest DELETE /tests/{testid}/questions/{questionid}. Remove question from test
// Auth: TestAdmin or Admin
// Req: Test.editable
func RemoveQuestionTest(params test.RemoveQuestionFromTestParams, u *models.User) middleware.Responder {
	if testEditable(params.Testid) && (isAdmin(u) || isTestAdmin(u, params.Testid)) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.RemoveQuestionTest(db, params.Testid, params.Questionid)
			if err == nil {
				return test.NewRemoveQuestionFromTestOK()
			}
			log.Println("Error en users_handler AddQuestionToTest(): ", err)
			return test.NewRemoveQuestionFromTestGone()
		}
	}
	return test.NewRemoveQuestionFromTestInternalServerError()
}
