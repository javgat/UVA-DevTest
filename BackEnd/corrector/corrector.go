// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package corrector executes Pruebas for the Code QuestionAnswers
package corrector

import (
	"errors"
	"log"
	"uva-devtest/models"
	"uva-devtest/persistence/dao"
	"uva-devtest/persistence/dbconnection"
)

func ExecutePrePruebas(answerid int64, questionid int64) {
	// STUB ACTION: errorCompilacion
	db, err := dbconnection.ConnectDb()
	if err == nil {
		err = dao.SetQuestionAnswerErrorCompilacion(db, answerid, questionid)
	}

	if err != nil {
		log.Println("Error en corrector en ExecutePrePruebas(): ", err)
	}
}

func ExecuteFullPruebas(answerid int64, questionid int64) {
	// STUB ACTION: errorCompilacion
	db, err := dbconnection.ConnectDb()
	if err == nil {
		err = dao.SetQuestionAnswerErrorCompilacion(db, answerid, questionid)
		if err == nil {
			var puntuacion int64 = 0
			review := &models.Review{
				Puntuacion: &puntuacion,
			}
			err = UpdateReview(answerid, questionid, review)
		}
	}

	if err != nil {
		log.Println("Error en corrector en ExecuteFullPruebas(): ", err)
	}
}

func AddAnswerPuntuacion(aid int64, qid int64, puntuacion int64) error {
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

func SubstractAnswerPuntuacion(aid int64, qid int64, puntuacion int64) error {
	return AddAnswerPuntuacion(aid, qid, -puntuacion)
}

func UpdateReview(aid int64, qid int64, review *models.Review) error {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var qa *dao.QuestionAnswer
		qa, err = dao.GetQuestionAnswerFromAnswer(db, aid, qid)
		if qa != nil && err == nil {
			err = SubstractAnswerPuntuacion(aid, qid, *qa.Puntuacion)
		}
		if err == nil {
			log.Println(*review.Puntuacion)
			err = AddAnswerPuntuacion(aid, qid, *review.Puntuacion)
			if err == nil {
				err = dao.PutReview(db, aid, qid, review)
			}
		}
	}
	return err
}
