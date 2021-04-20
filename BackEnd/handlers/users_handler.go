// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package handlers provides functions that handle http requests
package handlers

import (
	"log"
	"uva-devtest/models"
	"uva-devtest/persistence/dao"
	"uva-devtest/persistence/dbconnection"
	"uva-devtest/restapi/operations/user"

	"github.com/go-openapi/runtime/middleware"
	"golang.org/x/crypto/bcrypt"
)

func userOrAdmin(username string, u *models.User) bool {
	return username == *u.Username || *u.Rol == models.UserRolAdministrador
}

func isAdmin(u *models.User) bool {
	return *u.Rol == models.UserRolAdministrador
}

func isTeacher(u *models.User) bool {
	return *u.Rol == models.UserRolProfesor
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
	if role != nil && *role.Role == dao.TeamRoleRoleAdmin {
		return true, nil
	}
	return false, nil
}

func isTeamMember(teamname string, u *models.User) bool {
	db, err := dbconnection.ConnectDb()
	if err != nil {
		return false
	}
	user, err := dao.GetUserFromTeam(db, teamname, *u.Username)
	if err != nil || user == nil {
		return false
	}
	return true
}

// GetUsers GET /users. Returns all users.
// Auth: All
func GetUsers(params user.GetUsersParams, u *models.User) middleware.Responder {
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
// Auth: All
func GetUser(params user.GetUserParams, u *models.User) middleware.Responder {
	if *u.Username == params.Username {
		return user.NewGetUserOK().WithPayload(u)
	}
	db, err := dbconnection.ConnectDb()
	if err != nil {
		log.Println("Error en users_handler GetUser(): ", err)
		return user.NewGetUserInternalServerError()
	}
	us, err := dao.GetUserUsername(db, params.Username)
	if err != nil {
		log.Println("Error en users_handler GetUser(): ", err)
		return user.NewGetUserInternalServerError()
	} else if us == nil {
		us, err = dao.GetUserEmail(db, params.Username)
		if err != nil || us == nil {
			log.Println("No existe user en users_handler GetUser(): ", err)
			return user.NewGetUserGone() //410
		}
	}
	return user.NewGetUserOK().WithPayload(dao.ToModelUser(us))
}

func userUpdateToUser(uu *models.UserUpdate) *models.User {
	u := &models.User{
		Email:    uu.Email,
		Fullname: uu.Fullname,
		Username: uu.Username,
	}
	return u
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
			if *u.Rol != models.UserRolAdministrador {
				log.Println("Error en users_handler PutUsers(), intento de quitar admin de ultimo admin")
				s := "Es el unico administrador existente"
				return user.NewPutUserConflict().WithPayload(&models.Error{Message: &s})
			}
		}

		ud, _ := dao.GetUserUsername(db, params.Username)
		if bcrypt.CompareHashAndPassword([]byte(*ud.Pwhash), []byte(*params.UserUpdate.Password)) == nil {
			u = userUpdateToUser(params.UserUpdate)
			err = dao.UpdateUser(db, u, params.Username)
			if err != nil {
				log.Println("Error en users_handler PutUsers(): ", err)
				return user.NewPutUserGone()
			}
			return user.NewPutUserOK()
		}
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
		if len(admins) == 1 && *admins[0].Username == params.Username {
			log.Println("Error en users_handler DeleteUser(), intento de borrar ultimo admin")
			s := "Es el unico administrador existente"
			return user.NewDeleteUserBadRequest().WithPayload(&models.Error{Message: &s})
			// BadRequest en vez de Conflict ????
		}
		teams, err := dao.GetTeamsTeamRoleAdmin(db, params.Username)
		if err != nil {
			return user.NewDeleteUserGone()
		}
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
		err = dao.DeleteUser(db, params.Username) // en principio borra cascade
		if err != nil {
			log.Println("Error en users_handler DeleteUser(): ", err)
			return user.NewDeleteUserGone()
		}
		return user.NewPutUserOK()
	}
	return user.NewDeleteUserForbidden()
}

