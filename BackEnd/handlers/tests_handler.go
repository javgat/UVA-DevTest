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
	"uva-devtest/restapi/operations/user"

	"github.com/go-openapi/runtime/middleware"
)

// GetAllTests GET /tests. Returns all tests.
// Auth: Admin
func GetAllTests(params test.GetAllTestsParams, u *models.User) middleware.Responder {
	if isTeacherOrAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var ts []*dao.Test
			ts, err = dao.GetAllTests(db)
			if err == nil {
				var mts []*models.Test
				mts, err = dao.ToModelTests(ts)
				if err == nil {
					return test.NewGetAllTestsOK().WithPayload(mts)
				}
			}
		}
		log.Println("Error en users_handler GetAllTests(): ", err)
		return test.NewGetAllTestsInternalServerError()
	}
	return test.NewGetAllTestsForbidden()
}

// GetAllEditTests GET /editTests. Returns all non-published tests.
// Auth: Admin
func GetAllEditTests(params test.GetAllEditTestsParams, u *models.User) middleware.Responder {
	if isTeacherOrAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var ts []*dao.Test
			ts, err = dao.GetAllEditTests(db)
			if err == nil {
				var mts []*models.Test
				mts, err = dao.ToModelTests(ts)
				if err == nil {
					return test.NewGetAllEditTestsOK().WithPayload(mts)
				}
			}
		}
		log.Println("Error en users_handler GetAllEditTests(): ", err)
		return test.NewGetAllEditTestsInternalServerError()
	}
	return test.NewGetAllEditTestsForbidden()
}

// GetPublicTests GET /publicTests. Returns all public tests.
// Auth: Teacher or Admin
func GetPublicTests(params test.GetPublicTestsParams, u *models.User) middleware.Responder {
	if isTeacherOrAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var ts []*dao.Test
			ts, err = dao.GetPublicTests(db)
			if err == nil {
				var mts []*models.Test
				mts, err = dao.ToModelTests(ts)
				if err == nil {
					return test.NewGetPublicTestsOK().WithPayload(mts)
				}
			}
		}
		log.Println("Error en users_handler GetTests(): ", err)
		return test.NewGetPublicTestsInternalServerError()
	}
	return test.NewGetPublicTestsForbidden()
}

// GetPublicEditTests GET /publicEditTests. Returns all public non-published tests.
// Auth: Teacher or Admin
func GetPublicEditTests(params test.GetPublicEditTestsParams, u *models.User) middleware.Responder {
	if isTeacherOrAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var ts []*dao.Test
			ts, err = dao.GetPublicEditTests(db, params.Tags, params.LikeTitle)
			if err == nil {
				var mts []*models.Test
				mts, err = dao.ToModelTests(ts)
				if err == nil {
					return test.NewGetPublicEditTestsOK().WithPayload(mts)
				}
			}
		}
		log.Println("Error en users_handler GetPublicEditTests(): ", err)
		return test.NewGetPublicEditTestsInternalServerError()
	}
	return test.NewGetPublicEditTestsForbidden()
}

// GetTest GET /tests/{testid}. Returns a test.
// Auth: Teacher or Admin si accesoPublicoNoPublicada=true, admin o testAdmin si false
func GetTest(params test.GetTestParams, u *models.User) middleware.Responder {
	if isTeacherOrAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var ts *dao.Test
			ts, err = dao.GetTest(db, params.Testid)
			if err == nil && ts != nil {
				if !*ts.AccesoPublicoNoPublicado {
					if !(isAdmin(u) || isTestAdmin(u, params.Testid)) {
						return test.NewGetTestForbidden()
					}
				}

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

// GetAdminTeamsFromTest GET /tests/{testid}/adminTeams. Get teams that admin a test
// Auth: Teacher or Admin
func GetAdminTeamsFromTest(params test.GetAdminTeamsFromTestParams, u *models.User) middleware.Responder {
	if isTeacherOrAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var ts []*dao.Team
			ts, err = dao.GetAdminTeamsFromTest(db, params.Testid)
			var mts []*models.Team
			if err == nil {
				mts = dao.ToModelsTeams(ts)
				return test.NewGetAdminTeamsFromTestOK().WithPayload(mts)
			}
		}
		log.Println("Error en users_handler GetAdminTeamsFromTest(): ", err)
		return test.NewGetAdminTeamsFromTestInternalServerError()
	}
	return test.NewGetAdminTeamsFromTestForbidden()
}

