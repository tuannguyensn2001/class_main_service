package routes

import (
	"class_main_service/src/config"
	auth_repository "class_main_service/src/features/auth/repository"
	auth_transport "class_main_service/src/features/auth/transport"
	auth_usecase "class_main_service/src/features/auth/usecase"
	"class_main_service/src/services/internal_service"
	"github.com/gin-gonic/gin"
)

func Bootstrap(r *gin.Engine, cfg config.Config) {

	internalService := internal_service.New()

	authRepository := auth_repository.New(cfg.Db)
	authUsecase := auth_usecase.New(authRepository, internalService, cfg.SecretKey)
	authTransport := auth_transport.New(authUsecase)

	r.POST("/api/v1/auth/register", authTransport.Register)
	r.POST("/api/v1/auth/login", authTransport.Login)
}
