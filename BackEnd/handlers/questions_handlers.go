// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package handlers provides functions that handle http requests
package handlers

import (
	"log"
	"uva-devtest/models"
	"uva-devtest/permissions"
	"uva-devtest/persistence/dao"
	"uva-devtest/persistence/dbconnection"
	"uva-devtest/restapi/operations/question"

	"github.com/go-openapi/runtime/middleware"
)

// GetAllEditQuestions GET /allEditQuestions. Returns all non-published questions
// Auth: CanAdminEQuestions
func GetAllEditQuestions(params question.GetAllEditQuestionsParams, u *models.User) middleware.Responder {
	if permissions.CanAdminEQuestions(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var qs []*dao.Question
			qs, err = dao.GetAllEditQuestions(db, params.Tags, params.LikeTitle, params.Orderby,
				params.Limit, params.Offset)
			if err == nil {
				var mqs []*models.Question
				mqs, err = dao.ToModelQuestions(qs)
				if err == nil {
					return question.NewGetAllEditQuestionsOK().WithPayload(mqs)
				}
			}
		}
		log.Println("Error en questions_handler GetAllEditQuestions(): ", err)
		return question.NewGetAllEditQuestionsInternalServerError()
	}
	return question.NewGetAllEditQuestionsForbidden()
}

// GetAllQuestions GET /allQuestions. Returns all questions
// Auth: CanAdminQuestions
func GetAllQuestions(params question.GetAllQuestionsParams, u *models.User) middleware.Responder {
	if permissions.CanAdminQuestions(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var qs []*dao.Question
			qs, err = dao.GetAllQuestions(db, params.Tags, params.LikeTitle, params.Orderby,
				params.Limit, params.Offset)
			if err == nil {
				var mqs []*models.Question
				mqs, err = dao.ToModelQuestions(qs)
				if err == nil {
					return question.NewGetAllQuestionsOK().WithPayload(mqs)
				}
			}
		}
		log.Println("Error en questions_handler GetAllQuestions(): ", err)
		return question.NewGetAllQuestionsInternalServerError()
	}
	return question.NewGetAllQuestionsForbidden()
}

// GetEditQuestions GET /editQuestions. Returns all public non-published questions.
// Auth: CanVerEQuestions
func GetEditQuestions(params question.GetEditQuestionsParams, u *models.User) middleware.Responder {
	if permissions.CanVerEQuestions(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var qs []*dao.Question
			qs, err = dao.GetEditQuestions(db, params.Tags, params.LikeTitle, params.Orderby,
				params.Limit, params.Offset)
			if err == nil {
				var mqs []*models.Question
				mqs, err = dao.ToModelQuestions(qs)
				if err == nil {
					return question.NewGetEditQuestionsOK().WithPayload(mqs)
				}
			}
		}
		log.Println("Error en questions_handlers GetEditQuestions(): ", err)
		return question.NewGetEditQuestionsInternalServerError()
	}
	return question.NewGetEditQuestionsForbidden()
}

// GetQuestions GET /questions. Returns all public questions.
// Auth: CanVerQuestions
func GetQuestions(params question.GetQuestionsParams, u *models.User) middleware.Responder {
	if permissions.CanVerQuestions(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var qs []*dao.Question
			qs, err = dao.GetQuestions(db, params.Tags, params.LikeTitle, params.Orderby,
				params.Limit, params.Offset)
			if err == nil {
				var mqs []*models.Question
				mqs, err = dao.ToModelQuestions(qs)
				if err == nil {
					return question.NewGetQuestionsOK().WithPayload(mqs)
				}
			}
		}
		log.Println("Error en questions_handlers GetQuestions(): ", err)
		return question.NewGetQuestionsInternalServerError()
	}
	return question.NewGetQuestionsForbidden()
}

