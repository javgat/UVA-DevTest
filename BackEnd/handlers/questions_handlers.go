// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package handlers provides functions that handle http requests
package handlers

import (
	"log"
	"uva-devtest/models"
	"uva-devtest/persistence/dao"
	"uva-devtest/persistence/dbconnection"
	"uva-devtest/restapi/operations/question"

	"github.com/go-openapi/runtime/middleware"
)

// GetEditQuestions GET /editQuestions. Returns all non-published questions.
// Auth: Teacher or Admin
func GetEditQuestions(params question.GetEditQuestionsParams, u *models.User) middleware.Responder {
	if isTeacherOrAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var qs []*dao.Question
			qs, err = dao.GetEditQuestions(db)
			if err == nil {
				var mqs []*models.Question
				mqs, err = dao.ToModelQuestions(qs)
				if err == nil {
					return question.NewGetEditQuestionsOK().WithPayload(mqs)
				}
			}
		}
		log.Println("Error en users_handler GetEditQuestions(): ", err)
		return question.NewGetEditQuestionsInternalServerError()
	}
	return question.NewGetEditQuestionsForbidden()
}

// GetQuestions GET /questions. Returns all questions.
// Auth: Teacher or Admin
func GetQuestions(params question.GetQuestionsParams, u *models.User) middleware.Responder {
	if isTeacherOrAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var qs []*dao.Question
			qs, err = dao.GetQuestions(db)
			if err == nil {
				var mqs []*models.Question
				mqs, err = dao.ToModelQuestions(qs)
				if err == nil {
					return question.NewGetQuestionsOK().WithPayload(mqs)
				}
			}
		}
		log.Println("Error en users_handler GetQuestions(): ", err)
		return question.NewGetQuestionsInternalServerError()
	}
	return question.NewGetQuestionsForbidden()
}

// GetQuestion GET /questions/{questionid}. Returns a question.
// Auth: All
func GetQuestion(params question.GetQuestionParams, u *models.User) middleware.Responder {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var qs *dao.Question
		qs, err = dao.GetQuestion(db, params.Questionid)
		if err == nil && qs != nil {
			var mqs *models.Question
			mqs, err = dao.ToModelQuestion(qs)
			if err == nil {
				return question.NewGetQuestionOK().WithPayload(mqs)
			}
		}
	}
	log.Println("Error en users_handler GetQuestions(): ", err)
	return question.NewGetQuestionInternalServerError()
}

// Owner o miembro de team que admin question
func isQuestionAdmin(u *models.User, questionid int64) bool {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var q *dao.Question
		q, err = dao.GetQuestionOfUser(db, *u.Username, questionid)
		if q != nil && err == nil {
			return true
		}
		var ts []*dao.Team
		ts, err = dao.GetTeamsUsername(db, *u.Username)
		if err == nil {
			for _, itemCopy := range ts {
				q, err = dao.GetQuestionFromTeam(db, *itemCopy.Teamname, questionid)
				if q != nil && err == nil {
					return true
				}
			}
		}
	}
	return false
}

func isQuestionEditable(questionid int64) bool {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var q *dao.Question
		q, err = dao.GetQuestion(db, questionid)
		if err == nil && q != nil {
			return *q.Editable
		}
	}
	return false
}

func isTeamSoloProfesores(teamname string) bool {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		t, err := dao.GetTeam(db, teamname)
		if err == nil && t != nil {
			return *t.SoloProfesores
		}
	}
	return false
}

// PutQuestion PUT /questions/{questionid}. Modifies a question
// Auth: QuestionOwner or Admin
// Req: Question.Editable
func PutQuestion(params question.PutQuestionParams, u *models.User) middleware.Responder {
	if isQuestionEditable(params.Questionid) && (isAdmin(u) || isQuestionAdmin(u, params.Questionid)) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.PutQuestion(db, params.Questionid, params.Question)
			if err == nil {
				return question.NewPutQuestionOK()
			}
		}
	}
	return question.NewPutQuestionForbidden()
}

// DeleteQuestion DELETE /questions/{questionid}. Deletes a question
// Auth: QuestionOwner or Admin
// Req: Question.Editable
func DeleteQuestion(params question.DeleteQuestionParams, u *models.User) middleware.Responder {
	if isQuestionEditable(params.Questionid) && (isAdmin(u) || isQuestionAdmin(u, params.Questionid)) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.DeleteQuestion(db, params.Questionid)
			if err == nil {
				return question.NewDeleteQuestionOK()
			}
		}
	}
	return question.NewDeleteQuestionForbidden()
}

