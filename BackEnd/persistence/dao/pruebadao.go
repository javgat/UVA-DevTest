// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package dao acts as a Data Access Object for the Types
package dao

import (
	"database/sql"
	"errors"
	"uva-devtest/models"
)

// Transforms some sql.Rows into a slice(array) of Pruebas
// Param rows: Rows which contains database information returned
// Return []models.Prueba: Pruebas represented in rows
// Return error if any
func rowsToPruebas(rows *sql.Rows) ([]*Prueba, error) {
	var pruebas []*Prueba
	for rows.Next() {
		var t Prueba
		err := rows.Scan(&t.ID, &t.Preguntaid, &t.Entrada, &t.Salida, &t.Visible, &t.PostEntrega, &t.Valor)
		if err != nil {
			return pruebas, err
		}
		pruebas = append(pruebas, &t)
	}
	return pruebas, nil
}

// Transforms rows into a single prueba
// Param rows: Rows which contains database info of 1 prueba
// Return *models.Prueba: Prueba represented in rows
// Return error if something happens
func rowsToPrueba(rows *sql.Rows) (*Prueba, error) {
	var prueba *Prueba
	pruebas, err := rowsToPruebas(rows)
	if len(pruebas) >= 1 {
		prueba = pruebas[0]
	}
	return prueba, err
}

// Transforms some sql.Rows into a slice(array) of Pruebas, with more values
// Param rows: Rows which contains database information returned
// Return []models.Prueba: Pruebas represented in rows
// Return error if any
func rowsToPublishedPruebas(rows *sql.Rows) ([]*Prueba, error) {
	var pruebas []*Prueba
	for rows.Next() {
		var t Prueba
		var salidaReal, estado sql.NullString
		err := rows.Scan(&t.ID, &t.Preguntaid, &t.Entrada, &t.Salida, &t.Visible, &t.PostEntrega, &t.Valor, &salidaReal, &estado)
		if salidaReal.Valid {
			t.SalidaReal = salidaReal.String
		}
		if estado.Valid {
			t.Estado = estado.String
		}
		if err != nil {
			return pruebas, err
		}
		pruebas = append(pruebas, &t)
	}
	return pruebas, nil
}

func ToModelPrueba(p *Prueba) *models.Prueba {
	mp := &models.Prueba{
		ID:          p.ID,
		Preguntaid:  p.Preguntaid,
		Entrada:     p.Entrada,
		Salida:      p.Salida,
		Visible:     p.Visible,
		PostEntrega: p.PostEntrega,
		Valor:       p.Valor,
		SalidaReal:  p.SalidaReal,
		Estado:      p.Estado,
	}
	return mp
}

func ToModelPruebas(ps []*Prueba) []*models.Prueba {
	var mps = []*models.Prueba{}
	for _, itemCopy := range ps {
		mp := ToModelPrueba(itemCopy)
		mps = append(mps, mp)
	}
	return mps
}

func GetPruebas(db *sql.DB, questionid int64) ([]*Prueba, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var p []*Prueba
	query, err := db.Prepare("SELECT * FROM Prueba WHERE preguntaid=? ")
	if err == nil {
		defer query.Close()
		rows, err := query.Query(questionid)
		if err == nil {
			p, err = rowsToPruebas(rows)
			return p, err
		}
	}
	return nil, err
}

func PostPrueba(db *sql.DB, questionid int64, mp *models.Prueba) (*Prueba, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var p *Prueba
	visible := *mp.Visible && !*mp.PostEntrega
	query, err := db.Prepare("INSERT INTO Prueba(preguntaid, entrada, salida, visible, postEntrega, valor) VALUES(?,?,?,?,?,?) ")
	if err == nil {
		defer query.Close()
		var res sql.Result
		res, err = query.Exec(questionid, mp.Entrada, mp.Salida, visible, mp.PostEntrega, mp.Valor)
		if err == nil {
			var lid int64
			lid, err = res.LastInsertId()
			if err == nil {
				p = &Prueba{
					ID:          lid,
					Preguntaid:  questionid,
					Entrada:     mp.Entrada,
					Salida:      mp.Salida,
					Visible:     &visible,
					PostEntrega: mp.PostEntrega,
					Valor:       mp.Valor,
				}
				return p, err
			}
		}
	}
	return nil, err
}

func GetPrueba(db *sql.DB, questionid int64, pruebaid int64) (*Prueba, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var p *Prueba
	query, err := db.Prepare("SELECT * FROM Prueba WHERE preguntaid=? AND id=? ")
	if err == nil {
		defer query.Close()
		rows, err := query.Query(questionid, pruebaid)
		if err == nil {
			p, err = rowsToPrueba(rows)
			return p, err
		}
	}
	return nil, err
}

func PutPrueba(db *sql.DB, questionid int64, pruebaid int64, mp *models.Prueba) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	visible := *mp.Visible && !*mp.PostEntrega
	query, err := db.Prepare("UPDATE Prueba SET entrada=?, salida=?, visible=?, postEntrega=?, valor=? WHERE preguntaid=? AND id=? ")
	if err == nil {
		defer query.Close()
		_, err := query.Exec(mp.Entrada, mp.Salida, visible, mp.PostEntrega, mp.Valor, questionid, pruebaid)
		return err
	}
	return err
}

func DeletePrueba(db *sql.DB, questionid int64, pruebaid int64) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	query, err := db.Prepare("DELETE FROM Prueba WHERE preguntaid=? AND id=? ")
	if err == nil {
		defer query.Close()
		_, err := query.Exec(questionid, pruebaid)
		return err
	}
	return err
}

func GetVisiblePruebasQuestion(db *sql.DB, questionid int64) ([]*Prueba, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var p []*Prueba
	query, err := db.Prepare("SELECT * FROM Prueba WHERE preguntaid=? AND visible=1 ")
	if err == nil {
		defer query.Close()
		rows, err := query.Query(questionid)
		if err == nil {
			p, err = rowsToPruebas(rows)
			return p, err
		}
	}
	return nil, err
}

func GetPublishedPruebasQuestion(db *sql.DB, questionid int64, answerid int64) ([]*Prueba, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var p []*Prueba
	query, err := db.Prepare("SELECT P.*, E.salidaReal, E.estado FROM Prueba P JOIN Ejecucion E ON E.pruebaid=P.id WHERE P.preguntaid=?" +
		" AND E.respuestaExamenid=?")
	if err == nil {
		defer query.Close()
		rows, err := query.Query(questionid, answerid)
		if err == nil {
			p, err = rowsToPublishedPruebas(rows)
			return p, err
		}
	}
	return nil, err
}

func GetVisiblePublishedPruebasQuestion(db *sql.DB, questionid int64, answerid int64) ([]*Prueba, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var p []*Prueba
	query, err := db.Prepare("SELECT P.*, E.salidaReal, E.estado FROM Prueba P JOIN Ejecucion E ON E.pruebaid=P.id WHERE P.preguntaid=? " +
		" AND E.respuestaExamenid=? AND P.visible=1 ")
	if err == nil {
		defer query.Close()
		rows, err := query.Query(questionid, answerid)
		if err == nil {
			p, err = rowsToPublishedPruebas(rows)
			return p, err
		}
	}
	return nil, err
}
