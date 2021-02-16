// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package service handles the http request
package service

import (
  "net/http"
  "log"
  "encoding/json"
  "golang.org/x/crypto/bcrypt"

  "gitlab.com/HP-SCDS/Observatorio/2020-2021/uva-devtest/BackEnd/common/model"
  "gitlab.com/HP-SCDS/Observatorio/2020-2021/uva-devtest/BackEnd/common/dbconnection"
  "gitlab.com/HP-SCDS/Observatorio/2020-2021/uva-devtest/BackEnd/common/daos/userdao"
  "gitlab.com/HP-SCDS/Observatorio/2020-2021/uva-devtest/BackEnd/common/response"
)

// Responds with a Status Internal Server Error
func serverError(w http.ResponseWriter, err error){
  log.Println(err)
  response.RespondError(w, http.StatusInternalServerError, "Error interno del servidor")
}

// Responds with a Conflict Error, user already exists
func conflictError(w http.ResponseWriter, err error){
  log.Println(err)
  response.RespondError(w, http.StatusConflict, "Usuario ya existe")
}

// Responds with a 201 Created object and the user created
func success(w http.ResponseWriter, u model.User){
  log.Println("Usuario registrado")
  response.RespondJSON(w, http.StatusCreated, u)
}

// Main handler function
func RegisterUser(w http.ResponseWriter, r *http.Request) {
  log.Println("Registrando usuario...")
  var lu SigninUser
  json.NewDecoder(r.Body).Decode(&lu)
  log.Printf("Nombre de usuario: %v\n", lu.Username)
  log.Println("Email: "+lu.Email)
  bytes, err := bcrypt.GenerateFromPassword([]byte(lu.Pass), 14)
  if err != nil{
    serverError(w, err)
  }else{
    u := model.User{
      Username: lu.Username,
      Email: lu.Email,
      PwHash: string(bytes),
    }
    db, err := dbconnection.ConnectDb()

    if err != nil {
      serverError(w, err)
    }else{
      log.Println("Conectado a la base de datos")
      err = userdao.InsertUser(db, u)
      if err != nil {
        conflictError(w, err)
      }else{
        success(w, u)
      }
    }
  }
}