// GET /questions/{questionid}/tags
// Auth: Teacher or admin
func GetQuestionTags(params question.GetTagsFromQuestionParams, u *models.User) middleware.Responder {
	if isTeacherOrAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var ts []*dao.Tag
			ts, err = dao.GetQuestionTags(db, params.Questionid)
			var mts []*models.Tag
			if err == nil {
				mts = dao.ToModelTags(ts)
				if err == nil {
					return question.NewGetTagsFromQuestionOK().WithPayload(mts)
				}
			}
		}
		log.Println("Error en users_handler GetQuestionTags(): ", err)
		return question.NewGetTagsFromQuestionInternalServerError()
	}
	return question.NewGetTagsFromQuestionForbidden()
}

// GET /questions/{questionid}/tags/{tag}
// Auth: Teacher or admin
func GetQuestionTag(params question.GetTagFromQuestionParams, u *models.User) middleware.Responder {
	if isTeacherOrAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var t *dao.Tag
			t, err = dao.GetQuestionTag(db, params.Questionid, params.Tag)
			var mt *models.Tag
			if err == nil && t != nil {
				mt = dao.ToModelTag(t)
				if err == nil {
					return question.NewGetTagFromQuestionOK().WithPayload(mt)
				}
			}
			return question.NewGetTagFromQuestionGone()
		}
		log.Println("Error en users_handler GetQuestionTag(): ", err)
		return question.NewGetTagFromQuestionInternalServerError()
	}
	return question.NewGetTagFromQuestionForbidden()
}

// PUT /questions/{questionid}/tags/{tag}
// Auth: QuestionOwner or Admin
// Req: Question.Editable
func AddQuestionTag(params question.AddTagToQuestionParams, u *models.User) middleware.Responder {
	if isQuestionEditable(params.Questionid) && (isAdmin(u) || isQuestionAdmin(u, params.Questionid)) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.AddQuestionTag(db, params.Questionid, params.Tag)
			if err == nil {
				if err == nil {
					return question.NewAddTagToQuestionOK()
				}
			}
		}
		log.Println("Error en users_handler AddQuestionTag(): ", err)
		return question.NewAddTagToQuestionGone()
	}
	return question.NewAddTagToQuestionForbidden()
}

// DELETE /questions/{questionid}/tags/{tag}
// Auth: QuestionOwner or Admin
// Req: Question.Editable
func RemoveQuestionTag(params question.RemoveTagFromQuestionParams, u *models.User) middleware.Responder {
	if isQuestionEditable(params.Questionid) && (isAdmin(u) || isQuestionAdmin(u, params.Questionid)) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.RemoveQuestionTag(db, params.Questionid, params.Tag)
			if err == nil {
				if err == nil {
					return question.NewRemoveTagFromQuestionOK()
				}
			}
		}
		log.Println("Error en users_handler RemoveQuestionTag(): ", err)
		return question.NewRemoveTagFromQuestionGone()
	}
	return question.NewRemoveTagFromQuestionForbidden()
}

// PUT /questions/{questionid}/teams/{teamid}
// Auth: QuestionOwner or Admin
// Req: Question.Editable
func AddQuestionTeam(params question.AddTeamToQuestionParams, u *models.User) middleware.Responder {
	if isQuestionEditable(params.Questionid) && (isAdmin(u) || isQuestionAdmin(u, params.Questionid)) {
		if isTeamSoloProfesores(params.Teamname) {
			db, err := dbconnection.ConnectDb()
			if err == nil {
				err = dao.AddQuestionTeam(db, params.Questionid, params.Teamname)
				if err == nil {
					if err == nil {
						return question.NewAddTeamToQuestionOK()
					}
				}
				log.Println("Error en users_handler AddQuestionTeam(): ", err)
				return question.NewAddTagToQuestionGone()
			}
		}
		s := "El equipo a añadir tiene que ser de profesores"
		return question.NewAddTeamToQuestionBadRequest().WithPayload(&models.Error{Message: &s})
	}
	return question.NewAddTeamToQuestionForbidden()
}

