// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package dao acts as a Data Access Object for the Types
package dao

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
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
				Penalizacion:             q.Penalizacion,
				CantidadFavoritos:        q.CantidadFavoritos,
				Posicion:                 q.Posicion,
				NextID:                   q.NextID,
				PrevID:                   q.PrevID,
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
		err := rows.Scan(&q.ID, &q.Title, &q.Question, &q.EstimatedTime, &q.AutoCorrect, &q.Editable, &q.Usuarioid, &eleUni,
			&solu, &q.AccesoPublicoNoPublicada, &q.Penalizacion, &q.CantidadFavoritos)
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

func prepareQueryOrderBy(initQuery string, orderBy *string, idNombreConsulta string, cantidadFavoritosConsulta string, tiempoConsulta string) string {
	query := initQuery + " ORDER BY "
	if orderBy == nil {
		newDate := OrderByNewDate
		orderBy = &newDate
	}
	switch *orderBy {
	case OrderByLessFav:
		query += cantidadFavoritosConsulta + " ASC "
	case OrderByMoreFav:
		query += cantidadFavoritosConsulta + " DESC "
	case OrderByLessTime:
		query += tiempoConsulta + " ASC "
	case OrderByMoreTime:
		query += tiempoConsulta + " DESC "
	case OrderByOldDate:
		query += idNombreConsulta + " ASC "
	default: // order by new date
		query += idNombreConsulta + " DESC "
	}
	query += ", " + idNombreConsulta + " DESC "
	return query
}

func prepareQueryLimit(initQuery string, limit *int64, offset *int64) string {
	query := initQuery
	if limit != nil {
		query += " LIMIT " + strconv.FormatInt(*limit, 10) + " "
	}
	if offset != nil {
		query += " OFFSET " + strconv.FormatInt(*offset, 10) + " "
	}
	return query
}

func AddFiltersToQuery(hayWhere bool, initQuery string, tags [][]string, likeTitle *string, orderBy *string,
	limit *int64, offset *int64, idNombreConsulta string, idNombreSubconsulta string, tablaRelacionNombre string,
	titleConsulta string, cantidadFavoritosConsulta string, tiempoConsulta string) string {
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
	stPrepare = prepareQueryOrderBy(stPrepare, orderBy, idNombreConsulta, cantidadFavoritosConsulta, tiempoConsulta)
	stPrepare = prepareQueryLimit(stPrepare, limit, offset)
	return stPrepare
}

func addFiltersToQueryQuestion(hayWhere bool, initQuery string, tags [][]string, likeTitle *string, orderBy *string,
	limit *int64, offset *int64) string {
	return AddFiltersToQuery(hayWhere, initQuery, tags, likeTitle, orderBy, limit, offset, "id", "preguntaid", "PreguntaEtiqueta",
		"title", "cantidadFavoritos", "estimatedTime")
}

func addFiltersToQueryQuestionLongNames(hayWhere bool, initQuery string, tags [][]string, likeTitle *string, orderBy *string,
	limit *int64, offset *int64) string {
	return AddFiltersToQuery(hayWhere, initQuery, tags, likeTitle, orderBy, limit, offset, "P.id", "preguntaid", "PreguntaEtiqueta",
		"title", "P.cantidadFavoritos", "P.estimatedTime")
}