// GetQuestion GET /questions/{questionid}. Returns a public question.
// Auth: CanVerQuestions si accesoPublicoNoPublicada=true, CanAdminQuestions o questionAdmin si false
func GetQuestion(params question.GetQuestionParams, u *models.User) middleware.Responder {
	if permissions.CanVerQuestions(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var qs *dao.Question
			qs, err = dao.GetQuestion(db, params.Questionid)
			if err == nil && qs != nil {
				if !*qs.AccesoPublicoNoPublicada {
					if !(permissions.CanAdminQuestions(u) || isQuestionAdmin(u, params.Questionid)) {
						return question.NewGetQuestionForbidden()
					}
				}
				var mqs *models.Question
				mqs, err = dao.ToModelQuestion(qs)
				if err == nil {
					return question.NewGetQuestionOK().WithPayload(mqs)
				}
			}
			return question.NewGetQuestionGone()
		}
		log.Println("Error en questions_handlers GetQuestions(): ", err)
		return question.NewGetQuestionInternalServerError()
	}
	return question.NewGetQuestionForbidden()
}

// Owner o miembro de team que admin question
func isQuestionAdmin(u *models.User, questionid int64) bool {
	if u == nil {
		return false
	}
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var q *dao.Question
		q, err = dao.GetQuestionOfUser(db, *u.Username, questionid)
		if q != nil && err == nil {
			return true
		}
		q, err = dao.GetSharedQuestionFromUser(db, *u.Username, questionid)
		if q != nil && err == nil {
			return true
		}
	}
	return false
}

func isQuestionEditable(questionid int64) bool {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var q *dao.Question
		q, err = dao.GetQuestion(db, questionid)
		if err == nil && q != nil {
			return *q.Editable
		}
	}
	return false
}

func isTeamSoloProfesores(teamname string) bool {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		t, err := dao.GetTeam(db, teamname)
		if err == nil && t != nil {
			return *t.SoloProfesores
		}
	}
	return false
}

// PutQuestion PUT /questions/{questionid}. Modifies an editable question
// Auth: QuestionOwner or CanAdminEQuestions
// Req: Question.Editable
func PutQuestion(params question.PutQuestionParams, u *models.User) middleware.Responder {
	if isQuestionEditable(params.Questionid) && (permissions.CanAdminEQuestions(u) || isQuestionAdmin(u, params.Questionid)) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.PutQuestion(db, params.Questionid, params.Question)
			if err == nil {
				return question.NewPutQuestionOK()
			}
		}
	}
	return question.NewPutQuestionForbidden()
}

// DeleteQuestion DELETE /questions/{questionid}. Deletes an editable question
// Auth: QuestionOwner or CanAdminEQuestions
// Req: Question.Editable
func DeleteQuestion(params question.DeleteQuestionParams, u *models.User) middleware.Responder {
	if isQuestionEditable(params.Questionid) && (permissions.CanAdminEQuestions(u) || isQuestionAdmin(u, params.Questionid)) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.DeleteQuestion(db, params.Questionid)
			if err == nil {
				return question.NewDeleteQuestionOK()
			}
		}
	}
	return question.NewDeleteQuestionForbidden()
}

// GET /questions/{questionid}/tags
// Auth: CanVerEQuestions
func GetQuestionTags(params question.GetTagsFromQuestionParams, u *models.User) middleware.Responder {
	if permissions.CanVerEQuestions(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var ts []*dao.Tag
			ts, err = dao.GetQuestionTags(db, params.Questionid)
			var mts []*models.Tag
			if err == nil {
				mts = dao.ToModelTags(ts)
				if err == nil {
					return question.NewGetTagsFromQuestionOK().WithPayload(mts)
				}
			}
		}
		log.Println("Error en questions_handlers GetQuestionTags(): ", err)
		return question.NewGetTagsFromQuestionInternalServerError()
	}
	return question.NewGetTagsFromQuestionForbidden()
}

