// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package dao acts as a Data Access Object for the Types
package dao

import (
	"database/sql"
	"errors"

	// Blank import of mysql driver
	_ "github.com/go-sql-driver/mysql"
)

func GetPublicPublishedTests(db *sql.DB) ([]*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var ts []*Test
	query, err := db.Prepare("SELECT * FROM Test WHERE editable=0 AND T.accesoPublico=1")
	if err == nil {
		defer query.Close()
		rows, err := query.Query()
		if err == nil {
			ts, err = rowsToTests(rows)
			return ts, err
		}
	}
	return nil, err
}

func GetPublicPublishedTest(db *sql.DB, testid int64) (*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var ts *Test
	query, err := db.Prepare("SELECT * FROM Test WHERE id=? AND editable=0 AND T.accesoPublico=1")
	if err == nil {
		defer query.Close()
		rows, err := query.Query(testid)
		if err == nil {
			ts, err = rowsToTest(rows)
			return ts, err
		}
	}
	return nil, err
}

func GetPublishedTests(db *sql.DB) ([]*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var ts []*Test
	query, err := db.Prepare("SELECT * FROM Test WHERE editable=0")
	if err == nil {
		defer query.Close()
		rows, err := query.Query()
		if err == nil {
			ts, err = rowsToTests(rows)
			return ts, err
		}
	}
	return nil, err
}

func GetPublishedTest(db *sql.DB, testid int64) (*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var ts *Test
	query, err := db.Prepare("SELECT * FROM Test WHERE id=? AND editable=0")
	if err == nil {
		defer query.Close()
		rows, err := query.Query(testid)
		if err == nil {
			ts, err = rowsToTest(rows)
			return ts, err
		}
	}
	return nil, err
}

func GetUsersInvitedPTest(db *sql.DB, testid int64) ([]*User, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var us []*User
	query, err := db.Prepare("SELECT U.* FROM Usuario U JOIN InvitacionTestUsuario I ON I.usuarioid=U.id WHERE I.testid=?")
	if err == nil {
		defer query.Close()
		rows, err := query.Query(testid)
		if err == nil {
			us, err = rowsToUsers(rows)
			return us, err
		}
	}
	return nil, err
}

func InviteUserPTest(db *sql.DB, testid int64, username string) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err == nil {
		RemoveUserPTest(db, testid, username)
		var query *sql.Stmt
		query, err = db.Prepare("INSERT INTO InvitacionTestUsuario(usuarioid, testid) VALUES(?,?)")
		if err == nil {
			defer query.Close()
			_, err = query.Exec(u.ID, testid)
		}
	}
	return err
}

func RemoveUserPTest(db *sql.DB, testid int64, username string) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err == nil {
		var query *sql.Stmt
		query, err = db.Prepare("DELETE FROM InvitacionTestUsuario WHERE usuarioid=? AND testid=?")
		if err == nil {
			defer query.Close()
			_, err = query.Exec(u.ID, testid)
		}
	}
	return err
}

func GetTeamsInvitedPTest(db *sql.DB, testid int64) ([]*Team, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var ts []*Team
	query, err := db.Prepare("SELECT E.* FROM Equipo E JOIN InvitacionTestEquipo I ON I.equipoid=E.id WHERE I.testid=?")
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

func InviteTeamPTest(db *sql.DB, testid int64, teamname string) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	t, err := GetTeam(db, teamname)
	if err == nil {
		RemoveTeamPTest(db, testid, teamname)
		var query *sql.Stmt
		query, err = db.Prepare("INSERT INTO InvitacionTestEquipo(equipoid, testid) VALUES(?,?)")
		if err == nil {
			defer query.Close()
			_, err = query.Exec(t.ID, testid)
		}
	}
	return err
}

func RemoveTeamPTest(db *sql.DB, testid int64, teamname string) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	t, err := GetTeam(db, teamname)
	if err == nil {
		var query *sql.Stmt
		query, err = db.Prepare("DELETE FROM InvitacionTestEquipo WHERE equipoid=? AND testid=?")
		if err == nil {
			defer query.Close()
			_, err = query.Exec(t.ID, testid)
		}
	}
	return err
}

func GetAnswersFromPTest(db *sql.DB, testid int64) ([]*Answer, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var as []*Answer
	query, err := db.Prepare("SELECT * FROM RespuestaExamen WHERE testid=?")
	if err == nil {
		defer query.Close()
		rows, err := query.Query(testid)
		if err == nil {
			as, err = rowsToAnswers(rows)
			return as, err
		}
	}
	return nil, err
}

func GetInvitedTestsFromTeam(db *sql.DB, teamname string) ([]*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetTeam(db, teamname)
	if err == nil {
		var t []*Test
		query, err := db.Prepare("SELECT T.* FROM Test T JOIN InvitacionTestEquipo I ON T.id=I.testid WHERE I.equipoid=? AND T.editable=0")
		if err == nil {
			defer query.Close()
			rows, err := query.Query(u.ID)
			if err == nil {
				t, err = rowsToTests(rows)
				return t, err
			}
		}
	}
	return nil, err
}

func GetInvitedTestFromTeam(db *sql.DB, teamname string, testid int64) (*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetTeam(db, teamname)
	if err == nil {
		var t *Test
		query, err := db.Prepare("SELECT T.* FROM Test T JOIN InvitacionTestEquipo I ON T.id=I.testid WHERE I.equipoid=? AND T.id=? AND T.editable=0")
		if err == nil {
			defer query.Close()
			rows, err := query.Query(u.ID, testid)
			if err == nil {
				t, err = rowsToTest(rows)
				return t, err
			}
		}
	}
	return nil, err
}

