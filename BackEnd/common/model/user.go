package model

type User struct{
  Username string `json: "username"`
  Email string `json: "email"`
  PwHash string `json: "pwhash"`
}
