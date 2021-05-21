// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package handlers provides functions that handle http Authuests
package handlers

import (
	"database/sql"
	"errors"
	"log"
	"strings"
	"uva-devtest/models"
	"uva-devtest/permissions"
	"uva-devtest/persistence/dao"
	"uva-devtest/persistence/dbconnection"
	"uva-devtest/restapi/operations/tiporol"

	"github.com/go-openapi/runtime/middleware"
)

// GetTipoRoles GET /tiporoles
// Auth: ALL (none)
func GetTipoRoles(params tiporol.GetTipoRolesParams) middleware.Responder {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var trs []*dao.TipoRol
		trs, err = dao.GetTipoRoles(db)
		if err == nil {
			mtrs := dao.ToModelTipoRoles(trs)
			return tiporol.NewGetTipoRolesOK().WithPayload(mtrs)
		}
	}
	log.Println("Error en GetTipoRoles(): ", err)
	return tiporol.NewGetTipoRolesInternalServerError()
}

// PostTipoRol POST /tiporoles
// Auth: CanAdminPermissions
func PostTipoRol(params tiporol.PostTipoRolParams, u *models.User) middleware.Responder {
	if permissions.CanAdminPermissions(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			var ntr *models.TipoRol
			ntr, err = dao.PostTipoRol(db, params.NewTipoRol)
			if err == nil {
				return tiporol.NewPostTipoRolCreated().WithPayload(ntr)
			}
			log.Println("Error en PostTipoRol(): ", err)
			return tiporol.NewPostTipoRolConflict()
		}
		log.Println("Error en PostTipoRol(): ", err)
		return tiporol.NewPostTipoRolInternalServerError()
	}
	return tiporol.NewPostTipoRolForbidden()
}

// GetTipoRol GET /tiporoles/{rolNombre}
// Auth: ALL (none)
func GetTipoRol(params tiporol.GetTipoRolParams) middleware.Responder {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var tr *dao.TipoRol
		tr, err = dao.GetTipoRolByNombre(db, &params.RolNombre)
		if err == nil {
			mtr := dao.ToModelTipoRol(tr)
			return tiporol.NewGetTipoRolOK().WithPayload(mtr)
		}
	}
	log.Println("Error en GetTipoRol(): ", err)
	return tiporol.NewGetTipoRolInternalServerError()
}

// Comprueba si el rol es el unico rol con el rolbase
func checkIfTipoUnico(db *sql.DB, rolNombre *string) bool {
	if db != nil && rolNombre != nil {
		trs, err := dao.GetTipoRoles(db)
		if err == nil {
			var rb *string
			for _, tr := range trs {
				if strings.EqualFold(*tr.Nombre, *rolNombre) {
					rb = tr.RolBase
					break
				}
			}
			if rb != nil {
				for _, tr := range trs {
					if strings.EqualFold(*tr.RolBase, *rb) {
						return false
					}
				}
			}
		}
	}
	return true
}

// Comprueba si hay cambio de rolbase respecto a la BD
func checkIfChangeRolBase(db *sql.DB, rolNombre *string, ntr *models.TipoRol) bool {
	if db != nil && rolNombre != nil {
		tr, err := dao.GetTipoRolByNombre(db, rolNombre)
		if err == nil && tr != nil {
			if strings.EqualFold(*ntr.RolBase, *tr.RolBase) {
				return false
			}
		}
	}
	return true
}

// PutTipoRol PUT /tiporoles/{rolNombre}
// Auth: CanAdminPermissions
// Req: Si es unico tiporol de su rolBase, no se puede cambiar el rolBase
func PutTipoRol(params tiporol.PutTipoRolParams, u *models.User) middleware.Responder {
	if permissions.CanAdminPermissions(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			if !(checkIfTipoUnico(db, &params.RolNombre) && checkIfChangeRolBase(db, &params.RolNombre, params.NewTipoRol)) {
				err = dao.PutTipoRol(db, &params.RolNombre, params.NewTipoRol)
				if err == nil {
					return tiporol.NewPutTipoRolOK()
				}
			} else {
				err = errors.New("si es el unico rol de su rolBase, no se puede cambiar el rolBase")
			}
			log.Println("Error en PutTipoRol(): ", err)
			return tiporol.NewPutTipoRolConflict()
		}
		log.Println("Error en PutTipoRol(): ", err)
		return tiporol.NewPutTipoRolInternalServerError()
	}
	return tiporol.NewPutTipoRolForbidden()
}

