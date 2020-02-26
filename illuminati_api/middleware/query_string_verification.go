package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"343-Group-K-Illuminati/illuminati_api/controllers/error_handler"
	"343-Group-K-Illuminati/illuminati_api/models/filters"
	"net/http"
)

func UsersQueryStringVerification() gin.HandlerFunc {
	return func(c *gin.Context) {
		variables := []string{"id", "createdAt", "updatedAt", "username", "email", "admin", "verified", "mmr", "friend_list"}
		if QueryStringVerification(variables, c.Keys["research_data"].(filters.ResearchData).QueryStringToVerify.Parameters, c) {
			c.Next()
		}
	}
}

func QueryStringVerification(variables []string, parameters []filters.Parameter, c *gin.Context) bool {
	var errorsCustom error_handler.ErrorMulti

	for _, param := range parameters {
		for i, variable := range variables {
			if variable == param.Name {
				break
			} else if len(variables) - 1 == i {
				err := errors.New("invalid querystring parameter " + param.Name)
				errorsCustom.AddError(error_handler.PayloadError, err.Error())
				error_handler.HandleCustomError(err, errorsCustom, c, http.StatusBadRequest)
				return false
			}
		}
	}
	return true
}


