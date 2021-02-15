package service

import (
  "encoding/json"
  "log"
  "net/http"
  "golang.org/x/crypto/bcrypt"

  "gitlab.com/HP-SCDS/Observatorio/2020-2021/uva-devtest/BackEnd/common/daos/userdao"
  "gitlab.com/HP-SCDS/Observatorio/2020-2021/uva-devtest/BackEnd/common/dbconnection"
  "gitlab.com/HP-SCDS/Observatorio/2020-2021/uva-devtest/BackEnd/common/jwtauth"
  "gitlab.com/HP-SCDS/Observatorio/2020-2021/uva-devtest/BackEnd/common/model"
  "gitlab.com/HP-SCDS/Observatorio/2020-2021/uva-devtest/BackEnd/common/response"
)

func serverError(w http.ResponseWriter, err error) {
  log.Println(err)
  response.RespondError(w, http.StatusInternalServerError, err.Error())
}

func badReqError(w http.ResponseWriter, err error) {
  log.Println(err)
  response.RespondError(w, http.StatusBadRequest, "Datos de log invalidos")
}

func authFailError(w http.ResponseWriter, err error, info string) {
  log.Println(err)
  response.RespondError(w, http.StatusGone, info)
}

func success(w http.ResponseWriter, u model.User) {
  log.Println("Usuario logged in")
  var wrap jwtauth.JwtWrapper
  wrap.SecretKey = u.PwHash
  wrap.Issuer = "DevTest"
  wrap.ExpirationHours = 24
  signedToken, err := wrap.GenerateToken(u.Email)
  tokenJson := jwtauth.CreateJwtJson(signedToken)
  if err != nil {
    badReqError(w, err)
  } else {
    response.RespondJSON(w, http.StatusOK, tokenJson)
  }
}

func Login(w http.ResponseWriter, r *http.Request) {
  log.Println("Generando Token JWT de usuario...")
  var lu model.LoginUser
  json.NewDecoder(r.Body).Decode(&lu)
  log.Printf("Nombre de usuario: %v\n", lu.Username)
  log.Println("Email: " + lu.Email)
    db, err := dbconnection.ConnectDb()

  if err != nil {
    serverError(w, err)
  } else {
    log.Println("Conectado a la base de datos")
    var u *model.User
    if lu.Username != "" {
      u, err = userdao.GetUserUsername(db, lu.Username)
    } else if lu.Email != "" {
      u, err = userdao.GetUserEmail(db, lu.Email)
    }
    if lu.Email == "" && lu.Username == "" {
      badReqError(w, nil)
    } else if err != nil {
      badReqError(w, err)
    } else if u == nil {
      authFailError(w, err, "Usuario no existe")
    } else {
      authErr := bcrypt.CompareHashAndPassword([]byte(u.PwHash), []byte(lu.Pass))
      if authErr != nil {
	authFailError(w, err, "Password incorrecto")
      } else {
	success(w, *u)
      }
    }
  }
}
