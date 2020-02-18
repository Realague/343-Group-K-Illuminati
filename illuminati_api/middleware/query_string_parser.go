package middleware

import (
	"343-Group-K-Illuminati/illuminati_api/controllers/error_handler"
	"343-Group-K-Illuminati/illuminati_api/models/filters"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

func parseParameter(c *gin.Context, queryStringObject filters.QueryString, parameterString string) filters.QueryString {
	param := strings.SplitN(parameterString, "=", 2)
	var err error
	var errorsCustom error_handler.ErrorMulti

	if len(param) != 2 {
		err = errors.New("bad query string parameter")
		errorsCustom.AddError(error_handler.PayloadError, err.Error())
		error_handler.HandleCustomError(err, errorsCustom, c, http.StatusBadRequest)
		return queryStringObject
	}

	var attributes []filters.Attribute
	var attribute filters.Attribute
	if param[1] == "" {
		attribute.Key = "$exists"
		attribute.Value = true
	} else if param[1][0] == '!' {
		if len(param[1]) == 1 {
			attribute.Key = "$exists"
			attribute.Value = false
		} else {
			attribute.Key = "$ne"
			attribute.Value = strings.Split(param[1], "!")[1]
		}
	} else if param[1][0] == '>' {
		if param[1][1] == '=' {
			attribute.Key = "$gte"
			attribute.Value, err = strconv.ParseInt(strings.Split(param[1], ">=")[1], 10, 64)
		} else {
			attribute.Key = "$gt"
			attribute.Value, err = strconv.ParseInt(strings.Split(param[1], ">")[1], 10, 64)
		}
	} else if param[1][0] == '<' {
		if param[1][1] == '=' {
			attribute.Key = "$lte"
			attribute.Value, err = strconv.ParseInt(strings.Split(param[1], "<=")[1], 10, 64)
		} else {
			attribute.Key = "$lt"
			attribute.Value, err = strconv.ParseInt(strings.Split(param[1], "<")[1], 10, 64)
		}
	} else if param[1][0] == '/' {
		attribute.Key = "$regex"
		attribute.Value = strings.Split(param[1], "/")[1]
		attributes = append(attributes, attribute)
		attribute.Key = "$options"
		attribute.Value = "i"
	} else if param[0][len(param)] == ']' {
		if param[1][0] == '!' {
			attribute.Key = "$nin"
			param[1] = strings.Split(param[1], "!")[1]
			attribute.Value = strings.Split(param[1], "&")
		} else {
			attribute.Key = "$in"
			attribute.Value = strings.Split(param[1], "&")
		}
	} else {
		attribute.Key = ""
		attribute.Value = param[1]
	}

	if err != nil {
		errorsCustom.AddError(error_handler.PayloadError, err.Error())
		error_handler.HandleCustomError(err, errorsCustom, c, http.StatusBadRequest)
		return queryStringObject
	}

	attributes = append(attributes, attribute)
	for i, parameter := range queryStringObject.Parameters {
		if parameter.Name == param[0] {
			queryStringObject.Parameters[i].Attributes = append(parameter.Attributes, attributes...)
			return queryStringObject
		}
	}

	queryStringObject.Parameters = append(queryStringObject.Parameters, filters.Parameter{attributes, param[0]})
	return queryStringObject
}

func QueryStringToBsonMap(queryString filters.QueryString) bson.M {
	result := bson.M{}

	for _, parameter := range queryString.Parameters {
		param := bson.M{}
		value := ""
		for _, attribute := range parameter.Attributes {
			if attribute.Key != "" {
				param[attribute.Key] = attribute.Value
			} else {
				value = attribute.Value.(string)
				break
			}
		}

		if value != "" {
			result[parameter.Name] = value
		} else {
			result[parameter.Name] = param
		}
	}
	return result
}

func parseQueryString(c *gin.Context, query string) filters.QueryString {
	queryString := filters.QueryString{}
	if query == "" {
		return queryString
	}

	arrayString := strings.Split(query, ",")

	for _, parameter := range arrayString {
		queryString = parseParameter(c, queryString, parameter)
	}

	return queryString
}
