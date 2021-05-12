// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas

// Package emailHelper provides functions to send proper emails to users
package emailHelper

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const errorResourceNotFound = "no se encontro el recurso"

const emailinfo_filename = "./config/emailinfo.json"

type EmailInfo struct {
	From        string `json:"from"`
	Password    string `json:"password"`
	Serverhost  string `json:"serverhost"`
	Serverport  string `json:"serverport"`
	FrontEndUrl string `json:"frontendurl"`
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
func GetEmailInfo(filename string) (*EmailInfo, error) {
	data, err := ioutil.ReadFile(filename)
	var emailInfo *EmailInfo
	if err != nil {
		return emailInfo, err
	}
	err = json.Unmarshal(data, &emailInfo)
	return emailInfo, err
}

func GetOwnEmailInfo() (*EmailInfo, error) {
	return GetEmailInfo(emailinfo_filename)
}

func PutEmailInfo(filename string, emailInfo *EmailInfo) error {
	data, err := json.Marshal(emailInfo)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, data, os.ModePerm)
	return err
}

func PutOwnEmailInfo(emailInfo *EmailInfo) error {
	return PutEmailInfo(emailinfo_filename, emailInfo)
}
