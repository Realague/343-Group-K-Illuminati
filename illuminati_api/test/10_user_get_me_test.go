package test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"343-Group-K-Illuminati/illuminati_api/config"
	"343-Group-K-Illuminati/illuminati_api/models/db"
	"testing"
)

func TestUserMeGetSuccessful(t *testing.T) {
	var jsonStr = []byte(`{
			"identifier": "admin",
			"password": "admin"
			}`)
	var w = PerformRequest(router, "POST", "/api/auth/local", jsonStr, config.Config.TokenAuthenticationKey)
	var response loginResponse
	_ = json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Equal(t, 200, w.Code)

	w = PerformRequest(router, "GET", "/api/users/me", nil, response.AccessToken)
	var resp db.User
	_ = json.Unmarshal([]byte(w.Body.String()), &resp)
	assert.Equal(t, 200, w.Code)
}
func TestUserMeGet401(t *testing.T) {
	var jsonStr = []byte(`{
		"identifier": "testtest",
		"password": "admin"
		}`)
	var w = PerformRequest(router, "POST", "/api/auth/local", jsonStr, config.Config.TokenAuthenticationKey)
	var response loginResponse
	_ = json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Equal(t, 200, w.Code)

	w = PerformRequest(router, "GET", "/api/users/me", nil, "")
	var resp db.User
	_ = json.Unmarshal([]byte(w.Body.String()), &resp)
	assert.Equal(t, 401, w.Code)
}
