package test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"343-Group-K-Illuminati/illuminati_api/config"
	"testing"
)

func TestFeedback400(t *testing.T) {
	var jsonStr = []byte(`{
			"identifier": "admin",
			"password": "admin"
			}`)
	var w = PerformRequest(router, "POST", "/api/auth/local", jsonStr, config.Config.TokenAuthenticationKey)
	var response loginResponse
	_ = json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Equal(t, 200, w.Code)

	jsonStr = []byte(`{{
			"title": ""
			}`)
	w = PerformRequest(router, "POST", "/api/feedback", jsonStr, response.AccessToken)
	assert.Equal(t, 400, w.Code)
}

func TestFeedback400WithoutBody(t *testing.T) {
	var jsonStr = []byte(`{
			"identifier": "admin",
			"password": "admin"
			}`)
	var w = PerformRequest(router, "POST", "/api/auth/local", jsonStr, config.Config.TokenAuthenticationKey)
	var response loginResponse
	_ = json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Equal(t, 200, w.Code)

	jsonStr = []byte(``)
	w = PerformRequest(router, "POST", "/api/feedback", jsonStr, response.AccessToken)
	assert.Equal(t, 400, w.Code)
}

func TestFeedbackSuccessful(t *testing.T) {
	var jsonStr = []byte(`{
			"identifier": "admin",
			"password": "admin"
			}`)
	var w = PerformRequest(router, "POST", "/api/auth/local", jsonStr, config.Config.TokenAuthenticationKey)
	var response loginResponse
	_ = json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Equal(t, 200, w.Code)

	jsonStr = []byte(`{
			"title": "tesr",
			"body": "test"
			}`)
	w = PerformRequest(router, "POST", "/api/feedback", jsonStr, response.AccessToken)
	assert.Equal(t, 204, w.Code)
}

func TestFeedback401(t *testing.T) {
	var jsonStr = []byte(`{
			"title": "",
			"body": ""
			}`)
	var w = PerformRequest(router, "POST", "/api/feedback", jsonStr, config.Config.TokenAuthenticationKey)
	assert.Equal(t, 401, w.Code)
}

func TestLoginSuccessful(t *testing.T) {
	var jsonStr = []byte(`{
		"identifier": "admin",
		"password": "admin"
		}`)
	var w = PerformRequest(router, "POST", "/api/auth/local", jsonStr, config.Config.TokenAuthenticationKey)
	var response loginResponse
	_ = json.Unmarshal([]byte(w.Body.String()), &response)
}

