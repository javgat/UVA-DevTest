// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package dbconnection provides a connection with the mysql/mariadb database
package dbconnection

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
)

// Opens the database information file and returns a DbInfo struct
// Param filename: String containing the route to dbinfo file.
// Returns DbInfo struct, or err if something failed.
func getDbInfo(filename string) (*DbInfo, error) {
	data, err := ioutil.ReadFile(filename)
	var dbinfo *DbInfo
	if err != nil {
		return dbinfo, err
	}
	err = json.Unmarshal(data, &dbinfo)
	return dbinfo, err
}

func connectDb(filename string) (*sql.DB, error) {
	dbinfo, err := getDbInfo(filename)
	if err != nil {
		return nil, err
	}
	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbinfo.Username, dbinfo.Pass, dbinfo.Host, dbinfo.Port, dbinfo.Name)
	db, err := sql.Open("mysql", dbSource)
	return db, err
}

// Connects with the database and returns its sql.DB representation
// Returns *sql.DB pointing to the MySQL/MariaDB database.
func ConnectDb() (*sql.DB, error) {
	return connectDb("../config/dbinfo.json")
}
