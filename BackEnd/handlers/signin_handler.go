// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package handlers handles the http request
package handlers

import (
	"log"

	"github.com/go-openapi/runtime/middleware"
	"golang.org/x/crypto/bcrypt"

	"uva-devtest/models"
	"uva-devtest/persistence/daos/userdao"
	"uva-devtest/persistence/dbconnection"
	"uva-devtest/restapi/operations/user"
)

// Responds with a Status Internal Server Error
func serverErrorSignin(err error) middleware.Responder {
	log.Println(err)
	return user.NewRegisterUserBadRequest()
}

// Responds with a Conflict Error, user already exists
func conflictErrorSignin(err error) middleware.Responder {
	log.Println(err)
	errSt := "Usuario ya existe"
	prerr := models.Error{
		Message: &errSt,
	}
	return user.NewRegisterUserConflict().WithPayload(&prerr)
}

// Responds with a 201 Created object and the user created
func successSignin(u *userdao.User) middleware.Responder {
	log.Println("Usuario registrado")
	mu := &models.User{
		Username: u.Username,
		Email:    u.Email,
	}
	return user.NewRegisterUserCreated().WithPayload(mu)
}

// RegisterUser is the main handler function for Sign In functionality
// Param params Parametros de entrada que tiene la peticion http
// Return middleware.Responder
func RegisterUser(params user.RegisterUserParams) middleware.Responder {

	log.Println("Registrando usuario...")
	var lu *models.SigninUser = params.SigninUser
	log.Printf("Nombre de usuario: %v\n", *lu.Username)
	log.Println("Email: " + *lu.Email)
	bytes, err := bcrypt.GenerateFromPassword([]byte(*lu.Pass), 14)
	pwhashstring := string(bytes)
	if err != nil {
		return serverErrorSignin(err)
	}
	u := &userdao.User{
		Username: lu.Username,
		Email:    lu.Email,
		Pwhash:   &pwhashstring,
	}
	db, err := dbconnection.ConnectDb()

	if err != nil {
		return serverErrorSignin(err)
	}
	log.Println("Conectado a la base de datos")
	err = userdao.InsertUser(db, u)
	if err != nil {
		return conflictErrorSignin(err)
	}
	return successSignin(u)
}
