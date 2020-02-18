package middleware

import (
	"343-Group-K-Illuminati/illuminati_api/config"
	"343-Group-K-Illuminati/illuminati_api/database"
	"343-Group-K-Illuminati/illuminati_api/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func IsTokenValidForAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Not authorized",
			})
			return
		}

		encryptedToken := strings.Replace(header, "Bearer", "", 1)

		if encryptedToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Not authorized",
			})
			return
		}

		encryptedToken = strings.Replace(encryptedToken, " ", "", 1)
		if encryptedToken != config.Config.TokenAuthenticationKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Not authorized",
			})
			return
		}
		c.Next()
	}
}

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {

		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Admin reserved action",
			})
			return
		}

		encryptedToken := strings.Replace(header, "Bearer", "", 1)

		if encryptedToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Not authorized",
			})
			return
		}

		encryptedToken = strings.Replace(encryptedToken, " ", "", 1)
		claims, valid, err := utils.DecryptToken(encryptedToken)
		if !valid {
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "Admin reserved action",
				})
			}
			return
		}

		id := claims["id"].(string)

		user, err := database.Database.FindUserByID(id)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Admin reserved action",
			})
			c.Abort()
			return
		}
		if user.Admin == false {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Admin reserved action",
			})
		}
		c.Set("user", user)
		c.Next()
	}
}

func IsAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {

		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Not authorized",
			})
			return
		}

		encryptedToken := strings.Replace(header, "Bearer", "", 1)

		if encryptedToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Not authorized",
			})
			return
		}

		encryptedToken = strings.Replace(encryptedToken, " ", "", 1)
		claims, valid, err := utils.DecryptToken(encryptedToken)
		if !valid || err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Not authorized",
			})
			return
		}

		id := claims["id"].(string)

		user, err := database.Database.FindUserByID(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Not authorized",
			})
			return
		}

		if user.Verified == false {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Not authorized",
			})
			return
		}
		c.Set("user", user)
		c.Next()
	}
}
