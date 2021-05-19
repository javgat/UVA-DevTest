// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package handlers provides functions that handle http requests
package handlers

import (
	"log"
	"math/rand"
	"strings"
	"time"
	"uva-devtest/emailHelper"
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

// GET /users/{username}/availableEditQuestions
// Auth: Current User or Admin
func GetAvailableEditQuestions(params user.GetAvailableEditQuestionsOfUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var q []*dao.Question
			q, err = dao.GetAvailableEditQuestionsOfUser(db, params.Username, params.Tags, params.LikeTitle, params.Orderby,
				params.Limit, params.Offset)
			if err == nil {
				if q != nil {
					mq, err := dao.ToModelQuestions(q)
					if mq != nil && err == nil {
						return user.NewGetAvailableEditQuestionsOfUserOK().WithPayload(mq)
					}
					user.NewGetAvailableEditQuestionsOfUserInternalServerError()
				}
				mq := []*models.Question{}
				return user.NewGetAvailableEditQuestionsOfUserOK().WithPayload(mq)
			}
		}
		return user.NewGetAvailableEditQuestionsOfUserInternalServerError()
	}
	return user.NewGetAvailableEditQuestionsOfUserForbidden()
}

// GET /users/{username}/availableQuestions
// Auth: Current User or Admin
func GetAvailableQuestions(params user.GetAvailableQuestionsOfUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var q []*dao.Question
			q, err = dao.GetAvailableQuestionsOfUser(db, params.Username, params.Tags, params.LikeTitle, params.Orderby,
				params.Limit, params.Offset)
			if err == nil {
				if q != nil {
					mq, err := dao.ToModelQuestions(q)
					if mq != nil && err == nil {
						return user.NewGetAvailableQuestionsOfUserOK().WithPayload(mq)
					}
					user.NewGetAvailableQuestionsOfUserInternalServerError()
				}
				mq := []*models.Question{}
				return user.NewGetAvailableQuestionsOfUserOK().WithPayload(mq)
			}
		}
		return user.NewGetAvailableQuestionsOfUserInternalServerError()
	}
	return user.NewGetAvailableQuestionsOfUserForbidden()
}

// GET /users/{username}/sharedQuestions
// Auth: Current User or Admin
func GetSharedQuestions(params user.GetSharedQuestionsOfUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var q []*dao.Question
			q, err = dao.GetSharedQuestionsOfUser(db, params.Username, params.Tags, params.LikeTitle, params.Orderby,
				params.Limit, params.Offset)
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
			q, err = dao.GetPublicEditQuestionsOfUser(db, params.Username, params.Tags, params.LikeTitle, params.Orderby,
				params.Limit, params.Offset)
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
			q, err = dao.GetEditQuestionsOfUser(db, params.Username, params.Tags, params.LikeTitle, params.Orderby,
				params.Limit, params.Offset)
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
			q, err = dao.GetQuestionsOfUser(db, params.Username, params.Tags, params.LikeTitle, params.Orderby,
				params.Limit, params.Offset)
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

func copyOption(opt *dao.Option, targetQID int64) (*models.Option, error) {
	opt.Preguntaid = targetQID
	db, err := dbconnection.ConnectDb()
	if err == nil {
		mo := dao.ToModelOption(opt)
		var o *dao.Option
		o, err = dao.PostOption(db, targetQID, mo)
		if err == nil && o != nil {
			mo = dao.ToModelOption(o)
			return mo, nil
		}
	}
	return nil, err
}

func copyOptions(sourceQID int64, targetQID int64) error {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var opts []*dao.Option
		opts, err = dao.GetOptionsQuestion(db, sourceQID)
		if err == nil {
			for _, opt := range opts {
				_, err = copyOption(opt, targetQID)
				if err != nil {
					return err
				}
			}
			return nil
		}
	}
	return err
}

func copyAddTagQuestion(tag *dao.Tag, targetQID int64) error {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		err = dao.AddQuestionTag(db, targetQID, *tag.Tag)
	}
	return err
}

func copyTagsQuestion(sourceQID int64, targetQID int64) error {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var tags []*dao.Tag
		tags, err = dao.GetQuestionTags(db, sourceQID)
		if err == nil {
			for _, tag := range tags {
				err = copyAddTagQuestion(tag, targetQID)
				if err != nil {
					return err
				}
			}
			return nil
		}
	}
	return err
}

