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
		Fullname: *su.Username,
	}
	return u
}

func expectInsert(mock sqlmock.Sqlmock, u *User) {
	mock.ExpectPrepare("INSERT INTO Users").ExpectExec().
		WithArgs(u.Username, u.Email, u.Pwhash, u.Type, u.Fullname).WillReturnResult(sqlmock.NewResult(1, 1))
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

func expectUpdateUser(mock sqlmock.Sqlmock, u *User) {
	mock.ExpectPrepare("UPDATE Users").ExpectExec().
		WithArgs(u.Username, u.Email, u.Fullname, u.Type, u.Username).WillReturnResult(sqlmock.NewResult(1, 1))
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
		WithArgs(u.Username, u.Email, u.Fullname, u.Type, username).WillReturnError(fmt.Errorf("Error"))
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
	mock.ExpectPrepare("DELETE FROM Users").ExpectExec().
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

// AddUserTeam
var usernameAddUser string = "username"
var teamnameAddUser string = "teamname"
var useridAddUser int = 1
var teamidAddUser int = 1

func defaultTeam() *Team {
	t := &Team{
		ID:          1,
		Teamname:    &teamnameAddUser,
		Description: "description",
	}
	return t
}

func expectGetTeam(mock sqlmock.Sqlmock, arg string, team *Team) {
	columns := []string{"id", "teamname", "description"}
	mock.ExpectPrepare("SELECT (.+) FROM Teams").ExpectQuery().WithArgs(arg).
		WillReturnRows(sqlmock.NewRows(columns).AddRow(team.ID, team.Teamname, team.Description))
}

func expectGetTeamEmpty(mock sqlmock.Sqlmock, arg string) {
	columns := []string{"id", "teamname", "description"}
	mock.ExpectPrepare("SELECT (.+) FROM Teams").ExpectQuery().WithArgs(arg).
		WillReturnRows(sqlmock.NewRows(columns))
}

func expectAddUserTeam(mock sqlmock.Sqlmock, userid int, teamid int) {
	expectGetUser(mock, usernameAddUser, usernameAddUser, emailS, pwhash, typeTeacher)
	expectGetTeam(mock, teamnameAddUser, defaultTeam())
	mock.ExpectPrepare("INSERT INTO Teamroles").ExpectExec().
		WithArgs(userid, teamid, models.TeamRoleRoleMember).WillReturnResult(sqlmock.NewResult(1, 1))
}

func expectAddUserTeamDefault(mock sqlmock.Sqlmock) {
	expectAddUserTeam(mock, useridAddUser, teamidAddUser)
}

func expectAddUserTeamUsernameNotFound(mock sqlmock.Sqlmock) {
	expectGetUserEmpty(mock, usernameAddUser)
	expectGetTeam(mock, teamnameAddUser, defaultTeam())
	mock.ExpectPrepare("INSERT INTO Teamroles").ExpectExec().
		WithArgs(useridAddUser, teamidAddUser, models.TeamRoleRoleMember).WillReturnResult(sqlmock.NewResult(1, 1))
}

func expectAddUserTeamTeamnameNotFound(mock sqlmock.Sqlmock) {
	expectGetUser(mock, usernameAddUser, usernameAddUser, emailS, pwhash, typeTeacher)
	expectGetTeamEmpty(mock, teamnameAddUser)
	mock.ExpectPrepare("INSERT INTO Teamroles").ExpectExec().
		WithArgs(useridAddUser, teamidAddUser, models.TeamRoleRoleMember).WillReturnResult(sqlmock.NewResult(1, 1))
}

func TestAddUserTeamNilDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectAddUserTeamDefault(mock)
	err = AddUserTeam(nil, usernameAddUser, teamnameAddUser)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestAddUserTeamClosedDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db.Close()
	expectAddUserTeamDefault(mock)
	err = AddUserTeam(db, usernameAddUser, teamnameAddUser)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestAddUserTeamUsernameNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectAddUserTeamUsernameNotFound(mock)
	err = AddUserTeam(db, usernameAddUser, teamnameAddUser)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestAddUserTeamTeamnameNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectAddUserTeamTeamnameNotFound(mock)
	err = AddUserTeam(db, usernameAddUser, teamnameAddUser)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestAddUserTeamError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetUser(mock, usernameAddUser, usernameAddUser, emailS, pwhash, typeTeacher)
	expectGetTeam(mock, teamnameAddUser, defaultTeam())
	mock.ExpectPrepare("INSERT INTO Teamroles").ExpectExec().
		WithArgs(useridAddUser, teamidAddUser, models.TeamRoleRoleMember).WillReturnError(fmt.Errorf("Error"))
	err = AddUserTeam(db, usernameAddUser, teamnameAddUser)
	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestAddUserTeamCorrect(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectAddUserTeamDefault(mock)
	err = AddUserTeam(db, usernameAddUser, teamnameAddUser)
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}

// ExitUserTeam

func expectExitUserTeam(mock sqlmock.Sqlmock, userid int, teamid int) {
	expectGetUser(mock, usernameAddUser, usernameAddUser, emailS, pwhash, typeTeacher)
	expectGetTeam(mock, teamnameAddUser, defaultTeam())
	mock.ExpectPrepare("DELETE FROM Teamroles").ExpectExec().
		WithArgs(userid, teamid).WillReturnResult(sqlmock.NewResult(1, 1))
}

func expectExitUserTeamDefault(mock sqlmock.Sqlmock) {
	expectExitUserTeam(mock, useridAddUser, teamidAddUser)
}

func expectExitUserTeamUsernameNotFound(mock sqlmock.Sqlmock) {
	expectGetUserEmpty(mock, usernameAddUser)
	expectGetTeam(mock, teamnameAddUser, defaultTeam())
	mock.ExpectPrepare("DELETE FROM Teamroles").ExpectExec().
		WithArgs(useridAddUser, teamidAddUser).WillReturnResult(sqlmock.NewResult(1, 1))
}

func expectExitUserTeamTeamnameNotFound(mock sqlmock.Sqlmock) {
	expectGetUser(mock, usernameAddUser, usernameAddUser, emailS, pwhash, typeTeacher)
	expectGetTeamEmpty(mock, teamnameAddUser)
	mock.ExpectPrepare("DELETE FROM Teamroles").ExpectExec().
		WithArgs(useridAddUser, teamidAddUser).WillReturnResult(sqlmock.NewResult(1, 1))
}

func expectExitUserTeamError(mock sqlmock.Sqlmock) {
	expectGetUser(mock, usernameAddUser, usernameAddUser, emailS, pwhash, typeTeacher)
	expectGetTeam(mock, teamnameAddUser, defaultTeam())
	mock.ExpectPrepare("DELETE FROM Teamroles").ExpectExec().
		WithArgs(useridAddUser, teamidAddUser).WillReturnError(fmt.Errorf("Error"))
}
func TestExitUserTeamNilDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectExitUserTeamDefault(mock)
	err = ExitUserTeam(nil, usernameAddUser, teamnameAddUser)
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
	err = ExitUserTeam(db, usernameAddUser, teamnameAddUser)
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
	err = ExitUserTeam(db, usernameAddUser, teamnameAddUser)
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
	err = ExitUserTeam(db, usernameAddUser, teamnameAddUser)
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
	err = ExitUserTeam(db, usernameAddUser, teamnameAddUser)
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
	err = ExitUserTeam(db, usernameAddUser, teamnameAddUser)
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}

// GetUsersFromTeam

func expectGetUsersFromTeam(mock sqlmock.Sqlmock, teamname string, users []*User) {
	expectGetTeam(mock, teamnameAddUser, defaultTeam())
	columns := []string{"id", "username", "email", "pwhash", "type"}
	rows := sqlmock.NewRows(columns)
	for _, u := range users {
		rows.AddRow(u.ID, u.Username, u.Email, u.Pwhash, u.Type)
	}
	mock.ExpectPrepare("SELECT (.+) FROM Users").ExpectQuery().WithArgs(defaultTeam().ID).WillReturnRows(rows)
}

func expectGetUsersFromTeamDefault(mock sqlmock.Sqlmock) {
	expectGetUsersFromTeam(mock, teamnameAddUser, defaultUsers())
}

func defaultUsersFromTeam() []*User {
	return defaultUsers()
}

func expectGetUsersFromTeamTeamnameNotFound(mock sqlmock.Sqlmock) {
	expectGetTeamEmpty(mock, teamnameAddUser)
	columns := []string{"id", "username", "email", "pwhash", "type"}
	rows := sqlmock.NewRows(columns)
	mock.ExpectPrepare("SELECT (.+) FROM Users").ExpectQuery().WithArgs(defaultTeam().ID).WillReturnRows(rows)
}

func expectGetUsersFromTeamError(mock sqlmock.Sqlmock) {
	expectGetTeam(mock, teamnameAddUser, defaultTeam())
	mock.ExpectPrepare("SELECT (.+) FROM Users").ExpectQuery().WithArgs(defaultTeam().ID).
		WillReturnError(fmt.Errorf("Error"))
}

func expectGetUsersFromTeamEmpty(mock sqlmock.Sqlmock) {
	expectGetTeam(mock, teamnameAddUser, defaultTeam())
	columns := []string{"id", "username", "email", "pwhash", "type"}
	rows := sqlmock.NewRows(columns)
	mock.ExpectPrepare("SELECT (.+) FROM Users").ExpectQuery().WithArgs(defaultTeam().ID).WillReturnRows(rows)
}

func TestGetUsersFromTeamNilDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	expectGetUsersFromTeamDefault(mock)
	u, err := GetUsersFromTeam(nil, teamnameAddUser)
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
	u, err := GetUsersFromTeam(db, teamnameAddUser)
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
	u, err := GetUsersFromTeam(db, teamnameAddUser)
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
	u, err := GetUsersFromTeam(db, teamnameAddUser)
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
	u, err := GetUsersFromTeam(db, teamnameAddUser)
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
	u, err := GetUsersFromTeam(db, teamnameAddUser)
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
	expectGetTeam(mock, teamnameAddUser, defaultTeam())
	columns := []string{"id", "username", "email", "pwhash", "type"}
	rows := sqlmock.NewRows(columns)
	for _, u := range users {
		rows.AddRow(u.ID, u.Username, u.Email, u.Pwhash, u.Type)
	}
	mock.ExpectPrepare("SELECT (.+) FROM Users").ExpectQuery().WithArgs(defaultTeam().ID).WillReturnRows(rows)
}

