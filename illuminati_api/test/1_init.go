package test

import (
	"343-Group-K-Illuminati/illuminati_api/config"
	"343-Group-K-Illuminati/illuminati_api/controllers"
	"343-Group-K-Illuminati/illuminati_api/database"
	"343-Group-K-Illuminati/illuminati_api/models/db"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

type loginResponse struct {
	User        db.User `json:"user"`
	AccessToken string  `json:"access_token"`
}

type EmailValidResponse struct {
	User  db.User `json:"user"`
	Token string  `json:"token"`
}

var (
	router     http.Handler
	adminToken string
	userToken  string
)

func init() {
	config.InitConfig()
	router = controllers.InitRouter()

	adminUser := &db.User{
		Id:        bson.NewObjectId(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Username:  "admin",
		Password:  "$2y$10$aShL5BCnDwdGcoEcwBCkl.CCtHj4t5CokaJOmqYDOdL7up2yArRA.",
		Email:     "admin",
		Verified:  true,
		Admin:     true,
	}

	testUser := &db.User{
		Id:        bson.NewObjectId(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Username:  "test",
		Password:  "$2y$10$aShL5BCnDwdGcoEcwBCkl.CCtHj4t5CokaJOmqYDOdL7up2yArRA.",
		Email:     "test",
		Verified:  false,
		Admin:     false,
	}

	testUser2 := &db.User{
		Id:        bson.NewObjectId(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Username:  "testtest",
		Password:  "$2y$10$aShL5BCnDwdGcoEcwBCkl.CCtHj4t5CokaJOmqYDOdL7up2yArRA.",
		Email:     "testtest",
		Verified:  true,
		Admin:     false,
	}

	_ = database.Database.InsertUser(adminUser)
	_ = database.Database.InsertUser(testUser)
	_ = database.Database.InsertUser(testUser2)

	var jsonStr = []byte(`{
			"identifier": "admin",
			"password": "admin"
			}`)
	var w = PerformRequest(router, "POST", "/api/auth/local", jsonStr, config.Config.TokenAuthenticationKey)
	var response loginResponse
	_ = json.Unmarshal([]byte(w.Body.String()), &response)
	adminToken = response.AccessToken

	jsonStr = []byte(`{
		"identifier": "testtest",
		"password": "admin"
		}`)
	w = PerformRequest(router, "POST", "/api/auth/local", jsonStr, config.Config.TokenAuthenticationKey)
	_ = json.Unmarshal([]byte(w.Body.String()), &response)
	userToken = response.AccessToken
}
