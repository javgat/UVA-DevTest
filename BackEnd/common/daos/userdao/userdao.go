// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package userdao acts as a Data Access Object for the User Type
package userdao

import(
  "database/sql"
  _ "github.com/go-sql-driver/mysql"

  "gitlab.com/HP-SCDS/Observatorio/2020-2021/uva-devtest/BackEnd/common/model"
)

// Inserts user <u> into the database <db>
func InsertUser(db *sql.DB, u model.User) error{
  query, err := db.Prepare("INSERT INTO users(username, email, pwhash) VALUES (?,?,?)")

  if err != nil {
    return err
  }

  _, err = query.Exec(u.Username, u.Email, u.PwHash)
  defer query.Close()
  return err
}

// Transforms <rows> into a slice of users
func rowsToUsers(rows *sql.Rows) ([]model.User, error){
  var users []model.User
  var trash int
  for rows.Next() {
    var us model.User
    err := rows.Scan(&trash, &us.Username, &us.Email, &us.PwHash)
    if err != nil{
      return users, err
    }
    users = append(users, us)
  }
  return users, nil
}

// Transforms <rows> into a single user
func rowsToUser(rows *sql.Rows) (*model.User, error){
  var user *model.User
  users, err := rowsToUsers(rows)
  if len(users) >= 1{
    user = &users[0]
  }
  return user, err
}

// Returns the user whose username is <username>. If there is no such user,
// returns nil
func GetUserUsername(db *sql.DB, username string) (*model.User, error){
  query, err := db.Prepare("SELECT * FROM users WHERE username=?")
  var u *model.User
  if err != nil {
    return u, err
  }
  rows, err := query.Query(username)
  u, err = rowsToUser(rows)
  defer query.Close()
  return u, err
}

// Returns the user whose email is <email>. If there is no such user,
// returns nil
func GetUserEmail(db *sql.DB, email string) (*model.User, error){
  query, err := db.Prepare("SELECT * FROM users WHERE email=?")
  var u *model.User
  if err != nil {
    return u, err
  }
  rows, err := query.Query(email)
  u, err = rowsToUser(rows)
  defer query.Close()
  return u, err
}
