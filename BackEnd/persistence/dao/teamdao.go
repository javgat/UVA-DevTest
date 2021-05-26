// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package dao acts as a Data Access Object for the Types
package dao

import (
	"database/sql"
	"errors"
	"strconv"
	"uva-devtest/models"
)

// ToModelTeam converts a teamdao.Team into a models.Team
func ToModelTeam(t *Team) *models.Team {
	mt := &models.Team{
		Teamname:       t.Teamname,
		Description:    t.Description,
		SoloProfesores: t.SoloProfesores,
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
	teams := []*Team{}
	for rows.Next() {
		var t Team
		err := rows.Scan(&t.ID, &t.Teamname, &t.Description, &t.SoloProfesores)
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
		return nil, errors.New(errorDBNil)
	}
	query, err := db.Prepare("SELECT * FROM Equipo WHERE teamname=?")
	var t *Team
	if err != nil {
		return t, err
	}
	defer query.Close()
	rows, err := query.Query(teamname)
	if err == nil {
		t, err = rowsToTeam(rows)
	}
	return t, err
}

// PostTeam creates a team into the database
// Param db: Database to use
// Param t: Team data to create
// Return error if something wrong happens
func PostTeam(db *sql.DB, t *models.Team, username string) error {
	if db == nil || t == nil {
		return errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err != nil {
		return err
	} else if u == nil {
		return errors.New("User not found")
	}
	query, err := db.Prepare("INSERT INTO Equipo(teamname, description, soloProfesores) VALUES(?,?,?)")
	if err != nil {
		return err
	}
	defer query.Close()
	_, err = query.Exec(t.Teamname, t.Description, t.SoloProfesores)
	if err != nil {
		return err
	}
	err = AddUserTeamAdmin(db, username, *t.Teamname)
	if err != nil {
		return err
	}
	roleString := TeamRoleRoleAdmin
	role := &TeamRole{
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
		return errors.New(errorDBNil)
	}
	query, err := db.Prepare("UPDATE Equipo SET teamname=?, description=? WHERE teamname = ? ")
	if err != nil {
		return err
	}
	defer query.Close()
	_, err = query.Exec(t.Teamname, t.Description, teamname)
	return err
}

// DeleteTeam Deletes a team
func DeleteTeam(db *sql.DB, teamname string) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	query, err := db.Prepare("DELETE FROM Equipo WHERE teamname = ? ") //ESTO se supone que borra en cascade
	if err != nil {
		return err
	}
	defer query.Close()
	_, err = query.Exec(teamname)
	return err
}

func addFiltersTeams(hayWhere bool, initQuery string, likeStartTeamname *string, limit *int64, offset *int64) string {
	query := initQuery
	nexoQuery := " WHERE "
	if hayWhere {
		nexoQuery = " AND "
	}
	if likeStartTeamname != nil && *likeStartTeamname != "" {
		query = query + nexoQuery + " teamname LIKE ? "
		nexoQuery = " AND "
	}
	query += " ORDER BY teamname ASC"
	if limit != nil {
		query = query + " LIMIT " + strconv.FormatInt(*limit, 10) + " "
	}
	if offset != nil {
		query = query + " OFFSET " + strconv.FormatInt(*offset, 10) + " "
	}
	return query
}

func FilterTeamParamsToInterfaceArr(likeStartTeamname *string) []interface{} {
	hayLikeStartTeamname := 0
	if likeStartTeamname != nil && *likeStartTeamname != "" {
		hayLikeStartTeamname = 1
	}
	interfaceParams := make([]interface{}, hayLikeStartTeamname)
	if hayLikeStartTeamname == 1 {
		interfaceParams[0] = *likeStartTeamname + "%"
	}
	return interfaceParams
}

// GetTeams gets all teams
func GetTeams(db *sql.DB, likeStartTeamname *string, limit *int64, offset *int64) ([]*Team, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	stPrepare := "SELECT * FROM Equipo "
	stPrepare = addFiltersTeams(false, stPrepare, likeStartTeamname, limit, offset)
	query, err := db.Prepare(stPrepare)
	var ts []*Team
	if err != nil {
		return ts, err
	}
	defer query.Close()
	interfaceParams := FilterTeamParamsToInterfaceArr(likeStartTeamname)
	rows, err := query.Query(interfaceParams...)
	if err == nil {
		ts, err = rowsToTeams(rows)
	}
	return ts, err
}

// GetTeamsUsername gets all teams from user <username>
// Param username: Username of the user
func GetTeamsUsername(db *sql.DB, username string) ([]*Team, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err != nil {
		return nil, err
	} else if u == nil {
		return nil, errors.New(errorResourceNotFound)
	}
	var ts []*Team
	query, err := db.Prepare("SELECT T.* FROM Equipo T JOIN EquipoUsuario R ON	T.id=R.equipoid WHERE R.usuarioid = ?")
	if err != nil {
		return ts, err
	}
	defer query.Close()
	rows, err := query.Query(u.ID)
	if err == nil {
		ts, err = rowsToTeams(rows)
	}
	return ts, err
}

func GetTeamFromUser(db *sql.DB, teamname string, username string) (*Team, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err != nil {
		return nil, err
	} else if u == nil {
		return nil, errors.New("no se encontro al usuario")
	}
	var t *Team
	query, err := db.Prepare("SELECT T.* FROM Equipo T JOIN EquipoUsuario R ON	T.id=R.equipoid WHERE R.usuarioid = ? AND R.equipoid = ?")
	if err != nil {
		return nil, err
	}
	defer query.Close()
	th, err := GetTeam(db, teamname)
	if err != nil {
		return nil, err
	} else if th == nil {
		return nil, errors.New("no se encontro al equipo")
	}
	rows, err := query.Query(u.ID, th.ID)
	if err == nil {
		t, err = rowsToTeam(rows)
	}
	return t, err
}

// GetTeamsTeamRoleAdmin gets all teams where user is Admin
// Param username: Username of the user
func GetTeamsTeamRoleAdmin(db *sql.DB, username string) ([]*Team, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err != nil {
		return nil, err
	} else if u == nil {
		return nil, errors.New("User not found")
	}
	var ts []*Team
	query, err := db.Prepare("SELECT * FROM Equipo T JOIN EquipoUsuario R ON R.equipoid=T.id WHERE R.usuarioid=? AND R.rol='Admin'")
	if err != nil {
		return ts, err
	}
	defer query.Close()
	rows, err := query.Query(u.ID)
	if err == nil {
		ts, err = rowsToTeams(rows)
	}
	return ts, err
}

func GetTeamsQuestion(db *sql.DB, questionid int64) ([]*Team, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var ts []*Team
	query, err := db.Prepare("SELECT E.* FROM Equipo E JOIN PreguntaEquipo P ON P.equipoid=E.id WHERE P.preguntaid=?")
	if err == nil {
		defer query.Close()
		rows, err := query.Query(questionid)
		if err == nil {
			ts, err = rowsToTeams(rows)
			return ts, err
		}
	}
	return nil, err
}

func GetAdminTeamsFromTest(db *sql.DB, testid int64) ([]*Team, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var ts []*Team
	query, err := db.Prepare("SELECT E.* FROM Equipo E JOIN GestionTestEquipo G ON G.equipoid=E.id WHERE G.testid=?")
	if err == nil {
		defer query.Close()
		rows, err := query.Query(testid)
		if err == nil {
			ts, err = rowsToTeams(rows)
			return ts, err
		}
	}
	return nil, err
}

func AddAdminTeamToTest(db *sql.DB, testid int64, teamname string) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	ts, err := GetTeam(db, teamname)
	if err != nil || ts == nil {
		return err
	}
	query, err := db.Prepare("INSERT INTO GestionTestEquipo(equipoid, testid) VALUES(?,?)")
	if err == nil {
		defer query.Close()
		_, err = query.Exec(ts.ID, testid)
	}
	return err
}

func RemoveAdminTeamFromTest(db *sql.DB, testid int64, teamname string) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	ts, err := GetTeam(db, teamname)
	if err != nil || ts == nil {
		return err
	}
	query, err := db.Prepare("DELETE FROM GestionTestEquipo WHERE equipoid=? AND testid=?")
	if err == nil {
		defer query.Close()
		_, err = query.Exec(ts.ID, testid)
	}
	return err
}
