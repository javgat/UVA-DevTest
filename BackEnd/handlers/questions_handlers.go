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
		for _, itemCopy := range ts {
			q, err = dao.GetQuestionFromTeam(db, *itemCopy.Teamname, questionid)
			if q != nil && err == nil {
				return true
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
// Auth: All
func GetQuestionTags(params question.GetTagsFromQuestionParams, u *models.User) middleware.Responder {
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

// GET /questions/{questionid}/tags/{tag}
// Auth: All
func GetQuestionTag(params question.GetTagFromQuestionParams, u *models.User) middleware.Responder {
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
	if isQuestionEditable(params.Questionid) && isTeamSoloProfesores(params.Teamname) && (isAdmin(u) || isQuestionAdmin(u, params.Questionid)) {
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
