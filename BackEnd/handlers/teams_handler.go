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

// PostTeam POST /users/{username}/teams
// Auth: Teacher or Admin
// Req: Meterle el usuario como TeamAdmin
func PostTeam(params team.PostTeamParams, u *models.User) middleware.Responder {
	if isTeacherOrAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err := dao.PostTeam(db, params.Team, params.Username)
			if err == nil {
				return team.NewPostTeamCreated().WithPayload(params.Team)
			} else {
				mess := err.Error()
				return team.NewPostTeamConflict().WithPayload(&models.Error{
					Message: &mess,
				})
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

// GetAdminsFromTeam returns users from team GET /teams/{teamname}/admins
// Auth: TeamMember or Admin
func GetAdminsFromTeam(params team.GetAdminsParams, u *models.User) middleware.Responder {
	teamMember, err := isTeamMember(params.Teamname, u)
	if err != nil {
		return team.NewGetAdminsInternalServerError()
	}
	if teamMember || isAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			users, err := dao.GetTeamAdmins(db, params.Teamname)
			if err == nil {
				return team.NewGetAdminsOK().WithPayload(dao.ToModelsUser(users))
			}
		}
		log.Println("Error en teams_handler GetUsersFromTeam(): ", err)
		return team.NewGetAdminsInternalServerError()
	}
	return team.NewGetAdminsForbidden()
}

// AddAdminToTeam adds user to team PUT /teams/{teamname}/admins/{username}
// Auth: TeamAdmin or Admin
// DEBERIA: Si ya existe cambiar rol
func AddAdminToTeam(params team.AddAdminParams, u *models.User) middleware.Responder {
	teamAdmin, err := isTeamAdmin(params.Teamname, u)
	if err != nil {
		return team.NewAddAdminInternalServerError()
	}
	if teamAdmin || isAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err := dao.AddUserTeamAdmin(db, params.Username, params.Teamname)
			if err == nil {
				return team.NewAddAdminOK()
			}
		}
		log.Println("Error en teams_handler AddUserFromTeam(): ", err)
		return team.NewAddAdminInternalServerError()
	}
	return team.NewAddAdminForbidden()
}

// GetAdminFromTeam returns users from team GET /teams/{teamname}/admins/{username}
// Auth: TeamMember or Admin
func GetAdminFromTeam(params team.GetAdminParams, u *models.User) middleware.Responder {
	teamMember, err := isTeamMember(params.Teamname, u)
	if err != nil {
		return team.NewGetAdminInternalServerError()
	}
	if teamMember || isAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			user, err := dao.GetTeamAdmin(db, params.Teamname, params.Username)
			if err == nil {
				return team.NewGetAdminOK().WithPayload(dao.ToModelUser(user))
			}
		}
		log.Println("Error en teams_handler GetUsersFromTeam(): ", err)
		return team.NewGetAdminInternalServerError()
	}
	return team.NewGetAdminForbidden()
}

// GetMembersFromTeam returns users from team GET /teams/{teamname}/members
// Auth: TeamMember or Admin
func GetMembersFromTeam(params team.GetMembersParams, u *models.User) middleware.Responder {
	teamMember, err := isTeamMember(params.Teamname, u)
	if err != nil {
		return team.NewGetMembersInternalServerError()
	}
	if teamMember || isAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			users, err := dao.GetTeamMembers(db, params.Teamname)
			if err == nil {
				return team.NewGetMembersOK().WithPayload(dao.ToModelsUser(users))
			}
		}
		log.Println("Error en teams_handler GetUsersFromTeam(): ", err)
		return team.NewGetMembersInternalServerError()
	}
	return team.NewGetMembersForbidden()
}

// AddMemberToTeam adds user to team PUT /teams/{teamname}/members/{username}
// Auth: TeamMember or Member
// DEBERIA: Si ya existe cambiar rol
func AddMemberToTeam(params team.AddMemberParams, u *models.User) middleware.Responder {
	teamAdmin, err := isTeamAdmin(params.Teamname, u)
	if err != nil {
		return team.NewAddMemberInternalServerError()
	}
	if teamAdmin || isAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err := dao.AddUserTeamMember(db, params.Username, params.Teamname)
			if err == nil {
				return team.NewAddMemberOK()
			}
		}
		log.Println("Error en teams_handler AddUserFromTeam(): ", err)
		return team.NewAddMemberInternalServerError()
	}
	return team.NewAddMemberForbidden()
}