// PutPassword PUT /users/{username}/password Modifies the password of a user.
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
		return user.NewPutPasswordBadRequest() //O forbidden?
	}
	return user.NewPutPasswordForbidden()
}

// PutRole PUT /users/{username}/role Modifies the role of a user.
// Auth: Admin
func PutRole(params user.PutRoleParams, u *models.User) middleware.Responder {
	if isAdmin(u) {
		r := params.Role
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var admins []*dao.User
			admins, err = dao.GetAdmins(db)
			if err == nil {
				if len(admins) == 1 && *admins[0].Username == params.Username {
					log.Println("Error en users_handler PutRole(), intento de quitar admin de ultimo admin")
					s := "Es el unico administrador existente"
					return user.NewPutRoleBadRequest().WithPayload(&models.Error{Message: &s})
				}
				err = dao.PutRole(db, params.Username, r)
				if err == nil {
					return user.NewPutRoleOK()
				}
			}
		}
		log.Println("Error en users_handler PutRole(): ", err)
		return user.NewPutRoleInternalServerError()
	}
	return user.NewPutRoleForbidden()
}

// GetTeamsOfUser GET /users/{username}/teams
// Auth: All
func GetTeamsOfUser(params user.GetTeamsOfUserParams, u *models.User) middleware.Responder {
	var teams []*dao.Team
	db, err := dbconnection.ConnectDb()
	if err == nil {
		teams, err = dao.GetTeamsUsername(db, params.Username)
		if err == nil && teams != nil {
			return user.NewGetTeamsOfUserOK().WithPayload(dao.ToModelsTeams(teams))
		}
	}
	log.Println("Error en users_handler GetTeamsOfUser(): ", err)
	return user.NewGetTeamsOfUserInternalServerError()
}

// GET /users/{username}/teams/{teamname}
// Auth: All
func GetTeamFromUser(params user.GetTeamFromUserParams, u *models.User) middleware.Responder {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		t, err := dao.GetTeamFromUser(db, params.Teamname, params.Username)
		if err == nil {
			if t != nil {
				return user.NewGetTeamFromUserOK().WithPayload(dao.ToModelTeam(t))
			}
			return user.NewGetTeamFromUserGone()
		}
		return user.NewGetTeamFromUserGone()
	}
	return user.NewGetTeamFromUserInternalServerError()
}

// GET /users/{username}/sharedQuestions
// Auth: Current User or Admin
func GetSharedQuestions(params user.GetSharedQuestionsOfUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var q []*dao.Question
			if len(params.Tags) == 0 {
				q, err = dao.GetSharedQuestionsOfUser(db, params.Username)
			} else {
				q, err = dao.GetSharedQuestionsOfUserTags(db, params.Username, params.Tags)
			}
			if err == nil {
				if q != nil {
					mq, err := dao.ToModelQuestions(q)
					if mq != nil && err == nil {
						return user.NewGetSharedQuestionsOfUserOK().WithPayload(mq)
					}
					user.NewGetSharedQuestionsOfUserInternalServerError()
				}
				mq := []*models.Question{}
				return user.NewGetSharedQuestionsOfUserOK().WithPayload(mq)
			}
		}
		return user.NewGetSharedQuestionsOfUserInternalServerError()
	}
	return user.NewGetSharedQuestionsOfUserForbidden()
}

// GET /users/{username}/sharedQuestions/{questionid}
// Auth: Current User or Admin
func GetSharedQuestion(params user.GetSharedQuestionFromUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			q, err := dao.GetSharedQuestionFromUser(db, params.Username, params.Questionid)
			if err == nil {
				if q != nil {
					mq, err := dao.ToModelQuestion(q)
					if mq != nil && err == nil {
						return user.NewGetSharedQuestionFromUserOK().WithPayload(mq)
					}
					user.NewGetSharedQuestionFromUserInternalServerError()
				}
				return user.NewGetSharedQuestionFromUserGone()
			}
		}
		return user.NewGetSharedQuestionFromUserInternalServerError()
	}
	return user.NewGetSharedQuestionFromUserForbidden()
}

