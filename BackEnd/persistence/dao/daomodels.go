package dao

import "github.com/go-openapi/strfmt"

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
	Fullname string `json:"fullname,omitempty"`

	// rol
	// Enum: [student teacher admin]
	Rol string `json:"rol,omitempty"`
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

	// testid
	// Example: 1
	Testid int64 `json:"testid,omitempty"`

	// title
	// Example: Paralelismo en C
	// Required: true
	Title *string `json:"title"`

	// usuarioid
	// Example: 3
	// Required: true
	Usuarioid int64 `json:"usuarioid"`
}

// Test test
//
// swagger:model Test
type Test struct {

	// acceso publico
	// Example: true
	// Required: true
	AccesoPublico *bool `json:"accesoPublico"`

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

	// max seconds
	// Example: 32600
	// Required: true
	MaxSeconds *int64 `json:"maxSeconds"`

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
