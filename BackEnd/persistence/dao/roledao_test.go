// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

package dao

import (
	"github.com/DATA-DOG/go-sqlmock"
)

func expectUpdateRole(mock sqlmock.Sqlmock, role string) {
	expectGetUser(mock, username, defaultUser())
	expectGetTeam(mock, teamname, defaultTeam())
	mock.ExpectPrepare("UPDATE Teamroles").ExpectExec().
		WithArgs(role, defaultUser().ID, defaultTeam().ID).WillReturnResult(sqlmock.NewResult(1, 1))
}
