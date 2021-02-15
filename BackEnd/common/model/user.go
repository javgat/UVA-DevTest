package model

type LoginUser struct {
  Username string `json: "username"`
  Email    string `json: "email"`
  Pass   string `json: "pass"`
}

type User struct{
  Username string `json: "username"`
  Email string `json: "email"`
  PwHash string `json: "pwhash"`
}