// InvitedTests

func GetInvitedPTestsFromUser(db *sql.DB, username string) ([]*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err == nil {
		var ts []*Test
		query, err := db.Prepare("SELECT T.* FROM Test T JOIN InvitacionTestUsuario I ON T.id=I.testid WHERE I.usuarioid=? AND T.editable=0")

		if err == nil {
			defer query.Close()
			rows, err := query.Query(u.ID)
			if err == nil {
				ts, err = rowsToTests(rows)
				return ts, err
			}
		}
	}
	return nil, err
}

func GetInvitedPTestFromUser(db *sql.DB, username string, testid int64) (*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err == nil {
		var ts *Test
		query, err := db.Prepare("SELECT T.* FROM Test T JOIN InvitacionTestUsuario I ON T.id=I.testid WHERE I.usuarioid=? AND I.testid=? AND T.editable=0")
		if err == nil {
			defer query.Close()
			rows, err := query.Query(u.ID, testid)
			if err == nil {
				ts, err = rowsToTest(rows)
				return ts, err
			}
		}
	}
	return nil, err
}

// PublishedTests

func GetPublishedTestsFromUser(db *sql.DB, username string) ([]*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var ts []*Test
	query, err := db.Prepare("SELECT T.* FROM Test T JOIN Usuario U ON T.usuarioid=U.id WHERE U.username=? AND T.editable=0")

	if err == nil {
		defer query.Close()
		rows, err := query.Query(username)
		if err == nil {
			ts, err = rowsToTests(rows)
			return ts, err
		}
	}
	return nil, err
}

func GetPublicPublishedTestsFromUser(db *sql.DB, username string) ([]*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var ts []*Test
	query, err := db.Prepare("SELECT T.* FROM Test T JOIN Usuario U ON T.usuarioid=U.id WHERE U.username=? AND T.editable=0 AND T.accesoPublico=1")

	if err == nil {
		defer query.Close()
		rows, err := query.Query(username)
		if err == nil {
			ts, err = rowsToTests(rows)
			return ts, err
		}
	}
	return nil, err
}

func GetSolvableTestsFromUser(db *sql.DB, username string) ([]*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var ts []*Test
	query, err := db.Prepare(
		"WITH TestsUserInvited AS ( SELECT T.* FROM Test T JOIN InvitacionTestUsuario I ON T.id=I.testid JOIN Usuario U ON U.id=I.usuarioid WHERE U.username=? AND T.editable=0)," +
			"TestsUserTeamInvited AS ( SELECT T.* FROM Test T JOIN InvitacionTestEquipo I ON T.id=I.testid JOIN EquipoUsuario E ON E.equipoid=I.equipoid JOIN Usuario U ON U.id=E.usuarioid WHERE U.username=? AND T.editable=0)" +
			"SELECT DISTINCT T.* FROM Test T LEFT JOIN TestsUserInvited U ON T.id=U.id " +
			"LEFT JOIN TestsUserTeamInvited E ON E.id=T.id " +
			"WHERE T.editable=0 AND (T.accesoPublico==1 OR U.accesoPublico==0 OR E.accesoPublico==0)")

	if err == nil {
		defer query.Close()
		rows, err := query.Query(username, username)
		if err == nil {
			ts, err = rowsToTests(rows)
			return ts, err
		}
	}
	return nil, err
}

func GetSolvableTestFromUser(db *sql.DB, username string, testid int64) (*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var ts *Test
	query, err := db.Prepare(
		"WITH TestsUserInvited AS ( SELECT T.* FROM Test T JOIN InvitacionTestUsuario I ON T.id=I.testid JOIN Usuario U ON U.id=I.usuarioid WHERE U.username=? AND T.id=? AND T.editable=0)," +
			"TestsUserTeamInvited AS ( SELECT T.* FROM Test T JOIN InvitacionTestEquipo I ON T.id=I.testid JOIN EquipoUsuario E ON E.equipoid=I.equipoid JOIN Usuario U ON U.id=E.usuarioid WHERE U.username=? AND T.id=? AND T.editable=0)" +
			"SELECT DISTINCT T.* FROM Test T LEFT JOIN TestsUserInvited U ON T.id=U.id " +
			"LEFT JOIN TestsUserTeamInvited E ON E.id=T.id " +
			"WHERE T.id=? AND T.editable=0 AND (T.accesoPublico==1 OR U.accesoPublico==0 OR E.accesoPublico==0)")

	if err == nil {
		defer query.Close()
		rows, err := query.Query(username, testid, username, testid, testid)
		if err == nil {
			ts, err = rowsToTest(rows)
			return ts, err
		}
	}
	return nil, err
}

// AnsweredTests

func GetATestsFromUser(db *sql.DB, username string) ([]*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err == nil {
		var ts []*Test
		query, err := db.Prepare("SELECT T.* FROM Test T JOIN RespuestaExamen R ON T.id=R.testid WHERE R.usuarioid=?")

		if err == nil {
			defer query.Close()
			rows, err := query.Query(u.ID)
			if err == nil {
				ts, err = rowsToTests(rows)
				return ts, err
			}
		}
	}
	return nil, err
}

func GetATestFromUser(db *sql.DB, username string, testid int64) (*Test, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err == nil {
		var ts *Test
		query, err := db.Prepare("SELECT T.* FROM Test T JOIN RespuestaExamen R ON T.id=R.testid WHERE R.usuarioid=? AND R.testid=?")
		if err == nil {
			defer query.Close()
			rows, err := query.Query(u.ID, testid)
			if err == nil {
				ts, err = rowsToTest(rows)
				return ts, err
			}
		}
	}
	return nil, err
}
