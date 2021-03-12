// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

package dao

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

// GetTeams

func defaultTeams() []*Team {
	ts := []*Team{}
	ts = append(ts, defaultTeam())
	return ts
}

func expectGetTeams(mock sqlmock.Sqlmock, teams []*Team) {
	columns := []string{"id", "teamname", "description"}
	rows := sqlmock.NewRows(columns)
	for _, team := range teams {
		rows.AddRow(team.ID, team.Teamname, team.Description)
	}
	mock.ExpectPrepare("SELECT (.+) FROM Teams").ExpectQuery().
		WillReturnRows(rows)
}

func expectGetTeamsDefault(mock sqlmock.Sqlmock) {
	expectGetTeams(mock, defaultTeams())
}

func expectGetTeamsError(mock sqlmock.Sqlmock) {
	mock.ExpectPrepare("SELECT (.+) FROM Teams").ExpectQuery().
		WillReturnError(fmt.Errorf("Error"))
}

func expectGetTeamsEmpty(mock sqlmock.Sqlmock) {
	expectGetTeams(mock, []*Team{})
}

func TestGetTeamsNilBD(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetTeamsDefault(mock)
	teams, err := GetTeams(nil)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if teams != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetTeamsClosedBD(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db.Close()
	expectGetTeamsDefault(mock)
	teams, err := GetTeams(db)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if teams != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetTeamsError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetTeamsError(mock)
	teams, err := GetTeams(db)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if teams != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetTeamsEmpty(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetTeamsEmpty(mock)
	teams, err := GetTeams(db)
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	} else if teams != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetTeamsFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetTeamsDefault(mock)
	teams, err := GetTeams(db)
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	} else if teams == nil {
		t.Log("u should not be nil")
		t.Fail()
	}
	ts := defaultTeams()
	if *teams[0].Teamname != *ts[0].Teamname || len(teams) != len(ts) {
		t.Log("teamnames incorrectos")
	}
}

// GetTeamsUsername

func expectGetTeamsUsername(mock sqlmock.Sqlmock, teams []*Team, username string) {
	columns := []string{"id", "teamname", "description"}
	rows := sqlmock.NewRows(columns)
	for _, team := range teams {
		rows.AddRow(team.ID, team.Teamname, team.Description)
	}
	expectGetUser(mock, username, defaultUser())
	mock.ExpectPrepare("SELECT (.+) FROM Teams").ExpectQuery().WithArgs(defaultUser().ID).
		WillReturnRows(rows)
}

func expectGetTeamsUsernameDefault(mock sqlmock.Sqlmock) {
	expectGetTeamsUsername(mock, defaultTeams(), username)
}

func expectGetTeamsUsernameUsernameNotFound(mock sqlmock.Sqlmock) {
	columns := []string{"id", "teamname", "description"}
	rows := sqlmock.NewRows(columns)
	expectGetUserEmpty(mock, username)
	mock.ExpectPrepare("SELECT (.+) FROM Teams").ExpectQuery().WithArgs(defaultUser().ID).
		WillReturnRows(rows)
}

func expectGetTeamsUsernameError(mock sqlmock.Sqlmock) {
	expectGetUser(mock, username, defaultUser())
	mock.ExpectPrepare("SELECT (.+) FROM Teams").ExpectQuery().WithArgs(defaultUser().ID).
		WillReturnError(fmt.Errorf("Error"))
}

func expectGetTeamsUsernameEmpty(mock sqlmock.Sqlmock) {
	columns := []string{"id", "teamname", "description"}
	rows := sqlmock.NewRows(columns)
	expectGetUser(mock, username, defaultUser())
	mock.ExpectPrepare("SELECT (.+) FROM Teams").ExpectQuery().WithArgs(defaultUser().ID).
		WillReturnRows(rows)
}

func TestGetTeamsUsernameNilDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetTeamsUsernameDefault(mock)
	teams, err := GetTeamsUsername(nil, username)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if teams != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetTeamsUsernameClosedDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db.Close()
	expectGetTeamsUsernameDefault(mock)
	teams, err := GetTeamsUsername(db, username)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if teams != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetTeamsUsernameUsernameNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetTeamsUsernameUsernameNotFound(mock)
	teams, err := GetTeamsUsername(db, username)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if teams != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetTeamsUsernameError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetTeamsUsernameError(mock)
	teams, err := GetTeamsUsername(db, username)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if teams != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetTeamsUsernameEmpty(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetTeamsUsernameEmpty(mock)
	teams, err := GetTeamsUsername(db, username)
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	} else if teams != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetTeamsUsernameFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetTeamsUsernameDefault(mock)
	teams, err := GetTeamsUsername(db, username)
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	} else if teams == nil {
		t.Log("u should not be nil")
		t.Fail()
	}
	ts := defaultTeams()
	if *teams[0].Teamname != *ts[0].Teamname || len(ts) != len(teams) {
		t.Log("teamnames incorrectos")
		t.Fail()
	}
}
