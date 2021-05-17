// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package dao acts as a Data Access Object for the Types
package dao

import (
	"database/sql"
	"errors"
	"log"
	"time"
	"uva-devtest/models"
	"uva-devtest/persistence/dbconnection"

	// Blank import of mysql driver
	"github.com/go-openapi/strfmt"
	_ "github.com/go-sql-driver/mysql"
)

func ToModelTest(t *Test) (*models.Test, error) {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		u, err := GetUserByID(db, t.Usuarioid)
		if err == nil {
			mt := &models.Test{
				AccesoPublico:            t.AccesoPublico,
				AccesoPublicoNoPublicado: t.AccesoPublicoNoPublicado,
				Editable:                 t.Editable,
				Description:              t.Description,
				ID:                       t.ID,
				MaxMinutes:               t.MaxMinutes,
				Title:                    t.Title,
				Username:                 u.Username,
				HoraCreacion:             t.HoraCreacion,
				OriginalTestID:           t.OriginalTestID,
				NotaMaxima:               t.NotaMaxima,
				AutoCorrect:              t.AutoCorrect,
				Visibilidad:              t.Visibilidad,
				CantidadFavoritos:        t.CantidadFavoritos,
			}
			return mt, nil
		}
	}
	return nil, errors.New(errorResourceNotFound)
}

func ToModelTests(ts []*Test) ([]*models.Test, error) {
	var mts = []*models.Test{}
	for _, itemCopy := range ts {
		mt, err := ToModelTest(itemCopy)
		if err != nil {
			return nil, err
		}
		mts = append(mts, mt)
	}
	return mts, nil
}

// Transforms some sql.Rows into a slice(array) of tests
// Param rows: Rows which contains database information returned
// Return []models.Test: Tests represented in rows
// Return error if any
func rowsToTests(rows *sql.Rows) ([]*Test, error) {
	var tests []*Test
	for rows.Next() {
		var t Test
		var timeNull sql.NullTime
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.MaxMinutes, &t.AccesoPublico, &t.Editable, &t.Usuarioid,
			&t.AccesoPublicoNoPublicado, &timeNull, &t.OriginalTestID, &t.NotaMaxima, &t.AutoCorrect, &t.Visibilidad,
			&t.CantidadFavoritos)
		if err != nil {
			return tests, err
		}
		if !timeNull.Valid {
			return tests, errors.New("tiempo SQL no valido")
		}
		t.HoraCreacion = strfmt.DateTime(timeNull.Time)
		tests = append(tests, &t)
	}
	return tests, nil
}

// Transforms rows into a single test
// Param rows: Rows which contains database info of 1 test
// Return *models.Test: Test represented in rows
// Return error if something happens
func rowsToTest(rows *sql.Rows) (*Test, error) {
	var test *Test
	tests, err := rowsToTests(rows)
	if len(tests) >= 1 {
		test = tests[0]
	}
	return test, err
}

func addFiltersToQueryTest(hayWhere bool, initQuery string, tags [][]string, likeTitle *string, orderBy *string,
	limit *int64, offset *int64) string {
	return AddFiltersToQuery(hayWhere, initQuery, tags, likeTitle, orderBy, limit, offset, "id", "testid", "TestEtiqueta",
		"title", "cantidadFavoritos", "maxMinutes")
}

func addFiltersToQueryTestLong(hayWhere bool, initQuery string, tags [][]string, likeTitle *string, orderBy *string,
	limit *int64, offset *int64) string {
	return AddFiltersToQuery(hayWhere, initQuery, tags, likeTitle, orderBy, limit, offset, "T.id", "testid", "TestEtiqueta",
		"title", "T.cantidadFavoritos", "T.maxMinutes")
}

func GetAllTests(db *sql.DB) ([]*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var ts []*Test
	query, err := db.Prepare("SELECT * FROM Test")
	if err == nil {
		defer query.Close()
		rows, err := query.Query()
		if err == nil {
			ts, err = rowsToTests(rows)
			return ts, err
		}
	}
	return nil, err
}

func GetAllEditTests(db *sql.DB) ([]*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var ts []*Test
	query, err := db.Prepare("SELECT * FROM Test WHERE editable=1")
	if err == nil {
		defer query.Close()
		rows, err := query.Query()
		if err == nil {
			ts, err = rowsToTests(rows)
			return ts, err
		}
	}
	return nil, err
}

