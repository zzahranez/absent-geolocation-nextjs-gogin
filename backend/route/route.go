package route

import (
	"gin/controller"

	"github.com/gin-gonic/gin"
)

type Routes interface {
	SetupRoutes(r *gin.Engine)
}

type routes struct {
	authController controller.AuthController
}

func NewRoutes(auth controller.AuthController) Routes {
	return &routes{
		authController: auth,
	}
}

func (r *routes) SetupRoutes(router *gin.Engine) {
	// Public route
	auth := router.Group("/auth")
	{
		auth.POST("/login", r.authController.Login)
	}

	// Example: another module route group
	// user := router.Group("/users")
	// {
	// 	user.GET("/", r.userController.GetAllUsers)
	// }
}
