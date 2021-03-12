// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

package dao

import (
	"fmt"
	"testing"
	"uva-devtest/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-openapi/strfmt"
	"golang.org/x/crypto/bcrypt"
)

// Testing InsertUser

func signinUserToInsert() *models.SigninUser {
	uname := "Test"
	var email strfmt.Email = "Test@mail.com"
	var pass strfmt.Password = "Testpass"
	u := &models.SigninUser{
		Username: &uname,
		Email:    &email,
		Pass:     &pass,
	}
	return u
}

func userToInsert() *User {
	su := signinUserToInsert()
	bytes, _ := bcrypt.GenerateFromPassword([]byte(*su.Pass), 14)
	pwhashstring := string(bytes)
	studentType := models.UserTypeStudent
	u := &User{
		Username: su.Username,
		Email:    su.Email,
		Pwhash:   &pwhashstring,
		Type:     studentType,
	}
	return u
}

func expectInsert(mock sqlmock.Sqlmock, u *User) {
	mock.ExpectPrepare("INSERT INTO Users").ExpectExec().
		WithArgs(u.Username, u.Email, u.Pwhash, u.Type).WillReturnResult(sqlmock.NewResult(1, 1))
}

func TestInsertUserNilDB(t *testing.T) {
	u := userToInsert()
	err := InsertUser(nil, u)
	if err == nil {
		t.Log("error should not be nil")
		t.Fail()
	}
}

func TestInsertUserNilUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	u := userToInsert()
	expectInsert(mock, u)
	err = InsertUser(db, nil)
	if err == nil {
		t.Log("error should not be nil")
		t.Fail()
	}
}

func TestInsertUserWrongDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	u := userToInsert()
	expectInsert(mock, u)
	db.Close()
	err = InsertUser(db, u)
	if err == nil {
		t.Log("error should not be nil")
		t.Fail()
	}
}

func TestInsertUserRepeated(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	u := userToInsert()
	mock.ExpectPrepare("INSERT INTO users").ExpectExec().
		WithArgs(u.Username, u.Email, u.Pwhash).WillReturnError(fmt.Errorf("Dato repetido"))
	err = InsertUser(db, u)
	if err == nil {
		t.Log("error should be nil")
		t.Fail()
	}
}

func TestInsertUserCorrect(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	u := userToInsert()
	expectInsert(mock, u)
	err = InsertUser(db, u)
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}

// Testing getUserUsername y getUserEmail

var username string = "Test"
var email strfmt.Email = "Test@mail.com"
var pwhash string = "AAAAAAAAAAA"
var emailS = email.String()
var typeTeacher string = models.UserTypeTeacher

func expectGetUser(mock sqlmock.Sqlmock, arg string, u string, e string, p string, t string) {
	columns := []string{"id", "username", "email", "pwhash", "type"}
	mock.ExpectPrepare("SELECT (.+) FROM Users").ExpectQuery().WithArgs(arg).
		WillReturnRows(sqlmock.NewRows(columns).AddRow(1, u, e, p, t))
}

func expectGetEmail(mock sqlmock.Sqlmock) {
	expectGetUser(mock, emailS, username, emailS, pwhash, typeTeacher)
}

func expectGetUsername(mock sqlmock.Sqlmock) {
	expectGetUser(mock, username, username, emailS, pwhash, typeTeacher)
}

func expectGetUserEmpty(mock sqlmock.Sqlmock, arg string) {
	columns := []string{"id", "username", "email", "pwhash", "type"}
	mock.ExpectPrepare("SELECT (.+) FROM Users").ExpectQuery().WithArgs(arg).
		WillReturnRows(sqlmock.NewRows(columns))
}

func expectGetEmailEmpty(mock sqlmock.Sqlmock) {
	expectGetUserEmpty(mock, emailS)
}

func expectGetUsernameEmpty(mock sqlmock.Sqlmock) {
	expectGetUserEmpty(mock, username)
}