func GetPublicTests(db *sql.DB) ([]*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var ts []*Test
	query, err := db.Prepare("SELECT * FROM Test WHERE (editable=1 AND accesoPublicoNoPublicado=1) OR (editable=0 AND accesoPublico=1)")
	if err == nil {
		defer query.Close()
		rows, err := query.Query()
		if err == nil {
			ts, err = rowsToTests(rows)
			return ts, err
		}
	}
	return nil, err
}

func GetPublicEditTests(db *sql.DB, tags [][]string, likeTitle *string, orderBy *string,
	limit *int64, offset *int64) ([]*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var ts []*Test
	stPrepare := "SELECT * FROM Test WHERE editable=1 AND accesoPublicoNoPublicado=1 "
	stPrepare = addFiltersToQueryTest(true, stPrepare, tags, likeTitle, orderBy, limit, offset)
	query, err := db.Prepare(stPrepare)
	if err == nil {
		defer query.Close()
		interfaceParams := FilterParamsSlicesToInterfaceArr(tags, likeTitle)
		rows, err := query.Query(interfaceParams...)
		if err == nil {
			ts, err = rowsToTests(rows)
			return ts, err
		}
	}
	return nil, err
}

func GetTest(db *sql.DB, testid int64) (*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var ts *Test
	query, err := db.Prepare("SELECT * FROM Test WHERE id=?")
	if err == nil {
		defer query.Close()
		rows, err := query.Query(testid)
		if err == nil {
			ts, err = rowsToTest(rows)
			return ts, err
		}
	}
	return nil, err
}

func PutTest(db *sql.DB, testid int64, t *models.Test) error {
	if db == nil || t == nil {
		return errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, *t.Username)
	if err != nil || u == nil {
		return errors.New(errorResourceNotFound)
	}
	query, err := db.Prepare("UPDATE Test SET title=?, description=?, maxMinutes=?, accesoPublico=?, usuarioid=?, " +
		"accesoPublicoNoPublicado=?, notaMaxima=?, autoCorrect=?, visibilidad=? WHERE editable=1 AND id=?")

	if err != nil {
		return err
	}
	defer query.Close()
	_, err = query.Exec(t.Title, t.Description, t.MaxMinutes, *t.AccesoPublico, u.ID, *t.AccesoPublicoNoPublicado, t.NotaMaxima,
		t.AutoCorrect, t.Visibilidad, testid)
	return err
}

func DeleteTest(db *sql.DB, testid int64) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	query, err := db.Prepare("DELETE FROM Test WHERE id=?")
	if err == nil {
		defer query.Close()
		_, err = query.Exec(testid)
	}
	return err
}

func GetPublicEditTestsFromUser(db *sql.DB, username string, tags [][]string, likeTitle *string, orderBy *string,
	limit *int64, offset *int64) ([]*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err == nil {
		var ts []*Test
		stPrepare := "SELECT * FROM Test WHERE usuarioid=? AND accesoPublicoNoPublicado=1"
		stPrepare = addFiltersToQueryTest(true, stPrepare, tags, likeTitle, orderBy, limit, offset)
		query, err := db.Prepare(stPrepare)
		if err == nil {
			defer query.Close()
			interfaceParams := FilterParamsSlicesToInterfaceArr(tags, likeTitle)
			var paramsSlice []interface{}
			paramsSlice = append(paramsSlice, u.ID)
			interfaceParams = append(paramsSlice, interfaceParams...)
			rows, err := query.Query(interfaceParams...)
			if err == nil {
				ts, err = rowsToTests(rows)
				return ts, err
			}
		}
	}
	return nil, err
}

func GetEditTestsFromUser(db *sql.DB, username string, tags [][]string, likeTitle *string, orderBy *string,
	limit *int64, offset *int64) ([]*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err == nil {
		var ts []*Test
		stPrepare := "SELECT * FROM Test WHERE usuarioid=? AND editable=1"
		stPrepare = addFiltersToQueryTest(true, stPrepare, tags, likeTitle, orderBy, limit, offset)
		query, err := db.Prepare(stPrepare)
		if err == nil {
			defer query.Close()
			interfaceParams := FilterParamsSlicesToInterfaceArr(tags, likeTitle)
			var paramsSlice []interface{}
			paramsSlice = append(paramsSlice, u.ID)
			interfaceParams = append(paramsSlice, interfaceParams...)
			rows, err := query.Query(interfaceParams...)
			if err == nil {
				ts, err = rowsToTests(rows)
				return ts, err
			}
		}
	}
	return nil, err
}

