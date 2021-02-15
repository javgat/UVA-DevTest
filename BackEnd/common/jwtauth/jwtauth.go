package jwtauth

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type JwtJson struct{
  Token	string `json: "token"`
}

func CreateJwtJson(signedToken string) JwtJson{
  jsonToken := JwtJson{
    Token: signedToken,
  }
  return jsonToken
}

type JwtWrapper struct {
  SecretKey       string
  Issuer          string
  ExpirationHours int64
}

type JwtClaim struct {
  Email string
  jwt.StandardClaims
}

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

func (j *JwtWrapper) ValidateToken(signedToken string) (claims *JwtClaim, err error) {
  token, err := jwt.ParseWithClaims(
    signedToken,
    &JwtClaim{},
    func(token *jwt.Token) (interface{}, error) {
      return []byte(j.SecretKey), nil
    },
  )
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
