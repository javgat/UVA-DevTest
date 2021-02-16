package service

type SigninUser struct {
  Username string `json: "username"`
  Email    string `json: "email"`
  Pass   string `json: "pass"`
}
