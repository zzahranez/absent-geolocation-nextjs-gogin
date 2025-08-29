package route

import (
	"gin/controller"
	"gin/repository"
	"gin/service"

	"gorm.io/gorm"
)

func InjecRepository(db *gorm.DB) Routes {
	// Users
	usersRepo := repository.NewUsersRepository(db)

	authRepo := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepo)
	authController := controller.NewAuthController(authService)

	//Presence History
	presenceHistoryRepo := repository.NewPresenceHistoryRepository(db)
	presenceHistoryService := service.NewPresenceHistoryService(presenceHistoryRepo)
	presenceHistoryController := controller.NewPresenceHistoryController(presenceHistoryService)

	// Presence
	presenceRepo := repository.NewPresenceRepository(db)
	presenceSerice := service.NewPresenceService(presenceRepo, usersRepo)
	presenceController := controller.NewPresenceController(presenceSerice)

	return NewRoutes(authController, presenceHistoryController, presenceController)
}
