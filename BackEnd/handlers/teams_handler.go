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

//TODO: Auth Check

// GetTeams returns all teams GET /teams
// Auth: Admin
func GetTeams(params team.GetTeamsParams, u *models.User) middleware.Responder {
	if isAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			teams, err := dao.GetTeams(db)
			if err == nil {
				return team.NewGetTeamsOK().WithPayload(dao.DaoToModelsTeams(teams))
			}
		}
		log.Println("Error en teams_handler GetTeams(): ", err)
		return team.NewGetTeamsInternalServerError()
	}
	return team.NewGetTeamsForbidden()
}

// PostTeam POST /teams
// Auth: Teacher or Admin
func PostTeam(params team.PostTeamParams, u *models.User) middleware.Responder {
	if isUser(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err := dao.PostTeam(db, params.Team)
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
	if isUser(u) { //De momento TODOS los iniciados pueden ver los equipos
		db, err := dbconnection.ConnectDb()
		if err == nil {
			t, err := dao.GetTeam(db, params.Teamname)
			if err == nil {
				return team.NewGetTeamOK().WithPayload(dao.DaoToModelTeam(t))
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
	if isUser(u) { //De momento TODOS los iniciados pueden ver los equipos
		// CAMBIAAAAAAAAAAAAR ACCESO
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
	if isUser(u) { //De momento TODOS los iniciados pueden ver los equipos
		// CAMBIAAAAAAAAAAAAR ACCESO
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err := dao.DeleteTeam(db, params.Teamname)
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
	if isUser(u) { //De momento TODOS los iniciados pueden ver los equipos
		// CAMBIAAAAAAAAAAAAR ACCESO
		db, err := dbconnection.ConnectDb()
		if err == nil {
			users, err := dao.GetUsersFromTeam(db, params.Teamname)
			if err == nil {
				return user.NewGetUsersFromTeamOK().WithPayload(users)
			}
		}
		log.Println("Error en teams_handler GetUsersFromTeam(): ", err)
		return user.NewGetUsersFromTeamInternalServerError()
	}
	return user.NewGetUsersFromTeamForbidden()
}

// AddUserFromTeam adds user to team PUT /teams/{teamname}/users/{username}
// Auth: TeamAdmin or Admin
func AddUserFromTeam(params user.AddUserFromTeamParams, u *models.User) middleware.Responder {
	if isUser(u) { //De momento TODOS los iniciados pueden ver los equipos
		// CAMBIAAAAAAAAAAAAR ACCESO
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err := dao.AddUserTeam(db, params.Username, params.Teamname)
			if err == nil {
				return user.NewAddUserFromTeamOK()
			}
		}
		log.Println("Error en teams_handler GetUsersFromTeam(): ", err)
		return user.NewAddUserFromTeamInternalServerError()
	}
	return user.NewAddUserFromTeamForbidden()
}

// DeleteUserFromTeam kicks user from team DELETE /teams/{teamname}/users/{username}
// Auth: TeamAdmin or Admin
func DeleteUserFromTeam(params user.DeleteUserFromTeamParams, u *models.User) middleware.Responder {
	if isUser(u) { //De momento TODOS los iniciados pueden ver los equipos
		// CAMBIAAAAAAAAAAAAR ACCESO
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err := dao.ExitUserTeam(db, params.Username, params.Teamname)
			if err == nil {
				return user.NewDeleteUserFromTeamOK()
			}
		}
		log.Println("Error en teams_handler GetUsersFromTeam(): ", err)
		return user.NewDeleteUserFromTeamInternalServerError()
	}
	return user.NewDeleteUserFromTeamForbidden()
}
