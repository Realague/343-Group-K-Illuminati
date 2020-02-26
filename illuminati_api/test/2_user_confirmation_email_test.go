package test

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/dgrijalva/jwt-go.v3"
	"343-Group-K-Illuminati/illuminati_api/config"
	"testing"
	"time"
)

func TestResendConfirmationMailSuccessful(t *testing.T) {
	var jsonStr = []byte(`{
			"email": "admin"
			}`)
	w := PerformRequest(router, "POST", "/api/auth/send-confirmation-mail", jsonStr, config.Config.TokenAuthenticationKey)
	assert.Equal(t, 204, w.Code)
}

func TestResendConfirmationMail401(t *testing.T) {
	var jsonStr = []byte(`{
			"email": "admin"
			}`)
	w := PerformRequest(router, "POST", "/api/auth/send-confirmation-mail", jsonStr, "")
	assert.Equal(t, 401, w.Code)
}

func TestResendConfirmationMail400WithoutBody(t *testing.T) {
	var jsonStr = []byte(``)
	w := PerformRequest(router, "POST", "/api/auth/send-confirmation-mail", jsonStr, config.Config.TokenAuthenticationKey)
	assert.Equal(t, 400, w.Code)
}

func TestResendConfirmationMail400(t *testing.T) {
	var jsonStr = []byte(`{
			"test": "admin"
			}`)
	w := PerformRequest(router, "POST", "/api/auth/send-confirmation-mail", jsonStr, config.Config.TokenAuthenticationKey)
	assert.Equal(t, 400, w.Code)
}

func TestResendConfirmationMail404(t *testing.T) {
	var jsonStr = []byte(`{
			"email": "notfound"
			}`)
	w := PerformRequest(router, "POST", "/api/auth/send-confirmation-mail", jsonStr, config.Config.TokenAuthenticationKey)
	assert.Equal(t, 404, w.Code)
}

func TestConfirmationMail400(t *testing.T) {
	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  "oo",
		"exp": time.Now().Add(config.Config.RecoverPasswordTokenValidityTime * time.Minute).Unix(),
	}).SignedString([]byte(config.Config.TokenEncryptionKey))

	w := PerformRequest(router, "GET", "/api/auth/confirm-email/"+token, nil, "")
	assert.Equal(t, 404, w.Code)
}