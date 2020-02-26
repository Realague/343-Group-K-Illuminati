package error_handler

import (
	"github.com/gin-gonic/gin"
)

type ErrorMulti struct {
	Errors []*ErrorCustom
}

func (e *ErrorMulti) AddError(code int, message string) {
	e.Errors = append(e.Errors, &ErrorCustom{Code: code, Message: message})
}

type ErrorCustom struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

const (
	PayloadError                 = 1000
	UserDoestNotExist            = 2000
	UserSaveError                = 2001
	UserNewError                 = 2002
	AdminReservedAction          = 3000
	SystemErrorPassword          = 4000
	SystemErrorEmailNotConfirmed = 4000
	SystemErrorDecrypt           = 4001
	SystemErrorMail              = 4002
	DatabaseErrorUpdate          = 5000
	DatabaseErrorQuery           = 5002
	DatabaseErrorRemove          = 5001
)

func HandleBasicError(err error, c *gin.Context, code int) bool {
	if err != nil {
		c.JSON(code, gin.H{
			"error": err.Error(),
		})
		c.Done()
		return true
	}
	return false
}

func HandleCustomError(err error, errs ErrorMulti, c *gin.Context, code int) bool {
	if len(errs.Errors) > 0 {
		c.AbortWithStatusJSON(code, gin.H{
			"error":  err.Error(),
			"errors": errs.Errors,
		})
		c.Done()
		return true
	}
	return false
}
