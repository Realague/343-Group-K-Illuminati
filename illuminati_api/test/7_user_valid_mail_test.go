package test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"343-Group-K-Illuminati/illuminati_api/config"
	"343-Group-K-Illuminati/illuminati_api/models/db"
	"testing"
)

func TestValidEmailSuccessful(t *testing.T) {
	var jsonStr = []byte(`{
		"identifier": "admin",
		"password": "admin"
		}`)
	var w = PerformRequest(router, "POST", "/api/auth/local", jsonStr, config.Config.TokenAuthenticationKey)
	var response loginResponse
	_ = json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Equal(t, 200, w.Code)

	jsonStr = []byte(`{
		"username":"createnew",
		"email": "testing@gmail.com",
		"password": "admin"
		}`)
	w = PerformRequest(router, "POST", "/api/auth/local", jsonStr, config.Config.TokenAuthenticationKey)
	var resp EmailValidResponse
	_ = json.Unmarshal([]byte(w.Body.String()), &resp)
	assert.Equal(t, 201, w.Code)

	w = PerformRequest(router, "GET", "/api/auth/confirm-email/"+resp.Token, nil, "")
	assert.Equal(t, 308, w.Code)

	jsonStr = []byte(`{
		"identifier": "createnew",
		"password": "admin"
		}`)
	w = PerformRequest(router, "POST", "/api/auth/local", jsonStr, config.Config.TokenAuthenticationKey)
	assert.Equal(t, 200, w.Code)

	w = PerformRequest(router, "GET", "/api/users", nil, response.AccessToken)
	var users []db.User
	err := json.Unmarshal([]byte(w.Body.String()), &users)
	if err != nil {
	}
	for _, v := range users {
		if v.Username == "createnew" {
			w = PerformRequest(router, "DELETE", "/api/users/"+v.Id.Hex(), nil, response.AccessToken)
			assert.Equal(t, 204, w.Code)
			break
		}
	}
}

func TestValidEmail400(t *testing.T) {
	var jsonStr = []byte(`{
		"identifier": "admin",
		"password": "admin"
		}`)
	var w = PerformRequest(router, "POST", "/api/auth/local", jsonStr, config.Config.TokenAuthenticationKey)
	var response loginResponse
	_ = json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Equal(t, 200, w.Code)

	jsonStr = []byte(`{
		"username":"createnew",
		"email": "testing@gmail.com",
		"password": "admin"
		}`)
	w = PerformRequest(router, "POST", "/api/auth/local", jsonStr, config.Config.TokenAuthenticationKey)
	var resp EmailValidResponse
	_ = json.Unmarshal([]byte(w.Body.String()), &resp)
	assert.Equal(t, 201, w.Code)

	w = PerformRequest(router, "GET", "/api/auth/confirm-email/eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c", nil, "")
	assert.Equal(t, 400, w.Code)

	jsonStr = []byte(`{
		"identifier": "createnew",
		"password": "admin"
		}`)
	w = PerformRequest(router, "POST", "/api/auth/local", jsonStr, config.Config.TokenAuthenticationKey)
	assert.Equal(t, 401, w.Code)

	w = PerformRequest(router, "GET", "/api/users", nil, response.AccessToken)
	var users []db.User
	err := json.Unmarshal([]byte(w.Body.String()), &users)
	if err != nil {
	}
	for _, v := range users {
		if v.Username == "createnew" {
			w = PerformRequest(router, "DELETE", "/api/users/"+v.Id.Hex(), nil, response.AccessToken)
			assert.Equal(t, 204, w.Code)
			break
		}
	}
}
