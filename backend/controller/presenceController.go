package controller

import (
	"gin/dto"
	"gin/service"

	"github.com/gin-gonic/gin"
)

type PresenceController interface {
	CreatePresence(ctx *gin.Context)
}

type presencecontroller struct {
	service service.PresenceService
}

func NewPresenceController(serv service.PresenceService) PresenceController {
	return &presencecontroller{
		service: serv,
	}
}

func (c *presencecontroller) CreatePresence(ctx *gin.Context) {
	userLogin := "jack@gmail.com"

	var req dto.PresenceRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(400, gin.H{
			"status":  false,
			"message": "request data is nothing happen",
			"error":   err.Error(),
		})

		return
	}

	presence, err := c.service.HandleCreatePresence(userLogin, req)
	if err != nil {
		ctx.JSON(500, gin.H{
			"status":  false,
			"message": "can't create presence",
			"error":   err.Error(),
		})

		return
	}

	ctx.JSON(200, gin.H{
		"status":  true,
		"message": "presence has been create successfully",
		"data":    presence,
	})
}