// AddAdminTeamToTest PUT /tests/{testid}/adminTeams/{teamname}. Adds team to admin a test
// Auth: TestAdmin or Admin
func AddAdminTeamToTest(params test.AddAdminTeamToTestParams, u *models.User) middleware.Responder {
	if isAdmin(u) || isTestAdmin(u, params.Testid) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			if isTeamSoloProfesores(params.Teamname) {
				err = dao.AddAdminTeamToTest(db, params.Testid, params.Teamname)
				if err == nil {
					return test.NewAddAdminTeamToTestOK()
				}
			}
			s := "El equipo a añadir tiene que ser de profesores"
			if err != nil {
				s = "El equipo ya esta añadido"
			} else {
				var t *dao.Team
				t, err = dao.GetTeam(db, params.Teamname)
				if err != nil || t == nil {
					s = "El equipo no existe"
				}
			}
			return test.NewAddAdminTeamToTestBadRequest().WithPayload(&models.Error{Message: &s})
		}
		log.Println("Error en users_handler AddAdminTeamToTest(): ", err)
		return test.NewAddAdminTeamToTestInternalServerError()
	}
	return test.NewAddAdminTeamToTestForbidden()
}

// RemoveAdminTeamTest DELETE /tests/{testid}/adminTeams/{teamname}. Removes team from admin a test
// Auth: TestAdmin or Admin
func RemoveAdminTeamTest(params test.RemoveAdminTeamToTestParams, u *models.User) middleware.Responder {
	if isAdmin(u) || isTestAdmin(u, params.Testid) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.RemoveAdminTeamFromTest(db, params.Testid, params.Teamname)
			if err == nil {
				return test.NewRemoveAdminTeamToTestOK()
			}
		}
		log.Println("Error en users_handler RemoveTeamTest(): ", err)
		return test.NewRemoveAdminTeamToTestInternalServerError()
	}
	return test.NewRemoveAdminTeamToTestForbidden()
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

