// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package handlers provides functions that handle http requests
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

func isTeacher(u *models.User) bool {
	return u.Type == models.UserTypeTeacher
}

func isTeacherOrAdmin(u *models.User) bool {
	return isAdmin(u) || isTeacher(u)
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
// Auth: Teacher or Admin
func GetUsers(params user.GetUsersParams, u *models.User) middleware.Responder {
	if !isTeacherOrAdmin(u) {
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
	return user.NewGetUsersOK().WithPayload(dao.ToModelsUser(us))
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
// Req: No puede quedarse sin users type Admin el sistema
func PutUser(params user.PutUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err != nil {
			log.Println("Error en users_handler PutUsers(): ", err)
			return user.NewPutUserInternalServerError()
		}
		admins, err := dao.GetAdmins(db)
		if err != nil {
			log.Println("Error en users_handler PutUsers(): ", err)
			return user.NewPutUserInternalServerError()
		}
		if len(admins) == 1 && admins[0].Username == &params.Username {
			if u.Type != models.UserTypeAdmin {
				log.Println("Error en users_handler PutUsers(), intento de quitar admin de ultimo admin")
				s := "Es el unico administrador existente"
				return user.NewPutUserConflict().WithPayload(&models.Error{Message: &s})
			}
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
// Req: No puede quedarse sin users type Admin el sistema
// Req: Eliminar usuario y todas relaciones. Si en algun equipo es unico admin, error
func DeleteUser(params user.DeleteUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err != nil {
			log.Println("Error en users_handler DeleteUser(): ", err)
			return user.NewDeleteUserInternalServerError()
		}
		admins, err := dao.GetAdmins(db)
		if err != nil {
			log.Println("Error en users_handler DeleteUser(): ", err)
			return user.NewDeleteUserInternalServerError()
		}
		if len(admins) == 1 && admins[0].Username == &params.Username {
			if u.Type != models.UserTypeAdmin {
				log.Println("Error en users_handler DeleteUser(), intento de borrar ultimo admin")
				s := "Es el unico administrador existente"
				return user.NewDeleteUserBadRequest().WithPayload(&models.Error{Message: &s})
			} // BadRequest en vez de Conflict ????
		}
		teams, err := dao.GetTeamsTeamRoleAdmin(db, params.Username)
		for _, team := range teams {
			admins, err := dao.GetTeamAdmins(db, *team.Teamname)
			if err != nil {
				log.Println("Error en users_handler DeleteUser(): ", err)
				return user.NewDeleteUserInternalServerError()
			}
			if len(admins) == 1 {
				log.Println("Error en users_handler DeleteUser(), intento de borrar ultimo admin de un equipo")
				s := "Es el unico administrador existente en equipo"
				return user.NewDeleteUserBadRequest().WithPayload(&models.Error{Message: &s})
				// BadRequest en vez de Conflict ????
			}
		}
		err = dao.DeleteUser(db, params.Username) // EN principio borra cascade
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
				return user.NewGetTeamsOfUserOK().WithPayload(dao.ToModelsTeams(teams))
			}
		}
		log.Println("Error en users_handler GetTeamsOfUser(): ", err)
		return user.NewGetTeamsOfUserInternalServerError()
	}
	return user.NewGetTeamsOfUserForbidden()
}

// AddTeamOfUser PUT /users/{username}/teams/{teamname}
// Auth: TeamAdmin or Admin
// DEBERIA devolver error o no modificar si ya exisite uno y hace PUT (no quitar de admin)
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
			log.Println("Error en users_handler AddTeamOfUser(): ", err)
			return user.NewAddTeamOfUserConflict()
		}
		log.Println("Error en users_handler AddTeamOfUser(): ", err)
		return user.NewAddTeamOfUserInternalServerError()
	}
	return user.NewAddTeamOfUserForbidden()
}

// DeleteTeamOfUser DELETE /users/{username}/teams/{teamname}
// Auth: Current User, TeamAdmin or Admin
// Req: No puede quedarse sin admins (ni miembros, pero el ultimo siempre sera admin)
func DeleteTeamOfUser(params user.DeleteTeamOfUserParams, u *models.User) middleware.Responder {
	teamAdmin, err := isTeamAdmin(params.Teamname, u)
	if err != nil {
		return user.NewDeleteTeamOfUserInternalServerError()
	}
	if userOrAdmin(params.Username, u) || teamAdmin {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			admins, err := dao.GetTeamAdmins(db, params.Teamname)
			if err != nil {
				log.Println("Error en users_handler DeleteTeamOfUser(): ", err)
				return user.NewDeleteTeamOfUserInternalServerError()
			}
			if len(admins) == 1 && admins[0].Username == &params.Username {
				log.Println("Error en users_handler DeleteTeamOfUser(): ", err)
				s := "Es el unico administrador existente en el equipo"
				return user.NewDeleteTeamOfUserBadRequest().WithPayload(&models.Error{Message: &s}) //Conflict???
			}
			err = dao.ExitUserTeam(db, params.Username, params.Teamname)
			if err == nil {
				return user.NewDeleteTeamOfUserOK()
			}
			log.Println("Error en users_handler DeleteTeamOfUser(): ", err)
			return user.NewDeleteTeamOfUserGone()
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
// Req: NO puede quedarse sin TeamAdmins el equipo
func PutUserTeamRole(params team.PutUserTeamRoleParams, u *models.User) middleware.Responder {
	teamAdmin, err := isTeamAdmin(params.Teamname, u)
	if err != nil {
		return team.NewPutUserTeamRoleInternalServerError()
	}
	if teamAdmin || isAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			role, err := dao.GetRole(db, params.Username, params.Teamname)
			if err != nil {
				return team.NewPutUserTeamRoleInternalServerError()
			}
			if *role.Role == models.TeamRoleRoleAdmin && *params.Role.Role != models.TeamRoleRoleAdmin {
				admins, err := dao.GetTeamAdmins(db, params.Teamname)
				if err != nil {
					log.Println("Error en users_handler PutUserTeamRole(): ", err)
					return team.NewPutUserTeamRoleInternalServerError()
				}
				if len(admins) == 1 {
					log.Println("Error en users_handler PutUserTeamRole(): ", err)
					s := "Es el unico administrador existente en el equipo"
					return team.NewPutUserTeamRoleBadRequest().WithPayload(&models.Error{Message: &s}) //Conflict???
				}
			}
			err = dao.UpdateRole(db, params.Username, params.Teamname, params.Role)
			if err == nil {
				return team.NewPutUserTeamRoleOK()
			}
		}
		log.Println("Error en users_handler PutUserTeamRole(): ", err)
		return team.NewPutUserTeamRoleInternalServerError()
	}
	return team.NewPutUserTeamRoleForbidden()
}