func expectGetUserTablesWrong(mock sqlmock.Sqlmock, arg string, u string, e string) {
	columns := []string{"id", "username", "email"}
	mock.ExpectPrepare("SELECT (.+) FROM users").ExpectQuery().WithArgs(arg).
		WillReturnRows(sqlmock.NewRows(columns).AddRow(1, u, e))
}

func expectGetUsernameTablesWrong(mock sqlmock.Sqlmock) {
	expectGetUserTablesWrong(mock, username, username, emailS)
}

func expectGetEmailTablesWrong(mock sqlmock.Sqlmock) {
	expectGetUserTablesWrong(mock, emailS, username, emailS)
}

// Testing GetUserUsername

func TestGetUserUsernameNilDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetUsername(mock)
	u, err := GetUserUsername(nil, username)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if u != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetUserUsernameClosedDb(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db.Close()
	expectGetUsername(mock)
	u, err := GetUserUsername(db, username)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if u != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetUserUsernameError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectPrepare("SELECT (.+) FROM users").ExpectQuery().WithArgs(username).
		WillReturnError(fmt.Errorf("Error"))
	u, err := GetUserUsername(db, username)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if u != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetUserUsernameEmpty(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetUsernameEmpty(mock)
	u, err := GetUserUsername(db, username)
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	} else if u != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetUserUsernameTablesWrong(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetUsernameTablesWrong(mock)
	u, err := GetUserUsername(db, username)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if u != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetUserUsernameFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetUsername(mock)
	u, err := GetUserUsername(db, username)
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	} else if u == nil {
		t.Log("u should not be nil")
		t.Fail()
	}
	if u.Email.String() != emailS {
		t.Log("email incorrect")
		t.Fail()
	}
}

// Testing GetUserEmail

func TestGetUserEmailNilDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetEmail(mock)
	u, err := GetUserEmail(nil, emailS)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if u != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetUserEmailClosedDb(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db.Close()
	expectGetEmail(mock)
	u, err := GetUserEmail(db, emailS)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if u != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetUserEmailError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectPrepare("SELECT (.+) FROM users").ExpectQuery().WithArgs(email).
		WillReturnError(fmt.Errorf("Error"))
	u, err := GetUserEmail(db, emailS)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if u != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetUserEmailEmpty(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetEmailEmpty(mock)
	u, err := GetUserEmail(db, emailS)
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	} else if u != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetUserEmailTablesWrong(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetEmailTablesWrong(mock)
	u, err := GetUserEmail(db, emailS)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if u != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetUserEmailFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetEmail(mock)
	u, err := GetUserEmail(db, emailS)
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	} else if u == nil {
		t.Log("u should not be nil")
		t.Fail()
	}
	if *u.Username != username {
		t.Log("username incorrect")
		t.Fail()
	}
}

// GetUsers

var username2 string = "Test2"
var email2 strfmt.Email = "Test2@mail.com"
var pwhash2 string = "AAAAAAAAAAA"
var typeAdmin string = models.UserTypeAdmin

func defaultAdmin1() *User {
	u := &User{
		ID:       2,
		Username: &username2,
		Email:    &email2,
		Pwhash:   &pwhash2,
		Type:     typeAdmin,
	}
	return u
}

func defaultUsers() []*User {
	us := []*User{}
	u1 := &User{
		ID:       1,
		Username: &username,
		Email:    &email,
		Pwhash:   &pwhash,
		Type:     typeTeacher,
	}
	u2 := defaultAdmin1()
	us = append(us, u1)
	us = append(us, u2)
	return us
}

func defaultAdmins() []*User {
	us := []*User{}
	u2 := defaultAdmin1()
	us = append(us, u2)
	return us
}

