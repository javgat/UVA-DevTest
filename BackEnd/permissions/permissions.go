// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package handlers provides functions that handle http requests
package permissions

import (
	"uva-devtest/models"
	"uva-devtest/persistence/dao"
	"uva-devtest/persistence/dbconnection"
)

func isAdmin(u *models.User) bool {
	if u == nil {
		return false
	}
	return *u.Rol == models.UserRolAdministrador
}

func isTeacher(u *models.User) bool {
	if u == nil {
		return false
	}
	return *u.Rol == models.UserRolProfesor
}

func isTeacherOrAdmin(u *models.User) bool {
	return isAdmin(u) || isTeacher(u)
}

func getUserTipoRol(u *models.User) *dao.TipoRol {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var tr *dao.TipoRol
		if u == nil {
			tr, err = dao.GetTipoRolNoRegistrado(db)
			if err == nil && tr != nil {
				return tr
			}
			return dao.GetDefaultTipoRolNoRegistrado()
		}
		tr, err = dao.GetTipoRolByNombre(db, &u.Tiporol)
		if err == nil && tr != nil {
			return tr
		}
	}
	return dao.GetDefaultTipoRolStudent()
}

// CanVerPTests indica si el usuario puede ver tests publicados públicos
func CanVerPTests(u *models.User) bool {
	return *getUserTipoRol(u).VerPTests
}

// CanVerETests indica si el usuario puede ver tests editables públicos
func CanVerETests(u *models.User) bool {
	return *getUserTipoRol(u).VerETests
}

// CanVerTests
func CanVerTests(u *models.User) bool {
	return CanVerPTests(u) && CanVerETests(u)
}

// CanVerEQuestions indica si el usuario puede ver preguntas editables públicos
func CanVerEQuestions(u *models.User) bool {
	return *getUserTipoRol(u).VerEQuestions
}

// CanVerPQuestions indica si el usuario puede ver preguntas publicadas públicos
func CanVerPQuestions(u *models.User) bool {
	return *getUserTipoRol(u).VerPQuestions
}

func CanVerQuestions(u *models.User) bool {
	return CanVerPQuestions(u) && CanVerEQuestions(u)
}

// CanVerPQuestions indica si el usuario puede ver respuestas de otros usuarios
func CanVerAnswers(u *models.User) bool {
	return *getUserTipoRol(u).VerAnswers
}

// CanChangeRoles indica si el usuario puede modificar el rol de otro, de oldpriority a newpriority
// Usa parametro change roles, y que el old rol sea de prioridad menor que el del usuario, y el new como maximo igual (la prioridad tiene orden inverso, mayor es 0)
func CanChangeRoles(u *models.User, otherRolPriorityOld int64, otherRolPriorityNew int64) bool {
	tipo := *getUserTipoRol(u)
	return *tipo.ChangeRoles && (otherRolPriorityOld >= *tipo.Prioridad) && (otherRolPriorityNew >= *tipo.Prioridad)
}

// CanTenerTeams indica si el usuario puede crear o ser administrador de un equipo
func CanTenerTeams(u *models.User) bool {
	return *getUserTipoRol(u).TenerTeams
}

func CanTenerEQuestions(u *models.User) bool {
	return *getUserTipoRol(u).TenerEQuestions
}

func CanTenerETests(u *models.User) bool {
	return *getUserTipoRol(u).TenerETests
}

func CanTenerPTests(u *models.User) bool {
	return *getUserTipoRol(u).TenerPTests
}

// ADMIN

// CanAdminPTests indica si el usuario puede ver y editar tests publicados públicos y privados (y sus preguntas)
func CanAdminPTests(u *models.User) bool {
	return *getUserTipoRol(u).AdminPTests
}

// CanAdminETests indica si el usuario puede ver y editar tests editables públicos y privados
func CanAdminETests(u *models.User) bool {
	return *getUserTipoRol(u).AdminETests
}

func CanAdminTests(u *models.User) bool {
	return CanAdminPTests(u) && CanAdminETests(u)
}

// CanAdminQuestions indica si el usuario puede ver y editar preguntas editables públicos y privados
func CanAdminEQuestions(u *models.User) bool {
	return *getUserTipoRol(u).AdminEQuestions
}

func CanAdminQuestions(u *models.User) bool {
	return CanAdminEQuestions(u) && CanAdminPTests(u)
}

// CanAdminAnswers indica si el usuario puede ver y corregir respuestas a cualquier test
func CanAdminAnswers(u *models.User) bool {
	return *getUserTipoRol(u).AdminAnswers
}

// CanAdminUsers indica si el usuario puede modificar datos del usuario
func CanAdminUsers(u *models.User) bool {
	return *getUserTipoRol(u).AdminUsers
}

// CanAdminUsers indica si el usuario puede modificar datos de equipos
func CanAdminTeams(u *models.User) bool {
	return *getUserTipoRol(u).AdminUsers
}

// CanAdminConfiguration indica si el usuario puede modificar otras configuraciones del servidor
func CanAdminConfiguration(u *models.User) bool {
	return *getUserTipoRol(u).AdminConfiguration
}

// CanAdminPermissions indica si el usuario puede modificar los roles existentes y sus privilegios
func CanAdminPermissions(u *models.User) bool {
	return *getUserTipoRol(u).AdminPermissions
}
