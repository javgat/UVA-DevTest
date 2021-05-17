// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package dao acts as a Data Access Object for the Types
package dao

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
	"strings"
	"uva-devtest/models"
)

func ToModelTag(dt *Tag) *models.Tag {
	mt := &models.Tag{
		Tag: dt.Tag,
	}
	return mt
}

func ToModelTags(ts []*Tag) []*models.Tag {
	var mts = []*models.Tag{}
	for _, itemCopy := range ts {
		mt := ToModelTag(itemCopy)
		mts = append(mts, mt)
	}
	return mts
}

// Transforms some sql.Rows into a slice(array) of tags
// Param rows: Rows which contains database information returned
// Return []Tag: Tag represented in rows
// Return error if any
func rowsToTags(rows *sql.Rows) ([]*Tag, error) {
	var tags []*Tag
	for rows.Next() {
		var tag Tag
		err := rows.Scan(&tag.Tag)
		if err != nil {
			log.Print(err)
			return tags, err
		}
		tags = append(tags, &tag)
	}
	return tags, nil
}

// Transforms rows into a single tag
// Param rows: Rows which contains database info of 1 tag
// Return *Tag: Question represented in rows
// Return error if something happens
func rowsToTag(rows *sql.Rows) (*Tag, error) {
	var tag *Tag
	tags, err := rowsToTags(rows)
	if len(tags) >= 1 {
		tag = tags[0]
	}
	return tag, err
}

func getTagsQuery(orderBy *string, limit *int64, offset *int64) string {
	var order, orderComp string
	if orderBy == nil {
		orderComp = ""
	} else {
		orderComp = *orderBy
	}
	switch orderComp {
	case TagOrderByLastAlpha:
		order = " nombre ASC "
	case TagOrderByMoreQuestion:
		order = " countpreg DESC "
	case TagOrderByLessQuestion:
		order = " countpreg ASC "
	case TagOrderByMoreTest:
		order = " counttest DESC "
	case TagOrderByLessTest:
		order = " counttest ASC "
	default: //case TagOrderByFirstAlpha:
		order = " nombre DESC "
	}
	withs := "WITH ECountPregunta AS ( SELECT etiquetanombre, COUNT(*) AS countpreg FROM PreguntaEtiqueta GROUP BY etiquetanombre), " +
		"ECountTest AS ( SELECT etiquetanombre, COUNT(*) AS counttest FROM TestEtiqueta GROUP BY etiquetanombre) "
	stPrepare := withs + " SELECT E.* FROM Etiqueta E LEFT JOIN ECountPregunta P ON P.etiquetanombre=E.nombre LEFT JOIN " +
		" ECountTest T ON T.etiquetanombre=E.nombre WHERE E.nombre LIKE ? ORDER BY " + order
	if limit != nil {
		stPrepare = stPrepare + " LIMIT " + strconv.FormatInt(*limit, 10)
	}
	if offset != nil {
		stPrepare = stPrepare + " OFFSET " + strconv.FormatInt(*offset, 10)
	}
	return stPrepare
}

func GetTags(db *sql.DB, likeTag *string, orderBy *string, limit *int64, offset *int64) ([]*Tag, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var ts []*Tag
	stPrepare := getTagsQuery(orderBy, limit, offset)
	query, err := db.Prepare(stPrepare)
	likeCom := "%"
	if likeTag != nil {
		likeCom = likeCom + *likeTag + likeCom
	}
	if err == nil {
		defer query.Close()
		rows, err := query.Query(likeCom)
		if err == nil {
			ts, err = rowsToTags(rows)
			return ts, err
		}
	}
	return nil, err
}

func GetTag(db *sql.DB, nombre string) (*Tag, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var ts *Tag
	query, err := db.Prepare("SELECT * FROM Etiqueta WHERE nombre=?")
	if err == nil {
		defer query.Close()
		rows, err := query.Query(nombre)
		if err == nil {
			ts, err = rowsToTag(rows)
			return ts, err
		}
	}
	return nil, err
}

