// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas

// Package emailHelper provides functions to send proper emails to users
package emailHelper

import (
	"errors"
	"log"
	"net/smtp"
	"uva-devtest/persistence/dao"
	"uva-devtest/persistence/dbconnection"
)

func sendEmailGoroutine(emailBody []byte, emailAddress string) {
	emailInfo, err := GetOwnEmailInfo()
	if err == nil {
		smtpServer := smtpServer{host: emailInfo.Serverhost, port: emailInfo.Serverport}
		auth := smtp.PlainAuth("", emailInfo.From, emailInfo.Password, smtpServer.host)
		to := []string{emailAddress}
		err = smtp.SendMail(smtpServer.Address(), auth, emailInfo.From, to, emailBody)
	}
	if err != nil {
		log.Println(err)
	}
}

// sendEmail sends content "emailBody" to the address, in a goroutine
func sendEmail(emailBody []byte, emailAddress string) {
	go sendEmailGoroutine(emailBody, emailAddress)
}

func getEmailFromUsername(username string) (string, error) {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var u *dao.User
		u, err = dao.GetUserUsername(db, username)
		if err == nil {
			if u == nil {
				return "", errors.New(errorResourceNotFound)
			}
			return u.Email.String(), nil
		}
	}
	return "", err
}