// GET /questions/{questionid}/tags/{tag}
// Auth: CanVerEQuestions
func GetQuestionTag(params question.GetTagFromQuestionParams, u *models.User) middleware.Responder {
	if permissions.CanVerEQuestions(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var t *dao.Tag
			t, err = dao.GetQuestionTag(db, params.Questionid, params.Tag)
			var mt *models.Tag
			if err == nil && t != nil {
				mt = dao.ToModelTag(t)
				if err == nil {
					return question.NewGetTagFromQuestionOK().WithPayload(mt)
				}
			}
			return question.NewGetTagFromQuestionGone()
		}
		log.Println("Error en questions_handlers GetQuestionTag(): ", err)
		return question.NewGetTagFromQuestionInternalServerError()
	}
	return question.NewGetTagFromQuestionForbidden()
}

// PUT /questions/{questionid}/tags/{tag}
// Auth: QuestionOwner or CanAdminEQuestions
// Req: Question.Editable
func AddQuestionTag(params question.AddTagToQuestionParams, u *models.User) middleware.Responder {
	if isQuestionEditable(params.Questionid) && (permissions.CanAdminEQuestions(u) || isQuestionAdmin(u, params.Questionid)) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var t *dao.Tag
			t, _ = dao.GetQuestionTag(db, params.Questionid, params.Tag)
			if t != nil {
				return question.NewAddTagToQuestionConflict()
			}
			err = dao.AddQuestionTag(db, params.Questionid, params.Tag)
			if err == nil {
				if err == nil {
					return question.NewAddTagToQuestionOK()
				}
			}
		}
		log.Println("Error en questions_handlers AddQuestionTag(): ", err)
		return question.NewAddTagToQuestionGone()
	}
	return question.NewAddTagToQuestionForbidden()
}

// DELETE /questions/{questionid}/tags/{tag}
// Auth: QuestionOwner or CanAdminEQuestions
// Req: Question.Editable
func RemoveQuestionTag(params question.RemoveTagFromQuestionParams, u *models.User) middleware.Responder {
	if isQuestionEditable(params.Questionid) && (permissions.CanAdminEQuestions(u) || isQuestionAdmin(u, params.Questionid)) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.RemoveQuestionTag(db, params.Questionid, params.Tag)
			if err == nil {
				if err == nil {
					return question.NewRemoveTagFromQuestionOK()
				}
			}
		}
		log.Println("Error en questions_handlers RemoveQuestionTag(): ", err)
		return question.NewRemoveTagFromQuestionGone()
	}
	return question.NewRemoveTagFromQuestionForbidden()
}

// PUT /questions/{questionid}/teams/{teamid}
// Auth: QuestionOwner or CanAdminEQuestions
// Req: Question.Editable
func AddQuestionTeam(params question.AddTeamToQuestionParams, u *models.User) middleware.Responder {
	if isQuestionEditable(params.Questionid) && (permissions.CanAdminEQuestions(u) || isQuestionAdmin(u, params.Questionid)) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			if isTeamSoloProfesores(params.Teamname) {
				err = dao.AddQuestionTeam(db, params.Questionid, params.Teamname)
				if err == nil {
					if err == nil {
						return question.NewAddTeamToQuestionOK()
					}
				}
				log.Println("Error en questions_handlers AddQuestionTeam(): ", err)
			}
			s := "El equipo a añadir tiene que ser de profesores"
			if err != nil {
				s = "El equipo ya esta añadido"
			} else {
				var t *dao.Team
				t, err = dao.GetTeam(db, params.Teamname)
				if err != nil || t == nil {
					s = "El equipo no existe"
				}
			}
			return question.NewAddTeamToQuestionBadRequest().WithPayload(&models.Error{Message: &s})
		}
		return question.NewAddTeamToQuestionInternalServerError()
	}
	return question.NewAddTeamToQuestionForbidden()
}

