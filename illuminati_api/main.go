package main

import (
	"343-Group-K-Illuminati/illuminati_api/config"
	"343-Group-K-Illuminati/illuminati_api/controllers"
	"343-Group-K-Illuminati/illuminati_api/database"
	"github.com/gin-contrib/cors"
)

func main() {
	config.InitConfig()
	if !config.Config.IsUnitTest {
		db := database.Connect()
		defer db.GetDatabase().Close()
	}
	r := controllers.InitRouter()
	r.Use(cors.Default())
	_ = r.Run(config.Config.Port)
}
