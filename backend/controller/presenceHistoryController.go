package controller

import (
	"gin/service"

	"github.com/gin-gonic/gin"
)

type PresenceHistoryController interface {
	GetPresenceByEmail(ctx *gin.Context)
}

type presencehistorycontroller struct {
	service service.PresenceHistoryService
}

func NewPresenceHistoryController(service service.PresenceHistoryService) PresenceHistoryController {
	return &presencehistorycontroller{service: service}
}

func (c *presencehistorycontroller) GetPresenceByEmail(ctx *gin.Context) {
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
