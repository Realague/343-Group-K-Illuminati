package controllers

import (
	"343-Group-K-Illuminati/illuminati_api/config"
	"343-Group-K-Illuminati/illuminati_api/controllers/error_handler"
	"343-Group-K-Illuminati/illuminati_api/middleware"
	"343-Group-K-Illuminati/illuminati_api/models/payload"
	"343-Group-K-Illuminati/illuminati_api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func InitFeedbackController(route *gin.RouterGroup) {

	route.POST("", middleware.IsAuthenticated(), func(c *gin.Context) {
		feedbackPayload := payload.Feedback{}
		var errorsCustom error_handler.ErrorMulti

		if err := c.ShouldBindBodyWith(&feedbackPayload, binding.JSON); err != nil {
			errorsCustom.AddError(error_handler.PayloadError, err.Error())
			error_handler.HandleCustomError(err, errorsCustom, c, http.StatusBadRequest)
			return
		}

		if !config.Config.IsUnitTest {
			if err := utils.SendEmail(feedbackPayload.Title, feedbackPayload.Body, "bolmog.noreply@gmail.com"); err != nil {
				errorsCustom.AddError(error_handler.PayloadError, err.Error())
				error_handler.HandleCustomError(err, errorsCustom, c, http.StatusBadRequest)
				return
			}
		}

		c.Status(http.StatusNoContent)
	})

}