// GET /users/{username}/publicEditQuestions
// Auth: Teacher or admin
func GetPublicEditQuestionsOfUser(params user.GetPublicEditQuestionsOfUserParams, u *models.User) middleware.Responder {
	if isTeacherOrAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var q []*dao.Question
			if len(params.Tags) == 0 {
				q, err = dao.GetPublicEditQuestionsOfUser(db, params.Username)
			} else {
				q, err = dao.GetPublicEditQuestionsOfUserTags(db, params.Username, params.Tags)
			}
			if err == nil {
				if q != nil {
					mq, err := dao.ToModelQuestions(q)
					if mq != nil && err == nil {
						return user.NewGetPublicEditQuestionsOfUserOK().WithPayload(mq)
					}
					user.NewGetEditQuestionsOfUserInternalServerError()
				}
				mq := []*models.Question{}
				return user.NewGetPublicEditQuestionsOfUserOK().WithPayload(mq)
			}
		}
		return user.NewGetPublicEditQuestionsOfUserInternalServerError()
	}
	return user.NewGetPublicEditQuestionsOfUserForbidden()
}

// GET /users/{username}/editQuestions
// Auth: Current User or Admin
func GetEditQuestionsOfUser(params user.GetEditQuestionsOfUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var q []*dao.Question
			if len(params.Tags) == 0 {
				q, err = dao.GetEditQuestionsOfUser(db, params.Username)
			} else {
				q, err = dao.GetEditQuestionsOfUserTags(db, params.Username, params.Tags)
			}
			if err == nil {
				if q != nil {
					mq, err := dao.ToModelQuestions(q)
					if mq != nil && err == nil {
						return user.NewGetEditQuestionsOfUserOK().WithPayload(mq)
					}
					user.NewGetEditQuestionsOfUserInternalServerError()
				}
				mq := []*models.Question{}
				return user.NewGetEditQuestionsOfUserOK().WithPayload(mq)
			}
		}
		return user.NewGetEditQuestionsOfUserInternalServerError()
	}
	return user.NewGetEditQuestionsOfUserForbidden()
}

// GET /users/{username}/questions
// Auth: Current User or Admin
func GetQuestionsOfUser(params user.GetQuestionsOfUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var q []*dao.Question
			if len(params.Tags) == 0 {
				q, err = dao.GetQuestionsOfUser(db, params.Username)
			} else {
				q, err = dao.GetQuestionsOfUserTags(db, params.Username, params.Tags)
			}
			if err == nil {
				if q != nil {
					mq, err := dao.ToModelQuestions(q)
					if mq != nil && err == nil {
						return user.NewGetQuestionsOfUserOK().WithPayload(mq)
					}
					user.NewGetQuestionsOfUserInternalServerError()
				}
				mq := []*models.Question{}
				return user.NewGetQuestionsOfUserOK().WithPayload(mq)
			}
		}
		return user.NewGetQuestionsOfUserInternalServerError()
	}
	return user.NewGetQuestionsOfUserForbidden()
}

// POST /users/{username}/questions
// Auth: Current User or Admin
func PostQuestionOfUser(params user.PostQuestionParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var mq *models.Question
			mq, err = dao.PostQuestion(db, params.Question, params.Username)
			if err == nil && mq != nil {
				return user.NewPostQuestionCreated().WithPayload(mq)
			}
			errSt := err.Error()
			return user.NewPostQuestionGone().WithPayload(&models.Error{Message: &errSt})
		}
		return user.NewPostQuestionInternalServerError()
	}
	return user.NewPostQuestionForbidden()
}

// GET /users/{username}/questions/{questionid}
// Auth: Teacher or Admin
func GetQuestionOfUser(params user.GetQuestionFromUserParams, u *models.User) middleware.Responder {
	if isTeacherOrAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			q, err := dao.GetQuestionOfUser(db, params.Username, params.Questionid)
			if err == nil && q != nil {
				mq, err := dao.ToModelQuestion(q)
				if mq != nil && err == nil {
					return user.NewGetQuestionFromUserOK().WithPayload(mq)
				}
			}
			return user.NewGetQuestionFromUserGone()
		}
		return user.NewGetQuestionFromUserInternalServerError()
	}
	return user.NewGetQuestionFromUserForbidden()
}