// GetMemberFromTeam returns users from team GET /teams/{teamname}/members/{username}
// Auth: TeamMember or Admin
func GetMemberFromTeam(params team.GetMemberParams, u *models.User) middleware.Responder {
	teamMember, err := isTeamMember(params.Teamname, u)
	if err != nil {
		return team.NewGetMemberInternalServerError()
	}
	if teamMember || isAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			user, err := dao.GetTeamMember(db, params.Teamname, params.Username)
			if err == nil {
				return team.NewGetMemberOK().WithPayload(dao.ToModelUser(user))
			}
		}
		log.Println("Error en teams_handler GetUsersFromTeam(): ", err)
		return team.NewGetMemberInternalServerError()
	}
	return team.NewGetMemberForbidden()
}

// GetUsersFromTeam returns users from team GET /teams/{teamname}/users
// Auth: TeamMember or Admin
func GetUsersFromTeam(params team.GetUsersFromTeamParams, u *models.User) middleware.Responder {
	teamMember, err := isTeamMember(params.Teamname, u)
	if err != nil {
		return team.NewGetUsersFromTeamInternalServerError()
	}
	if teamMember || isAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			users, err := dao.GetUsersFromTeam(db, params.Teamname)
			if err == nil {
				return team.NewGetUsersFromTeamOK().WithPayload(dao.ToModelsUser(users))
			}
		}
		log.Println("Error en teams_handler GetUsersFromTeam(): ", err)
		return team.NewGetUsersFromTeamInternalServerError()
	}
	return team.NewGetUsersFromTeamForbidden()
}

// GetUserFromTeam returns users from team GET /teams/{teamname}/users/{username}
// Auth: TeamMember or Admin
func GetUserFromTeam(params team.GetUserFromTeamParams, u *models.User) middleware.Responder {
	teamMember, err := isTeamMember(params.Teamname, u)
	if err != nil {
		return team.NewGetUserFromTeamInternalServerError()
	}
	if teamMember || isAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			user, err := dao.GetUserFromTeam(db, params.Teamname, params.Username)
			if err == nil {
				return team.NewGetUserFromTeamOK().WithPayload(dao.ToModelUser(user))
			}
		}
		log.Println("Error en teams_handler GetUsersFromTeam(): ", err)
		return team.NewGetUserFromTeamInternalServerError()
	}
	return team.NewGetAdminForbidden()
}

// DeleteUserFromTeam kicks user from team DELETE /teams/{teamname}/users/{username}
// Auth: Current User, TeamAdmin or Admin
// Req: No quedarse sin TeamAdmins en Teams
func DeleteUserFromTeam(params team.DeleteUserFromTeamParams, u *models.User) middleware.Responder {
	teamAdmin, err := isTeamAdmin(params.Teamname, u)
	if err != nil {
		return user.NewDeleteUserFromTeamInternalServerError()
	}
	if teamAdmin || userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			admins, err := dao.GetTeamAdmins(db, params.Teamname)
			if err == nil {
				if len(admins) == 1 && *admins[0].Username == params.Username {
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

// /teams/{teamname}/questions
func GetQuestionsFromTeam(params team.GetQuestionsFromTeamParams, u *models.User) middleware.Responder {
	teamMember, err := isTeamMember(params.Teamname, u)
	if err != nil {
		return team.NewGetQuestionsFromTeamInternalServerError()
	}
	if teamMember || isAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			questions, err := dao.GetQuestionsFromTeam(db, params.Teamname)
			if err == nil {
				mq, err := dao.ToModelQuestions(questions)
				if err == nil && mq != nil {
					return team.NewGetQuestionsFromTeamOK().WithPayload(mq)
				}
				return team.NewGetQuestionsFromTeamGone()
			}
		}
		log.Println("Error en teams_handler GetUsersFromTeam(): ", err)
		return team.NewGetQuestionsFromTeamInternalServerError()
	}
	return team.NewGetQuestionsFromTeamForbidden()
}

// /teams/{teamname}/questions/{questionid}
func GetQuestionFromTeam(params team.GetQuestionFromTeamParams, u *models.User) middleware.Responder {
	teamMember, err := isTeamMember(params.Teamname, u)
	if err != nil {
		return team.NewGetQuestionFromTeamInternalServerError()
	}
	if teamMember || isAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			question, err := dao.GetQuestionFromTeam(db, params.Teamname, params.Questionid)
			if err == nil {
				mq, err := dao.ToModelQuestion(question)
				if err == nil && mq != nil {
					return team.NewGetQuestionFromTeamOK().WithPayload(mq)
				}
				return team.NewGetQuestionFromTeamGone()
			}
		}
		log.Println("Error en teams_handler GetUsersFromTeam(): ", err)
		return team.NewGetQuestionFromTeamInternalServerError()
	}
	return team.NewGetQuestionFromTeamForbidden()
}

