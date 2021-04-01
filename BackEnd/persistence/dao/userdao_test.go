// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

package dao

import (
	"database/sql/driver"
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
	studentRol := models.UserRolStudent
	u := &User{
		Username: su.Username,
		Email:    su.Email,
		Pwhash:   &pwhashstring,
		Rol:      studentRol,
		Fullname: *su.Username,
	}
	return u
}

func expectInsert(mock sqlmock.Sqlmock, u *User) {
	mock.ExpectPrepare("INSERT INTO Usuario").ExpectExec().
		WithArgs(u.Username, u.Email, u.Pwhash, u.Rol, u.Fullname).WillReturnResult(sqlmock.NewResult(1, 1))
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
	mock.ExpectPrepare("INSERT INTO Usuario").ExpectExec().
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
var rolTeacher string = models.UserRolTeacher
var userid int64 = 1

func rowsUsers(us []*User) *sqlmock.Rows {
	columns := []string{"id", "username", "email", "pwhash", "rol", "fullname"}
	sqlcols := sqlmock.NewRows(columns)
	for _, user := range us {
		sqlcols.AddRow(user.ID, *user.Username, user.Email.String(), *user.Pwhash, user.Rol, user.Fullname)
	}
	return sqlcols
}

func rowsUser(u *User) *sqlmock.Rows {
	users := []*User{u}
	return rowsUsers(users)
}

func expectGetUser(mock sqlmock.Sqlmock, arg driver.Value, user *User) {
	rows := rowsUser(user)
	mock.ExpectPrepare("SELECT (.+) FROM Usuario").ExpectQuery().WithArgs(arg).
		WillReturnRows(rows)
}

func defaultUser() *User {
	return &User{
		ID:       1,
		Username: &username,
		Email:    &email,
		Pwhash:   &pwhash,
		Rol:      rolTeacher,
		Fullname: username,
	}
}

func expectGetEmail(mock sqlmock.Sqlmock) {
	expectGetUser(mock, emailS, defaultUser())
}

func expectGetUsername(mock sqlmock.Sqlmock) {
	expectGetUser(mock, username, defaultUser())
}

func expectGetUserID(mock sqlmock.Sqlmock) {
	expectGetUser(mock, userid, defaultUser())
}

func expectGetUserEmpty(mock sqlmock.Sqlmock, arg driver.Value) {
	rows := rowsUsers([]*User{})
	mock.ExpectPrepare("SELECT (.+) FROM Usuario").ExpectQuery().WithArgs(arg).
		WillReturnRows(rows)
}

func expectGetEmailEmpty(mock sqlmock.Sqlmock) {
	expectGetUserEmpty(mock, emailS)
}

func expectGetUsernameEmpty(mock sqlmock.Sqlmock) {
	expectGetUserEmpty(mock, username)
}

func expectGetUserIDEmpty(mock sqlmock.Sqlmock) {
	expectGetUserEmpty(mock, userid)
}

func expectGetUserTablesWrong(mock sqlmock.Sqlmock, arg string, u string, e string) {
	columns := []string{"id", "username", "email"}
	mock.ExpectPrepare("SELECT (.+) FROM Usuario").ExpectQuery().WithArgs(arg).
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
	mock.ExpectPrepare("SELECT (.+) FROM Usuario").ExpectQuery().WithArgs(username).
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
	mock.ExpectPrepare("SELECT (.+) FROM Usuario").ExpectQuery().WithArgs(email).
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

// GetUserByID

func TestGetUserByIDNilDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetUserID(mock)
	u, err := GetUserByID(nil, userid)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if u != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetUserByIDClosedDb(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db.Close()
	expectGetUserID(mock)
	u, err := GetUserByID(nil, userid)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if u != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetUserByIDError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectPrepare("SELECT (.+) FROM Usuario").ExpectQuery().WithArgs(userid).
		WillReturnError(fmt.Errorf("Error"))
	u, err := GetUserByID(nil, userid)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if u != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetUserByIDEmpty(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetUserIDEmpty(mock)
	u, err := GetUserByID(nil, userid)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if u != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetUserByIDNilFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetUserID(mock)
	u, err := GetUserByID(db, userid)
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	} else if u == nil {
		t.Log("u should not be nil")
		t.Fail()
	}
}

// GetUsers

var username2 string = "Test2"
var email2 strfmt.Email = "Test2@mail.com"
var pwhash2 string = "AAAAAAAAAAA"
var rolAdmin string = models.UserRolAdmin

func defaultAdmin1() *User {
	u := &User{
		ID:       2,
		Username: &username2,
		Email:    &email2,
		Pwhash:   &pwhash2,
		Rol:      rolAdmin,
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
		Rol:      rolTeacher,
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
	rows := rowsUsers(us)
	mock.ExpectPrepare("SELECT (.+) FROM Usuario").ExpectQuery().WillReturnRows(rows)
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
	mock.ExpectPrepare("SELECT (.+) FROM Usuario").ExpectQuery().WillReturnRows(rows)
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
	mock.ExpectPrepare("SELECT (.+) FROM Usuario").ExpectQuery().
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
	rows := rowsUsers(us)
	for _, u := range us {
		rows.AddRow(u.ID, u.Username, u.Email, u.Pwhash, u.Rol, u.Fullname)
	}
	mock.ExpectPrepare("SELECT (.+) FROM Usuario WHERE rol='Admin'").ExpectQuery().WillReturnRows(rows)
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
	mock.ExpectPrepare("SELECT (.+) FROM Usuario WHERE rol='Admin'").ExpectQuery().WillReturnRows(rows)
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
	mock.ExpectPrepare("SELECT (.+) FROM Usuario WHERE rol='Admin'").ExpectQuery().
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

func expectUpdateUser(mock sqlmock.Sqlmock, u *User) {
	mock.ExpectPrepare("UPDATE Users").ExpectExec().
		WithArgs(u.Username, u.Email, u.Fullname, u.Rol, u.Username).WillReturnResult(sqlmock.NewResult(1, 1))
}

func expectUpdateUserDefault(mock sqlmock.Sqlmock) {
	u := defaultAdmin1()
	expectUpdateUser(mock, u)
}

func TestUpdateUserNilDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectUpdateUserDefault(mock)
	u := ToModelUser(defaultAdmin1())
	err = UpdateUser(nil, u, *u.Username)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestUpdateUserNilUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectUpdateUserDefault(mock)
	u := ToModelUser(defaultAdmin1())
	err = UpdateUser(db, nil, *u.Username)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestUpdateUserClosedDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db.Close()
	expectUpdateUserDefault(mock)
	u := ToModelUser(defaultAdmin1())
	err = UpdateUser(db, u, *u.Username)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestUpdateUserError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	u := ToModelUser(defaultAdmin1())
	mock.ExpectPrepare("UPDATE Users").ExpectExec().
		WithArgs(u.Username, u.Email, u.Fullname, u.Rol, username).WillReturnError(fmt.Errorf("Error"))
	err = UpdateUser(db, u, *u.Username)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestUpdateUserCorrect(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectUpdateUserDefault(mock)
	u := ToModelUser(defaultAdmin1())
	err = UpdateUser(db, u, *u.Username)
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}

// DeleteUser

var usernameDelete string = "username"

func expectDeleteUser(mock sqlmock.Sqlmock, username string) {
	mock.ExpectPrepare("DELETE FROM Usuario").ExpectExec().
		WithArgs(username).WillReturnResult(sqlmock.NewResult(1, 1))
}

func expectDeleteUserDefault(mock sqlmock.Sqlmock) {
	expectDeleteUser(mock, usernameDelete)
}

func TestDeleteUserNilDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectDeleteUserDefault(mock)
	err = DeleteUser(nil, usernameDelete)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestDeleteUserClosedDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db.Close()
	expectDeleteUserDefault(mock)
	err = DeleteUser(db, usernameDelete)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestDeleteUserError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectPrepare("DELETE Users").ExpectExec().
		WithArgs(username).WillReturnError(fmt.Errorf("Error"))
	err = DeleteUser(db, usernameDelete)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestDeleteUserCorrect(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectDeleteUserDefault(mock)
	err = DeleteUser(db, usernameDelete)
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}

// AddUserTeamMember
var teamname string = "teamname"
var useridAddUser int = 1
var teamidAddUser int = 1

func defaultTeam() *Team {
	t := &Team{
		ID:          1,
		Teamname:    &teamname,
		Description: "description",
	}
	return t
}

func rowsTeams(ts []*Team) *sqlmock.Rows {
	columns := []string{"id", "teamname", "description"}
	sqlcols := sqlmock.NewRows(columns)
	for _, team := range ts {
		sqlcols.AddRow(team.ID, team.Teamname, team.Description)
	}
	return sqlcols
}

func rowsTeam(t *Team) *sqlmock.Rows {
	ts := []*Team{t}
	return rowsTeams(ts)
}

func rowsTeamEmpty() *sqlmock.Rows {
	return rowsTeams([]*Team{})
}

func expectGetTeam(mock sqlmock.Sqlmock, arg string, team *Team) {
	rows := rowsTeam(team)
	mock.ExpectPrepare("SELECT (.+) FROM Equipo").ExpectQuery().WithArgs(arg).
		WillReturnRows(rows)
}

func expectGetTeamEmpty(mock sqlmock.Sqlmock, arg string) {
	rows := rowsTeamEmpty()
	mock.ExpectPrepare("SELECT (.+) FROM Equipo").ExpectQuery().WithArgs(arg).
		WillReturnRows(rows)
}

func expectAddUserTeamMember(mock sqlmock.Sqlmock, userid int, teamid int) {
	expectGetUser(mock, username, defaultUser())
	expectGetTeam(mock, teamname, defaultTeam())
	mock.ExpectPrepare("INSERT INTO EquipoUsuario").ExpectExec().
		WithArgs(userid, teamid, models.TeamRoleRoleMember).WillReturnResult(sqlmock.NewResult(1, 1))
}

func expectAddUserTeamMemberDefault(mock sqlmock.Sqlmock) {
	expectAddUserTeamMember(mock, useridAddUser, teamidAddUser)
}

func expectAddUserTeamMemberUsernameNotFound(mock sqlmock.Sqlmock) {
	expectGetUserEmpty(mock, username)
	expectGetTeam(mock, teamname, defaultTeam())
	mock.ExpectPrepare("INSERT INTO EquipoUsuario").ExpectExec().
		WithArgs(useridAddUser, teamidAddUser, models.TeamRoleRoleMember).WillReturnResult(sqlmock.NewResult(1, 1))
}

func expectAddUserTeamMemberTeamnameNotFound(mock sqlmock.Sqlmock) {
	expectGetUser(mock, username, defaultUser())
	expectGetTeamEmpty(mock, teamname)
	mock.ExpectPrepare("INSERT INTO EquipoUsuario").ExpectExec().
		WithArgs(useridAddUser, teamidAddUser, models.TeamRoleRoleMember).WillReturnResult(sqlmock.NewResult(1, 1))
}

func TestAddUserTeamMemberNilDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectAddUserTeamMemberDefault(mock)
	err = AddUserTeamMember(nil, username, teamname)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestAddUserTeamMemberUsernameNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectAddUserTeamMemberUsernameNotFound(mock)
	err = AddUserTeamMember(db, username, teamname)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestAddUserTeamMemberTeamnameNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectAddUserTeamMemberTeamnameNotFound(mock)
	err = AddUserTeamMember(db, username, teamname)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestAddUserTeamMemberError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetUser(mock, username, defaultUser())
	expectGetTeam(mock, teamname, defaultTeam())
	mock.ExpectPrepare("INSERT INTO EquipoUsuario").ExpectExec().
		WithArgs(useridAddUser, teamidAddUser, models.TeamRoleRoleMember).WillReturnError(fmt.Errorf("Error"))
	err = AddUserTeamMember(db, username, teamname)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestAddUserTeamMemberCorrect(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectAddUserTeamMemberDefault(mock)
	err = AddUserTeamMember(db, username, teamname)
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}

// AddUserTeamAdmin

func expectAddUserTeamAdmin(mock sqlmock.Sqlmock, userid int, teamid int) {
	expectGetUser(mock, username, defaultUser())
	expectGetTeam(mock, teamname, defaultTeam())
	mock.ExpectPrepare("INSERT INTO EquipoUsuario").ExpectExec().
		WithArgs(userid, teamid, models.TeamRoleRoleAdmin).WillReturnResult(sqlmock.NewResult(1, 1))
}

func expectAddUserTeamAdminDefault(mock sqlmock.Sqlmock) {
	expectAddUserTeamAdmin(mock, useridAddUser, teamidAddUser)
}

func expectAddUserTeamAdminUsernameNotFound(mock sqlmock.Sqlmock) {
	expectGetUserEmpty(mock, username)
	expectGetTeam(mock, teamname, defaultTeam())
	mock.ExpectPrepare("INSERT INTO EquipoUsuario").ExpectExec().
		WithArgs(useridAddUser, teamidAddUser, models.TeamRoleRoleAdmin).WillReturnResult(sqlmock.NewResult(1, 1))
}

func expectAddUserTeamAdminTeamnameNotFound(mock sqlmock.Sqlmock) {
	expectGetUser(mock, username, defaultUser())
	expectGetTeamEmpty(mock, teamname)
	mock.ExpectPrepare("INSERT INTO EquipoUsuario").ExpectExec().
		WithArgs(useridAddUser, teamidAddUser, models.TeamRoleRoleAdmin).WillReturnResult(sqlmock.NewResult(1, 1))
}

func TestAddUserTeamAdminNilDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectAddUserTeamAdminDefault(mock)
	err = AddUserTeamAdmin(nil, username, teamname)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestAddUserTeamAdminUsernameNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectAddUserTeamAdminUsernameNotFound(mock)
	err = AddUserTeamAdmin(db, username, teamname)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestAddUserTeamAdminTeamnameNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectAddUserTeamAdminTeamnameNotFound(mock)
	err = AddUserTeamAdmin(db, username, teamname)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestAddUserTeamAdminError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetUser(mock, username, defaultUser())
	expectGetTeam(mock, teamname, defaultTeam())
	mock.ExpectPrepare("INSERT INTO EquipoUsuario").ExpectExec().
		WithArgs(useridAddUser, teamidAddUser, models.TeamRoleRoleAdmin).WillReturnError(fmt.Errorf("Error"))
	err = AddUserTeamAdmin(db, username, teamname)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestAddUserTeamAdminCorrect(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectAddUserTeamAdminDefault(mock)
	err = AddUserTeamAdmin(db, username, teamname)
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}

// ExitUserTeam

func expectExitUserTeam(mock sqlmock.Sqlmock, userid int, teamid int) {
	expectGetUser(mock, username, defaultUser())
	expectGetTeam(mock, teamname, defaultTeam())
	mock.ExpectPrepare("DELETE FROM EquipoUsuario").ExpectExec().
		WithArgs(userid, teamid).WillReturnResult(sqlmock.NewResult(1, 1))
}

func expectExitUserTeamDefault(mock sqlmock.Sqlmock) {
	expectExitUserTeam(mock, useridAddUser, teamidAddUser)
}

func expectExitUserTeamUsernameNotFound(mock sqlmock.Sqlmock) {
	expectGetUserEmpty(mock, username)
	expectGetTeam(mock, teamname, defaultTeam())
	mock.ExpectPrepare("DELETE FROM EquipoUsuario").ExpectExec().
		WithArgs(useridAddUser, teamidAddUser).WillReturnResult(sqlmock.NewResult(1, 1))
}

func expectExitUserTeamTeamnameNotFound(mock sqlmock.Sqlmock) {
	expectGetUser(mock, username, defaultUser())
	expectGetTeamEmpty(mock, teamname)
	mock.ExpectPrepare("DELETE FROM EquipoUsuario").ExpectExec().
		WithArgs(useridAddUser, teamidAddUser).WillReturnResult(sqlmock.NewResult(1, 1))
}

func expectExitUserTeamError(mock sqlmock.Sqlmock) {
	expectGetUser(mock, username, defaultUser())
	expectGetTeam(mock, teamname, defaultTeam())
	mock.ExpectPrepare("DELETE FROM EquipoUsuario").ExpectExec().
		WithArgs(useridAddUser, teamidAddUser).WillReturnError(fmt.Errorf("Error"))
}
func TestExitUserTeamNilDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectExitUserTeamDefault(mock)
	err = ExitUserTeam(nil, username, teamname)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestExitUserTeamClosedDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db.Close()
	expectExitUserTeamDefault(mock)
	err = ExitUserTeam(db, username, teamname)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestExitUserTeamUsernameNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectExitUserTeamUsernameNotFound(mock)
	err = ExitUserTeam(db, username, teamname)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestExitUserTeamTeamnameNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectExitUserTeamTeamnameNotFound(mock)
	err = ExitUserTeam(db, username, teamname)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestExitUserTeamError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectExitUserTeamError(mock)
	err = ExitUserTeam(db, username, teamname)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestExitUserTeamCorrect(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectExitUserTeamDefault(mock)
	err = ExitUserTeam(db, username, teamname)
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}

// GetUsersFromTeam

func expectGetUsersFromTeam(mock sqlmock.Sqlmock, teamname string, users []*User) {
	expectGetTeam(mock, teamname, defaultTeam())
	rows := rowsUsers(users)
	mock.ExpectPrepare("SELECT (.+) FROM Usuario").ExpectQuery().WithArgs(defaultTeam().ID).WillReturnRows(rows)
}

func expectGetUsersFromTeamDefault(mock sqlmock.Sqlmock) {
	expectGetUsersFromTeam(mock, teamname, defaultUsersFromTeam())
}

func defaultUsersFromTeam() []*User {
	return defaultUsers()
}

func expectGetUsersFromTeamTeamnameNotFound(mock sqlmock.Sqlmock) {
	expectGetTeamEmpty(mock, teamname)
	rows := rowsUsers([]*User{})
	mock.ExpectPrepare("SELECT (.+) FROM Usuario").ExpectQuery().WithArgs(defaultTeam().ID).WillReturnRows(rows)
}

func expectGetUsersFromTeamError(mock sqlmock.Sqlmock) {
	expectGetTeam(mock, teamname, defaultTeam())
	mock.ExpectPrepare("SELECT (.+) FROM Usuario").ExpectQuery().WithArgs(defaultTeam().ID).
		WillReturnError(fmt.Errorf("Error"))
}

func expectGetUsersFromTeamEmpty(mock sqlmock.Sqlmock) {
	expectGetTeam(mock, teamname, defaultTeam())
	rows := rowsUsers([]*User{})
	mock.ExpectPrepare("SELECT (.+) FROM Usuario").ExpectQuery().WithArgs(defaultTeam().ID).WillReturnRows(rows)
}

func TestGetUsersFromTeamNilDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetUsersFromTeamDefault(mock)
	u, err := GetUsersFromTeam(nil, teamname)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if u != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetUsersFromTeamClosedDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db.Close()
	expectGetUsersFromTeamDefault(mock)
	u, err := GetUsersFromTeam(db, teamname)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if u != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetUsersFromTeamTeamnameNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetUsersFromTeamTeamnameNotFound(mock)
	u, err := GetUsersFromTeam(db, teamname)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if u != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetUsersFromTeamError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetUsersFromTeamError(mock)
	u, err := GetUsersFromTeam(db, teamname)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if u != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetUsersFromTeamEmpty(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetUsersFromTeamEmpty(mock)
	u, err := GetUsersFromTeam(db, teamname)
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	} else if u != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetUsersFromTeamFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetUsersFromTeamDefault(mock)
	u, err := GetUsersFromTeam(db, teamname)
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	} else if u == nil {
		t.Log("u should not be nil")
		t.Fail()
	}
	us := defaultUsersFromTeam()
	if len(u) != len(us) || u[0].Email.String() != us[0].Email.String() {
		t.Log("emails incorrectos")
		t.Fail()
	}
}

// GetTeamAdmins

func expectGetTeamAdmins(mock sqlmock.Sqlmock, teamname string, users []*User) {
	expectGetTeam(mock, teamname, defaultTeam())
	rows := rowsUsers(users)
	mock.ExpectPrepare("SELECT (.+) FROM Usuario").ExpectQuery().WithArgs(defaultTeam().ID).WillReturnRows(rows)
}

func expectGetTeamAdminsDefault(mock sqlmock.Sqlmock) {
	expectGetTeamAdmins(mock, teamname, defaultTeamAdmins())
}

func expectGetTeamAdminsTeamnameNotFound(mock sqlmock.Sqlmock) {
	expectGetTeamEmpty(mock, teamname)
	rows := rowsUsers([]*User{})
	mock.ExpectPrepare("SELECT (.+) FROM Usuario").ExpectQuery().WithArgs(defaultTeam().ID).WillReturnRows(rows)
}

func expectGetTeamAdminsError(mock sqlmock.Sqlmock) {
	expectGetTeam(mock, teamname, defaultTeam())
	mock.ExpectPrepare("SELECT (.+) FROM Usuario").ExpectQuery().WithArgs(defaultTeam().ID).
		WillReturnError(fmt.Errorf("Error"))
}

func expectGetTeamAdminsEmpty(mock sqlmock.Sqlmock) {
	expectGetTeam(mock, teamname, defaultTeam())
	rows := rowsUsers([]*User{})
	mock.ExpectPrepare("SELECT (.+) FROM Usuario").ExpectQuery().WithArgs(defaultTeam().ID).WillReturnRows(rows)
}

func defaultTeamAdmins() []*User {
	return defaultAdmins()
}

func TestGetTeamAdminsNilDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetTeamAdminsDefault(mock)
	u, err := GetTeamAdmins(nil, teamname)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if u != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetTeamAdminsClosedDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db.Close()
	expectGetTeamAdminsDefault(mock)
	u, err := GetTeamAdmins(db, teamname)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if u != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetTeamAdminsTeamnameNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetTeamAdminsTeamnameNotFound(mock)
	u, err := GetTeamAdmins(db, teamname)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if u != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetTeamAdminsError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetTeamAdminsError(mock)
	u, err := GetTeamAdmins(db, teamname)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	} else if u != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetTeamAdminsEmpty(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetTeamAdminsEmpty(mock)
	u, err := GetTeamAdmins(db, teamname)
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	} else if u != nil {
		t.Log("u should be nil")
		t.Fail()
	}
}

func TestGetTeamAdminsFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetTeamAdminsDefault(mock)
	u, err := GetTeamAdmins(db, teamname)
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	} else if u == nil {
		t.Log("u should not be nil")
		t.Fail()
	}
	us := defaultTeamAdmins()
	if len(u) != len(us) || u[0].Email.String() != us[0].Email.String() {
		t.Log("emails incorrectos")
		t.Fail()
	}
}
