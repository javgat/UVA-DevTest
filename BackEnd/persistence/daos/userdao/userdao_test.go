// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

package userdao

import (
	"fmt"
	"testing"
	"uva-devtest/models"

	"github.com/DATA-DOG/go-sqlmock"
	"golang.org/x/crypto/bcrypt"
)

// Testing InsertUser

func signinUserToInsert() *models.SigninUser {
	uname := "Test"
	email := "Test@mail.com"
	pass := "Test"
	u := &models.SigninUser{
		Username: &uname,
		Email:    &email,
		Pass:     &pass,
	}
	return u
}

func userToInsert() *models.User {
	su := signinUserToInsert()
	bytes, _ := bcrypt.GenerateFromPassword([]byte(*su.Pass), 14)
	pwhashstring := string(bytes)
	u := &models.User{
		Username: su.Username,
		Email:    su.Email,
		Pwhash:   &pwhashstring,
	}
	return u
}

func expectInsert(mock sqlmock.Sqlmock, u *models.User) {
	mock.ExpectPrepare("INSERT INTO users").ExpectExec().
		WithArgs(u.Username, u.Email, u.Pwhash).WillReturnResult(sqlmock.NewResult(1, 1))
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
		t.Log("error should be nil")
		t.Fail()
	}
}

// Testing getUserUsername y getUserEmail

const username string = "Test"
const email string = "Test@mail.com"
const pwhash string = "AAAAAAAAAAA"

func expectGetUser(mock sqlmock.Sqlmock, arg string, u string, e string, p string) {
	columns := []string{"id", "username", "email", "pwhash"}
	mock.ExpectPrepare("SELECT (.+) FROM users").ExpectQuery().WithArgs(arg).
		WillReturnRows(sqlmock.NewRows(columns).AddRow(1, u, e, p))
}

func expectGetEmail(mock sqlmock.Sqlmock) {
	expectGetUser(mock, email, username, email, pwhash)
}

func expectGetUsername(mock sqlmock.Sqlmock) {
	expectGetUser(mock, username, username, email, pwhash)
}

func expectGetUserEmpty(mock sqlmock.Sqlmock, arg string) {
	columns := []string{"id", "username", "email", "pwhash"}
	mock.ExpectPrepare("SELECT (.+) FROM users").ExpectQuery().WithArgs(arg).
		WillReturnRows(sqlmock.NewRows(columns))
}

func expectGetEmailEmpty(mock sqlmock.Sqlmock) {
	expectGetUserEmpty(mock, email)
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
	expectGetUserTablesWrong(mock, username, username, email)
}

func expectGetEmailTablesWrong(mock sqlmock.Sqlmock) {
	expectGetUserTablesWrong(mock, email, username, email)
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
	if *u.Email != email {
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
	u, err := GetUserEmail(nil, email)
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
	u, err := GetUserEmail(db, email)
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
	u, err := GetUserEmail(db, email)
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
	u, err := GetUserEmail(db, email)
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
	u, err := GetUserEmail(db, email)
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
	u, err := GetUserEmail(db, email)
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
