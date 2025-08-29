package route

import (
	"gin/controller"

	"github.com/gin-gonic/gin"
)

type Routes interface {
	SetupRoutes(r *gin.Engine)
}

type routes struct {
	authController            controller.AuthController
	presenceHistoryController controller.PresenceHistoryController
	presencecontroller        controller.PresenceController
}

func NewRoutes(
	auth controller.AuthController,
	presencehistory controller.PresenceHistoryController,
	presence controller.PresenceController,
) Routes {
	return &routes{
		authController:            auth,
		presenceHistoryController: presencehistory,
		presencecontroller:        presence,
	}
}

func (r *routes) SetupRoutes(router *gin.Engine) {
	// Public route
	auth := router.Group("/auth")
	{
		auth.POST("/login", r.authController.Login)
	}

	// Create Presence
	presence := router.Group("/presence")
	{
		presence.POST("", r.presencecontroller.CreatePresence)
	}

	presenceHistory := router.Group("/presence/history")
	{
		presenceHistory.GET("", r.presenceHistoryController.GetPresenceByEmail)
	}
}
