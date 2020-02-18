package test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"343-Group-K-Illuminati/illuminati_api/config"
	"343-Group-K-Illuminati/illuminati_api/models/db"
	"testing"
)

func TestUserPutByIDSuccessful(t *testing.T) {
	var jsonStr = []byte(`{
		"identifier": "admin",
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
			"verified": true,
			"admin": true
		}`)
	w = PerformRequest(router, "PUT", "/api/users/"+response.User.Id.Hex(), jsonStr, response.AccessToken)
	var resp db.User
	_ = json.Unmarshal([]byte(w.Body.String()), &resp)
	assert.Equal(t, 200, w.Code)
}

func TestUserPutByID401(t *testing.T) {
	var jsonStr = []byte(`{
		"identifier": "admin",
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
	w = PerformRequest(router, "PUT", "/api/users/"+response.User.Id.Hex(), jsonStr, userToken)
	var resp db.User
	_ = json.Unmarshal([]byte(w.Body.String()), &resp)
	assert.Equal(t, 401, w.Code)
}

func TestUserPutByID400WithoutBody(t *testing.T) {
	var jsonStr = []byte(`{
		"identifier": "admin",
		"password": "admin"
		}`)
	var w = PerformRequest(router, "POST", "/api/auth/local", jsonStr, config.Config.TokenAuthenticationKey)
	var response loginResponse
	_ = json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Equal(t, 200, w.Code)

	jsonStr = []byte(`{
		}`)
	w = PerformRequest(router, "PUT", "/api/users/"+response.User.Id.Hex(), jsonStr, adminToken)
	var resp db.User
	_ = json.Unmarshal([]byte(w.Body.String()), &resp)
	assert.Equal(t, 400, w.Code)
}

func TestUserPutByID404(t *testing.T) {
	var jsonStr = []byte(`{
		"identifier": "admin",
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
	w = PerformRequest(router, "PUT", "/api/users/507f1f77bcf86cd799439011", jsonStr, response.AccessToken)
	var resp db.User
	_ = json.Unmarshal([]byte(w.Body.String()), &resp)
	assert.Equal(t, 404, w.Code)
}