func expectGetUsers(mock sqlmock.Sqlmock, us []*User) {
	columns := []string{"id", "username", "email", "pwhash", "type"}
	rows := sqlmock.NewRows(columns)
	for _, u := range us {
		rows.AddRow(u.ID, u.Username, u.Email, u.Pwhash, u.Type)
	}
	mock.ExpectPrepare("SELECT (.+) FROM Users").ExpectQuery().WillReturnRows(rows)
}

func expectGetUsersDefault(mock sqlmock.Sqlmock) {
	us := defaultUsers()
	expectGetUsers(mock, us)
}

func expectGetUsersEmpty(mock sqlmock.Sqlmock) {
	us := []*User{}
	expectGetUsers(mock, us)
}

func expectGetUsersWrong(mock sqlmock.Sqlmock) {
	columns := []string{"id", "username", "email"}
	us := defaultUsers()
	rows := sqlmock.NewRows(columns)
	for _, u := range us {
		rows.AddRow(u.ID, u.Username, u.Email)
	}
	mock.ExpectPrepare("SELECT (.+) FROM Users").ExpectQuery().WillReturnRows(rows)
}

func TestGetUsersNilDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetUsersDefault(mock)
	u, err := GetUsers(nil)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if u != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetUsersClosedDb(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db.Close()
	expectGetUsersDefault(mock)
	u, err := GetUsers(db)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if u != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetUsersError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectPrepare("SELECT (.+) FROM Users").ExpectQuery().
		WillReturnError(fmt.Errorf("Error"))
	u, err := GetUsers(db)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if u != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetUsersEmpty(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetUsersEmpty(mock)
	u, err := GetUsers(db)
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	} else if u != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetUsersTablesWrong(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetUsersWrong(mock)
	u, err := GetUsers(db)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if u != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetUsersFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetUsersDefault(mock)
	u, err := GetUsers(db)
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	} else if u == nil {
		t.Log("u should not be nil")
		t.Fail()
	}
	us := defaultUsers()
	if u[0].Email.String() != us[0].Email.String() || u[1].Email.String() != us[1].Email.String() {
		t.Log("emails incorrectos")
		t.Fail()
	}
}

// GetAdmins

func expectGetAdmins(mock sqlmock.Sqlmock, us []*User) {
	columns := []string{"id", "username", "email", "pwhash", "type"}
	rows := sqlmock.NewRows(columns)
	for _, u := range us {
		rows.AddRow(u.ID, u.Username, u.Email, u.Pwhash, u.Type)
	}
	mock.ExpectPrepare("SELECT (.+) FROM Users WHERE type='Admin'").ExpectQuery().WillReturnRows(rows)
}

func expectGetAdminsDefault(mock sqlmock.Sqlmock) {
	us := defaultAdmins()
	expectGetAdmins(mock, us)
}

func expectGetAdminsEmpty(mock sqlmock.Sqlmock) {
	us := []*User{}
	expectGetAdmins(mock, us)
}

func expectGetAdminsWrong(mock sqlmock.Sqlmock) {
	columns := []string{"id", "username", "email"}
	rows := sqlmock.NewRows(columns)
	us := defaultAdmins()
	for _, u := range us {
		rows.AddRow(u.ID, u.Username, u.Email)
	}
	mock.ExpectPrepare("SELECT (.+) FROM Users WHERE type='Admin'").ExpectQuery().WillReturnRows(rows)
}

func TestGetAdminsNilDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetAdminsDefault(mock)
	u, err := GetAdmins(nil)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if u != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetAdminsClosedDb(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db.Close()
	expectGetAdminsDefault(mock)
	u, err := GetAdmins(db)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if u != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetAdminsError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectPrepare("SELECT (.+) FROM Users WHERE type='Admin'").ExpectQuery().
		WillReturnError(fmt.Errorf("Error"))
	u, err := GetAdmins(db)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if u != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetAdminsEmpty(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetAdminsEmpty(mock)
	u, err := GetAdmins(db)
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	} else if u != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetAdminsTablesWrong(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetAdminsWrong(mock)
	u, err := GetAdmins(db)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if u != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetAdminsFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetAdminsDefault(mock)
	u, err := GetAdmins(db)
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	} else if u == nil {
		t.Log("u should not be nil")
		t.Fail()
	}
	us := defaultAdmins()
	if u[0].Email.String() != us[0].Email.String() {
		t.Log("emails incorrectos")
		t.Fail()
	}
}

