// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package handlers provides functions that handle http Authuests
package handlers

import (
	"database/sql"
	"errors"
	"log"
	"strings"
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

func autoCorrigeString(daq *dao.QuestionAnswer, dq *dao.Question) error {
	var puntuacion int64 = -*dq.Penalizacion
	if strings.EqualFold(daq.Respuesta, dq.Solucion) {
		puntuacion = 100
	}
	review := &models.Review{
		Puntuacion: &puntuacion,
	}
	err := updateReview(*daq.IDRespuesta, dq.ID, review)
	return err
}

func autoCorrigeOpcionesEleUni(daq *dao.QuestionAnswer, dq *dao.Question) error {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var puntuacion int64 = -*dq.Penalizacion
		if len(daq.IndicesOpciones) > 0 {
			ind := daq.IndicesOpciones[0]
			var opt *dao.Option
			opt, err = dao.GetOptionQuestion(db, dq.ID, ind)
			if err != nil {
				return err
			}
			if opt == nil {
				return errors.New("opcion seleccionada no existe")
			}
			if *opt.Correcta {
				puntuacion = 100
			}
		}
		review := &models.Review{
			Puntuacion: &puntuacion,
		}
		err = updateReview(*daq.IDRespuesta, dq.ID, review)
	}
	return err
}

func int64sliceContains(slice []int64, value int64) bool {
	for _, val := range slice {
		if val == value {
			return true
		}
	}
	return false
}

func autoCorrigeOpcionesEleMulti(daq *dao.QuestionAnswer, dq *dao.Question) error {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var puntuacion int64 = 100
		var opts []*dao.Option
		opts, err = dao.GetOptionsQuestion(db, dq.ID)
		if err == nil {
			for _, opt := range opts {
				if (*opt.Correcta && !int64sliceContains(daq.IndicesOpciones, opt.Indice)) ||
					(!*opt.Correcta && int64sliceContains(daq.IndicesOpciones, opt.Indice)) {
					puntuacion = -*dq.Penalizacion
					break
				}
			}
			review := &models.Review{
				Puntuacion: &puntuacion,
			}
			err = updateReview(*daq.IDRespuesta, dq.ID, review)
		}
	}
	return err
}

func autoCorrigeRespuestaPregunta(aid int64, dq *dao.Question) error {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var daq *dao.QuestionAnswer
		daq, err = dao.GetQuestionAnswerFromAnswer(db, aid, dq.ID)
		if err == nil {
			if daq == nil {
				return nil // Es valido que no haya respondido a una pregunta
			}
			switch *dq.TipoPregunta {
			case models.QuestionTipoPreguntaString:
				err = autoCorrigeString(daq, dq)
			case models.QuestionTipoPreguntaOpciones:
				if err == nil {
					if dq.EleccionUnica {
						err = autoCorrigeOpcionesEleUni(daq, dq)
					} else {
						err = autoCorrigeOpcionesEleMulti(daq, dq)
					}
				}
			default:
				return errors.New("tipo de pregunta extraño")
			}
		}
	}
	return err
}

func marcarRespuestaComoCorregida(aid int64) error {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var da *dao.Answer
		var dt *dao.Test
		da, err = dao.GetAnswer(db, aid)
		if err == nil {
			if da == nil {
				return errors.New("respuesta no encontrada con el id")
			}
			dt, err = dao.GetTest(db, da.Testid)
			if err == nil {
				if dt == nil {
					return errors.New("test no encontrado con el id")
				}
				err = dao.SetAnswerCorrected(db, aid)
				if err == nil {
					if strings.EqualFold(*dt.Visibilidad, models.TestVisibilidadAlCorregir) {
						err = dao.SetAnswerVisible(db, aid)
					}
				}
			}
		}
	}
	return err
}

