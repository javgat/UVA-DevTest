// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package dao acts as a Data Access Object for the Types
package dao

import (
	"database/sql"
	"errors"
	"fmt"
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
				AutoCorrect:              q.AutoCorrect,
				Editable:                 q.Editable,
				EstimatedTime:            q.EstimatedTime,
				ID:                       q.ID,
				Question:                 q.Question,
				Title:                    q.Title,
				Username:                 u.Username,
				EleccionUnica:            q.EleccionUnica, //Puede ser nil
				Solucion:                 q.Solucion,      //Puede ser nil
				TipoPregunta:             q.TipoPregunta,
				ValorFinal:               q.ValorFinal,
				AccesoPublicoNoPublicada: q.AccesoPublicoNoPublicada,
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
		err := rows.Scan(&q.ID, &q.Title, &q.Question, &q.EstimatedTime, &q.AutoCorrect, &q.Editable, &q.Usuarioid, &eleUni, &solu, &q.AccesoPublicoNoPublicada)
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

func tagsSlicesToSlice(tags [][]string) []string {
	var retTags []string
	for _, arr := range tags {
		retTags = append(retTags, arr...)
	}
	return retTags

}

func FilterParamsSlicesToInterfaceArr(tags [][]string, likeTitle *string) []interface{} {
	arrTags := tagsSlicesToSlice(tags)
	hayTitle := 0
	if likeTitle != nil && *likeTitle != "" {
		hayTitle = 1
	}
	interfaceParams := make([]interface{}, len(arrTags)+hayTitle)
	for i := range arrTags {
		interfaceParams[i] = arrTags[i]
	}
	if hayTitle == 1 {
		interfaceParams[len(arrTags)] = "%" + *likeTitle + "%"
	}
	return interfaceParams
}

func prepareQueryTags(initQuery string, tags [][]string, idNombreConsulta string, idNombreSubconsulta string, tablaRelacionNombre string) string {
	query := initQuery + " ( "
	primerOr := true
	for _, arr_ands := range tags {
		if !primerOr {
			query = query + "OR "
		} else {
			primerOr = false
		}
		primerAnd := true
		query = query + idNombreConsulta + " IN ( SELECT " + idNombreSubconsulta + " FROM " + tablaRelacionNombre +
			" WHERE etiquetanombre IN ( "
		for range arr_ands {
			if !primerAnd {
				query = query + ", "
			} else {
				primerAnd = false
			}
			query = query + "? "
		}
		query = query + ") GROUP BY " + idNombreSubconsulta + " HAVING COUNT(DISTINCT etiquetanombre) = " + fmt.Sprint(len(arr_ands)) + ") "
	}
	query = query + " ) "
	return query
}

func prepareQueryLikeTitle(initQuery string, likeTitle *string, titleConsulta string) string {
	query := initQuery + " ( " + titleConsulta + " LIKE ? )"
	return query
}

func addFiltersToQuery(hayWhere bool, initQuery string, tags [][]string, likeTitle *string, idNombreConsulta string, idNombreSubconsulta string, tablaRelacionNombre string, titleConsulta string) string {
	stPrepare := initQuery
	nexoString := " AND "
	if !hayWhere {
		nexoString = " WHERE "
	}
	if len(tags) > 0 {
		stPrepare = stPrepare + nexoString
		stPrepare = prepareQueryTags(stPrepare, tags, idNombreConsulta, idNombreSubconsulta, tablaRelacionNombre)
		nexoString = " AND "
	}
	if likeTitle != nil && *likeTitle != "" {
		stPrepare = stPrepare + nexoString
		stPrepare = prepareQueryLikeTitle(stPrepare, likeTitle, titleConsulta)
	}
	return stPrepare
}

func addFiltersToQueryQuestion(hayWhere bool, initQuery string, tags [][]string, likeTitle *string) string {
	return addFiltersToQuery(hayWhere, initQuery, tags, likeTitle, "id", "preguntaid", "PreguntaEtiqueta", "title")
}

func addFiltersToQueryQuestionLongNames(hayWhere bool, initQuery string, tags [][]string, likeTitle *string) string {
	return addFiltersToQuery(hayWhere, initQuery, tags, likeTitle, "P.id", "preguntaid", "PreguntaEtiqueta", "title")
}

func GetAllEditQuestions(db *sql.DB, tags [][]string, likeTitle *string) ([]*Question, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var qs []*Question
	stPrepare := "SELECT * FROM Pregunta WHERE editable=1 "
	stPrepare = addFiltersToQueryQuestion(true, stPrepare, tags, likeTitle)
	query, err := db.Prepare(stPrepare)
	if err == nil {
		defer query.Close()
		interfaceParams := FilterParamsSlicesToInterfaceArr(tags, likeTitle)
		rows, err := query.Query(interfaceParams...)
		if err == nil {
			qs, err = rowsToQuestions(rows)
			return qs, err
		}
	}
	return nil, err
}

func GetEditQuestions(db *sql.DB, tags [][]string, likeTitle *string) ([]*Question, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var qs []*Question
	stPrepare := "SELECT * FROM Pregunta WHERE editable=1 AND accesoPublicoNoPublicada=1"
	stPrepare = addFiltersToQueryQuestion(true, stPrepare, tags, likeTitle)
	query, err := db.Prepare(stPrepare)
	if err == nil {
		defer query.Close()
		interfaceParams := FilterParamsSlicesToInterfaceArr(tags, likeTitle)
		rows, err := query.Query(interfaceParams...)
		if err == nil {
			qs, err = rowsToQuestions(rows)
			return qs, err
		}
	}
	return nil, err
}

func GetAllQuestions(db *sql.DB, tags [][]string, likeTitle *string) ([]*Question, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var qs []*Question
	stPrepare := "SELECT * FROM Pregunta"
	stPrepare = addFiltersToQueryQuestion(false, stPrepare, tags, likeTitle)
	query, err := db.Prepare(stPrepare)
	if err == nil {
		defer query.Close()
		interfaceParams := FilterParamsSlicesToInterfaceArr(tags, likeTitle)
		rows, err := query.Query(interfaceParams...)
		if err == nil {
			qs, err = rowsToQuestions(rows)
			return qs, err
		}
	}
	return nil, err
}

func GetQuestions(db *sql.DB, tags [][]string, likeTitle *string) ([]*Question, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var qs []*Question
	stPrepare := "SELECT * FROM Pregunta WHERE accesoPublicoNoPublicada=1"
	stPrepare = addFiltersToQueryQuestion(true, stPrepare, tags, likeTitle)
	query, err := db.Prepare(stPrepare)
	if err == nil {
		defer query.Close()
		interfaceParams := FilterParamsSlicesToInterfaceArr(tags, likeTitle)
		rows, err := query.Query(interfaceParams...)
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
	query, err := db.Prepare("UPDATE Pregunta SET title=?, question=?, estimatedTime=?, autoCorrect=?, editable=?, usuarioid=?, eleccionUnica=?, solucion=?, accesoPublicoNoPublicada=? WHERE id=? ")
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
	_, err = query.Exec(q.Title, q.Question, q.EstimatedTime, q.AutoCorrect, q.Editable, u.ID, eleUni, solucion, q.AccesoPublicoNoPublicada, questionid)
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

func GetQuestionsOfUser(db *sql.DB, username string, tags [][]string, likeTitle *string) ([]*Question, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err == nil && u != nil {
		var qs []*Question
		stPrepare := "SELECT * FROM Pregunta WHERE usuarioid=?"
		stPrepare = addFiltersToQueryQuestion(true, stPrepare, tags, likeTitle)
		query, err := db.Prepare(stPrepare)
		if err == nil {
			defer query.Close()
			interfaceParams := FilterParamsSlicesToInterfaceArr(tags, likeTitle)
			var paramsSlice []interface{}
			paramsSlice = append(paramsSlice, u.ID)
			interfaceParams = append(paramsSlice, interfaceParams...)
			rows, err := query.Query(interfaceParams...)
			if err == nil {
				qs, err = rowsToQuestions(rows)
				return qs, err
			}
		}
	}
	return nil, err
}

func GetPublicEditQuestionsOfUser(db *sql.DB, username string, tags [][]string, likeTitle *string) ([]*Question, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err == nil && u != nil {
		var qs []*Question
		stPrepare := "SELECT * FROM Pregunta WHERE usuarioid=? AND editable=1 AND accesoPublicoNoPublicada=1"
		stPrepare = addFiltersToQueryQuestion(true, stPrepare, tags, likeTitle)
		query, err := db.Prepare(stPrepare)
		if err == nil {
			defer query.Close()
			interfaceParams := FilterParamsSlicesToInterfaceArr(tags, likeTitle)
			var paramsSlice []interface{}
			paramsSlice = append(paramsSlice, u.ID)
			interfaceParams = append(paramsSlice, interfaceParams...)
			rows, err := query.Query(interfaceParams...)
			if err == nil {
				qs, err = rowsToQuestions(rows)
				return qs, err
			}
		}
	}
	return nil, err
}

func GetEditQuestionsOfUser(db *sql.DB, username string, tags [][]string, likeTitle *string) ([]*Question, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err == nil && u != nil {
		var qs []*Question
		stPrepare := "SELECT * FROM Pregunta WHERE usuarioid=? AND editable=1"
		stPrepare = addFiltersToQueryQuestion(true, stPrepare, tags, likeTitle)
		query, err := db.Prepare(stPrepare)
		if err == nil {
			defer query.Close()
			interfaceParams := FilterParamsSlicesToInterfaceArr(tags, likeTitle)
			var paramsSlice []interface{}
			paramsSlice = append(paramsSlice, u.ID)
			interfaceParams = append(paramsSlice, interfaceParams...)
			rows, err := query.Query(interfaceParams...)
			if err == nil {
				qs, err = rowsToQuestions(rows)
				return qs, err
			}
		}
	}
	return nil, err
}

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

func PostQuestion(db *sql.DB, q *models.Question, username string) (*models.Question, error) {
	if db == nil || q == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err != nil || u == nil {
		return nil, errors.New(errorResourceNotFound)
	}
	query, err := db.Prepare("INSERT INTO Pregunta(title, question, estimatedTime, autoCorrect, editable, usuarioid, eleccionUnica, solucion, accesoPublicoNoPublicada) " +
		"VALUES (?,?,?,?,?,?,?,?,?)")

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
	sol, err := query.Exec(q.Title, q.Question, q.EstimatedTime, q.AutoCorrect, q.Editable, u.ID, eleUni, solucion, q.AccesoPublicoNoPublicada)
	if err == nil {
		qs := q
		qs.ID, err = sol.LastInsertId()
		return qs, err
	}
	return nil, err
}

func GetQuestionsFromTeam(db *sql.DB, teamname string, tags [][]string, likeTitle *string) ([]*Question, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetTeam(db, teamname)
	if err == nil && u != nil {
		var qs []*Question
		stPrepare := "SELECT P.* FROM Pregunta P JOIN PreguntaEquipo E ON P.id=E.preguntaid WHERE E.equipoid=?"
		stPrepare = addFiltersToQueryQuestionLongNames(true, stPrepare, tags, likeTitle)
		query, err := db.Prepare(stPrepare)
		if err == nil {
			defer query.Close()
			interfaceParams := FilterParamsSlicesToInterfaceArr(tags, likeTitle)
			var paramsSlice []interface{}
			paramsSlice = append(paramsSlice, u.ID)
			interfaceParams = append(paramsSlice, interfaceParams...)
			rows, err := query.Query(interfaceParams...)
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

func addValorFinal(qs *Question, testid int64) error {
	db, err := dbconnection.ConnectDb()
	if err != nil {
		return err
	}
	var vF *int64
	vF, err = GetValorFinal(db, qs.ID, testid)
	if err != nil || vF == nil {
		return errors.New("no se pudo leer un valor final")
	}
	qs.ValorFinal = vF
	return nil
}

func addValoresFinales(qs []*Question, testid int64) error {
	for _, q := range qs {
		err := addValorFinal(q, testid)
		if err != nil {
			return err
		}
	}
	return nil
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
			if err == nil {
				err = addValoresFinales(qs, testid)
				return qs, err
			}
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
			if qs != nil && err == nil {
				err = addValorFinal(qs, testid)
				return qs, err
			}
		}
	} else {
		log.Print(err)
	}
	return nil, err
}

func AddQuestionTest(db *sql.DB, questionid int64, testid int64, valorFinal int64) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	query, err := db.Prepare("INSERT INTO TestPregunta(testid, preguntaid, valorFinal) VALUES(?,?,?)")
	if err == nil {
		defer query.Close()
		_, err = query.Exec(testid, questionid, valorFinal)
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

func GetEditQuestionsFromTag(db *sql.DB, tag string) ([]*Question, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var qs []*Question
	query, err := db.Prepare("SELECT P.* FROM Pregunta P JOIN PreguntaEtiqueta E ON P.id=E.preguntaid WHERE E.etiquetaNombre=? AND P.editable=1")
	if err == nil {
		defer query.Close()
		rows, err := query.Query(tag)
		if err == nil {
			qs, err = rowsToQuestions(rows)
			return qs, err
		}
	} else {
		log.Print(err)
	}
	return nil, err
}

func GetQuestionsFromTag(db *sql.DB, tag string) ([]*Question, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var qs []*Question
	query, err := db.Prepare("SELECT P.* FROM Pregunta P JOIN PreguntaEtiqueta E ON P.id=E.preguntaid WHERE E.etiquetaNombre=?")
	if err == nil {
		defer query.Close()
		rows, err := query.Query(tag)
		if err == nil {
			qs, err = rowsToQuestions(rows)
			return qs, err
		}
	} else {
		log.Print(err)
	}
	return nil, err
}

func rowsToInts64(rows *sql.Rows) ([]*int64, error) {
	var ints []*int64
	for rows.Next() {
		var i int64
		err := rows.Scan(&i)
		if err != nil {
			log.Print(err)
			return ints, err
		}
		ints = append(ints, &i)
	}
	return ints, nil
}

func rowsToInt64(rows *sql.Rows) (*int64, error) {
	var i *int64
	ints, err := rowsToInts64(rows)
	if len(ints) >= 1 {
		i = ints[0]
	}
	return i, err
}

func GetValorFinal(db *sql.DB, questionid int64, testid int64) (*int64, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var vF *int64
	query, err := db.Prepare("SELECT valorFinal FROM TestPregunta WHERE preguntaid=? AND testid=?")
	if err == nil {
		defer query.Close()
		rows, err := query.Query(questionid, testid)
		if err == nil {
			vF, err = rowsToInt64(rows)
			return vF, err
		}
	} else {
		log.Print(err)
	}
	return nil, err
}

func GetSharedQuestionsOfUser(db *sql.DB, username string, tags [][]string, likeTitle *string) ([]*Question, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var qs []*Question
	stPrepare := "SELECT DISTINCT P.* FROM Pregunta P JOIN PreguntaEquipo E ON P.id=E.preguntaid JOIN EquipoUsuario U ON U.equipoid=E.equipoid JOIN Usuario V ON V.id=U.usuarioid WHERE V.username=?"
	stPrepare = addFiltersToQueryQuestionLongNames(true, stPrepare, tags, likeTitle)
	query, err := db.Prepare(stPrepare)
	if err == nil {
		defer query.Close()
		interfaceParams := FilterParamsSlicesToInterfaceArr(tags, likeTitle)
		var paramsSlice []interface{}
		paramsSlice = append(paramsSlice, username)
		interfaceParams = append(paramsSlice, interfaceParams...)
		rows, err := query.Query(interfaceParams...)
		if err == nil {
			qs, err = rowsToQuestions(rows)
			return qs, err
		}
	} else {
		log.Print(err)
	}
	return nil, err
}

func GetSharedQuestionFromUser(db *sql.DB, username string, questionid int64) (*Question, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var qs *Question
	query, err := db.Prepare("SELECT P.* FROM Pregunta P JOIN PreguntaEquipo E ON P.id=E.preguntaid JOIN EquipoUsuario U ON U.equipoid=E.equipoid JOIN Usuario V ON V.id=U.usuarioid WHERE V.username=? AND P.id=?")
	if err == nil {
		defer query.Close()
		rows, err := query.Query(username, questionid)
		if err == nil {
			qs, err = rowsToQuestion(rows)
			return qs, err
		}
	} else {
		log.Print(err)
	}
	return nil, err
}
