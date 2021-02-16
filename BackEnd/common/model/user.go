// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package model provides structures used in DevTest project
package model

type User struct{
  Username string `json: "username"`
  Email string `json: "email"`
  PwHash string `json: "pwhash"`
}
