// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package dao acts as a Data Access Object for the Types
package dao

import (
	"database/sql"
	"errors"
	"uva-devtest/models"
)

// Transforms some sql.Rows into a slice(array) of roles
// Param rows: Rows which contains database information returned
// Return []models.TeamRole: TeamRoles represented in rows
// Return error if any
func rowsToRoles(rows *sql.Rows) ([]*models.TeamRole, error) {
	var roles []*models.TeamRole
	var userid, teamid int
	for rows.Next() {
		var r models.TeamRole
		err := rows.Scan(&userid, &teamid, &r.Role)
		if err != nil {
			return roles, err
		}
		roles = append(roles, &r)
	}
	return roles, nil
}

// Transforms rows into a single role
// Param rows: Rows which contains database info of 1 Role
// Return *models.TeamRole: TeamRole that was represented in rows
// Return error if something happens
func rowsToRole(rows *sql.Rows) (*models.TeamRole, error) {
	var role *models.TeamRole
	roles, err := rowsToRoles(rows)
	if len(roles) >= 1 {
		role = roles[0]
	}
	return role, err
}

// GetRole gets the TeamRole of user <username> at team <teamname>
// Param username: Username of the user
// Param teamname: Teamname of the team
func GetRole(db *sql.DB, username string, teamname string) (role *models.TeamRole, err error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err != nil {
		return nil, err
	} else if u == nil {
		return nil, errors.New(errorResourceNotFound)
	}
	t, err := GetTeam(db, teamname)
	if err != nil {
		return nil, err
	} else if t == nil {
		return nil, errors.New("Team no existe")
	}
	query, err := db.Prepare("SELECT * FROM EquipoUsuario WHERE userid = ? AND teamid = ? ")
	defer query.Close()
	if err != nil {
		return nil, err
	}
	rows, err := query.Query(u.ID, t.ID)
	if err == nil {
		role, err = rowsToRole(rows)
	}
	return role, err
}

// UpdateRole updates a role to the database
// Param username: Username of the user
// Param teamname: Teamname of the team
// Param role: New role to update
func UpdateRole(db *sql.DB, username string, teamname string, role *models.TeamRole) error {
	if db == nil || role == nil {
		return errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err != nil {
		return err
	} else if u == nil {
		return errors.New(errorResourceNotFound)
	}
	t, err := GetTeam(db, teamname)
	if err != nil {
		return err
	} else if t == nil {
		return errors.New("Team no existe")
	}
	query, err := db.Prepare("UPDATE EquipoUsuario SET rol = ? WHERE userid = ? AND teamid = ? ")
	defer query.Close()
	if err != nil {
		return err
	}
	_, err = query.Exec(role.Role, u.ID, t.ID)
	return err
}