func copyQuestion(q *dao.Question, username string, userID int64) (*models.Question, error) {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		q.Usuarioid = userID
		var mq *models.Question
		mq, err = dao.ToModelQuestion(q)
		if err == nil {
			btrue := true
			mq.Editable = &btrue
			mq, err = dao.PostQuestion(db, mq, username)
			if err == nil {
				err = copyOptions(q.ID, mq.ID)
				if err == nil {
					err = copyTagsQuestion(q.ID, mq.ID)
					if err == nil {
						return mq, nil
					}
				}
			}
		}
	}
	return nil, err
}

// POST /users/{username}/questions/{questionid}/copiedQuestions
// Auth: Teacher or Admin if accesoPublicoNoPublicada=true, else QuestionAdmin or Admin. Admin o Current User
func CopyQuestion(params user.CopyQuestionParams, u *models.User) middleware.Responder {
	if isTeacherOrAdmin(u) && userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var q *dao.Question
			q, err = dao.GetQuestion(db, params.Questionid)
			if err == nil {
				if q == nil {
					return user.NewCopyQuestionGone()
				}
				if !*q.AccesoPublicoNoPublicada {
					if !(isQuestionAdmin(u, params.Questionid) || isAdmin(u)) {
						return user.NewCopyQuestionForbidden()
					}
				}
				var us *dao.User
				us, err = dao.GetUserUsername(db, params.Username)
				if err == nil || us != nil {
					var mq *models.Question
					mq, err = copyQuestion(q, params.Username, us.ID)
					if err == nil && mq != nil {
						return user.NewCopyQuestionCreated().WithPayload(mq)
					}
				}
			}
		}
		log.Println("Error en CopyQuestion(): ", err)
		return user.NewCopyQuestionInternalServerError()
	}
	return user.NewCopyQuestionForbidden()
}

// GET /users/{username}/sharedEditTests
// Auth: Current User or Admin
func GetSharedEditTests(params user.GetSharedEditTestsFromUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var t []*dao.Test
			t, err = dao.GetSharedEditTestsFromUser(db, params.Username, params.Tags, params.LikeTitle, params.Orderby,
				params.Limit, params.Offset)
			if err == nil {
				var mt []*models.Test
				mt, err = dao.ToModelTests(t)
				if err == nil {
					return user.NewGetSharedEditTestsFromUserOK().WithPayload(mt)
				}
			}
		}
		log.Println("Error en GetSharedEditTests: ", err)
		return user.NewGetSharedEditTestsFromUserInternalServerError()
	}
	return user.NewGetSharedEditTestsFromUserForbidden()
}

// GET /users/{username}/sharedPublishedTests
// Auth: Current User or Admin
func GetSharedPublishedTests(params user.GetSharedPublishedTestsFromUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var t []*dao.Test
			t, err = dao.GetSharedPublishedTestsFromUser(db, params.Username, params.Tags, params.LikeTitle, params.Orderby,
				params.Limit, params.Offset)
			if err == nil {
				var mt []*models.Test
				mt, err = dao.ToModelTests(t)
				if err == nil {
					return user.NewGetSharedPublishedTestsFromUserOK().WithPayload(mt)
				}
			}
		}
		log.Println("Error en GetSharedPublishedTests: ", err)
		return user.NewGetSharedPublishedTestsFromUserInternalServerError()
	}
	return user.NewGetSharedPublishedTestsFromUserForbidden()
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
			t, err := dao.GetPublicEditTestsFromUser(db, params.Username, params.Tags, params.LikeTitle, params.Orderby,
				params.Limit, params.Offset)
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
func GetEditTestsFromUser(params user.GetEditTestsFromUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			t, err := dao.GetEditTestsFromUser(db, params.Username, params.Tags, params.LikeTitle, params.Orderby,
				params.Limit, params.Offset)
			if err == nil {
				mt, err := dao.ToModelTests(t)
				if mt != nil && err == nil {
					return user.NewGetEditTestsFromUserOK().WithPayload(mt)
				}
				return user.NewGetEditTestsFromUserGone()
			}
		}
		return user.NewGetEditTestsFromUserInternalServerError()
	}
	return user.NewGetEditTestsFromUserForbidden()
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
			var menosuno int64 = -1
			params.Test.OriginalTestID = &menosuno
			horaCreacion := time.Now()
			params.Test.NotaMaxima = 0
			t, err := dao.PostTest(db, params.Username, params.Test, horaCreacion)
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

