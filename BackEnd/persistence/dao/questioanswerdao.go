// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package dao acts as a Data Access Object for the Types
package dao

import (
	"database/sql"
	"errors"
	"log"
	"uva-devtest/models"
	"uva-devtest/persistence/dbconnection"

	// Blank import of mysql driver
	_ "github.com/go-sql-driver/mysql"
)

func ToModelQuestionAnswer(q *QuestionAnswer) *models.QuestionAnswer {
	db, err := dbconnection.ConnectDb()
	username := ""
	if err == nil {
		var a *Answer
		a, err = GetAnswer(db, *q.IDRespuesta)
		if err == nil && a != nil {
			var u *User
			u, err = GetUserByID(db, a.Usuarioid)
			if err == nil && u != nil {
				username = *u.Username
			}
		}
	}
	mq := &models.QuestionAnswer{
		IDPregunta:  q.IDPregunta,
		IDRespuesta: q.IDRespuesta,
		Respuesta:   q.Respuesta,
		Corregida:   q.Corregida,
		Puntuacion:  q.Puntuacion,
		Username:    username,
	}
	mq.IndicesOpciones = append(mq.IndicesOpciones, q.IndicesOpciones...)
	return mq
}

func ToModelQuestionAnswers(qs []*QuestionAnswer) []*models.QuestionAnswer {
	var mqs = []*models.QuestionAnswer{}
	for _, itemCopy := range qs {
		mq := ToModelQuestionAnswer(itemCopy)
		mqs = append(mqs, mq)
	}
	return mqs
}

// Transforms some sql.Rows into a slice(array) of questionAnswers
// Param rows: Rows which contains database information returned
// Return []models.QuestionAnswer: QuestionsAnswer represented in rows
// Return error if any
func rowsToQuestionAnswers(rows *sql.Rows) ([]*QuestionAnswer, error) {
	var qas []*QuestionAnswer
	for rows.Next() {
		var qa QuestionAnswer
		err := rows.Scan(&qa.IDRespuesta, &qa.IDPregunta, &qa.Puntuacion, &qa.Corregida, &qa.Respuesta)
		if err != nil {
			log.Print(err)
			return qas, err
		}

		qas = append(qas, &qa)
	}
	return qas, nil
}

// Transforms rows into a single question
// Param rows: Rows which contains database info of 1 questionAnswer
// Return *models.QuestionAnswer: Question represented in rows
// Return error if something happens
func rowsToQuestionAnswer(rows *sql.Rows) (*QuestionAnswer, error) {
	var question *QuestionAnswer
	questions, err := rowsToQuestionAnswers(rows)
	if len(questions) >= 1 {
		question = questions[0]
	}
	return question, err
}

func addOptionsChosen(db *sql.DB, qa *QuestionAnswer) error {
	var opts []*Option
	var err error
	opts, err = GetOptionsQuestionAnswer(db, qa)
	if err == nil {
		for _, opt := range opts {
			qa.IndicesOpciones = append(qa.IndicesOpciones, opt.Indice)
		}
	}
	return err
}

func GetQuestionAnswersFromPTestQuestion(db *sql.DB, testid int64, questionid int64) ([]*QuestionAnswer, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var qas []*QuestionAnswer
	query, err := db.Prepare("SELECT R.* FROM RespuestaPregunta R JOIN RespuestaExamen E ON R.respuestaExamenid=E.id WHERE R.preguntaid=? AND E.testid=?")
	if err == nil {
		defer query.Close()
		rows, err := query.Query(questionid, testid)
		if err == nil {
			qas, err = rowsToQuestionAnswers(rows)
			if err == nil {
				for _, qa := range qas {
					err = addOptionsChosen(db, qa)
					if err != nil {
						return nil, err
					}
				}
			}
			return qas, err
		}
	}
	return nil, err
}

func GetQuestionAnswersFromAnswer(db *sql.DB, answerid int64) ([]*QuestionAnswer, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var qas []*QuestionAnswer
	query, err := db.Prepare("SELECT * FROM RespuestaPregunta WHERE respuestaExamenid=?")
	if err == nil {
		defer query.Close()
		rows, err := query.Query(answerid)
		if err == nil {
			qas, err = rowsToQuestionAnswers(rows)
			if err == nil {
				for _, qa := range qas {
					err = addOptionsChosen(db, qa)
					if err != nil {
						return nil, err
					}
				}
			}
			return qas, err
		}
	}
	return nil, err
}

