package service

import (
  "net/http"
  "log"
  "encoding/json"
  "gitlab.com/HP-SCDS/Observatorio/2020-2021/uva-devtest/BackEnd/common/model"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
  log.Println("Registrando usuario...")
  var u model.User
  json.NewDecoder(r.Body).Decode(&u)
  log.Printf("Nombre de usuario: %v\n", u.Username)
  log.Println("Email: "+u.Email)
}
