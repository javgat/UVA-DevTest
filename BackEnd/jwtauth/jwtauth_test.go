package jwtauth

import "testing"

const email string = "Test@mail.com"

func TestGenerateToken(t *testing.T) {
	wrap := JwtWrapper{
		SecretKey:       "secret",
		Issuer:          "yo",
		ExpirationHours: 1,
	}
	_, err := wrap.GenerateToken(email)
	if err != nil {
		t.Log("err should be nil")
		t.Fail()
	}
}

func TestGenerateValidateToken(t *testing.T) {
	wrap := JwtWrapper{
		SecretKey:       "secret",
		Issuer:          "yo",
		ExpirationHours: 1,
	}
	signed, _ := wrap.GenerateToken(email)
	claims, err := wrap.ValidateToken(signed)
	if err != nil {
		t.Log("err should be nil")
		t.Fail()
	} else if claims == nil {
		t.Log("claims should not be nil")
		t.Fail()
	} else if claims.Email != email {
		t.Log("claims.Email should be email")
		t.Fail()
	}
}

func TestGenerateValidateTokenChangeSecret(t *testing.T) {
	wrap := JwtWrapper{
		SecretKey:       "secret",
		Issuer:          "yo",
		ExpirationHours: 1,
	}
	signed, _ := wrap.GenerateToken(email)
	wrap.SecretKey = "secret2"
	claims, err := wrap.ValidateToken(signed)
	if err == nil {
		t.Log("err should not be nil")
		t.Fail()
	} else if claims != nil {
		t.Log("claims should be nil")
		t.Fail()
	}
}
