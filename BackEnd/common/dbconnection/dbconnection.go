// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package dbconnection provides a connection with the mysql/mariadb database
package dbconnection

import(
  "fmt"
  "database/sql"
  "encoding/json"
  "io/ioutil"

  _ "github.com/go-sql-driver/mysql"

)

// Returns database information stored at ./dbinfo.json
func getDbInfo() (DbInfo, error) {
  data, err := ioutil.ReadFile("./dbinfo.json")
  var dbinfo DbInfo
  if err!= nil {
    return dbinfo, err
  }
  err = json.Unmarshal(data, &dbinfo)
  return dbinfo, err
}

// Connects with the database and returns its sql.DB representation
func ConnectDb() (*sql.DB, error) {
  dbinfo, err := getDbInfo()
  if err != nil{
    return nil, err
  }
  dbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbinfo.Username, dbinfo.Pass, dbinfo.Host, dbinfo.Port, dbinfo.Name)
  db, err := sql.Open("mysql", dbSource)
  return db, err
}