// GET /users/{username}/sharedTests
// Auth: Current User or Admin
func GetSharedTests(params user.GetSharedTestsFromUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var t []*dao.Test
			t, err = dao.GetSharedTestsFromUser(db, params.Username)
			if err == nil {
				var mt []*models.Test
				mt, err = dao.ToModelTests(t)
				if err == nil {
					return user.NewGetSharedTestsFromUserOK().WithPayload(mt)
				}
			}
		}
		log.Println("Error en GetSharedTests: ", err)
		return user.NewGetSharedTestsFromUserInternalServerError()
	}
	return user.NewGetSharedTestsFromUserForbidden()
}

// GET /users/{username}/sharedTests/{testid}
// Auth: Current User or Admin
func GetSharedTest(params user.GetSharedTestFromUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			t, err := dao.GetSharedTestFromUser(db, params.Username, params.Testid)
			if err == nil {
				if t != nil {
					mt, err := dao.ToModelTest(t)
					if mt != nil && err == nil {
						return user.NewGetSharedTestFromUserOK().WithPayload(mt)
					}
					return user.NewGetSharedTestFromUserInternalServerError()
				}
				return user.NewGetSharedTestFromUserGone()
			}
		}
		return user.NewGetSharedTestFromUserInternalServerError()
	}
	return user.NewGetSharedTestFromUserForbidden()
}

// GET /users/{username}/publicEditTests
// Auth: Teacher or Admin
func GetPublicETestsFromUser(params user.GetPublicEditTestsFromUserParams, u *models.User) middleware.Responder {
	if isTeacherOrAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			t, err := dao.GetPublicEditTestsFromUser(db, params.Username)
			if err == nil {
				mt, err := dao.ToModelTests(t)
				if mt != nil && err == nil {
					return user.NewGetPublicEditTestsFromUserOK().WithPayload(mt)
				}
				return user.NewGetPublicEditTestsFromUserGone()
			}
		}
		return user.NewGetPublicEditTestsFromUserInternalServerError()
	}
	return user.NewGetPublicEditTestsFromUserForbidden()
}

// GET /users/{username}/tests
// Auth: Current User or Admin
func GetTestsFromUser(params user.GetTestsFromUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			t, err := dao.GetTestsFromUser(db, params.Username)
			if err == nil {
				mt, err := dao.ToModelTests(t)
				if mt != nil && err == nil {
					return user.NewGetTestsFromUserOK().WithPayload(mt)
				}
				return user.NewGetTestsFromUserGone()
			}
		}
		return user.NewGetTestsFromUserInternalServerError()
	}
	return user.NewGetTestsFromUserForbidden()
}

// POST /users/{username}/tests
// Auth: Current User or Admin
func PostTest(params user.PostTestParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			t, err := dao.PostTest(db, params.Username, params.Test)
			if err == nil && t != nil {
				return user.NewPostTestCreated().WithPayload(t)
			}
			log.Print("Error en PostTest(): ", err)
			return user.NewPostTestGone()
		}
		return user.NewPostTestInternalServerError()
	}
	return user.NewPostTestForbidden()
}

// GET /users/{username}/tests/{testid}
// Auth: Current User or Admin
func GetTestFromUser(params user.GetTestFromUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			t, err := dao.GetTestFromUser(db, params.Username, params.Testid)
			if err == nil && t != nil {
				mt, err := dao.ToModelTest(t)
				if mt != nil && err == nil {
					return user.NewGetTestFromUserOK().WithPayload(mt)
				}
			}
			return user.NewGetTestFromUserGone()
		}
		return user.NewGetTestFromUserInternalServerError()
	}
	return user.NewGetTestFromUserForbidden()
}

