// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package userdao acts as a Data Access Object for the User Type
package userdao

import (
	"database/sql"
	"errors"
	"uva-devtest/models"

	// Blank import of mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// DaoToModelUser converts a userdao.User into a models.User
func DaoToModelUser(u *User) *models.User {
	mu := &models.User{
		Username: u.Username,
		Email:    u.Email,
		Type:     "admin", // CAMBIAR !!!!
	}
	return mu
}

// DaoToModelsUser converts a splice of userdao.User into models.User
func DaoToModelsUser(us []*User) []*models.User {
	var mus = []*models.User{}
	for _, itemCopy := range us {
		mus = append(mus, DaoToModelUser(itemCopy))
	}
	return mus
}

// InsertUser inserts a user into the database
// Param db: Database to use
// Param u: User to insert
// Return error if something wrong happens
func InsertUser(db *sql.DB, u *User) error {
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
func rowsToUsers(rows *sql.Rows) ([]*User, error) {
	var users []*User
	var trash int
	for rows.Next() {
		var us User
		err := rows.Scan(&trash, &us.Username, &us.Email, &us.Pwhash)
		if err != nil {
			return users, err
		}
		users = append(users, &us)
	}
	return users, nil
}

// Transforms rows into a single user
// Param rows: Rows which contains database info of 1 user
// Return *models.User: User represented in rows
// Return error if something happens, or if there is more than 1 user
func rowsToUser(rows *sql.Rows) (*User, error) {
	var user *User
	users, err := rowsToUsers(rows)
	if len(users) >= 1 {
		user = users[0]
	}
	return user, err
}

// GetUserUsername returns the user based on their username.
// Param db: Database in which the user will be looked for
// Param username: Username of the user
// Return *models.User: User found, or nil if not found
// Return error if something happened
func GetUserUsername(db *sql.DB, username string) (*User, error) {
	if db == nil {
		return nil, errors.New("Parametro db nil")
	}
	query, err := db.Prepare("SELECT * FROM users WHERE username=?")
	var u *User
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

// GetUserEmail returns the user based on their email.
// Param db: Database in which the user will be looked for
// Param email: Email of the user
// Return *models.User: User found, or nil if not found
// Return error if something happened
func GetUserEmail(db *sql.DB, email string) (*User, error) {
	if db == nil {
		return nil, errors.New("Parametro db nil")
	}
	query, err := db.Prepare("SELECT * FROM users WHERE email=?")
	var u *User
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

// GetUsers returns all users
func GetUsers(db *sql.DB) ([]*User, error) {
	if db == nil {
		return nil, errors.New("Parametro db nil")
	}
	query, err := db.Prepare("SELECT * FROM users")
	var us []*User
	if err != nil {
		return us, err
	}
	rows, err := query.Query()
	if err == nil {
		us, err = rowsToUsers(rows)
	}
	defer query.Close()
	return us, err
}

// PutPasswordUsername modifies the pwhash of user <username> in database <db>
func PutPasswordUsername(db *sql.DB, username string, newpwhash string) error {
	if db == nil {
		return errors.New("Parametro db nil")
	}
	query, err := db.Prepare("UPDATE users SET pwhash = ? WHERE username = ?")
	if err != nil {
		return err
	}
	_, err = query.Exec(newpwhash, username)
	defer query.Close()
	return err
}

// UpdateUser updates a user from the database
// Param db: Database to use
// Param u: User data to update
// Param username: Username of the user to update
// Return error if something wrong happens
func UpdateUser(db *sql.DB, u *models.User, username string) error {
	if db == nil || u == nil {
		return errors.New("Argumento de entrada nil")
	}
	query, err := db.Prepare("UPDATE users SET email=?, fullname=?, username=?, type=? WHERE username = ? ")
	if err != nil {
		return err
	}
	_, err = query.Exec(u.Email, u.Fullname, u.Username, u.Type, username)
	defer query.Close()
	return err
}

// DeleteUser deletes user <username> from the database
func DeleteUser(db *sql.DB, username string) error {
	if db == nil {
		return errors.New("Argumento de entrada nil")
	}
	query, err := db.Prepare("DELETE FROM users WHERE username = ? ") //ESTO TENDRA QUE SER MAS COMPLEJO, RELACIONES
	if err != nil {
		return err
	}
	_, err = query.Exec(username)
	defer query.Close()
	return err
}

// AddUserTeam adds a user to a team
func AddUserTeam(db *sql.DB, username string, teamname string) error {
	if db == nil {
		return errors.New("Argumento de entrada nil")
	}
	query, err := db.Prepare("INSERT INTO teamroles(username, teamname, role) VALUES (?, ?, ?) ")
	if err != nil {
		return err
	}
	_, err = query.Exec(username, teamname, models.TeamRoleRoleMember)
	defer query.Close()
	return err
}

// ExitUserTeam gets out a user from a team
func ExitUserTeam(db *sql.DB, username string, teamname string) error {
	if db == nil {
		return errors.New("Argumento de entrada nil")
	}
	query, err := db.Prepare("DELETE FROM teamroles WHERE username = ? AND teamname = ? ")
	if err != nil {
		return err
	}
	_, err = query.Exec(username, teamname)
	defer query.Close()
	return err
}

// GetUsersFromTeam returns all users
func GetUsersFromTeam(db *sql.DB, teamname string) ([]*models.User, error) {
	if db == nil {
		return nil, errors.New("Parametro db nil")
	}
	query, err := db.Prepare("SELECT U FROM users U JOIN teamroles R WHERE R.teamname=?")
	var us []*User
	if err != nil {
		return nil, err
	}
	rows, err := query.Query(teamname)
	if err == nil {
		us, err = rowsToUsers(rows)
	}
	defer query.Close()
	return DaoToModelsUser(us), err
}
