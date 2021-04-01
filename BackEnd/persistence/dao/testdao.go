// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package dao acts as a Data Access Object for the Types
package dao

import (
	"database/sql"
	"errors"
	"uva-devtest/models"
	"uva-devtest/persistence/dbconnection"

	// Blank import of mysql driver
	_ "github.com/go-sql-driver/mysql"
)

func ToModelTest(t *Test) (*models.Test, error) {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		u, err := GetUserByID(db, t.Usuarioid)
		if err == nil {
			mt := &models.Test{
				AccesoPublico: t.AccesoPublico,
				Editable:      t.Editable,
				Description:   t.Description,
				ID:            t.ID,
				MaxSeconds:    t.MaxSeconds,
				Title:         t.Title,
				Username:      u.Username,
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
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.MaxSeconds, &t.AccesoPublico, &t.Editable, &t.Usuarioid)
		if err != nil {
			return tests, err
		}
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

func PostTest(db *sql.DB, username string, t *models.Test) error {
	if db == nil || t == nil {
		return errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err != nil || u == nil {
		return errors.New(errorResourceNotFound)
	}
	query, err := db.Prepare("INSERT INTO Test(title, description, maxSeconds, accesoPublico, editable, usuarioid) " +
		"VALUES (?,?,?,?,?,?)")

	if err != nil {
		return err
	}
	defer query.Close()
	_, err = query.Exec(t.Title, t.Description, t.MaxSeconds, t.AccesoPublico, t.Editable, u.ID)
	return err
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

// PublishedTests

func GetPTestsFromUser(db *sql.DB, username string) ([]*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err == nil {
		var ts []*Test
		query, err := db.Prepare("SELECT T.* FROM Test T JOIN InvitacionTestUsuario I ON T.id=I.testid WHERE I.usuarioid=?")

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

func GetPTestFromUser(db *sql.DB, username string, testid int64) (*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err == nil {
		var ts *Test
		query, err := db.Prepare("SELECT T.* FROM Test T JOIN InvitacionTestUsuario I ON T.id=I.testid WHERE I.usuarioid=? AND I.testid=?")
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

// AnsweredTests

func GetATestsFromUser(db *sql.DB, username string) ([]*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err == nil {
		var ts []*Test
		query, err := db.Prepare("SELECT T.* FROM Test T JOIN Pregunta P ON T.id=P.testid WHERE P.usuarioid=?")

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

func GetATestFromUser(db *sql.DB, username string, testid int64) (*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err == nil {
		var ts *Test
		query, err := db.Prepare("SELECT T.* FROM Test T JOIN Pregunta P ON T.id=P.testid WHERE P.usuarioid=? AND P.testid=?")
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
