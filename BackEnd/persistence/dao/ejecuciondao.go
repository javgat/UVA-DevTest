// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package dao acts as a Data Access Object for the Types
package dao

import (
	"database/sql"
	"errors"
)

func DeleteEjecucion(db *sql.DB, pruebaid int64, answerid int64, questionid int64) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	query, err := db.Prepare("DELETE FROM Ejecucion WHERE pruebaid=? AND respuestaExamenid=? AND preguntaid=? ")
	if err == nil {
		defer query.Close()
		_, err = query.Exec(pruebaid, answerid, questionid)
	}
	return err
}

func SaveEjecucion(db *sql.DB, ej *Ejecucion) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	DeleteEjecucion(db, *ej.Pruebaid, *ej.RespuestaExamenid, *ej.Preguntaid)
	query, err := db.Prepare("INSERT INTO Ejecucion(pruebaid, respuestaExamenid, preguntaid, estado, salidaReal) VALUES(?,?,?,?,?)")
	if err == nil {
		defer query.Close()
		_, err = query.Exec(ej.Pruebaid, ej.RespuestaExamenid, ej.Preguntaid, ej.Estado, ej.SalidaReal)
	}
	return err
}
