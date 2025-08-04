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

	//Presence Page
	presenceRepo := repository.NewPresenceRepository(db)
	presenceService := service.NewPresenceService(presenceRepo)
	presenceController := controller.NewPresenceController(presenceService)

	return NewRoutes(authController, presenceController)
}
