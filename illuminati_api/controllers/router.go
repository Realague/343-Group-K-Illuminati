package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORS
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers, authorization, content-type")
		c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Content-Type", "application/json")
	}
}

// InitRouter
func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(CORS())

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.OPTIONS("/api/*id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	})

	api := r.Group("/api")

	auth := api.Group("/auth")

	feedback := api.Group("/feedback")

	InitFeedbackController(feedback)

	InitAuthenticationController(auth)

	user := api.Group("/users")

	InitUserController(user)

	return r
}