// DELETE /questions/{questionid}/teams/{teamid}
// Auth: QuestionOwner or CanAdminEQuestions
// Req: Question.Editable
func RemoveQuestionTeam(params question.RemoveTeamToQuestionParams, u *models.User) middleware.Responder {
	if isQuestionEditable(params.Questionid) && (permissions.CanAdminEQuestions(u) || isQuestionAdmin(u, params.Questionid)) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.RemoveQuestionTeam(db, params.Questionid, params.Teamname)
			if err == nil {
				if err == nil {
					return question.NewRemoveTeamToQuestionOK()
				}
			}
			log.Println("Error en questions_handlers RemoveQuestionTeam(): ", err)
			return question.NewRemoveTeamToQuestionGone()
		}
	}
	return question.NewRemoveTeamToQuestionForbidden()
}

// GET /questions/{questionid}/teams
// Auth: CanVerEQuestions
func GetTeamsFromQuestion(params question.GetTeamsFromQuestionParams, u *models.User) middleware.Responder {
	if permissions.CanVerEQuestions(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var ts []*dao.Team
			ts, err = dao.GetTeamsQuestion(db, params.Questionid)
			if err == nil {
				mts := dao.ToModelsTeams(ts)
				if err == nil {
					return question.NewGetTeamsFromQuestionOK().WithPayload(mts)
				}
			}
		}
		log.Println("Error en questions_handlers GetTeamsFromQuestion(): ", err)
		return question.NewGetTeamsFromQuestionInternalServerError()
	}
	return question.NewGetTeamsFromQuestionForbidden()
}

// GET /questions/{questionid}/options
// Auth: CanVerEQuestions
func GetOptions(params question.GetOptionsFromQuestionParams, u *models.User) middleware.Responder {
	if permissions.CanVerEQuestions(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var os []*dao.Option
			os, err = dao.GetOptionsQuestion(db, params.Questionid)
			if err == nil {
				mos := dao.ToModelOptions(os)
				if err == nil {
					return question.NewGetOptionsFromQuestionOK().WithPayload(mos)
				}
			}
		}
		log.Println("Error en questions_handlers GetOptions(): ", err)
		return question.NewGetOptionsFromQuestionInternalServerError()
	}
	return question.NewGetOptionsFromQuestionForbidden()
}

// POST /questions/{questionid}/options
// Auth: QuestionAdmin or CanAdminEQuestions
func PostOption(params question.PostOptionParams, u *models.User) middleware.Responder {
	if isQuestionAdmin(u, params.Questionid) || permissions.CanAdminEQuestions(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var os *dao.Option
			os, err = dao.PostOption(db, params.Questionid, params.Option)
			if err == nil && os != nil {
				mos := dao.ToModelOption(os)
				if err == nil {
					return question.NewPostOptionCreated().WithPayload(mos)
				}
			}
		}
		log.Print("Error en crear opcion: ", err)
		return question.NewPostOptionInternalServerError()
	}
	return question.NewPostOptionForbidden()
}

// GET /questions/{questionid}/options/{optionindex}
// Auth: CanVerEQuestions
func GetOption(params question.GetOptionFromQuestionParams, u *models.User) middleware.Responder {
	if permissions.CanVerEQuestions(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var os *dao.Option
			os, err = dao.GetOptionQuestion(db, params.Questionid, params.Optionindex)
			if err == nil {
				if os == nil {
					return question.NewGetOptionFromQuestionGone()
				}
				mos := dao.ToModelOption(os)
				if err == nil {
					return question.NewGetOptionFromQuestionOK().WithPayload(mos)
				}
			}
		}
		log.Println("Error en questions_handlers GetOption(): ", err)
		return question.NewGetOptionFromQuestionInternalServerError()
	}
	return question.NewGetOptionFromQuestionForbidden()
}

// PUT /questions/{questionid}/options/{optionindex}
// Auth: QuestionAdmin or CanAdminEQuestions
// Req: Question.Editable
func PutOption(params question.PutOptionParams, u *models.User) middleware.Responder {
	if isQuestionEditable(params.Questionid) && (isQuestionAdmin(u, params.Questionid) || permissions.CanAdminEQuestions(u)) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.PutOption(db, params.Questionid, params.Optionindex, params.Option)
			if err == nil {
				return question.NewPutOptionOK()
			}
		}
		log.Println("Error en questions_handlers PutOption(): ", err)
		return question.NewPutOptionInternalServerError()
	}
	return question.NewPutOptionForbidden()
}

