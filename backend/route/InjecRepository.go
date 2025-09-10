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

	// Cuti
	cutiRepo := repository.NewCutiRepository(db)
	cuttService := service.NewCutiService(cutiRepo, usersRepo)
	cutiController := controller.NewCutiController(cuttService)

	// profile
	profileRepo := repository.NewProfileRepository(db)
	profileService := service.NewProfileService(profileRepo, usersRepo)
	profileController := controller.NewProfileController(profileService)

	return NewRoutes(authController, presenceHistoryController, presenceController, profileController, cutiController)
}
