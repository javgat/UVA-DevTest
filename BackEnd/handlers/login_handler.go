// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package handlers provides functions that handle http requests
package handlers

import (
	"errors"
	"log"
	"uva-devtest/models"
	"uva-devtest/persistence/daos/userdao"
	"uva-devtest/persistence/dbconnection"
	"uva-devtest/restapi/operations/auth"

	"github.com/go-openapi/runtime/middleware"
	"golang.org/x/crypto/bcrypt"

	"uva-devtest/jwtauth"
)

// Responds with Server Error
func serverErrorLogin(err error) middleware.Responder {
	log.Println(err)
	return auth.NewLoginInternalServerError()
}

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
func successLogin(u userdao.User) middleware.Responder {
	log.Println("Usuario logged in")
	var wrap jwtauth.JwtWrapper
	wrap.SecretKey = *u.Pwhash
	wrap.Issuer = "DevTest"
	wrap.ExpirationHours = 24
	signedToken, err := wrap.GenerateToken(u.Email.String())
	log.Println(wrap.SecretKey, u.Email.String(), signedToken)
	tokenJSON := jwtauth.CreateJwtJSON(signedToken)
	if err != nil {
		return badReqErrorLogin(err)
	}
	return auth.NewLoginOK().WithPayload(&tokenJSON)
}

// Login is the main handler function for the login functionality
// Param params Parametros de entrada que tiene la peticion http
// Return middleware.Responder
func Login(params auth.LoginParams) middleware.Responder {
	log.Println("Generando Token JWT de usuario...")
	var lu *models.LoginUser = params.LoginUser
	if lu == nil {
		return badReqErrorLogin(errors.New("Parametros de entrada vacios"))
	}
	log.Printf("Login id: %v\n", *lu.Loginid)
	db, err := dbconnection.ConnectDb()
	if err != nil {
		return serverErrorLogin(err)
	}
	log.Println("Conectado a la base de datos")
	var u *userdao.User
	if *lu.Loginid == "" {
		return badReqErrorLogin(nil)
	}
	// Primero compruebo si la LoginId corresponde a un username
	u, err = userdao.GetUserUsername(db, *lu.Loginid)
	if err != nil {
		return serverErrorLogin(err)
	}
	if u == nil {
		//Si no corresponde, compruebo con un email
		u, err = userdao.GetUserEmail(db, *lu.Loginid)
	}
	if u == nil {
		return authFailErrorLogin(err, "Usuario no existe")
	}
	authErr := bcrypt.CompareHashAndPassword([]byte(*u.Pwhash), []byte(*lu.Pass))
	if authErr != nil {
		return authFailErrorLogin(err, "Password incorrecto")
	}
	return successLogin(*u)

}
