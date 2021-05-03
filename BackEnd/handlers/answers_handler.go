// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package handlers provides functions that handle http Authuests
package handlers

import (
	"database/sql"
	"log"
	"uva-devtest/models"
	"uva-devtest/persistence/dao"
	"uva-devtest/persistence/dbconnection"
	"uva-devtest/restapi/operations/answer"

	"github.com/go-openapi/runtime/middleware"
)

// GET /answers
// Auth: Teacher or Admin
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
// Auth: Teacher or Admin OR AnswerOwner
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
// Auth: Admin OR AnswerOwner

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
// Auth: Teacher or Admin OR AnswerOwner
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

func isAnswerFinished(answerid int64) bool {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		a, err := dao.GetAnswer(db, answerid)
		if err == nil && a != nil {
			return *a.Finished
		}
	}
	return false
}

// Comprueba que una pregunta, en caso de ser de opciones y eleccion unica, sea valida
func isAnswerOpcionesUnicaValida(qa *models.QuestionAnswer) bool {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var q *dao.Question
		q, err = dao.GetQuestion(db, *qa.IDPregunta)
		if err == nil {
			if *q.TipoPregunta != models.QuestionTipoPreguntaOpciones || !q.EleccionUnica {
				return true
			}
			return (len(qa.IndicesOpciones) < 2)
		}
	}
	return false
}

// isQAnswerValida comprueba que la respuesta a la pregunta es valida para el tipo de pregunta
func isQAnswerValida(qa *models.QuestionAnswer) bool {
	return isAnswerOpcionesUnicaValida(qa)
}

// POST /answers/{answerid}/qanswers
// Auth: Admin OR AnswerOwner
// Req: Question no finished. If Question Tipo Opciones & eleccionUnica -> Solo una opcion marcada
func PostQuestionAnswer(params answer.PostQuestionAnswerParams, u *models.User) middleware.Responder {
	if !isAnswerFinished(params.Answerid) && (isAdmin(u) || isAnswerOwner(params.Answerid, u) &&
		isQAnswerValida(params.QuestionAnswer)) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var pqa *models.QuestionAnswer
			pqa, err = dao.PostQuestionAnswer(db, params.Answerid, params.QuestionAnswer)
			if err == nil && pqa != nil {
				return answer.NewPostQuestionAnswerCreated().WithPayload(pqa)
			}
		}
		log.Println("Error en answers_handler PostQuestionAnswer(): ", err)
		return answer.NewPostQuestionAnswerInternalServerError()
	}
	return answer.NewPostQuestionAnswerForbidden()
}

// GET /answers/{answerid}/qanswers/{questionid}
// Auth: Teacher or Admin OR AnswerOwner
func GetQuestionAnswer(params answer.GetQuestionAnswerFromAnswerParams, u *models.User) middleware.Responder {
	if isTeacherOrAdmin(u) || isAnswerOwner(params.Answerid, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var qas *dao.QuestionAnswer
			qas, err = dao.GetQuestionAnswerFromAnswer(db, params.Answerid, params.Questionid)
			if err == nil {
				if qas == nil {
					return answer.NewGetQuestionAnswerFromAnswerGone()
				}
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
// Auth: Admin OR AnswerOwner
// Req: Question no finished. If Question Tipo Opciones & eleccionUnica -> Solo una opcion marcada
func PutQuestionAnswer(params answer.PutQuestionAnswerFromAnswerParams, u *models.User) middleware.Responder {
	if !isAnswerFinished(params.Answerid) && (isAdmin(u) || isAnswerOwner(params.Answerid, u) && isQAnswerValida(params.QuestionAnswer)) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.PutQuestionAnswer(db, params.Answerid, params.Questionid, params.QuestionAnswer)
			if err == nil {
				return answer.NewPutQuestionAnswerFromAnswerOK()
			}
		}
		log.Println("Error en answers_handler PutQuestionAnswer(): ", err)
		return answer.NewPutQuestionAnswerFromAnswerInternalServerError()
	}
	return answer.NewPutQuestionAnswerFromAnswerForbidden()
}

// DELETE /answers/{answerid}/qanswers/{questionid}
// Auth: Admin OR AnswerOwner
// Req: Question no finished
func DeleteQuestionAnswer(params answer.DeleteQuestionAnswerFromAnswerParams, u *models.User) middleware.Responder {
	if !isAnswerFinished(params.Answerid) && (isAdmin(u) || isAnswerOwner(params.Answerid, u)) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.DeleteQuestionAnswer(db, params.Answerid, params.Questionid)
			if err == nil {
				return answer.NewDeleteQuestionAnswerFromAnswerOK()
			}
		}
		log.Println("Error en answers_handler DeleteQuestionAnswer(): ", err)
		return answer.NewDeleteQuestionAnswerFromAnswerInternalServerError()
	}
	return answer.NewDeleteQuestionAnswerFromAnswerForbidden()
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
// Auth: TestAdmin or Admin
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

