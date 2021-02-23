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
	//response.RespondError(w, http.StatusInternalServerError, "Error interno del servidor")
}

// Responds with a Conflict Error, user already exists
func conflictErrorSignin(err error) middleware.Responder {
	log.Println(err)
	errSt := "Usuario ya existe"
	prerr := models.Error{
		Message: &errSt,
	}
	return user.NewRegisterUserConflict().WithPayload(&prerr)
	//response.RespondError(w, http.StatusConflict, "Usuario ya existe")
}

// Responds with a 201 Created object and the user created
func successSignin(u *models.User) middleware.Responder {
	log.Println("Usuario registrado")
	return user.NewRegisterUserCreated().WithPayload(u)
	//response.RespondJSON(w, http.StatusCreated, u)
}

// Main handler function
func RegisterUser(params user.RegisterUserParams) middleware.Responder {

	log.Println("Registrando usuario...")
	var lu *models.SigninUser = params.SigninUser
	log.Printf("Nombre de usuario: %v\n", *lu.Username)
	log.Println("Email: " + *lu.Email)
	bytes, err := bcrypt.GenerateFromPassword([]byte(*lu.Pass), 14)
	pwhashstring := string(bytes)
	if err != nil {
		return serverErrorSignin(err)
	} else {
		u := &models.User{
			Username: lu.Username,
			Email:    lu.Email,
			Pwhash:   &pwhashstring,
		}
		db, err := dbconnection.ConnectDb()

		if err != nil {
			return serverErrorSignin(err)
		} else {
			log.Println("Conectado a la base de datos")
			err = userdao.InsertUser(db, u)
			if err != nil {
				return conflictErrorSignin(err)
			} else {
				return successSignin(u)
			}
		}
	}
}
