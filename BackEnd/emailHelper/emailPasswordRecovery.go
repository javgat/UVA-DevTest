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

// sendEmail sends content "emailBody" to the address
func sendEmail(emailBody []byte, emailAddress string) {
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

func generateEmailBodyRecoveryPassword(username string, token string, email string, frontEnd string) []byte {
	msg := []byte("To: " + email + "\r\n" +
		"Subject: [NO RESPONDER] Recuperación de contraseña\r\n" +
		"\r\n" +
		"Para recuperar tu contraseña asociada a " + username + ", haz click en: " + frontEnd + "/recoverPassword/" + username + "?token=" + token + "\r\n")
	return msg
}

// SendPasswordRecoveryMail sends an email to the user username, with a link to recover the password
func SendPasswordRecoveryMail(username string, token string) {
	email, err := getEmailFromUsername(username)
	if err == nil {
		emailInfo, err := GetOwnEmailInfo()
		if err == nil {
			emailBody := generateEmailBodyRecoveryPassword(username, token, email, emailInfo.FrontEndUrl)
			sendEmail(emailBody, email)
		}
	}
}