func addAdminTeamsToTest(db *sql.DB, teams []*dao.Team, testid int64) error {
	var err error
	for _, team := range teams {
		if err == nil {
			err = dao.AddAdminTeamToTest(db, testid, *team.Teamname)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func addTagsToTest(db *sql.DB, tags []*dao.Tag, testid int64) error {
	var err error
	for _, tag := range tags {
		if err == nil {
			err = dao.AddTestTag(db, testid, *tag.Tag)
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
									var teams []*dao.Team
									teams, err = dao.GetAdminTeamsFromTest(db, params.Testid)
									if err == nil {
										err = addAdminTeamsToTest(db, teams, newModelTest.ID)
										if err == nil {
											var tags []*dao.Tag
											tags, err = dao.GetTestTags(db, params.Testid)
											if err == nil {
												err = addTagsToTest(db, tags, newModelTest.ID)
												if err == nil {
													return test.NewPostPublishedTestCreated().WithPayload(newModelTest)
												}
											}
										}
									}
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

func addQuestionTimeToTest(qt *dao.Question, testid int64) error {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var t *dao.Test
		t, err = dao.GetTest(db, testid)
		if t != nil && err == nil {
			var mt *models.Test
			mt, err = dao.ToModelTest(t)
			if err == nil && mt != nil {
				nuevosMins := *mt.MaxMinutes + *qt.EstimatedTime
				mt.MaxMinutes = &nuevosMins
				err = dao.PutTest(db, testid, mt)
				return err
			}
		}
	}
	return err
}

// AddQuestionToTest PUT /tests/{testid}/questions/{questionid}. Add question to test
// Auth: TestAdmin or Admin. Si question no publica => además questionAdmin or Admin, o que ya esté en el test
// Req: Test.editable
func AddQuestionToTest(params test.AddQuestionToTestParams, u *models.User) middleware.Responder {
	if testEditable(params.Testid) && (isAdmin(u) || isTestAdmin(u, params.Testid)) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var q, qt *dao.Question
			q, err = dao.GetQuestion(db, params.Questionid)
			if q != nil && err == nil {
				puedeHacerse := false
				qt, err = dao.GetQuestionFromTest(db, params.Questionid, params.Testid)
				if qt != nil && err == nil {
					puedeHacerse = true
				} else if *q.AccesoPublicoNoPublicada {
					puedeHacerse = true
				} else if isAdmin(u) || isQuestionAdmin(u, params.Questionid) {
					puedeHacerse = true
				}
				if puedeHacerse {
					if qt != nil {
						err = dao.RemoveQuestionTest(db, params.Questionid, params.Testid)
					} else {
						err = addQuestionTimeToTest(q, params.Testid)
					}
					if err == nil {
						err = dao.AddQuestionTest(db, params.Questionid, params.Testid, *params.ValorFinal.ValorFinal)
						if err == nil {
							return test.NewAddQuestionToTestOK()
						}
					}
				} else {
					return test.NewAddQuestionToTestForbidden()
				}
			}
			log.Println("Error en users_handler AddQuestionToTest(): ", err)
			return test.NewAddQuestionToTestGone()
		}
	}
	return test.NewAddQuestionToTestInternalServerError()
}

func substractQuestionTimeTest(qid int64, tid int64) error {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var q *dao.Question
		q, err = dao.GetQuestion(db, qid)
		if err == nil && q != nil {
			var t *dao.Test
			t, err = dao.GetTest(db, tid)
			if err == nil && t != nil {
				var mt *models.Test
				mt, err = dao.ToModelTest(t)
				if err == nil && mt != nil {
					if *mt.MaxMinutes >= *q.EstimatedTime {
						nuevoTiempo := *mt.MaxMinutes - *q.EstimatedTime
						mt.MaxMinutes = &nuevoTiempo
						err = dao.PutTest(db, tid, mt)
						return err
					} else {
						return nil
					}
				}
			}
		}
	}
	return err
}

// RemoveQuestionTest DELETE /tests/{testid}/questions/{questionid}. Remove question from test
// Auth: TestAdmin or Admin
// Req: Test.editable
func RemoveQuestionTest(params test.RemoveQuestionFromTestParams, u *models.User) middleware.Responder {
	if testEditable(params.Testid) && (isAdmin(u) || isTestAdmin(u, params.Testid)) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.RemoveQuestionTest(db, params.Questionid, params.Testid)
			if err == nil {
				err = substractQuestionTimeTest(params.Questionid, params.Testid)
				if err == nil {
					return test.NewRemoveQuestionFromTestOK()
				}
			}
			log.Println("Error en users_handler AddQuestionToTest(): ", err)
			return test.NewRemoveQuestionFromTestGone()
		}
	}
	return test.NewRemoveQuestionFromTestForbidden()
}

// GetTagsFromTests GET /tests/{testid}/tags
// Auth: Teacher or Admin
func GetTagsFromTest(params test.GetTagsFromTestParams, u *models.User) middleware.Responder {
	if isTeacherOrAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var dts []*dao.Tag
			dts, err = dao.GetTestTags(db, params.Testid)
			if err == nil {
				ts := dao.ToModelTags(dts)
				return test.NewGetTagsFromTestOK().WithPayload(ts)
			}
		}
		return test.NewGetTagsFromTestInternalServerError()
	}
	return test.NewGetTagsFromTestForbidden()
}

// GetTagFromTest GET /tests/{testid}/tags/{tag}
// Auth: Teacher or Admin
func GetTagFromTest(params test.GetTagFromTestParams, u *models.User) middleware.Responder {
	if isTeacherOrAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var dt *dao.Tag
			dt, err = dao.GetTestTag(db, params.Testid, params.Tag)
			if err == nil {
				if dt != nil {
					t := dao.ToModelTag(dt)
					return test.NewGetTagFromTestOK().WithPayload(t)
				}
				return test.NewGetTagFromTestGone()
			}
		}
		return test.NewGetTagFromTestInternalServerError()
	}
	return test.NewGetTagFromTestForbidden()
}

