// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas

// Package emailHelper provides functions to send proper emails to users
package emailHelper

func generateEmailBodyUserCreated(username string, pass string, address string, frontEnd string) []byte {

	msg := []byte("To: " + address + "\r\n" +
		"Subject: [NO RESPONDER] Usuario Creado\r\n" +
		"\r\n" +
		"Hola " + username + ", se ha creado una cuenta asociada a este correo electronico.\r\n" +
		"Puedes acceder mediante la web, o pulsando el siguiente enlace: " + frontEnd + "/login \r\n" +
		"\r\n" +
		"Los datos de inicio de sesion son:\r\n" +
		"ID: " + username + "\r\n" +
		"Contraseña: " + pass)
	return msg

}

// SendEmailUserCreated sends an email to the user with username username,
// notifying them that they have been created
func SendEmailUserCreated(username string, pass string) {
	address, err := getEmailFromUsername(username)
	if err == nil {
		emailInfo, err := getOwnEmailInfo()
		if err == nil {
			emailBody := generateEmailBodyUserCreated(username, pass, address, emailInfo.FrontEndUrl)
			sendEmail(emailBody, address)
		}
	}
}