// DELETE /questions/{questionid}/options/{optionindex}
// Auth: QuestionAdmin or CanAdminEQuestions
// Req: Question.Editable
func DeleteOption(params question.DeleteOptionParams, u *models.User) middleware.Responder {
	if isQuestionEditable(params.Questionid) && (isQuestionAdmin(u, params.Questionid) || permissions.CanAdminEQuestions(u)) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.DeleteOption(db, params.Questionid, params.Optionindex)
			if err == nil {
				return question.NewDeleteOptionOK()
			}
		}
		return question.NewDeleteOptionInternalServerError()
	}
	return question.NewDeleteOptionForbidden()
}

// GET /questions/{questionid}/pruebas
// Auth: (CanVerQuestions and Question.AccesoPublicoNoPublicada==true) OR (QuestionAdmin or CanAdminEQuestions)
func GetPruebas(params question.GetPruebasFromQuestionParams, u *models.User) middleware.Responder {
	if permissions.CanVerQuestions(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var qs *dao.Question
			qs, err = dao.GetQuestion(db, params.Questionid)
			if err == nil && qs != nil {
				if !*qs.AccesoPublicoNoPublicada {
					if !(permissions.CanAdminQuestions(u) || isQuestionAdmin(u, params.Questionid)) {
						return question.NewGetPruebasFromQuestionForbidden()
					}
				}
				var ps []*dao.Prueba
				ps, err = dao.GetPruebas(db, params.Questionid)
				if err == nil {
					mps := dao.ToModelPruebas(ps)
					return question.NewGetPruebasFromQuestionOK().WithPayload(mps)
				}
				log.Println("Error en questions_handlers GetPruebas(): ", err)
				return question.NewGetPruebasFromQuestionInternalServerError()
			}
			return question.NewGetPruebasFromQuestionGone()
		}
		log.Println("Error en questions_handlers GetPruebas(): ", err)
		return question.NewGetPruebasFromQuestionInternalServerError()
	}
	return question.NewGetPruebasFromQuestionForbidden()
}

// POST /questions/{questionid}/pruebas
// Auth: (CanVerQuestions and Question.AccesoPublicoNoPublicada==true) OR (QuestionAdmin or CanAdminEQuestions)
// Req: Question.Editable
func PostPrueba(params question.PostPruebaParams, u *models.User) middleware.Responder {
	if permissions.CanVerQuestions(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var qs *dao.Question
			qs, err = dao.GetQuestion(db, params.Questionid)
			if err == nil && qs != nil {
				if !*qs.AccesoPublicoNoPublicada {
					if !(permissions.CanAdminQuestions(u) || isQuestionAdmin(u, params.Questionid)) {
						return question.NewPostPruebaForbidden()
					}
				}
				if !*qs.Editable {
					return question.NewPostPruebaForbidden()
				}
				var p *dao.Prueba
				p, err = dao.PostPrueba(db, params.Questionid, params.Prueba)
				if err == nil {
					mp := dao.ToModelPrueba(p)
					return question.NewPostPruebaCreated().WithPayload(mp)
				}
				log.Println("Error en questions_handlers PostPrueba(): ", err)
				return question.NewPostPruebaInternalServerError()
			}
			return question.NewPostPruebaGone()
		}
		log.Println("Error en questions_handlers PostPrueba(): ", err)
		return question.NewPostPruebaInternalServerError()
	}
	return question.NewPostPruebaForbidden()
}