// AddTagToTest PUT /tests/{testid}/tags/{tag}
// Auth: TestAdmin or Admin
// Req: Test.editable
func AddTagToTest(params test.AddTagToTestParams, u *models.User) middleware.Responder {
	if testEditable(params.Testid) && (isAdmin(u) || isTestAdmin(u, params.Testid)) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.AddTestTag(db, params.Testid, params.Tag)
			if err == nil {
				return test.NewAddTagToTestOK()
			}
		}
		return test.NewAddTagToTestInternalServerError()
	}
	return test.NewAddTagToTestForbidden()
}

// RemoveTagFromTest DELETE /tests/{testid}/tags/{tag}
// Auth: TestAdmin or Admin
// Req: Test.editable
func RemoveTagFromTest(params test.RemoveTagFromTestParams, u *models.User) middleware.Responder {
	if testEditable(params.Testid) && (isAdmin(u) || isTestAdmin(u, params.Testid)) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.RemoveTestTag(db, params.Testid, params.Tag)
			if err == nil {
				return test.NewRemoveTagFromTestOK()
			}
		}
		return test.NewRemoveTagFromTestInternalServerError()
	}
	return test.NewRemoveTagFromTestForbidden()
}

// GET /users/{username}/favoriteEditTests
// Auth: Current User or Admin
// Req: Fav+available+editable (SQL)
func GetFavoriteEditTests(params user.GetFavoriteEditTestsParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var qs []*dao.Test
			qs, err = dao.GetFavoriteEditTests(db, params.Username, params.Tags, params.LikeTitle)
			if err == nil {
				var mqs []*models.Test
				mqs, err = dao.ToModelTests(qs)
				if err == nil {
					return user.NewGetFavoriteEditTestsOK().WithPayload(mqs)
				}
			}
		}
		return user.NewGetFavoriteEditTestsInternalServerError()
	}
	return user.NewGetFavoriteEditTestsForbidden()
}

// GET /users/{username}/favoriteTests
// Auth: Current User or Admin
// Req: Fav+available (SQL)
func GetFavoriteTests(params user.GetFavoriteTestsParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var qs []*dao.Test
			qs, err = dao.GetFavoriteTests(db, params.Username, params.Tags, params.LikeTitle)
			if err == nil {
				var mqs []*models.Test
				mqs, err = dao.ToModelTests(qs)
				if err == nil {
					return user.NewGetFavoriteTestsOK().WithPayload(mqs)
				}
			}
		}
		return user.NewGetFavoriteTestsInternalServerError()
	}
	return user.NewGetFavoriteTestsForbidden()
}

// GET /users/{username}/favoriteTests/{testid}
// Auth: Current User or Admin
// Req: Fav+available (SQL)
func GetFavoriteTest(params user.GetFavoriteTestParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var qs *dao.Test
			qs, err = dao.GetFavoriteTest(db, params.Username, params.Testid)
			if err == nil {
				if qs == nil {
					return user.NewGetFavoriteTestGone()
				}
				var mqs *models.Test
				mqs, err = dao.ToModelTest(qs)
				if err == nil {
					return user.NewGetFavoriteTestOK().WithPayload(mqs)
				}
			}
		}
		return user.NewGetFavoriteTestInternalServerError()
	}
	return user.NewGetFavoriteTestForbidden()
}

// PUT /users/{username}/favoriteTests/{testid}
// Auth: Current User or Admin
func AddFavoriteTest(params user.AddTestFavoriteParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.AddFavoriteTest(db, params.Username, params.Testid)
			if err == nil {
				return user.NewAddTestFavoriteOK()
			}
		}
		return user.NewAddTestFavoriteInternalServerError()
	}
	return user.NewAddTestFavoriteForbidden()
}

// DELETE /users/{username}/favoriteTests/{testid}
// Auth: Current User or Admin
func RemoveFavoriteTest(params user.RemoveTestFavoriteParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.RemoveFavoriteTest(db, params.Username, params.Testid)
			if err == nil {
				return user.NewRemoveTestFavoriteOK()
			}
		}
		return user.NewRemoveTestFavoriteInternalServerError()
	}
	return user.NewRemoveTestFavoriteForbidden()
}
