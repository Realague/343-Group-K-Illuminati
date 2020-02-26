package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"343-Group-K-Illuminati/illuminati_api/controllers/error_handler"
	"343-Group-K-Illuminati/illuminati_api/models/filters"
	"net/http"
	"strconv"
	"strings"
)

func pagination(c *gin.Context, data filters.ResearchData) filters.ResearchData {
	var errorsCustom error_handler.ErrorMulti
	var err error
	var offset int64 = 0
	param := c.Query("offset")
	if param != "" {
		offset, err = strconv.ParseInt(param, 10, 64)
		if err != nil {
			errorsCustom.AddError(error_handler.PayloadError, err.Error())
			error_handler.HandleCustomError(err, errorsCustom, c, http.StatusBadRequest)
			return data
		}
	}

	var limit int64 = 0
	param = c.Query("limit")
	if param != "" {
		limit, err = strconv.ParseInt(c.Query("limit"), 10, 64)
		if err != nil {
			errorsCustom.AddError(error_handler.PayloadError, err.Error())
			error_handler.HandleCustomError(err, errorsCustom, c, http.StatusBadRequest)
			return data
		}
	}

	data.Pagination = filters.PaginationData{
		Offset: int(offset),
		Limit:  int(limit),
	}
	return data
}

func sort(c *gin.Context, data filters.ResearchData) filters.ResearchData {
	var errorsCustom error_handler.ErrorMulti
	param := c.Query("sort_by")
	if param == "" {
		return data
	}
	sortBy := strings.Split(param, ",")

	for _, arg := range sortBy {
		if arg[0] != '-' {
			split := strings.Split(arg, "+")
			if len(split) != 2 {
				errorsCustom.AddError(error_handler.PayloadError, "sort by error")
				error_handler.HandleCustomError(errors.New("sort by error"), errorsCustom, c, http.StatusBadRequest)
				return data
			} else {
				arg = split[0]
			}
		}
		data.Sorting = append(data.Sorting, arg)
	}

	return data
}

func query(c *gin.Context, data filters.ResearchData) filters.ResearchData {
	queryString := c.Query("q")
	data.QueryStringToVerify = parseQueryString(c, queryString)
	data.QueryString = QueryStringToBsonMap(data.QueryStringToVerify)

	return data
}

func Research() gin.HandlerFunc {
	return func(c *gin.Context) {
		data := filters.ResearchData{}
		data = pagination(c, data)
		data = sort(c, data)
		data = query(c, data)

		c.Set("research_data", data)
		c.Next()
	}
}