// GET /users/{username}/invitedTests
// Auth: Current User or Admin
func GetInvitedTests(params user.GetInvitedTestsFromUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			t, err := dao.GetInvitedPTestsFromUser(db, params.Username)
			if err == nil {
				mt, err := dao.ToModelTests(t)
				if mt != nil && err == nil {
					return user.NewGetInvitedTestsFromUserOK().WithPayload(mt)
				}
				return user.NewGetInvitedTestsFromUserGone()
			}
		}
		return user.NewGetInvitedTestsFromUserInternalServerError()
	}
	return user.NewGetInvitedTestsFromUserForbidden()
}

// GET /users/{username}/invitedTests/{testid}
// Auth: Current User or Admin
func GetInvitedTest(params user.GetInvitedTestFromUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			t, err := dao.GetInvitedPTestFromUser(db, params.Username, params.Testid)
			if err == nil {
				mt, err := dao.ToModelTest(t)
				if mt != nil && err == nil {
					return user.NewGetInvitedTestFromUserOK().WithPayload(mt)
				}
				return user.NewGetInvitedTestFromUserGone()
			}
		}
		return user.NewGetInvitedTestFromUserInternalServerError()
	}
	return user.NewGetInvitedTestFromUserForbidden()
}

// GET /users/{username}/solvableTests
// Auth: Current User or Admin
func GetPublishedTestsFromUser(params user.GetPublishedTestsFromUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			t, err := dao.GetPublishedTestsFromUser(db, params.Username)
			if err == nil {
				mt, err := dao.ToModelTests(t)
				if mt != nil && err == nil {
					return user.NewGetPublishedTestsFromUserOK().WithPayload(mt)
				}
				return user.NewGetPublishedTestsFromUserGone()
			}
		}
		return user.NewGetPublishedTestsFromUserInternalServerError()
	}
	return user.NewGetPublishedTestsFromUserForbidden()
}

// GET /users/{username}/publicPublishedTests
// Auth: All
func GetPublicPublishedTestsFromUser(params user.GetPublicPublishedTestsFromUserParams, u *models.User) middleware.Responder {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		t, err := dao.GetPublicPublishedTestsFromUser(db, params.Username)
		if err == nil {
			mt, err := dao.ToModelTests(t)
			if mt != nil && err == nil {
				return user.NewGetPublicPublishedTestsFromUserOK().WithPayload(mt)
			}
			return user.NewGetPublicPublishedTestsFromUserGone()
		}
	}
	return user.NewGetPublicPublishedTestsFromUserInternalServerError()
}

// GET /users/{username}/publishedTests
// Auth: Current User or Admin
func GetSolvableTestsFromUser(params user.GetSolvableTestsFromUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			t, err := dao.GetSolvableTestsFromUser(db, params.Username)
			if err == nil {
				mt, err := dao.ToModelTests(t)
				if mt != nil && err == nil {
					return user.NewGetSolvableTestsFromUserOK().WithPayload(mt)
				}
				return user.NewGetSolvableTestsFromUserGone()
			}
		}
		return user.NewGetSolvableTestsFromUserInternalServerError()
	}
	return user.NewGetSolvableTestsFromUserForbidden()
}

// GET /users/{username}/solvableTests/{testid}
// Auth: Current User or Admin
func GetSolvableTestFromUser(params user.GetSolvableTestFromUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			t, err := dao.GetSolvableTestFromUser(db, params.Username, params.Testid)
			if err == nil && t != nil {
				mt, err := dao.ToModelTest(t)
				if mt != nil && err == nil {
					return user.NewGetSolvableTestFromUserOK().WithPayload(mt)
				}
			}
			return user.NewGetSolvableTestFromUserGone()
		}
		return user.NewGetSolvableTestFromUserInternalServerError()
	}
	return user.NewGetSolvableTestFromUserForbidden()
}

