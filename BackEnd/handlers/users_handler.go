package handlers

import (
	"log"
	"uva-devtest/models"
	"uva-devtest/persistence/daos/roledao"
	"uva-devtest/persistence/daos/teamdao"
	"uva-devtest/persistence/daos/userdao"
	"uva-devtest/persistence/dbconnection"
	"uva-devtest/restapi/operations/team"
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

// DeleteUser DELETE /users/{username}
func DeleteUser(params user.DeleteUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err != nil {
			log.Println("Error en users_handler DeleteUser(): ", err)
			return user.NewDeleteUserInternalServerError()
		}
		err = userdao.DeleteUser(db, params.Username)
		if err != nil {
			log.Println("Error en users_handler DeleteUser(): ", err)
			return user.NewDeleteUserGone()
		}
		return user.NewPutUserOK()
	}
	return user.NewDeleteUserForbidden()
}

// GetTeamsOfUser GET /users/{username}/teams
func GetTeamsOfUser(params user.GetTeamsOfUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			teams, err := teamdao.GetTeamsUsername(db, params.Username)
			if err == nil && teams != nil {
				return user.NewGetTeamsOfUserOK().WithPayload(teams)
			}
		}
		log.Println("Error en users_handler GetTeamsOfUser(): ", err)
		return user.NewGetTeamsOfUserInternalServerError()
	}
	return user.NewGetTeamsOfUserForbidden()
}

// AddTeamOfUser PUT /users/{username}/teams/{teamname}
func AddTeamOfUser(params user.AddTeamOfUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = userdao.AddUserTeam(db, params.Username, params.Teamname)
			if err == nil {
				return user.NewAddTeamOfUserOK()
			}
		}
		log.Println("Error en users_handler AddTeamOfUser(): ", err)
		return user.NewAddTeamOfUserInternalServerError()
	}
	return user.NewAddTeamOfUserForbidden()
}

// DeleteTeamOfUser DELETE /users/{username}/teams/{teamname}
func DeleteTeamOfUser(params user.DeleteTeamOfUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = userdao.ExitUserTeam(db, params.Username, params.Teamname)
			if err == nil {
				return user.NewDeleteTeamOfUserOK()
			}
		}
		log.Println("Error en users_handler DeleteTeamOfUser(): ", err)
		return user.NewDeleteTeamOfUserInternalServerError()
	}
	return user.NewDeleteTeamOfUserForbidden()
}

// GetUserTeamRole GET /users/{username}/teams/{teamname}/role
func GetUserTeamRole(params team.GetUserTeamRoleParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			role, err := roledao.GetRole(db, params.Username, params.Teamname)
			if err == nil {
				return team.NewGetUserTeamRoleOK().WithPayload(role)
			}
		}
		log.Println("Error en users_handler GetUserTeamRole(): ", err)
		return team.NewGetUserTeamRoleInternalServerError()
	}
	return team.NewGetUserTeamRoleForbidden()
}

// PutUserTeamRole PUT /users/{username}/teams/{teamname}/role
func PutUserTeamRole(params team.PutUserTeamRoleParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err := roledao.UpdateRole(db, params.Username, params.Teamname, params.Role)
			if err == nil {
				return team.NewPutUserTeamRoleOK()
			}
		}
		log.Println("Error en users_handler PutUserTeamRole(): ", err)
		return team.NewPutUserTeamRoleInternalServerError()
	}
	return team.NewPutUserTeamRoleForbidden()
}
