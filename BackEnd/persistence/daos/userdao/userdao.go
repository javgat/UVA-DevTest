// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package userdao acts as a Data Access Object for the User Type
package userdao

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"

	"uva-devtest/models"
)

// Inserts a user into the database
// Param db: Database to use
// Param u: User to insert
// Return error if something wrong happens
func InsertUser(db *sql.DB, u *models.User) error {
	if db == nil || u == nil {
		return errors.New("Argumento de entrada nil")
	}
	query, err := db.Prepare("INSERT INTO users(username, email, pwhash) VALUES (?,?,?)")

	if err != nil {
		return err
	}

	_, err = query.Exec(u.Username, u.Email, u.Pwhash)
	defer query.Close()
	return err
}

// Transforms some sql.Rows into a slice(array) of users
// Param rows: Rows which contains database information returned
// Return []models.User: Users represented in rows
// Return error if any
func rowsToUsers(rows *sql.Rows) ([]models.User, error) {
	var users []models.User
	var trash int
	for rows.Next() {
		var us models.User
		err := rows.Scan(&trash, &us.Username, &us.Email, &us.Pwhash)
		if err != nil {
			return users, err
		}
		users = append(users, us)
	}
	return users, nil
}

// Transforms rows into a single user
// Param rows: Rows which contains database info of 1 user
// Return *models.User: User represented in rows
// Return error if something happens, or if there is more than 1 user
func rowsToUser(rows *sql.Rows) (*models.User, error) {
	var user *models.User
	users, err := rowsToUsers(rows)
	if len(users) >= 1 {
		user = &users[0]
	}
	return user, err
}

// Returns the user based on their username.
// Param db: Database in which the user will be looked for
// Param username: Username of the user
// Return *models.User: User found, or nil if not found
// Return error if something happened
func GetUserUsername(db *sql.DB, username string) (*models.User, error) {
	if db == nil {
		return nil, errors.New("Parametro db nil")
	}
	query, err := db.Prepare("SELECT * FROM users WHERE username=?")
	var u *models.User
	if err != nil {
		return u, err
	}
	rows, err := query.Query(username)
	if err == nil {
		u, err = rowsToUser(rows)
	}
	defer query.Close()
	return u, err
}

// Returns the user based on their email.
// Param db: Database in which the user will be looked for
// Param email: Email of the user
// Return *models.User: User found, or nil if not found
// Return error if something happened
func GetUserEmail(db *sql.DB, email string) (*models.User, error) {
	if db == nil {
		return nil, errors.New("Parametro db nil")
	}
	query, err := db.Prepare("SELECT * FROM users WHERE email=?")
	var u *models.User
	if err != nil {
		return u, err
	}
	rows, err := query.Query(email)
	if err == nil {
		u, err = rowsToUser(rows)
	}
	defer query.Close()
	return u, err
}
