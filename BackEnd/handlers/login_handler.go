// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package handlers provides functions that handle http requests
package handlers

import (
	"log"
	"uva-devtest/models"
	"uva-devtest/persistence/dao"
	"uva-devtest/persistence/dbconnection"
	"uva-devtest/restapi/operations/auth"

	"github.com/go-openapi/runtime/middleware"
	"golang.org/x/crypto/bcrypt"
)

// Responds with Bad Request Error
func badReqErrorLogin(err error) middleware.Responder {
	log.Println(err)
	errSt := "Datos de log invalidos"
	prerr := models.Error{
		Message: &errSt,
	}
	return auth.NewLoginBadRequest().WithPayload(&prerr)
}

// Responds with authentication Failure Error
func authFailErrorLogin(err error, info string) middleware.Responder {
	log.Println(err)
	errSt := info
	prerr := models.Error{
		Message: &errSt,
	}
	return auth.NewLoginGone().WithPayload(&prerr)
}

// The user is logged in, the handler will try to respond with a JWT
func successLogin(u dao.User) middleware.Responder {
	log.Println("Usuario logged in")

	signedToken, err := CreateJWT(u, int64(AuthHours))
	if err != nil {
		return badReqErrorLogin(err)
	}
	resignedToken, err := CreateJWT(u, int64(ReauthHours))
	if err != nil {
		return badReqErrorLogin(err)
	}
	cookie := CreateCookie(BearerCookieName, signedToken, hoursToSeconds(AuthHours))
	recookie := CreateCookie(ReauthCookieName, resignedToken, hoursToSeconds(ReauthHours))
	return auth.NewLoginCreated().WithAuth(cookie).WithReAuth(recookie)
}

// Login is the main handler function for the login functionality
// Param params Parametros de entrada que tiene la peticion http
// Return middleware.Responder
func Login(params auth.LoginParams) middleware.Responder {
	log.Println("Generando Token JWT de usuario...")
	var lu *models.LoginUser = params.LoginUser
	log.Printf("Login id: %v\n", *lu.Loginid)
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var u *dao.User
		if *lu.Loginid == "" {
			return badReqErrorLogin(nil)
		}
		// Primero compruebo si la LoginId corresponde a un username
		u, err = dao.GetUserUsername(db, *lu.Loginid)
		if err == nil {
			if u == nil {
				//Si no corresponde, compruebo con un email
				u, err = dao.GetUserEmail(db, *lu.Loginid)
			}
			if err == nil {
				if u == nil {
					return authFailErrorLogin(err, "Usuario no existe")
				}
				authErr := bcrypt.CompareHashAndPassword([]byte(*u.Pwhash), []byte(*lu.Pass))
				if authErr != nil {
					return authFailErrorLogin(err, "Password incorrecto")
				}
				return successLogin(*u)
			}
		}
	}
	log.Println("Error en Login, ", err)
	return auth.NewLoginInternalServerError()

}