func copyQuestions(sourceTID int64, targetTID int64, username string, userid int64) error {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var qs []*dao.Question
		qs, err = dao.GetQuestionsFromTest(db, sourceTID)
		if err == nil {
			for _, q := range qs {
				var mq *models.Question
				mq, err = copyQuestion(q, username, userid)
				if err != nil {
					return err
				}
				err = dao.AddQuestionTest(db, mq.ID, targetTID, *mq.ValorFinal)
				if err != nil {
					return err
				}
			}
			return nil
		}
	}
	return err
}

func copyAddTagTest(tag *dao.Tag, targetTID int64) error {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		err = dao.AddTestTag(db, targetTID, *tag.Tag)
	}
	return err
}

func copyTagsTest(sourceTID int64, targetTID int64) error {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var tags []*dao.Tag
		tags, err = dao.GetTestTags(db, sourceTID)
		if err == nil {
			for _, tag := range tags {
				err = copyAddTagTest(tag, targetTID)
				if err != nil {
					return err
				}
			}
			return nil
		}
	}
	return err
}

func copyTest(t *dao.Test, username string) (*models.Test, error) {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var u *dao.User
		u, err = dao.GetUserUsername(db, username)
		if err == nil || u != nil {
			t.Usuarioid = u.ID
			var mt *models.Test
			mt, err = dao.ToModelTest(t)
			if err == nil {
				btrue := true
				mt.Editable = &btrue
				var menosuno int64 = -1
				mt.OriginalTestID = &menosuno
				horaCreacion := time.Now()
				mt, err = dao.PostTest(db, username, mt, horaCreacion)
				if err == nil {
					err = copyQuestions(t.ID, mt.ID, username, u.ID)
					if err == nil {
						err = copyTagsTest(t.ID, mt.ID)
						if err == nil {
							return mt, nil
						}
					}
				}
			}
		}
	}
	return nil, err
}

// POST /users/{username}/tests/{testid}/copiedTests
// Auth: Teacher or Admin if accesoPublicoNoPublicada=true, else TestAdmin or Admin. Admin o Current User
func CopyTest(params user.CopyTestParams, u *models.User) middleware.Responder {
	if isTeacherOrAdmin(u) && userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var t *dao.Test
			t, err = dao.GetTest(db, params.Testid)
			if err == nil {
				if t == nil {
					return user.NewCopyTestGone()
				}
				if !*t.AccesoPublicoNoPublicado {
					if !(isAdmin(u) || isTestAdmin(u, params.Testid)) {
						return user.NewCopyTestForbidden()
					}
				}
				var mt *models.Test
				mt, err = copyTest(t, params.Username)
				if err == nil && mt != nil {
					return user.NewCopyTestCreated().WithPayload(mt)
				}
			}
		}
		return user.NewCopyTestInternalServerError()
	}
	return user.NewCopyTestForbidden()
}

// GET /users/{username}/invitedTestsByTeamsAndUser
// Auth: Current User or Admin
func GetInvitedTestsByTeamsAndUser(params user.GetInvitedTestsByTeamsAndUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			t, err := dao.GetInvitedPTestsByTeamsAndUser(db, params.Username, params.Tags, params.LikeTitle, params.Orderby,
				params.Limit, params.Offset)
			if err == nil {
				mt, err := dao.ToModelTests(t)
				if mt != nil && err == nil {
					return user.NewGetInvitedTestsByTeamsAndUserOK().WithPayload(mt)
				}
				return user.NewGetInvitedTestsByTeamsAndUserGone()
			}
		}
		return user.NewGetInvitedTestsByTeamsAndUserInternalServerError()
	}
	return user.NewGetInvitedTestsByTeamsAndUserForbidden()
}

