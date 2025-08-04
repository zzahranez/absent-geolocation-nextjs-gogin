package controller

import (
	"gin/service"

	"github.com/gin-gonic/gin"
)

type PresenceController interface {
	GetPresenceByEmail(ctx *gin.Context)
}

type presencecontroller struct {
	service service.PresenceService
}

func NewPresenceController(service service.PresenceService) PresenceController {
	return &presencecontroller{service: service}
}

func (c *presencecontroller) GetPresenceByEmail(ctx *gin.Context) {
	userLogin := "jack@gmail.com"

	presence, err := c.service.GetPresenceByEmail(userLogin)
	if err != nil {
		ctx.JSON(400, gin.H{
			"status":  false,
			"error":   err.Error(),
			"message": "Terjadi kesalahan",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"status":  true,
		"message": "Data berhasil ditemukan",
		"data":    presence,
	})
}
