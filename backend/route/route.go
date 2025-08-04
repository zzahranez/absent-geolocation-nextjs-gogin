package route

import (
	"gin/controller"

	"github.com/gin-gonic/gin"
)

type Routes interface {
	SetupRoutes(r *gin.Engine)
}

type routes struct {
	authController     controller.AuthController
	presenceController controller.PresenceController
}

func NewRoutes(
	auth controller.AuthController,
	presence controller.PresenceController,
) Routes {
	return &routes{
		authController:     auth,
		presenceController: presence,
	}
}

func (r *routes) SetupRoutes(router *gin.Engine) {
	// Public route
	auth := router.Group("/auth")
	{
		auth.POST("/login", r.authController.Login)
	}

	presence := router.Group("/presence")
	{
		presence.GET("", r.presenceController.GetPresenceByEmail)
	}
}
