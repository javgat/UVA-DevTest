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
	return username == *u.Username || u.Rol == models.UserRolAdministrador
}

func isAdmin(u *models.User) bool {
	return u.Rol == models.UserRolAdministrador
}

func isTeacher(u *models.User) bool {
	return u.Rol == models.UserRolAdministrador
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
			log.Println("No existe user en users_handler GetUser(): ", err)
			return user.NewGetUserGone() //410
		}
		return user.NewGetUserOK().WithPayload(dao.ToModelUser(us))

	}
	return user.NewGetUserForbidden()
}

func userUpdateToUser(uu *models.UserUpdate) *models.User {
	u := &models.User{
		Email:    uu.Email,
		Fullname: *uu.Fullname,
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
			if u.Rol != models.UserRolAdministrador {
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
			err = dao.PutRole(db, params.Username, r)
			if err == nil {
				return user.NewPutRoleOK()
			}
		}
		log.Println("Error en users_handler PutRole(): ", err)
		return user.NewPutRoleInternalServerError()
	}
	return user.NewPutRoleForbidden()
}

// GetTeamsOfUser GET /users/{username}/teams
// Auth: Current User or Admin
func GetTeamsOfUser(params user.GetTeamsOfUserParams, u *models.User) middleware.Responder {
	var teams []*dao.Team
	if userOrAdmin(params.Username, u) {
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
	return user.NewGetTeamsOfUserForbidden()
}

// GET /users/{username}/teams/{teamname}
func GetTeamFromUser(params user.GetTeamFromUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
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
	return user.NewGetTeamFromUserForbidden()
}

// GET /users/{username}/questions
func GetQuestionsOfUser(params user.GetQuestionsOfUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			q, err := dao.GetQuestionsOfUser(db, params.Username)
			if err == nil {
				if q != nil {
					mq, err := dao.ToModelQuestions(q)
					if mq != nil && err == nil {
						return user.NewGetQuestionsOfUserOK().WithPayload(mq)
					}
					user.NewGetQuestionsOfUserGone()
				}
			}
		}
		return user.NewGetQuestionsOfUserInternalServerError()
	}
	return user.NewGetQuestionsOfUserForbidden()
}

// POST /users/{username}/questions
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

// GET /users/{username}/tests
func GetTestsFromUser(params user.GetTestsFromUserParams, u *models.User) middleware.Responder {
	if isTeacherOrAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			t, err := dao.GetTestsFromUser(db, params.Username)
			if err == nil {
				log.Print(t)
				log.Print(err)
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
func PostTest(params user.PostTestParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			t, err := dao.PostTest(db, params.Username, params.Test)
			if err == nil && t == nil {
				return user.NewPostTestCreated().WithPayload(t)
			}
			return user.NewPostTestGone()
		}
		return user.NewPostTestInternalServerError()
	}
	return user.NewPostTestForbidden()
}

// GET /users/{username}/tests/{testid}
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

// GET /users/{username}/publishedTests
func GetPTestsFromUser(params user.GetPublishedTestsFromUserParams, u *models.User) middleware.Responder {
	if isTeacherOrAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			t, err := dao.GetPTestsFromUser(db, params.Username)
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

// GET /users/{username}/publishedTests/{testid}
func GetPTestFromUser(params user.GetPublishedTestFromUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			t, err := dao.GetPTestFromUser(db, params.Username, params.Testid)
			if err == nil && t != nil {
				mt, err := dao.ToModelTest(t)
				if mt != nil && err == nil {
					return user.NewGetPublishedTestFromUserOK().WithPayload(mt)
				}
			}
			return user.NewGetPublishedTestFromUserGone()
		}
		return user.NewGetPublishedTestFromUserInternalServerError()
	}
	return user.NewGetPublishedTestFromUserForbidden()
}

// POST /users/{username}/publishedTests/{testid}/answers
func StartAnswer(params user.StartAnswerParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			t, err := dao.GetPTestFromUser(db, params.Username, params.Testid)
			if err == nil && t != nil {
				a, err := dao.StartAnswer(db, params.Username, params.Testid)
				if err == nil && a != nil {
					ma, err := dao.ToModelAnswer(a)
					if ma != nil && err == nil {
						return user.NewStartAnswerCreated().WithPayload(ma)
					}
				}
			}
			return user.NewStartAnswerGone()
		}
		return user.NewStartAnswerInternalServerError()
	}
	return user.NewStartAnswerForbidden()
}

// GET /users/{username}/answeredTests
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
