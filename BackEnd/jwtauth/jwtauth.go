// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package jwtauth provides functions and structs that allowes the generation
// and validation of JWTs
package jwtauth

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"uva-devtest/models"
)

// CreateJwtJSON creates a JwtJson from the signed jwt in the string
// Param signedToken represents the JWT Token
// Returns models.JWTJSON to send
func CreateJwtJSON(signedToken string) models.JWTJSON {
	jsonToken := models.JWTJSON{
		Token: &signedToken,
	}
	return jsonToken
}

// JwtWrapper represents the wrapper that contains values of the JWT except for
// the claims
type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

// JwtClaim contains the claims that will issue the JWT.
type JwtClaim struct {
	Email string
	jwt.StandardClaims
}

// GenerateToken creates a signed token with an email claim, identifying the user
// Param email: Email identifying the user
// Return signedToken: Signed JWT Token
// Return err: Error, if any
func (j *JwtWrapper) GenerateToken(email string) (signedToken string, err error) {
	claims := &JwtClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(j.ExpirationHours)).Unix(),
			Issuer:    j.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString([]byte(j.SecretKey))
	return signedToken, err
}

// ValidateToken validates the token signedToken and returns the claims
// Param signedToken: Token that will be validated
// Return claims: *JwtClaim that the token is claiming
// Return err: error if something is wrong or if validation is unsuccessful
func (j *JwtWrapper) ValidateToken(signedToken string) (claims *JwtClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
		},
	)
	//Para validar en realidad tendre que hacer nuevas funcioness que
	//separen validacion de secreto de obtencion de claims
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JwtClaim)
	if !ok {
		err = errors.New("Couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("JWT is expired")
		return
	}
	return
}
