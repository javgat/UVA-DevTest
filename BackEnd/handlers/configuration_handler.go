// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package handlers provides functions that handle http Authuests
package handlers

import (
	"log"
	"strconv"
	"uva-devtest/emailHelper"
	"uva-devtest/models"
	"uva-devtest/permissions"
	"uva-devtest/persistence/dao"
	"uva-devtest/persistence/dbconnection"
	"uva-devtest/restapi/operations/configuration"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// GetEmailConfiguration GET /emailConfiguration
// Auth: CanAdminConfiguration
func GetEmailConfiguration(params configuration.GetEmailConfigurationParams, u *models.User) middleware.Responder {
	if !permissions.CanAdminConfiguration(u) {
		return configuration.NewGetEmailConfigurationForbidden()
	}
	emailInfo, err := emailHelper.GetOwnEmailInfo()
	if err == nil {
		var puerto int64
		puerto, err = strconv.ParseInt(emailInfo.Serverport, 10, 64)
		if err == nil {
			emailConfig := &models.EmailConfiguration{
				From:        (*strfmt.Email)(&emailInfo.From),
				Frontendurl: &emailInfo.FrontEndUrl,
				Password:    &emailInfo.Password,
				Serverhost:  &emailInfo.Serverhost,
				Serverport:  &puerto,
			}
			return configuration.NewGetEmailConfigurationOK().WithPayload(emailConfig)
		}
	}
	log.Println("Error en GetEmailConfiguration(): ", err)
	return configuration.NewGetEmailConfigurationInternalServerError()
}

// PutEmailConfiguration PUT /emailConfiguration
// Auth: CanAdminConfiguration
func PutEmailConfiguration(params configuration.PutEmailConfigurationParams, u *models.User) middleware.Responder {
	if !permissions.CanAdminConfiguration(u) {
		return configuration.NewPutEmailConfigurationForbidden()
	}
	emailConfig := params.EmailConfiguration
	port := strconv.FormatInt(*emailConfig.Serverport, 10)
	emailInfo := &emailHelper.EmailInfo{
		From:        emailConfig.From.String(),
		FrontEndUrl: *emailConfig.Frontendurl,
		Password:    *emailConfig.Password,
		Serverhost:  *emailConfig.Serverhost,
		Serverport:  port,
	}
	err := emailHelper.PutOwnEmailInfo(emailInfo)
	if err == nil {
		return configuration.NewPutEmailConfigurationOK()
	}
	log.Println("Error en PutEmailConfiguration(): ", err)
	return configuration.NewPutEmailConfigurationInternalServerError()
}

func GetCViews(params configuration.GetCViewsParams, u *models.User) middleware.Responder {
	if permissions.CanAdminConfiguration(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var cvs []*dao.CustomizedView
			cvs, err = dao.GetCViews(db)
			if err == nil {
				mcvs := dao.ToModelCustomizedViews(cvs)
				return configuration.NewGetCViewsOK().WithPayload(mcvs)
			}
		}
		log.Println("Error en GetCViews(): ", err)
		return configuration.NewGetCViewsInternalServerError()
	}
	return configuration.NewGetCViewsForbidden()
}

func GetCView(params configuration.GetCViewParams, u *models.User) middleware.Responder {
	if permissions.CanAdminConfiguration(u) || params.RolBase == models.CustomizedViewRolBaseNoRegistrado || (u != nil && params.RolBase == *u.Rol) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var cv *dao.CustomizedView
			cv, err = dao.GetCView(db, params.RolBase)
			if err == nil {
				mcv := dao.ToModelCustomizedView(cv)
				return configuration.NewGetCViewOK().WithPayload(mcv)
			}
		}
		log.Println("Error en GetCView(): ", err)
		return configuration.NewGetCViewInternalServerError()
	}
	return configuration.NewGetCViewForbidden()
}

func PutCView(params configuration.PutCViewParams, u *models.User) middleware.Responder {
	if permissions.CanAdminConfiguration(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			err = dao.PutCView(db, params.RolBase, params.NewView)
			if err == nil {
				return configuration.NewPutCViewOK()
			}
		}
		log.Println("Error en PutCView(): ", err)
		return configuration.NewPutCViewInternalServerError()
	}
	return configuration.NewPutCViewForbidden()
}