// GET /users/{username}/invitedTests
// Auth: Current User or Admin
func GetInvitedTests(params user.GetInvitedTestsFromUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			t, err := dao.GetInvitedPTestsFromUser(db, params.Username, params.Tags, params.LikeTitle, params.Orderby,
				params.Limit, params.Offset)
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

// GET /users/{username}/publishedTests
// Auth: Current User or Admin
func GetPublishedTestsFromUser(params user.GetPublishedTestsFromUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			t, err := dao.GetPublishedTestsFromUser(db, params.Username, params.Tags, params.LikeTitle, params.Orderby,
				params.Limit, params.Offset)
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
		t, err := dao.GetPublicPublishedTestsFromUser(db, params.Username, params.Tags, params.LikeTitle, params.Orderby,
			params.Limit, params.Offset)
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

// GET /users/{username}/solvableTests
// Auth: Current User or Admin
func GetSolvableTestsFromUser(params user.GetSolvableTestsFromUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var t []*dao.Test
			t, err = dao.GetSolvableTestsFromUser(db, params.Username, params.Tags, params.LikeTitle, params.Orderby,
				params.Limit, params.Offset)
			if err == nil {
				mt, err := dao.ToModelTests(t)
				if mt != nil && err == nil {
					return user.NewGetSolvableTestsFromUserOK().WithPayload(mt)
				}
				return user.NewGetSolvableTestsFromUserGone()
			}
		}
		log.Println("Error en GetSolvableTestsFromUser(): ", err)
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
					var cant int64
					cant, err = dao.GetNumberAnswersUserTest(db, params.Username, params.Testid)
					if err == nil {
						mt.CantidadRespuestasDelUsuario = cant
						return user.NewGetSolvableTestFromUserOK().WithPayload(mt)
					}
				}
			}
			return user.NewGetSolvableTestFromUserGone()
		}
		log.Println("Error en GetSolvableTestFromUser(): ", err)
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
				visible := strings.EqualFold(*t.Visibilidad, models.TestVisibilidadAlEntregar)
				var a *dao.Answer
				a, err = dao.StartAnswer(db, params.Username, params.Testid, visible)
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

// GET /users/{username}/pendingTests
// Auth: Current User or Admin
func GetPendingTestsFromUser(params user.GetPendingTestsFromUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var t []*dao.Test
			t, err = dao.GetPendingTestsFromUser(db, params.Username, params.Tags, params.LikeTitle, params.Orderby,
				params.Limit, params.Offset)
			if err == nil {
				mt, err := dao.ToModelTests(t)
				if mt != nil && err == nil {
					return user.NewGetPendingTestsFromUserOK().WithPayload(mt)
				}
				return user.NewGetPendingTestsFromUserGone()
			}
		}
		log.Println("Error en GetPendingTestsUser(): ", err)
		return user.NewGetPendingTestsFromUserInternalServerError()
	}
	return user.NewGetPendingTestsFromUserForbidden()
}

// GET /users/{username}/answeredTests
// Auth: Current User or Admin
func GetATestsFromUser(params user.GetAnsweredTestsFromUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var t []*dao.Test
			t, err = dao.GetATestsFromUser(db, params.Username, params.Tags, params.LikeTitle, params.Orderby,
				params.Limit, params.Offset)
			if err == nil {
				mt, err := dao.ToModelTests(t)
				if mt != nil && err == nil {
					return user.NewGetAnsweredTestsFromUserOK().WithPayload(mt)
				}
				return user.NewGetAnsweredTestsFromUserGone()
			}
		}
		log.Println("Error en GetATestsFromUser(): ", err)
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

// GET /users/{username}/solvableTests/{testid}/openAnswers
// Auth: Current User or Admin
func GetOpenAnswersTestUser(params user.GetOpenAnswersFromUserTestParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var a []*dao.Answer
			a, err = dao.GetOpenAnswersFromUserTest(db, params.Username, params.Testid)
			if err == nil {
				var ma []*models.Answer
				ma, err = dao.ToModelAnswers(a)
				if err == nil {
					return user.NewGetOpenAnswersFromUserTestOK().WithPayload(ma)
				}
				return user.NewGetOpenAnswersFromUserTestGone()
			}
		}
		log.Println("Error en GetOpenAnswersTestUser(): ", err)
		return user.NewGetOpenAnswersFromUserTestInternalServerError()
	}
	return user.NewGetOpenAnswersFromUserTestForbidden()
}

