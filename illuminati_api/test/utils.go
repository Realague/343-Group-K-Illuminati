package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
)

func PerformRequest(r http.Handler, method, path string, toPost []byte, bearer string) *httptest.ResponseRecorder {
	var req *http.Request
	var token = ""
	if toPost != nil {
		req, _ = http.NewRequest(method, path, bytes.NewBuffer(toPost))
	} else {
		req, _ = http.NewRequest(method, path, nil)
	}
	if bearer != "" {
		token = "Bearer " + bearer
		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("Authorization", token)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
