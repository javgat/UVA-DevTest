package handlers

import (
	"uva-devtest/jwtauth"
	"uva-devtest/models"
	"uva-devtest/persistence/dao"
	"uva-devtest/persistence/dbconnection"
)

// BearerAuth gets the model User for the token, if valid JWT
func BearerAuth(token string) (*models.User, error) {
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
	mu := dao.DaoToModelUser(u)
	return mu, err
}
