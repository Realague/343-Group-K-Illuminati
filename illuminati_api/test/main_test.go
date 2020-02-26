package test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"343-Group-K-Illuminati/illuminati_api/config"
	"testing"
)

func TestUserMeDelete401(t *testing.T) {
	var jsonStr = []byte(`{
		"identifier": "testtest",
		"password": "admin"
		}`)
	var w = PerformRequest(router, "POST", "/api/auth/local", jsonStr, config.Config.TokenAuthenticationKey)
	var response loginResponse
	_ = json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Equal(t, 200, w.Code)

	w = PerformRequest(router, "DELETE", "/api/users/me", jsonStr, "")
	assert.Equal(t, 401, w.Code)
}

func TestUserDeleteById401(t *testing.T) {
	var jsonStr = []byte(`{
			"identifier": "admin",
			"password": "admin"
			}`)
	var w = PerformRequest(router, "POST", "/api/auth/local", jsonStr, config.Config.TokenAuthenticationKey)
	var response loginResponse
	_ = json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Equal(t, 200, w.Code)

	w = PerformRequest(router, "DELETE", "/api/users/"+response.User.Id.Hex(), jsonStr, userToken)
	assert.Equal(t, 401, w.Code)
}

func TestUserMeDeleteSuccessful(t *testing.T) {
	var jsonStr = []byte(`{
		"identifier": "testtest",
		"password": "admin"
		}`)
	var w = PerformRequest(router, "POST", "/api/auth/local", jsonStr, config.Config.TokenAuthenticationKey)
	var response loginResponse
	_ = json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Equal(t, 200, w.Code)
	w = PerformRequest(router, "DELETE", "/api/users/me", jsonStr, response.AccessToken)
	assert.Equal(t, 204, w.Code)
}

func TestUserDeleteById404(t *testing.T) {
	var jsonStr = []byte(`{
			"identifier": "admin",
			"password": "admin"
			}`)
	var w = PerformRequest(router, "POST", "/api/auth/local", jsonStr, config.Config.TokenAuthenticationKey)
	var response loginResponse
	_ = json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Equal(t, 200, w.Code)

	w = PerformRequest(router, "DELETE", "/api/users/507f1f77bcf86cd799439011", jsonStr, response.AccessToken)
	assert.Equal(t, 404, w.Code)
}

func TestUserDeleteByIdSuccessful(t *testing.T) {
	var jsonStr = []byte(`{
			"identifier": "admin",
			"password": "admin"
			}`)
	var w = PerformRequest(router, "POST", "/api/auth/local", jsonStr, config.Config.TokenAuthenticationKey)
	var response loginResponse
	_ = json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Equal(t, 200, w.Code)

	w = PerformRequest(router, "DELETE", "/api/users/"+response.User.Id.Hex(), jsonStr, response.AccessToken)
	assert.Equal(t, 204, w.Code)
}

func TestOptionsRoute(t *testing.T) {
	var w = PerformRequest(router, "OPTIONS", "/api/auth/local", nil, config.Config.TokenAuthenticationKey)
	assert.Equal(t, 200, w.Code)
}
