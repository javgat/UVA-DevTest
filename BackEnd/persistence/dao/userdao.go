// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package dao acts as a Data Access Object for the Types
package dao

import (
	"database/sql"
	"errors"
	"uva-devtest/models"

	// Blank import of mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ToModelUser converts a dao.User into a models.User
// Param u: dao.User to convert
func ToModelUser(u *User) *models.User {
	mu := &models.User{
		Username: u.Username,
		Email:    u.Email,
		Rol:      u.Rol,
		Fullname: u.Fullname,
	}
	return mu
}

// ToModelsUser converts a splice of dao.User into models.User
// Param us: slice of dao.User to convert
func ToModelsUser(us []*User) []*models.User {
	var mus = []*models.User{}
	for _, itemCopy := range us {
		mus = append(mus, ToModelUser(itemCopy))
	}
	return mus
}

// InsertUser inserts a user into the database
// Param db: Database to use
// Param u: User to insert
// Return error if something wrong happens
func InsertUser(db *sql.DB, u *User) error {
	if db == nil || u == nil {
		return errors.New(errorDBNil)
	}
	query, err := db.Prepare("INSERT INTO Usuario(username, email, pwhash, rol, fullname) VALUES (?,?,?,?,?)")

	if err != nil {
		return err
	}
	defer query.Close()

	_, err = query.Exec(u.Username, u.Email, u.Pwhash, u.Rol, u.Fullname)
	return err
}

// Transforms some sql.Rows into a slice(array) of users
// Param rows: Rows which contains database information returned
// Return []models.User: Users represented in rows
// Return error if any
func rowsToUsers(rows *sql.Rows) ([]*User, error) {
	var users []*User
	for rows.Next() {
		var us User
		err := rows.Scan(&us.ID, &us.Username, &us.Email, &us.Pwhash, &us.Rol, &us.Fullname)
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
// Return error if something happens
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
		return nil, errors.New(errorDBNil)
	}
	var u *User
	query, err := db.Prepare("SELECT * FROM Usuario WHERE username=?")
	if err != nil {
		return u, err
	}
	defer query.Close()
	rows, err := query.Query(username)
	if err == nil {
		u, err = rowsToUser(rows)
	}
	return u, err
}

// GetUserEmail returns the user based on their email.
// Param db: Database in which the user will be looked for
// Param email: Email of the user
// Return *models.User: User found, or nil if not found
// Return error if something happened
func GetUserEmail(db *sql.DB, email string) (*User, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var u *User
	query, err := db.Prepare("SELECT * FROM Usuario WHERE email=?")
	if err != nil {
		return u, err
	}
	defer query.Close()
	rows, err := query.Query(email)
	if err == nil {
		u, err = rowsToUser(rows)
	}
	return u, err
}

func GetUserByID(db *sql.DB, ID int64) (*User, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var u *User
	query, err := db.Prepare("SELECT * FROM Usuario WHERE id=?")
	if err != nil {
		return u, err
	}
	defer query.Close()
	rows, err := query.Query(ID)
	if err == nil {
		u, err = rowsToUser(rows)
	}
	return u, err
}

// GetUsers returns all users
func GetUsers(db *sql.DB) ([]*User, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	query, err := db.Prepare("SELECT * FROM Usuario")
	var us []*User
	if err != nil {
		return us, err
	}
	defer query.Close()
	rows, err := query.Query()
	if err == nil {
		us, err = rowsToUsers(rows)
	}
	return us, err
}

// GetAdmins returns all users that are admins
func GetAdmins(db *sql.DB) ([]*User, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	query, err := db.Prepare("SELECT * FROM Usuario WHERE rol=?")
	var us []*User
	if err != nil {
		return us, err
	}
	defer query.Close()
	rows, err := query.Query(models.UserRolAdministrador)
	if err == nil {
		us, err = rowsToUsers(rows)
	}
	return us, err
}

// PutPasswordUsername modifies the pwhash of user <username> in database <db>
// Param username: Username of the user
// Param newpwhash: New Password Hash to insert in the database
func PutPasswordUsername(db *sql.DB, username string, newpwhash string) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	query, err := db.Prepare("UPDATE Usuario SET pwhash = ? WHERE username = ?")
	if err != nil {
		return err
	}
	defer query.Close()
	_, err = query.Exec(newpwhash, username)
	return err
}