// PutPasswordUsername

const usernamePp string = "Test"
const newPp string = "AAAAAAAAAAAAA"

func expectPutPassword(mock sqlmock.Sqlmock, p string, username string) {
	mock.ExpectPrepare("UPDATE Users").ExpectExec().
		WithArgs(p, username).WillReturnResult(sqlmock.NewResult(1, 1))
}

func expectPutPasswordDefault(mock sqlmock.Sqlmock) {
	expectPutPassword(mock, newPp, usernamePp)
}
func TestPutPasswordNilDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectPutPasswordDefault(mock)
	err = PutPasswordUsername(nil, usernamePp, newPp)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestPutPasswordClosedDb(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db.Close()
	expectPutPasswordDefault(mock)
	err = PutPasswordUsername(db, usernamePp, newPp)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestPutPasswordError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectPrepare("UPDATE Users").ExpectExec().
		WithArgs(newPp, usernamePp).WillReturnError(fmt.Errorf("Error"))
	err = PutPasswordUsername(db, usernamePp, newPp)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestPutPasswordCorrect(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectPutPasswordDefault(mock)
	err = PutPasswordUsername(db, usernamePp, newPp)
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}

// UpdateUser

func TestUpdateUserNilDB(t *testing.T) {

}

func TestUpdateUserNilUser(t *testing.T) {

}

func TestUpdateUserClosedDB(t *testing.T) {

}

func TestUpdateUserError(t *testing.T) {

}

func TestUpdateUserCorrect(t *testing.T) {

}

// DeleteUser

func TestDeleteUserNilDB(t *testing.T) {

}

func TestDeleteUserClosedDB(t *testing.T) {

}

func TestDeleteUserError(t *testing.T) {

}

func TestDeleteUserCorrect(t *testing.T) {

}

// AddUserTeam

func TestAddUserTeamNilDB(t *testing.T) {

}

func TestAddUserTeamClosedDB(t *testing.T) {

}

func TestAddUserTeamUsernameError(t *testing.T) {

}

func TestAddUserTeamTeamnameError(t *testing.T) {

}

func TestAddUserTeamError(t *testing.T) {

}

func TestAddUserTeamCorrect(t *testing.T) {

}

// ExitUserTeam

func TestExitUserTeamNilDB(t *testing.T) {

}

func TestExitUserTeamClosedDB(t *testing.T) {

}

func TestExitUserTeamUsernameError(t *testing.T) {

}

func TestExitUserTeamTeamnameError(t *testing.T) {

}

func TestExitUserTeamError(t *testing.T) {

}

func TestExitUserTeamCorrect(t *testing.T) {

}

// GetUsersFromTeam

func TestGetUsersFromTeamNilDB(t *testing.T) {
}

func TestGetUsersFromTeamClosedDB(t *testing.T) {
}

func TestGetUsersFromTeamTeamnameError(t *testing.T) {
}

func TestGetUsersFromTeamError(t *testing.T) {
}

func TestGetUsersFromTeamEmpty(t *testing.T) {
}

func TestGetUsersFromTeamFound(t *testing.T) {
}

// GetTeamAdmins

func TestGetTeamAdminsNilDB(t *testing.T) {
}

func TestGetTeamAdminsClosedDB(t *testing.T) {
}

func TestGetTeamAdminsTeamnameError(t *testing.T) {
}

func TestGetTeamAdminsError(t *testing.T) {
}

func TestGetTeamAdminsEmpty(t *testing.T) {
}

func TestGetTeamAdminsFound(t *testing.T) {
}
