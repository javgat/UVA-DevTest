// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package dao acts as a Data Access Object for the Types
package dao

import (
	"database/sql"
	"errors"
	"uva-devtest/models"
	"uva-devtest/persistence/dbconnection"

	// Blank import of mysql driver
	_ "github.com/go-sql-driver/mysql"
)

func ToModelQuestion(q *Question) (*models.Question, error) {

	db, err := dbconnection.ConnectDb()
	if err == nil {
		u, err := GetUserByID(db, q.Usuarioid)
		if err == nil {
			mq := &models.Question{
				AutoCorrect:   q.AutoCorrect,
				Editable:      q.Editable,
				EstimatedTime: q.EstimatedTime,
				ID:            q.ID,
				Question:      q.Question,
				Title:         q.Title,
				Username:      u.Username,
				Testid:        q.Testid,        //Puede ser nil
				EleccionUnica: q.EleccionUnica, //Puede ser nil
				Solucion:      q.Solucion,      //Puede ser nil
			}
			return mq, nil
		}
	}
	return nil, errors.New(errorResourceNotFound)
}

func ToModelQuestions(qs []*Question) ([]*models.Question, error) {
	var mqs = []*models.Question{}
	for _, itemCopy := range qs {
		mq, err := ToModelQuestion(itemCopy)
		if err != nil {
			return nil, err
		}
		mqs = append(mqs, mq)
	}
	return mqs, nil
}

// Transforms some sql.Rows into a slice(array) of questions
// Param rows: Rows which contains database information returned
// Return []models.Question: Questions represented in rows
// Return error if any
func rowsToQuestions(rows *sql.Rows) ([]*Question, error) {
	var questions []*Question
	for rows.Next() {
		var q Question
		err := rows.Scan(&q.ID, &q.Title, &q.Question, &q.EstimatedTime, &q.AutoCorrect, &q.Editable, &q.Usuarioid, &q.Testid, &q.EleccionUnica, &q.Solucion)
		if err != nil {
			return questions, err
		}

		questions = append(questions, &q)
	}
	return questions, nil
}

// Transforms rows into a single question
// Param rows: Rows which contains database info of 1 question
// Return *models.Question: Question represented in rows
// Return error if something happens
func rowsToQuestion(rows *sql.Rows) (*Question, error) {
	var question *Question
	questions, err := rowsToQuestions(rows)
	if len(questions) >= 1 {
		question = questions[0]
	}
	return question, err
}

//NOTESTED:

func GetQuestionsOfUser(db *sql.DB, username string) ([]*Question, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err == nil {
		var qs []*Question
		query, err := db.Prepare("SELECT * FROM Pregunta WHERE usuarioid=?")
		if err == nil {
			defer query.Close()
			rows, err := query.Query(u.ID)
			if err == nil {
				qs, err = rowsToQuestions(rows)
				return qs, err
			}
		}
	}
	return nil, err
}

//NOTESTED:

func GetQuestionOfUser(db *sql.DB, username string, qid int64) (*Question, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err == nil {
		var q *Question
		query, err := db.Prepare("SELECT * FROM Pregunta WHERE usuarioid=? AND id=?")
		if err == nil {
			defer query.Close()
			rows, err := query.Query(u.ID, qid)
			if err == nil {
				q, err = rowsToQuestion(rows)
				return q, err
			}
		}
	}
	return nil, err
}

//NOTESTED:

func PostQuestion(db *sql.DB, q *models.Question, username string) error {
	if db == nil || q == nil {
		return errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err != nil || u == nil {
		return errors.New(errorResourceNotFound)
	}
	query, err := db.Prepare("INSERT INTO Pregunta(title, question, estimatedTime, autoCorrect, editable, usuarioid, testid, eleccionUnica, solucion) " +
		"VALUES (?,?,?,?,?,?,NULL,?,?)")

	if err != nil {
		return err
	}
	defer query.Close()
	_, err = query.Exec(q.Title, q.Question, q.EstimatedTime, q.AutoCorrect, q.Editable, u.ID, q.EleccionUnica, q.Solucion)
	return err
}
