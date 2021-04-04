// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package handlers provides functions that handle http requests
package handlers

import (
	"log"
	"uva-devtest/models"
	"uva-devtest/persistence/dao"
	"uva-devtest/persistence/dbconnection"
	"uva-devtest/restapi/operations/tag"

	"github.com/go-openapi/runtime/middleware"
)

// GetTests GET /tags. Returns all tags.
// Auth: ALL
func GetTags(params tag.GetTagsParams, u *models.User) middleware.Responder {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var ts []*dao.Tag
		ts, err = dao.GetTags(db)
		if err == nil {
			mts := dao.ToModelTags(ts)
			return tag.NewGetTagsOK().WithPayload(mts)
		}
	}
	log.Println("Error en users_handler GetTags(): ", err)
	return tag.NewGetTagsInternalServerError()
}

// GetTests GET /tags/{tag}. Returns a tag.
// Auth: ALL
func GetTag(params tag.GetTagParams, u *models.User) middleware.Responder {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var ts *dao.Tag
		ts, err = dao.GetTag(db, params.Tag)
		if err == nil && ts != nil {
			mts := dao.ToModelTag(ts)
			return tag.NewGetTagOK().WithPayload(mts)
		}
	}
	log.Println("Error en users_handler GetTag(): ", err)
	return tag.NewGetTagInternalServerError()
}

// GetQuestionsFromTag GET /tags/{tag}/questions. Returns all questions related to tag.
// Auth: Teacher Or Admin
func GetQuestionsFromTag(params tag.GetQuestionsFromTagParams, u *models.User) middleware.Responder {
	if isTeacherOrAdmin(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var ts []*dao.Question
			ts, err = dao.GetQuestionsFromTag(db, params.Tag)
			if err == nil {
				var mts []*models.Question
				mts, err = dao.ToModelQuestions(ts)
				if err == nil {
					return tag.NewGetQuestionsFromTagOK().WithPayload(mts)
				}
			}
		}
		log.Println("Error en users_handler GetTags(): ", err)
		return tag.NewGetQuestionsFromTagInternalServerError()
	}
	return tag.NewGetQuestionsFromTagForbidden()
}
