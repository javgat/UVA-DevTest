// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package dao acts as a Data Access Object for the Types
package dao

import (
	// Blank import of mysql driver
	"database/sql"
	"errors"
	"time"
	"uva-devtest/models"
	"uva-devtest/persistence/dbconnection"

	_ "github.com/go-sql-driver/mysql"
)

// Transforms some sql.Rows into a slice(array) of Answers
// Param rows: Rows which contains database information returned
// Return []models.Answer: Answers represented in rows
// Return error if any
func rowsToAnswers(rows *sql.Rows) ([]*Answer, error) {
	var answers []*Answer
	for rows.Next() {
		var t Answer
		err := rows.Scan(&t.ID, &t.Startime, &t.Finished, &t.Testid, &t.Usuarioid)
		if err != nil {
			return answers, err
		}
		answers = append(answers, &t)
	}
	return answers, nil
}

// Transforms rows into a single answer
// Param rows: Rows which contains database info of 1 answer
// Return *models.Answer: Answer represented in rows
// Return error if something happens
func rowsToAnswer(rows *sql.Rows) (*Answer, error) {
	var answer *Answer
	answers, err := rowsToAnswers(rows)
	if len(answers) >= 1 {
		answer = answers[0]
	}
	return answer, err
}

func ToModelAnswer(a *Answer) (*models.Answer, error) {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		u, err := GetUserByID(db, a.Usuarioid)
		if err == nil {
			mt := &models.Answer{
				Finished: a.Finished,
				Startime: a.Startime,
				Testid:   a.Testid,
				ID:       a.ID,
				Username: *u.Username,
			}
			return mt, nil
		}
	}
	return nil, errors.New(errorResourceNotFound)
}

func ToModelAnswers(as []*Answer) ([]*models.Answer, error) {
	var mas = []*models.Answer{}
	for _, itemCopy := range as {
		ma, err := ToModelAnswer(itemCopy)
		if err != nil {
			return nil, err
		}
		mas = append(mas, ma)
	}
	return mas, nil
}

func GetAnswers(db *sql.DB) ([]*Answer, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var a []*Answer
	query, err := db.Prepare("SELECT * FROM RespuestaExamen")
	if err == nil {
		defer query.Close()
		rows, err := query.Query()
		if err == nil {
			a, err = rowsToAnswers(rows)
			return a, err
		}
	}
	return nil, err
}

func GetAnswer(db *sql.DB, answerid int64) (*Answer, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var a *Answer
	query, err := db.Prepare("SELECT * FROM RespuestaExamen WHERE id=?")
	if err == nil {
		defer query.Close()
		rows, err := query.Query(answerid)
		if err == nil {
			a, err = rowsToAnswer(rows)
			return a, err
		}
	}
	return nil, err
}

func StartAnswer(db *sql.DB, username string, testid int64) (*Answer, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err != nil || u == nil {
		return nil, errors.New(errorResourceNotFound)
	}
	query, err := db.Prepare("INSERT INTO RespuestaExamen(startTime, finished, testid, usuarioid) VALUES (?,?,?,?)")

	if err != nil {
		return nil, err
	}
	defer query.Close()
	now := time.Now()
	var res sql.Result
	res, err = query.Exec(now, false, testid, u.ID)
	if err == nil {
		var id int64
		id, err = res.LastInsertId()
		if err == nil {
			bfalse := false
			ar := &Answer{
				Finished:  &bfalse,
				Testid:    testid,
				Usuarioid: u.ID,
				Startime:  now.String(),
				ID:        id,
			}
			return ar, err
		}
	}
	return nil, err
}

func FinishAnswer(db *sql.DB, answerid int64) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	query, err := db.Prepare("UPDATE RespuestaExamen SET finished=1 WHERE id=?")
	if err == nil {
		defer query.Close()
		_, err = query.Exec(answerid)
	}
	return err
}

func GetAnswersFromUserAnsweredTest(db *sql.DB, username string, testid int64) ([]*Answer, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err == nil {
		var a []*Answer
		query, err := db.Prepare("SELECT * FROM RespuestaExamen WHERE usuarioid=? AND testid=?")
		if err == nil {
			defer query.Close()
			rows, err := query.Query(u.ID, testid)
			if err == nil {
				a, err = rowsToAnswers(rows)
				return a, err
			}
		}
	}
	return nil, err
}

func GetAnswersFromUser(db *sql.DB, username string) ([]*Answer, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err == nil {
		var a []*Answer
		query, err := db.Prepare("SELECT * FROM RespuestaExamen WHERE usuarioid=?")
		if err == nil {
			defer query.Close()
			rows, err := query.Query(u.ID)
			if err == nil {
				a, err = rowsToAnswers(rows)
				return a, err
			}
		}
	}
	return nil, err
}

func GetAnswerFromUser(db *sql.DB, username string, answerid int64) (*Answer, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err == nil {
		var a *Answer
		query, err := db.Prepare("SELECT * FROM RespuestaExamen WHERE usuarioid=? AND id=?")
		if err == nil {
			defer query.Close()
			rows, err := query.Query(u.ID, answerid)
			if err == nil {
				a, err = rowsToAnswer(rows)
				return a, err
			}
		}
	}
	return nil, err
}
