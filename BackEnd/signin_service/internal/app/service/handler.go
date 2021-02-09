package service

import (
  "net/http"
  "log"
  "encoding/json"

  "gitlab.com/HP-SCDS/Observatorio/2020-2021/uva-devtest/BackEnd/common/model"
  "gitlab.com/HP-SCDS/Observatorio/2020-2021/uva-devtest/BackEnd/common/dbconnection"
  "gitlab.com/HP-SCDS/Observatorio/2020-2021/uva-devtest/BackEnd/common/daos/userdao"
  "gitlab.com/HP-SCDS/Observatorio/2020-2021/uva-devtest/BackEnd/common/response"
)

func serverError(w http.ResponseWriter, err error){
  log.Println(err)
  response.RespondError(w, http.StatusInternalServerError, err.Error())
}

func conflictError(w http.ResponseWriter, err error){
  log.Println(err)
  response.RespondError(w, http.StatusConflict, err.Error())
}

func success(w http.ResponseWriter, u model.User){
  log.Println("Usuario registrado")
  response.RespondJSON(w, http.StatusCreated, u)
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
  log.Println("Registrando usuario...")
  var u model.User
  json.NewDecoder(r.Body).Decode(&u)
  log.Printf("Nombre de usuario: %v\n", u.Username)
  log.Println("Email: "+u.Email)

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