// Obtiene el valor mas alto del atributo prioridad, que es el menos prioritario
func getMaxValorPrioridad(trs []*dao.TipoRol) int64 {
	var ret int64 = -1
	for _, tr := range trs {
		if *tr.Prioridad >= ret {
			ret = *tr.Prioridad
		}
	}
	return ret
}

func getSubstituteTipoRol(tra *dao.TipoRol, trs []*dao.TipoRol) *dao.TipoRol {
	maxValorPrioridad := getMaxValorPrioridad(trs)
	// Busca el que tenga prioridad igual o mayor (menos importante)
	for i := *tra.Prioridad; i < maxValorPrioridad; i++ {
		for _, tr := range trs {
			if *tr.Prioridad == i && *tr.ID != *tra.ID {
				return tr
			}
		}
	}
	for i := *tra.Prioridad - 1; i >= 0; i-- {
		for _, tr := range trs {
			if *tr.Prioridad == i && *tr.ID != *tra.ID {
				return tr
			}
		}
	}
	return nil
}

func substituteUsersTipoRol(db *sql.DB, tra *dao.TipoRol, ntr *dao.TipoRol) error {
	return dao.ChangeTipoRolUsers(db, tra.ID, ntr.ID)
}

// Pasa a todos los usuarios al rol inmediatamente menos importante o igual, si no hay ninguno al inmediatamente mas importante
func adaptUsersTipoRol(db *sql.DB, nombre *string) error {
	if db == nil || nombre == nil {
		return errors.New("parametros nil en adaptUsersTipoRol")
	}
	tra, err := dao.GetTipoRolByNombre(db, nombre)
	if err == nil {
		if tra == nil {
			return errors.New("tiporol no existe, adaptUsersTipoRol()")
		}
		var trs []*dao.TipoRol
		trs, err = dao.GetTipoRoles(db)
		if err == nil {
			ntr := getSubstituteTipoRol(tra, trs)
			if ntr == nil {
				return errors.New("no se pudo encontrar un tiporol substituto para los estudiantes")
			}
			err = substituteUsersTipoRol(db, tra, ntr)
		}
	}
	return err
}

// DeleteTipoRol DELETE /tiporoles/{rolNombre}
// Auth: CanAdminPermissions
// Req: Si es unico tiporol de su rolBase, no se puede borrar
// Al borrar pasar todos los usuarios al rol inmediatamente menos importante o igual, si no hay ninguno al inmediatamente mas importante
func DeleteTipoRol(params tiporol.DeleteTipoRolParams, u *models.User) middleware.Responder {
	if permissions.CanAdminPermissions(u) {
		db, err := dbconnection.ConnectDb()
		if err == nil {
			if !checkIfTipoUnico(db, &params.RolNombre) {
				err = adaptUsersTipoRol(db, &params.RolNombre)
				if err == nil {
					err = dao.DeleteTipoRol(db, &params.RolNombre)
					if err == nil {
						return tiporol.NewDeleteTipoRolOK()
					}
				}
			} else {
				err = errors.New("si es el unico rol de su rolBase, no se puede eliminar")
			}
			log.Println("Error en DeleteTipoRol(): ", err)
			return tiporol.NewDeleteTipoRolBadRequest()
		}
		log.Println("Error en DeleteTipoRol(): ", err)
		return tiporol.NewDeleteTipoRolInternalServerError()
	}
	return tiporol.NewDeleteTipoRolForbidden()
}
