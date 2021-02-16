// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

package service

// LoginUser represents de json given by the client trying to log in
type LoginUser struct {
  LoginId string `json: "loginid"`
  Pass   string `json: "pass"`
}