// GET /users/{username}/answeredTests/{testid}/answers
// Auth: Current User or Admin
func GetAnswersFromUserATest(params user.GetAnswersFromUserAnsweredTestParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			a, err := dao.GetAnswersFromUserAnsweredTest(db, params.Username, params.Testid)
			if err == nil {
				ma, err := dao.ToModelAnswers(a)
				if err == nil {
					return user.NewGetAnswersFromUserAnsweredTestOK().WithPayload(ma)
				}
				return user.NewGetAnswersFromUserAnsweredTestGone()
			}
		}
		return user.NewGetAnswersFromUserAnsweredTestInternalServerError()
	}
	return user.NewGetAnswersFromUserAnsweredTestForbidden()
}

// GET /users/{username}/answeredTests/{testid}/correctedAnswers
// Auth: Current User or Admin
func GetCAnswersFromUserATest(params user.GetCorrectedAnswersFromUserAnsweredTestParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			a, err := dao.GetCorrectedAnswersFromUserAnsweredTest(db, params.Username, params.Testid)
			if err == nil {
				ma, err := dao.ToModelAnswers(a)
				if err == nil {
					return user.NewGetCorrectedAnswersFromUserAnsweredTestOK().WithPayload(ma)
				}
				return user.NewGetCorrectedAnswersFromUserAnsweredTestGone()
			}
		}
		return user.NewGetCorrectedAnswersFromUserAnsweredTestInternalServerError()
	}
	return user.NewGetCorrectedAnswersFromUserAnsweredTestForbidden()
}

// GET /users/{username}/answeredTests/{testid}/uncorrectedAnswers
// Auth: Current User or Admin
func GetUCAnswersFromUserATest(params user.GetUncorrectedAnswersFromUserAnsweredTestParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			a, err := dao.GetUncorrectedAnswersFromUserAnsweredTest(db, params.Username, params.Testid)
			if err == nil {
				ma, err := dao.ToModelAnswers(a)
				if err == nil {
					return user.NewGetUncorrectedAnswersFromUserAnsweredTestOK().WithPayload(ma)
				}
				return user.NewGetUncorrectedAnswersFromUserAnsweredTestGone()
			}
		}
		return user.NewGetUncorrectedAnswersFromUserAnsweredTestInternalServerError()
	}
	return user.NewGetUncorrectedAnswersFromUserAnsweredTestForbidden()
}

// GET /users/{username}/correctedAnswers
// Auth: Current User or Admin
func GetCorrectedAnswersFromUser(params user.GetCorrectedAnswersFromUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var a []*dao.Answer
			a, err = dao.GetCorrectedAnswersFromUser(db, params.Username)
			if err == nil {
				ma, err := dao.ToModelAnswers(a)
				if ma != nil && err == nil {
					return user.NewGetCorrectedAnswersFromUserOK().WithPayload(ma)
				}
				return user.NewGetCorrectedAnswersFromUserGone()
			}
		}
		log.Println("Error en GetCorrectedAnswersFromUser(): ", err)
		return user.NewGetCorrectedAnswersFromUserInternalServerError()
	}
	return user.NewGetCorrectedAnswersFromUserForbidden()
}

// GET /users/{username}/uncorrectedAnswers
// Auth: Current User or Admin
func GetUncorrectedAnswersFromUser(params user.GetUncorrectedAnswersFromUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var a []*dao.Answer
			a, err = dao.GetUncorrectedAnswersFromUser(db, params.Username)
			if err == nil {
				ma, err := dao.ToModelAnswers(a)
				if ma != nil && err == nil {
					return user.NewGetUncorrectedAnswersFromUserOK().WithPayload(ma)
				}
				return user.NewGetUncorrectedAnswersFromUserGone()
			}
		}
		log.Println("Error en GetUncorrectedAnswersFromUser(): ", err)
		return user.NewGetUncorrectedAnswersFromUserInternalServerError()
	}
	return user.NewGetAnswersFromUserForbidden()
}

// GET /users/{username}/answers
// Auth: Current User or Admin
func GetAnswersFromUser(params user.GetAnswersFromUserParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var a []*dao.Answer
			a, err = dao.GetAnswersFromUser(db, params.Username)
			if err == nil {
				ma, err := dao.ToModelAnswers(a)
				if ma != nil && err == nil {
					return user.NewGetAnswersFromUserOK().WithPayload(ma)
				}
				return user.NewGetAnswersFromUserGone()
			}
		}
		log.Println("Error en GetAnswersFromUser(): ", err)
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

