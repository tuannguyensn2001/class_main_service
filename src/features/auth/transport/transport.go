package auth_transport

import (
	"class_main_service/src/app"
	auth_struct "class_main_service/src/features/auth/struct"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type IUsecase interface {
	Register(ctx context.Context, input auth_struct.RegisterInput) error
}

type transport struct {
	usecase IUsecase
}

func New(usecase IUsecase) *transport {
	return &transport{usecase: usecase}
}

func (t *transport) Register(ctx *gin.Context) {
	var input auth_struct.RegisterInput
	if err := ctx.ShouldBind(&input); err != nil {
		zap.S().Error(err)
		panic(app.BadRequestHttpError("data not valid", err))
	}

	err := t.usecase.Register(ctx.Request.Context(), input)
	if err != nil {

		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": "success",
	})

}
