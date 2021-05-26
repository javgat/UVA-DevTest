// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package handlers provides functions that handle http requests
package handlers

import (
	"log"
	"uva-devtest/emailHelper"
	"uva-devtest/models"
	"uva-devtest/permissions"
	"uva-devtest/persistence/dao"
	"uva-devtest/persistence/dbconnection"
	"uva-devtest/restapi/operations/team"

	"github.com/go-openapi/runtime/middleware"
)

// GetTeams returns all teams GET /teams
// Auth: CanAdminTeams
func GetTeams(params team.GetTeamsParams, u *models.User) middleware.Responder {
	if permissions.CanAdminTeams(u) {
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
// Auth: CanTenerTeams
// Req: Meterle el usuario como TeamAdmin
func PostTeam(params team.PostTeamParams, u *models.User) middleware.Responder {
	if permissions.CanTenerTeams(u) {
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
// Auth: TeamMember or CanAdminTeams
func GetTeam(params team.GetTeamParams, u *models.User) middleware.Responder {
	if isTeamMember(params.Teamname, u) || permissions.CanAdminTeams(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			t, err := dao.GetTeam(db, params.Teamname)
			if err == nil && t != nil {
				return team.NewGetTeamOK().WithPayload(dao.ToModelTeam(t))
			}
		}
		log.Println("Error en teams_handler GetTeam(): ", err)
		return team.NewGetTeamInternalServerError()
	}
	return team.NewGetTeamForbidden()
}

// PutTeam updates team PUT /teams/{teamname}
// Auth: TeamAdmin or CanAdminTeams
func PutTeam(params team.PutTeamParams, u *models.User) middleware.Responder {
	teamAdmin, err := isTeamAdmin(params.Teamname, u)
	if err != nil {
		return team.NewPutTeamInternalServerError()
	}
	if teamAdmin || permissions.CanAdminTeams(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err := dao.UpdateTeam(db, params.Team, params.Teamname) // NO puede cambiar el soloProfesores
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
// Auth: TeamAdmin or CanAdminTeams
func DeleteTeam(params team.DeleteTeamParams, u *models.User) middleware.Responder {
	teamAdmin, err := isTeamAdmin(params.Teamname, u)
	if err != nil {
		return team.NewDeleteTeamInternalServerError()
	}
	if teamAdmin || permissions.CanAdminTeams(u) {
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
// Auth: TeamMember or CanAdminTeams
func GetAdminsFromTeam(params team.GetAdminsParams, u *models.User) middleware.Responder {
	if isTeamMember(params.Teamname, u) || permissions.CanAdminTeams(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var users []*dao.User
			users, err = dao.GetTeamAdmins(db, params.Teamname)
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
// Auth: TeamAdmin or CanAdminTeams
// DEBERIA: Si ya existe cambiar rol
// Req: username un user que CanTenerTeams
func AddAdminToTeam(params team.AddAdminParams, u *models.User) middleware.Responder {
	teamAdmin, err := isTeamAdmin(params.Teamname, u)
	if err != nil {
		return team.NewAddAdminInternalServerError()
	}
	if teamAdmin || permissions.CanAdminTeams(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var ous *dao.User
			ous, err = dao.GetUserUsername(db, params.Username)
			if err == nil {
				if ous == nil {
					return team.NewAddAdminGone()
				}
				mous := dao.ToModelUser(ous)
				if !permissions.CanTenerTeams(mous) {
					s := "No se puede añadir ese tipo de usuario como administrador"
					return team.NewAddAdminBadRequest().WithPayload(&models.Error{Message: &s})
				}
			}
			if canBeAddedTeam(params.Username, params.Teamname) {
				var us *dao.User
				us, err = dao.GetTeamMember(db, params.Teamname, params.Username)
				if err == nil {
					newAdmin := us == nil
					err = dao.AddUserTeamAdmin(db, params.Username, params.Teamname)
					if err == nil {
						if params.Message == nil || *params.Message.SendEmail {
							if newAdmin {
								emailHelper.SendEmailUserAddedToTeamAsAdmin(params.Username, params.Teamname, params.Message)
							} else {
								emailHelper.SendEmailUserChangedToTeamAdmin(params.Username, params.Teamname, params.Message)
							}
						}
						return team.NewAddAdminOK()
					}
				}
			} else {
				s := "No se puede añadir un estudiante a un equipo de solo profesores"
				return team.NewAddAdminBadRequest().WithPayload(&models.Error{Message: &s})
			}
		}
		log.Println("Error en teams_handler AddUserFromTeam(): ", err)
		return team.NewAddAdminInternalServerError()
	}
	return team.NewAddAdminForbidden()
}

// GetAdminFromTeam returns users from team GET /teams/{teamname}/admins/{username}
// Auth: TeamMember or CanAdminTeams
func GetAdminFromTeam(params team.GetAdminParams, u *models.User) middleware.Responder {
	if isTeamMember(params.Teamname, u) || permissions.CanAdminTeams(u) {
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
// Auth: TeamMember or CanAdminTeams
func GetMembersFromTeam(params team.GetMembersParams, u *models.User) middleware.Responder {
	if isTeamMember(params.Teamname, u) || permissions.CanAdminTeams(u) {
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

func canBeAddedTeam(username string, teamname string) bool {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		team, err := dao.GetTeam(db, teamname)
		if err == nil {
			if !*team.SoloProfesores {
				return true
			}
			user, err := dao.GetUserUsername(db, username)
			var muser *models.User
			if err == nil && user != nil {
				muser = dao.ToModelUser(user)
				if muser != nil {
					return *muser.Rol != models.UserRolEstudiante
				}
			}
		}
	}
	return false
}

// AddMemberToTeam adds user to team PUT /teams/{teamname}/members/{username}
// Auth: TeamAdmin or CanAdminTeams
// DEBERIA: Si ya existe cambiar rol
func AddMemberToTeam(params team.AddMemberParams, u *models.User) middleware.Responder {
	teamAdmin, err := isTeamAdmin(params.Teamname, u)
	if err != nil {
		return team.NewAddMemberInternalServerError()
	}
	if teamAdmin || permissions.CanAdminTeams(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			if canBeAddedTeam(params.Username, params.Teamname) {
				admins, err := dao.GetTeamAdmins(db, params.Teamname)
				if err == nil {
					if len(admins) == 1 && *admins[0].Username == params.Username {
						log.Println("Error en users_handler AddMemberToTeam(): ", err)
						s := "Es el unico administrador existente en el equipo"
						return team.NewAddMemberBadRequest().WithPayload(&models.Error{Message: &s}) //Conflict???
					}
					var us *dao.User
					us, err = dao.GetTeamAdmin(db, params.Teamname, params.Username)
					if err == nil {
						newMember := us == nil
						err = dao.AddUserTeamMember(db, params.Username, params.Teamname)
						if err == nil {
							if params.Message == nil || *params.Message.SendEmail {
								if newMember {
									emailHelper.SendEmailUserAddedToTeamAsMember(params.Username, params.Teamname, params.Message)
								} else {
									emailHelper.SendEmailUserChangedToTeamMember(params.Username, params.Teamname, params.Message)
								}
							}
							return team.NewAddMemberOK()
						}
					}
				}
			} else {
				s := "No se puede añadir un estudiante a un equipo de solo profesores"
				return team.NewAddMemberBadRequest().WithPayload(&models.Error{Message: &s})
			}
		}
		log.Println("Error en teams_handler AddUserFromTeam(): ", err)
		return team.NewAddMemberInternalServerError()
	}
	return team.NewAddMemberForbidden()
}

// GetMemberFromTeam returns users from team GET /teams/{teamname}/members/{username}
// Auth: TeamMember or CanAdminTeams
func GetMemberFromTeam(params team.GetMemberParams, u *models.User) middleware.Responder {
	if isTeamMember(params.Teamname, u) || permissions.CanAdminTeams(u) {
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
// Auth: TeamMember or CanAdminTeams
func GetUsersFromTeam(params team.GetUsersFromTeamParams, u *models.User) middleware.Responder {
	if isTeamMember(params.Teamname, u) || permissions.CanAdminTeams(u) {
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
// Auth: TeamMember or CanAdminTeams
func GetUserFromTeam(params team.GetUserFromTeamParams, u *models.User) middleware.Responder {
	if isTeamMember(params.Teamname, u) || permissions.CanAdminTeams(u) {
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
// Auth: Current User, TeamAdmin or CanAdminTeams
// Req: No quedarse sin TeamAdmins en Teams
func DeleteUserFromTeam(params team.DeleteUserFromTeamParams, u *models.User) middleware.Responder {
	teamAdmin, err := isTeamAdmin(params.Teamname, u)
	if err != nil {
		return team.NewDeleteUserFromTeamInternalServerError()
	}
	if teamAdmin || isUser(params.Username, u) || permissions.CanAdminTeams(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			admins, err := dao.GetTeamAdmins(db, params.Teamname)
			if err == nil {
				if len(admins) == 1 && *admins[0].Username == params.Username {
					log.Println("Error en users_handler DeleteUserFromTeam(): ", err)
					s := "Es el unico administrador existente en el equipo"
					return team.NewDeleteUserFromTeamBadRequest().WithPayload(&models.Error{Message: &s}) //Conflict???
				}
				err = dao.ExitUserTeam(db, params.Username, params.Teamname)
				if err == nil {
					return team.NewDeleteUserFromTeamOK()
				}
			}
		}
		log.Println("Error en teams_handler DeleteUserFromTeam(): ", err)
		return team.NewDeleteUserFromTeamInternalServerError()
	}
	return team.NewDeleteUserFromTeamForbidden()
}

// GetQuestionsFromTeam GET /teams/{teamname}/questions
// Auth: TeamMember OR CanAdminTeams
func GetQuestionsFromTeam(params team.GetQuestionsFromTeamParams, u *models.User) middleware.Responder {
	if isTeamMember(params.Teamname, u) || permissions.CanAdminTeams(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var questions []*dao.Question
			questions, err = dao.GetQuestionsFromTeam(db, params.Teamname, params.Tags, params.LikeTitle, params.Orderby,
				params.Limit, params.Offset)
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

// GetQuestionFromTeam GET /teams/{teamname}/questions/{questionid}
// Auth: TeamMember OR CanAdminTeams
func GetQuestionFromTeam(params team.GetQuestionFromTeamParams, u *models.User) middleware.Responder {
	if isTeamMember(params.Teamname, u) || permissions.CanAdminTeams(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var question *dao.Question
			question, err = dao.GetQuestionFromTeam(db, params.Teamname, params.Questionid)
			if err == nil && question != nil {
				mq, err := dao.ToModelQuestion(question)
				if err == nil && mq != nil {
					return team.NewGetQuestionFromTeamOK().WithPayload(mq)
				}
			}
			return team.NewGetQuestionFromTeamGone()
		}
		log.Println("Error en teams_handler GetUsersFromTeam(): ", err)
		return team.NewGetQuestionFromTeamInternalServerError()
	}
	return team.NewGetQuestionFromTeamForbidden()
}

// GetTestsFromTeam GET /teams/{teamname}/tests
// Auth: TeamMember OR CanAdminTeams
func GetTestsFromTeam(params team.GetTestsFromTeamParams, u *models.User) middleware.Responder {
	if isTeamMember(params.Teamname, u) || permissions.CanAdminTeams(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			tests, err := dao.GetTestsFromTeam(db, params.Teamname, params.Tags, params.LikeTitle, params.Orderby,
				params.Limit, params.Offset)
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
// Auth: TeamMember OR CanAdminTeams
func GetTestFromTeam(params team.GetTestFromTeamParams, u *models.User) middleware.Responder {
	if isTeamMember(params.Teamname, u) || permissions.CanAdminTeams(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			test, err := dao.GetTestFromTeam(db, params.Teamname, params.Testid)
			if err == nil && test != nil {
				mt, err := dao.ToModelTest(test)
				if err == nil && mt != nil {
					return team.NewGetTestFromTeamOK().WithPayload(mt)
				}
			}
			return team.NewGetTestFromTeamGone()
		}
		log.Println("Error en teams_handler GetUsersFromTeam(): ", err)
		return team.NewGetTestFromTeamInternalServerError()
	}
	return team.NewGetTestFromTeamForbidden()
}

// /teams/{teamname}/publishedTests
// Auth: TeamMember OR CanAdminTeams
func GetPTestsFromTeam(params team.GetPublishedTestsFromTeamParams, u *models.User) middleware.Responder {
	if isTeamMember(params.Teamname, u) || permissions.CanAdminTeams(u) {
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
		log.Println("Error en teams_handler GetPTestsFromTeam(): ", err)
		return team.NewGetPublishedTestsFromTeamInternalServerError()
	}
	return team.NewGetPublishedTestsFromTeamForbidden()
}

// /teams/{teamname}/publishedTests/{testid}
// Auth: TeamMember OR CanAdminTeams
func GetPTestFromTeam(params team.GetPublishedTestFromTeamParams, u *models.User) middleware.Responder {
	if isTeamMember(params.Teamname, u) || permissions.CanAdminTeams(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			test, err := dao.GetPTestFromTeam(db, params.Teamname, params.Testid)
			if err == nil && test != nil {
				mt, err := dao.ToModelTest(test)
				if err == nil && mt != nil {
					return team.NewGetPublishedTestFromTeamOK().WithPayload(mt)
				}
			}
			return team.NewGetPublishedTestFromTeamGone()
		}
		log.Println("Error en teams_handler GetPTestFromTeam(): ", err)
		return team.NewGetPublishedTestFromTeamInternalServerError()
	}
	return team.NewGetPublishedTestFromTeamForbidden()
}

// /teams/{teamname}/invitedTests
// Auth: TeamMember OR CanAdminTeams
func GetInvitedTestsFromTeam(params team.GetInvitedTestsFromTeamParams, u *models.User) middleware.Responder {
	if isTeamMember(params.Teamname, u) || permissions.CanAdminTeams(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			tests, err := dao.GetInvitedTestsFromTeam(db, params.Teamname, params.Tags, params.LikeTitle, params.Orderby,
				params.Limit, params.Offset)
			if err == nil {
				mt, err := dao.ToModelTests(tests)
				if err == nil && mt != nil {
					return team.NewGetInvitedTestsFromTeamOK().WithPayload(mt)
				}
				return team.NewGetInvitedTestsFromTeamGone()
			}
		}
		log.Println("Error en teams_handler GetInvitedTestsFromTeam(): ", err)
		return team.NewGetInvitedTestsFromTeamInternalServerError()
	}
	return team.NewGetInvitedTestsFromTeamForbidden()
}

// /teams/{teamname}/invitedTests/{publishedTestsid}
// Auth: TeamMember OR CanAdminTeams
func GetInvitedTestFromTeam(params team.GetInvitedTestFromTeamParams, u *models.User) middleware.Responder {
	if isTeamMember(params.Teamname, u) || permissions.CanAdminTeams(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			tests, err := dao.GetInvitedTestFromTeam(db, params.Teamname, params.Testid)
			if err == nil && tests != nil {
				mt, err := dao.ToModelTest(tests)
				if err == nil && mt != nil {
					return team.NewGetInvitedTestFromTeamOK().WithPayload(mt)
				}
			}
			return team.NewGetInvitedTestFromTeamGone()
		}
		log.Println("Error en teams_handler GetUsersFromTeam(): ", err)
		return team.NewGetInvitedTestFromTeamInternalServerError()
	}
	return team.NewGetInvitedTestFromTeamForbidden()
}
