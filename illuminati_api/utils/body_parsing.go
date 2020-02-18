package utils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"343-Group-K-Illuminati/illuminati_api/controllers/error_handler"
	"net/http"
)

func GetBodyAsStringMap(c *gin.Context) (map[string]string, bool) {
	var body map[string]string
	var errorsCustom error_handler.ErrorMulti

	byteArray, err := c.GetRawData()
	if err != nil {
		errorsCustom.AddError(error_handler.PayloadError, err.Error())
		error_handler.HandleCustomError(err, errorsCustom, c, http.StatusBadRequest)
		return nil, true
	}

	err = json.Unmarshal(byteArray, &body)
	if err != nil {
		errorsCustom.AddError(error_handler.PayloadError, err.Error())
		error_handler.HandleCustomError(err, errorsCustom, c, http.StatusBadRequest)
		return nil, true
	}
	return body, false
}