func marcarRespuestaComoNoCorregida(aid int64) error {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var da *dao.Answer
		var dt *dao.Test
		da, err = dao.GetAnswer(db, aid)
		if err == nil {
			if da == nil {
				return errors.New("respuesta no encontrada con el id")
			}
			dt, err = dao.GetTest(db, da.Testid)
			if err == nil {
				if dt == nil {
					return errors.New("test no encontrado con el id")
				}
				err = dao.SetAnswerNotCorrected(db, aid)
				if err == nil {
					if strings.EqualFold(*dt.Visibilidad, models.TestVisibilidadAlCorregir) {
						err = dao.SetAnswerNotVisible(db, aid)
					}
				}
			}
		}
	}
	return err
}

func autoCorregirLoPosible(aid int64) error {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var da *dao.Answer
		da, err = dao.GetAnswer(db, aid)
		if err == nil {
			if da == nil {
				return errors.New("respuesta no existe con ese id")
			}
			var dqs []*dao.Question
			dqs, err = dao.GetQuestionsFromTest(db, da.Testid)
			if err == nil {
				for _, dq := range dqs {
					if *dq.AutoCorrect {
						err = autoCorrigeRespuestaPregunta(aid, dq)
						if err != nil {
							return err
						}
					}
				}
			}
			var dt *dao.Test
			dt, err = dao.GetTest(db, da.Testid)
			if err == nil {
				if dt == nil {
					return errors.New("test no existe con ese id")
				}
				if *dt.AutoCorrect {
					err = marcarRespuestaComoCorregida(aid)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return err
}

// PUT /answers/{answerid}
// Auth: Admin OR AnswerOwner
func FinishAnswer(params answer.FinishAnswerParams, u *models.User) middleware.Responder {
	if isAdmin(u) || isAnswerOwner(params.Answerid, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.FinishAnswer(db, params.Answerid)
			if err == nil {
				err = autoCorregirLoPosible(params.Answerid)
				if err == nil {
					return answer.NewFinishAnswerOK()
				}
			}
		}
		log.Println("Error en answers_handler FinishAnswer(): ", err)
		return answer.NewFinishAnswerInternalServerError()
	}
	return answer.NewFinishAnswerForbidden()
}

// PUT /answers/{answerid}/corrected
// Auth: Admin OR TestAdmin
func SetAnswerCorrected(params answer.SetAnswerCorrectedParams, u *models.User) middleware.Responder {
	if isAdmin(u) || isAnswerTestAdmin(u, params.Answerid) {
		err := marcarRespuestaComoCorregida(params.Answerid)
		if err == nil {
			return answer.NewSetAnswerCorrectedOK()
		}
		log.Println("Error en answers_handler SetAnswerCorrected(): ", err)
		return answer.NewSetAnswerCorrectedInternalServerError()
	}
	return answer.NewSetAnswerCorrectedForbidden()
}

// DELETE /answers/{answerid}/corrected
// Auth: Admin OR TestAdmin
func SetAnswerNotCorrected(params answer.SetAnswerNotCorrectedParams, u *models.User) middleware.Responder {
	if isAdmin(u) || isAnswerTestAdmin(u, params.Answerid) {
		err := marcarRespuestaComoNoCorregida(params.Answerid)
		if err == nil {
			return answer.NewSetAnswerNotCorrectedOK()
		}
		log.Println("Error en answers_handler SetAnswerNotCorrected(): ", err)
		return answer.NewSetAnswerNotCorrectedInternalServerError()
	}
	return answer.NewSetAnswerNotCorrectedForbidden()
}

// PUT /answers/{answerid}/visible
// Auth: Admin OR TestAdmin
func SetAnswerVisible(params answer.SetAnswerVisibleParams, u *models.User) middleware.Responder {
	if isAdmin(u) || isAnswerTestAdmin(u, params.Answerid) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.SetAnswerVisible(db, params.Answerid)
			if err == nil {
				return answer.NewSetAnswerVisibleOK()
			}
		}
		log.Println("Error en answers_handler SetAnswerVisible(): ", err)
		return answer.NewSetAnswerVisibleInternalServerError()
	}
	return answer.NewSetAnswerVisibleForbidden()
}

