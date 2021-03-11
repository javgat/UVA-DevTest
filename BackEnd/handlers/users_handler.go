package handlers

import (
	"log"
	"uva-devtest/models"
	"uva-devtest/persistence/dao"
	"uva-devtest/persistence/dbconnection"
	"uva-devtest/restapi/operations/team"
	"uva-devtest/restapi/operations/user"

	"github.com/go-openapi/runtime/middleware"
	"golang.org/x/crypto/bcrypt"
)

func userOrAdmin(username string, u *models.User) bool {
	return username == *u.Username || u.Type == models.UserTypeAdmin
}

func isAdmin(u *models.User) bool {
	return u.Type == models.UserTypeAdmin
}

func isUser(u *models.User) bool {
	return u.Username != nil
}

func isTeamAdmin(teamname string, u *models.User) (bool, error) {
	db, err := dbconnection.ConnectDb()
	if err != nil {
		log.Println("Error en users_handler isTeamAdmin(): ", err)
		return false, err
	}
	role, err := dao.GetRole(db, *u.Username, teamname)
	if err != nil {
		log.Println("Error en users_handler isTeamAdmin(): ", err)
		return false, err
	}
	if *role.Role == models.TeamRoleRoleAdmin {
		return true, nil
	}
	return false, nil
}

func isTeamMember(teamname string, u *models.User) (bool, error) {
	db, err := dbconnection.ConnectDb()
	if err != nil {
		log.Println("Error en users_handler isTeamMember(): ", err)
		return false, err
	}
	users, err := dao.GetUsersFromTeam(db, teamname)
	if err != nil {
		log.Println("Error en users_handler isTeamMember(): ", err)
		return false, err
	}
	for _, us := range users {
		if us.Username == u.Username {
			return true, nil
		}
	}
	return false, nil
}

// PutPassword PUT /password/{username} Modifies the password of a user.
// Auth: Current User or Admin
func PutPassword(params user.PutPasswordParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		pu := params.PasswordUpdate
		db, err := dbconnection.ConnectDb()
		if err != nil {
			log.Println("Error en users_handler PutPassword(): ", err)
			return user.NewPutPasswordInternalServerError()
		}
		ud, _ := dao.GetUserUsername(db, params.Username)
		if bcrypt.CompareHashAndPassword([]byte(*ud.Pwhash), []byte(*pu.Oldpass)) == nil {
			bytes, errBcrypt := bcrypt.GenerateFromPassword([]byte(*pu.Newpass), Cost)
			newpwhash := string(bytes)
			err = dao.PutPasswordUsername(db, params.Username, newpwhash)
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

// GetUsers GET /users. Returns all users.
// Auth: Admin
func GetUsers(params user.GetUsersParams, u *models.User) middleware.Responder {
	if !isAdmin(u) {
		return user.NewGetUsersForbidden()
	}
	db, err := dbconnection.ConnectDb()
	if err != nil {
		log.Println("Error en users_handler GetUsers(): ", err)
		return user.NewGetUsersInternalServerError()
	}
	log.Println("Conectado a la base de datos")
	us, err := dao.GetUsers(db)
	if err != nil {
		log.Println("Error en users_handler GetUsers(): ", err)
		return user.NewGetUsersBadRequest()
	}
	return user.NewGetUsersOK().WithPayload(dao.DaoToModelsUser(us))
}

// GetUser GET /users/{username}
// Auth: Current User or Admin
func GetUser(params user.GetUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		return user.NewGetUserOK().WithPayload(u)
	}
	return user.NewGetUserForbidden()
}

// PutUser PUT /users/{username}
// Auth: Current User or Admin
// TODO: No puede quedarse sin admins !!!!
func PutUser(params user.PutUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err != nil {
			log.Println("Error en users_handler PutUsers(): ", err)
			return user.NewPutUserInternalServerError()
		}
		err = dao.UpdateUser(db, params.User, params.Username)
		if err != nil {
			log.Println("Error en users_handler PutUsers(): ", err)
			return user.NewPutUserGone()
		}
		return user.NewPutUserOK()
	}
	return user.NewPutUserForbidden()
}

// DeleteUser DELETE /users/{username}
// Auth: Current User or Admin
// TODO: Eliminar todas las relaciones de pertenencia en equipos !!!??
func DeleteUser(params user.DeleteUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err != nil {
			log.Println("Error en users_handler DeleteUser(): ", err)
			return user.NewDeleteUserInternalServerError()
		}
		err = dao.DeleteUser(db, params.Username)
		if err != nil {
			log.Println("Error en users_handler DeleteUser(): ", err)
			return user.NewDeleteUserGone()
		}
		return user.NewPutUserOK()
	}
	return user.NewDeleteUserForbidden()
}

// GetTeamsOfUser GET /users/{username}/teams
// Auth: Current User or Admin
func GetTeamsOfUser(params user.GetTeamsOfUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			teams, err := dao.GetTeamsUsername(db, params.Username)
			if err == nil && teams != nil {
				return user.NewGetTeamsOfUserOK().WithPayload(dao.DaoToModelsTeams(teams))
			}
		}
		log.Println("Error en users_handler GetTeamsOfUser(): ", err)
		return user.NewGetTeamsOfUserInternalServerError()
	}
	return user.NewGetTeamsOfUserForbidden()
}

// AddTeamOfUser PUT /users/{username}/teams/{teamname}
// Auth: TeamAdmin or Admin
func AddTeamOfUser(params user.AddTeamOfUserParams, u *models.User) middleware.Responder {
	teamAdmin, err := isTeamAdmin(params.Teamname, u)
	if err != nil {
		return user.NewAddTeamOfUserInternalServerError()
	}
	if teamAdmin || isAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.AddUserTeam(db, params.Username, params.Teamname)
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
// Auth: Current User, TeamAdmin or Admin
// TODO: No puede quedarse sin miembros!!!!!
func DeleteTeamOfUser(params user.DeleteTeamOfUserParams, u *models.User) middleware.Responder {
	teamAdmin, err := isTeamAdmin(params.Teamname, u)
	if err != nil {
		return user.NewDeleteTeamOfUserInternalServerError()
	}
	if userOrAdmin(params.Username, u) || teamAdmin {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.ExitUserTeam(db, params.Username, params.Teamname)
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
// Auth: Team or Admin
func GetUserTeamRole(params team.GetUserTeamRoleParams, u *models.User) middleware.Responder {
	teamMember, err := isTeamMember(params.Teamname, u)
	if err != nil {
		return team.NewGetUserTeamRoleInternalServerError()
	}
	if teamMember || isAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			role, err := dao.GetRole(db, params.Username, params.Teamname)
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
// Auth: TeamAdmin or Admin
// TODO: NO puede quedarse sin ADMINS!!!!
func PutUserTeamRole(params team.PutUserTeamRoleParams, u *models.User) middleware.Responder {
	teamAdmin, err := isTeamAdmin(params.Teamname, u)
	if err != nil {
		return team.NewPutUserTeamRoleInternalServerError()
	}
	if teamAdmin || isAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err := dao.UpdateRole(db, params.Username, params.Teamname, params.Role)
			if err == nil {
				return team.NewPutUserTeamRoleOK()
			}
		}
		log.Println("Error en users_handler PutUserTeamRole(): ", err)
		return team.NewPutUserTeamRoleInternalServerError()
	}
	return team.NewPutUserTeamRoleForbidden()
}
