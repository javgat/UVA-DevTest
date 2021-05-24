// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package dao acts as a Data Access Object for the Types
package dao

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"sync"
	"uva-devtest/models"
)

type customizedViewsSingle struct {
	cviews []*CustomizedView
}

var customizedViewsSingleton *customizedViewsSingle

var lockCVS = &sync.Mutex{}

func getCVSInstance(db *sql.DB) ([]*CustomizedView, error) {
	var err error
	if customizedViewsSingleton == nil {
		lockCVS.Lock()
		defer lockCVS.Unlock()
		if customizedViewsSingleton == nil {
			fmt.Println("Reading CustomizedViews from Database...")
			var trs []*CustomizedView
			trs, err = getCViews(db)
			customizedViewsSingleton = &customizedViewsSingle{
				cviews: trs,
			}
		}
	}
	return customizedViewsSingleton.cviews, err
}

func markAsInvalidCustomizedViewsSingleton() {
	lockCVS.Lock()
	defer lockCVS.Unlock()
	customizedViewsSingleton = nil
}

// Transforms some sql.Rows into a slice(array) of CustomizedViews
// Param rows: Rows which contains database information returned
// Return []models.CustomizedView: CustomizedViews represented in rows
// Return error if any
func rowsToCustomizedViews(rows *sql.Rows) ([]*CustomizedView, error) {
	var cvs []*CustomizedView
	for rows.Next() {
		var t CustomizedView
		err := rows.Scan(&t.RolBase, &t.MensajeInicio)
		if err != nil {
			return cvs, err
		}
		cvs = append(cvs, &t)
	}
	return cvs, nil
}

// Transforms rows into a single CustomizedView
// Param rows: Rows which contains database info of 1 CustomizedView
// Return *models.CustomizedView: CustomizedView represented in rows
// Return error if something happens
func rowsToCustomizedView(rows *sql.Rows) (*CustomizedView, error) {
	var cv *CustomizedView
	cvs, err := rowsToCustomizedViews(rows)
	if len(cvs) >= 1 {
		cv = cvs[0]
	}
	return cv, err
}

func ToModelCustomizedView(cv *CustomizedView) *models.CustomizedView {
	mcv := &models.CustomizedView{
		RolBase:       cv.RolBase,
		MensajeInicio: cv.MensajeInicio,
	}
	return mcv
}

func ToModelCustomizedViews(cvs []*CustomizedView) []*models.CustomizedView {
	var mcvs = []*models.CustomizedView{}
	for _, itemCopy := range cvs {
		mcv := ToModelCustomizedView(itemCopy)
		mcvs = append(mcvs, mcv)
	}
	return mcvs
}

func GetCViews(db *sql.DB) ([]*CustomizedView, error) {
	return getCVSInstance(db)
}

func getCViews(db *sql.DB) ([]*CustomizedView, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var cvs []*CustomizedView
	query, err := db.Prepare("SELECT * FROM VistaPersonalizada")
	if err == nil {
		defer query.Close()
		var rows *sql.Rows
		rows, err = query.Query()
		if err == nil {
			cvs, err = rowsToCustomizedViews(rows)
			return cvs, err
		}
	}
	return nil, err
}

func GetCView(db *sql.DB, rolBase string) (*CustomizedView, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	cvs, err := getCVSInstance(db)
	if err == nil {
		for _, cv := range cvs {
			if strings.EqualFold(*cv.RolBase, rolBase) {
				return cv, nil
			}
		}
		return nil, nil
	}
	return nil, err
}

func PutCView(db *sql.DB, rolBase string, newView *models.CustomizedView) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	query, err := db.Prepare("UPDATE VistaPersonalizada SET mensajeInicio=? WHERE rolBase=?")
	if err == nil {
		defer query.Close()
		_, err = query.Exec(newView.MensajeInicio, rolBase)
	}
	markAsInvalidCustomizedViewsSingleton()
	return err
}
