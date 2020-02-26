package test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"343-Group-K-Illuminati/illuminati_api/config"
	"testing"
)

func TestRefreshTokenSuccessful(t *testing.T) {
	var jsonStr = []byte(`{
		"identifier": "admin",
		"password": "admin"
		}`)
	var w = PerformRequest(router, "POST", "/api/auth/local", jsonStr, config.Config.TokenAuthenticationKey)
	var response loginResponse
	_ = json.Unmarshal([]byte(w.Body.String()), &response)

	jsonStr = []byte(`{
		"refresh_token":"` + response.AccessToken + `"
		}`)
	w = PerformRequest(router, "POST", "/api/auth/refresh-token", jsonStr, config.Config.TokenAuthenticationKey)
	assert.Equal(t, 200, w.Code)
}

func TestRefreshToken400(t *testing.T) {
	var jsonStr = []byte(`{
		"identifier": "admin",
		"password": "admin"
		}`)
	var w = PerformRequest(router, "POST", "/api/auth/local", jsonStr, config.Config.TokenAuthenticationKey)
	var response loginResponse
	_ = json.Unmarshal([]byte(w.Body.String()), &response)

	jsonStr = []byte(`{
		"refresh_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
		}`)
	w = PerformRequest(router, "POST", "/api/auth/refresh-token", jsonStr, config.Config.TokenAuthenticationKey)
	assert.Equal(t, 400, w.Code)
}

func TestRefreshToken400WithoutBody(t *testing.T) {
	var jsonStr = []byte(`{
		"identifier": "admin",
		"password": "admin"
		}`)
	var w = PerformRequest(router, "POST", "/api/auth/local", jsonStr, config.Config.TokenAuthenticationKey)
	var response loginResponse
	_ = json.Unmarshal([]byte(w.Body.String()), &response)

	jsonStr = []byte(`{
		}`)
	w = PerformRequest(router, "POST", "/api/auth/refresh-token", jsonStr, config.Config.TokenAuthenticationKey)
	assert.Equal(t, 400, w.Code)
}