// POST /users/{username}/solvableTests/{testid}/answers
// Auth: Current User or Admin
func StartAnswer(params user.StartAnswerParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			t, err := dao.GetSolvableTestFromUser(db, params.Username, params.Testid)
			if t == nil && err == nil && isAdmin(u) {
				t, err = dao.GetPublishedTest(db, params.Testid)
			}
			if err == nil && t != nil {
				var a *dao.Answer
				a, err = dao.StartAnswer(db, params.Username, params.Testid)
				if err == nil && a != nil {
					var ma *models.Answer
					ma, err = dao.ToModelAnswer(a)
					if ma != nil && err == nil {
						return user.NewStartAnswerCreated().WithPayload(ma)
					}
				}
			}
			log.Print("error StartUser: ", err)
			return user.NewStartAnswerGone()
		}
		return user.NewStartAnswerInternalServerError()
	}
	return user.NewStartAnswerForbidden()
}

// GET /users/{username}/answeredTests
// Auth: Current User or Admin
func GetATestsFromUser(params user.GetAnsweredTestsFromUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			t, err := dao.GetATestsFromUser(db, params.Username)
			if err == nil {
				mt, err := dao.ToModelTests(t)
				if mt != nil && err == nil {
					return user.NewGetAnsweredTestsFromUserOK().WithPayload(mt)
				}
				return user.NewGetAnsweredTestsFromUserGone()
			}
		}
		return user.NewGetAnsweredTestsFromUserInternalServerError()
	}
	return user.NewGetAnsweredTestsFromUserForbidden()
}

// GET /users/{username}/answeredTests/{testid}
// Auth: Current User or Admin
func GetATestFromUser(params user.GetAnsweredTestFromUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			t, err := dao.GetATestFromUser(db, params.Username, params.Testid)
			if err == nil && t != nil {
				mt, err := dao.ToModelTest(t)
				if mt != nil && err == nil {
					return user.NewGetAnsweredTestFromUserOK().WithPayload(mt)
				}
			}
			return user.NewGetAnsweredTestFromUserGone()
		}
		return user.NewGetAnsweredTestFromUserInternalServerError()
	}
	return user.NewGetAnsweredTestFromUserForbidden()
}

// GET /users/{username}/answeredTests/{testid}/answers
// Auth: Current User or Admin
func GetAnswersFromUserATest(params user.GetAnswersFromUserAnsweredTestParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			a, err := dao.GetAnswersFromUserAnsweredTest(db, params.Username, params.Testid)
			if err == nil && a != nil {
				ma, err := dao.ToModelAnswers(a)
				if ma != nil && err == nil {
					return user.NewGetAnswersFromUserAnsweredTestOK().WithPayload(ma)
				}
				return user.NewGetAnswersFromUserAnsweredTestGone()
			}
		}
		return user.NewGetAnswersFromUserAnsweredTestInternalServerError()
	}
	return user.NewGetAnswersFromUserAnsweredTestForbidden()
}

// GET /users/{username}/answers
// Auth: Current User or Admin
func GetAnswersFromUser(params user.GetAnswersFromUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			a, err := dao.GetAnswersFromUser(db, params.Username)
			if err == nil && a != nil {
				ma, err := dao.ToModelAnswers(a)
				if ma != nil && err == nil {
					return user.NewGetAnswersFromUserOK().WithPayload(ma)
				}
				return user.NewGetAnswersFromUserGone()
			}
		}
		return user.NewGetAnswersFromUserInternalServerError()
	}
	return user.NewGetAnswersFromUserForbidden()
}

// GET /users/{username}/answers/{answerid}
// Auth: Current User or Admin
func GetAnswerFromUser(params user.GetAnswerFromUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			a, err := dao.GetAnswerFromUser(db, params.Username, params.Answerid)
			if err == nil && a != nil {
				ma, err := dao.ToModelAnswer(a)
				if ma != nil && err == nil {
					return user.NewGetAnswerFromUserOK().WithPayload(ma)
				}
			}
			return user.NewGetAnswerFromUserGone()
		}
		return user.NewGetAnswerFromUserInternalServerError()
	}
	return user.NewGetAnswerFromUserForbidden()
}
