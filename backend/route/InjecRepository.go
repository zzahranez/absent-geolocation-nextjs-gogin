package route

import (
	"gin/controller"
	"gin/repository"
	"gin/service"

	"gorm.io/gorm"
)

func InjecRepository(db *gorm.DB) Routes {
	authRepo := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepo)
	authController := controller.NewAuthController(authService)

	return NewRoutes(authController)
}
