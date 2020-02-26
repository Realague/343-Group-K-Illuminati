package test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"343-Group-K-Illuminati/illuminati_api/config"
	"testing"
)

func TestUsersGetSuccessful(t *testing.T) {
	var jsonStr = []byte(`{
		"identifier": "admin",
		"password": "admin"
		}`)
	var w = PerformRequest(router, "POST", "/api/auth/local", jsonStr, config.Config.TokenAuthenticationKey)
	var response loginResponse
	_ = json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Equal(t, 200, w.Code)

	w = PerformRequest(router, "GET", "/api/users", nil, response.AccessToken)
	assert.Equal(t, 200, w.Code)
}

func TestUserGetByIdSuccessful(t *testing.T) {
	var jsonStr = []byte(`{
			"identifier": "admin",
			"password": "admin"
			}`)
	var w = PerformRequest(router, "POST", "/api/auth/local", jsonStr, config.Config.TokenAuthenticationKey)
	var response loginResponse
	_ = json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Equal(t, 200, w.Code)

	w = PerformRequest(router, "GET", "/api/users/"+response.User.Id.Hex(), jsonStr, response.AccessToken)
	assert.Equal(t, 200, w.Code)
}

func TestUserGetById401(t *testing.T) {
	var jsonStr = []byte(`{
			"identifier": "admin",
			"password": "admin"
			}`)
	var w = PerformRequest(router, "POST", "/api/auth/local", jsonStr, config.Config.TokenAuthenticationKey)
	var response loginResponse
	_ = json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Equal(t, 200, w.Code)

	w = PerformRequest(router, "GET", "/api/users/"+response.User.Id.Hex(), jsonStr, userToken)
	assert.Equal(t, 401, w.Code)
}

func TestUserGetById404(t *testing.T) {
	var jsonStr = []byte(`{
			"identifier": "admin",
			"password": "admin"
			}`)
	var w = PerformRequest(router, "POST", "/api/auth/local", jsonStr, config.Config.TokenAuthenticationKey)
	var response loginResponse
	_ = json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Equal(t, 200, w.Code)

	w = PerformRequest(router, "GET", "/api/users/507f1f77bcf86cd799439011", jsonStr, response.AccessToken)
	assert.Equal(t, 404, w.Code)
}
