// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package dao acts as a Data Access Object for the Types
package dao

import (
	"database/sql"
	"log"
	"uva-devtest/models"

	// Blank import of mysql driver
	_ "github.com/go-sql-driver/mysql"
)

func ToModelQuestionAnswer(q *QuestionAnswer) *models.QuestionAnswer {
	mq := &models.QuestionAnswer{
		IDPregunta:   q.IDPregunta,
		IDRespuesta:  q.IDRespuesta,
		IndiceOpcion: q.IndiceOpcion,
		Respuesta:    q.Respuesta,
		Corregida:    q.Corregida,
		Puntuacion:   q.Puntuacion,
	}
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
		err := rows.Scan(&qa.IDRespuesta, &qa.IDPregunta, &qa.Puntuacion, &qa.Corregida, &qa.IndiceOpcion, &qa.Respuesta)
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
