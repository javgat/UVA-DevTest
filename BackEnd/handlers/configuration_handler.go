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
