package test

import (
	"github.com/stretchr/testify/assert"
	"343-Group-K-Illuminati/illuminati_api/config"
	"testing"
)

func TestRecoverPasswordSuccessful(t *testing.T) {
	var jsonStr = []byte(`{
			"email": "admin"
			}`)
	w := PerformRequest(router, "POST", "/api/auth/recover-password", jsonStr, config.Config.TokenAuthenticationKey)
	assert.Equal(t, 204, w.Code)
}

func TestRecoverPassword401(t *testing.T) {
	var jsonStr = []byte(`{
			"email": "admin"
			}`)
	w := PerformRequest(router, "POST", "/api/auth/recover-password", jsonStr, "")
	assert.Equal(t, 401, w.Code)
}

func TestRecoverPassword400WithoutBody(t *testing.T) {
	var jsonStr = []byte(``)
	w := PerformRequest(router, "POST", "/api/auth/recover-password", jsonStr, config.Config.TokenAuthenticationKey)
	assert.Equal(t, 400, w.Code)
}

func TestRecoverPassword400(t *testing.T) {
	var jsonStr = []byte(`{
			"test": "admin"
			}`)
	w := PerformRequest(router, "POST", "/api/auth/recover-password", jsonStr, config.Config.TokenAuthenticationKey)
	assert.Equal(t, 400, w.Code)
}

func TestRecoverPassword404(t *testing.T) {
	var jsonStr = []byte(`{
			"email": "404"
			}`)
	w := PerformRequest(router, "POST", "/api/auth/recover-password", jsonStr, config.Config.TokenAuthenticationKey)
	assert.Equal(t, 404, w.Code)
}
