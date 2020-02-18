package utils

import (
	"343-Group-K-Illuminati/illuminati_api/config"
	"343-Group-K-Illuminati/illuminati_api/controllers/error_handler"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/dgrijalva/jwt-go.v3"
)

func DecryptToken(token string) (jwt.MapClaims, bool, error) {
	if token == "" {
		return nil, false, errors.New("nill token")
	}
	decryptToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.Config.TokenEncryptionKey), nil
	})

	if err != nil {
		return nil, false, err
	}

	claims, valid := decryptToken.Claims.(jwt.MapClaims)
	if !decryptToken.Valid {
		valid = false
	}
	return claims, valid, err
}

func GenerateToken(c *gin.Context, claims jwt.Claims) (string, bool) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.Config.TokenEncryptionKey))

	if error_handler.HandleBasicError(err, c, http.StatusInternalServerError) {
		return token, true
	}
	return token, false
}

func GenerateAccessToken(c *gin.Context, id string) (string, bool) {
	return GenerateToken(c, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(config.Config.AccessTokenValidityTime * time.Hour).Unix(),
	})
}

func GenerateRefreshToken(c *gin.Context, id string) (string, bool) {
	return GenerateToken(c, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(config.Config.RefreshTokenValidityTime * time.Hour).Unix(),
	})
}