func GetTestsFromUser(db *sql.DB, username string) ([]*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err == nil {
		var ts []*Test
		query, err := db.Prepare("SELECT * FROM Test WHERE usuarioid=?")
		if err == nil {
			defer query.Close()
			rows, err := query.Query(u.ID)
			if err == nil {
				ts, err = rowsToTests(rows)
				return ts, err
			}
		}
	}
	return nil, err
}

func PostTest(db *sql.DB, username string, t *models.Test, horaCreacion time.Time) (*models.Test, error) {
	if db == nil || t == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err != nil || u == nil {
		return nil, errors.New(errorResourceNotFound)
	}
	query, err := db.Prepare("INSERT INTO Test(title, description, maxMinutes, accesoPublico, editable, usuarioid, " +
		" accesoPublicoNoPublicado, horaCreacion, origenTestid, notaMaxima, autoCorrect, visibilidad) " +
		"VALUES (?,?,?,?,?,?,?,?,?,?,?,?)")

	if err != nil {
		return nil, err
	}
	defer query.Close()
	origID := &t.OriginalTestID
	if *t.OriginalTestID == -1 {
		origID = nil
	}
	sol, err := query.Exec(t.Title, t.Description, t.MaxMinutes, t.AccesoPublico, t.Editable, u.ID,
		t.AccesoPublicoNoPublicado, horaCreacion, origID, t.NotaMaxima, t.AutoCorrect, t.Visibilidad)
	if err == nil {
		ts := t
		ts.ID, err = sol.LastInsertId()
		return ts, err
	}
	return nil, err
}

func GetTestFromUser(db *sql.DB, username string, testid int64) (*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err == nil {
		var ts *Test
		query, err := db.Prepare("SELECT * FROM Test WHERE usuarioid=? AND id=?")
		if err == nil {
			defer query.Close()
			rows, err := query.Query(u.ID, testid)
			if err == nil {
				ts, err = rowsToTest(rows)
				return ts, err
			}
		}
	}
	return nil, err
}

func GetTestsFromTeam(db *sql.DB, teamname string, tags [][]string, likeTitle *string, orderBy *string,
	limit *int64, offset *int64) ([]*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetTeam(db, teamname)
	if err == nil {
		var t []*Test
		stPrepare := "SELECT T.* FROM Test T JOIN GestionTestEquipo G ON T.id=G.testid WHERE G.equipoid=? "
		stPrepare = addFiltersToQueryTest(true, stPrepare, tags, likeTitle, orderBy, limit, offset)
		query, err := db.Prepare(stPrepare)
		if err == nil {
			defer query.Close()
			interfaceParams := FilterParamsSlicesToInterfaceArr(tags, likeTitle)
			var paramsSlice []interface{}
			paramsSlice = append(paramsSlice, u.ID)
			interfaceParams = append(paramsSlice, interfaceParams...)
			rows, err := query.Query(interfaceParams...)
			if err == nil {
				t, err = rowsToTests(rows)
				return t, err
			}
		}
	}
	return nil, err
}

func GetTestFromTeam(db *sql.DB, teamname string, testid int64) (*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetTeam(db, teamname)
	if err == nil {
		var t *Test
		query, err := db.Prepare("SELECT T.* FROM Test T JOIN GestionTestEquipo G ON T.id=G.testid WHERE G.equipoid=? AND T.id=?")
		if err == nil {
			defer query.Close()
			rows, err := query.Query(u.ID, testid)
			if err == nil {
				t, err = rowsToTest(rows)
				return t, err
			}
		}
	}
	return nil, err
}

func GetPTestsFromTeam(db *sql.DB, teamname string) ([]*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetTeam(db, teamname)
	if err == nil {
		var t []*Test
		query, err := db.Prepare("SELECT T.* FROM Test T JOIN GestionTestEquipo G ON T.id=G.testid WHERE G.equipoid=? AND T.editable=0")
		if err == nil {
			defer query.Close()
			rows, err := query.Query(u.ID)
			if err == nil {
				t, err = rowsToTests(rows)
				return t, err
			}
		}
	}
	return nil, err
}

func GetPTestFromTeam(db *sql.DB, teamname string, testid int64) (*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetTeam(db, teamname)
	if err == nil {
		var t *Test
		query, err := db.Prepare("SELECT T.* FROM Test T JOIN GestionTestEquipo G ON T.id=G.testid WHERE G.equipoid=? AND T.id=? AND T.editable=0")
		if err == nil {
			defer query.Close()
			rows, err := query.Query(u.ID, testid)
			if err == nil {
				t, err = rowsToTest(rows)
				return t, err
			}
		}
	}
	return nil, err
}

