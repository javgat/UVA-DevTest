// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas

// Package emailHelper provides functions to send proper emails to users
package emailHelper

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/smtp"
	"uva-devtest/persistence/dao"
	"uva-devtest/persistence/dbconnection"
)

const errorResourceNotFound = "no se encontro el recurso"

type EmailInfo struct {
	From       string `json:"from"`
	Password   string `json:"password"`
	Serverhost string `json:"serverhost"`
	Serverport string `json:"serverport"`
}

// smtpServer data to smtp server
type smtpServer struct {
	host string
	port string
}

// Address URI to smtp server
func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

// Opens the email information file and returns a EmailInfo struct
// Param filename: String containing the route to dbinfo file.
// Returns DbInfo struct, or err if something failed.
func getEmailInfo(filename string) (*EmailInfo, error) {
	data, err := ioutil.ReadFile(filename)
	var emailInfo *EmailInfo
	if err != nil {
		return emailInfo, err
	}
	err = json.Unmarshal(data, &emailInfo)
	return emailInfo, err
}

func getOwnEmailInfo() (*EmailInfo, error) {
	return getEmailInfo("./config/emailinfo.json")
}

func sendEmail(emailBody []byte, email string) {
	emailInfo, err := getOwnEmailInfo()
	if err == nil {
		smtpServer := smtpServer{host: emailInfo.Serverhost, port: emailInfo.Serverport}
		auth := smtp.PlainAuth("", emailInfo.From, emailInfo.Password, smtpServer.host)
		to := []string{email}
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

func generateEmailBodyRecoveryPassword(username string, token string, email string) []byte {
	msg := []byte("To: " + email + "\r\n" +
		"Subject: Recuperación de contraseña\r\n" +
		"\r\n" +
		"Para recuperar tu contraseña asociada a " + username + ", el token a usar es: " + token + "\r\n")
	return msg
}

func SendPasswordRecoveryMail(username string, token string) {
	email, err := getEmailFromUsername(username)
	if err == nil {
		emailBody := generateEmailBodyRecoveryPassword(username, token, email)
		sendEmail(emailBody, email)
	}
}
