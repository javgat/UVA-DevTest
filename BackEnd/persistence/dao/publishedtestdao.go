// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package dao acts as a Data Access Object for the Types
package dao

import (
	"database/sql"
	"errors"

	// Blank import of mysql driver
	_ "github.com/go-sql-driver/mysql"
)

func GetPublishedTests(db *sql.DB) ([]*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var ts []*Test
	query, err := db.Prepare("SELECT * FROM Test WHERE editable=0")
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

func GetPublishedTest(db *sql.DB, testid int64) (*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var ts *Test
	query, err := db.Prepare("SELECT * FROM Test WHERE id=? AND editable=0")
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
