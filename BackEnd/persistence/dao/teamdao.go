package dao

import (
	"database/sql"
	"errors"
	"uva-devtest/models"
)

// DaoToModelTeam converts a teamdao.Team into a models.Team
func DaoToModelTeam(t *Team) *models.Team {
	mt := &models.Team{
		Teamname:    t.Teamname,
		Description: t.Description,
	}
	return mt
}

// DaoToModelsTeams converts a splice of teamdao.Team into models.Team
func DaoToModelsTeams(us []*Team) []*models.Team {
	var mts = []*models.Team{}
	for _, itemCopy := range us {
		mts = append(mts, DaoToModelTeam(itemCopy))
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

// GetTeamsUsername gets all teams from user <username>
func GetTeamsUsername(db *sql.DB, username string) ([]*Team, error) {
	if db == nil {
		return nil, errors.New("Parametro db nil")
	}
	u, err := GetUserUsername(db, username)
	if err != nil {
		return nil, err
	}
	query, err := db.Prepare("SELECT T FROM teams T JOIN teamroles R ON	T.id=R.teamid WHERE R.userid = ?")
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

// GetTeams gets all teams
func GetTeams(db *sql.DB) ([]*Team, error) {
	if db == nil {
		return nil, errors.New("Parametro db nil")
	}
	query, err := db.Prepare("SELECT * FROM teams")
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

// GetTeam gets team <teamname>
func GetTeam(db *sql.DB, teamname string) (*Team, error) {
	if db == nil {
		return nil, errors.New("Parametro db nil")
	}
	query, err := db.Prepare("SELECT * FROM teams")
	var t *Team
	if err != nil {
		return t, err
	}
	rows, err := query.Query()
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
func PostTeam(db *sql.DB, t *models.Team) error {
	if db == nil || t == nil {
		return errors.New("Argumento de entrada nil")
	}
	query, err := db.Prepare("INSERT INTO teams(teamname, description) VALUES(?, ?)")
	if err != nil {
		return err
	}
	_, err = query.Exec(t.Teamname, t.Description)
	defer query.Close()
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
	query, err := db.Prepare("UPDATE teams SET teamname=?, description=? WHERE teamname = ? ")
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
	query, err := db.Prepare("DELETE FROM teams WHERE teamname = ? ") //ESTO TENDRA QUE SER MAS COMPLEJO, RELACIONES
	if err != nil {
		return err
	}
	_, err = query.Exec(teamname)
	defer query.Close()
	return err
}
