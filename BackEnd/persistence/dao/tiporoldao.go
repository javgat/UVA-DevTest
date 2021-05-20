// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package dao acts as a Data Access Object for the Types
package dao

import (
	"database/sql"
	"errors"
	"fmt"
	"sync"
	"uva-devtest/models"
)

type tipoRolesSingle struct {
	roles []*TipoRol
}

var tipoRolesSingleton *tipoRolesSingle

var lockTRS = &sync.Mutex{}

func getTRSInstance(db *sql.DB) ([]*TipoRol, error) {
	var err error
	if tipoRolesSingleton == nil {
		lockTRS.Lock()
		defer lockTRS.Unlock()
		if tipoRolesSingleton == nil {
			fmt.Println("Reading TipoRoles from Database...")
			var trs []*TipoRol
			trs, err = getTipoRoles(db)
			tipoRolesSingleton = &tipoRolesSingle{
				roles: trs,
			}
		}
	}
	return tipoRolesSingleton.roles, err
}

func markAsInvalidTipoRolesSingleton() {
	lockTRS.Lock()
	defer lockTRS.Unlock()
	tipoRolesSingleton = nil
}

// Transforms some sql.Rows into a slice(array) of TipoRol
// Param rows: Rows which contains database information returned
// Return []models.TipoRol: Users represented in rows
// Return error if any
func rowsToTipoRoles(rows *sql.Rows) ([]*TipoRol, error) {
	var trs []*TipoRol
	for rows.Next() {
		var tr TipoRol
		err := rows.Scan(&tr.ID, &tr.RolBase, &tr.Nombre, &tr.Prioridad, &tr.VerPTests, &tr.VerETests, &tr.VerEQuestions,
			&tr.VerPQuestions, &tr.VerAnswers, &tr.ChangeRoles, &tr.TenerTeams, &tr.TenerEQuestions, &tr.TenerETests, &tr.TenerPTests,
			&tr.AdminPTests, &tr.AdminETests, &tr.AdminEQuestions, &tr.AdminAnswers, &tr.AdminUsers, &tr.AdminTeams, &tr.AdminConfiguration,
			&tr.AdminPermissions, &tr.TipoInicial)
		if err != nil {
			return trs, err
		}
		trs = append(trs, &tr)
	}
	return trs, nil
}

// Transforms rows into a single TipoRol
// Param rows: Rows which contains database info of 1 TipoRol
// Return *models.TipoRol: TipoRol represented in rows
// Return error if something happens
func rowsToTipoRol(rows *sql.Rows) (*TipoRol, error) {
	var tr *TipoRol
	trs, err := rowsToTipoRoles(rows)
	if len(trs) >= 1 {
		tr = trs[0]
	}
	return tr, err
}

// ToModelTipoRol converts a dao.TipoRol into a models.TipoRol
// Param t: dao.TipoRol to convert
func ToModelTipoRol(t *TipoRol) *models.TipoRol {
	mt := &models.TipoRol{
		ID:                 t.ID,
		RolBase:            t.RolBase,
		Nombre:             t.Nombre,
		Prioridad:          t.Prioridad,
		VerPTests:          t.VerPTests,
		VerETests:          t.VerETests,
		VerEQuestions:      t.VerEQuestions,
		VerPQuestions:      t.VerPQuestions,
		VerAnswers:         t.VerAnswers,
		ChangeRoles:        t.ChangeRoles,
		TenerTeams:         t.TenerTeams,
		TenerEQuestions:    t.TenerEQuestions,
		TenerETests:        t.TenerETests,
		TenerPTests:        t.TenerPTests,
		AdminPTests:        t.AdminPTests,
		AdminETests:        t.AdminETests,
		AdminEQuestions:    t.AdminEQuestions,
		AdminAnswers:       t.AdminAnswers,
		AdminUsers:         t.AdminUsers,
		AdminTeams:         t.AdminTeams,
		AdminConfiguration: t.AdminConfiguration,
		AdminPermissions:   t.AdminPermissions,
		TipoInicial:        t.TipoInicial,
	}
	return mt
}

func GetTipoRolByID(db *sql.DB, id *int64) (*TipoRol, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	trs, err := getTRSInstance(db)
	if err == nil {
		for _, tr := range trs {
			if *tr.ID == *id {
				return tr, nil
			}
		}
		return nil, nil
	}
	return nil, err
}

func getTipoRoles(db *sql.DB) ([]*TipoRol, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var trs []*TipoRol
	query, err := db.Prepare("SELECT * FROM TipoRol")
	if err != nil {
		return trs, err
	}
	defer query.Close()
	rows, err := query.Query()
	if err == nil {
		trs, err = rowsToTipoRoles(rows)
	}
	return trs, err
}

func GetTipoRoles(db *sql.DB) ([]*TipoRol, error) {
	return getTRSInstance(db)
}

