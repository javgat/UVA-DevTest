// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

package dao

import (
	"fmt"
	"testing"
	"uva-devtest/models"

	"github.com/DATA-DOG/go-sqlmock"
)

// GetRole

func defaultRole() *models.TeamRole {
	rol := models.TeamRoleRoleMember
	r := &models.TeamRole{
		Role: &rol,
	}
	return r
}

func rowsRoles(rs []*models.TeamRole) *sqlmock.Rows {
	columns := []string{"userid", "teamid", "role"}
	sqlcols := sqlmock.NewRows(columns)
	for _, role := range rs {
		sqlcols.AddRow(defaultUser().ID, defaultTeam().ID, role.Role)
	}
	return sqlcols
}

func rowsRole(rs *models.TeamRole) *sqlmock.Rows {
	return rowsRoles([]*models.TeamRole{rs})
}

func expectGetRole(mock sqlmock.Sqlmock, role *models.TeamRole) {
	expectGetUser(mock, *defaultUser().Username, defaultUser())
	expectGetTeam(mock, *defaultTeam().Teamname, defaultTeam())
	rows := rowsRole(role)
	mock.ExpectPrepare("SELECT (.+) FROM Teamroles").ExpectQuery().
		WithArgs(defaultUser().ID, defaultTeam().ID).WillReturnRows(rows)
}

func expectGetRoleDefault(mock sqlmock.Sqlmock) {
	expectGetRole(mock, defaultRole())
}

func expectGetRoleUsernameNotFound(mock sqlmock.Sqlmock) {
	expectGetUserEmpty(mock, *defaultUser().Username)
	expectGetTeam(mock, *defaultTeam().Teamname, defaultTeam())
	mock.ExpectPrepare("SELECT (.+) FROM Teamroles").ExpectQuery().
		WithArgs(defaultUser().ID, defaultTeam().ID).WillReturnError(fmt.Errorf("Error"))
}

func expectGetRoleTeamnameNotFound(mock sqlmock.Sqlmock) {
	expectGetUser(mock, *defaultUser().Username, defaultUser())
	expectGetTeamEmpty(mock, *defaultTeam().Teamname)
	mock.ExpectPrepare("SELECT (.+) FROM Teamroles").ExpectQuery().
		WithArgs(defaultUser().ID, defaultTeam().ID).WillReturnError(fmt.Errorf("Error"))
}

func expectGetRoleError(mock sqlmock.Sqlmock) {
	expectGetUser(mock, *defaultUser().Username, defaultUser())
	expectGetTeam(mock, *defaultTeam().Teamname, defaultTeam())
	mock.ExpectPrepare("SELECT (.+) FROM Teamroles").ExpectQuery().
		WithArgs(defaultUser().ID, defaultTeam().ID).WillReturnError(fmt.Errorf("Error"))
}

func TestGetRoleNilDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetRoleDefault(mock)
	role, err := GetRole(nil, username, teamname)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if role != nil {
		t.Log("role should be nil", err)
		t.Fail()
	}
}
func TestGetRoleClosedDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db.Close()
	expectGetRoleDefault(mock)
	role, err := GetRole(db, username, teamname)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if role != nil {
		t.Log("role should be nil", err)
		t.Fail()
	}
}
func TestGetRoleUsernameNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetRoleUsernameNotFound(mock)
	role, err := GetRole(db, username, teamname)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if role != nil {
		t.Log("role should be nil", err)
		t.Fail()
	}
}
func TestGetRoleTeamnameNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetRoleTeamnameNotFound(mock)
	role, err := GetRole(db, username, teamname)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if role != nil {
		t.Log("role should be nil", err)
		t.Fail()
	}
}
func TestGetRoleError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetRoleError(mock)
	role, err := GetRole(db, username, teamname)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if role != nil {
		t.Log("role should be nil", err)
		t.Fail()
	}
}
func TestGetRoleFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetRoleDefault(mock)
	role, err := GetRole(db, username, teamname)
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	} else if role == nil {
		t.Log("role should not be nil", err)
		t.Fail()
	}
	if *role.Role != *defaultRole().Role {
		t.Log("roles are different", *role.Role, *defaultRole().Role)
		t.Fail()
	}
}

// UpdateRole

func expectUpdateRole(mock sqlmock.Sqlmock, role string) {
	expectGetUser(mock, username, defaultUser())
	expectGetTeam(mock, teamname, defaultTeam())
	mock.ExpectPrepare("UPDATE Teamroles").ExpectExec().
		WithArgs(role, defaultUser().ID, defaultTeam().ID).WillReturnResult(sqlmock.NewResult(1, 1))
}

func expectUpdateRoleDefault(mock sqlmock.Sqlmock) {
	expectUpdateRole(mock, *defaultRole().Role)
}

func expectUpdateRoleUsernameNotFound(mock sqlmock.Sqlmock) {
	role := defaultRole().Role
	expectGetUserEmpty(mock, username)
	expectGetTeam(mock, teamname, defaultTeam())
	mock.ExpectPrepare("UPDATE Teamroles").ExpectExec().
		WithArgs(role, defaultUser().ID, defaultTeam().ID).WillReturnError(fmt.Errorf("Error"))

}

func expectUpdateRoleTeamnameNotFound(mock sqlmock.Sqlmock) {
	role := defaultRole().Role
	expectGetUser(mock, username, defaultUser())
	expectGetTeamEmpty(mock, teamname)
	mock.ExpectPrepare("UPDATE Teamroles").ExpectExec().
		WithArgs(role, defaultUser().ID, defaultTeam().ID).WillReturnError(fmt.Errorf("Error"))

}

func expectUpdateRoleError(mock sqlmock.Sqlmock) {
	role := defaultRole().Role
	expectGetUser(mock, username, defaultUser())
	expectGetTeam(mock, teamname, defaultTeam())
	mock.ExpectPrepare("UPDATE Teamroles").ExpectExec().
		WithArgs(role, defaultUser().ID, defaultTeam().ID).WillReturnError(fmt.Errorf("Error"))

}

func TestUpdateRoleNilDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectUpdateRoleDefault(mock)
	err = UpdateRole(nil, username, teamname, defaultRole())
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}
func TestUpdateRoleClosedDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db.Close()
	expectUpdateRoleDefault(mock)
	err = UpdateRole(db, username, teamname, defaultRole())
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}
func TestUpdateRoleUsernameNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectUpdateRoleUsernameNotFound(mock)
	err = UpdateRole(db, username, teamname, defaultRole())
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}
func TestUpdateRoleTeamnameNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectUpdateRoleTeamnameNotFound(mock)
	err = UpdateRole(db, username, teamname, defaultRole())
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestUpdateRoleError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectUpdateRoleError(mock)
	err = UpdateRole(db, username, teamname, defaultRole())
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}
func TestUpdateRoleCorrect(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectUpdateRoleDefault(mock)
	err = UpdateRole(db, username, teamname, defaultRole())
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}
