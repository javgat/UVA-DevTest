// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package dao acts as a Data Access Object for the Types
package dao

import (
	"database/sql"
	"errors"
	"log"
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

func GetTags(db *sql.DB) ([]*Tag, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var ts []*Tag
	query, err := db.Prepare("SELECT * FROM Etiqueta")
	if err == nil {
		defer query.Close()
		rows, err := query.Query()
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
