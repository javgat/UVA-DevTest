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
				EleccionUnica: q.EleccionUnica, //Puede ser nil
				Solucion:      q.Solucion,      //Puede ser nil
				TipoPregunta:  q.TipoPregunta,
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
		var eleUni sql.NullBool
		var solu sql.NullString
		err := rows.Scan(&q.ID, &q.Title, &q.Question, &q.EstimatedTime, &q.AutoCorrect, &q.Editable, &q.Usuarioid, &eleUni, &solu)
		var tipo string
		if eleUni.Valid {
			q.EleccionUnica = eleUni.Bool
			tipo = models.QuestionTipoPreguntaOpciones
		}
		if solu.Valid {
			q.Solucion = solu.String
			tipo = models.QuestionTipoPreguntaString
		}
		q.TipoPregunta = &tipo
		if err != nil {
			log.Print(err)
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

func GetQuestions(db *sql.DB) ([]*Question, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var qs []*Question
	query, err := db.Prepare("SELECT * FROM Pregunta")
	if err == nil {
		defer query.Close()
		rows, err := query.Query()
		if err == nil {
			qs, err = rowsToQuestions(rows)
			return qs, err
		}
	}
	return nil, err
}

func GetQuestion(db *sql.DB, questionid int64) (*Question, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var qs *Question
	query, err := db.Prepare("SELECT * FROM Pregunta WHERE id=?")
	if err == nil {
		defer query.Close()
		rows, err := query.Query(questionid)
		if err == nil {
			qs, err = rowsToQuestion(rows)
			return qs, err
		}
	}
	return nil, err
}

func PutQuestion(db *sql.DB, questionid int64, q *models.Question) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, *q.Username)
	if err != nil || u == nil {
		return errors.New(errorResourceNotFound)
	}
	query, err := db.Prepare("UPDATE Pregunta SET title=?, question=?, estimatedTime=?, autoCorrect=?, editable=?, usuarioid=?, eleccionUnica=?, solucion=? WHERE id=? ")
	if err != nil {
		return err
	}
	var solucion *string = nil
	var eleUni *bool = nil
	if *q.TipoPregunta == models.QuestionTipoPreguntaOpciones {
		eleUni = &q.EleccionUnica
	} else if *q.TipoPregunta == models.QuestionTipoPreguntaString {
		solucion = &q.Solucion
	}
	defer query.Close()
	_, err = query.Exec(q.Title, q.Question, q.EstimatedTime, q.AutoCorrect, q.Editable, u.ID, eleUni, solucion, questionid)
	return err
}

func DeleteQuestion(db *sql.DB, questionid int64) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	query, err := db.Prepare("DELETE FROM Pregunta WHERE id=? ")
	if err != nil {
		return err
	}
	defer query.Close()
	_, err = query.Exec(questionid)
	return err
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
	if err == nil && u != nil {
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

func PostQuestion(db *sql.DB, q *models.Question, username string) (*models.Question, error) {
	if db == nil || q == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err != nil || u == nil {
		return nil, errors.New(errorResourceNotFound)
	}
	query, err := db.Prepare("INSERT INTO Pregunta(title, question, estimatedTime, autoCorrect, editable, usuarioid, testid, eleccionUnica, solucion) " +
		"VALUES (?,?,?,?,?,?,NULL,?,?)")

	if err != nil {
		return nil, err
	}
	var solucion *string = nil
	var eleUni *bool = nil
	if *q.TipoPregunta == models.QuestionTipoPreguntaOpciones {
		eleUni = &q.EleccionUnica
	} else if *q.TipoPregunta == models.QuestionTipoPreguntaString {
		solucion = &q.Solucion
	}
	defer query.Close()
	sol, err := query.Exec(q.Title, q.Question, q.EstimatedTime, q.AutoCorrect, q.Editable, u.ID, eleUni, solucion)
	if err == nil {
		qs := q
		qs.ID, err = sol.LastInsertId()
		return qs, err
	}
	return nil, err
}

func GetQuestionsFromTeam(db *sql.DB, teamname string) ([]*Question, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetTeam(db, teamname)
	if err == nil {
		var qs []*Question
		query, err := db.Prepare("SELECT P.* FROM Pregunta P JOIN PreguntaEquipo E ON P.id=E.preguntaid WHERE E.equipoid=?")
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

func GetQuestionFromTeam(db *sql.DB, teamname string, questionid int64) (*Question, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetTeam(db, teamname)
	if err == nil {
		var qs *Question
		var query *sql.Stmt
		query, err = db.Prepare("SELECT P.* FROM Pregunta P JOIN PreguntaEquipo E ON P.id=E.preguntaid WHERE E.equipoid=? AND P.id=?")
		if err == nil {
			defer query.Close()
			rows, err := query.Query(u.ID, questionid)
			if err == nil {
				qs, err = rowsToQuestion(rows)
				return qs, err
			}
		} else {
			log.Print(err)
		}
	}
	return nil, err
}

func AddQuestionTeam(db *sql.DB, questionid int64, teamname string) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	t, err := GetTeam(db, teamname)
	if err == nil {
		var query *sql.Stmt
		query, err = db.Prepare("INSERT INTO PreguntaEquipo(preguntaid, equipoid) VALUES(?,?)")
		if err == nil {
			defer query.Close()
			_, err = query.Exec(questionid, t.ID)
			return err
		}
	}
	return err
}

func RemoveQuestionTeam(db *sql.DB, questionid int64, teamname string) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	t, err := GetTeam(db, teamname)
	if err == nil {
		var query *sql.Stmt
		query, err = db.Prepare("DELETE FROM PreguntaEquipo WHERE preguntaid=? AND equipoid=?")
		if err == nil {
			defer query.Close()
			_, err = query.Exec(questionid, t.ID)
			return err
		}
	}
	return err
}

func GetQuestionsFromTest(db *sql.DB, testid int64) ([]*Question, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var qs []*Question
	query, err := db.Prepare("SELECT P.* FROM Pregunta P JOIN TestPregunta T ON P.id=T.preguntaid WHERE T.testid=?")
	if err == nil {
		defer query.Close()
		rows, err := query.Query(testid)
		if err == nil {
			qs, err = rowsToQuestions(rows)
			return qs, err
		}
	} else {
		log.Print(err)
	}
	return nil, err
}

func GetQuestionFromTest(db *sql.DB, testid int64, questionid int64) (*Question, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var qs *Question
	query, err := db.Prepare("SELECT P.* FROM Pregunta P JOIN TestPregunta T ON P.id=T.preguntaid WHERE T.testid=? AND P.id=?")
	if err == nil {
		defer query.Close()
		rows, err := query.Query(testid, questionid)
		if err == nil {
			qs, err = rowsToQuestion(rows)
			return qs, err
		}
	} else {
		log.Print(err)
	}
	return nil, err
}

func AddQuestionTest(db *sql.DB, questionid int64, testid int64) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	query, err := db.Prepare("INSERT INTO TestPregunta(testid, preguntaid, valorFinal) VALUES(?,?,1)")
	if err == nil {
		defer query.Close()
		_, err = query.Exec(testid, questionid)
		return err
	}
	return err
}

func RemoveQuestionTest(db *sql.DB, questionid int64, testid int64) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	query, err := db.Prepare("DELETE FROM TestPregunta WHERE testid=? AND preguntaid=?")
	if err == nil {
		defer query.Close()
		_, err = query.Exec(testid, questionid)
		return err
	}
	return err
}