// /teams/{teamname}/tests
func GetTestsFromTeam(params team.GetTestsFromTeamParams, u *models.User) middleware.Responder {
	teamMember, err := isTeamMember(params.Teamname, u)
	if err != nil {
		return team.NewGetTestsFromTeamInternalServerError()
	}
	if teamMember || isAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			tests, err := dao.GetTestsFromTeam(db, params.Teamname)
			if err == nil {
				mt, err := dao.ToModelTests(tests)
				if err == nil && mt != nil {
					return team.NewGetTestsFromTeamOK().WithPayload(mt)
				}
				return team.NewGetTestsFromTeamGone()
			}
		}
		log.Println("Error en teams_handler GetUsersFromTeam(): ", err)
		return team.NewGetTestsFromTeamInternalServerError()
	}
	return team.NewGetTestsFromTeamForbidden()
}

// /teams/{teamname}/tests/{testid}
func GetTestFromTeam(params team.GetTestFromTeamParams, u *models.User) middleware.Responder {
	teamMember, err := isTeamMember(params.Teamname, u)
	if err != nil {
		return team.NewGetTestFromTeamInternalServerError()
	}
	if teamMember || isAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			test, err := dao.GetTestFromTeam(db, params.Teamname, params.Testid)
			if err == nil {
				mt, err := dao.ToModelTest(test)
				if err == nil && mt != nil {
					return team.NewGetTestFromTeamOK().WithPayload(mt)
				}
				return team.NewGetTestFromTeamGone()
			}
		}
		log.Println("Error en teams_handler GetUsersFromTeam(): ", err)
		return team.NewGetTestFromTeamInternalServerError()
	}
	return team.NewGetTestFromTeamForbidden()
}

// /teams/{teamname}/publishedTests
func GetPTestsFromTeam(params team.GetPublishedTestsFromTeamParams, u *models.User) middleware.Responder {
	teamMember, err := isTeamMember(params.Teamname, u)
	if err != nil {
		return team.NewGetPublishedTestsFromTeamInternalServerError()
	}
	if teamMember || isAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			tests, err := dao.GetPTestsFromTeam(db, params.Teamname)
			if err == nil {
				mt, err := dao.ToModelTests(tests)
				if err == nil && mt != nil {
					return team.NewGetPublishedTestsFromTeamOK().WithPayload(mt)
				}
				return team.NewGetPublishedTestsFromTeamGone()
			}
		}
		log.Println("Error en teams_handler GetUsersFromTeam(): ", err)
		return team.NewGetPublishedTestsFromTeamInternalServerError()
	}
	return team.NewGetPublishedTestsFromTeamForbidden()
}

// /teams/{teamname}/publishedTests/{publishedTestsid}
func GetPTestFromTeam(params team.GetPublishedTestFromTeamParams, u *models.User) middleware.Responder {
	teamMember, err := isTeamMember(params.Teamname, u)
	if err != nil {
		return team.NewGetPublishedTestFromTeamInternalServerError()
	}
	if teamMember || isAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			tests, err := dao.GetPTestFromTeam(db, params.Teamname, params.Testid)
			if err == nil {
				mt, err := dao.ToModelTest(tests)
				if err == nil && mt != nil {
					return team.NewGetPublishedTestFromTeamOK().WithPayload(mt)
				}
				return team.NewGetPublishedTestFromTeamGone()
			}
		}
		log.Println("Error en teams_handler GetUsersFromTeam(): ", err)
		return team.NewGetPublishedTestFromTeamInternalServerError()
	}
	return team.NewGetPublishedTestFromTeamForbidden()
}