func GetAllEditQuestions(db *sql.DB, tags [][]string, likeTitle *string, orderBy *string,
	limit *int64, offset *int64) ([]*Question, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var qs []*Question
	stPrepare := "SELECT * FROM Pregunta WHERE editable=1 "
	stPrepare = addFiltersToQueryQuestion(true, stPrepare, tags, likeTitle, orderBy, limit, offset)
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

func GetEditQuestions(db *sql.DB, tags [][]string, likeTitle *string, orderBy *string,
	limit *int64, offset *int64) ([]*Question, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var qs []*Question
	stPrepare := "SELECT * FROM Pregunta WHERE editable=1 AND accesoPublicoNoPublicada=1"
	stPrepare = addFiltersToQueryQuestion(true, stPrepare, tags, likeTitle, orderBy, limit, offset)
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

func GetAllQuestions(db *sql.DB, tags [][]string, likeTitle *string, orderBy *string,
	limit *int64, offset *int64) ([]*Question, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var qs []*Question
	stPrepare := "SELECT * FROM Pregunta"
	stPrepare = addFiltersToQueryQuestion(false, stPrepare, tags, likeTitle, orderBy, limit, offset)
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

func GetQuestions(db *sql.DB, tags [][]string, likeTitle *string, orderBy *string,
	limit *int64, offset *int64) ([]*Question, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var qs []*Question
	stPrepare := "SELECT * FROM Pregunta WHERE accesoPublicoNoPublicada=1"
	stPrepare = addFiltersToQueryQuestion(true, stPrepare, tags, likeTitle, orderBy, limit, offset)
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
	query, err := db.Prepare("UPDATE Pregunta SET title=?, question=?, estimatedTime=?, autoCorrect=?, editable=?, usuarioid=?, " +
		" eleccionUnica=?, solucion=?, accesoPublicoNoPublicada=?, penalizacion=? WHERE id=? ")
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
	_, err = query.Exec(q.Title, q.Question, q.EstimatedTime, q.AutoCorrect, q.Editable, u.ID, eleUni, solucion, q.AccesoPublicoNoPublicada, q.Penalizacion, questionid)
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

func GetQuestionsOfUser(db *sql.DB, username string, tags [][]string, likeTitle *string, orderBy *string,
	limit *int64, offset *int64) ([]*Question, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err == nil && u != nil {
		var qs []*Question
		stPrepare := "SELECT * FROM Pregunta WHERE usuarioid=?"
		stPrepare = addFiltersToQueryQuestion(true, stPrepare, tags, likeTitle, orderBy, limit, offset)
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

func GetPublicEditQuestionsOfUser(db *sql.DB, username string, tags [][]string, likeTitle *string, orderBy *string,
	limit *int64, offset *int64) ([]*Question, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err == nil && u != nil {
		var qs []*Question
		stPrepare := "SELECT * FROM Pregunta WHERE usuarioid=? AND editable=1 AND accesoPublicoNoPublicada=1"
		stPrepare = addFiltersToQueryQuestion(true, stPrepare, tags, likeTitle, orderBy, limit, offset)
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

func GetEditQuestionsOfUser(db *sql.DB, username string, tags [][]string, likeTitle *string, orderBy *string,
	limit *int64, offset *int64) ([]*Question, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err == nil && u != nil {
		var qs []*Question
		stPrepare := "SELECT * FROM Pregunta WHERE usuarioid=? AND editable=1"
		stPrepare = addFiltersToQueryQuestion(true, stPrepare, tags, likeTitle, orderBy, limit, offset)
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
	query, err := db.Prepare("INSERT INTO Pregunta(title, question, estimatedTime, autoCorrect, editable, usuarioid, " +
		" eleccionUnica, solucion, accesoPublicoNoPublicada, penalizacion) " +
		"VALUES (?,?,?,?,?,?,?,?,?,?)")

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
	sol, err := query.Exec(q.Title, q.Question, q.EstimatedTime, q.AutoCorrect, q.Editable,
		u.ID, eleUni, solucion, q.AccesoPublicoNoPublicada, q.Penalizacion)
	if err == nil {
		qs := q
		qs.ID, err = sol.LastInsertId()
		return qs, err
	}
	return nil, err
}

func GetQuestionsFromTeam(db *sql.DB, teamname string, tags [][]string, likeTitle *string, orderBy *string,
	limit *int64, offset *int64) ([]*Question, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetTeam(db, teamname)
	if err == nil && u != nil {
		var qs []*Question
		stPrepare := "SELECT P.* FROM Pregunta P JOIN PreguntaEquipo E ON P.id=E.preguntaid WHERE E.equipoid=?"
		stPrepare = addFiltersToQueryQuestionLongNames(true, stPrepare, tags, likeTitle, orderBy, limit, offset)
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

func addNextPrevIdQuestion(qs *Question, testid int64) error {
	db, err := dbconnection.ConnectDb()
	if err != nil {
		return err
	}
	if err == nil {
		var prevQ, nextQ *Question
		prevQ, err = GetPreviousQuestionFromTest(db, testid, qs)
		if err == nil {
			if prevQ != nil {
				qs.PrevID = prevQ.ID // En FE tienes que poner por defecto 0 en ambos, ya que no se enviara nada si es 0
			} else {
				qs.PrevID = -1
			}
			if err == nil {
				nextQ, err = GetNextQuestionFromTest(db, testid, qs)
				if err == nil {
					if nextQ != nil {
						qs.NextID = nextQ.ID
					} else {
						qs.NextID = -1
					}
				}
			}
		}
	}
	return err
}

func addTestPreguntaAtributos(qs *Question, testid int64) error {
	db, err := dbconnection.ConnectDb()
	if err != nil {
		return err
	}
	var tp *TestPregunta
	tp, err = GetTestPreguntaAtributos(db, qs.ID, testid)
	if err != nil || tp == nil {
		return errors.New("no se pudo leer un valor de TestPregunta")
	}
	qs.ValorFinal = tp.ValorFinal
	qs.Posicion = *tp.Posicion
	return nil
}

func addTestPreguntasAtributos(qs []*Question, testid int64) error {
	for _, q := range qs {
		err := addTestPreguntaAtributos(q, testid)
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
	query, err := db.Prepare("SELECT P.* FROM Pregunta P JOIN TestPregunta T ON P.id=T.preguntaid WHERE T.testid=? ORDER BY T.posicion ASC")
	if err == nil {
		defer query.Close()
		rows, err := query.Query(testid)
		if err == nil {
			qs, err = rowsToQuestions(rows)
			if err == nil {
				err = addTestPreguntasAtributos(qs, testid)
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
				err = addTestPreguntaAtributos(qs, testid)
				if err == nil {
					err = addNextPrevIdQuestion(qs, testid)
					return qs, err
				}
			}
		}
	} else {
		log.Print(err)
	}
	return nil, err
}

func GetPreviousQuestionFromTest(db *sql.DB, testid int64, question *Question) (*Question, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var qs *Question
	query, err := db.Prepare("SELECT P.* FROM Pregunta P JOIN TestPregunta T ON P.id=T.preguntaid WHERE T.testid=? AND T.posicion<? " +
		" ORDER BY T.posicion DESC LIMIT 1;")
	if err == nil {
		defer query.Close()
		rows, err := query.Query(testid, question.Posicion)
		if err == nil {
			qs, err = rowsToQuestion(rows)
			return qs, err
		}
	} else {
		log.Print(err)
	}
	return nil, err
}

func GetNextQuestionFromTest(db *sql.DB, testid int64, question *Question) (*Question, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var qs *Question
	query, err := db.Prepare("SELECT P.* FROM Pregunta P JOIN TestPregunta T ON P.id=T.preguntaid WHERE T.testid=? AND T.posicion>? " +
		" ORDER BY T.posicion ASC LIMIT 1;")
	if err == nil {
		defer query.Close()
		rows, err := query.Query(testid, question.Posicion)
		if err == nil {
			qs, err = rowsToQuestion(rows)
			return qs, err
		}
	} else {
		log.Print(err)
	}
	return nil, err
}

func AddQuestionTest(db *sql.DB, questionid int64, testid int64, valorFinal int64, posicion int64) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	query, err := db.Prepare("INSERT INTO TestPregunta(testid, preguntaid, valorFinal, posicion) VALUES(?,?,?,?)")
	if err == nil {
		defer query.Close()
		_, err = query.Exec(testid, questionid, valorFinal, posicion)
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

func rowsToTestPreguntas(rows *sql.Rows) ([]*TestPregunta, error) {
	var tps []*TestPregunta
	for rows.Next() {
		var i TestPregunta
		var t1, t2 int64
		err := rows.Scan(&t1, &t2, &i.ValorFinal, &i.Posicion)
		if err != nil {
			log.Print(err)
			return tps, err
		}
		tps = append(tps, &i)
	}
	return tps, nil
}

func rowsToTestPregunta(rows *sql.Rows) (*TestPregunta, error) {
	var i *TestPregunta
	ts, err := rowsToTestPreguntas(rows)
	if len(ts) >= 1 {
		i = ts[0]
	}
	return i, err
}

func GetTestPreguntaAtributos(db *sql.DB, questionid int64, testid int64) (*TestPregunta, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var tP *TestPregunta
	query, err := db.Prepare("SELECT * FROM TestPregunta WHERE preguntaid=? AND testid=?")
	if err == nil {
		defer query.Close()
		rows, err := query.Query(questionid, testid)
		if err == nil {
			tP, err = rowsToTestPregunta(rows)
			return tP, err
		}
	} else {
		log.Print(err)
	}
	return nil, err
}

func GetFavoriteEditQuestions(db *sql.DB, username string, tags [][]string, likeTitle *string, orderBy *string,
	limit *int64, offset *int64) ([]*Question, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var qs []*Question
	stPrepare := "SELECT DISTINCT P.* FROM Pregunta P LEFT JOIN PreguntaEquipo E ON P.id=E.preguntaid LEFT JOIN EquipoUsuario U ON U.equipoid=E.equipoid " +
		" LEFT JOIN Usuario V ON V.id=U.usuarioid LEFT JOIN Usuario W ON W.id=P.usuarioid JOIN PreguntaFavorita F ON P.id=F.preguntaid JOIN Usuario Y ON Y.id=F.usuarioid " +
		" WHERE P.editable=1 AND Y.username=? AND (V.username=? OR P.accesoPublicoNoPublicada=1 OR W.username=?) "
	stPrepare = addFiltersToQueryQuestionLongNames(true, stPrepare, tags, likeTitle, orderBy, limit, offset)
	query, err := db.Prepare(stPrepare)
	if err == nil {
		defer query.Close()
		interfaceParams := FilterParamsSlicesToInterfaceArr(tags, likeTitle)
		var paramsSlice []interface{}
		paramsSlice = append(paramsSlice, username)
		paramsSlice = append(paramsSlice, username)
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

func GetFavoriteQuestions(db *sql.DB, username string, tags [][]string, likeTitle *string, orderBy *string,
	limit *int64, offset *int64) ([]*Question, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var qs []*Question
	stPrepare := "SELECT DISTINCT P.* FROM Pregunta P LEFT JOIN PreguntaEquipo E ON P.id=E.preguntaid LEFT JOIN EquipoUsuario U ON U.equipoid=E.equipoid " +
		" LEFT JOIN Usuario V ON V.id=U.usuarioid LEFT JOIN Usuario W ON W.id=P.usuarioid JOIN PreguntaFavorita F ON P.id=F.preguntaid JOIN Usuario Y ON Y.id=F.usuarioid " +
		" WHERE Y.username=? AND (V.username=? OR P.accesoPublicoNoPublicada=1 OR W.username=?) "
	stPrepare = addFiltersToQueryQuestionLongNames(true, stPrepare, tags, likeTitle, orderBy, limit, offset)
	query, err := db.Prepare(stPrepare)
	if err == nil {
		defer query.Close()
		interfaceParams := FilterParamsSlicesToInterfaceArr(tags, likeTitle)
		var paramsSlice []interface{}
		paramsSlice = append(paramsSlice, username)
		paramsSlice = append(paramsSlice, username)
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

func GetFavoriteQuestion(db *sql.DB, username string, questionid int64) (*Question, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var qs *Question
	stPrepare := "SELECT DISTINCT P.* FROM Pregunta P LEFT JOIN PreguntaEquipo E ON P.id=E.preguntaid LEFT JOIN EquipoUsuario U ON U.equipoid=E.equipoid " +
		" LEFT JOIN Usuario V ON V.id=U.usuarioid LEFT JOIN Usuario W ON W.id=P.usuarioid JOIN PreguntaFavorita F ON P.id=F.preguntaid JOIN Usuario Y ON Y.id=F.usuarioid " +
		" WHERE Y.username=? AND (V.username=? OR P.accesoPublicoNoPublicada=1 OR W.username=?) AND P.id=?"
	query, err := db.Prepare(stPrepare)
	if err == nil {
		defer query.Close()
		rows, err := query.Query(username, username, username, questionid)
		if err == nil {
			qs, err = rowsToQuestion(rows)
			if qs == nil {
				return nil, nil
			}
			return qs, err
		}
	} else {
		log.Print(err)
	}
	return nil, err
}

func AddFavoriteQuestion(db *sql.DB, username string, questionid int64) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err == nil && u != nil {
		stPrepare := "INSERT INTO PreguntaFavorita(usuarioid, preguntaid) VALUES(?,?)"
		var query *sql.Stmt
		query, err = db.Prepare(stPrepare)
		if err == nil {
			defer query.Close()
			_, err := query.Exec(u.ID, questionid)
			return err
		} else {
			log.Print(err)
		}
	} else if err == nil {
		err = errors.New(errorResourceNotFound)
	}
	return err
}

func RemoveFavoriteQuestion(db *sql.DB, username string, questionid int64) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err == nil && u != nil {
		stPrepare := "DELETE FROM PreguntaFavorita WHERE usuarioid=? AND preguntaid=?"
		var query *sql.Stmt
		query, err = db.Prepare(stPrepare)
		if err == nil {
			defer query.Close()
			_, err := query.Exec(u.ID, questionid)
			return err
		} else {
			log.Print(err)
		}
	} else if err == nil {
		err = errors.New(errorResourceNotFound)
	}
	return err
}

func GetAvailableEditQuestionsOfUser(db *sql.DB, username string, tags [][]string, likeTitle *string, orderBy *string,
	limit *int64, offset *int64) ([]*Question, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var qs []*Question
	stPrepare := "SELECT DISTINCT P.* FROM Pregunta P LEFT JOIN PreguntaEquipo E ON P.id=E.preguntaid LEFT JOIN EquipoUsuario U ON U.equipoid=E.equipoid " +
		" LEFT JOIN Usuario V ON V.id=U.usuarioid LEFT JOIN Usuario W ON W.id=P.usuarioid WHERE P.editable=1 AND (V.username=? OR P.accesoPublicoNoPublicada=1 OR W.username=?) "
	stPrepare = addFiltersToQueryQuestionLongNames(true, stPrepare, tags, likeTitle, orderBy, limit, offset)
	query, err := db.Prepare(stPrepare)
	if err == nil {
		defer query.Close()
		interfaceParams := FilterParamsSlicesToInterfaceArr(tags, likeTitle)
		var paramsSlice []interface{}
		paramsSlice = append(paramsSlice, username)
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

func GetAvailableQuestionsOfUser(db *sql.DB, username string, tags [][]string, likeTitle *string, orderBy *string,
	limit *int64, offset *int64) ([]*Question, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var qs []*Question
	stPrepare := "SELECT DISTINCT P.* FROM Pregunta P LEFT JOIN PreguntaEquipo E ON P.id=E.preguntaid LEFT JOIN EquipoUsuario U ON U.equipoid=E.equipoid " +
		" LEFT JOIN Usuario V ON V.id=U.usuarioid LEFT JOIN Usuario W ON W.id=P.usuarioid WHERE (V.username=? OR P.accesoPublicoNoPublicada=1 OR W.username=?) "
	stPrepare = addFiltersToQueryQuestionLongNames(true, stPrepare, tags, likeTitle, orderBy, limit, offset)
	query, err := db.Prepare(stPrepare)
	if err == nil {
		defer query.Close()
		interfaceParams := FilterParamsSlicesToInterfaceArr(tags, likeTitle)
		var paramsSlice []interface{}
		paramsSlice = append(paramsSlice, username)
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

func GetSharedQuestionsOfUser(db *sql.DB, username string, tags [][]string, likeTitle *string, orderBy *string,
	limit *int64, offset *int64) ([]*Question, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var qs []*Question
	stPrepare := "SELECT DISTINCT P.* FROM Pregunta P JOIN PreguntaEquipo E ON P.id=E.preguntaid JOIN EquipoUsuario U ON U.equipoid=E.equipoid JOIN Usuario V ON V.id=U.usuarioid WHERE V.username=?"
	stPrepare = addFiltersToQueryQuestionLongNames(true, stPrepare, tags, likeTitle, orderBy, limit, offset)
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
