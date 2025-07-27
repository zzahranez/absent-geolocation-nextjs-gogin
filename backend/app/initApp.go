package app

import (
	"gin/database"
	"gin/route"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}
}

func InitApp() {
	//Load Environtment
	LoadEnv()
	//Database Connection
	db := database.ConnectionToDataBase()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	//injecRepository And Route SetUp
	routes := route.InjecRepository(db)
	routes.SetupRoutes(r)

	r.Run()

}
