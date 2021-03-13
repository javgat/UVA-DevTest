// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package dao acts as a Data Access Object for the Types
package dao

import (
	"database/sql"
	"errors"
	"uva-devtest/models"
)

// ToModelTeam converts a teamdao.Team into a models.Team
func ToModelTeam(t *Team) *models.Team {
	mt := &models.Team{
		Teamname:    t.Teamname,
		Description: t.Description,
	}
	return mt
}

// ToModelsTeams converts a splice of teamdao.Team into models.Team
func ToModelsTeams(us []*Team) []*models.Team {
	var mts = []*models.Team{}
	for _, itemCopy := range us {
		mts = append(mts, ToModelTeam(itemCopy))
	}
	return mts
}

// Transforms some sql.Rows into a slice(array) of teams
// Param rows: Rows which contains database information returned
// Return []Team: Team represented in rows
// Return error if any
func rowsToTeams(rows *sql.Rows) ([]*Team, error) {
	var teams []*Team
	for rows.Next() {
		var t Team
		err := rows.Scan(&t.ID, &t.Teamname, &t.Description)
		if err != nil {
			return teams, err
		}
		teams = append(teams, &t)
	}
	return teams, nil
}

// Transforms rows into a single team
// Param rows: Rows which contains database info of 1 Team
// Return *Team: Team that was represented in rows
// Return error if something happens
func rowsToTeam(rows *sql.Rows) (*Team, error) {
	var team *Team
	teams, err := rowsToTeams(rows)
	if len(teams) >= 1 {
		team = teams[0]
	}
	return team, err
}

// GetTeam gets team <teamname>
// Param teamname: Teamname of the team
func GetTeam(db *sql.DB, teamname string) (*Team, error) {
	if db == nil {
		return nil, errors.New("Parametro db nil")
	}
	query, err := db.Prepare("SELECT * FROM Teams WHERE teamname=?")
	var t *Team
	if err != nil {
		return t, err
	}
	rows, err := query.Query(teamname)
	if err == nil {
		t, err = rowsToTeam(rows)
	}
	defer query.Close()
	return t, err
}

// PostTeam creates a team into the database
// Param db: Database to use
// Param t: Team data to create
// Return error if something wrong happens
func PostTeam(db *sql.DB, t *models.Team, username string) error {
	if db == nil || t == nil {
		return errors.New("Argumento de entrada nil")
	}
	u, err := GetUserUsername(db, username)
	if err != nil {
		return err
	} else if u == nil {
		return errors.New("User not found")
	}
	query, err := db.Prepare("INSERT INTO Teams(teamname, description) VALUES(?, ?)")
	if err != nil {
		return err
	}
	_, err = query.Exec(t.Teamname, t.Description)
	defer query.Close()
	if err != nil {
		return err
	}
	err = AddUserTeam(db, username, *t.Teamname)
	if err != nil {
		return err
	}
	roleString := models.TeamRoleRoleAdmin
	role := &models.TeamRole{
		Role: &roleString,
	}
	err = UpdateRole(db, username, *t.Teamname, role)
	return err
}

// UpdateTeam updates a team to the database
// Param db: Database to use
// Param t: Team data to update
// Param teamname: Teamname of the team to update
// Return error if something wrong happens
func UpdateTeam(db *sql.DB, t *models.Team, teamname string) error {
	if db == nil || t == nil {
		return errors.New("Argumento de entrada nil")
	}
	query, err := db.Prepare("UPDATE Teams SET teamname=?, description=? WHERE teamname = ? ")
	if err != nil {
		return err
	}
	_, err = query.Exec(t.Teamname, t.Description, teamname)
	defer query.Close()
	return err
}

// DeleteTeam Deletes a team
func DeleteTeam(db *sql.DB, teamname string) error {
	if db == nil {
		return errors.New("Argumento de entrada nil")
	}
	query, err := db.Prepare("DELETE FROM Teams WHERE teamname = ? ") //ESTO se supone que borra en cascade
	if err != nil {
		return err
	}
	_, err = query.Exec(teamname)
	defer query.Close()
	return err
}

// GetTeams gets all teams
func GetTeams(db *sql.DB) ([]*Team, error) {
	if db == nil {
		return nil, errors.New("Parametro db nil")
	}
	query, err := db.Prepare("SELECT * FROM Teams")
	var ts []*Team
	if err != nil {
		return ts, err
	}
	rows, err := query.Query()
	if err == nil {
		ts, err = rowsToTeams(rows)
	}
	defer query.Close()
	return ts, err
}

// GetTeamsUsername gets all teams from user <username>
// Param username: Username of the user
func GetTeamsUsername(db *sql.DB, username string) ([]*Team, error) {
	if db == nil {
		return nil, errors.New("Parametro db nil")
	}
	u, err := GetUserUsername(db, username)
	if err != nil {
		return nil, err
	} else if u == nil {
		return nil, errors.New("No se encontro al usuario")
	}
	query, err := db.Prepare("SELECT T FROM Teams T JOIN Teamroles R ON	T.id=R.teamid WHERE R.userid = ?")
	var ts []*Team
	if err != nil {
		return ts, err
	}
	rows, err := query.Query(u.ID)
	if err == nil {
		ts, err = rowsToTeams(rows)
	}
	defer query.Close()
	return ts, err
}

// GetTeamsTeamRoleAdmin gets all teams where user is Admin
// Param username: Username of the user
func GetTeamsTeamRoleAdmin(db *sql.DB, username string) ([]*Team, error) {
	if db == nil {
		return nil, errors.New("Parametro db nil")
	}
	u, err := GetUserUsername(db, username)
	if err != nil {
		return nil, err
	} else if u == nil {
		return nil, errors.New("User not found")
	}
	query, err := db.Prepare("SELECT * FROM Teams T JOIN Teamroles R ON R.teamid=T.id WHERE R.userid=? AND R.role='Admin'")
	var ts []*Team
	if err != nil {
		return ts, err
	}
	rows, err := query.Query(u.ID)
	if err == nil {
		ts, err = rowsToTeams(rows)
	}
	defer query.Close()
	return ts, err
}