func GetSharedEditTestsFromUser(db *sql.DB, username string, tags [][]string, likeTitle *string, orderBy *string,
	limit *int64, offset *int64) ([]*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var t []*Test
	stPrepare := "SELECT DISTINCT T.* FROM Test T JOIN GestionTestEquipo G ON T.id=G.testid JOIN EquipoUsuario E ON G.equipoid=E.equipoid JOIN Usuario U ON U.id=E.usuarioid WHERE U.username=? AND T.editable=1"
	stPrepare = addFiltersToQueryTestLong(true, stPrepare, tags, likeTitle, orderBy, limit, offset)
	query, err := db.Prepare(stPrepare)
	if err == nil {
		defer query.Close()
		interfaceParams := FilterParamsSlicesToInterfaceArr(tags, likeTitle)
		var paramsSlice []interface{}
		paramsSlice = append(paramsSlice, username)
		interfaceParams = append(paramsSlice, interfaceParams...)
		rows, err := query.Query(interfaceParams...)
		if err == nil {
			t, err = rowsToTests(rows)
			return t, err
		}
	}
	return nil, err
}

func GetSharedPublishedTestsFromUser(db *sql.DB, username string, tags [][]string, likeTitle *string, orderBy *string,
	limit *int64, offset *int64) ([]*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var t []*Test
	stPrepare := "SELECT DISTINCT T.* FROM Test T JOIN GestionTestEquipo G ON T.id=G.testid JOIN EquipoUsuario E ON G.equipoid=E.equipoid JOIN Usuario U ON U.id=E.usuarioid WHERE U.username=? AND T.editable=0"
	stPrepare = addFiltersToQueryTestLong(true, stPrepare, tags, likeTitle, orderBy, limit, offset)
	query, err := db.Prepare(stPrepare)
	if err == nil {
		defer query.Close()
		interfaceParams := FilterParamsSlicesToInterfaceArr(tags, likeTitle)
		var paramsSlice []interface{}
		paramsSlice = append(paramsSlice, username)
		interfaceParams = append(paramsSlice, interfaceParams...)
		rows, err := query.Query(interfaceParams...)
		if err == nil {
			t, err = rowsToTests(rows)
			return t, err
		}
	}
	return nil, err
}

func GetSharedTestsFromUser(db *sql.DB, username string) ([]*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var t []*Test
	query, err := db.Prepare("SELECT DISTINCT T.* FROM Test T JOIN GestionTestEquipo G ON T.id=G.testid JOIN EquipoUsuario E ON G.equipoid=E.equipoid JOIN Usuario U ON U.id=E.usuarioid WHERE U.username=?")
	if err == nil {
		defer query.Close()
		rows, err := query.Query(username)
		if err == nil {
			t, err = rowsToTests(rows)
			return t, err
		}
	}
	return nil, err
}

func GetSharedTestFromUser(db *sql.DB, username string, testid int64) (*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var t *Test
	query, err := db.Prepare("SELECT DISTINCT T.* FROM Test T JOIN GestionTestEquipo G ON T.id=G.testid JOIN EquipoUsuario E ON G.equipoid=E.equipoid JOIN Usuario U ON U.id=E.usuarioid WHERE U.username=? AND T.id=?")
	if err == nil {
		defer query.Close()
		var rows *sql.Rows
		rows, err = query.Query(username, testid)
		if err == nil {
			t, err = rowsToTest(rows)
			return t, err
		}
	}
	return nil, err
}

func GetTestsFromTag(db *sql.DB, nombre string) ([]*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var t []*Test
	query, err := db.Prepare("SELECT T.* FROM Test T JOIN TestEtiqueta E ON T.id=E.testid WHERE E.etiquetanombre=?")
	if err == nil {
		defer query.Close()
		var rows *sql.Rows
		rows, err = query.Query(nombre)
		if err == nil {
			t, err = rowsToTests(rows)
			return t, err
		}
	}
	return nil, err
}

func GetEditTestsFromTag(db *sql.DB, nombre string) ([]*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var t []*Test
	query, err := db.Prepare("SELECT T.* FROM Test T JOIN TestEtiqueta E ON T.id=E.testid WHERE E.etiquetanombre=? AND T.editable=1")
	if err == nil {
		defer query.Close()
		var rows *sql.Rows
		rows, err = query.Query(nombre)
		if err == nil {
			t, err = rowsToTests(rows)
			return t, err
		}
	}
	return nil, err
}

