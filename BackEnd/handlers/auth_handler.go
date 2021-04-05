package handlers

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strings"
	"uva-devtest/jwtauth"
	"uva-devtest/models"
	"uva-devtest/persistence/dao"
	"uva-devtest/persistence/dbconnection"
	"uva-devtest/restapi/operations/auth"

	"github.com/go-openapi/runtime/middleware"
)

const BearerCookieName string = "Bearer-Cookie"
const ReauthCookieName string = "ReAuth-Cookie"

const AuthHours int = 1
const ReauthHours int = 48

func CreateCookie(name string, token string, maxage int) *http.Cookie {

	cookie := &http.Cookie{
		Name:     name,
		Value:    token,
		Path:     "/",
		HttpOnly: true,                    // Evita ataques XSS
		Secure:   true,                    // Fuerza HTTPS
		MaxAge:   maxage,                  //Poner fin en 24h
		SameSite: http.SameSiteStrictMode, // Evita ataques XSRF
	}
	/*path := "/"
	maxage := "86400"
	samesite := "strict"
	cookie := fmt.Sprintf("%s=%s; Path=%s; Secure; SameSite=%s; HttpOnly; Max-Age=%s ", name, token, path, samesite, maxage)*/
	return cookie
}

func CreateDeprecatedCookie(name string) *http.Cookie {
	/*path := "/"
	samesite := "strict"
	cookie := fmt.Sprintf("%s=; Path=%s; Secure; SameSite=%s; HttpOnly; expires=Thu, 01 Jan 1970 00:00:00 GMT", name, path, samesite)*/
	return CreateCookie(name, "", 1)
}

func hoursToSeconds(hours int) int {
	return hours * 3600
}

func CreateJWTWrapper(u dao.User, expirationHours int64) jwtauth.JwtWrapper {
	var wrap jwtauth.JwtWrapper
	wrap.SecretKey = *u.Pwhash
	wrap.Issuer = "DevTest"
	wrap.ExpirationHours = expirationHours
	return wrap
}

func CreateJWT(u dao.User, expirationHours int64) (string, error) {
	wrap := CreateJWTWrapper(u, expirationHours)
	signedToken, err := wrap.GenerateToken(u.Email.String())
	log.Println(wrap.SecretKey, u.Email.String(), signedToken)
	return signedToken, err
}

func GetJWTModelUserCookies(cookieSlice []string, expectedName string) (*models.User, error) {
	for _, cookie := range cookieSlice {
		cookieName := cookie[0:14]
		var err error
		if expectedName == cookieName {
			token := cookie[14:]
			var email string
			email, err = jwtauth.GetEmailToken(token)
			if err == nil {
				var db *sql.DB
				db, err = dbconnection.ConnectDb()
				if err == nil {
					var u *dao.User
					u, err = dao.GetUserEmail(db, email)
					if u != nil || err == nil {
						wrap := CreateJWTWrapper(*u, int64(AuthHours))
						_, err = wrap.ValidateToken(token)
						if err == nil {
							mu := dao.ToModelUser(u)
							return mu, err
						}
					}
				}
			}
		}
		log.Println("Cookie incorrecta: ", err)
	}
	errMsg := strings.Join([]string{"no se puede leer la cookie", expectedName}, " ")
	return nil, errors.New(errMsg)
}

// BearerAuth gets the model User for the token, if valid JWT
func BearerAuth(cookies string) (*models.User, error) {
	// PRECACUCION Si hay mas de una cookie esto petaria si no va la primera, hacer bien
	expectedName := "Bearer-Cookie="
	cookieSlice := strings.Split(cookies, ";")
	mu, err := GetJWTModelUserCookies(cookieSlice, expectedName)
	if err == nil && mu != nil {
		return mu, err
	}
	return nil, err
}

func Logout(auth.LogoutParams) middleware.Responder {
	bcookie := CreateDeprecatedCookie(BearerCookieName)
	rcookie := CreateDeprecatedCookie(ReauthCookieName)
	return auth.NewLogoutOK().WithAuth(bcookie).WithReAuth(rcookie)
}

func Relogin(auth.ReloginParams) middleware.Responder {

}