func getDefaultTipoRolStudent() *TipoRol {
	var uno int64 = 1
	var tipoEst string = models.TipoRolRolBaseEstudiante
	var btrue bool = true
	var bfalse bool = false
	return &TipoRol{
		ID:                 &uno,
		RolBase:            &tipoEst,
		Nombre:             &tipoEst,
		Prioridad:          &uno,
		VerPTests:          &btrue,
		VerETests:          &bfalse,
		VerEQuestions:      &bfalse,
		VerPQuestions:      &bfalse,
		VerAnswers:         &bfalse,
		ChangeRoles:        &bfalse,
		TenerTeams:         &bfalse,
		TenerEQuestions:    &bfalse,
		TenerETests:        &bfalse,
		TenerPTests:        &bfalse,
		AdminPTests:        &bfalse,
		AdminETests:        &bfalse,
		AdminEQuestions:    &bfalse,
		AdminAnswers:       &bfalse,
		AdminUsers:         &bfalse,
		AdminTeams:         &bfalse,
		AdminConfiguration: &bfalse,
		AdminPermissions:   &bfalse,
		TipoInicial:        &btrue,
	}
}

func GetTipoRolNewUser(db *sql.DB) (*TipoRol, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	trs, err := getTRSInstance(db)
	if err == nil {
		for _, tr := range trs {
			if *tr.TipoInicial {
				return tr, nil
			}
		}
		trdef := getDefaultTipoRolStudent()
		_, err = PostTipoRol(db, ToModelTipoRol(trdef))
		if err == nil {
			return trdef, nil
		}
	}
	return nil, err
}

func PostTipoRol(db *sql.DB, t *models.TipoRol) (*models.TipoRol, error) {
	if db == nil || t == nil {
		return nil, errors.New(errorDBNil)
	}
	query, err := db.Prepare("INSERT INTO TipoRol(rolBase, nombre, prioridad, verPTests, verETests, verEQuestions, verPQuestions, " +
		" verAnswers, changeRoles, tenerTeams, tenerEQuestions, tenerETests, tenerPTests, adminPTests, adminETests, adminEQuestions, " +
		" adminAnswers, adminUsers, adminTeams, adminConfiguration, adminPermissions, tipoInicial) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")

	if err == nil {
		defer query.Close()
		var res sql.Result
		res, err = query.Exec(t.RolBase, t.Nombre, t.Prioridad, t.VerPTests, t.VerETests, t.VerEQuestions, t.VerPQuestions,
			t.VerAnswers, t.ChangeRoles, t.TenerTeams, t.TenerEQuestions, t.TenerETests, t.TenerPTests, t.AdminPTests, t.AdminETests, t.AdminEQuestions,
			t.AdminAnswers, t.AdminUsers, t.AdminTeams, t.AdminConfiguration, t.AdminPermissions, t.TipoInicial)
		if err == nil {
			var lastInt int64
			lastInt, err = res.LastInsertId()
			if err == nil {
				newTr := t
				newTr.ID = &lastInt
				markAsInvalidTipoRolesSingleton()
				return newTr, err
			}
		}
	}
	return nil, err
}

func PutTipoRol(db *sql.DB, id *int64, t *TipoRol) error {
	if db == nil || t == nil || id == nil {
		return errors.New(errorDBNil)
	}
	query, err := db.Prepare("UPDATE TipoRol SET rolBase=?, nombre=?, prioridad=?, verPTests=?, verETests=?, verEQuestions=?, verPQuestions=?, " +
		" verAnswers=?, changeRoles=?, tenerTeams=?, tenerEQuestions=?, tenerETests=?, tenerPTests=?, adminPTests=?, adminETests=?, adminEQuestions=?, " +
		" adminAnswers=?, adminUsers=?, adminTeams=?, adminConfiguration=?, adminPermissions=?, tipoInicial=? WHERE id=?")

	if err == nil {
		defer query.Close()
		_, err = query.Exec(t.RolBase, t.Nombre, t.Prioridad, t.VerPTests, t.VerETests, t.VerEQuestions, t.VerPQuestions,
			t.VerAnswers, t.ChangeRoles, t.TenerTeams, t.TenerEQuestions, t.TenerETests, t.TenerPTests, t.AdminPTests, t.AdminETests, t.AdminEQuestions,
			t.AdminAnswers, t.AdminUsers, t.AdminTeams, t.AdminConfiguration, t.AdminPermissions, t.TipoInicial, id)
		if err == nil {
			markAsInvalidTipoRolesSingleton()
		}
	}
	return err
}

func DeleteTipoRol(db *sql.DB, id *int64) error {
	if db == nil || id == nil {
		return errors.New(errorDBNil)
	}
	query, err := db.Prepare("DELETE FROM TipoRol WHERE id=?")
	if err == nil {
		defer query.Close()
		_, err = query.Exec(id)
	}
	return err
}
