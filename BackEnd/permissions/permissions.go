// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package handlers provides functions that handle http requests
package permissions

import "uva-devtest/models"

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

// CanVerPTests indica si el usuario puede ver tests publicados públicos
func CanVerPTests(u *models.User) bool {
	return u != nil
}

// CanVerETests indica si el usuario puede ver tests editables públicos
func CanVerETests(u *models.User) bool {
	return isTeacherOrAdmin(u)
}

// CanVerTests
func CanVerTests(u *models.User) bool {
	return CanVerPTests(u) && CanVerETests(u)
}

// CanVerEQuestions indica si el usuario puede ver preguntas editables públicos
func CanVerEQuestions(u *models.User) bool {
	return isTeacherOrAdmin(u)
}

// CanVerPQuestions indica si el usuario puede ver preguntas publicadas públicos
func CanVerPQuestions(u *models.User) bool {
	return isTeacherOrAdmin(u)
}

func CanVerQuestions(u *models.User) bool {
	return CanVerPQuestions(u) && CanVerEQuestions(u)
}

// CanVerPQuestions indica si el usuario puede ver respuestas de otros usuarios
func CanVerAnswers(u *models.User) bool {
	return isTeacherOrAdmin(u)
}

// CanChangeRoles indica si el usuario puede modificar el rol de otro, de oldpriority a newpriority
func CanChangeRoles(u *models.User, otherRolPriorityOld int64, otherRolPriorityNew int64) bool {
	return isAdmin(u)
}

// CanTenerTeams indica si el usuario puede crear o ser administrador de un equipo
func CanTenerTeams(u *models.User) bool {
	return isTeacherOrAdmin(u)
}

func CanTenerEQuestions(u *models.User) bool {
	return isTeacherOrAdmin(u)
}

func CanTenerETests(u *models.User) bool {
	return isTeacherOrAdmin(u)
}

func CanTenerPTests(u *models.User) bool {
	return isTeacherOrAdmin(u)
}

// ADMIN

// CanAdminPTests indica si el usuario puede ver y editar tests publicados públicos y privados (y sus preguntas)
func CanAdminPTests(u *models.User) bool {
	return isAdmin(u)
}

// CanAdminETests indica si el usuario puede ver y editar tests editables públicos y privados
func CanAdminETests(u *models.User) bool {
	return isAdmin(u)
}

func CanAdminTests(u *models.User) bool {
	return CanAdminPTests(u) && CanAdminETests(u)
}

// CanAdminQuestions indica si el usuario puede ver y editar preguntas editables públicos y privados
func CanAdminEQuestions(u *models.User) bool {
	return isAdmin(u)
}

func CanAdminQuestions(u *models.User) bool {
	return CanAdminEQuestions(u) && CanAdminPTests(u)
}

// CanAdminAnswers indica si el usuario puede ver y corregir respuestas a cualquier test
func CanAdminAnswers(u *models.User) bool {
	return isAdmin(u)
}

// CanAdminUsers indica si el usuario puede modificar datos del usuario
func CanAdminUsers(u *models.User) bool {
	return isAdmin(u)
}

// CanAdminUsers indica si el usuario puede modificar datos de equipos
func CanAdminTeams(u *models.User) bool {
	return isAdmin(u)
}

// CanAdminConfiguration indica si el usuario puede modificar otras configuraciones del servidor
func CanAdminConfiguration(u *models.User) bool {
	return isAdmin(u)
}

// CanAdminPermissions indica si el usuario puede modificar los roles existentes y sus privilegios
func CanAdminPermissions(u *models.User) bool {
	return isAdmin(u)
}
