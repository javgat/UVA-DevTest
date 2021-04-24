// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package dao acts as a Data Access Object for the Types
package dao

import (
	"database/sql"
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Transforms some sql.Rows into a slice(array) of MailTokens
// Param rows: Rows which contains database information returned
// Return []models.MailToken: MailTokens represented in rows
// Return error if any
func rowsToMailTokens(rows *sql.Rows) ([]*MailToken, error) {
	var MailTokens []*MailToken
	for rows.Next() {
		var t MailToken
		err := rows.Scan(&t.Userid, &t.Mailtoken, &t.Caducidad)
		if err != nil {
			return MailTokens, err
		}
		MailTokens = append(MailTokens, &t)
	}
	return MailTokens, nil
}

// Transforms rows into a single MailToken
// Param rows: Rows which contains database info of 1 MailToken
// Return *models.MailToken: MailToken represented in rows
// Return error if something happens
func rowsToMailToken(rows *sql.Rows) (*MailToken, error) {
	var MailToken *MailToken
	MailTokens, err := rowsToMailTokens(rows)
	if len(MailTokens) >= 1 {
		MailToken = MailTokens[0]
	}
	return MailToken, err
}

func PostRecoveryToken(db *sql.DB, username string, token string) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	u, err := GetUserUsername(db, username)
	if err == nil {
		if u == nil {
			return errors.New(errorResourceNotFound)
		}
		var query *sql.Stmt
		query, err = db.Prepare("INSERT INTO TokenCorreo(usuarioid, token, caducidad) VALUES(?,?,?)")
		caducidad := time.Now()
		tokenMinutes := 45
		caducidad.Add(time.Minute * time.Duration(tokenMinutes))
		if err == nil {
			defer query.Close()
			_, err = query.Exec(u.ID, token, caducidad)
		}
	}
	return err
}

// Obtains a MailToken, returns nil, nil if it does not exist
func GetMailToken(db *sql.DB, token *string) (*MailToken, error) {
	if db == nil {
		return nil, errors.New(errorDBNil)
	}
	var mt *MailToken
	query, err := db.Prepare("SELECT * FROM TokenCorreo WHERE token=?")
	if err == nil {
		defer query.Close()
		var rows *sql.Rows
		rows, err = query.Query(token)
		if err == nil {
			mt, err = rowsToMailToken(rows)
			return mt, err
		}
	}
	return nil, err
}

func DeleteMailToken(db *sql.DB, token *string) error {
	if db == nil {
		return errors.New(errorDBNil)
	}
	query, err := db.Prepare("DELETE FROM TokenCorreo WHERE token=?")
	if err == nil {
		defer query.Close()
		_, err = query.Exec(token)
	}
	return err
}