func GetQuestionAnswerFromAnswer(db *sql.DB, answerid int64, questionid int64) (*QuestionAnswer, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var qas *QuestionAnswer
	query, err := db.Prepare("SELECT * FROM RespuestaPregunta WHERE preguntaid=? AND respuestaExamenid=?")
	if err == nil {
		defer query.Close()
		rows, err := query.Query(questionid, answerid)
		if err == nil {
			qas, err = rowsToQuestionAnswer(rows)
			if err == nil && qas != nil {
				err = addOptionsChosen(db, qas)
				if err != nil {
					return nil, err
				}
			}
			return qas, err
		}
	}
	return nil, err
}

func AddOptionQuestionAnswer(db *sql.DB, questionid int64, answerid int64, optionindex int64) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	query, err := db.Prepare("INSERT INTO OpcionRespuesta(respuestaExamenid, preguntaid, opcionindice) VALUES(?,?,?)")
	if err == nil {
		defer query.Close()
		_, err = query.Exec(answerid, questionid, optionindex)
	}
	return err
}

func PostQuestionAnswer(db *sql.DB, answerid int64, qa *models.QuestionAnswer) (*models.QuestionAnswer, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	query, err := db.Prepare("INSERT INTO RespuestaPregunta(respuestaExamenid, preguntaid, puntuacion, corregida, respuesta) VALUES(?,?,0,0,?)")
	if err == nil {
		defer query.Close()
		_, err = query.Exec(answerid, qa.IDPregunta, qa.Respuesta)
		if err == nil {
			qa.IDRespuesta = &answerid
			bfalse := false
			qa.Corregida = &bfalse
			var bzero int64 = 0
			qa.Puntuacion = &bzero
			for _, i := range qa.IndicesOpciones {
				err = AddOptionQuestionAnswer(db, *qa.IDPregunta, *qa.IDRespuesta, i)
				if err != nil {
					return nil, err
				}
			}
			return qa, err
		}
	}
	return nil, err
}

func DeleteIndiceOpcionesQuestionAnswer(db *sql.DB, answerid int64, questionid int64) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	query, err := db.Prepare("DELETE FROM OpcionRespuesta WHERE respuestaExamenid=? AND preguntaid=?")
	if err == nil {
		defer query.Close()
		_, err = query.Exec(answerid, questionid)
	}
	return err
}

func PutQuestionAnswer(db *sql.DB, answerid int64, questionid int64, qa *models.QuestionAnswer) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	query, err := db.Prepare("UPDATE RespuestaPregunta SET respuesta=? WHERE respuestaExamenid=? AND preguntaid=?")
	if err == nil {
		defer query.Close()
		_, err = query.Exec(qa.Respuesta, answerid, questionid)
		if err == nil {
			err = DeleteIndiceOpcionesQuestionAnswer(db, answerid, questionid)
			if err == nil {
				for _, i := range qa.IndicesOpciones {
					err = AddOptionQuestionAnswer(db, questionid, answerid, i)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return err
}

func DeleteQuestionAnswer(db *sql.DB, answerid int64, questionid int64) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	query, err := db.Prepare("DELETE FROM RespuestaPregunta WHERE respuestaExamenid=? AND preguntaid=?")
	if err == nil {
		defer query.Close()
		_, err = query.Exec(answerid, questionid)
		if err == nil {
			err = DeleteIndiceOpcionesQuestionAnswer(db, answerid, questionid)
			if err == nil {
				return err
			}
		}
	}
	return err
}

func PutReview(db *sql.DB, answerid int64, questionid int64, rev *models.Review) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	query, err := db.Prepare("UPDATE RespuestaPregunta SET puntuacion=?, corregida=1 WHERE respuestaExamenid=? AND preguntaid=?")
	if err == nil {
		defer query.Close()
		_, err = query.Exec(rev.Puntuacion, answerid, questionid)
	}
	return err
}
