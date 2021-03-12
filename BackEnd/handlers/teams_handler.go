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
)

// GetTeams returns all teams GET /teams
// Auth: Admin
func GetTeams(params team.GetTeamsParams, u *models.User) middleware.Responder {
	if isAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			teams, err := dao.GetTeams(db)
			if err == nil {
				return team.NewGetTeamsOK().WithPayload(dao.ToModelsTeams(teams))
			}
		}
		log.Println("Error en teams_handler GetTeams(): ", err)
		return team.NewGetTeamsInternalServerError()
	}
	return team.NewGetTeamsForbidden()
}

// PostTeam POST /teams
// Auth: Teacher or Admin
// Req: Meterle el usuario como TeamAdmin
func PostTeam(params team.PostTeamParams, u *models.User) middleware.Responder {
	if isTeacherOrAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err := dao.PostTeam(db, params.Team, *u.Username)
			if err == nil {
				return team.NewPostTeamCreated().WithPayload(params.Team)
			}
		}
		log.Println("Error en teams_handler PostTeam(): ", err)
		return team.NewPostTeamInternalServerError()
	}
	return team.NewPostTeamForbidden()
}

// GetTeam returns team GET /teams/{teamname}
// Auth: TeamMember or Admin
func GetTeam(params team.GetTeamParams, u *models.User) middleware.Responder {
	teamMember, err := isTeamMember(params.Teamname, u)
	if err != nil {
		return team.NewGetTeamInternalServerError()
	}
	if teamMember || isAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			t, err := dao.GetTeam(db, params.Teamname)
			if err == nil {
				return team.NewGetTeamOK().WithPayload(dao.ToModelTeam(t))
			}
		}
		log.Println("Error en teams_handler GetTeam(): ", err)
		return team.NewGetTeamInternalServerError()
	}
	return team.NewGetTeamForbidden()
}

// PutTeam updates team PUT /teams/{teamname}
// Auth: TeamAdmin or Admin
func PutTeam(params team.PutTeamParams, u *models.User) middleware.Responder {
	teamAdmin, err := isTeamAdmin(params.Teamname, u)
	if err != nil {
		return team.NewPutTeamInternalServerError()
	}
	if teamAdmin || isAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err := dao.UpdateTeam(db, params.Team, params.Teamname)
			if err == nil {
				return team.NewPutTeamOK()
			}
		}
		log.Println("Error en teams_handler PutTeam(): ", err)
		return team.NewPutTeamInternalServerError()
	}
	return team.NewPutTeamForbidden()
}

// DeleteTeam deletes team DELETE /teams/{teamname}
// Auth: TeamAdmin or Admin
func DeleteTeam(params team.DeleteTeamParams, u *models.User) middleware.Responder {
	teamAdmin, err := isTeamAdmin(params.Teamname, u)
	if err != nil {
		return team.NewDeleteTeamInternalServerError()
	}
	if teamAdmin || isAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err := dao.DeleteTeam(db, params.Teamname) // En principio borra en cascade las relaciones
			if err == nil {
				return team.NewDeleteTeamOK()
			}
		}
		log.Println("Error en teams_handler DeleteTeam(): ", err)
		return team.NewDeleteTeamInternalServerError()
	}
	return team.NewDeleteTeamForbidden()
}

// GetUsersFromTeam returns users from team GET /teams/{teamname}/users
// Auth: TeamMember or Admin
func GetUsersFromTeam(params user.GetUsersFromTeamParams, u *models.User) middleware.Responder {
	teamMember, err := isTeamMember(params.Teamname, u)
	if err != nil {
		return user.NewGetUsersFromTeamInternalServerError()
	}
	if teamMember || isAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			users, err := dao.GetUsersFromTeam(db, params.Teamname)
			if err == nil {
				return user.NewGetUsersFromTeamOK().WithPayload(dao.ToModelsUser(users))
			}
		}
		log.Println("Error en teams_handler GetUsersFromTeam(): ", err)
		return user.NewGetUsersFromTeamInternalServerError()
	}
	return user.NewGetUsersFromTeamForbidden()
}

// AddUserFromTeam adds user to team PUT /teams/{teamname}/users/{username}
// Auth: TeamAdmin or Admin
// DEBERIA: Si ya existe que pete o que no se cargue el rol (no quite de TeamAdmin)
func AddUserFromTeam(params user.AddUserFromTeamParams, u *models.User) middleware.Responder {
	teamAdmin, err := isTeamAdmin(params.Teamname, u)
	if err != nil {
		return user.NewAddUserFromTeamInternalServerError()
	}
	if teamAdmin || isAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err := dao.AddUserTeam(db, params.Username, params.Teamname)
			if err == nil {
				return user.NewAddUserFromTeamOK()
			}
		}
		log.Println("Error en teams_handler AddUserFromTeam(): ", err)
		return user.NewAddUserFromTeamInternalServerError()
	}
	return user.NewAddUserFromTeamForbidden()
}

// DeleteUserFromTeam kicks user from team DELETE /teams/{teamname}/users/{username}
// Auth: Current User, TeamAdmin or Admin
// Req: No quedarse sin TeamAdmins en Teams
func DeleteUserFromTeam(params user.DeleteUserFromTeamParams, u *models.User) middleware.Responder {
	teamAdmin, err := isTeamAdmin(params.Teamname, u)
	if err != nil {
		return user.NewDeleteUserFromTeamInternalServerError()
	}
	if teamAdmin || userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			admins, err := dao.GetTeamAdmins(db, params.Teamname)
			if err == nil {
				if len(admins) == 1 && admins[0].Username == &params.Username {
					log.Println("Error en users_handler DeleteUserFromTeam(): ", err)
					s := "Es el unico administrador existente en el equipo"
					return user.NewDeleteUserFromTeamBadRequest().WithPayload(&models.Error{Message: &s}) //Conflict???
				}
				err = dao.ExitUserTeam(db, params.Username, params.Teamname)
				if err == nil {
					return user.NewDeleteUserFromTeamOK()
				}
			}
		}
		log.Println("Error en teams_handler DeleteUserFromTeam(): ", err)
		return user.NewDeleteUserFromTeamInternalServerError()
	}
	return user.NewDeleteUserFromTeamForbidden()
}
