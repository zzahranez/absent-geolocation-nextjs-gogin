package controller

import (
	"gin/dto"
	"gin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
}

type authcontroller struct {
	service service.AuthService
}

func NewAuthController(se service.AuthService) AuthController {
	return &authcontroller{service: se}
}

func (c *authcontroller) Login(ctx *gin.Context) {
	var req dto.LoginRequest

	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "invalid input",
		})
		return
	}

	users, err := c.service.Login(&req)
	if err != nil {
		ctx.JSON(401, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	//Ini servicec
	ctx.JSON(200, gin.H{
		"status": true,
		"user":   users,
	})
}
