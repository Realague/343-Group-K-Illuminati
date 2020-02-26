package database

import (
	"343-Group-K-Illuminati/illuminati_api/config"
	"343-Group-K-Illuminati/illuminati_api/utils"

	mgo "gopkg.in/mgo.v2"
)

// DatabaseService is a struct that contain all information about the database
type DatabaseService struct {
	host        string
	database    string
	collections map[string]string
	db          *mgo.Database
}

var (
	// Database is the entire service
	Database DatabaseService
)

// Connect is used to config & connect the api to the database
func Connect() *DatabaseService {
	Database.Config(config.Config.DatabaseHost, config.Config.DatabaseName, map[string]string{
		"users":                    "users",
		"forbidden_mail_addresses": "forbidden_mail_addresses",
	})
	if !config.Config.IsUnitTest {
		Database.Connect()
	}
	if hash, err := utils.HashPassword("admin123"); err == nil {
		Database.FindOrCreate("admin", hash, "admin@apinit-go.eu", true)
	}
	return &Database
}
