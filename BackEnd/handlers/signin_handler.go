// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package handlers handles the http request
package handlers

import (
	"errors"
	"log"
	"strings"

	"github.com/go-openapi/runtime/middleware"
	"golang.org/x/crypto/bcrypt"

	"uva-devtest/models"
	"uva-devtest/persistence/dao"
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
	errSt := "Ya existe un usuario registrado con ese nombre de usuario"
	if strings.Contains(err.Error(), "email") {
		errSt = "Ya existe un usuario registrado con ese email"
	}
	prerr := models.Error{
		Message: &errSt,
	}
	return user.NewRegisterUserConflict().WithPayload(&prerr)
}

// Responds with a 201 Created object and the user created
func successSignin(u *dao.User) middleware.Responder {
	log.Println("Usuario registrado")
	mu := dao.ToModelUser(u)
	return user.NewRegisterUserCreated().WithPayload(mu)
}

// Cost used in bcrypt.GenerateFromPassword functions
const Cost = 14

// RegisterUser is the main handler function for Sign In functionality
// Param params Parametros de entrada que tiene la peticion http
// Return middleware.Responder
func RegisterUser(params user.RegisterUserParams) middleware.Responder {

	log.Println("Registrando usuario...")
	var lu *models.SigninUser = params.SigninUser
	if lu == nil {
		return serverErrorSignin(errors.New("Parametros de entrada vacios"))
	}
	log.Printf("Nombre de usuario: %v\n", *lu.Username)
	log.Println("Email: " + *lu.Email)
	bytes, err := bcrypt.GenerateFromPassword([]byte(*lu.Pass), Cost)
	if err != nil {
		return serverErrorSignin(err)
	}
	pwhashstring := string(bytes)
	u := &dao.User{
		Username: lu.Username,
		Email:    lu.Email,
		Pwhash:   &pwhashstring,
		Type:     models.UserTypeStudent,
		Fullname: *lu.Username,
	}
	db, err := dbconnection.ConnectDb()

	if err != nil {
		return serverErrorSignin(err)
	}
	log.Println("Conectado a la base de datos")
	err = dao.InsertUser(db, u)
	if err != nil {
		return conflictErrorSignin(err)
	}
	return successSignin(u)
}
