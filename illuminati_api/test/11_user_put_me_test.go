package test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"343-Group-K-Illuminati/illuminati_api/config"
	"343-Group-K-Illuminati/illuminati_api/models/db"
	"testing"
)

func TestUserMePutSuccessful(t *testing.T) {
	var jsonStr = []byte(`{
		"identifier": "testtest",
		"password": "admin"
		}`)
	var w = PerformRequest(router, "POST", "/api/auth/local", jsonStr, config.Config.TokenAuthenticationKey)
	var response loginResponse
	_ = json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Equal(t, 200, w.Code)

	jsonStr = []byte(`{
			"username": "` + response.User.Username + `",
			"password": "admin",
			"email": "` + response.User.Email + `",
			"verified": true
		}`)
	w = PerformRequest(router, "PUT", "/api/users/me", jsonStr, response.AccessToken)
	var resp db.User
	_ = json.Unmarshal([]byte(w.Body.String()), &resp)
	assert.Equal(t, 200, w.Code)
}

func TestUserMePut401(t *testing.T) {
	var jsonStr = []byte(`{
		"identifier": "testtest",
		"password": "admin"
		}`)
	var w = PerformRequest(router, "POST", "/api/auth/local", jsonStr, config.Config.TokenAuthenticationKey)
	var response loginResponse
	_ = json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Equal(t, 200, w.Code)

	jsonStr = []byte(`{
			"username": "` + response.User.Username + `",
			"password": "admin",
			"email": "` + response.User.Email + `",
			"verified": true
		}`)
	w = PerformRequest(router, "PUT", "/api/users/me", jsonStr, "")
	var resp db.User
	_ = json.Unmarshal([]byte(w.Body.String()), &resp)
	assert.Equal(t, 401, w.Code)
}

func TestUserMePut401WithoutBody(t *testing.T) {
	var jsonStr = []byte(`{
		"identifier": "testtest",
		"password": "admin"
		}`)
	var w = PerformRequest(router, "POST", "/api/auth/local", jsonStr, config.Config.TokenAuthenticationKey)
	var response loginResponse
	_ = json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Equal(t, 200, w.Code)

	jsonStr = []byte(`{
		}`)
	w = PerformRequest(router, "PUT", "/api/users/me", jsonStr, "")
	var resp db.User
	_ = json.Unmarshal([]byte(w.Body.String()), &resp)
	assert.Equal(t, 401, w.Code)
}