func GetQuestionTags(db *sql.DB, questionid int64) ([]*Tag, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var ts []*Tag
	query, err := db.Prepare("SELECT E.* FROM Etiqueta E JOIN PreguntaEtiqueta P ON E.nombre=P.etiquetaNombre WHERE P.preguntaid=?")
	if err == nil {
		defer query.Close()
		rows, err := query.Query(questionid)
		if err == nil {
			ts, err = rowsToTags(rows)
			return ts, err
		}
	}
	return nil, err
}

func GetQuestionTag(db *sql.DB, questionid int64, nombre string) (*Tag, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var ts *Tag
	query, err := db.Prepare("SELECT E.* FROM Etiqueta E JOIN PreguntaEtiqueta P ON E.nombre=P.etiquetaNombre WHERE P.preguntaid=? AND E.nombre=?")
	if err == nil {
		defer query.Close()
		rows, err := query.Query(questionid, nombre)
		if err == nil {
			ts, err = rowsToTag(rows)
			return ts, err
		}
	}
	return nil, err
}

func GetTestTags(db *sql.DB, testid int64) ([]*Tag, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var ts []*Tag
	query, err := db.Prepare("SELECT E.* FROM Etiqueta E JOIN TestEtiqueta T ON E.nombre=T.etiquetaNombre WHERE T.testid=?")
	if err == nil {
		defer query.Close()
		rows, err := query.Query(testid)
		if err == nil {
			ts, err = rowsToTags(rows)
			return ts, err
		}
	}
	return nil, err
}

func GetTestTag(db *sql.DB, testid int64, nombre string) (*Tag, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var ts *Tag
	query, err := db.Prepare("SELECT E.* FROM Etiqueta E JOIN TestEtiqueta T ON E.nombre=T.etiquetaNombre WHERE T.testid=? AND E.nombre=?")
	if err == nil {
		defer query.Close()
		rows, err := query.Query(testid, nombre)
		if err == nil {
			ts, err = rowsToTag(rows)
			return ts, err
		}
	}
	return nil, err
}

func CreateTag(db *sql.DB, nombre string) (*Tag, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	nombre = strings.ToLower(nombre)
	ts := &Tag{
		Tag: &nombre,
	}
	query, err := db.Prepare("INSERT INTO Etiqueta(nombre) VALUES(?)")
	if err == nil {
		defer query.Close()
		_, err = query.Exec(ts.Tag)
	}
	return ts, err
}

func AddQuestionTag(db *sql.DB, questionid int64, nombre string) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	nombre = strings.ToLower(nombre)
	tag, err := GetTag(db, nombre)
	if err == nil {
		if tag == nil {
			_, err = CreateTag(db, nombre)
			if err != nil {
				return err
			}
		}
		query, err := db.Prepare("INSERT INTO PreguntaEtiqueta(preguntaid, etiquetanombre) VALUES(?,?)")
		if err == nil {
			defer query.Close()
			_, err = query.Exec(questionid, nombre)
		}
		return err

	}
	return err
}

func RemoveQuestionTag(db *sql.DB, questionid int64, nombre string) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	nombre = strings.ToLower(nombre)
	query, err := db.Prepare("DELETE FROM PreguntaEtiqueta WHERE preguntaid=? AND etiquetanombre=?")
	if err == nil {
		defer query.Close()
		_, err = query.Exec(questionid, nombre)
	}
	return err
}

func AddTestTag(db *sql.DB, testid int64, nombre string) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	nombre = strings.ToLower(nombre)
	tag, err := GetTag(db, nombre)
	if err == nil {
		if tag == nil {
			_, err = CreateTag(db, nombre)
			if err != nil {
				return err
			}
		}
		query, err := db.Prepare("INSERT INTO TestEtiqueta(testid, etiquetanombre) VALUES(?,?)")
		if err == nil {
			defer query.Close()
			_, err = query.Exec(testid, nombre)
		}
		return err

	}
	return err
}

func RemoveTestTag(db *sql.DB, testid int64, nombre string) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	nombre = strings.ToLower(nombre)
	query, err := db.Prepare("DELETE FROM TestEtiqueta WHERE testid=? AND etiquetanombre=?")
	if err == nil {
		defer query.Close()
		_, err = query.Exec(testid, nombre)
	}
	return err
}
