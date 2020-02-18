package test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"343-Group-K-Illuminati/illuminati_api/config"
	"343-Group-K-Illuminati/illuminati_api/models/db"
	"testing"
)

func TestCreateUserSuccessful(t *testing.T) {
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
	assert.Equal(t, 201, w.Code)

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

func TestCreateUser409(t *testing.T) {
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
	assert.Equal(t, 201, w.Code)

	w = PerformRequest(router, "POST", "/api/auth/local", jsonStr, config.Config.TokenAuthenticationKey)
	assert.Equal(t, 409, w.Code)

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

func TestCreateUser409WithSameEmail(t *testing.T) {
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
	assert.Equal(t, 201, w.Code)

	jsonStr = []byte(`{
		"username":"gaming",
		"email": "testing@gmail.com",
		"password": "admin"
	}`)
	w = PerformRequest(router, "POST", "/api/auth/local", jsonStr, config.Config.TokenAuthenticationKey)
	assert.Equal(t, 409, w.Code)

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

func TestCreateUser409WithSameUsername(t *testing.T) {
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
	assert.Equal(t, 201, w.Code)

	jsonStr = []byte(`{
		"username":"createnew",
		"email": "newemail@gmail.com",
		"password": "admin"
	}`)
	w = PerformRequest(router, "POST", "/api/auth/local", jsonStr, config.Config.TokenAuthenticationKey)
	assert.Equal(t, 409, w.Code)

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

func TestCreateUser400WithoutBody(t *testing.T) {
	var jsonStr = []byte(`{
		}`)
	var w = PerformRequest(router, "POST", "/api/auth/local", jsonStr, config.Config.TokenAuthenticationKey)
	assert.Equal(t, 400, w.Code)
}

