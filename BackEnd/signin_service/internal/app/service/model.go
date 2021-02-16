// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

package service

// SigninUser represents the json given by the client when registering new user
type SigninUser struct {
  Username string `json: "username"`
  Email    string `json: "email"`
  Pass   string `json: "pass"`
}
