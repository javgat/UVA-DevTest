package dao

import "github.com/go-openapi/strfmt"

// User user
//
// userdao User
type User struct {
	// id
	// Identificador SQL
	ID int

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
