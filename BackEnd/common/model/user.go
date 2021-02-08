package model

type User struct{
  // Quiza añadir ID, a la tabla seguro
  Username string `json: "username"`
  Email string `json: "email"`
  PwHash string `json: "pwhash"`
}
