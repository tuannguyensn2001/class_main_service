package middlewares

import (
	"class_main_service/src/app"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
)

func Recover(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			ctx.Header("Content-Type", "application/json")

			if httpError, ok := err.(*app.HttpError); ok {
				ctx.AbortWithStatusJSON(httpError.StatusCode, httpError)
				return
			}

			if errors.Is(err.(error), gorm.ErrRecordNotFound) {
				ctx.AbortWithStatusJSON(http.StatusNotFound, app.NotFoundHttpError(err.(error).Error(), err.(error)))
				return
			}

			if validationError, ok := err.(validator.ValidationErrors); ok {
				ctx.AbortWithStatusJSON(422, app.HttpError{
					Message: validationError[0].Error(),
				})
				return
			}

			httpError := app.InternalHttpError("internal_service server", err.(error))
			ctx.AbortWithStatusJSON(httpError.StatusCode, httpError)
			return
		}
	}()

	ctx.Next()
}
