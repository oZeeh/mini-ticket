package main

import (
	"backend/config"
	"backend/users"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "backend/docs"
)

// @title Mini Ticket API
// @version 1.0
// @description Helpdesk API
// @host localhost:8080
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("no .env file found")
	}

	db := config.ConnectMongo()

	userRepository := users.NewRepository(db)
	userService := users.NewService(userRepository)
	userController := *users.NewController(userService)

	r := gin.Default()

	userController.RegisterRoutes(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