func RecoverPassword(params user.RecoverPasswordParams) middleware.Responder {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var tok *dao.MailToken
		tok, err = dao.GetMailToken(db, params.PasswordRecovery.Mailtoken)
		if err == nil {
			if tok == nil {
				log.Println("Error al recuperar contraseña, el token no existe")
				return user.NewRecoverPasswordForbidden()
			}
			err = dao.DeleteMailToken(db, tok.Mailtoken) // Token gets consumed
			if err == nil {
				var u *dao.User
				u, err = dao.GetUserUsername(db, params.Username)
				if err == nil && u != nil {
					//If caducidad is before (antes de) now OR not his token
					if tok.Caducidad.Before(time.Now()) {
						log.Println("Error al recuperar contraseña, el token ha caducado")
						return user.NewRecoverPasswordForbidden()
					}
					if tok.Userid != u.ID {

						log.Println("Error al recuperar contraseña, el token no es suyo")
						return user.NewRecoverPasswordForbidden()
					}
					// Correct Token
					bytes, errBcrypt := bcrypt.GenerateFromPassword([]byte(*params.PasswordRecovery.Newpass), Cost)
					newpwhash := string(bytes)
					err = dao.PutPasswordUsername(db, params.Username, newpwhash)
					if err != nil || errBcrypt != nil {
						log.Println("Error al modificar la contraseña: ", err, errBcrypt)
						return user.NewRecoverPasswordInternalServerError()
					}
					return user.NewRecoverPasswordOK()
				}
			}
		}
	}
	return user.NewRecoverPasswordInternalServerError()
}

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func NewUniqueMailToken() (string, error) {
	db, err := dbconnection.ConnectDb()
	var tk *dao.MailToken
	var candidate string
	if err == nil {
		for {
			candidate = RandomString(40)
			tk, err = dao.GetMailToken(db, &candidate)
			if err != nil {
				return "", err
			}
			if tk == nil {
				return candidate, nil
			}
		}
	}
	return "", err
}

func PostRecoveryToken(params user.PostRecoveryTokenParams) middleware.Responder {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var token string
		token, err = NewUniqueMailToken()
		if err == nil {
			var u *dao.User
			u, err = dao.GetUserUsername(db, params.Username)
			if u == nil || err != nil {
				u, err = dao.GetUserEmail(db, params.Username)
				if u == nil || err != nil {
					return user.NewPostRecoveryTokenGone()
				}
			}
			err = dao.PostRecoveryToken(db, *u.Username, token)
			if err == nil {
				emailHelper.SendPasswordRecoveryMail(*u.Username, token)
				return user.NewPostRecoveryTokenCreated()
			}
		}
	}
	log.Println("error en PostRecoveryToken()", err)
	return user.NewPostRecoveryTokenInternalServerError()
}

func PostEmailUser(params user.PostEmailUserParams, u *models.User) middleware.Responder {
	if isTeacherOrAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			eu := params.EmailUser
			username := eu.Email.String()
			studentRol := models.UserRolEstudiante
			pass := RandomString(40)
			var bytes []byte
			var du *dao.User
			du, err = dao.GetUserEmail(db, eu.Email.String())
			if err == nil {
				if du != nil {
					return user.NewPostEmailUserConflict()
				}
				bytes, err = bcrypt.GenerateFromPassword([]byte(pass), Cost)
				if err == nil {
					pwhashstring := string(bytes)
					du := &dao.User{
						Username: &username,
						Email:    eu.Email,
						Pwhash:   &pwhashstring,
						Rol:      &studentRol,
						Fullname: &username,
					}
					err = dao.InsertUser(db, du)
					if err == nil {
						mu := dao.ToModelUser(du)
						// ENVIAR MAIL
						emailHelper.SendEmailUserCreated(username, pass)
						return user.NewPostEmailUserCreated().WithPayload(mu)
					}
				}
			}
		}
		log.Println("Error en PostEmailUser(): ", err)
		return user.NewPostEmailUserInternalServerError()
	}
	return user.NewPostEmailUserForbidden()
}
