package controllers

import (
	"343-Group-K-Illuminati/illuminati_api/config"
	"343-Group-K-Illuminati/illuminati_api/controllers/error_handler"
	"343-Group-K-Illuminati/illuminati_api/database"
	"343-Group-K-Illuminati/illuminati_api/middleware"
	"343-Group-K-Illuminati/illuminati_api/models/db"
	"343-Group-K-Illuminati/illuminati_api/models/payload"
	"343-Group-K-Illuminati/illuminati_api/utils"
	"errors"
	"gopkg.in/dgrijalva/jwt-go.v3"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/mgo.v2/bson"
)

func InitAuthenticationController(route *gin.RouterGroup) {

	route.POST("/local", middleware.IsTokenValidForAuthentication(), func(c *gin.Context) {
		loginPayload := payload.LocalLogin{}
		registrationPayload := payload.LocalRegistration{}
		var errorsCustom error_handler.ErrorMulti

		if err := c.ShouldBindBodyWith(&loginPayload, binding.JSON); err == nil {
			login(c, loginPayload)
		} else if err := c.ShouldBindBodyWith(&registrationPayload, binding.JSON); err == nil {
			register(c, registrationPayload)
		} else {
			errorsCustom.AddError(error_handler.PayloadError, err.Error())
			error_handler.HandleCustomError(err, errorsCustom, c, http.StatusBadRequest)
		}
	})

	route.POST("/refresh-token", middleware.IsTokenValidForAuthentication(), func(c *gin.Context) {
		body, failed := utils.GetBodyAsStringMap(c)
		var errorsCustom error_handler.ErrorMulti

		if failed {
			return
		}

		claims, valid, err := utils.DecryptToken(body["refresh_token"])
		if !valid {
			if err != nil {
				errorsCustom.AddError(error_handler.UserSaveError, "Invalid token supplied")
				error_handler.HandleCustomError(errors.New("invalid token supplied"), errorsCustom, c, http.StatusBadRequest)
			}
			return
		}
		id := claims["id"].(string)
		accessToken, failed := utils.GenerateAccessToken(c, id)
		refreshToken, failed := utils.GenerateRefreshToken(c, id)

		if failed {
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"refresh_token": refreshToken,
			"access_token":  accessToken,
		})
	})

	route.GET("/confirm-email/:token", func(c *gin.Context) {
		token := c.Param("token")
		var errorsCustom error_handler.ErrorMulti

		claims, valid, err := utils.DecryptToken(token)

		if !valid {
			if err != nil {
				errorsCustom.AddError(error_handler.SystemErrorDecrypt, err.Error())
				error_handler.HandleCustomError(err, errorsCustom, c, http.StatusBadRequest)
				return
			}
		}

		id := claims["id"]

		user, err := database.Database.FindUserByID(id.(string))
		if err != nil {
			errorsCustom.AddError(error_handler.UserDoestNotExist, err.Error())
			error_handler.HandleCustomError(err, errorsCustom, c, http.StatusNotFound)
			return
		}

		user.Verified = true

		err = database.Database.UpdateUser(user.Id.Hex(), &user)
		if err != nil {
			errorsCustom.AddError(error_handler.DatabaseErrorUpdate, err.Error())
			error_handler.HandleCustomError(err, errorsCustom, c, http.StatusInternalServerError)
			return
		}

		c.Redirect(http.StatusPermanentRedirect, config.Config.LoginRedirectionUrl)
	})

	route.POST("/recover-password", middleware.IsTokenValidForAuthentication(), func(c *gin.Context) {
		var errorsCustom error_handler.ErrorMulti

		body, failed := utils.GetBodyAsStringMap(c)
		if failed {
			return
		}

		email := body["email"]
		if email == "" {
			errorsCustom.AddError(error_handler.PayloadError, "Bad request")
			error_handler.HandleCustomError(errors.New("bad request"), errorsCustom, c, http.StatusBadRequest)
			return
		}

		user, err := database.Database.FindUserByKey("email", email)
		if err != nil {
			errorsCustom.AddError(error_handler.UserDoestNotExist, err.Error())
			error_handler.HandleCustomError(err, errorsCustom, c, http.StatusNotFound)
			return
		}

		if !config.Config.IsUnitTest && utils.SendRecoverPasswordEmail(c, *user) != nil {
			return
		}

		c.Status(http.StatusNoContent)
	})

	route.POST("/send-confirmation-mail", middleware.IsTokenValidForAuthentication(), func(c *gin.Context) {
		var errorsCustom error_handler.ErrorMulti

		body, failed := utils.GetBodyAsStringMap(c)
		if failed {
			return
		}

		email := body["email"]
		if email == "" {
			errorsCustom.AddError(error_handler.PayloadError, "Bad request")
			error_handler.HandleCustomError(errors.New("bad request"), errorsCustom, c, http.StatusBadRequest)
			return
		}

		user, err := database.Database.FindUserByKey("email", email)
		if err != nil {
			errorsCustom.AddError(error_handler.DatabaseErrorQuery, err.Error())
			error_handler.HandleCustomError(err, errorsCustom, c, http.StatusNotFound)
			return
		}

		if !config.Config.IsUnitTest {
			err = utils.SendConfirmationEmail(c, *user)
			if err != nil {
				errorsCustom.AddError(error_handler.SystemErrorMail, err.Error())
				error_handler.HandleCustomError(err, errorsCustom, c, http.StatusInternalServerError)
				return
			}
		}

		c.Status(http.StatusNoContent)
	})
}

