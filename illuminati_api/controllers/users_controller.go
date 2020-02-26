package controllers

import (
	"343-Group-K-Illuminati/illuminati_api/config"
	"343-Group-K-Illuminati/illuminati_api/controllers/error_handler"
	"343-Group-K-Illuminati/illuminati_api/database"
	"343-Group-K-Illuminati/illuminati_api/middleware"
	"343-Group-K-Illuminati/illuminati_api/models/db"
	"343-Group-K-Illuminati/illuminati_api/models/filters"
	"343-Group-K-Illuminati/illuminati_api/models/payload"
	"343-Group-K-Illuminati/illuminati_api/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func InitUserController(route *gin.RouterGroup) {
	route.Use()
	{
		route.GET("",  middleware.IsAuthenticated(), middleware.Research(), middleware.UsersQueryStringVerification(), func(c *gin.Context) {
			var errorsCustom error_handler.ErrorMulti
			research := c.Keys["research_data"].(filters.ResearchData)

			users, err := database.Database.FindUsersByQuery(research)
			if err != nil {
				errorsCustom.AddError(error_handler.UserDoestNotExist, err.Error())
				error_handler.HandleCustomError(err, errorsCustom, c, http.StatusInternalServerError)
				return
			}
			if config.Config.IsUnitTest == true {
				c.JSON(http.StatusOK, users)
			} else {
				c.JSON(http.StatusOK, gin.H{
					"users": users,
				})
			}
		})

		route.GET("/:id", middleware.IsAuthenticated(), func(c *gin.Context) {
			id := c.Param("id")
			var errorsCustom error_handler.ErrorMulti

			if id == "me" {
				id = c.Keys["user"].(db.User).Id.Hex()
			} else if !c.Keys["user"].(db.User).Admin {
				errorsCustom.AddError(error_handler.AdminReservedAction, "Admin reserved action")
				error_handler.HandleCustomError(errors.New("admin reserved action"), errorsCustom, c, http.StatusUnauthorized)
				return
			}

			user, err := database.Database.FindUserByID(id)
			if err != nil {
				errorsCustom.AddError(error_handler.UserDoestNotExist, err.Error())
				error_handler.HandleCustomError(err, errorsCustom, c, http.StatusNotFound)
				return
			}

			c.JSON(http.StatusOK, user)
		})

		route.POST("", middleware.IsAdmin(), func(c *gin.Context) {
			registrationPayload := payload.LocalRegistration{}
			var errorsCustom error_handler.ErrorMulti

			if err := c.ShouldBind(&registrationPayload); err == nil {
				register(c, registrationPayload)
			} else {
				errorsCustom.AddError(error_handler.PayloadError, err.Error())
				error_handler.HandleCustomError(err, errorsCustom, c, http.StatusBadRequest)
			}
		})

		route.PUT("/:id", middleware.IsAuthenticated(), func(c *gin.Context) {
			updatePayload := &db.UpdateUser{}
			var errorsCustom error_handler.ErrorMulti

			id := c.Param("id")

			if id == "me" {
				id = c.Keys["user"].(db.User).Id.Hex()
			} else if !c.Keys["user"].(db.User).Admin {
				errorsCustom.AddError(error_handler.AdminReservedAction, "Admin reserved action")
				error_handler.HandleCustomError(errors.New("admin reserved action"), errorsCustom, c, http.StatusUnauthorized)
				return
			}

			err := c.ShouldBind(updatePayload)
			if err != nil {
				errorsCustom.AddError(error_handler.PayloadError, err.Error())
				error_handler.HandleCustomError(err, errorsCustom, c, http.StatusBadRequest)
				return
			}

			user, err := database.Database.FindUserByID(id)
			if err != nil {
				errorsCustom.AddError(error_handler.UserDoestNotExist, err.Error())
				error_handler.HandleCustomError(err, errorsCustom, c, http.StatusNotFound)
				return
			}

			if updatePayload.Password != "" {
				password, err := utils.HashPassword(updatePayload.Password)
				if err != nil {
					errorsCustom.AddError(error_handler.SystemErrorPassword, err.Error())
					error_handler.HandleCustomError(err, errorsCustom, c, http.StatusInternalServerError)
					return
				}
				user.Password = password
			}
			user.UpdatedAt = time.Now()
			user.Username = updatePayload.Username
			user.Admin = updatePayload.Admin
			user.Email = updatePayload.Email
			user.Verified = updatePayload.Verified
			user.FriendList = updatePayload.FriendList
			user.Mmr = updatePayload.Mmr

			err = database.Database.UpdateUser(id, &user)
			if err != nil {
				errorsCustom.AddError(error_handler.DatabaseErrorUpdate, err.Error())
				error_handler.HandleCustomError(err, errorsCustom, c, http.StatusBadRequest)
				return
			}

			c.JSON(http.StatusOK, user)
		})

		route.DELETE("/:id", middleware.IsAuthenticated(), func(c *gin.Context) {
			id := c.Param("id")
			var errorsCustom error_handler.ErrorMulti

			if id == "me" {
				id = c.Keys["user"].(db.User).Id.Hex()
			} else if !c.Keys["user"].(db.User).Admin {
				errorsCustom.AddError(error_handler.AdminReservedAction, "Admin reserved action")
				error_handler.HandleCustomError(errors.New("admin reserved action"), errorsCustom, c, http.StatusUnauthorized)
				return
			}

			err := database.Database.DeleteUserByID(id)
			if err != nil {
				errorsCustom.AddError(error_handler.DatabaseErrorRemove, err.Error())
				error_handler.HandleCustomError(err, errorsCustom, c, http.StatusNotFound)
				return
			}

			c.Status(http.StatusNoContent)
		})

	}

}
