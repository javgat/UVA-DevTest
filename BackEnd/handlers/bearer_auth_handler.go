package handlers

import (
	"errors"
	"uva-devtest/jwtauth"
	"uva-devtest/models"
	"uva-devtest/persistence/dao"
	"uva-devtest/persistence/dbconnection"
)

// BearerAuth gets the model User for the token, if valid JWT
func BearerAuth(cookie string) (*models.User, error) {
	// PRECACUCION Si hay mas de una cookie esto petaria si no va la primera, hacer bien
	expectedName := "Bearer-Cookie="
	cookieName := cookie[0:14]
	if expectedName != cookieName {
		return nil, errors.New("no se puede leer la cookie Bearer-Cookie")
	}
	token := cookie[14:]
	email, err := jwtauth.GetEmailToken(token)
	if err != nil {
		return nil, err
	}
	db, err := dbconnection.ConnectDb()
	if err != nil {
		return nil, err
	}
	u, err := dao.GetUserEmail(db, email)
	if u == nil || err != nil {
		return nil, err
	}
	var wrap jwtauth.JwtWrapper
	wrap.SecretKey = *u.Pwhash
	wrap.Issuer = "DevTest"
	wrap.ExpirationHours = 24
	_, err = wrap.ValidateToken(token)
	if err != nil {
		return nil, err
	}
	mu := dao.ToModelUser(u)
	return mu, err
}