func GetFavoriteEditTests(db *sql.DB, username string, tags [][]string, likeTitle *string, orderBy *string,
	limit *int64, offset *int64) ([]*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var qs []*Test
	stPrepare := "SELECT DISTINCT T.* FROM Test T LEFT JOIN GestionTestEquipo E ON T.id=E.testid LEFT JOIN EquipoUsuario U ON U.equipoid=E.equipoid " +
		" LEFT JOIN Usuario V ON V.id=U.usuarioid LEFT JOIN Usuario W ON W.id=T.usuarioid JOIN TestFavorito F ON T.id=F.testid JOIN Usuario Y ON Y.id=F.usuarioid " +
		" WHERE Y.username=? AND T.editable=1 AND (V.username=? OR T.accesoPublicoNoPublicado=1 OR W.username=?) "
	stPrepare = addFiltersToQueryTestLong(true, stPrepare, tags, likeTitle, orderBy, limit, offset)
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
			qs, err = rowsToTests(rows)
			return qs, err
		}
	} else {
		log.Print(err)
	}
	return nil, err
}

func GetFavoriteTests(db *sql.DB, username string, tags [][]string, likeTitle *string, orderBy *string,
	limit *int64, offset *int64) ([]*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var qs []*Test
	stPrepare := "SELECT DISTINCT T.* FROM Test T LEFT JOIN GestionTestEquipo E ON T.id=E.testid LEFT JOIN EquipoUsuario U ON U.equipoid=E.equipoid " +
		" LEFT JOIN Usuario V ON V.id=U.usuarioid LEFT JOIN Usuario W ON W.id=T.usuarioid JOIN TestFavorito F ON T.id=F.testid JOIN Usuario Y ON Y.id=F.usuarioid " +
		" WHERE Y.username=? AND (V.username=? OR T.accesoPublicoNoPublicado=1 OR (T.editable=0 AND T.accesoPublico=1) OR W.username=?) "
	stPrepare = addFiltersToQueryTestLong(true, stPrepare, tags, likeTitle, orderBy, limit, offset)
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
			qs, err = rowsToTests(rows)
			return qs, err
		}
	} else {
		log.Print(err)
	}
	return nil, err
}

func GetFavoriteTest(db *sql.DB, username string, testid int64) (*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var qs *Test
	stPrepare := "SELECT DISTINCT T.* FROM Test T LEFT JOIN GestionTestEquipo E ON T.id=E.testid LEFT JOIN EquipoUsuario U ON U.equipoid=E.equipoid " +
		" LEFT JOIN Usuario V ON V.id=U.usuarioid LEFT JOIN Usuario W ON W.id=T.usuarioid JOIN TestFavorito F ON T.id=F.testid JOIN Usuario Y ON Y.id=F.usuarioid " +
		" WHERE Y.username=? AND (V.username=? OR T.accesoPublicoNoPublicado=1 OR (T.editable=0 AND T.accesoPublico=1) OR W.username=?) AND T.id=?"
	query, err := db.Prepare(stPrepare)
	if err == nil {
		defer query.Close()
		rows, err := query.Query(username, username, username, testid)
		if err == nil {
			qs, err = rowsToTest(rows)
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

func AddFavoriteTest(db *sql.DB, username string, testid int64) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err == nil && u != nil {
		stPrepare := "INSERT INTO TestFavorito(usuarioid, testid) VALUES(?,?)"
		var query *sql.Stmt
		query, err = db.Prepare(stPrepare)
		if err == nil {
			defer query.Close()
			_, err := query.Exec(u.ID, testid)
			return err
		} else {
			log.Print(err)
		}
	} else if err == nil {
		err = errors.New(errorResourceNotFound)
	}
	return err
}

func RemoveFavoriteTest(db *sql.DB, username string, testid int64) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err == nil && u != nil {
		stPrepare := "DELETE FROM TestFavorito WHERE usuarioid=? AND testid=?"
		var query *sql.Stmt
		query, err = db.Prepare(stPrepare)
		if err == nil {
			defer query.Close()
			_, err := query.Exec(u.ID, testid)
			return err
		} else {
			log.Print(err)
		}
	} else if err == nil {
		err = errors.New(errorResourceNotFound)
	}
	return err
}

func GetTestsByOrigenTestid(db *sql.DB, testid int64) ([]*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var ts []*Test
	stPrepare := "SELECT * FROM Test WHERE origenTestid=? ORDER BY id DESC"
	query, err := db.Prepare(stPrepare)
	if err == nil {
		defer query.Close()
		rows, err := query.Query(testid)
		if err == nil {
			ts, err = rowsToTests(rows)
			if ts == nil {
				return nil, nil
			}
			return ts, err
		}
	} else {
		log.Print(err)
	}
	return nil, err
}
