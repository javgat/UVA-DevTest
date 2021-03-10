package handlers

import (
	"log"
	"uva-devtest/models"
	"uva-devtest/persistence/daos/userdao"
	"uva-devtest/persistence/dbconnection"
	"uva-devtest/restapi/operations/user"

	"github.com/go-openapi/runtime/middleware"
	"golang.org/x/crypto/bcrypt"
)

func userOrAdmin(username string, u *models.User) bool {
	return username == *u.Username || u.Type == models.UserTypeAdmin
}

// PutPassword PUT /password/{username} Modifies the password of a user
func PutPassword(params user.PutPasswordParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		pu := params.PasswordUpdate
		db, err := dbconnection.ConnectDb()
		if err != nil {
			log.Println("Error en users_handler PutPassword(): ", err)
			return user.NewPutPasswordInternalServerError()
		}
		ud, _ := userdao.GetUserUsername(db, params.Username)
		if bcrypt.CompareHashAndPassword([]byte(*ud.Pwhash), []byte(*pu.Oldpass)) == nil {
			bytes, errBcrypt := bcrypt.GenerateFromPassword([]byte(*pu.Newpass), Cost)
			newpwhash := string(bytes)
			err = userdao.PutPasswordUsername(db, params.Username, newpwhash)
			if err != nil || errBcrypt != nil {
				log.Println("Error al modificar la contraseña: ", err, errBcrypt)
				return user.NewPutPasswordInternalServerError()
			}
			return user.NewPutPasswordOK()
		}
		return user.NewPutPasswordBadRequest()
	}
	return user.NewPutPasswordForbidden()
}

// GetUsers GET /users. Returns all users
func GetUsers(params user.GetUsersParams, u *models.User) middleware.Responder {
	if u.Type != "admin" {
		return user.NewGetUsersForbidden()
	}
	db, err := dbconnection.ConnectDb()
	if err != nil {
		log.Println("Error en users_handler GetUsers(): ", err)
		return user.NewGetUsersInternalServerError()
	}
	log.Println("Conectado a la base de datos")
	us, err := userdao.GetUsers(db)
	if err != nil {
		log.Println("Error en users_handler GetUsers(): ", err)
		return user.NewGetUsersBadRequest()
	}
	return user.NewGetUsersOK().WithPayload(userdao.DaoToModelsUser(us))
}

// GetUser GET /users/{username}
func GetUser(params user.GetUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		return user.NewGetUserOK().WithPayload(u)
	}
	return user.NewGetUserForbidden()
}

// PutUser PUT /users/{username}
func PutUser(params user.PutUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err != nil {
			log.Println("Error en users_handler PutUsers(): ", err)
			return user.NewPutUserInternalServerError()
		}
		err = userdao.UpdateUser(db, params.User, params.Username)
		if err != nil {
			log.Println("Error en users_handler PutUsers(): ", err)
			return user.NewPutUserGone()
		}
		return user.NewPutUserOK()
	}
	return user.NewPutUserForbidden()
}
