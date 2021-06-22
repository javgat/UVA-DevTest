// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package dao acts as a Data Access Object for the Types
package dao

import (
	"database/sql"
	"errors"
	"strings"
	"uva-devtest/models"
)

type EjecucionTesting struct {
	Estado   *string
	Cantidad *int64
}

// Transforms rows into a single Testing
// Param rows: Rows which contains database info of 1 Testing
// Return *models.Testing: Testing represented in rows
// Return error if something happens
func rowsToTesting(rows *sql.Rows) (*Testing, error) {
	var t *Testing
	*t.PruebasTotales = 0
	//var ejs []*EjecucionTesting
	for rows.Next() {
		var ej EjecucionTesting
		err := rows.Scan(&ej.Estado, &ej.Cantidad)
		if err != nil {
			return t, err
		}
		if strings.EqualFold(*ej.Estado, EstadoEjecucionCorrecto) {
			*t.PruebasSuperadas = *ej.Cantidad
		}
		*t.PruebasTotales += *ej.Cantidad
		//ejs = append(ejs, &ej)
	}
	return t, nil
}

func ToModelTesting(t *Testing) *models.Testing {
	mt := &models.Testing{
		PruebasSuperadas: t.PruebasSuperadas,
		PruebasTotales:   t.PruebasTotales,
	}
	return mt
}

func GetPreTesting(db *sql.DB, answerid int64, questionid int64) (*Testing, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var t *Testing
	query, err := db.Prepare("SELECT E.estado, COUNT(*) FROM Ejecucion E JOIN Prueba P ON P.id=E.pruebaid WHERE E.respuestaExamenid=? " +
		" AND E.preguntaid=? AND P.postEntrega=0 GROUP BY E.estado")
	if err == nil {
		defer query.Close()
		rows, err := query.Query(answerid, questionid)
		if err == nil {
			t, err = rowsToTesting(rows)
			return t, err
		}
	}
	return nil, err
}

func GetFullTesting(db *sql.DB, answerid int64, questionid int64) (*Testing, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var t *Testing
	query, err := db.Prepare("SELECT E.estado, COUNT(*) FROM Ejecucion E JOIN Prueba P ON P.id=E.pruebaid WHERE E.respuestaExamenid=? " +
		" AND E.preguntaid=? GROUP BY E.estado")
	if err == nil {
		defer query.Close()
		rows, err := query.Query(answerid, questionid)
		if err == nil {
			t, err = rowsToTesting(rows)
			return t, err
		}
	}
	return nil, err
}
