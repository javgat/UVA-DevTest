package teamdao

import (
	"database/sql"
	"errors"
	"uva-devtest/models"
)

// Transforms some sql.Rows into a slice(array) of teams
// Param rows: Rows which contains database information returned
// Return []models.Team: Team represented in rows
// Return error if any
func rowsToTeams(rows *sql.Rows) ([]*models.Team, error) {
	var teams []*models.Team
	var trash int
	for rows.Next() {
		var t models.Team
		err := rows.Scan(&trash, &t.Teamname, &t.Description)
		if err != nil {
			return teams, err
		}
		teams = append(teams, &t)
	}
	return teams, nil
}

// Transforms rows into a single tea,
// Param rows: Rows which contains database info of 1 Role
// Return *models.Team: Team that was represented in rows
// Return error if something happens
func rowsToTeam(rows *sql.Rows) (*models.Team, error) {
	var team *models.Team
	teams, err := rowsToTeams(rows)
	if len(teams) >= 1 {
		team = teams[0]
	}
	return team, err
}

// GetTeamsUsername gets all teams from user <username>
func GetTeamsUsername(db *sql.DB, username string) ([]*models.Team, error) {
	if db == nil {
		return nil, errors.New("Parametro db nil")
	}
	query, err := db.Prepare("SELECT T FROM teams T JOIN users U WHERE U.username = ?")
	var ts []*models.Team
	if err != nil {
		return ts, err
	}
	rows, err := query.Query(username)
	if err == nil {
		ts, err = rowsToTeams(rows)
	}
	defer query.Close()
	return ts, err
}
