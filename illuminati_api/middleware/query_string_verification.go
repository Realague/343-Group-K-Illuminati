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
		variables := []string{"id", "createdAt", "updatedAt", "auth_method", "username", "email", "admin", "verified", "gardening_info", "gardening_info.time_to_spend",
			"gardening_info.gardening_level", "gardening_info.latitude", "gardening_info.longitude"}
		if QueryStringVerification(variables, c.Keys["research_data"].(filters.ResearchData).QueryStringToVerify.Parameters, c) {
			c.Next()
		}
	}
}

func PlantsQueryStringVerification() gin.HandlerFunc {
	return func(c *gin.Context) {
		variables := []string{"id", "createdAt", "updatedAt", "common_name", "scientific_name", "temperature", "humidity", "ph_tolerance", "sunshine", "climate_types", "soils"}
		if QueryStringVerification(variables, c.Keys["research_data"].(filters.ResearchData).QueryStringToVerify.Parameters, c) {
			c.Next()
		}
	}
}

func GardensQueryStringVerification() gin.HandlerFunc {
	return func(c *gin.Context) {
		variables := []string{"id", "createdAt", "updatedAt", "garden", "garden.cases", "garden.cases.plant_id", "garden.cases.shading", "position", "position.lat", "position.lng", "size", "size.height", "size.width", "user_id"}
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


