package userdao

import(
  "database/sql"
  _ "github.com/go-sql-driver/mysql"

  "gitlab.com/HP-SCDS/Observatorio/2020-2021/uva-devtest/BackEnd/common/model"
)

func InsertUser(db *sql.DB, u model.User) error{
  query, err := db.Prepare("INSERT INTO users(username, email, pwhash) VALUES (?,?,?)")

  if err != nil {
    return err
  }

  _, err = query.Exec(u.Username, u.Email, u.PwHash)
  defer query.Close()
  return err
}
