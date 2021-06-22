// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package probador executes Pruebas for the Code QuestionAnswers
package probador

import (
	"log"
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
		log.Println("Error en probador en ExecutePrePruebas(): ", err)
	}
}