// DELETE /questions/{questionid}/teams/{teamid}
// Auth: QuestionOwner or Admin
// Req: Question.Editable
func RemoveQuestionTeam(params question.RemoveTeamToQuestionParams, u *models.User) middleware.Responder {
	if isQuestionEditable(params.Questionid) && (isAdmin(u) || isQuestionAdmin(u, params.Questionid)) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.RemoveQuestionTeam(db, params.Questionid, params.Teamname)
			if err == nil {
				if err == nil {
					return question.NewRemoveTeamToQuestionOK()
				}
			}
			log.Println("Error en users_handler RemoveQuestionTeam(): ", err)
			return question.NewRemoveTeamToQuestionGone()
		}
	}
	return question.NewRemoveTeamToQuestionForbidden()
}

// GET /questions/{questionid}/teams
// Auth: All
func GetTeamsFromQuestion(params question.GetTeamsFromQuestionParams, u *models.User) middleware.Responder {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var ts []*dao.Team
		ts, err = dao.GetTeamsQuestion(db, params.Questionid)
		if err == nil {
			mts := dao.ToModelsTeams(ts)
			if err == nil {
				return question.NewGetTeamsFromQuestionOK().WithPayload(mts)
			}
		}
	}
	log.Println("Error en users_handler GetTeamsFromQuestion(): ", err)
	return question.NewGetTeamsFromQuestionInternalServerError()
}

// GET /questions/{questionid}/options
// Auth: Teacher or admin
func GetOptions(params question.GetOptionsFromQuestionParams, u *models.User) middleware.Responder {
	if isTeacherOrAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var os []*dao.Option
			os, err = dao.GetOptionsQuestion(db, params.Questionid)
			if err == nil {
				mos := dao.ToModelOptions(os)
				if err == nil {
					return question.NewGetOptionsFromQuestionOK().WithPayload(mos)
				}
			}
		}
		log.Println("Error en users_handler GetOptions(): ", err)
		return question.NewGetOptionsFromQuestionInternalServerError()
	}
	return question.NewGetOptionsFromQuestionForbidden()
}

// POST /questions/{questionid}/options
// Auth: QuestionAdmin or Admin
func PostOption(params question.PostOptionParams, u *models.User) middleware.Responder {
	if isQuestionAdmin(u, params.Questionid) || isAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var os *dao.Option
			os, err = dao.PostOption(db, params.Questionid, params.Option)
			if err == nil && os != nil {
				mos := dao.ToModelOption(os)
				if err == nil {
					return question.NewPostOptionCreated().WithPayload(mos)
				}
			}
		}
		log.Print("Error en crear opcion: ", err)
		return question.NewPostOptionInternalServerError()
	}
	return question.NewPostOptionForbidden()
}

// GET /questions/{questionid}/options/{optionindex}
// Auth: Teacher or admin
func GetOption(params question.GetOptionFromQuestionParams, u *models.User) middleware.Responder {
	if isTeacherOrAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var os *dao.Option
			os, err = dao.GetOptionQuestion(db, params.Questionid, params.Optionindex)
			if err == nil {
				if os == nil {
					return question.NewGetOptionFromQuestionGone()
				}
				mos := dao.ToModelOption(os)
				if err == nil {
					return question.NewGetOptionFromQuestionOK().WithPayload(mos)
				}
			}
		}
		log.Println("Error en users_handler GetOption(): ", err)
		return question.NewGetOptionFromQuestionInternalServerError()
	}
	return question.NewGetOptionFromQuestionForbidden()
}

// PUT /questions/{questionid}/options/{optionindex}
// Auth: QuestionAdmin or Admin
// Req: Question.Editable
func PutOption(params question.PutOptionParams, u *models.User) middleware.Responder {
	if isQuestionEditable(params.Questionid) && (isQuestionAdmin(u, params.Questionid) || isAdmin(u)) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.PutOption(db, params.Questionid, params.Optionindex, params.Option)
			if err == nil {
				return question.NewPutOptionOK()
			}
		}
		log.Println("Error en users_handler PutOption(): ", err)
		return question.NewPutOptionInternalServerError()
	}
	return question.NewPutOptionForbidden()
}

// DELETE /questions/{questionid}/options/{optionindex}
// Auth: QuestionAdmin or Admin
// Req: Question.Editable
func DeleteOption(params question.DeleteOptionParams, u *models.User) middleware.Responder {
	if isQuestionEditable(params.Questionid) && (isQuestionAdmin(u, params.Questionid) || isAdmin(u)) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.DeleteOption(db, params.Questionid, params.Optionindex)
			if err == nil {
				return question.NewDeleteOptionOK()
			}
		}
		return question.NewDeleteOptionInternalServerError()
	}
	return question.NewDeleteOptionForbidden()
}
