// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package handlers provides functions that handle http requests
package handlers

import (
	"log"
	"uva-devtest/models"
	"uva-devtest/persistence/dao"
	"uva-devtest/persistence/dbconnection"
	"uva-devtest/restapi/operations/answer"

	"github.com/go-openapi/runtime/middleware"
)

// GET /answers
// Req: Teacher or Admin
func GetAnswers(params answer.GetAnswersParams, u *models.User) middleware.Responder {
	if isTeacherOrAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var as []*dao.Answer
			as, err = dao.GetAnswers(db)
			if err == nil {
				var mas []*models.Answer
				mas, err = dao.ToModelAnswers(as)
				if err == nil {
					return answer.NewGetAnswersOK().WithPayload(mas)
				}
			}
		}
		log.Println("Error en answers_handler GetAnswers(): ", err)
		return answer.NewGetAnswersInternalServerError()
	}
	return answer.NewGetAnswersForbidden()
}

func isAnswerOwner(answerid int64, u *models.User) bool {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var a *dao.Answer
		a, err = dao.GetAnswer(db, answerid)
		if err == nil && a != nil {
			var du *dao.User
			du, err = dao.GetUserUsername(db, *u.Username)
			if err == nil && du != nil {
				return du.ID == a.Usuarioid
			}
		}
	}
	return false
}

// GET /answers/{answerid}
// Req: Teacher or Admin OR AnswerOwner
func GetAnswer(params answer.GetAnswerParams, u *models.User) middleware.Responder {
	if isTeacherOrAdmin(u) || isAnswerOwner(params.Answerid, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var as *dao.Answer
			as, err = dao.GetAnswer(db, params.Answerid)
			if err == nil {
				var mas *models.Answer
				mas, err = dao.ToModelAnswer(as)
				if err == nil {
					return answer.NewGetAnswerOK().WithPayload(mas)
				}
			}
		}
		log.Println("Error en answers_handler GetAnswer(): ", err)
		return answer.NewGetAnswerInternalServerError()
	}
	return answer.NewGetAnswerForbidden()
}

// PUT /answers/{answerid}
// Req: Admin OR AnswerOwner
func FinishAnswer(params answer.FinishAnswerParams, u *models.User) middleware.Responder {
	if isAdmin(u) || isAnswerOwner(params.Answerid, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.FinishAnswer(db, params.Answerid)
			if err == nil {
				return answer.NewFinishAnswerOK()
			}
		}
		log.Println("Error en answers_handler FinishAnswer(): ", err)
		return answer.NewFinishAnswerInternalServerError()
	}
	return answer.NewFinishAnswerForbidden()
}

// GET /answers/{answerid}/qanswers
// Req: Teacher or Admin OR AnswerOwner
func GetQuestionAnswers(params answer.GetQuestionAnswersFromAnswerParams, u *models.User) middleware.Responder {
	if isTeacherOrAdmin(u) || isAnswerOwner(params.Answerid, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var qas []*dao.QuestionAnswer
			qas, err = dao.GetQuestionAnswersFromAnswer(db, params.Answerid)
			if err == nil {
				mqas := dao.ToModelQuestionAnswers(qas)
				return answer.NewGetQuestionAnswersFromAnswerOK().WithPayload(mqas)
			}
		}
		log.Println("Error en answers_handler GetQuestionAnswers(): ", err)
		return answer.NewGetQuestionAnswersFromAnswerInternalServerError()
	}
	return answer.NewGetQuestionAnswersFromAnswerForbidden()
}

// POST /answers/{answerid}/qanswers
// Req: Admin OR AnswerOwner
func PostQuestionAnswer(params answer.PostQuestionAnswerParams, u *models.User) middleware.Responder {
	if isAdmin(u) || isAnswerOwner(params.Answerid, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var pqa *models.QuestionAnswer
			pqa, err = dao.PostQuestionAnswer(db, params.Answerid, params.QuestionAnswer)
			return answer.NewPostQuestionAnswerCreated().WithPayload(pqa)
		}
		log.Println("Error en answers_handler PostQuestionAnswer(): ", err)
		return answer.NewPostQuestionAnswerInternalServerError()
	}
	return answer.NewPostQuestionAnswerForbidden()
}

// GET /answers/{answerid}/qanswers/{questionid}
// Req: Teacher or Admin OR AnswerOwner
func GetQuestionAnswer(params answer.GetQuestionAnswerFromAnswerParams, u *models.User) middleware.Responder {
	if isTeacherOrAdmin(u) || isAnswerOwner(params.Answerid, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var qas *dao.QuestionAnswer
			qas, err = dao.GetQuestionAnswerFromAnswer(db, params.Answerid, params.Questionid)
			if err == nil && qas != nil {
				mqas := dao.ToModelQuestionAnswer(qas)
				return answer.NewGetQuestionAnswerFromAnswerOK().WithPayload(mqas)
			}
		}
		log.Println("Error en answers_handler GetQuestionAnswer(): ", err)
		return answer.NewGetQuestionAnswerFromAnswerInternalServerError()
	}
	return answer.NewGetQuestionAnswerFromAnswerForbidden()
}

// PUT /answers/{answerid}/qanswers/{questionid}
// Req: Admin OR AnswerOwner
func PutQuestionAnswer(params answer.PutQuestionAnswerFromAnswerParams, u *models.User) middleware.Responder {
	if isAdmin(u) || isAnswerOwner(params.Answerid, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.PutQuestionAnswer(db, params.Answerid, params.Questionid, params.QuestionAnswer)
			return answer.NewPutQuestionAnswerFromAnswerOK()
		}
		log.Println("Error en answers_handler PutQuestionAnswer(): ", err)
		return answer.NewPutQuestionAnswerFromAnswerInternalServerError()
	}
	return answer.NewPutQuestionAnswerFromAnswerForbidden()
}

func isAnswerTestAdmin(u *models.User, answerid int64) bool {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var a *dao.Answer
		a, err = dao.GetAnswer(db, answerid)
		if err == nil && a != nil {
			return isTestAdmin(u, a.Testid)
		}
	}
	return false
}

// PUT /answers/{answerid}/qanswers/{questionid}/review
// Req: TestAdmin or Admin
func PutReview(params answer.PutReviewParams, u *models.User) middleware.Responder {
	if isAnswerTestAdmin(u, params.Answerid) || isAnswerOwner(params.Answerid, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.PutReview(db, params.Answerid, params.Questionid, params.Review)
			if err == nil {
				return answer.NewPutReviewOK()
			}
		}
		log.Println("Error en answers_handler PutReview(): ", err)
		return answer.NewPutReviewInternalServerError()

	}
	return answer.NewPutReviewForbidden()
}