// GET /questions/{questionid}/pruebas/{pruebaid}
// Auth: (CanVerQuestions and Question.AccesoPublicoNoPublicada==true) OR (QuestionAdmin or CanAdminEQuestions)
func GetPrueba(params question.GetPruebaFromQuestionParams, u *models.User) middleware.Responder {
	if permissions.CanVerQuestions(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var qs *dao.Question
			qs, err = dao.GetQuestion(db, params.Questionid)
			if err == nil && qs != nil {
				if !*qs.AccesoPublicoNoPublicada {
					if !(permissions.CanAdminQuestions(u) || isQuestionAdmin(u, params.Questionid)) {
						return question.NewGetPruebaFromQuestionForbidden()
					}
				}
				var p *dao.Prueba
				p, err = dao.GetPrueba(db, params.Questionid, params.Pruebaid)
				if err == nil {
					mp := dao.ToModelPrueba(p)
					return question.NewGetPruebaFromQuestionOK().WithPayload(mp)
				}
				log.Println("Error en questions_handlers GetPrueba(): ", err)
				return question.NewGetPruebaFromQuestionInternalServerError()
			}
			return question.NewGetPruebaFromQuestionGone()
		}
		log.Println("Error en questions_handlers GetPrueba(): ", err)
		return question.NewGetPruebaFromQuestionInternalServerError()
	}
	return question.NewGetPruebaFromQuestionForbidden()
}

// PUT /questions/{questionid}/pruebas/{pruebaid}
// Auth: (CanVerQuestions and Question.AccesoPublicoNoPublicada==true) OR (QuestionAdmin or CanAdminEQuestions)
// Req: Question.Editable
func PutPrueba(params question.PutPruebaParams, u *models.User) middleware.Responder {
	if permissions.CanVerQuestions(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var qs *dao.Question
			qs, err = dao.GetQuestion(db, params.Questionid)
			if err == nil && qs != nil {
				if !*qs.AccesoPublicoNoPublicada {
					if !(permissions.CanAdminQuestions(u) || isQuestionAdmin(u, params.Questionid)) {
						return question.NewPutPruebaForbidden()
					}
				}
				if !*qs.Editable {
					return question.NewPutPruebaForbidden()
				}
				err = dao.PutPrueba(db, params.Questionid, params.Pruebaid, params.Prueba)
				if err == nil {
					return question.NewPutPruebaOK()
				}
				log.Println("Error en questions_handlers PutPrueba(): ", err)
				return question.NewPutPruebaInternalServerError()
			}
			return question.NewPutPruebaGone()
		}
		log.Println("Error en questions_handlers PutPrueba(): ", err)
		return question.NewPutPruebaInternalServerError()
	}
	return question.NewPutPruebaForbidden()
}

// DELETE /questions/{questionid}/pruebas/{pruebaid}
// Auth: (CanVerQuestions and Question.AccesoPublicoNoPublicada==true) OR (QuestionAdmin or CanAdminEQuestions)
// Req: Question.Editable
func DeletePrueba(params question.DeletePruebaParams, u *models.User) middleware.Responder {
	if permissions.CanVerQuestions(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var qs *dao.Question
			qs, err = dao.GetQuestion(db, params.Questionid)
			if err == nil && qs != nil {
				if !*qs.AccesoPublicoNoPublicada {
					if !(permissions.CanAdminQuestions(u) || isQuestionAdmin(u, params.Questionid)) {
						return question.NewDeletePruebaForbidden()
					}
				}
				if !*qs.Editable {
					return question.NewDeletePruebaForbidden()
				}
				err = dao.DeletePrueba(db, params.Questionid, params.Pruebaid)
				if err == nil {
					return question.NewDeletePruebaOK()
				}
				log.Println("Error en questions_handlers DeletePrueba(): ", err)
				return question.NewDeletePruebaInternalServerError()
			}
			return question.NewDeletePruebaGone()
		}
		log.Println("Error en questions_handlers DeletePrueba(): ", err)
		return question.NewDeletePruebaInternalServerError()
	}
	return question.NewDeletePruebaForbidden()

}
