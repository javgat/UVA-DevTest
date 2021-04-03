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
	rows := rowsTeams(teams)
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
	} else if teams == nil {
		t.Log("u should not be nil")
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
	rows := rowsTeams(teams)
	expectGetUser(mock, username, defaultUser())
	mock.ExpectPrepare("SELECT (.+) FROM Teams").ExpectQuery().WithArgs(defaultUser().ID).
		WillReturnRows(rows)
}

func expectGetTeamsUsernameDefault(mock sqlmock.Sqlmock) {
	expectGetTeamsUsername(mock, defaultTeams(), username)
}

func expectGetTeamsUsernameUsernameNotFound(mock sqlmock.Sqlmock) {
	rows := rowsTeamEmpty()
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
	rows := rowsTeamEmpty()
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
	} else if teams == nil {
		t.Log("u should not be nil")
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

// GetTeamsTeamRoleAdmin

func expectGetTeamsTeamRoleAdminDefault(mock sqlmock.Sqlmock) {
	expectGetTeamsUsernameDefault(mock)
}

func expectGetTeamsTeamRoleAdminUsernameNotFound(mock sqlmock.Sqlmock) {
	expectGetTeamsUsernameUsernameNotFound(mock)
}

func expectGetTeamsTeamRoleAdminError(mock sqlmock.Sqlmock) {
	expectGetTeamsUsernameError(mock)
}

func TestGetTeamsTeamRoleAdminNilDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetTeamsTeamRoleAdminDefault(mock)
	teams, err := GetTeamsTeamRoleAdmin(nil, username)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if teams != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetTeamsTeamRoleAdminClosedDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db.Close()
	expectGetTeamsTeamRoleAdminDefault(mock)
	teams, err := GetTeamsTeamRoleAdmin(db, username)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if teams != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetTeamsTeamRoleAdminUsernameNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetTeamsTeamRoleAdminUsernameNotFound(mock)
	teams, err := GetTeamsTeamRoleAdmin(db, username)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if teams != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetTeamsTeamRoleAdminError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetTeamsTeamRoleAdminError(mock)
	teams, err := GetTeamsTeamRoleAdmin(db, username)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if teams != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetTeamsTeamRoleAdminFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetTeamsTeamRoleAdminDefault(mock)
	teams, err := GetTeamsTeamRoleAdmin(db, username)
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

// GetTeam

func expectGetTeamDefault(mock sqlmock.Sqlmock) {
	expectGetTeam(mock, teamname, defaultTeam())
}

func expectGetTeamError(mock sqlmock.Sqlmock) {
	mock.ExpectPrepare("SELECT (.+) FROM Teams").ExpectQuery().WithArgs(teamname).
		WillReturnError(fmt.Errorf("Error"))
}

func TestGetTeamNilDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetTeamDefault(mock)
	team, err := GetTeam(nil, teamname)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if team != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetTeamClosedDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db.Close()
	expectGetTeamDefault(mock)
	team, err := GetTeam(db, teamname)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if team != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetTeamError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetTeamError(mock)
	team, err := GetTeam(db, teamname)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if team != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetTeamFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetTeamDefault(mock)
	team, err := GetTeam(db, teamname)
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	} else if team == nil {
		t.Log("u should not be nil")
		t.Fail()
	}
	if *team.Teamname != teamname {
		t.Log("teamnames incorrectos")
		t.Fail()
	}
}

// PostTeam

func expectPostTeam(mock sqlmock.Sqlmock, team *Team) {
	expectGetUser(mock, username, defaultUser())
	mock.ExpectPrepare("INSERT INTO Teams").ExpectExec().
		WithArgs(team.Teamname, team.Description).WillReturnResult(sqlmock.NewResult(1, 1))
	expectAddUserTeamAdmin(mock, defaultUser().ID, team.ID)
	expectUpdateRole(mock, TeamRoleRoleAdmin)
}

func expectPostTeamDefault(mock sqlmock.Sqlmock) {
	expectPostTeam(mock, defaultTeam())
}