func expectGetTeamAdminsDefault(mock sqlmock.Sqlmock) {
	expectGetTeamAdmins(mock, teamnameAddUser, defaultTeamAdmins())
}

func expectGetTeamAdminsTeamnameNotFound(mock sqlmock.Sqlmock) {
	expectGetTeamEmpty(mock, teamnameAddUser)
	columns := []string{"id", "username", "email", "pwhash", "type"}
	rows := sqlmock.NewRows(columns)
	mock.ExpectPrepare("SELECT (.+) FROM Users").ExpectQuery().WithArgs(defaultTeam().ID).WillReturnRows(rows)
}

func expectGetTeamAdminsError(mock sqlmock.Sqlmock) {
	expectGetTeam(mock, teamnameAddUser, defaultTeam())
	mock.ExpectPrepare("SELECT (.+) FROM Users").ExpectQuery().WithArgs(defaultTeam().ID).
		WillReturnError(fmt.Errorf("Error"))
}

func expectGetTeamAdminsEmpty(mock sqlmock.Sqlmock) {
	expectGetTeam(mock, teamnameAddUser, defaultTeam())
	columns := []string{"id", "username", "email", "pwhash", "type"}
	rows := sqlmock.NewRows(columns)
	mock.ExpectPrepare("SELECT (.+) FROM Users").ExpectQuery().WithArgs(defaultTeam().ID).WillReturnRows(rows)
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
	u, err := GetTeamAdmins(nil, teamnameAddUser)
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
	u, err := GetTeamAdmins(db, teamnameAddUser)
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
	u, err := GetTeamAdmins(db, teamnameAddUser)
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
	u, err := GetTeamAdmins(db, teamnameAddUser)
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
	u, err := GetTeamAdmins(db, teamnameAddUser)
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
	u, err := GetTeamAdmins(db, teamnameAddUser)
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
