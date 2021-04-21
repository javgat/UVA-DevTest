// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

package dao

import (
	"database/sql/driver"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func rowsQuestions(qs []*Question) *sqlmock.Rows {
	columns := []string{"title", "question", "estimatedTime", "autoCorrect", "editable", "usuarioid", "eleccionUnica", "solucion"}
	sqlcols := sqlmock.NewRows(columns)
	for _, q := range qs {
		sqlcols.AddRow(q.Title, q.Question, q.EstimatedTime, q.AutoCorrect, q.Editable, q.Usuarioid, q.EleccionUnica, q.Solucion)
	}
	return sqlcols
}

/*
func rowsQuestion(q *Question) *sqlmock.Rows {
	questions := []*Question{q}
	return rowsQuestions(questions)
}*/

func defaultQuestion() *Question {
	titulo := "Titulo"
	pregunta := "Pregunta?"
	btrue := true
	var estT int64 = 30
	q := &Question{
		Title:       &titulo,
		Question:    &pregunta,
		AutoCorrect: &btrue,
		Editable:    &btrue,
		//EleccionUnica: nil,
		EstimatedTime: &estT,
		ID:            1,
		Solucion:      "E",
		//Testid:        nil,
		Usuarioid: 1,
	}
	return q
}

func defaultQuestions() []*Question {
	qs := []*Question{defaultQuestion()}
	return qs
}

func expectGetQuestionsOfUser(mock sqlmock.Sqlmock, arg driver.Value, questions []*Question) {
	rows := rowsQuestions(questions)
	mock.ExpectPrepare("SELECT (.+) FROM Pregunta").ExpectQuery().WithArgs(arg).
		WillReturnRows(rows)
}

// Testing GetQuestionsOfUser
func TestGetQuestionsOfUserNilDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetQuestionsOfUser(mock, username, defaultQuestions())
	q, err := GetQuestionsOfUser(nil, username, nil, nil)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if q != nil {
		t.Log("q should be nil")
		t.Fail()
	}
}

func TestGetQuestionsOfUserFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetQuestionsOfUser(mock, username, defaultQuestions())
	q, err := GetQuestionsOfUser(nil, username, nil, nil)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if q != nil {
		t.Log("q should be nil")
		t.Fail()
	}
}