func expectPostTeamUsernameNotFound(mock sqlmock.Sqlmock) {
	expectGetUserEmpty(mock, username)
	team := defaultTeam()
	mock.ExpectPrepare("INSERT INTO Teams").ExpectExec().
		WithArgs(team.Teamname, team.Description).WillReturnError(fmt.Errorf("Error"))
}

func expectPostTeamError(mock sqlmock.Sqlmock) {
	team := defaultTeam()
	expectGetUser(mock, username, defaultUser())
	mock.ExpectPrepare("INSERT INTO Teams").ExpectExec().
		WithArgs(team.Teamname, team.Description).WillReturnError(fmt.Errorf("Error"))
}

func TestPostTeamNilDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectPostTeamDefault(mock)
	err = PostTeam(nil, ToModelTeam(defaultTeam()), username)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestPostTeamClosedDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db.Close()
	expectPostTeamDefault(mock)
	err = PostTeam(db, ToModelTeam(defaultTeam()), username)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestPostTeamUsernameNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectPostTeamUsernameNotFound(mock)
	err = PostTeam(db, ToModelTeam(defaultTeam()), username)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestPostTeamError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectPostTeamError(mock)
	err = PostTeam(db, ToModelTeam(defaultTeam()), username)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestPostTeamCorrect(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectPostTeamDefault(mock)
	err = PostTeam(db, ToModelTeam(defaultTeam()), username)
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}

// UpdateTeam

func expectUpdateTeam(mock sqlmock.Sqlmock, team *Team) {
	mock.ExpectPrepare("UPDATE Teams").ExpectExec().
		WithArgs(team.Teamname, team.Description, team.Teamname).WillReturnResult(sqlmock.NewResult(1, 1))
}

func expectUpdateTeamDefault(mock sqlmock.Sqlmock) {
	expectUpdateTeam(mock, defaultTeam())
}

func expectUpdateTeamError(mock sqlmock.Sqlmock) {
	team := defaultTeam()
	mock.ExpectPrepare("UPDATE Teams").ExpectExec().
		WithArgs(team.Teamname, team.Description, team.Teamname).WillReturnError(fmt.Errorf("Error"))
}

func TestUpdateTeamNilDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectUpdateTeamDefault(mock)
	err = UpdateTeam(nil, ToModelTeam(defaultTeam()), teamname)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestUpdateTeamClosedDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db.Close()
	expectUpdateTeamDefault(mock)
	err = UpdateTeam(db, ToModelTeam(defaultTeam()), teamname)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestUpdateTeamError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectUpdateTeamError(mock)
	err = UpdateTeam(db, ToModelTeam(defaultTeam()), teamname)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestUpdateTeamCorrect(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectUpdateTeamDefault(mock)
	err = UpdateTeam(db, ToModelTeam(defaultTeam()), teamname)
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}

// DeleteTeam

func expectDeleteTeam(mock sqlmock.Sqlmock, teamname string) {
	mock.ExpectPrepare("DELETE FROM Teams").ExpectExec().
		WithArgs(teamname).WillReturnResult(sqlmock.NewResult(1, 1))
}

func expectDeleteTeamDefault(mock sqlmock.Sqlmock) {
	expectDeleteTeam(mock, teamname)
}

func expectDeleteTeamError(mock sqlmock.Sqlmock) {
	mock.ExpectPrepare("DELETE FROM Teams").ExpectExec().
		WithArgs(teamname).WillReturnError(fmt.Errorf("Error"))
}

func TestDeleteTeamNilDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectDeleteTeamDefault(mock)
	err = DeleteTeam(nil, teamname)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestDeleteTeamClosedDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db.Close()
	expectDeleteTeamDefault(mock)
	err = DeleteTeam(db, teamname)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestDeleteTeamError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectDeleteTeamError(mock)
	err = DeleteTeam(db, teamname)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestDeleteTeamCorrect(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectDeleteTeamDefault(mock)
	err = DeleteTeam(db, teamname)
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}
