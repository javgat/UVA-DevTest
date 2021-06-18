// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas

// Package emailHelper provides functions to send proper emails to users
package emailHelper

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
