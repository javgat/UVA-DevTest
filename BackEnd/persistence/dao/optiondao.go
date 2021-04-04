// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package dao acts as a Data Access Object for the Types
package dao

import (
	// Blank import of mysql driver
	"database/sql"
	"errors"
	"log"
	"uva-devtest/models"

	_ "github.com/go-sql-driver/mysql"
)

func ToModelOption(o *Option) *models.Option {
	mo := &models.Option{
		Correcta:   o.Correcta,
		Indice:     o.Indice,
		Preguntaid: o.Preguntaid,
		Texto:      o.Texto,
	}
	return mo
}

func ToModelOptions(os []*Option) []*models.Option {
	var mos = []*models.Option{}
	for _, itemCopy := range os {
		mo := ToModelOption(itemCopy)
		mos = append(mos, mo)
	}
	return mos
}

func rowsToOptions(rows *sql.Rows) ([]*Option, error) {
	var options []*Option
	for rows.Next() {
		var o Option
		err := rows.Scan(&o.Indice, &o.Texto, &o.Correcta, &o.Preguntaid)
		if err != nil {
			log.Print(err)
			return options, err
		}
		options = append(options, &o)
	}
	return options, nil
}

func rowsToOption(rows *sql.Rows) (*Option, error) {
	var option *Option
	options, err := rowsToOptions(rows)
	if len(options) >= 1 {
		option = options[0]
	}
	return option, err
}

func GetOptionsQuestion(db *sql.DB, questionid int64) ([]*Option, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var os []*Option
	query, err := db.Prepare("SELECT * FROM Opcion WHERE preguntaid=?")
	if err == nil {
		defer query.Close()
		rows, err := query.Query(questionid)
		if err == nil {
			os, err = rowsToOptions(rows)
			return os, err
		}
	}
	return nil, err
}

func PostOption(db *sql.DB, qid int64, o *models.Option) (*Option, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var os *Option
	query, err := db.Prepare("INSERT INTO Opcion(texto, correcta, preguntaid) VALUES(?,?,?)")
	if err == nil {
		defer query.Close()
		sol, err := query.Exec(o.Texto, o.Correcta, qid)
		if err == nil {
			id, err := sol.LastInsertId()
			if err == nil {
				os, err = GetOptionQuestion(db, qid, id)
				return os, err
			}
		}
	}
	return nil, err
}

func GetOptionQuestion(db *sql.DB, questionid int64, optionindex int64) (*Option, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var os *Option
	query, err := db.Prepare("SELECT * FROM Opcion WHERE preguntaid=? AND indice=?")
	if err == nil {
		defer query.Close()
		rows, err := query.Query(questionid, optionindex)
		if err == nil {
			os, err = rowsToOption(rows)
			return os, err
		}
	}
	return nil, err
}

func PutOption(db *sql.DB, questionid int64, optionindex int64, o *models.Option) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	query, err := db.Prepare("UPDATE Opcion SET texto=?, correcta=? WHERE preguntaid=? AND optionindex=?")
	if err == nil {
		defer query.Close()
		_, err = query.Exec(o.Texto, o.Correcta, questionid, optionindex)
	}
	return err
}

func DeleteOption(db *sql.DB, questionid int64, optionindex int64) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	query, err := db.Prepare("DELETE FROM Opcion WHERE preguntaid=? AND optionindex=?")
	if err == nil {
		defer query.Close()
		_, err = query.Exec(questionid, optionindex)
	}
	return err
}