func login(c *gin.Context, loginPayload payload.LocalLogin) {
	user := &db.User{}
	var err error
	var errorsCustom error_handler.ErrorMulti

	param := ""
	if strings.Contains(loginPayload.Identifier, "@") {
		param = "email"
	} else {
		param = "username"
	}

	user, err = database.Database.FindUserByKey(param, loginPayload.Identifier)
	if err != nil {
		errorsCustom.AddError(error_handler.UserDoestNotExist, err.Error())
		error_handler.HandleCustomError(err, errorsCustom, c, http.StatusBadRequest)
		return
	}

	if !utils.ComparePassword(user.Password, loginPayload.Password) {
		errorsCustom.AddError(error_handler.SystemErrorPassword, "Invalid email/password supplied")
		error_handler.HandleCustomError(errors.New("invalid email/password supplied"), errorsCustom, c, http.StatusBadRequest)
		return
	}

	if !user.Verified {
		errorsCustom.AddError(error_handler.SystemErrorEmailNotConfirmed, "Email not confirmed")
		error_handler.HandleCustomError(errors.New("email not confirmed"), errorsCustom, c, http.StatusUnauthorized)
		return
	}

	accessToken, failed := utils.GenerateAccessToken(c, user.Id.Hex())
	refreshToken, failed := utils.GenerateRefreshToken(c, user.Id.Hex())
	if failed {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"user":          user,
	})
}

func register(c *gin.Context, registrationPayload payload.LocalRegistration) {
	var errorsCustom error_handler.ErrorMulti

	password, err := utils.HashPassword(registrationPayload.Password)
	if err != nil {
		errorsCustom.AddError(error_handler.SystemErrorDecrypt, err.Error())
		error_handler.HandleCustomError(err, errorsCustom, c, http.StatusInternalServerError)
		return
	}

	user := &db.User{
		Id:        bson.NewObjectId(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Username:  registrationPayload.Username,
		Password:  password,
		Email:     registrationPayload.Email,
	}

	nb, err := database.Database.CountUserByKey("email", user.Email)
	nbUsername, err := database.Database.CountUserByKey("username", user.Username)
	forbiddenMailAddresses, err := database.Database.FindAllForbiddenMailAddress()

	if valid, issues := user.Validate(nb, nbUsername, forbiddenMailAddresses, err); !valid {
		errorsCustom.AddError(error_handler.PayloadError, issues[0].Error())
		error_handler.HandleCustomError(issues[0], errorsCustom, c, http.StatusConflict)
		return
	}
	if err := database.Database.InsertUser(user); err != nil {
		errorsCustom.AddError(error_handler.UserSaveError, err.Error())
		error_handler.HandleCustomError(err, errorsCustom, c, http.StatusBadRequest)
		return
	}

	if config.Config.IsUnitTest {
		token, err := utils.GenerateToken(c, jwt.MapClaims{
			"id": user.Id.Hex(),
		})

		if err {
			errorsCustom.AddError(error_handler.SystemErrorDecrypt, "Internal server error")
			error_handler.HandleCustomError(errors.New("Internal server error"), errorsCustom, c, http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"token": token,
			"user":  user,
		})
	} else {
		err = utils.SendConfirmationEmail(c, *user)
		if err != nil {
			errorsCustom.AddError(error_handler.SystemErrorMail, err.Error())
			error_handler.HandleCustomError(err, errorsCustom, c, http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusCreated, user)
	}

}