// UpdateUser updates a user from the database
// Param db: Database to use
// Param u: User data to update
// Param username: Username of the user to update
// Return error if something wrong happens
func UpdateUser(db *sql.DB, u *models.User, username string) error {
	if db == nil || u == nil {
		return errors.New(errorDBNil)
	}
	query, err := db.Prepare("UPDATE Usuario SET username=?, email=?, fullname=?, rol=? WHERE username = ? ")
	if err != nil {
		return err
	}
	defer query.Close()
	_, err = query.Exec(u.Username, u.Email, u.Fullname, u.Rol, username)
	return err
}

// DeleteUser deletes user <username> from the database
// Param username: Username of the user
func DeleteUser(db *sql.DB, username string) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	query, err := db.Prepare("DELETE FROM Usuario WHERE username = ? ") //ESTO se supone que borra en cascade
	if err != nil {
		return err
	}
	defer query.Close()
	_, err = query.Exec(username)
	return err
}

//NOTESTED (no se testea)

func addUserTeam(db *sql.DB, username string, teamname string, teamrole string) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err != nil {
		return err
	}
	t, err := GetTeam(db, teamname)
	if err != nil {
		return err
	}
	if u == nil || t == nil {
		return errors.New(errorResourceNotFound)
	}
	query, err := db.Prepare("INSERT INTO EquipoUsuario(userid, teamid, role) VALUES (?, ?, ?) ")
	if err != nil {
		return err
	}
	defer query.Close()
	_, err = query.Exec(u.ID, t.ID, teamrole)
	return err
}

// AddUserTeamMember adds a user to a team as a member
// Param username: Username of the user
// Param teamname: Teamname of the team
func AddUserTeamMember(db *sql.DB, username string, teamname string) error {
	return addUserTeam(db, username, teamname, models.TeamRoleRoleMember)
}

// AddUserTeamAdmin adds a user to a team as an Admin
// Param username: Username of the user
// Param teamname: Teamname of the team
func AddUserTeamAdmin(db *sql.DB, username string, teamname string) error {
	return addUserTeam(db, username, teamname, models.TeamRoleRoleAdmin)
}

// ExitUserTeam gets out a user from a team
// Param username: Username of the user
// Param teamname: Teamname of the team
func ExitUserTeam(db *sql.DB, username string, teamname string) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err != nil {
		return err
	}
	t, err := GetTeam(db, teamname)
	if err != nil {
		return err
	}
	if u == nil || t == nil {
		return errors.New(errorResourceNotFound)
	}
	query, err := db.Prepare("DELETE FROM EquipoUsuario WHERE userid = ? AND teamid = ? ")
	if err != nil {
		return err
	}
	defer query.Close()
	_, err = query.Exec(u.ID, t.ID)
	return err
}

// GetUsersFromTeam returns all users
// Param teamname: Teamname of the team
func GetUsersFromTeam(db *sql.DB, teamname string) ([]*User, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	t, err := GetTeam(db, teamname)
	if err != nil {
		return nil, err
	} else if t == nil {
		return nil, errors.New(errorResourceNotFound)
	}
	query, err := db.Prepare("SELECT U.* FROM Usuario U JOIN EquipoUsuario R ON	U.id=R.userid WHERE R.teamid=?")
	var us []*User
	if err != nil {
		return nil, err
	}
	defer query.Close()
	rows, err := query.Query(t.ID)
	if err == nil {
		us, err = rowsToUsers(rows)
	}
	return us, err
}

func getTeamUsersByRole(db *sql.DB, teamname string, teamrole string) ([]*User, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	t, err := GetTeam(db, teamname)
	if err != nil {
		return nil, err
	} else if t == nil {
		return nil, errors.New(errorResourceNotFound)
	}
	query, err := db.Prepare("SELECT U.* FROM Users U JOIN Teamroles R ON	U.id=R.userid WHERE R.teamid=? AND R.role=?")
	var us []*User
	if err != nil {
		return nil, err
	}
	defer query.Close()
	rows, err := query.Query(t.ID, teamrole)
	if err == nil {
		us, err = rowsToUsers(rows)
	}
	return us, err
}

// GetTeamAdmins returns all users of team that are admins in team
// Param teamname: Teamname of the team
func GetTeamAdmins(db *sql.DB, teamname string) ([]*User, error) {
	return getTeamUsersByRole(db, teamname, models.TeamRoleRoleAdmin)
}

// GetTeamMembers returns all users of team that are role members in team
// Param teamname: Teamname of the team
func GetTeamMembers(db *sql.DB, teamname string) ([]*User, error) {
	return getTeamUsersByRole(db, teamname, models.TeamRoleRoleMember)
}
