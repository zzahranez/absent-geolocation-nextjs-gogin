package controller

import (
	"gin/database"
	"gin/dto"
	"gin/model"

	"github.com/gin-gonic/gin"
)

func GetDBUsers(ctx *gin.Context) {
	var user []model.User
	db := database.DB

	err := db.Find(&user).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	var responses []dto.ResponseUsers
	for _, u := range user {
		responses = append(responses, dto.ResponseUsers{
			Name:  u.Name,
			Email: u.Email,
			Role:  u.Role,
		})
	}

	ctx.JSON(200, gin.H{
		"status":  true,
		"message": "user berhasil didapatkan",
		"data":    responses,
	})
}
