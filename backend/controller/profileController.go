package controller

import (
	"gin/service"

	"github.com/gin-gonic/gin"
)

type ProfileController interface {
	GetAccountUser(ctx *gin.Context)
}

type profilecontroller struct {
	service service.ProfileService
}

func NewProfileController(serv service.ProfileService) ProfileController {
	return &profilecontroller{
		service: serv,
	}
}

func (c profilecontroller) GetAccountUser(ctx *gin.Context) {

	userLogin := "jack@gmail.com"

	accountData, err := c.service.HandleGetAccountUser(userLogin)
	if err != nil {
		ctx.JSON(500, gin.H{
			"status":  false,
			"message": "gagal mendapatkan data",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  true,
		"message": "data account berhasil didapatkan",
		"data":    accountData,
	})
}
