package controller

import (
	"gin/dto"
	"gin/service"

	"github.com/gin-gonic/gin"
)

type CutiController interface {
	CreateCuti(ctx *gin.Context)
}

type cuticontroller struct {
	serv service.CutiService
}

func NewCutiController(service service.CutiService) CutiController {
	return &cuticontroller{
		serv: service,
	}
}

func (c *cuticontroller) CreateCuti(ctx *gin.Context) {
	userLogin := "jack@gmail.com"

	var req dto.CreateCutiRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(400, gin.H{
			"status":  false,
			"message": "data request is nothing happen",
			"error":   err.Error(),
		})

		return
	}

	createCuti, err := c.serv.HandleCreateCuti(userLogin, req)
	if err != nil {
		ctx.JSON(500, gin.H{
			"status":  false,
			"message": "can't create cuti submission",
			"err":     err.Error(),
		})

		return
	}

	ctx.JSON(200, gin.H{
		"status":  true,
		"message": "cuti succesfully created",
		"data":    createCuti,
	})

}