func fillQuestionsIsRespondida(db *sql.DB, mqs []*models.Question, answerid int64) error {
	var qan *dao.QuestionAnswer
	var err error
	for _, mq := range mqs {
		qan, err = dao.GetQuestionAnswerFromAnswer(db, answerid, mq.ID)
		if err != nil {
			return err
		}
		mq.IsRespondida = qan != nil
	}
	return nil
}

// GET /answers/{answerid}/questions
// Auth: Admin or User with testStarted or TestAdmin
func GetQuestionsFromAnswer(params answer.GetQuestionsFromAnswerParams, u *models.User) middleware.Responder {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var ans *dao.Answer
		ans, err = dao.GetAnswer(db, params.Answerid)
		if err == nil {
			if ans == nil {
				return answer.NewGetQuestionsFromAnswerGone()
			}
			if isAdmin(u) || isTestOpenByUserAuth(u, ans.Testid) || isTestAdmin(u, ans.Testid) {
				var qs []*dao.Question
				qs, err = dao.GetQuestionsFromTest(db, ans.Testid)
				if err == nil {
					var mqs []*models.Question
					mqs, err = dao.ToModelQuestions(qs)
					if err == nil {
						err = fillQuestionsIsRespondida(db, mqs, params.Answerid)
						if err == nil {
							return answer.NewGetQuestionsFromAnswerOK().WithPayload(mqs)
						}
					}
				}
				log.Println("Error en GetQuestionsFromAnswer() ", err)
				return answer.NewGetQuestionsFromAnswerInternalServerError()
			}
			return answer.NewGetQuestionsFromAnswerForbidden()
		}
	}
	log.Println("Error en GetQuestionsFromAnswer() ", err)
	return answer.NewGetQuestionsFromAnswerInternalServerError()
}

// GET /answers/{answerid}/questions/{questionid}/qanswers
// Auth: Admin or User with testStarted or TestAdmin
func GetQAnswerFromAnswerAndQuestion(params answer.GetQuestionAnswersFromAnswerAndQuestionParams, u *models.User) middleware.Responder {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var ans *dao.Answer
		ans, err = dao.GetAnswer(db, params.Answerid)
		if err == nil {
			if ans == nil {
				return answer.NewGetQuestionAnswersFromAnswerAndQuestionGone()
			}
			if isAdmin(u) || isTestOpenByUserAuth(u, ans.Testid) || isTestAdmin(u, ans.Testid) {
				var qan *dao.QuestionAnswer
				qan, err = dao.GetQuestionAnswerFromAnswer(db, params.Answerid, params.Questionid)
				if err == nil {
					mqa := dao.ToModelQuestionAnswer(qan)
					if err == nil {
						return answer.NewGetQuestionAnswersFromAnswerAndQuestionOK().WithPayload([]*models.QuestionAnswer{mqa})
					}
				}
				log.Println("Error en GetQuestionAnswersFromAnswerAndQuestion() ", err)
				return answer.NewGetQuestionAnswersFromAnswerAndQuestionInternalServerError()
			}
			return answer.NewGetQuestionAnswersFromAnswerAndQuestionForbidden()
		}
	}
	log.Println("Error en GetQuestionAnswersFromAnswerAndQuestion() ", err)
	return answer.NewGetQuestionAnswersFromAnswerAndQuestionInternalServerError()
}