// DELETE /answers/{answerid}/visible
// Auth: Admin OR TestAdmin
func SetAnswerNotVisible(params answer.SetAnswerNotVisibleParams, u *models.User) middleware.Responder {
	if isAdmin(u) || isAnswerTestAdmin(u, params.Answerid) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.SetAnswerNotVisible(db, params.Answerid)
			if err == nil {
				return answer.NewSetAnswerNotVisibleOK()
			}
		}
		log.Println("Error en answers_handler SetAnswerNotVisible(): ", err)
		return answer.NewSetAnswerNotVisibleInternalServerError()
	}
	return answer.NewSetAnswerNotVisibleForbidden()
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
			return *a.Entregado
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

func addAnswerPuntuacion(aid int64, qid int64, puntuacion int64) error {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var a *dao.Answer
		var q *dao.Question
		a, err = dao.GetAnswer(db, aid)
		if err == nil {
			if a == nil {
				return errors.New("no se encontro el recurso")
			}
			q, err = dao.GetQuestionFromTest(db, a.Testid, qid)
			if err == nil {
				if q == nil {
					return errors.New("no se encontro el recurso")
				}
				punt := a.Puntuacion + float64(*q.ValorFinal*puntuacion)/float64(100)
				err = dao.PutAnswerPuntuacion(db, aid, punt)
				if err == nil {
					return nil
				}
			}
		}
	}
	return err
}

func substractAnswerPuntuacion(aid int64, qid int64, puntuacion int64) error {
	return addAnswerPuntuacion(aid, qid, -puntuacion)
}

func updateReview(aid int64, qid int64, review *models.Review) error {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var qa *dao.QuestionAnswer
		qa, err = dao.GetQuestionAnswerFromAnswer(db, aid, qid)
		if qa != nil && err == nil {
			err = substractAnswerPuntuacion(aid, qid, *qa.Puntuacion)
		}
		if err == nil {
			log.Println(*review.Puntuacion)
			err = addAnswerPuntuacion(aid, qid, *review.Puntuacion)
			if err == nil {
				err = dao.PutReview(db, aid, qid, review)
			}
		}
	}
	return err
}

// PUT /answers/{answerid}/qanswers/{questionid}/review
// Auth: TestAdmin or Admin
func PutReview(params answer.PutReviewParams, u *models.User) middleware.Responder {
	if isAnswerTestAdmin(u, params.Answerid) || isAnswerOwner(params.Answerid, u) {
		err := updateReview(params.Answerid, params.Questionid, params.Review)
		if err == nil {
			return answer.NewPutReviewOK()
		}
		log.Println("Error en answers_handler PutReview(): ", err)
		return answer.NewPutReviewInternalServerError()

	}
	return answer.NewPutReviewForbidden()
}

// DELETE /answers/{answerid}/qanswers/{questionid}/review
// Auth: TestAdmin or Admin
func DeleteReview(params answer.DeleteReviewParams, u *models.User) middleware.Responder {
	if isAnswerTestAdmin(u, params.Answerid) || isAnswerOwner(params.Answerid, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var qa *dao.QuestionAnswer
			qa, err = dao.GetQuestionAnswerFromAnswer(db, params.Answerid, params.Questionid)
			if qa != nil && err == nil {
				err = substractAnswerPuntuacion(params.Answerid, params.Questionid, *qa.Puntuacion)
			}
			if err == nil {
				err = dao.DeleteReview(db, params.Answerid, params.Questionid)
				if err == nil {
					return answer.NewDeleteReviewOK()
				}
			}
		}
		log.Println("Error en answers_handler DeleteReview(): ", err)
		return answer.NewDeleteReviewInternalServerError()

	}
	return answer.NewDeleteReviewForbidden()
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
