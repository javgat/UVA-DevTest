// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas

// Package emailHelper provides functions to send proper emails to users
package emailHelper

import (
	"fmt"
	"uva-devtest/models"
	"uva-devtest/persistence/dao"
	"uva-devtest/persistence/dbconnection"
)

func generateEmailBodyUserInvitedTest(username string, testid int64, address string, frontEnd string, m *models.Message) ([]byte, error) {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var test *dao.Test
		test, err = dao.GetTest(db, testid)
		mensInvitacion := ""
		if m != nil && m.Body != nil {
			mensInvitacion = *m.Body
		}
		if err == nil && test != nil {
			msg := []byte("To: " + address + "\r\n" +
				"Subject: [NO RESPONDER] Invitado al test: " + *test.Title + "\r\n" +
				"\r\n" +
				"Hola " + username + ", has sido invitado a realizar el test " + *test.Title + ".\r\n" +
				"Puedes acceder mediante la web, o pulsando el siguiente enlace: " + frontEnd + "/pt/" + fmt.Sprint(testid) + "\r\n" +
				"\r\n" +
				mensInvitacion)
			return msg, nil
		}
	}
	return nil, err
}

// SendEmailUserInvitedToTest sends an email to the user with username username,
// notifying them that they have been invited to a test
func SendEmailUserInvitedToTest(username string, testid int64, m *models.Message) {
	address, err := getEmailFromUsername(username)
	if err == nil {
		emailInfo, err := GetOwnEmailInfo()
		if err == nil {
			emailBody, err := generateEmailBodyUserInvitedTest(username, testid, address, emailInfo.FrontEndUrl, m)
			if err == nil {
				sendEmail(emailBody, address)
			}
		}
	}
}

func generateEmailBodyUserTeamInvitedTest(username string, testid int64, address string, frontEnd string, teamname string, m *models.Message) ([]byte, error) {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var test *dao.Test
		test, err = dao.GetTest(db, testid)
		mensInvitacion := ""
		if m != nil && m.Body != nil {
			mensInvitacion = *m.Body
		}
		if err == nil && test != nil {
			msg := []byte("To: " + address + "\r\n" +
				"Subject: [NO RESPONDER] Invitado al test: " + *test.Title + "\r\n" +
				"\r\n" +
				"Hola " + username + ", has sido invitado a realizar el test " + *test.Title + ", debido a pertenecer al equipo " + teamname + ".\r\n" +
				"Puedes acceder mediante la web, o pulsando el siguiente enlace: " + frontEnd + "/pt/" + fmt.Sprint(testid) + "\r\n" +
				"\r\n" +
				mensInvitacion)
			return msg, nil
		}
	}
	return nil, err
}

// SendEmailUserInvitedToTest sends an email to the user with username username,
// notifying them that they have been invited to a test
func SendEmailUserTeamInvitedToTest(username string, testid int64, teamname string, m *models.Message) {
	address, err := getEmailFromUsername(username)
	if err == nil {
		emailInfo, err := GetOwnEmailInfo()
		if err == nil {
			emailBody, err := generateEmailBodyUserTeamInvitedTest(username, testid, address, emailInfo.FrontEndUrl, teamname, m)
			if err == nil {
				sendEmail(emailBody, address)
			}
		}
	}
}

func generateEmailBodyUserInvitedTeam(address string, m *models.Message, titulo string, cuerpo string) []byte {
	mensInvitacion := ""
	if m != nil && m.Body != nil {
		mensInvitacion = *m.Body
	}
	msg := []byte("To: " + address + "\r\n" +
		titulo + "\r\n" +
		"\r\n" +
		cuerpo + "\r\n" +
		"\r\n" +
		mensInvitacion)
	return msg
}

func generateEmailBodyUserAddedToTeamAsMember(username string, address string, frontEnd string, teamname string, m *models.Message) []byte {
	titulo := "Subject: [NO RESPONDER] Añadido al equipo: " + teamname
	cuerpo := "Hola " + username + ", has sido añadido al equipo " + teamname + ".\r\n" +
		"Puedes acceder mediante la web, o pulsando el siguiente enlace: " + frontEnd + "/teams/" + teamname + "."
	return generateEmailBodyUserInvitedTeam(address, m, titulo, cuerpo)
}

func generateEmailBodyUserChangedToTeamMember(username string, address string, frontEnd string, teamname string, m *models.Message) []byte {
	titulo := "Subject: [NO RESPONDER] Retirados permisos de administración del equipo: " + teamname
	cuerpo := "Hola " + username + ", se te han retirado los permisos de administración en el equipo " + teamname + ".\r\n" +
		"Puedes acceder mediante la web, o pulsando el siguiente enlace: " + frontEnd + "/teams/" + teamname + "."
	return generateEmailBodyUserInvitedTeam(address, m, titulo, cuerpo)
}

func generateEmailBodyUserAddedToTeamAsAdmin(username string, address string, frontEnd string, teamname string, m *models.Message) []byte {
	titulo := "Subject: [NO RESPONDER] Añadido como administrador al equipo: " + teamname
	cuerpo := "Hola " + username + ", has sido añadido como administrador al equipo " + teamname + ".\r\n" +
		"Puedes acceder mediante la web, o pulsando el siguiente enlace: " + frontEnd + "/teams/" + teamname + "."
	return generateEmailBodyUserInvitedTeam(address, m, titulo, cuerpo)
}

func generateEmailBodyUserChangedToTeamAdmin(username string, address string, frontEnd string, teamname string, m *models.Message) []byte {
	titulo := "Subject: [NO RESPONDER] Obtenido permisos de administración del equipo: " + teamname
	cuerpo := "Hola " + username + ", has sido cambiado a administrador en el equipo " + teamname + ".\r\n" +
		"Puedes acceder mediante la web, o pulsando el siguiente enlace: " + frontEnd + "/teams/" + teamname + "."
	return generateEmailBodyUserInvitedTeam(address, m, titulo, cuerpo)
}

func SendEmailUserAddedToTeamAsMember(username string, teamname string, m *models.Message) {
	address, err := getEmailFromUsername(username)
	if err == nil {
		emailInfo, err := GetOwnEmailInfo()
		if err == nil {
			emailBody := generateEmailBodyUserAddedToTeamAsMember(username, address, emailInfo.FrontEndUrl, teamname, m)
			sendEmail(emailBody, address)
		}
	}
}

func SendEmailUserChangedToTeamMember(username string, teamname string, m *models.Message) {
	address, err := getEmailFromUsername(username)
	if err == nil {
		emailInfo, err := GetOwnEmailInfo()
		if err == nil {
			emailBody := generateEmailBodyUserChangedToTeamMember(username, address, emailInfo.FrontEndUrl, teamname, m)
			sendEmail(emailBody, address)
		}
	}
}

func SendEmailUserAddedToTeamAsAdmin(username string, teamname string, m *models.Message) {
	address, err := getEmailFromUsername(username)
	if err == nil {
		emailInfo, err := GetOwnEmailInfo()
		if err == nil {
			emailBody := generateEmailBodyUserAddedToTeamAsAdmin(username, address, emailInfo.FrontEndUrl, teamname, m)
			sendEmail(emailBody, address)
		}
	}
}

func SendEmailUserChangedToTeamAdmin(username string, teamname string, m *models.Message) {
	address, err := getEmailFromUsername(username)
	if err == nil {
		emailInfo, err := GetOwnEmailInfo()
		if err == nil {
			emailBody := generateEmailBodyUserChangedToTeamAdmin(username, address, emailInfo.FrontEndUrl, teamname, m)
			sendEmail(emailBody, address)
		}
	}
}
