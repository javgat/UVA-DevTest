package dao

import "github.com/go-openapi/strfmt"

const (

	// TeamRoleRoleAdmin captures enum value "admin"
	TeamRoleRoleAdmin string = "admin"

	// TeamRoleRoleMember captures enum value "miembro"
	TeamRoleRoleMember string = "miembro"
)

// User user
//
// userdao User
type User struct {
	// id
	// Identificador SQL
	ID int64

	// email
	// Example: carlos@mail.com
	// Required: true
	// Format: email
	Email *strfmt.Email `json:"email"`

	// pwhash
	// Example: e$ia9s7ATDGba39pakscAKs
	// Required: true
	Pwhash *string `json:"pwhash"`

	// username
	// Example: carlosg72
	// Required: true
	// Pattern: [^@]+
	Username *string `json:"username"`

	// fullname
	// Example: Javier Gatón Herguedas
	// Required: true
	Fullname *string `json:"fullname"`

	// rol
	// Required: true
	// Enum: [estudiante profesor administrador]
	Rol *string `json:"rol"`
}

// Team team
//
// teamdao team
type Team struct {
	// id
	// Identificador SQL
	ID int64

	// description
	// Example: DevTest Team
	Description string `json:"description,omitempty"`

	// solo profesores
	// Example: true
	// Required: true
	SoloProfesores *bool `json:"soloProfesores"`

	// teamname
	// Example: devtestTeam
	// Required: true
	// Pattern: ^[^@ \t\r\n]+$
	Teamname *string `json:"teamname"`
}

// Question question
//
// swagger:model Question
type Question struct {

	// acceso publico no publicada
	// Example: false
	// Required: true
	AccesoPublicoNoPublicada *bool `json:"accesoPublicoNoPublicada"`

	// auto correct
	// Example: true
	// Required: true
	AutoCorrect *bool `json:"autoCorrect"`

	// editable
	// Example: false
	// Required: true
	Editable *bool `json:"editable"`

	// eleccion unica
	// Example: false
	EleccionUnica bool `json:"eleccionUnica,omitempty"`

	// estimated time
	// Example: 32600
	// Required: true
	// Minimum: 0
	EstimatedTime *int64 `json:"estimatedTime"`

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// question
	// Example: ¿Cual es el lenguaje que tiene un nombre más largo de todos?
	// Required: true
	Question *string `json:"question"`

	// solucion
	// Example: Javadoc
	Solucion string `json:"solucion,omitempty"`

	// tipo pregunta
	// Required: true
	// Enum: [opciones string codigo]
	TipoPregunta *string `json:"tipoPregunta"`

	// title
	// Example: Paralelismo en C
	// Required: true
	Title *string `json:"title"`

	// usuarioid
	// Example: 3
	// Required: true
	Usuarioid int64 `json:"usuarioid"`

	// valor final
	// Required: true
	// Minimum: 0
	ValorFinal *int64 `json:"valorFinal"`
}

// Test test
//
// swagger:model Test
type Test struct {

	// acceso publico
	// Example: true
	// Required: true
	AccesoPublico *bool `json:"accesoPublico"`

	// acceso publico no publicado
	// Example: true
	// Required: true
	AccesoPublicoNoPublicado *bool `json:"accesoPublicoNoPublicado"`

	// description
	// Example: En este test se evaluaran los conocimientos respecto al lenguaje de programación Java
	// Required: true
	Description *string `json:"description"`

	// editable
	// Example: false
	// Required: true
	Editable *bool `json:"editable"`

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// max minutes
	// Example: 60
	// Required: true
	// Minimum: 0
	MaxMinutes *int64 `json:"maxMinutes"`

	// title
	// Example: Test de introduccion a Java
	// Required: true
	Title *string `json:"title"`

	// usuarioid
	// Example: 3
	// Required: true
	Usuarioid int64 `json:"usuarioid"`
}

// Answer answer
//
// swagger:model Answer
type Answer struct {

	// finished
	// Example: false
	// Required: true
	Finished *bool `json:"finished"`

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// startime
	// Example: 2021-02-25 14:44:55
	Startime string `json:"startime,omitempty"`

	// testid
	// Example: 343
	Testid int64 `json:"testid,omitempty"`

	// usuarioid
	// Example: 3
	// Required: true
	Usuarioid int64 `json:"usuarioid"`
}

// TeamRole team role
//
// swagger:model TeamRole
type TeamRole struct {

	// role
	// Required: true
	// Enum: [admin member]
	Role *string `json:"role"`
}

// Tag tag
//
// swagger:model Tag
type Tag struct {

	// tag
	// Example: Java
	// Required: true
	Tag *string `json:"tag"`
}

// QuestionAnswer question answer
//
// swagger:model QuestionAnswer
type QuestionAnswer struct {

	// corregida
	// Example: true
	// Required: true
	Corregida *bool `json:"corregida"`

	// id pregunta
	// Example: 1
	// Required: true
	IDPregunta *int64 `json:"idPregunta"`

	// indices opciones
	IndicesOpciones []int64 `json:"indicesOpciones"`

	// id respuesta
	// Example: 1
	// Required: true
	IDRespuesta *int64 `json:"idRespuesta"`

	// puntuacion
	// Example: 1
	// Required: true
	Puntuacion *int64 `json:"puntuacion"`

	// respuesta
	// Example: Javadoc
	Respuesta string `json:"respuesta,omitempty"`
}

// Option option
//
// swagger:model Option
type Option struct {

	// correcta
	// Example: false
	Correcta *bool `json:"correcta"`

	// indice
	// Example: 1
	Indice int64 `json:"indice,omitempty"`

	// preguntaid
	// Example: 1
	Preguntaid int64 `json:"preguntaid,omitempty"`

	// texto
	// Example: Esta opcion es la buena
	// Required: true
	Texto *string `json:"texto"`
}
